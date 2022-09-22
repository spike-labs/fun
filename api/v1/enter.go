package v1

import (
	logger "github.com/ipfs/go-log"
	"spike-mc-ops/api/v1/txApi"
)

var log = logger.Logger("api")

type RouterGroup struct {
	TxGroup txApi.TxGroup
}

func NewRouterGroup() (RouterGroup, error) {
	txGroup, err := txApi.NewTxGroup()
	if err != nil {
		log.Error("===Spike log:", err)
		return RouterGroup{}, err
	}
	return RouterGroup{
		TxGroup: txGroup,
	}, nil
}
