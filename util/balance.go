package util

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"spike-mc-ops/config"
)

func QueryNativeBalance(address string, client *ethclient.Client) (string, error) {
	bnbBalance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		log.Errorf("query balance err : %v", err)
		return "", err
	}
	return ParseBalance(bnbBalance), nil
}

func ParseBalance(balance *big.Int) string {
	decimalBalance := ToDecimal(balance, 18)
	return decimalBalance.String()
}

func TransferBNB(amount float64, publicKey, privateKey, toAddress string) (txHash string, err error) {

	ec, err := ethclient.Dial(config.Cfg.Chain.RpcNodeAddress)
	if err != nil {
		log.Errorf("dial client err : %v", err)
		return "", err
	}
	price, err := ec.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorf("SuggestGasPrice err : %v", err)
		return "", err
	}

	nonce, err := ec.PendingNonceAt(context.Background(), common.HexToAddress(publicKey))
	if err != nil {
		log.Errorf("PendingNonceAt err : %v", err)
		return "", err
	}

	toAddr := common.HexToAddress(toAddress)
	tx := &types.LegacyTx{
		Value:    ToWei(amount, 18),
		GasPrice: price,
		To:       &toAddr,
		Nonce:    nonce,
	}

	gas, err := ec.EstimateGas(context.Background(), ethereum.CallMsg{
		From:     common.HexToAddress(publicKey),
		To:       &toAddr,
		GasPrice: tx.GasPrice,
		Value:    tx.Value,
	})
	if err != nil {
		log.Error("Estimate gas", err)
		return "", err
	}

	chainId, err := ec.ChainID(context.Background())
	if err != nil {
		log.Error("ChainID ", err)
		return "", err
	}

	tx.Gas = uint64(float64(gas))
	ntx := types.NewTx(tx)
	priKey, err := crypto.HexToECDSA(privateKey[2:])
	if err != nil {
		log.Error(err)
		return "", err
	}
	signedTx, err := types.SignTx(ntx, types.NewEIP155Signer(chainId), priKey)
	if err != nil {
		log.Error("Sign transaction", err)
		return "", err
	}

	err = ec.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Error("Send transaction", err)
		return "", err
	}
	return signedTx.Hash().String(), nil
}
