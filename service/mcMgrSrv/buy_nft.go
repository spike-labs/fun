package mcMgrSrv

import (
	"context"
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
	buyNFTStrategy := &request.BuyNFTStrategy{
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
	}
}

func (m *MCManager) ExecBuyNFTStrategy(uuid string) {

	nftDistribution, err := contract.NewNFTDistribution(common.HexToAddress(config.Cfg.Contract.NFTDistributionAddress), m.Wallet.BscClient)
	if err != nil {
		log.Errorf("construct nftDistributionContract err :%v", err)
		return
	}

	leftNftAmount := big.NewInt(0)

	toTokenID, err := nftDistribution.ToTokenId(nil)
	if err != nil {
		log.Error(err)
		return
	}

	currentTokenID, err := nftDistribution.CurrentTokenId(nil)
	if err != nil {
		log.Error(err)
		return
	}

	leftNftAmount.Sub(big.NewInt(int64(toTokenID)), big.NewInt(int64(currentTokenID))).Add(leftNftAmount, big.NewInt(1))

	m.BuyNFTStrategies[uuid].LeftAmount = leftNftAmount.Int64() - m.BuyNFTStrategies[uuid].BuyAmount

	buyFrequency := m.BuyNFTStrategies[uuid].Time / m.BuyNFTStrategies[uuid].BuyAmount

	ticker := time.NewTicker(time.Second * time.Duration(buyFrequency+1))

loop:
	for {
		select {
		case <-ticker.C:
			m.BuyNFT(uuid)

			ticker.Reset(time.Duration(util.Random(10, int(buyFrequency))) * time.Second)
		case <-m.BuyNFTStrategies[uuid].Closing:
			log.Info("buyNFTstrategy close...")
			delete(m.BuyNFTStrategies, uuid)
			break loop
		}
	}
	return
}

func (m *MCManager) BuyNFT(uuid string) {
	nftDistribution, err := contract.NewNFTDistribution(common.HexToAddress(config.Cfg.Contract.NFTDistributionAddress), m.Wallet.BscClient)
	if err != nil {
		log.Errorf("construct nftDistributionContract err :%v", err)
		return
	}
	leftNftAmount := big.NewInt(0)

	toTokenID, err := nftDistribution.ToTokenId(nil)
	if err != nil {
		log.Error(err)
		return
	}

	currentTokenID, err := nftDistribution.CurrentTokenId(nil)
	if err != nil {
		log.Error(err)
		return
	}

	if toTokenID-currentTokenID < 0 {
		log.Info("The event is over and the sale is over.")
		m.BuyNFTStrategies[uuid].Closing <- struct{}{}
		return
	}

	leftNftAmount.Sub(big.NewInt(int64(toTokenID)), big.NewInt(int64(currentTokenID))).Add(leftNftAmount, big.NewInt(1))

	if leftNftAmount.Cmp(big.NewInt(m.BuyNFTStrategies[uuid].LeftAmount)) == 0 || leftNftAmount.Cmp(big.NewInt(m.BuyNFTStrategies[uuid].LeftAmount)) == -1 {
		m.BuyNFTStrategies[uuid].Closing <- struct{}{}
		return
	}

	buyer := m.PickWallet()

	rand.Seed(time.Now().UnixNano())
	whiteAddress := util.TailorWalletPrefix(config.Cfg.WhiteList.Address[rand.Intn(len(config.Cfg.WhiteList.Address))])
	proofPath, _, err := m.WhiteList.GetMerklePath(util.UserInfo{Address: whiteAddress})
	if err != nil {
		log.Error(err)
		return
	}
	log.Info(whiteAddress)
	proofRoot := m.WhiteList.MerkleRoot()
	log.Info("proof root", common.Bytes2Hex(proofRoot))

	var root [32]byte
	copy(root[:], proofRoot[:32])

	path := make([][32]byte, len(proofPath))
	for i := range proofPath {
		log.Info(common.Bytes2Hex(proofPath[i]))
		copy(path[i][:], proofPath[i][:32])
	}
	rule, err := nftDistribution.SpecialRateRule(nil, root)
	if err != nil {
		log.Info(err)
		return
	}

	if !rule.IsActive {
		log.Error("specialRateRule is notActive")
		return
	}

	price, err := nftDistribution.GetCurrentSalePrice(nil, buyer.From, common.HexToAddress(whiteAddress), rule.RateRule)
	if err != nil {
		log.Error(err)
		return
	}

	buyer.Value = price
	log.Info(util.ToDecimal(buyer.Value, 18).String())
	tx, err := nftDistribution.ClaimWithProof(buyer, common.HexToAddress(whiteAddress), root, path)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info("buyNFT tx has been send", tx.Hash().String())
}

func (m *MCManager) PickWallet() *bind.TransactOpts {

	nftPrice := util.ToWei("0.001", 18)
	for i := 0; i < len(m.Wallet.PuppetWallets); i++ {

		rand.Seed(time.Now().UnixNano())
		wallet := m.Wallet.PuppetWallets[rand.Intn(len(m.Wallet.PuppetWallets))]
		bnbBalance, err := m.Wallet.BscClient.BalanceAt(context.Background(), wallet.From, nil)
		if err != nil {
			log.Error(err)
			return nil
		}
		if bnbBalance.Cmp(nftPrice) == 1 || bnbBalance.Cmp(nftPrice) == 0 {
			return wallet
		}

	}
	return nil
}
