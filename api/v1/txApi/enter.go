package txApi

import (
	"github.com/gin-gonic/gin"
	logger "github.com/ipfs/go-log"
	"spike-mc-ops/request"
	"spike-mc-ops/response"
	"spike-mc-ops/service/mcMgrSrv"
)

var log = logger.Logger("txApi")

type TxGroup struct {
	m *mcMgrSrv.MCManager
}

func NewTxGroup() (TxGroup, error) {
	m := mcMgrSrv.NewMCManager()
	return TxGroup{
		m: m,
	}, nil
}

func (txGroup *TxGroup) InitTxGroup(g *gin.RouterGroup) {
	g.Use()
	trade := g.Group("trade")
	{
		trade.POST("/addStrategy", txGroup.AddStrategy)
		trade.POST("/cancelStrategy", txGroup.CancelStrategy)
	}
	funds := g.Group("fund")
	{
		funds.POST("/diversifyFunds", txGroup.DiversifyFunds)
		funds.POST("/recoveryFunds", txGroup.RecoveryFunds)
	}
	info := g.Group("info")
	{
		info.GET("/walletBalance", txGroup.m.QueryWalletBalance)
		info.GET("/queryStrategy", txGroup.m.QueryStrategy)
	}
}

func (txGroup *TxGroup) AddStrategy(c *gin.Context) {
	var service request.TradeStrategyService
	err := c.ShouldBindJSON(&service)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage("request params error", c)
		return
	}
	log.Infof("params: %v", service)
	uuid := txGroup.m.AddTradeStrategy(service)

	response.OkWithData(uuid, c)
}

func (txGroup *TxGroup) CancelStrategy(c *gin.Context) {
	var service request.DelTradeStrategyService
	err := c.ShouldBindJSON(&service)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage("request params error", c)
		return
	}
	txGroup.m.CancelStrategy(service.Uuid)

	response.Ok(c)
}

func (txGroup *TxGroup) DiversifyFunds(c *gin.Context) {
	var service request.DiversifyFundsService
	err := c.ShouldBindJSON(&service)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage("request params error", c)
		return
	}
	log.Infof("params: %v", service)
	err = txGroup.m.Wallet.DiversifyFunds(service)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.Fail(c)
		return
	}

	response.Ok(c)
}

func (txGroup *TxGroup) RecoveryFunds(c *gin.Context) {
	var service request.RecoveryFundsService
	err := c.ShouldBindJSON(&service)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage("request params error", c)
		return
	}
	log.Infof("params: %v", service)
	err = txGroup.m.Wallet.RecoveryFunds(service)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.Fail(c)
		return
	}

	response.Ok(c)
}
