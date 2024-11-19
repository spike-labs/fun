package merlin

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"
	"spike-mc-ops/service/merlin/contract"
	"spike-mc-ops/util"
	"sync"
	"testing"
	"time"
)

const BtcFunWalletMnemonic = ""
const mainWalletPrivateKey = ""
const btcAmount = "0.0001"
const merlAmount = "1"

func TestCheckAllowance(t *testing.T) {
	cli, err := ethclient.Dial(merlinTestNetRpcAddress)
	if err != nil {
		log.Errorf("failed to connect to rpc")
		return
	}
	client = cli
	id, err := client.ChainID(context.Background())
	if err != nil {
		log.Errorf("query chain id err: %v", err)
		return
	}
	chainId = id
	toAddressList := GetBtcFunAddressList(10)
	throttle := make(chan struct{}, 5)
	var wg sync.WaitGroup
	for _, toAddressInfo := range toAddressList {
		wg.Add(1)
		throttle <- struct{}{}
		go func(privateKeyHex string) {
			defer func() {
				wg.Done()
				<-throttle
			}()
			util.CheckAllowance(privateKeyHex, testnetMerlContractAddress, btcFunContractAddress, chainId, client)
			time.Sleep(3 * time.Second)
		}(toAddressInfo.PrivateKey)
	}
	wg.Wait()
}

func TestDeliverMerl(t *testing.T) {
	cli, err := ethclient.Dial(merlinTestNetRpcAddress)
	if err != nil {
		log.Errorf("failed to connect to rpc")
		return
	}
	client = cli
	id, err := client.ChainID(context.Background())
	if err != nil {
		log.Errorf("query chain id err: %v", err)
		return
	}
	chainId = id
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorf("query gas price err: %v", err)
		return
	}
	merlContract, err := contract.NewErc20(common.HexToAddress(testnetMerlContractAddress), client)
	if err != nil {
		log.Error("NewErc20 err: ", err)
		return
	}
	toAddressList := GetBtcFunAddressList(10)
	for _, toAddressInfo := range toAddressList {
		priKey, err := crypto.HexToECDSA(mainWalletPrivateKey[2:])
		if err != nil {
			return
		}

		opts, err := bind.NewKeyedTransactorWithChainID(priKey, id)
		if err != nil {
			return
		}
		toAddress := common.HexToAddress(toAddressInfo.Address)
		merlBalance, err := merlContract.BalanceOf(nil, toAddress)
		if err != nil {
			log.Error("Query balance err: ", err)
			return
		}
		log.Debugf("merl balance %s", merlBalance.String())
		if merlBalance.Int64() > 0 {
			log.Infof("merl balance enough, address: %s, balance: %d", toAddress.String(), merlBalance.Int64())
			continue
		}
		opts.GasLimit = 200000
		opts.GasPrice = big.NewInt(int64(math.Ceil(float64(gasPrice.Int64()) * gasPriceMultiples)))
		tx, err := merlContract.Transfer(opts, common.HexToAddress(toAddress.String()), util.ToWei(merlAmount, 18))
		if err != nil {
			log.Error("transfer merl err: ", err)
			return
		}
		log.Infof("send transaction success. tx: %s", tx.Hash().Hex())
		time.Sleep(5 * time.Second)
	}
}

func TestDeliverBtc(t *testing.T) {
	cli, err := ethclient.Dial(merlinTestNetRpcAddress)
	if err != nil {
		log.Errorf("failed to connect to rpc")
		return
	}
	client = cli
	id, err := client.ChainID(context.Background())
	if err != nil {
		log.Errorf("query chain id err: %v", err)
		return
	}
	chainId = id
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorf("query gas price err: %v", err)
		return
	}
	currentBlockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Error("query block number err: ", err)
		return
	}
	mainAddress, err := util.GenerateAddress(mainWalletPrivateKey)

	toAddressList := GetBtcFunAddressList(10)
	for _, toAddressInfo := range toAddressList {
		nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(mainAddress))
		if err != nil {
			log.Error("query nonce err: ", err)
			return
		}
		priKey, err := crypto.HexToECDSA(mainWalletPrivateKey[2:])
		if err != nil {
			return
		}

		opts, err := bind.NewKeyedTransactorWithChainID(priKey, id)
		if err != nil {
			return
		}
		toAddress := common.HexToAddress(toAddressInfo.Address)

		btcBalance, err := client.BalanceAt(context.Background(), common.HexToAddress(toAddress.String()), big.NewInt(int64(currentBlockNumber)))
		if err != nil {
			log.Error("query balance err: ", err)
			return
		}
		if btcBalance.Int64() > 0 {
			log.Infof("btc balance enough, %d", btcBalance.Int64())
			continue
		}
		gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
			From:     common.HexToAddress(mainAddress),
			To:       &toAddress,
			Value:    util.ToWei(btcAmount, 18),
			GasPrice: big.NewInt(int64(math.Ceil(float64(gasPrice.Int64()) * gasPriceMultiples))),
		})
		if err != nil {
			log.Error(err)
			return
		}

		signedTx, err := opts.Signer(common.HexToAddress(mainAddress), types.NewTx(
			&types.LegacyTx{
				Nonce:    nonce,
				Gas:      gasLimit,
				GasPrice: big.NewInt(int64(math.Ceil(float64(gasPrice.Int64()) * gasPriceMultiples))),
				Value:    util.ToWei(btcAmount, 18),
				To:       &toAddress,
			}))
		if err != nil {
			log.Error(err)
			return
		}

		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			log.Error(err)
			return
		}
		log.Infof("send transaction success. tx: %s", signedTx.Hash().Hex())
		time.Sleep(5 * time.Second)
	}
}

func GetBtcFunAddressList(num int) []AddressInfo {
	addressList := make([]AddressInfo, 0)
	for i := 0; i < num; i++ {
		priKeyString, err := util.DerivePrivateKeyWithNumber(BtcFunWalletMnemonic, i)
		if err != nil {
			return nil
		}
		address, err := util.GenerateAddress(priKeyString)
		log.Debugf("address: %s, privateKey: %s", address, priKeyString)
		addressList = append(addressList, AddressInfo{
			Address:    address,
			PrivateKey: priKeyString,
		})
	}
	return addressList
}

type AddressInfo struct {
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
}
