package merlin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-resty/resty/v2"
	logger "github.com/ipfs/go-log"
	"math/big"
	"math/rand"
	"spike-mc-ops/service/merlin/contract"
	"spike-mc-ops/util"
	"strings"
	"sync"
	"testing"
	"time"
)

var (
	walletMnemonic            = ""
	quoterContractAddress     = "0x2569bcE69287618e2cd004f785d016F7DF29232F"
	merlinSwapContractAddress = "0x1aFa5D7f89743219576Ef48a9826261bE6378a68" //
	wbtcTokenContractAddress  = "0xF6D226f9Dc15d9bB51182815b320D3fBE324e1bA"
	merlContractAddress       = "0x5c46bFF4B38dc1EAE09C5BAc65872a1D8bc87378"
	swapContractAddress       = "0x2426191006F378BF33445f87938d355096eE2e8C" //池子合约
	merlinMainNetRpcAddress   = "https://merlin-mainnet-rpc.merlinchain.io"
	SlippageTolerance         = 0.05
	SwapMerlAmount            = "5"
)

type TxOps struct {
	Ops   *bind.TransactOpts
	Index int64
}

var lock sync.Mutex
var merlinClient *ethclient.Client

func TestSwap(t *testing.T) {
	logger.SetAllLoggers(logger.LevelInfo)
	client, err := ethclient.Dial(merlinMainNetRpcAddress)
	if err != nil {
		log.Errorf("failed to connect to merlin-mainnet-rpc.merlinchain.io")
		return
	}
	merlinClient = client
	txOps, err := InitWallet(0, 1)
	if err != nil {
		log.Errorf("failed to init wallet")
		return
	}
	throttle := make(chan struct{}, 5)
	var wg sync.WaitGroup
	for _, ops := range txOps {
		wg.Add(1)
		throttle <- struct{}{}
		go func(opts *bind.TransactOpts, index int64) {
			defer func() {
				wg.Done()
				<-throttle
			}()
			wbtcContract, err := contract.NewErc20(common.HexToAddress(wbtcTokenContractAddress), merlinClient)
			if err != nil {
				log.Errorf("err: %v", err)
				return
			}
			wbtcWalletBalance, err := wbtcContract.BalanceOf(nil, opts.From)
			if err != nil {
				log.Errorf("err: %v", err)
				return
			}
			if wbtcWalletBalance.Int64() > 0 {
				log.Infof("address: %s,  index: %d, wbtcWalletBalance: %s", opts.From, index, wbtcWalletBalance.String())
				return
			}
			txHash, err := BuyToken(opts)
			if err != nil {
				log.Errorf("exec %d, err: %v", index, err)
				return
			}
			log.Infof("exec %d , tx hash: %s success", index, txHash)
		}(ops.Ops, ops.Index)
		time.Sleep(3 * time.Second)
	}
	wg.Wait()
}

func InitWallet(walletStartIndex int, walletEndIndex int) ([]TxOps, error) {
	var wg sync.WaitGroup
	id, err := merlinClient.ChainID(context.Background())
	if err != nil {
		log.Errorf("query chain id err: %v", err)
		return nil, err
	}
	log.Infof("id: %d", id.Int64())
	txOps := make([]TxOps, 0)
	throttle := make(chan struct{}, 5)
	for i := walletStartIndex; i < walletEndIndex; i++ {
		wg.Add(1)
		throttle <- struct{}{}
		go func(index int) {
			defer func() {
				wg.Done()
				<-throttle
			}()
			priKeyString, err := util.DerivePrivateKeyWithNumber(walletMnemonic, index)
			if err != nil {
				return
			}
			address, err := util.GenerateAddress(priKeyString)
			log.Debugf("address: %s, privateKey: %s", address, priKeyString)
			priKey, err := crypto.HexToECDSA(priKeyString[2:])
			if err != nil {
				return
			}

			opts, err := bind.NewKeyedTransactorWithChainID(priKey, id)
			if err != nil {
				return
			}
			lock.Lock()
			txOps = append(txOps, TxOps{opts, int64(index)})
			lock.Unlock()

			util.CheckAllowance(priKeyString, merlContractAddress, merlinSwapContractAddress, id, merlinClient)
			time.Sleep(5 * time.Second)
			//util.CheckAllowance(priKeyString, wbtcTokenContractAddress, merlinSwapContractAddress, id, merlinClient)
		}(i)
	}
	wg.Wait()
	return txOps, nil
}

func QueryTokenHold(poolAddress string) (tokenHold map[string]string, err error) {
	const tokenHoldings = "https://api.w3w.ai/merlin/v2/explorer/address/%s/token_holdings"

	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").Get(fmt.Sprintf(tokenHoldings, poolAddress))

	if err != nil {
		fmt.Printf("err ------: %v", err)
		return
	}
	var tokenHoldInfo TokenHoldInfo
	err = json.Unmarshal(resp.Body(), &tokenHoldInfo)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	tokenHoldInfoMap := make(map[string]string)
	for _, info := range tokenHoldInfo.Data {
		tokenHoldInfoMap[info.TokenAddress] = info.Balance
	}
	log.Infof("tokenHoldInfo: %v", tokenHoldInfoMap)
	return tokenHoldInfoMap, nil
}

func BuyToken(opts *bind.TransactOpts) (txHash string, err error) {
	tokenHoldMap, err := QueryTokenHold(swapContractAddress)
	if err != nil {
		log.Errorf("query token hold err: %v", err)
		return "", err
	}
	router, err := contract.NewRouter(common.HexToAddress(merlinSwapContractAddress), merlinClient)
	if err != nil {
		return "", err
	}
	merlTotalAmount := tokenHoldMap[strings.ToLower(merlContractAddress)]
	wbtcTotalAmount := tokenHoldMap[strings.ToLower(wbtcTokenContractAddress)]
	feePercent := big.NewFloat(0.01) // 1%
	merlTotalAmountFloat, _, err := new(big.Float).Parse(merlTotalAmount, 10)
	if err != nil {
		log.Errorf("err: %v", err)
		return "", err
	}
	wbtcTotalAmountFloat, _, err := new(big.Float).Parse(wbtcTotalAmount, 10)
	if err != nil {
		log.Errorf("err: %v", err)
		return "", err
	}
	swapMerlAmount, _, err := new(big.Float).Parse(SwapMerlAmount, 10)
	if err != nil {
		log.Errorf("err: %v", err)
		return "", err
	}

	effectiveSwapAmount := new(big.Float).Mul(swapMerlAmount, new(big.Float).Sub(big.NewFloat(1), feePercent))

	newA := new(big.Float).Add(merlTotalAmountFloat, effectiveSwapAmount)

	k := new(big.Float).Mul(wbtcTotalAmountFloat, merlTotalAmountFloat)

	newB := new(big.Float).Quo(k, newA)

	outputB := new(big.Float).Sub(wbtcTotalAmountFloat, newB)
	finalOutputB := new(big.Float).Mul(outputB, new(big.Float).Sub(big.NewFloat(1), big.NewFloat(SlippageTolerance)))
	//log.Infof("merlTotalAmountFloat, %s, wbtcTotalAmountFloat: %s", merlTotalAmountFloat.String(), wbtcTotalAmountFloat.String())
	//log.Infof("effectiveSwapAmount, %s, newA: %s, k: %s, newB: %s", effectiveSwapAmount.String(), newA.String(), k.String(), newB.String())
	router, err = contract.NewRouter(common.HexToAddress(merlinSwapContractAddress), merlinClient)
	if err != nil {
		log.Errorf("err: %v", err)
		return "", err
	}
	var fund int
	rand.Seed(time.Now().Unix())

	deadline := time.Now().Add(time.Minute * 10).Unix()

	fund = util.Random(10, 100)
	log.Info("singleAmount: ", fund)
	//fee3000 := []byte{0x0b, 0xb8} // 0.3% 的手续费编码
	//fee500 := []byte{0x01, 0xf4}  // 0.05% 的手续费编码
	fee10000 := "2710" //  1%的手续费编码

	// 按顺序将地址和手续费编码
	pathBytes := bytes.Join([][]byte{
		common.Hex2Bytes(merlContractAddress[2:]),
		common.Hex2Bytes("00"),
		common.Hex2Bytes(fee10000),
		common.Hex2Bytes(wbtcTokenContractAddress[2:]),
	}, nil)
	//log.Infof("addr: %s", opts.From.String())
	//log.Infof("path: %s", hexutil.Encode(pathBytes))

	tx, err := router.SwapAmount(opts, contract.SwapSwapAmountParams{
		pathBytes,
		opts.From,
		util.ToWei(swapMerlAmount.String(), 18),
		util.ToWei(finalOutputB.String(), 18),
		big.NewInt(deadline),
	})

	if err != nil {
		log.Error(err)
		return "", err
	}
	return tx.Hash().String(), nil
}

type TokenHoldInfo struct {
	Data []TokenHoldData `json:"data"`
}

type TokenHoldData struct {
	TokenAddress string `json:"token_address"`
	Balance      string `json:"balance"`
	TokenName    string `json:"token_name"`
}
