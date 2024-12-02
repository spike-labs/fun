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

var MainWalletPrivateKeyList = []string{
	"",
	"",
	"",
}

const btcAmount = "0.0001"
const merlAmount = "1"

func TestCheckAllowance(t *testing.T) {
	walletStartIndex := 0
	walletEndIndex := 10

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
	toAddressList := GetBtcFunAddressList(walletStartIndex, walletEndIndex)
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

// wbtc归集
func TestERC20Aggregation(t *testing.T) {
	mainWalletPrivateIndex := 0 //归集到第几个主钱包
	walletStartIndex := 0
	walletEndIndex := 10

	cli, err := ethclient.Dial(merlinMainNetRpcAddress)
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
	wbtcTokenContract, err := contract.NewErc20(common.HexToAddress(wbtcTokenContractAddress), client)
	if err != nil {
		log.Error("NewErc20 err: ", err)
		return
	}
	mainAddress, err := util.GenerateAddress(MainWalletPrivateKeyList[mainWalletPrivateIndex])
	if err != nil {
		log.Error("GenerateAddress err: ", err)
		return
	}
	fromAddressList := GetBtcFunAddressList(walletStartIndex, walletEndIndex)
	for _, fromAddressInfo := range fromAddressList {
		wbtcBalance, err := wbtcTokenContract.BalanceOf(nil, common.HexToAddress(fromAddressInfo.Address))
		if err != nil {
			log.Error("Query balance err: ", err)
			return
		}
		log.Debugf("wbtc balance %s", wbtcBalance.String())
		if wbtcBalance.Int64() == 0 {
			log.Debugf("wbtc balance is zero, address: %s, balance: %d", fromAddressInfo, wbtcBalance.Int64())
			continue
		}
		priKey, err := crypto.HexToECDSA(fromAddressInfo.PrivateKey[2:])
		if err != nil {
			return
		}

		opts, err := bind.NewKeyedTransactorWithChainID(priKey, id)
		if err != nil {
			return
		}
		opts.GasLimit = 200000
		opts.GasPrice = big.NewInt(int64(math.Ceil(float64(gasPrice.Int64()) * gasPriceMultiples)))
		tx, err := wbtcTokenContract.Transfer(opts, common.HexToAddress(mainAddress), wbtcBalance)
		if err != nil {
			log.Error("transfer wbtc err: ", err)
			return
		}
		log.Infof("send transaction success. tx: %s", tx.Hash().Hex())
		time.Sleep(5 * time.Second)
	}
}

func TestDeliverMerl(t *testing.T) {
	mainWalletPrivateIndex := 0
	walletStartIndex := 0
	walletEndIndex := 10

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
	toAddressList := GetBtcFunAddressList(walletStartIndex, walletEndIndex)
	for _, toAddressInfo := range toAddressList {
		priKey, err := crypto.HexToECDSA(MainWalletPrivateKeyList[mainWalletPrivateIndex][2:])
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
	mainWalletPrivateIndex := 0
	walletStartIndex := 0
	walletEndIndex := 10
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
	mainAddress, err := util.GenerateAddress(MainWalletPrivateKeyList[mainWalletPrivateIndex])

	toAddressList := GetBtcFunAddressList(walletStartIndex, walletEndIndex)
	for _, toAddressInfo := range toAddressList {
		nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(mainAddress))
		if err != nil {
			log.Error("query nonce err: ", err)
			return
		}
		priKey, err := crypto.HexToECDSA(MainWalletPrivateKeyList[mainWalletPrivateIndex][2:])
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

func GetBtcFunAddressList(start int, end int) []AddressInfo {
	addressList := make([]AddressInfo, 0)
	for i := start; i < end; i++ {
		priKeyString, err := util.DerivePrivateKeyWithNumber(BtcFunWalletMnemonic, i)
		if err != nil {
			return nil
		}
		address, err := util.GenerateAddress(priKeyString)
		//log.Debugf("address: %s, privateKey: %s", address, priKeyString)
		addressList = append(addressList, AddressInfo{
			Address:    address,
			PrivateKey: priKeyString,
			Index:      i,
		})
	}
	return addressList
}

type AddressInfo struct {
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
	Index      int    `json:"index"`
}
