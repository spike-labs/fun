package merlin

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	logger "github.com/ipfs/go-log"
	"math"
	"math/big"
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
	loginUrl                   = "https://api-testnet.btc.fun/user/login"
	signUrl                    = "https://api-testnet.btc.fun/token/sign"
	nonceUrl                   = "https://api-testnet.btc.fun/user/nonce?publicKey="
	partyTokenContractAddress  = "0xF4055a9DaE32eB96Fd55dc49d84f30727757fe3A"
	btcFunContractAddress      = "0xD63DE91412Be4817b4fDB39993cDE784b9497bcD"
	testnetMerlContractAddress = "0x5c9ad13be752a5e21e7ed393bb7d407144c6d550"
	merlinTestNetRpcAddress    = "https://testnet-rpc.merlinchain.io"
	singleAmount               = 2      //每次募资多少个merl
	gasLimit                   = 200000 //募资交易的gasLimit
	gasPriceMultiples          = 1.3
)

var client *ethclient.Client
var chainId *big.Int

func TestOffer(t *testing.T) {
	walletStartIndex := 0
	walletEndIndex := 10
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

	throttle := make(chan struct{}, 5)
	var wg sync.WaitGroup
	for _, toAddressInfo := range toAddressList {
		wg.Add(1)
		throttle <- struct{}{}
		go func(privateKeyHex string, address string) {
			defer func() {
				wg.Done()
				<-throttle
			}()
			if CheckOfferOf(address) {
				log.Infof("address: %s offered", address)
				return
			}
			nonce, err := service.QueryNonce(address, nonceUrl)
			if err != nil {
				log.Errorf("err: %v", err)
				return
			}
			log.Debugf("nonceRes: %s", nonce.Data.Nonce)
			accessToken, err := service.Login(privateKeyHex, address, nonce.Data.Nonce, loginUrl)
			if err != nil {
				log.Errorf("err: %v", err)
				return
			}
			log.Debugf("access token: %s", accessToken)
			signResp, err := service.Sign(accessToken, address, singleAmount, signUrl, partyTokenContractAddress)
			if err != nil {
				log.Errorf("err: %v", err)
				return
			}
			Offer(privateKeyHex, signResp)
			time.Sleep(3 * time.Second)
		}(toAddressInfo.PrivateKey, toAddressInfo.Address)
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
	return amount.Int64() > 0
}

func Offer(privateKeyHex string, signResp service.SignResp) {
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
	log.Debugf("expiry: %d, v: %d, s: %v, r: %v", signResp.Data.TimeStamp, uint8(signResp.Data.V), s, r)
	tx, err := btcFun.Offer(opts, common.HexToAddress(partyTokenContractAddress), util.ToWei(strconv.FormatInt(singleAmount, 10), 18), big.NewInt(signResp.Data.TimeStamp), uint8(signResp.Data.V), r, s)
	if err != nil {
		log.Errorf("err: %v", err)
		return
	}
	log.Debugf("tx: %s", tx.Hash().Hex())
}
