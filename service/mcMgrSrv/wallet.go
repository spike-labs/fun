package mcMgrSrv

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	contract "spike-mc-ops/chain/contract"
	"spike-mc-ops/config"
	"spike-mc-ops/util"
)

type PuppetWallet struct {
	PuppetWallets []*bind.TransactOpts
	BscClient     *ethclient.Client
	Router        *contract.PanCakeRouter
}

func NewPuppetWallet() *PuppetWallet {

	bscClient, err := ethclient.Dial(config.Cfg.Chain.RpcNodeAddress)
	if err != nil {
		panic("=== Spike log: ")
	}

	router, err := contract.NewPanCakeRouter(common.HexToAddress(config.Cfg.Contract.PanCakeRouterAddress), bscClient)
	if err != nil {
		panic("=== Spike log: ")
	}

	if len(config.Cfg.PuppetWallet.PrivateKey) == 0 {
		panic("Please Deploy PuppetWallet")
	}

	id, err := bscClient.ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	puppetWallets := make([]*bind.TransactOpts, 0)

	for i := range config.Cfg.PuppetWallet.PrivateKey {
		priKey, err := crypto.HexToECDSA(config.Cfg.PuppetWallet.PrivateKey[i][2:])
		if err != nil {
			panic(err)
		}
		util.CheckAllowance(config.Cfg.PuppetWallet.PrivateKey[i], config.Cfg.Contract.UsdcAddress, config.Cfg.Contract.PanCakeRouterAddress, id, bscClient)
		util.CheckAllowance(config.Cfg.PuppetWallet.PrivateKey[i], config.Cfg.Contract.GameTokenAddress, config.Cfg.Contract.PanCakeRouterAddress, id, bscClient)
		opts, err := bind.NewKeyedTransactorWithChainID(priKey, id)
		if err != nil {
			panic(err)
		}

		puppetWallets = append(puppetWallets, opts)
	}

	return &PuppetWallet{
		BscClient:     bscClient,
		Router:        router,
		PuppetWallets: puppetWallets,
	}
}
