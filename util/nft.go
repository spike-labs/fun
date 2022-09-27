package util

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-resty/resty/v2"
	"golang.org/x/xerrors"
	"math/big"
	"spike-mc-ops/chain/contract"
	"spike-mc-ops/config"
	"spike-mc-ops/constant"
	"spike-mc-ops/model"
)

var MoralisRateLimit = "{\"message\":\"Rate limit exceeded.\"}"

func TransferNft(publicKey, privateKey, toAddress string, tokenId int64) (txHash string, err error) {
	ethClient, err := ethclient.Dial(config.Cfg.Chain.RpcNodeAddress)
	if err != nil {
		log.Error(err)
		return "", err
	}
	gameNftContract, err := contract.NewGameNft(common.HexToAddress(config.Cfg.Contract.NftAddress), ethClient)
	if err != nil {
		log.Error(err)
		return "", err
	}
	priKey, err := crypto.HexToECDSA(privateKey[2:])
	if err != nil {
		log.Error(err)
		return "", err
	}

	chainId, err := ethClient.ChainID(context.Background())
	if err != nil {
		log.Error(err)
		return "", err
	}
	log.Infof("chainId : %d", chainId)
	signer, err := bind.NewKeyedTransactorWithChainID(priKey, chainId)
	if err != nil {
		log.Error("New Transactor", err)
		return "", err
	}
	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error(err)
		return "", err
	}
	nonce, err := ethClient.PendingNonceAt(context.Background(), common.HexToAddress(publicKey))
	if err != nil {
		log.Error(err)
		return "", err
	}

	signer.GasPrice = gasPrice
	signer.Nonce = new(big.Int).SetUint64(nonce)
	signer.NoSend = true
	simulatorTx, err := gameNftContract.TransferFrom(signer, common.HexToAddress(publicKey), common.HexToAddress(toAddress), big.NewInt(tokenId))
	if err != nil {
		log.Error(err)
		return "", err
	}
	signer.GasLimit = uint64(float64(simulatorTx.Gas()) * 1.2)
	signer.NoSend = false
	tx, err := gameNftContract.TransferFrom(signer, common.HexToAddress(publicKey), common.HexToAddress(toAddress), big.NewInt(tokenId))
	if err != nil {
		log.Error(err)
		return "", err
	}
	log.Infof("tx : %s", tx.Hash().String())
	return tx.Hash().String(), nil
}

func QueryWalletNft(cursor, walletAddr, network string, res []model.NftResult) ([]model.NftResult, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("x-api-key", config.Cfg.Moralis.XApiKey).
		Get(getQueryNftUrl(config.Cfg.Contract.NftAddress, walletAddr, network, cursor))
	if err != nil {
		log.Errorf("query wallet nft , wallet : %s, err : %+v", walletAddr, err)
		return res, err
	}

	if string(resp.Body()) == MoralisRateLimit {
		log.Errorf(MoralisRateLimit)
		return res, xerrors.New(MoralisRateLimit)
	}

	var nrs model.NftResults
	err = json.Unmarshal(resp.Body(), &nrs)
	if err != nil {
		log.Error("json unmarshal err : ", err)
		return res, err
	}

	res = append(res, nrs.Results...)
	if nrs.Page*nrs.PageSize >= nrs.Total {
		return res, nil
	}
	res, err = QueryWalletNft(nrs.Cursor, walletAddr, network, res)
	return res, nil
}

func getQueryNftUrl(contractAddr, walletAddr, network, cursor string) string {
	return fmt.Sprintf("%s%s/nft/%s?chain=%s&cursor=%s", constant.MORALIS_API, walletAddr, contractAddr, network, cursor)
}
