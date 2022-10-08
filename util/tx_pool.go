package util

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"spike-mc-ops/chain/contract"
	"spike-mc-ops/config"
	"strconv"
	"time"
)

func BuyGameToken(gameTokenAmount *big.Int, slippage int64, privateKey string, chainId *big.Int, client *ethclient.Client) error {
	panCakeRouter, err := contract.NewPanCakeRouter(common.HexToAddress(config.Cfg.Contract.PanCakeRouterAddress), client)
	if err != nil {
		log.Errorf("construct panCakeRouter err : %v", err)
		return err
	}
	txPool, err := contract.NewTxPool(common.HexToAddress(config.Cfg.Contract.GameTokenDexPoolAddress), client)
	if err != nil {
		log.Errorf("construct txPool err : %v", err)
		return err
	}
	reserves, err := txPool.GetReserves(nil)
	if err != nil {
		log.Errorf("query reserve err : %v", err)
		return err
	}
	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey[2:])
	if err != nil {
		log.Errorf("parse privateKey err : %v", err)
		return err
	}
	publicKey, err := GenerateAddress(privateKey)
	if err != nil {
		log.Errorf("generate publicKey err : %v", err)
		return err
	}
	txOps, err := bind.NewKeyedTransactorWithChainID(ecdsaPrivateKey, chainId)
	if err != nil {
		log.Errorf("construct txOps err : %v", err)
		return err
	}
	sksIn := gameTokenAmount
	usdcOut, err := panCakeRouter.GetAmountOut(nil, sksIn, reserves.Reserve0, reserves.Reserve1)
	if err != nil {
		log.Errorf("get amount out err : %v", err)
		return err
	}
	log.Infof("reserves : %s, %s", reserves.Reserve0.String(), reserves.Reserve1.String())

	var usdcOutMin big.Rat
	usdcOutMin.Mul(big.NewRat(usdcOut.Int64(), 1), big.NewRat(950, 1000))
	usdcOutMinResult, err := strconv.ParseInt(usdcOutMin.FloatString(0), 10, 64)
	if err != nil {
		log.Errorf("parse usdcInMax err : %v", err)
		return err
	}
	log.Infof("usdcOutMinResult : %d", usdcOutMinResult)
	path := []common.Address{common.HexToAddress(config.Cfg.Contract.GameTokenAddress), common.HexToAddress(config.Cfg.Contract.UsdcAddress)}
	deadline := big.NewInt(time.Now().Add(10 * time.Minute).Unix())
	log.Infof("sksIn : %s", sksIn.String())
	tx, err := panCakeRouter.SwapExactTokensForTokens(txOps, sksIn, big.NewInt(usdcOutMinResult), path, common.HexToAddress(publicKey), deadline)
	if err != nil {
		log.Errorf("swap err : %v", err)
		return err
	}
	log.Infof("swap tx hash : %s", tx.Hash())
	return nil
}
