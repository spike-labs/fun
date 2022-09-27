package chain

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	logger "github.com/ipfs/go-log"
	"math/big"
	"spike-mc-ops/chain/contract"
	"spike-mc-ops/config"
	"spike-mc-ops/util"
	"strings"
	"time"
)

var log = logger.Logger("chain")

func FilterPancakeLog() {
	ethClient, err := ethclient.Dial(config.Cfg.Chain.WsNodeAddress)

	if err != nil {
		log.Errorf("dial err : %v", err)
		return
	}
	contractAddress := common.HexToAddress(config.Cfg.Contract.GameTokenDexPoolAddress)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	ch := make(chan types.Log)
	sub, err := ethClient.SubscribeFilterLogs(context.Background(), query, ch)
	if err != nil {
		log.Errorf("nft subscribe event log  err : %v", err)
		return
	}
	for {
		select {
		case err := <-sub.Err():
			log.Errorf("SubscribeFilterLogs : err : %v", err)
			sub = event.Resubscribe(time.Millisecond, func(ctx context.Context) (event.Subscription, error) {
				return ethClient.SubscribeFilterLogs(ctx, query, ch)
			})
		case l := <-ch:
			switch l.Topics[0].String() {
			case util.EventSignHash("Swap(address,uint256,uint256,uint256,uint256,address)"):
				abi, _ := abi.JSON(strings.NewReader(contract.TxPoolMetaData.ABI))

				inputDatas, err := abi.Events["Swap"].Inputs.Unpack(l.Data)

				if err != nil {
					log.Errorf("err : %v", err)
					return
				}
				sender := common.HexToAddress(l.Topics[1].Hex()).String()
				toAddr := common.HexToAddress(l.Topics[2].Hex()).String()
				amount0In := inputDatas[0].(*big.Int)
				amount1In := inputDatas[1].(*big.Int)
				amount0Out := inputDatas[2].(*big.Int)
				amount1Out := inputDatas[3].(*big.Int)
				log.Infof("sender : %s, toAddr : %s, amount0In : %s,  amount1In : %s, amount0Out : %v, amount1Out : %s", sender, toAddr, amount0In.String(), amount1In.String(), amount0Out.String(), amount1Out.String())
			}
		}
	}
}
