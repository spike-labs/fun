package mcMgrSrv

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"math/big"
	"math/rand"
	"spike-mc-ops/chain/contract"
	"spike-mc-ops/config"
	"spike-mc-ops/request"
	"spike-mc-ops/util"
	"time"
)

func (m *MCManager) AddBuyNFTStrategy(t request.BuyNFTService) string {
	u := uuid.New().String()
	buyNFTStrategy := request.BuyNFTStrategy{
		Time:      t.Time,
		BuyAmount: t.BuyAmount,
		Closing:   make(chan struct{}, 0),
	}

	m.BuyNFTStrategies[u] = buyNFTStrategy

	go m.ExecBuyNFTStrategy(u)
	return u
}

func (m *MCManager) CancelBuyNFTStrategy(u string) {
	if _, have := m.BuyNFTStrategies[u]; have {
		m.BuyNFTStrategies[u].Closing <- struct{}{}
		delete(m.BuyNFTStrategies, u)
	}
}

func (m *MCManager) ExecBuyNFTStrategy(uuid string) {
	strategy := m.BuyNFTStrategies[uuid]
	buyFrequency := strategy.Time * 3600 / strategy.BuyAmount

	ticker := time.NewTicker(time.Second * time.Duration(buyFrequency+1))

loop:
	for {
		select {
		case <-ticker.C:
			m.BuyNFT(uuid)

			ticker.Reset(time.Duration(util.Random(0, int(buyFrequency))) * time.Second)
		case <-strategy.Closing:
			delete(m.BuyNFTStrategies, uuid)
			log.Infof("buyNFTstrategy close...")
			break loop
		}
	}
	return
}

func (m *MCManager) BuyNFT(uuid string) {
	strategy := m.BuyNFTStrategies[uuid]
	// todo 查询合约中剩余nft数量
	leftNftAmount := big.NewInt(0)
	if leftNftAmount.Cmp(big.NewInt(strategy.BuyAmount)) == 0 || leftNftAmount.Cmp(big.NewInt(strategy.BuyAmount)) == -1 {
		strategy.Closing <- struct{}{}
		return
	}

	m.PickWallet()
	// todo buy nft

}

func (m *MCManager) PickWallet() *bind.TransactOpts {
	usdcContract, err := contract.NewErc20Contract(common.HexToAddress(config.Cfg.Contract.UsdcAddress), m.Wallet.BscClient)
	if err != nil {
		log.Errorf("construct usdcContract err :%v", err)
		return nil
	}

	// todo 查询盲盒价格
	nftPrice := big.NewInt(0)
	for i := 0; i < len(m.Wallet.PuppetWallets); i++ {

		rand.Seed(time.Now().UnixNano())
		wallet := m.Wallet.PuppetWallets[rand.Intn(len(m.Wallet.PuppetWallets))]
		usdcBalance, err := usdcContract.BalanceOf(nil, wallet.From)
		if err != nil {
			log.Error(err)
			return nil
		}
		if usdcBalance.Cmp(nftPrice) == 1 || usdcBalance.Cmp(nftPrice) == 0 {
			return wallet
		}

	}
	return nil
}
