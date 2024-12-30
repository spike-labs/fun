package merlin

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	logger "github.com/ipfs/go-log"
	"math"
	"math/big"
	"os"
	"spike-mc-ops/service/merlin/contract"
	"spike-mc-ops/service/merlin/service"
	"spike-mc-ops/util"
	"strconv"
	"sync"
	"testing"
	"time"
)

var log = logger.Logger("merlin")

func init() {
	logger.SetAllLoggers(logger.LevelDebug) //
}

const (
	loginUrl                   = "https://api-testnet-new.btc.fun/api/v1/login"
	signUrl                    = "https://api-testnet-new.btc.fun/api/v1/token/sign"
	partyTokenContractAddress  = "0x4D9882a3BB13cc086367D0aE964367e6B7ea246f"
	btcFunContractAddress      = "0x15545789c664f9b92703324c0a6423ade08337d1"
	testnetMerlContractAddress = "0x5c46bFF4B38dc1EAE09C5BAc65872a1D8bc87378"
	merlinTestNetRpcAddress    = "https://merlin-mainnet-rpc.merlinchain.io"
	singleAmount               = 50     //每次募资多少个merl
	gasLimit                   = 200000 //募资交易的gasLimit
	gasPriceMultiples          = 1.3
)

var client *ethclient.Client
var chainId *big.Int

func TestOffer(t *testing.T) {
	walletStartIndex := 0
	walletEndIndex := 1
	cli, err := ethclient.Dial(merlinTestNetRpcAddress)
	if err != nil {
		log.Errorf("failed to connect to merlin-mainnet-rpc.merlinchain.io")
		return
	}
	client = cli
	id, err := client.ChainID(context.Background())
	if err != nil {
		log.Errorf("query chain id err: %v", err)
		return
	}
	chainId = id
	toAddressList := GetBtcFunAddressList(walletStartIndex, walletEndIndex)

	f, err := os.OpenFile("offer_log.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Infof("no such file or directory and create new file")
		f, err = os.Create("offer_log.csv")
		if err != nil {
			log.Errorf("err: %v", err)
			return
		}
		_, err = fmt.Fprintln(f, "index,address,tx_hash")
	}

	defer f.Close()

	if err != nil {
		log.Errorf("err: %v", err)
		return
	}
	throttle := make(chan struct{}, 5)
	var wg sync.WaitGroup
	for _, toAddressInfo := range toAddressList {
		wg.Add(1)
		throttle <- struct{}{}
		go func(privateKeyHex string, address string, index int) {
			defer func() {
				wg.Done()
				<-throttle
			}()
			if CheckOfferOf(address) {
				log.Infof("address: %s offered", address)
				return
			}
			accessToken, err := service.Login(privateKeyHex, address, loginUrl)
			if err != nil {
				log.Errorf("err: %v", err)
				return
			}
			signResp, err := service.Sign(accessToken, address, singleAmount, signUrl, partyTokenContractAddress)
			if err != nil {
				log.Errorf("err: %v", err)
				return
			}
			txHash, err := Offer(privateKeyHex, signResp)
			if err != nil {
				log.Errorf("err: %v", err)
				return
			}
			fmt.Fprintf(f, "%d,%s,%s\n", index, address, txHash)
			time.Sleep(3 * time.Second)
		}(toAddressInfo.PrivateKey, toAddressInfo.Address, toAddressInfo.Index)
	}
	wg.Wait()
}

func CheckOfferOf(address string) (flag bool) {
	btcFun, err := contract.NewBtcfun(common.HexToAddress(btcFunContractAddress), client)
	if err != nil {
		log.Errorf("failed to connect to merlin-mainnet-rpc.merlinchain.io")
		return
	}
	amount, err := btcFun.OfferedOf(nil, common.HexToAddress(partyTokenContractAddress), common.HexToAddress(address))
	if err != nil {
		log.Errorf("err: %v", err)
		return true
	}
	log.Debugf("amount: %s", amount.String())
	return amount.Cmp(big.NewInt(0)) > 1
}

func Offer(privateKeyHex string, signResp service.SignResp) (txHash string, err error) {
	btcFun, err := contract.NewBtcfun(common.HexToAddress(btcFunContractAddress), client)
	if err != nil {
		log.Errorf("failed to connect to merlin-mainnet-rpc.merlinchain.io")
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorf("query gas price err: %v", err)
		return
	}
	log.Debugf("gasPrice: %s", gasPrice.String())
	priKey, err := crypto.HexToECDSA(privateKeyHex[2:])
	if err != nil {
		return
	}
	opts, err := bind.NewKeyedTransactorWithChainID(priKey, chainId)
	if err != nil {
		return
	}
	r, err := service.HexToByte32(signResp.Data.R[2:])
	if err != nil {
		log.Errorf("err: %v", err)
		return
	}
	s, err := service.HexToByte32(signResp.Data.S[2:])
	if err != nil {
		log.Errorf("err: %v", err)
		return
	}
	opts.GasLimit = gasLimit
	opts.GasPrice = big.NewInt(int64(math.Ceil(float64(gasPrice.Int64()) * gasPriceMultiples)))
	log.Debugf("expiry: %d, v: %d, s: %v, r: %v", signResp.Data.Expiry, uint8(signResp.Data.V), s, r)
	tx, err := btcFun.Offer(opts, common.HexToAddress(partyTokenContractAddress), util.ToWei(strconv.FormatInt(singleAmount, 10), 18), big.NewInt(signResp.Data.Expiry), uint8(signResp.Data.V), r, s)
	if err != nil {
		log.Errorf("err: %v", err)
		return
	}
	log.Debugf("tx: %s", tx.Hash().Hex())
	return tx.Hash().Hex(), nil
}
