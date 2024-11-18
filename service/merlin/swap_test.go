package merlin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-resty/resty/v2"
	logger "github.com/ipfs/go-log"
	"math/big"
	"math/rand"
	"spike-mc-ops/service/merlin/contract"
	"spike-mc-ops/util"
	"strconv"
	"sync"
	"testing"
	"time"
)

var (
	walletMnemonic             = ""
	merlinSwapContractAddress  = "0x1aFa5D7f89743219576Ef48a9826261bE6378a68"
	bitmapTokenContractAddress = "0x7b0400231Cddf8a7ACa78D8c0483890cd0c6fFD6"
	merlContractAddress        = "0x5c46bFF4B38dc1EAE09C5BAc65872a1D8bc87378"
	swapContractAddress        = "0x501ca56E4b6Af84CBAAaaf2731D7C87Bed32ee65"
	merlinMainNetRpcAddress    = "https://merlin-mainnet-rpc.merlinchain.io"
)

var lock sync.Mutex
var merlinClient *ethclient.Client

func TestSwap(t *testing.T) {
	logger.SetAllLoggers(logger.LevelInfo)
	//QueryTokenHold(swapContractAddress)
	client, err := ethclient.Dial(merlinMainNetRpcAddress)
	if err != nil {
		log.Errorf("failed to connect to merlin-mainnet-rpc.merlinchain.io")
		return
	}
	merlinClient = client
	txOps, err := InitWallet()
	if err != nil {
		log.Errorf("failed to init wallet")
		return
	}
	_ = txOps
	//BuyToken(txOps[0])
}

func InitWallet() ([]*bind.TransactOpts, error) {
	var wg sync.WaitGroup
	id, err := merlinClient.ChainID(context.Background())
	if err != nil {
		log.Errorf("query chain id err: %v", err)
		return nil, err
	}
	log.Infof("id: %d", id.Int64())
	puppetWallets := make([]*bind.TransactOpts, 0)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			priKeyString, err := util.DerivePrivateKeyWithNumber(walletMnemonic, index)
			if err != nil {
				return
			}
			address, err := util.GenerateAddress(priKeyString)
			log.Infof("address: %s, privateKey: %s", address, priKeyString)
			priKey, err := crypto.HexToECDSA(priKeyString[2:])
			if err != nil {
				return
			}

			opts, err := bind.NewKeyedTransactorWithChainID(priKey, id)
			if err != nil {
				return
			}
			lock.Lock()
			puppetWallets = append(puppetWallets, opts)
			lock.Unlock()

			util.CheckAllowance(priKeyString, merlContractAddress, merlinSwapContractAddress, id, merlinClient)
			util.CheckAllowance(priKeyString, bitmapTokenContractAddress, merlinSwapContractAddress, id, merlinClient)
		}(i)
	}
	wg.Wait()
	return puppetWallets, nil
}

func QueryTokenHold(poolAddress string) {
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
	log.Infof("tokenHoldInfo: %v", tokenHoldInfo)
}

func BuyToken(opts *bind.TransactOpts) {

	router, err := contract.NewRouter(common.HexToAddress(merlinSwapContractAddress), merlinClient)
	if err != nil {
		return
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
		common.Hex2Bytes(bitmapTokenContractAddress[2:]),
	}, nil)
	log.Infof("addr: %s", opts.From.String())
	log.Infof("path: %s", hexutil.Encode(pathBytes))
	//swapContract, err := contract.NewSwap(common.HexToAddress(swapContractAddress), merlinClient)
	//if err != nil {
	//	return
	//}
	log.Infof("amount: %d", util.ToWei(strconv.FormatInt(100, 10), 18))
	//return
	tx, err := router.SwapAmount(opts, contract.SwapSwapAmountParams{
		pathBytes,
		opts.From,
		util.ToWei(strconv.FormatInt(88, 10), 18),
		util.ToWei(strconv.FormatInt(500, 10), 18),
		big.NewInt(deadline),
	})

	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("tx: %s", tx.Hash().String())
}

type TokenHoldInfo struct {
	Data []TokenHoldData `json:"data"`
}

type TokenHoldData struct {
	TokenAddress string `json:"token_address"`
	Balance      string `json:"balance"`
	TokenName    string `json:"token_name"`
}
