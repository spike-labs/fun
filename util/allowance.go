package util

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"
	"spike-mc-ops/service/merlin/contract"
)

const allowanceAmount = "10000000000000000000"

func CheckAllowance(privateKey string, erc20ContractAddr string, panCakeRouterAddress string, chainId *big.Int, client *ethclient.Client) {
	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey[2:])
	if err != nil {
		log.Errorf("parse privateKey err : %v", err)
		return
	}
	publicKey, err := GenerateAddress(privateKey)
	if err != nil {
		log.Errorf("generate publicKey err : %v", err)
		return
	}

	erc20Contract, err := contract.NewErc20(common.HexToAddress(erc20ContractAddr), client)
	if err != nil {
		log.Errorf("construct usdcContract err : %v", err)
		return
	}
	erc20Allowance, err := erc20Contract.Allowance(nil, common.HexToAddress(publicKey), common.HexToAddress(panCakeRouterAddress))
	if err != nil {
		log.Errorf("query usdcContract allowance err : %v", err)
		return
	}
	if erc20Allowance.Cmp(big.NewInt(0)) > 0 {
		log.Debugf("allowance : %s", erc20Allowance.String())
		return
	}

	txOps, err := bind.NewKeyedTransactorWithChainID(ecdsaPrivateKey, chainId)
	if err != nil {
		log.Errorf("construct txOps err : %v", err)
		return
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorf("query gas price err: %v", err)
		return
	}
	txOps.GasLimit = 80000
	txOps.GasPrice = big.NewInt(int64(math.Ceil(float64(gasPrice.Int64()) * 1.3)))
	tx, err := erc20Contract.Approve(txOps, common.HexToAddress(panCakeRouterAddress), ToWei(allowanceAmount, 18))
	if err != nil {
		log.Errorf("usdcContract Approve  err : %v", err)
		return
	}
	log.Infof("address: %s approve erc20Contract: %s,  txHash: %s", publicKey, erc20ContractAddr, tx.Hash().String())
}
