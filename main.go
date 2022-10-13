package main

import (
	logger "github.com/ipfs/go-log"
	"spike-mc-ops/config"
	"spike-mc-ops/global"
	"spike-mc-ops/initialize"
)

func main() {
	logger.SetLogLevel("*", "INFO")
	global.Viper = config.InitViper()
	//chain.BuyTxListener()
	initialize.RunServer()
}
