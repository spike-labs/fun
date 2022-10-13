package mcMgrSrv

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"math/rand"
	contract "spike-mc-ops/chain/contract"
	"spike-mc-ops/config"
	"spike-mc-ops/request"
	"spike-mc-ops/util"
	"sync"
	"time"
)

type PuppetWallet struct {
	PuppetWallets []*bind.TransactOpts
	BscClient     *ethclient.Client
	Router        *contract.PanCakeRouter
}

func NewPuppetWallet() *PuppetWallet {

	bscClient, err := ethclient.Dial(config.Cfg.Chain.RpcNodeAddress)
	if err != nil {
		panic("=== Spike log: ")
	}

	router, err := contract.NewPanCakeRouter(common.HexToAddress(config.Cfg.Contract.PanCakeRouterAddress), bscClient)
	if err != nil {
		panic("=== Spike log: ")
	}

	id, err := bscClient.ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	puppetWallets := make([]*bind.TransactOpts, 0)

	for i := 0; i < config.Cfg.PuppetWallet.AccountNumber; i++ {
		priKeyString, err := util.DerivePrivateKeyWithNumber(config.Cfg.PuppetWallet.Mnemonic, i)
		if err != nil {
			panic(err)
		}

		priKey, err := crypto.HexToECDSA(priKeyString[2:])
		if err != nil {
			panic(err)
		}

		opts, err := bind.NewKeyedTransactorWithChainID(priKey, id)
		if err != nil {
			panic(err)
		}

		puppetWallets = append(puppetWallets, opts)

		util.CheckAllowance(priKeyString, config.Cfg.Contract.UsdcAddress, config.Cfg.Contract.PanCakeRouterAddress, id, bscClient)
		util.CheckAllowance(priKeyString, config.Cfg.Contract.GameTokenAddress, config.Cfg.Contract.PanCakeRouterAddress, id, bscClient)

	}

	return &PuppetWallet{
		BscClient:     bscClient,
		Router:        router,
		PuppetWallets: puppetWallets,
	}
}

type WalletTree struct {
	Val   *bind.TransactOpts
	Child []*WalletTree
}

func CreateMultipleTree(dataLst []*bind.TransactOpts) *WalletTree {
	rand.Seed(time.Now().Unix())

	// 剔除第一个账户 原始账户
	data := make([]*bind.TransactOpts, len(dataLst))
	copy(data, dataLst)

	data = append(data[:0], data[1:]...)

	leftWallet := config.Cfg.PuppetWallet.AccountNumber

	treeItem := make([]*WalletTree, 0)
	treeRoot := &WalletTree{Val: dataLst[0], Child: make([]*WalletTree, 0)}
	treeItem = append(treeItem, treeRoot)

	for {
		if leftWallet == 0 || len(data) == 0 {
			return treeRoot
		}

		childNum := util.Random(2, 5)
		for i := 0; i < childNum; i++ {
			if len(data) == 0 || len(treeItem) == 0 {
				return treeRoot
			}
			leftWallet--
			index := rand.Intn(len(data))
			child := &WalletTree{Val: data[index], Child: make([]*WalletTree, 0)}
			treeItem[0].Child = append(treeItem[0].Child, child)
			treeItem = append(treeItem, child)
			data = append(data[:index], data[index+1:]...)
		}

		treeItem = append(treeItem[:0], treeItem[1:]...)
	}
}

// CountAmount Recursively query the total number below the minimum limit under the node.
func (p *PuppetWallet) CountAmount(root *WalletTree, contractAddress string, minBalance string, total *big.Int) error {
	var wg sync.WaitGroup
	if root == nil {
		return nil
	}
	minBalWei := util.ToWei(minBalance, 18)

	if contractAddress != "" {
		// erc20 balance
		erc20Contract, err := contract.NewErc20Contract(common.HexToAddress(contractAddress), p.BscClient)
		if err != nil {
			log.Error(err)
			return err
		}

		balance, err := erc20Contract.BalanceOf(nil, root.Val.From)
		if err != nil {
			log.Error(err)
			return err
		}

		log.Info("wallet address: ", root.Val.From, " erc20 balance:", util.ToDecimal(balance, 18).String(), "-", "minBalance: ", minBalance)

		if balance.Cmp(minBalWei) == -1 {
			balance.Sub(minBalWei, balance)
			total.Add(total, balance)
		}
	} else {
		// bnb balance
		balance, err := p.BscClient.BalanceAt(context.Background(), root.Val.From, nil)
		if err != nil {
			log.Error(err)
			return err
		}

		if balance.Cmp(minBalWei) == -1 {
			balance.Sub(minBalWei, balance)
			total.Add(total, balance)
			total.Add(total, util.ToWei("0.001", 18))
		}

	}

	for i := range root.Child {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			err := p.CountAmount(root.Child[index], contractAddress, minBalance, total)
			if err != nil {
				log.Error(err)
			}
		}(i)
	}

	wg.Wait()
	return nil
}

func (p *PuppetWallet) DiversifyFunds(request request.DiversifyFundsService) error {
	mulTreeRoot := CreateMultipleTree(p.PuppetWallets)
	PrintTree(mulTreeRoot)
	var err error
	totalAmount := big.NewInt(0)
	mainAccountBalance := big.NewInt(0)

	// query main Account Balance ,bnb or erc20
	if request.ContractAddress == "" {
		mainAccountBalance, err = p.BscClient.BalanceAt(context.Background(), mulTreeRoot.Val.From, nil)
		if err != nil {
			log.Error(err)
			return err
		}
	} else {
		erc20Contract, err := contract.NewErc20Contract(common.HexToAddress(request.ContractAddress), p.BscClient)
		if err != nil {
			log.Error(err)
			return err
		}

		mainAccountBalance, err = erc20Contract.BalanceOf(nil, mulTreeRoot.Val.From)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	// query the whole tree, need  fund number
	err = p.CountAmount(mulTreeRoot, request.ContractAddress, request.MinBalance, totalAmount)
	if err != nil {
		log.Error(err)
		return err
	}

	// check main Account Balance ,if main Account Balance is not enough , notify user
	if totalAmount.Cmp(mainAccountBalance) == 1 {
		mainAccountBalance.Sub(totalAmount, mainAccountBalance)
		return errors.New("mainAccount: " + p.PuppetWallets[0].From.String() + "balance is not enough, Please add" + util.ToDecimal(mainAccountBalance.String(), 18).String())
	}

	err = p.RecursionDiversifyFunds(request.ContractAddress, request.MinBalance, mulTreeRoot)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (p *PuppetWallet) RecursionDiversifyFunds(contractAddress string, MinBalance string, mulTreeItem *WalletTree) error {
	if mulTreeItem == nil {
		return nil
	}

	for i := range mulTreeItem.Child {

		log.Info("parentWalletAddress: ", mulTreeItem.Val.From, "walletAddress: ", mulTreeItem.Child[i].Val.From)

		totalAmount := big.NewInt(0)
		err := p.CountAmount(mulTreeItem.Child[i], contractAddress, MinBalance, totalAmount)
		if err != nil {
			log.Error(err)
			return err
		}

		if totalAmount.Cmp(util.ToWei("0", 18)) == 0 {
			log.Info("wallet balance is enough")
			continue
		}

		log.Info("from: ", mulTreeItem.Val, "to: ", mulTreeItem.Child[i].Val, "amount", util.ToDecimal(totalAmount, 18).String())
		err = p.SendFunds(contractAddress, totalAmount, mulTreeItem.Val, mulTreeItem.Child[i].Val)
		if err != nil {
			log.Error(err)
			return err
		}
		//  Monitor the success of the transaction or sleep directly before proceeding to the next step.
		time.Sleep(time.Second * 6)

		go func(index int) {
			err = p.RecursionDiversifyFunds(contractAddress, MinBalance, mulTreeItem.Child[index])
			if err != nil {
				log.Error(err)
			}
		}(i)

	}
	return nil
}

func (p *PuppetWallet) SendFunds(contractAddress string, sendAmount *big.Int, fromWallet *bind.TransactOpts, toWallet *bind.TransactOpts) error {

	if contractAddress != "" {
		// erc20 balance
		erc20Contract, err := contract.NewErc20Contract(common.HexToAddress(contractAddress), p.BscClient)
		if err != nil {
			log.Error(err)
			return err
		}

		signedTx, err := erc20Contract.Transfer(fromWallet, toWallet.From, sendAmount)
		if err != nil {
			log.Error(err)
			return err
		}
		log.Info("tx has been send ; hash", signedTx.Hash(), "from: ", fromWallet.From, "to: ", toWallet.From, "amount: ", util.ToDecimal(sendAmount, 18).String())
		return nil
	}

	nonce, err := p.BscClient.PendingNonceAt(context.Background(), fromWallet.From)
	if err != nil {
		log.Error(err)
		return err
	}
	gasPrice, err := p.BscClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error(err)
		return err
	}

	gasLimit, err := p.BscClient.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  fromWallet.From,
		To:    &toWallet.From,
		Value: sendAmount,
	})
	if err != nil {
		log.Error(err)
		return err
	}

	signedTx, err := fromWallet.Signer(fromWallet.From, types.NewTx(
		&types.LegacyTx{
			Nonce:    nonce,
			Gas:      gasLimit,
			GasPrice: gasPrice,
			Value:    sendAmount,
			To:       &toWallet.From,
		}))
	if err != nil {
		log.Error(err)
		return err
	}

	err = p.BscClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("tx has been send ; hash", signedTx.Hash(), "from: ", fromWallet.From, "to: ", toWallet.From, "amount: ", util.ToDecimal(sendAmount, 18).String())
	return nil
}

func (p *PuppetWallet) RecoveryFunds(request request.RecoveryFundsService) error {

	switch request.ContractAddress {
	case "":

		for i := 1; i < len(p.PuppetWallets); i++ {
			go func(index int) {

				gas := util.ToWei("0.0005", 18)

				bnbBalance, err := p.BscClient.BalanceAt(context.Background(), p.PuppetWallets[index].From, nil)
				if err != nil {
					log.Error(err)
					return
				}

				if bnbBalance.Cmp(gas) == -1 {
					return
				}

				sendBalance := big.NewInt(0)
				sendBalance.Sub(bnbBalance, gas)

				err = p.SendFunds(request.ContractAddress, sendBalance, p.PuppetWallets[index], p.PuppetWallets[0])
				if err != nil {
					log.Error(err)
					return
				}

			}(i)
		}

		break
	default:
		erc20Contract, err := contract.NewErc20Contract(common.HexToAddress(request.ContractAddress), p.BscClient)
		if err != nil {
			log.Error(err)
			return err
		}

		for i := 1; i < len(p.PuppetWallets); i++ {

			go func(index int, contract *contract.Erc20Contract) {

				tokenBalance, err := contract.BalanceOf(nil, p.PuppetWallets[index].From)
				if err != nil {
					log.Error(err)
					return
				}

				if tokenBalance.Cmp(big.NewInt(1)) == -1 {
					return
				}

				err = p.SendFunds(request.ContractAddress, tokenBalance, p.PuppetWallets[index], p.PuppetWallets[0])
				if err != nil {
					log.Error(err)
					return
				}
			}(i, erc20Contract)

		}
	}
	return nil
}

func PrintTree(s *WalletTree) {
	if s.Val == nil {
		return
	}
	log.Info("root node: ", s.Val.From)
	if len(s.Child) == 0 {
		return
	}
	for i := range s.Child {
		log.Info("child node: ", s.Child[i].Val.From, "-")
	}
	fmt.Println()
	if len(s.Child) != 0 {
		for i := range s.Child {
			PrintTree(s.Child[i])
		}
	}
}
