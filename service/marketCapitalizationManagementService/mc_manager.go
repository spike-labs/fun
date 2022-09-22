package marketCapitalizationManagementService

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	logger "github.com/ipfs/go-log"
	"github.com/shopspring/decimal"
	"math/big"
	contract "spike-mc-ops/chain/contract"
	"spike-mc-ops/config"
	"spike-mc-ops/constant"
	"spike-mc-ops/request"
	"spike-mc-ops/util"
)

var log = logger.Logger("mc")

type MCManager struct {
	BscClient  *ethclient.Client
	Router     *contract.PanCakeRouter
	Strategies map[string]request.TradeStrategy

	uChan chan string
}

func NewMCManager() *MCManager {
	bscClient, err := ethclient.Dial(config.Cfg.Chain.RpcNodeAddress)
	if err != nil {
		panic("=== Spike log: ")
	}

	log.Info(config.Cfg.Contract.PanCakeRouterAddress)
	router, err := contract.NewPanCakeRouter(common.HexToAddress(config.Cfg.Contract.PanCakeRouterAddress), bscClient)
	if err != nil {
		panic("=== Spike log: ")
	}

	m := &MCManager{
		Router:     router,
		BscClient:  bscClient,
		uChan:      make(chan string, 1),
		Strategies: make(map[string]request.TradeStrategy, 0),
	}

	go m.Schedule()

	return m
}

func (m *MCManager) AddStadeStrategy(t request.TradeStrategyService) string {
	u := uuid.New().String()
	tradeStrategy := request.TradeStrategy{
		Price:  make(map[int]string, 0),
		Amount: make(map[int]string, 0),
		Work:   true,
	}

	tradeStrategy.Price[constant.BUY_PRICE] = t.BuyPrice
	tradeStrategy.Price[constant.SELL_PRICE] = t.SellPrice
	tradeStrategy.Amount[constant.TOTAL_AMOUNT] = t.TotalAmount
	tradeStrategy.Amount[constant.SINGLE_AMOUNT] = t.SigleAmount
	tradeStrategy.Frequency = t.Frequency
	tradeStrategy.SlippageTolerance = t.SlippageTolerance

	m.Strategies[u] = tradeStrategy
	m.uChan <- u
	return u
}

func (m *MCManager) Schedule() {
	for {
		select {
		case u := <-m.uChan:
			m.ExecStrategy(u)
		}

	}
}

// ExecStrategy Buy or sell recursively according to the strategy.
func (m *MCManager) ExecStrategy(uuid string) []*big.Int {

	sksPrice, err := m.Router.GetAmountsOut(nil, util.ToWei("1", 18), []common.Address{common.HexToAddress(config.Cfg.Contract.SksAddress), common.HexToAddress(config.Cfg.Contract.UsdcAddress)})
	if err != nil {
		log.Error(err)
		return nil
	}
	strategy := m.Strategies[uuid]

	buyPrice, err := decimal.NewFromString(strategy.Price[constant.BUY_PRICE])
	if err != nil {
		log.Error(err)
		return nil
	}

	sellPrice, err := decimal.NewFromString(strategy.Price[constant.SELL_PRICE])
	if err != nil {
		log.Error(err)
		return nil
	}

	if buyPrice.LessThan(util.ToDecimal(sksPrice[len(sksPrice)-1], 18)) && sellPrice.LessThan(util.ToDecimal(sksPrice[len(sksPrice)-1], 18)) {
		//todo buy op
	} else if buyPrice.GreaterThan(util.ToDecimal(sksPrice[len(sksPrice)-1], 18)) && sellPrice.GreaterThanOrEqual(util.ToDecimal(sksPrice[len(sksPrice)-1], 18)) {
		//todo sell op
	}

	return sksPrice
}

func (m *MCManager) SwapToken() {

}
