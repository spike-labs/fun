package marketCapitalizationManagementService

import (
	"github.com/ethereum/go-ethereum/ethclient"
	logger "github.com/ipfs/go-log"
)

var log = logger.Logger("mc")

type MCManager struct {
	BscClient *ethclient.Client
}

func NewMCManager() {

}
