package initialize

import (
	"github.com/gin-gonic/gin"
	logger "github.com/ipfs/go-log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "spike-mc-ops/api/v1"
	"spike-mc-ops/middleware"
)

var log = logger.Logger("initialize")

func initRouter() (*gin.Engine, error) {
	var r = gin.Default()
	r.Use(middleware.Cors())
	r.Use(middleware.TokenAuth())
	publicGroup := r.Group("")
	{
		// health
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	routerGroupApp, err := v1.NewRouterGroup()
	if err != nil {
		return nil, err
	}

	txGroup := routerGroupApp.TxGroup
	txApiGroup := r.Group("/trade-api/v1")
	txGroup.InitTxGroup(txApiGroup)

	queryGroup := routerGroupApp.QueryGroup
	queryApiGroup := r.Group("/query-api/v1")
	queryGroup.InitQueryGroup(queryApiGroup)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r, nil
}
