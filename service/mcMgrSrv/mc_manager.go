package mcMgrSrv

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	logger "github.com/ipfs/go-log"
	"github.com/shopspring/decimal"
	"math/big"
	"math/rand"
	"spike-mc-ops/config"
	"spike-mc-ops/constant"
	"spike-mc-ops/request"
	"spike-mc-ops/util"
	"strconv"
	"time"
)

var log = logger.Logger("mc")

type MCManager struct {
	Strategies map[string]request.TradeStrategy
	Wallet     *PuppetWallet
}

func NewMCManager() *MCManager {

	puppetWallet := NewPuppetWallet()

	m := &MCManager{
		Wallet:     puppetWallet,
		Strategies: make(map[string]request.TradeStrategy, 0),
	}

	for i := range m.Wallet.PuppetWallets {
		log.Infof("MCManager info :%v", m.Wallet.PuppetWallets[i])
	}

	return m
}

func (m *MCManager) AddTradeStrategy(t request.TradeStrategyService) string {
	u := uuid.New().String()
	tradeStrategy := request.TradeStrategy{
		Price:   make(map[int]string, 0),
		Amount:  make(map[int]string, 0),
		Closing: make(chan struct{}, 0),
		Work:    true,
	}

	tradeStrategy.Price[constant.BUY_PRICE] = t.BuyPrice
	tradeStrategy.Price[constant.SELL_PRICE] = t.SellPrice
	tradeStrategy.Amount[constant.TOTAL_AMOUNT] = t.TotalAmount
	tradeStrategy.Amount[constant.SINGLE_AMOUNT] = t.SingleAmount
	tradeStrategy.Frequency = t.Frequency
	tradeStrategy.SlippageTolerance = t.SlippageTolerance
	tradeStrategy.ExecTime = t.ExecTime

	m.Strategies[u] = tradeStrategy

	go m.ExecStrategy(u)
	return u
}

func (m *MCManager) CancelStrategy(u string) {
	m.Strategies[u].Closing <- struct{}{}
}

// ExecStrategy Buy or sell recursively according to the strategy.
func (m *MCManager) ExecStrategy(uuid string) error {
	var alreadyUsedQuota int64

	strategy := m.Strategies[uuid]

	ticker := time.NewTicker(time.Duration(strategy.Frequency) * time.Minute)
	deadlineTicker := time.NewTicker(time.Duration(strategy.ExecTime) * time.Minute)
	amount, err := strconv.ParseInt(strategy.Amount[constant.TOTAL_AMOUNT], 10, 18)
	if err != nil {
		log.Error(err)
		return err
	}

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
loop:
	for {
		select {
		case <-ticker.C:

			if alreadyUsedQuota >= amount {
				log.Info("strategy fund use done")
				break loop
			}

			sksPrice, err := m.Wallet.Router.GetAmountsOut(nil, util.ToWei("1", 18), []common.Address{common.HexToAddress(config.Cfg.Contract.GameTokenAddress), common.HexToAddress(config.Cfg.Contract.UsdcAddress)})
			if err != nil {
				log.Error(err)
				return nil
			}

			log.Info("buyPrice: ", buyPrice, "sellPrice: ", sellPrice, "sksPrice:", util.ToDecimal(sksPrice[len(sksPrice)-1], 18))

			if buyPrice.GreaterThanOrEqual(util.ToDecimal(sksPrice[len(sksPrice)-1], 18)) && sellPrice.GreaterThanOrEqual(util.ToDecimal(sksPrice[len(sksPrice)-1], 18)) {
				log.Info("buy")
				tokenAmount, err := m.SwapToken(uuid, []common.Address{common.HexToAddress(config.Cfg.Contract.UsdcAddress), common.HexToAddress(config.Cfg.Contract.GameTokenAddress)}, constant.BUY_OP)
				if err != nil {
					log.Error(err)
					break
				}
				alreadyUsedQuota += tokenAmount

			} else if buyPrice.LessThanOrEqual(util.ToDecimal(sksPrice[len(sksPrice)-1], 18)) && sellPrice.LessThanOrEqual(util.ToDecimal(sksPrice[len(sksPrice)-1], 18)) {
				log.Info("sell")
				_, err := m.SwapToken(uuid, []common.Address{common.HexToAddress(config.Cfg.Contract.GameTokenAddress), common.HexToAddress(config.Cfg.Contract.UsdcAddress)}, constant.SELL_OP)
				if err != nil {
					log.Error(err)
					break
				}

			}
		case <-deadlineTicker.C:
			break loop

		case <-strategy.Closing:
			break loop
		}
	}

	return nil
}

func (m *MCManager) SwapToken(uuid string, path []common.Address, flag int) (int64, error) {
	var amountOutResult int64
	var amountOutInternal big.Rat
	var amountInResult int64
	var fund int

	rand.Seed(time.Now().Unix())
	singer := m.Wallet.PuppetWallets[rand.Intn(len(m.Wallet.PuppetWallets)-1)]

	deadline := time.Now().Add(time.Minute * 10).Unix()

	singleMaxAmount, err := strconv.Atoi(m.Strategies[uuid].Amount[constant.SINGLE_AMOUNT])
	if err != nil {
		log.Error(err)
		return 0, err
	}

	if singleMaxAmount-10 <= 0 {
		fund = util.Random(1, singleMaxAmount)
		log.Info("singleMaxAmount: ", singleMaxAmount)
	} else {
		fund = util.Random(singleMaxAmount-10, singleMaxAmount)
	}

	switch flag {
	case constant.BUY_OP:

		amountInResult = int64(fund)

		amountOut, err := m.Wallet.Router.GetAmountsOut(nil, util.ToWei(strconv.FormatInt(amountInResult, 10), 18), path)
		if err != nil {
			log.Error(err)
			return 0, err
		}

		amountOutInternal.Mul(big.NewRat(util.ToDecimal(amountOut[len(amountOut)-1], 18).IntPart(), 1), big.NewRat(int64(1000)-m.Strategies[uuid].SlippageTolerance, int64(1000)))

		log.Info("singleAmount: ", util.ToWei(strconv.FormatInt(amountInResult, 10), 18), "amountOut:", amountOut, "minAmount: ", amountOutInternal.FloatString(0), singer.From, big.NewInt(deadline), "int64:", util.ToDecimal(amountOut[len(amountOut)-1], 18).IntPart())

		amountOutResult, err = strconv.ParseInt(amountOutInternal.FloatString(0), 10, 64)
		if err != nil {
			log.Error(err)
			return 0, err
		}

		tokens, err := m.Wallet.Router.SwapExactTokensForTokens(singer, util.ToWei(strconv.FormatInt(amountInResult, 10), 18), big.NewInt(amountOutResult), path, singer.From, big.NewInt(deadline))
		if err != nil {
			log.Error(err)
			return 0, err
		}
		log.Info("transaction is already send", tokens)

	case constant.SELL_OP:
		var internalPath []common.Address
		internalPath = append(internalPath, path[1])
		internalPath = append(internalPath, path[0])

		amountIn, err := m.Wallet.Router.GetAmountsOut(nil, util.ToWei(strconv.FormatInt(int64(fund), 10), 18), internalPath)
		if err != nil {
			log.Error(err)
			return 0, err
		}

		amountInResult = amountIn[len(amountIn)-1].Int64()
		log.Info(amountIn, "---", amountInResult)
		amountOutInternal.Mul(big.NewRat(int64(fund), 1), big.NewRat(int64(1000)-m.Strategies[uuid].SlippageTolerance, int64(1000)))

		amountOutResult, err = strconv.ParseInt(amountOutInternal.FloatString(0), 10, 64)
		if err != nil {
			log.Error(err)
			return 0, err
		}
		log.Info("amountIn: ", amountIn[len(amountIn)-1], "minAmount: ", big.NewInt(amountOutResult), singer.From, big.NewInt(deadline), "interPath", internalPath)

		tokens, err := m.Wallet.Router.SwapExactTokensForTokens(singer, amountIn[len(amountIn)-1], big.NewInt(amountOutResult), path, singer.From, big.NewInt(deadline))
		if err != nil {
			log.Error(err)
			return 0, err
		}
		log.Info("transaction is already send", tokens)
	}

	return amountInResult, nil
}
