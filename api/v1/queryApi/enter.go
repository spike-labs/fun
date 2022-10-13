package queryApi

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	logger "github.com/ipfs/go-log"
	"spike-mc-ops/chain/contract"
	"spike-mc-ops/config"
	"spike-mc-ops/response"
	"spike-mc-ops/util"
)

var log = logger.Logger("queryApi")

type QueryGroup struct {
}

func NewQueryGroup() (QueryGroup, error) {
	return QueryGroup{}, nil
}

func (qg *QueryGroup) InitQueryGroup(g *gin.RouterGroup) {
	g.Use()

	{
		g.GET("/tokenPrice", qg.QueryGameTokenPrice)
	}
}

func (qg *QueryGroup) QueryGameTokenPrice(c *gin.Context) {
	client, err := ethclient.Dial(config.Cfg.Chain.RpcNodeAddress)
	if err != nil {
		log.Errorf("ethClient dial err : %v", err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	panCakeRouter, err := contract.NewPanCakeRouter(common.HexToAddress(config.Cfg.Contract.PanCakeRouterAddress), client)
	if err != nil {
		log.Errorf("construct panCakeRouter err : %v", err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	txPool, err := contract.NewTxPool(common.HexToAddress(config.Cfg.Contract.GameTokenDexPoolAddress), client)
	if err != nil {
		log.Errorf("construct txPool err : %v", err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	reserves, err := txPool.GetReserves(nil)
	if err != nil {
		log.Errorf("query reserve err : %v", err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	usdcOut, err := panCakeRouter.GetAmountOut(nil, util.ToWei("1", 18), reserves.Reserve0, reserves.Reserve1)
	if err != nil {
		log.Errorf("get amount out err : %v", err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(struct {
		Price string `json:"price"`
	}{Price: "$ " + util.ToDecimal(usdcOut, 18).Rat().FloatString(4)}, c)
}
