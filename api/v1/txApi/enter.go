package txApi

import (
	"github.com/gin-gonic/gin"
	logger "github.com/ipfs/go-log"
	"spike-mc-ops/request"
	"spike-mc-ops/response"
	"spike-mc-ops/service/marketCapitalizationManagementService"
)

var log = logger.Logger("txApi")

type TxGroup struct {
	m *marketCapitalizationManagementService.MCManager
}

func NewTxGroup() (TxGroup, error) {
	m := marketCapitalizationManagementService.NewMCManager()
	return TxGroup{
		m: m,
	}, nil
}

func (txGroup *TxGroup) InitTxGroup(g *gin.RouterGroup) {
	g.Use()
	trade := g.Group("trade")
	{
		trade.POST("/addStrategy", txGroup.AddStrategy)

	}
}

func (txGroup *TxGroup) AddStrategy(c *gin.Context) {
	var service request.TradeStrategyService
	err := c.ShouldBind(&service)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage("request params error", c)
	}
	uuid := txGroup.m.AddStadeStrategy(service)

	response.OkWithData(uuid, c)
}
