package mcMgrSrv

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	logger "github.com/ipfs/go-log"
	"github.com/shopspring/decimal"
	"math/big"
	"math/rand"
	"spike-mc-ops/chain/contract"
	"spike-mc-ops/config"
	"spike-mc-ops/constant"
	"spike-mc-ops/model"
	"spike-mc-ops/request"
	"spike-mc-ops/response"
	"spike-mc-ops/util"
	"strconv"
	"sync"
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
	delete(m.Strategies, u)
}

func (m *MCManager) QueryStrategy(c *gin.Context) {
	var strategyList []response.TradeStrategyResponse
	for key, value := range m.Strategies {
		strategyList = append(strategyList, response.TradeStrategyResponse{
			Uuid:              key,
			BuyPrice:          value.Price[constant.BUY_PRICE],
			SellPrice:         value.Price[constant.SELL_PRICE],
			TotalAmount:       value.Amount[constant.TOTAL_AMOUNT],
			SingleAmount:      value.Amount[constant.SINGLE_AMOUNT],
			ExecTime:          value.ExecTime,
			Frequency:         value.Frequency,
			SlippageTolerance: value.SlippageTolerance,
		})
	}
	log.Infof("startegyList : %v", strategyList)
	if len(strategyList) == 0 {
		response.OkWithData([]response.TradeStrategyResponse{}, c)
		return
	}
	response.OkWithData(strategyList, c)
}

// ExecStrategy Buy or sell recursively according to the strategy.
func (m *MCManager) ExecStrategy(uuid string) error {
	var alreadyUsedQuota int64

	strategy := m.Strategies[uuid]

	ticker := time.NewTicker(time.Duration(strategy.Frequency) * time.Second)
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
			log.Infof("strategy close...")
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
	singer := m.Wallet.PuppetWallets[rand.Intn(len(m.Wallet.PuppetWallets))]

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

		tx, err := m.Wallet.Router.SwapExactTokensForTokens(singer, util.ToWei(strconv.FormatInt(amountInResult, 10), 18), big.NewInt(amountOutResult), path, singer.From, big.NewInt(deadline))
		if err != nil {
			log.Error(err)
			return 0, err
		}
		log.Infof("transaction is already send : %s", tx.Hash().String())

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

func (m *MCManager) QueryWalletBalance(c *gin.Context) {
	var walletInfo []model.WalletInfo
	var lock sync.Mutex
	var wg sync.WaitGroup

	client, err := ethclient.Dial(config.Cfg.Chain.RpcNodeAddress)
	if err != nil {
		log.Errorf("ethClient dial err :%v", err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	gameTokenContract, err := contract.NewErc20Contract(common.HexToAddress(config.Cfg.Contract.GameTokenAddress), client)
	if err != nil {
		log.Errorf("construct gameTokenContract err :%v", err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	usdcContract, err := contract.NewErc20Contract(common.HexToAddress(config.Cfg.Contract.UsdcAddress), client)
	if err != nil {
		log.Errorf("construct usdcContract err :%v", err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	throttle := make(chan struct{}, 5)
	for _, wallet := range m.Wallet.PuppetWallets {
		wg.Add(1)
		throttle <- struct{}{}
		go func(walletAddr common.Address) {
			defer func() {
				wg.Done()
				<-throttle
			}()
			bnbBalance, err := client.BalanceAt(context.Background(), walletAddr, nil)
			if err != nil {
				return
			}
			gameTokenBalance, err := gameTokenContract.BalanceOf(nil, walletAddr)
			if err != nil {
				return
			}
			usdcBalance, err := usdcContract.BalanceOf(nil, walletAddr)
			if err != nil {
				return
			}
			lock.Lock()
			walletInfo = append(walletInfo, model.WalletInfo{
				WalletAddr:  walletAddr.String(),
				BnbBalance:  util.ToDecimal(bnbBalance.String(), 18).String(),
				USDCBalance: util.ToDecimal(usdcBalance.String(), 18).String(),
				SKSBalance:  util.ToDecimal(gameTokenBalance.String(), 18).String(),
			})
			lock.Unlock()
		}(wallet.From)
	}
	wg.Wait()
	response.OkWithData(walletInfo, c)
}
