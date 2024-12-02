package merlin

import (
	"bytes"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	logger "github.com/ipfs/go-log"
	"math/big"
	"spike-mc-ops/service/merlin/contract"
	"spike-mc-ops/util"
	"sync"
	"testing"
	"time"
)

var (
	walletMnemonic                  = ""
	merlinSwapContractAddress       = "0x1aFa5D7f89743219576Ef48a9826261bE6378a68" //
	wbtcTokenContractAddress        = "0xF6D226f9Dc15d9bB51182815b320D3fBE324e1bA"
	merlContractAddress             = "0x5c46bFF4B38dc1EAE09C5BAc65872a1D8bc87378"
	swapContractAddress             = "0x2426191006F378BF33445f87938d355096eE2e8C" //池子合约
	merlinMainNetRpcAddress         = "https://merlin-mainnet-rpc.merlinchain.io"
	merlinSwapQuoterContractAddress = "0x2569bce69287618e2cd004f785d016f7df29232f"
	SlippageTolerance               = 0.05
	SwapMerlAmount                  = "5"
)

type TxOps struct {
	Ops   *bind.TransactOpts
	Index int64
}

var lock sync.Mutex
var merlinClient *ethclient.Client

func TestQuoter(t *testing.T) {
	logger.SetAllLoggers(logger.LevelInfo)
	client, err := ethclient.Dial(merlinMainNetRpcAddress)
	if err != nil {
		log.Errorf("failed to connect to merlin-mainnet-rpc.merlinchain.io")
		return
	}
	merlinClient = client
	opts := &bind.CallOpts{
		Context: context.Background(),
	}
	quoter, err := contract.NewQuoter(common.HexToAddress("0x2569bce69287618e2cd004f785d016f7df29232f"), merlinClient)
	if err != nil {
		log.Errorf("failed to connect to merlin-mainnet-rpc.merlinchain.io")
		return
	}
	tokenX := common.HexToAddress("0x480e158395cc5b41e5584347c495584ca2caf78d")
	tokenY := common.HexToAddress("0xF6D226f9Dc15d9bB51182815b320D3fBE324e1bA")
	fee := big.NewInt(10000)         // 假设手续费等级 0.3%
	amount := util.ToWei("1000", 18) // 输入 1 token
	lowPt := big.NewInt(-887272)
	amountY, finalPoint, err := quoter.SwapX2YStatic(opts, tokenX, tokenY, fee, amount, lowPt)
	if err != nil {
		log.Errorf("failed to swap x2Y static quoter opts: %v", err)
		return
	}
	log.Infof("swap x2Y static quoter opts: %s, %s", amountY.String(), finalPoint.String())
}

func TestSwapMerl2Token(t *testing.T) {
	logger.SetAllLoggers(logger.LevelDebug)
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
	return
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
				log.Debugf("address: %s,  index: %d, wbtcWalletBalance: %s", opts.From, index, wbtcWalletBalance.String())
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
			time.Sleep(5 * time.Second)
		}(i)
	}
	wg.Wait()
	return txOps, nil
}

func BuyToken(opts *bind.TransactOpts) (txHash string, err error) {
	router, err := contract.NewRouter(common.HexToAddress(merlinSwapContractAddress), merlinClient)
	if err != nil {
		log.Errorf("err: %v", err)
		return "", err
	}
	deadline := time.Now().Add(time.Minute * 100).Unix()
	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}
	quoter, err := contract.NewQuoter(common.HexToAddress(merlinSwapQuoterContractAddress), merlinClient)
	if err != nil {
		log.Errorf("failed to connect to merlin-mainnet-rpc.merlinchain.io")
		return
	}
	tokenX := common.HexToAddress(merlContractAddress)
	tokenY := common.HexToAddress(wbtcTokenContractAddress)
	fee := big.NewInt(10000)                 // 假设手续费等级 0.3%
	amount := util.ToWei(SwapMerlAmount, 18) // 输入 merl token数量
	lowPt := big.NewInt(-887272)
	amountY, _, err := quoter.SwapX2YStatic(callOpts, tokenX, tokenY, fee, amount, lowPt)
	if err != nil {
		log.Errorf("failed to swap x2Y static quoter opts: %v", err)
		return
	}

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
	floatTokenY, _, err := new(big.Float).Parse(amountY.String(), 10)
	if err != nil {
		log.Errorf("failed to parse float quoter opts: %v", err)
		return
	}
	finalTokenY := new(big.Float).Mul(floatTokenY, new(big.Float).Sub(big.NewFloat(1), big.NewFloat(SlippageTolerance)))
	log.Debugf("final token y: %s, %s", amountY.String(), finalTokenY.String())
	tx, err := router.SwapAmount(opts, contract.SwapSwapAmountParams{
		pathBytes,
		opts.From,
		util.ToWei(SwapMerlAmount, 18),
		BigFloat2Int(finalTokenY),
		big.NewInt(deadline),
	})

	if err != nil {
		log.Error(err)
		return "", err
	}
	return tx.Hash().String(), nil
}

func BigFloat2Int(floatNum *big.Float) (intNum *big.Int) {

	// 定义精度
	scale := new(big.Float).SetFloat64(0.5)

	// 四舍五入
	rounded := new(big.Float).Add(floatNum, scale)

	// 转换为 big.Int
	bigInt := new(big.Int)
	rounded.Int(bigInt)
	return bigInt
}

type TokenHoldInfo struct {
	Data []TokenHoldData `json:"data"`
}

type TokenHoldData struct {
	TokenAddress string `json:"token_address"`
	Balance      string `json:"balance"`
	TokenName    string `json:"token_name"`
}
