// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// NFTDistributionTicker is an auto generated low-level Go binding around an user-defined struct.
type NFTDistributionTicker struct {
	Buyer     common.Address
	NftId     *big.Int
	Timestamp *big.Int
}

// NFTDistributionUserInfo is an auto generated low-level Go binding around an user-defined struct.
type NFTDistributionUserInfo struct {
	User               common.Address
	BuyCount           *big.Int
	InvitesCount       *big.Int
	Commission         *big.Int
	InTopCommissions   bool
	TopCommissionIndex *big.Int
}

// NFTDistributionMetaData contains all meta data concerning the NFTDistribution contract.
var NFTDistributionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fundMgr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nftToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"NotWhitelisted\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"inviter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"PreSaleClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"inviter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"PreSaled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"RolledOver\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds\",\"type\":\"uint256[]\"}],\"name\":\"claimPreSaledNfts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inviter\",\"type\":\"address\"}],\"name\":\"claimWithInviter\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inviter\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"proofRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofPath\",\"type\":\"bytes32[]\"}],\"name\":\"claimWithProof\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"commonNftTickers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discountRate\",\"type\":\"uint256\"}],\"name\":\"configPreSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configuredTopCommissionNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTokenId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultRateRule\",\"outputs\":[{\"internalType\":\"contractIRateRule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flipSaleStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fromTokenId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundMgr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"invitesCount\",\"type\":\"uint128\"}],\"name\":\"getCommissionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"inviter\",\"type\":\"address\"},{\"internalType\":\"contractIRateRule\",\"name\":\"currentRateRule\",\"type\":\"address\"}],\"name\":\"getCurrentSalePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastestBoughtUncommonNfts\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structNFTDistribution.Ticker[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestBoughtCommonNfts\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structNFTDistribution.Ticker[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inviter\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"proofRoot\",\"type\":\"bytes32\"}],\"name\":\"getSalePriceDiscount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTopCommissions\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"buyCount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"invitesCount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"inTopCommissions\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"topCommissionIndex\",\"type\":\"uint256\"}],\"internalType\":\"structNFTDistribution.UserInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserPreSaledNfts\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPreSaleTime\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftToken\",\"outputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preSaleConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"discountRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"preSaledNftClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"preSaledNftOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"preSaledNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"saleActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"salePrice\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRateRule\",\"name\":\"newRateRule\",\"type\":\"address\"}],\"name\":\"setDefaultRateRule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_salePrice\",\"type\":\"uint128\"}],\"name\":\"setSalePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proofRoot\",\"type\":\"bytes32\"},{\"internalType\":\"contractIRateRule\",\"name\":\"newRateRule\",\"type\":\"address\"}],\"name\":\"setSpecialRateRule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_fromTokenId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_toTokenId\",\"type\":\"uint32\"}],\"name\":\"setTokenScope\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"newTradeToken\",\"type\":\"address\"}],\"name\":\"setTradeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"isUncommon\",\"type\":\"bool\"}],\"name\":\"setUncommonNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"buyCount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"invitesCount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"inTopCommissions\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"topCommissionIndex\",\"type\":\"uint256\"}],\"internalType\":\"structNFTDistribution.UserInfo[]\",\"name\":\"_input\",\"type\":\"tuple[]\"}],\"name\":\"sort\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"buyCount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"invitesCount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"inTopCommissions\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"topCommissionIndex\",\"type\":\"uint256\"}],\"internalType\":\"structNFTDistribution.UserInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"specialRateRule\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"contractIRateRule\",\"name\":\"rateRule\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toTokenId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"topCommissions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"buyCount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"invitesCount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"inTopCommissions\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"topCommissionIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tradeToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uncommonNft\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uncommonNftTickers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFundMgr\",\"type\":\"address\"}],\"name\":\"updateFundMgr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"buyCount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"invitesCount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"inTopCommissions\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"topCommissionIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"proofRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proofPath\",\"type\":\"bytes32[]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NFTDistributionABI is the input ABI used to generate the binding from.
// Deprecated: Use NFTDistributionMetaData.ABI instead.
var NFTDistributionABI = NFTDistributionMetaData.ABI

// NFTDistribution is an auto generated Go binding around an Ethereum contract.
type NFTDistribution struct {
	NFTDistributionCaller     // Read-only binding to the contract
	NFTDistributionTransactor // Write-only binding to the contract
	NFTDistributionFilterer   // Log filterer for contract events
}

// NFTDistributionCaller is an auto generated read-only Go binding around an Ethereum contract.
type NFTDistributionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTDistributionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NFTDistributionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTDistributionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NFTDistributionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTDistributionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NFTDistributionSession struct {
	Contract     *NFTDistribution  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTDistributionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NFTDistributionCallerSession struct {
	Contract *NFTDistributionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// NFTDistributionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NFTDistributionTransactorSession struct {
	Contract     *NFTDistributionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NFTDistributionRaw is an auto generated low-level Go binding around an Ethereum contract.
type NFTDistributionRaw struct {
	Contract *NFTDistribution // Generic contract binding to access the raw methods on
}

// NFTDistributionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NFTDistributionCallerRaw struct {
	Contract *NFTDistributionCaller // Generic read-only contract binding to access the raw methods on
}

// NFTDistributionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NFTDistributionTransactorRaw struct {
	Contract *NFTDistributionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNFTDistribution creates a new instance of NFTDistribution, bound to a specific deployed contract.
func NewNFTDistribution(address common.Address, backend bind.ContractBackend) (*NFTDistribution, error) {
	contract, err := bindNFTDistribution(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NFTDistribution{NFTDistributionCaller: NFTDistributionCaller{contract: contract}, NFTDistributionTransactor: NFTDistributionTransactor{contract: contract}, NFTDistributionFilterer: NFTDistributionFilterer{contract: contract}}, nil
}

// NewNFTDistributionCaller creates a new read-only instance of NFTDistribution, bound to a specific deployed contract.
func NewNFTDistributionCaller(address common.Address, caller bind.ContractCaller) (*NFTDistributionCaller, error) {
	contract, err := bindNFTDistribution(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NFTDistributionCaller{contract: contract}, nil
}

// NewNFTDistributionTransactor creates a new write-only instance of NFTDistribution, bound to a specific deployed contract.
func NewNFTDistributionTransactor(address common.Address, transactor bind.ContractTransactor) (*NFTDistributionTransactor, error) {
	contract, err := bindNFTDistribution(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NFTDistributionTransactor{contract: contract}, nil
}

// NewNFTDistributionFilterer creates a new log filterer instance of NFTDistribution, bound to a specific deployed contract.
func NewNFTDistributionFilterer(address common.Address, filterer bind.ContractFilterer) (*NFTDistributionFilterer, error) {
	contract, err := bindNFTDistribution(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NFTDistributionFilterer{contract: contract}, nil
}

// bindNFTDistribution binds a generic wrapper to an already deployed contract.
func bindNFTDistribution(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NFTDistributionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTDistribution *NFTDistributionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTDistribution.Contract.NFTDistributionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTDistribution *NFTDistributionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTDistribution.Contract.NFTDistributionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTDistribution *NFTDistributionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTDistribution.Contract.NFTDistributionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTDistribution *NFTDistributionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTDistribution.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTDistribution *NFTDistributionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTDistribution.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTDistribution *NFTDistributionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTDistribution.Contract.contract.Transact(opts, method, params...)
}

// CommonNftTickers is a free data retrieval call binding the contract method 0x1fc3f9bc.
//
// Solidity: function commonNftTickers(uint256 ) view returns(address buyer, uint256 nftId, uint256 timestamp)
func (_NFTDistribution *NFTDistributionCaller) CommonNftTickers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Buyer     common.Address
	NftId     *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "commonNftTickers", arg0)

	outstruct := new(struct {
		Buyer     common.Address
		NftId     *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Buyer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CommonNftTickers is a free data retrieval call binding the contract method 0x1fc3f9bc.
//
// Solidity: function commonNftTickers(uint256 ) view returns(address buyer, uint256 nftId, uint256 timestamp)
func (_NFTDistribution *NFTDistributionSession) CommonNftTickers(arg0 *big.Int) (struct {
	Buyer     common.Address
	NftId     *big.Int
	Timestamp *big.Int
}, error) {
	return _NFTDistribution.Contract.CommonNftTickers(&_NFTDistribution.CallOpts, arg0)
}

// CommonNftTickers is a free data retrieval call binding the contract method 0x1fc3f9bc.
//
// Solidity: function commonNftTickers(uint256 ) view returns(address buyer, uint256 nftId, uint256 timestamp)
func (_NFTDistribution *NFTDistributionCallerSession) CommonNftTickers(arg0 *big.Int) (struct {
	Buyer     common.Address
	NftId     *big.Int
	Timestamp *big.Int
}, error) {
	return _NFTDistribution.Contract.CommonNftTickers(&_NFTDistribution.CallOpts, arg0)
}

// ConfiguredTopCommissionNum is a free data retrieval call binding the contract method 0x7a9c46ad.
//
// Solidity: function configuredTopCommissionNum() view returns(uint256)
func (_NFTDistribution *NFTDistributionCaller) ConfiguredTopCommissionNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "configuredTopCommissionNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConfiguredTopCommissionNum is a free data retrieval call binding the contract method 0x7a9c46ad.
//
// Solidity: function configuredTopCommissionNum() view returns(uint256)
func (_NFTDistribution *NFTDistributionSession) ConfiguredTopCommissionNum() (*big.Int, error) {
	return _NFTDistribution.Contract.ConfiguredTopCommissionNum(&_NFTDistribution.CallOpts)
}

// ConfiguredTopCommissionNum is a free data retrieval call binding the contract method 0x7a9c46ad.
//
// Solidity: function configuredTopCommissionNum() view returns(uint256)
func (_NFTDistribution *NFTDistributionCallerSession) ConfiguredTopCommissionNum() (*big.Int, error) {
	return _NFTDistribution.Contract.ConfiguredTopCommissionNum(&_NFTDistribution.CallOpts)
}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint32)
func (_NFTDistribution *NFTDistributionCaller) CurrentTokenId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "currentTokenId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint32)
func (_NFTDistribution *NFTDistributionSession) CurrentTokenId() (uint32, error) {
	return _NFTDistribution.Contract.CurrentTokenId(&_NFTDistribution.CallOpts)
}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint32)
func (_NFTDistribution *NFTDistributionCallerSession) CurrentTokenId() (uint32, error) {
	return _NFTDistribution.Contract.CurrentTokenId(&_NFTDistribution.CallOpts)
}

// DefaultRateRule is a free data retrieval call binding the contract method 0x5d4aa1e4.
//
// Solidity: function defaultRateRule() view returns(address)
func (_NFTDistribution *NFTDistributionCaller) DefaultRateRule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "defaultRateRule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultRateRule is a free data retrieval call binding the contract method 0x5d4aa1e4.
//
// Solidity: function defaultRateRule() view returns(address)
func (_NFTDistribution *NFTDistributionSession) DefaultRateRule() (common.Address, error) {
	return _NFTDistribution.Contract.DefaultRateRule(&_NFTDistribution.CallOpts)
}

// DefaultRateRule is a free data retrieval call binding the contract method 0x5d4aa1e4.
//
// Solidity: function defaultRateRule() view returns(address)
func (_NFTDistribution *NFTDistributionCallerSession) DefaultRateRule() (common.Address, error) {
	return _NFTDistribution.Contract.DefaultRateRule(&_NFTDistribution.CallOpts)
}

// FromTokenId is a free data retrieval call binding the contract method 0x112e0dee.
//
// Solidity: function fromTokenId() view returns(uint32)
func (_NFTDistribution *NFTDistributionCaller) FromTokenId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "fromTokenId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FromTokenId is a free data retrieval call binding the contract method 0x112e0dee.
//
// Solidity: function fromTokenId() view returns(uint32)
func (_NFTDistribution *NFTDistributionSession) FromTokenId() (uint32, error) {
	return _NFTDistribution.Contract.FromTokenId(&_NFTDistribution.CallOpts)
}

// FromTokenId is a free data retrieval call binding the contract method 0x112e0dee.
//
// Solidity: function fromTokenId() view returns(uint32)
func (_NFTDistribution *NFTDistributionCallerSession) FromTokenId() (uint32, error) {
	return _NFTDistribution.Contract.FromTokenId(&_NFTDistribution.CallOpts)
}

// FundMgr is a free data retrieval call binding the contract method 0x4b8539ea.
//
// Solidity: function fundMgr() view returns(address)
func (_NFTDistribution *NFTDistributionCaller) FundMgr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "fundMgr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FundMgr is a free data retrieval call binding the contract method 0x4b8539ea.
//
// Solidity: function fundMgr() view returns(address)
func (_NFTDistribution *NFTDistributionSession) FundMgr() (common.Address, error) {
	return _NFTDistribution.Contract.FundMgr(&_NFTDistribution.CallOpts)
}

// FundMgr is a free data retrieval call binding the contract method 0x4b8539ea.
//
// Solidity: function fundMgr() view returns(address)
func (_NFTDistribution *NFTDistributionCallerSession) FundMgr() (common.Address, error) {
	return _NFTDistribution.Contract.FundMgr(&_NFTDistribution.CallOpts)
}

// GetCommissionRate is a free data retrieval call binding the contract method 0xfdc6739b.
//
// Solidity: function getCommissionRate(uint128 invitesCount) view returns(uint256)
func (_NFTDistribution *NFTDistributionCaller) GetCommissionRate(opts *bind.CallOpts, invitesCount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "getCommissionRate", invitesCount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommissionRate is a free data retrieval call binding the contract method 0xfdc6739b.
//
// Solidity: function getCommissionRate(uint128 invitesCount) view returns(uint256)
func (_NFTDistribution *NFTDistributionSession) GetCommissionRate(invitesCount *big.Int) (*big.Int, error) {
	return _NFTDistribution.Contract.GetCommissionRate(&_NFTDistribution.CallOpts, invitesCount)
}

// GetCommissionRate is a free data retrieval call binding the contract method 0xfdc6739b.
//
// Solidity: function getCommissionRate(uint128 invitesCount) view returns(uint256)
func (_NFTDistribution *NFTDistributionCallerSession) GetCommissionRate(invitesCount *big.Int) (*big.Int, error) {
	return _NFTDistribution.Contract.GetCommissionRate(&_NFTDistribution.CallOpts, invitesCount)
}

// GetCurrentSalePrice is a free data retrieval call binding the contract method 0x7e2ada68.
//
// Solidity: function getCurrentSalePrice(address buyer, address inviter, address currentRateRule) view returns(uint256)
func (_NFTDistribution *NFTDistributionCaller) GetCurrentSalePrice(opts *bind.CallOpts, buyer common.Address, inviter common.Address, currentRateRule common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "getCurrentSalePrice", buyer, inviter, currentRateRule)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentSalePrice is a free data retrieval call binding the contract method 0x7e2ada68.
//
// Solidity: function getCurrentSalePrice(address buyer, address inviter, address currentRateRule) view returns(uint256)
func (_NFTDistribution *NFTDistributionSession) GetCurrentSalePrice(buyer common.Address, inviter common.Address, currentRateRule common.Address) (*big.Int, error) {
	return _NFTDistribution.Contract.GetCurrentSalePrice(&_NFTDistribution.CallOpts, buyer, inviter, currentRateRule)
}

// GetCurrentSalePrice is a free data retrieval call binding the contract method 0x7e2ada68.
//
// Solidity: function getCurrentSalePrice(address buyer, address inviter, address currentRateRule) view returns(uint256)
func (_NFTDistribution *NFTDistributionCallerSession) GetCurrentSalePrice(buyer common.Address, inviter common.Address, currentRateRule common.Address) (*big.Int, error) {
	return _NFTDistribution.Contract.GetCurrentSalePrice(&_NFTDistribution.CallOpts, buyer, inviter, currentRateRule)
}

// GetLastestBoughtUncommonNfts is a free data retrieval call binding the contract method 0x1eeeb41a.
//
// Solidity: function getLastestBoughtUncommonNfts() view returns((address,uint256,uint256)[])
func (_NFTDistribution *NFTDistributionCaller) GetLastestBoughtUncommonNfts(opts *bind.CallOpts) ([]NFTDistributionTicker, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "getLastestBoughtUncommonNfts")

	if err != nil {
		return *new([]NFTDistributionTicker), err
	}

	out0 := *abi.ConvertType(out[0], new([]NFTDistributionTicker)).(*[]NFTDistributionTicker)

	return out0, err

}

// GetLastestBoughtUncommonNfts is a free data retrieval call binding the contract method 0x1eeeb41a.
//
// Solidity: function getLastestBoughtUncommonNfts() view returns((address,uint256,uint256)[])
func (_NFTDistribution *NFTDistributionSession) GetLastestBoughtUncommonNfts() ([]NFTDistributionTicker, error) {
	return _NFTDistribution.Contract.GetLastestBoughtUncommonNfts(&_NFTDistribution.CallOpts)
}

// GetLastestBoughtUncommonNfts is a free data retrieval call binding the contract method 0x1eeeb41a.
//
// Solidity: function getLastestBoughtUncommonNfts() view returns((address,uint256,uint256)[])
func (_NFTDistribution *NFTDistributionCallerSession) GetLastestBoughtUncommonNfts() ([]NFTDistributionTicker, error) {
	return _NFTDistribution.Contract.GetLastestBoughtUncommonNfts(&_NFTDistribution.CallOpts)
}

// GetLatestBoughtCommonNfts is a free data retrieval call binding the contract method 0xdd027d9b.
//
// Solidity: function getLatestBoughtCommonNfts() view returns((address,uint256,uint256)[])
func (_NFTDistribution *NFTDistributionCaller) GetLatestBoughtCommonNfts(opts *bind.CallOpts) ([]NFTDistributionTicker, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "getLatestBoughtCommonNfts")

	if err != nil {
		return *new([]NFTDistributionTicker), err
	}

	out0 := *abi.ConvertType(out[0], new([]NFTDistributionTicker)).(*[]NFTDistributionTicker)

	return out0, err

}

// GetLatestBoughtCommonNfts is a free data retrieval call binding the contract method 0xdd027d9b.
//
// Solidity: function getLatestBoughtCommonNfts() view returns((address,uint256,uint256)[])
func (_NFTDistribution *NFTDistributionSession) GetLatestBoughtCommonNfts() ([]NFTDistributionTicker, error) {
	return _NFTDistribution.Contract.GetLatestBoughtCommonNfts(&_NFTDistribution.CallOpts)
}

// GetLatestBoughtCommonNfts is a free data retrieval call binding the contract method 0xdd027d9b.
//
// Solidity: function getLatestBoughtCommonNfts() view returns((address,uint256,uint256)[])
func (_NFTDistribution *NFTDistributionCallerSession) GetLatestBoughtCommonNfts() ([]NFTDistributionTicker, error) {
	return _NFTDistribution.Contract.GetLatestBoughtCommonNfts(&_NFTDistribution.CallOpts)
}

// GetSalePriceDiscount is a free data retrieval call binding the contract method 0x7d31dad9.
//
// Solidity: function getSalePriceDiscount(address inviter, bytes32 proofRoot) view returns(uint256)
func (_NFTDistribution *NFTDistributionCaller) GetSalePriceDiscount(opts *bind.CallOpts, inviter common.Address, proofRoot [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "getSalePriceDiscount", inviter, proofRoot)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSalePriceDiscount is a free data retrieval call binding the contract method 0x7d31dad9.
//
// Solidity: function getSalePriceDiscount(address inviter, bytes32 proofRoot) view returns(uint256)
func (_NFTDistribution *NFTDistributionSession) GetSalePriceDiscount(inviter common.Address, proofRoot [32]byte) (*big.Int, error) {
	return _NFTDistribution.Contract.GetSalePriceDiscount(&_NFTDistribution.CallOpts, inviter, proofRoot)
}

// GetSalePriceDiscount is a free data retrieval call binding the contract method 0x7d31dad9.
//
// Solidity: function getSalePriceDiscount(address inviter, bytes32 proofRoot) view returns(uint256)
func (_NFTDistribution *NFTDistributionCallerSession) GetSalePriceDiscount(inviter common.Address, proofRoot [32]byte) (*big.Int, error) {
	return _NFTDistribution.Contract.GetSalePriceDiscount(&_NFTDistribution.CallOpts, inviter, proofRoot)
}

// GetTopCommissions is a free data retrieval call binding the contract method 0x60044bce.
//
// Solidity: function getTopCommissions() view returns((address,uint128,uint128,uint256,bool,uint256)[])
func (_NFTDistribution *NFTDistributionCaller) GetTopCommissions(opts *bind.CallOpts) ([]NFTDistributionUserInfo, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "getTopCommissions")

	if err != nil {
		return *new([]NFTDistributionUserInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]NFTDistributionUserInfo)).(*[]NFTDistributionUserInfo)

	return out0, err

}

// GetTopCommissions is a free data retrieval call binding the contract method 0x60044bce.
//
// Solidity: function getTopCommissions() view returns((address,uint128,uint128,uint256,bool,uint256)[])
func (_NFTDistribution *NFTDistributionSession) GetTopCommissions() ([]NFTDistributionUserInfo, error) {
	return _NFTDistribution.Contract.GetTopCommissions(&_NFTDistribution.CallOpts)
}

// GetTopCommissions is a free data retrieval call binding the contract method 0x60044bce.
//
// Solidity: function getTopCommissions() view returns((address,uint128,uint128,uint256,bool,uint256)[])
func (_NFTDistribution *NFTDistributionCallerSession) GetTopCommissions() ([]NFTDistributionUserInfo, error) {
	return _NFTDistribution.Contract.GetTopCommissions(&_NFTDistribution.CallOpts)
}

// GetUserPreSaledNfts is a free data retrieval call binding the contract method 0x7aedd1ff.
//
// Solidity: function getUserPreSaledNfts(address user) view returns(uint256[])
func (_NFTDistribution *NFTDistributionCaller) GetUserPreSaledNfts(opts *bind.CallOpts, user common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "getUserPreSaledNfts", user)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserPreSaledNfts is a free data retrieval call binding the contract method 0x7aedd1ff.
//
// Solidity: function getUserPreSaledNfts(address user) view returns(uint256[])
func (_NFTDistribution *NFTDistributionSession) GetUserPreSaledNfts(user common.Address) ([]*big.Int, error) {
	return _NFTDistribution.Contract.GetUserPreSaledNfts(&_NFTDistribution.CallOpts, user)
}

// GetUserPreSaledNfts is a free data retrieval call binding the contract method 0x7aedd1ff.
//
// Solidity: function getUserPreSaledNfts(address user) view returns(uint256[])
func (_NFTDistribution *NFTDistributionCallerSession) GetUserPreSaledNfts(user common.Address) ([]*big.Int, error) {
	return _NFTDistribution.Contract.GetUserPreSaledNfts(&_NFTDistribution.CallOpts, user)
}

// IsPreSaleTime is a free data retrieval call binding the contract method 0xebab43e4.
//
// Solidity: function isPreSaleTime() view returns(bool)
func (_NFTDistribution *NFTDistributionCaller) IsPreSaleTime(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "isPreSaleTime")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPreSaleTime is a free data retrieval call binding the contract method 0xebab43e4.
//
// Solidity: function isPreSaleTime() view returns(bool)
func (_NFTDistribution *NFTDistributionSession) IsPreSaleTime() (bool, error) {
	return _NFTDistribution.Contract.IsPreSaleTime(&_NFTDistribution.CallOpts)
}

// IsPreSaleTime is a free data retrieval call binding the contract method 0xebab43e4.
//
// Solidity: function isPreSaleTime() view returns(bool)
func (_NFTDistribution *NFTDistributionCallerSession) IsPreSaleTime() (bool, error) {
	return _NFTDistribution.Contract.IsPreSaleTime(&_NFTDistribution.CallOpts)
}

// NftToken is a free data retrieval call binding the contract method 0xd06fcba8.
//
// Solidity: function nftToken() view returns(address)
func (_NFTDistribution *NFTDistributionCaller) NftToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "nftToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftToken is a free data retrieval call binding the contract method 0xd06fcba8.
//
// Solidity: function nftToken() view returns(address)
func (_NFTDistribution *NFTDistributionSession) NftToken() (common.Address, error) {
	return _NFTDistribution.Contract.NftToken(&_NFTDistribution.CallOpts)
}

// NftToken is a free data retrieval call binding the contract method 0xd06fcba8.
//
// Solidity: function nftToken() view returns(address)
func (_NFTDistribution *NFTDistributionCallerSession) NftToken() (common.Address, error) {
	return _NFTDistribution.Contract.NftToken(&_NFTDistribution.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTDistribution *NFTDistributionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTDistribution *NFTDistributionSession) Owner() (common.Address, error) {
	return _NFTDistribution.Contract.Owner(&_NFTDistribution.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTDistribution *NFTDistributionCallerSession) Owner() (common.Address, error) {
	return _NFTDistribution.Contract.Owner(&_NFTDistribution.CallOpts)
}

// PreSaleConfig is a free data retrieval call binding the contract method 0x61dd2173.
//
// Solidity: function preSaleConfig() view returns(uint256 endTime, uint256 discountRate)
func (_NFTDistribution *NFTDistributionCaller) PreSaleConfig(opts *bind.CallOpts) (struct {
	EndTime      *big.Int
	DiscountRate *big.Int
}, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "preSaleConfig")

	outstruct := new(struct {
		EndTime      *big.Int
		DiscountRate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EndTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DiscountRate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PreSaleConfig is a free data retrieval call binding the contract method 0x61dd2173.
//
// Solidity: function preSaleConfig() view returns(uint256 endTime, uint256 discountRate)
func (_NFTDistribution *NFTDistributionSession) PreSaleConfig() (struct {
	EndTime      *big.Int
	DiscountRate *big.Int
}, error) {
	return _NFTDistribution.Contract.PreSaleConfig(&_NFTDistribution.CallOpts)
}

// PreSaleConfig is a free data retrieval call binding the contract method 0x61dd2173.
//
// Solidity: function preSaleConfig() view returns(uint256 endTime, uint256 discountRate)
func (_NFTDistribution *NFTDistributionCallerSession) PreSaleConfig() (struct {
	EndTime      *big.Int
	DiscountRate *big.Int
}, error) {
	return _NFTDistribution.Contract.PreSaleConfig(&_NFTDistribution.CallOpts)
}

// PreSaledNftClaimed is a free data retrieval call binding the contract method 0xeb72aeb0.
//
// Solidity: function preSaledNftClaimed(uint256 ) view returns(bool)
func (_NFTDistribution *NFTDistributionCaller) PreSaledNftClaimed(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "preSaledNftClaimed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PreSaledNftClaimed is a free data retrieval call binding the contract method 0xeb72aeb0.
//
// Solidity: function preSaledNftClaimed(uint256 ) view returns(bool)
func (_NFTDistribution *NFTDistributionSession) PreSaledNftClaimed(arg0 *big.Int) (bool, error) {
	return _NFTDistribution.Contract.PreSaledNftClaimed(&_NFTDistribution.CallOpts, arg0)
}

// PreSaledNftClaimed is a free data retrieval call binding the contract method 0xeb72aeb0.
//
// Solidity: function preSaledNftClaimed(uint256 ) view returns(bool)
func (_NFTDistribution *NFTDistributionCallerSession) PreSaledNftClaimed(arg0 *big.Int) (bool, error) {
	return _NFTDistribution.Contract.PreSaledNftClaimed(&_NFTDistribution.CallOpts, arg0)
}

// PreSaledNftOwner is a free data retrieval call binding the contract method 0x706499ec.
//
// Solidity: function preSaledNftOwner(uint256 ) view returns(address)
func (_NFTDistribution *NFTDistributionCaller) PreSaledNftOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "preSaledNftOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreSaledNftOwner is a free data retrieval call binding the contract method 0x706499ec.
//
// Solidity: function preSaledNftOwner(uint256 ) view returns(address)
func (_NFTDistribution *NFTDistributionSession) PreSaledNftOwner(arg0 *big.Int) (common.Address, error) {
	return _NFTDistribution.Contract.PreSaledNftOwner(&_NFTDistribution.CallOpts, arg0)
}

// PreSaledNftOwner is a free data retrieval call binding the contract method 0x706499ec.
//
// Solidity: function preSaledNftOwner(uint256 ) view returns(address)
func (_NFTDistribution *NFTDistributionCallerSession) PreSaledNftOwner(arg0 *big.Int) (common.Address, error) {
	return _NFTDistribution.Contract.PreSaledNftOwner(&_NFTDistribution.CallOpts, arg0)
}

// PreSaledNfts is a free data retrieval call binding the contract method 0x24322bef.
//
// Solidity: function preSaledNfts(address , uint256 ) view returns(uint256)
func (_NFTDistribution *NFTDistributionCaller) PreSaledNfts(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "preSaledNfts", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreSaledNfts is a free data retrieval call binding the contract method 0x24322bef.
//
// Solidity: function preSaledNfts(address , uint256 ) view returns(uint256)
func (_NFTDistribution *NFTDistributionSession) PreSaledNfts(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _NFTDistribution.Contract.PreSaledNfts(&_NFTDistribution.CallOpts, arg0, arg1)
}

// PreSaledNfts is a free data retrieval call binding the contract method 0x24322bef.
//
// Solidity: function preSaledNfts(address , uint256 ) view returns(uint256)
func (_NFTDistribution *NFTDistributionCallerSession) PreSaledNfts(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _NFTDistribution.Contract.PreSaledNfts(&_NFTDistribution.CallOpts, arg0, arg1)
}

// SaleActive is a free data retrieval call binding the contract method 0x68428a1b.
//
// Solidity: function saleActive() view returns(bool)
func (_NFTDistribution *NFTDistributionCaller) SaleActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "saleActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SaleActive is a free data retrieval call binding the contract method 0x68428a1b.
//
// Solidity: function saleActive() view returns(bool)
func (_NFTDistribution *NFTDistributionSession) SaleActive() (bool, error) {
	return _NFTDistribution.Contract.SaleActive(&_NFTDistribution.CallOpts)
}

// SaleActive is a free data retrieval call binding the contract method 0x68428a1b.
//
// Solidity: function saleActive() view returns(bool)
func (_NFTDistribution *NFTDistributionCallerSession) SaleActive() (bool, error) {
	return _NFTDistribution.Contract.SaleActive(&_NFTDistribution.CallOpts)
}

// SalePrice is a free data retrieval call binding the contract method 0xf51f96dd.
//
// Solidity: function salePrice() view returns(uint128)
func (_NFTDistribution *NFTDistributionCaller) SalePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "salePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SalePrice is a free data retrieval call binding the contract method 0xf51f96dd.
//
// Solidity: function salePrice() view returns(uint128)
func (_NFTDistribution *NFTDistributionSession) SalePrice() (*big.Int, error) {
	return _NFTDistribution.Contract.SalePrice(&_NFTDistribution.CallOpts)
}

// SalePrice is a free data retrieval call binding the contract method 0xf51f96dd.
//
// Solidity: function salePrice() view returns(uint128)
func (_NFTDistribution *NFTDistributionCallerSession) SalePrice() (*big.Int, error) {
	return _NFTDistribution.Contract.SalePrice(&_NFTDistribution.CallOpts)
}

// Sort is a free data retrieval call binding the contract method 0x42e528c7.
//
// Solidity: function sort((address,uint128,uint128,uint256,bool,uint256)[] _input) pure returns((address,uint128,uint128,uint256,bool,uint256)[])
func (_NFTDistribution *NFTDistributionCaller) Sort(opts *bind.CallOpts, _input []NFTDistributionUserInfo) ([]NFTDistributionUserInfo, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "sort", _input)

	if err != nil {
		return *new([]NFTDistributionUserInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]NFTDistributionUserInfo)).(*[]NFTDistributionUserInfo)

	return out0, err

}

// Sort is a free data retrieval call binding the contract method 0x42e528c7.
//
// Solidity: function sort((address,uint128,uint128,uint256,bool,uint256)[] _input) pure returns((address,uint128,uint128,uint256,bool,uint256)[])
func (_NFTDistribution *NFTDistributionSession) Sort(_input []NFTDistributionUserInfo) ([]NFTDistributionUserInfo, error) {
	return _NFTDistribution.Contract.Sort(&_NFTDistribution.CallOpts, _input)
}

// Sort is a free data retrieval call binding the contract method 0x42e528c7.
//
// Solidity: function sort((address,uint128,uint128,uint256,bool,uint256)[] _input) pure returns((address,uint128,uint128,uint256,bool,uint256)[])
func (_NFTDistribution *NFTDistributionCallerSession) Sort(_input []NFTDistributionUserInfo) ([]NFTDistributionUserInfo, error) {
	return _NFTDistribution.Contract.Sort(&_NFTDistribution.CallOpts, _input)
}

// SpecialRateRule is a free data retrieval call binding the contract method 0x591e23bc.
//
// Solidity: function specialRateRule(bytes32 ) view returns(bool isActive, address rateRule)
func (_NFTDistribution *NFTDistributionCaller) SpecialRateRule(opts *bind.CallOpts, arg0 [32]byte) (struct {
	IsActive bool
	RateRule common.Address
}, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "specialRateRule", arg0)

	outstruct := new(struct {
		IsActive bool
		RateRule common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsActive = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.RateRule = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// SpecialRateRule is a free data retrieval call binding the contract method 0x591e23bc.
//
// Solidity: function specialRateRule(bytes32 ) view returns(bool isActive, address rateRule)
func (_NFTDistribution *NFTDistributionSession) SpecialRateRule(arg0 [32]byte) (struct {
	IsActive bool
	RateRule common.Address
}, error) {
	return _NFTDistribution.Contract.SpecialRateRule(&_NFTDistribution.CallOpts, arg0)
}

// SpecialRateRule is a free data retrieval call binding the contract method 0x591e23bc.
//
// Solidity: function specialRateRule(bytes32 ) view returns(bool isActive, address rateRule)
func (_NFTDistribution *NFTDistributionCallerSession) SpecialRateRule(arg0 [32]byte) (struct {
	IsActive bool
	RateRule common.Address
}, error) {
	return _NFTDistribution.Contract.SpecialRateRule(&_NFTDistribution.CallOpts, arg0)
}

// ToTokenId is a free data retrieval call binding the contract method 0x980f3700.
//
// Solidity: function toTokenId() view returns(uint32)
func (_NFTDistribution *NFTDistributionCaller) ToTokenId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "toTokenId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ToTokenId is a free data retrieval call binding the contract method 0x980f3700.
//
// Solidity: function toTokenId() view returns(uint32)
func (_NFTDistribution *NFTDistributionSession) ToTokenId() (uint32, error) {
	return _NFTDistribution.Contract.ToTokenId(&_NFTDistribution.CallOpts)
}

// ToTokenId is a free data retrieval call binding the contract method 0x980f3700.
//
// Solidity: function toTokenId() view returns(uint32)
func (_NFTDistribution *NFTDistributionCallerSession) ToTokenId() (uint32, error) {
	return _NFTDistribution.Contract.ToTokenId(&_NFTDistribution.CallOpts)
}

// TopCommissions is a free data retrieval call binding the contract method 0xabeadb0a.
//
// Solidity: function topCommissions(uint256 ) view returns(address user, uint128 buyCount, uint128 invitesCount, uint256 commission, bool inTopCommissions, uint256 topCommissionIndex)
func (_NFTDistribution *NFTDistributionCaller) TopCommissions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	User               common.Address
	BuyCount           *big.Int
	InvitesCount       *big.Int
	Commission         *big.Int
	InTopCommissions   bool
	TopCommissionIndex *big.Int
}, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "topCommissions", arg0)

	outstruct := new(struct {
		User               common.Address
		BuyCount           *big.Int
		InvitesCount       *big.Int
		Commission         *big.Int
		InTopCommissions   bool
		TopCommissionIndex *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.BuyCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InvitesCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Commission = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.InTopCommissions = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.TopCommissionIndex = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TopCommissions is a free data retrieval call binding the contract method 0xabeadb0a.
//
// Solidity: function topCommissions(uint256 ) view returns(address user, uint128 buyCount, uint128 invitesCount, uint256 commission, bool inTopCommissions, uint256 topCommissionIndex)
func (_NFTDistribution *NFTDistributionSession) TopCommissions(arg0 *big.Int) (struct {
	User               common.Address
	BuyCount           *big.Int
	InvitesCount       *big.Int
	Commission         *big.Int
	InTopCommissions   bool
	TopCommissionIndex *big.Int
}, error) {
	return _NFTDistribution.Contract.TopCommissions(&_NFTDistribution.CallOpts, arg0)
}

// TopCommissions is a free data retrieval call binding the contract method 0xabeadb0a.
//
// Solidity: function topCommissions(uint256 ) view returns(address user, uint128 buyCount, uint128 invitesCount, uint256 commission, bool inTopCommissions, uint256 topCommissionIndex)
func (_NFTDistribution *NFTDistributionCallerSession) TopCommissions(arg0 *big.Int) (struct {
	User               common.Address
	BuyCount           *big.Int
	InvitesCount       *big.Int
	Commission         *big.Int
	InTopCommissions   bool
	TopCommissionIndex *big.Int
}, error) {
	return _NFTDistribution.Contract.TopCommissions(&_NFTDistribution.CallOpts, arg0)
}

// TradeToken is a free data retrieval call binding the contract method 0xd83678ac.
//
// Solidity: function tradeToken() view returns(address)
func (_NFTDistribution *NFTDistributionCaller) TradeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "tradeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TradeToken is a free data retrieval call binding the contract method 0xd83678ac.
//
// Solidity: function tradeToken() view returns(address)
func (_NFTDistribution *NFTDistributionSession) TradeToken() (common.Address, error) {
	return _NFTDistribution.Contract.TradeToken(&_NFTDistribution.CallOpts)
}

// TradeToken is a free data retrieval call binding the contract method 0xd83678ac.
//
// Solidity: function tradeToken() view returns(address)
func (_NFTDistribution *NFTDistributionCallerSession) TradeToken() (common.Address, error) {
	return _NFTDistribution.Contract.TradeToken(&_NFTDistribution.CallOpts)
}

// UncommonNft is a free data retrieval call binding the contract method 0xfe9b7349.
//
// Solidity: function uncommonNft(uint256 ) view returns(bool)
func (_NFTDistribution *NFTDistributionCaller) UncommonNft(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "uncommonNft", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UncommonNft is a free data retrieval call binding the contract method 0xfe9b7349.
//
// Solidity: function uncommonNft(uint256 ) view returns(bool)
func (_NFTDistribution *NFTDistributionSession) UncommonNft(arg0 *big.Int) (bool, error) {
	return _NFTDistribution.Contract.UncommonNft(&_NFTDistribution.CallOpts, arg0)
}

// UncommonNft is a free data retrieval call binding the contract method 0xfe9b7349.
//
// Solidity: function uncommonNft(uint256 ) view returns(bool)
func (_NFTDistribution *NFTDistributionCallerSession) UncommonNft(arg0 *big.Int) (bool, error) {
	return _NFTDistribution.Contract.UncommonNft(&_NFTDistribution.CallOpts, arg0)
}

// UncommonNftTickers is a free data retrieval call binding the contract method 0xfd393ed2.
//
// Solidity: function uncommonNftTickers(uint256 ) view returns(address buyer, uint256 nftId, uint256 timestamp)
func (_NFTDistribution *NFTDistributionCaller) UncommonNftTickers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Buyer     common.Address
	NftId     *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "uncommonNftTickers", arg0)

	outstruct := new(struct {
		Buyer     common.Address
		NftId     *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Buyer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UncommonNftTickers is a free data retrieval call binding the contract method 0xfd393ed2.
//
// Solidity: function uncommonNftTickers(uint256 ) view returns(address buyer, uint256 nftId, uint256 timestamp)
func (_NFTDistribution *NFTDistributionSession) UncommonNftTickers(arg0 *big.Int) (struct {
	Buyer     common.Address
	NftId     *big.Int
	Timestamp *big.Int
}, error) {
	return _NFTDistribution.Contract.UncommonNftTickers(&_NFTDistribution.CallOpts, arg0)
}

// UncommonNftTickers is a free data retrieval call binding the contract method 0xfd393ed2.
//
// Solidity: function uncommonNftTickers(uint256 ) view returns(address buyer, uint256 nftId, uint256 timestamp)
func (_NFTDistribution *NFTDistributionCallerSession) UncommonNftTickers(arg0 *big.Int) (struct {
	Buyer     common.Address
	NftId     *big.Int
	Timestamp *big.Int
}, error) {
	return _NFTDistribution.Contract.UncommonNftTickers(&_NFTDistribution.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(address user, uint128 buyCount, uint128 invitesCount, uint256 commission, bool inTopCommissions, uint256 topCommissionIndex)
func (_NFTDistribution *NFTDistributionCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	User               common.Address
	BuyCount           *big.Int
	InvitesCount       *big.Int
	Commission         *big.Int
	InTopCommissions   bool
	TopCommissionIndex *big.Int
}, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "userInfo", arg0)

	outstruct := new(struct {
		User               common.Address
		BuyCount           *big.Int
		InvitesCount       *big.Int
		Commission         *big.Int
		InTopCommissions   bool
		TopCommissionIndex *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.BuyCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InvitesCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Commission = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.InTopCommissions = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.TopCommissionIndex = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(address user, uint128 buyCount, uint128 invitesCount, uint256 commission, bool inTopCommissions, uint256 topCommissionIndex)
func (_NFTDistribution *NFTDistributionSession) UserInfo(arg0 common.Address) (struct {
	User               common.Address
	BuyCount           *big.Int
	InvitesCount       *big.Int
	Commission         *big.Int
	InTopCommissions   bool
	TopCommissionIndex *big.Int
}, error) {
	return _NFTDistribution.Contract.UserInfo(&_NFTDistribution.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(address user, uint128 buyCount, uint128 invitesCount, uint256 commission, bool inTopCommissions, uint256 topCommissionIndex)
func (_NFTDistribution *NFTDistributionCallerSession) UserInfo(arg0 common.Address) (struct {
	User               common.Address
	BuyCount           *big.Int
	InvitesCount       *big.Int
	Commission         *big.Int
	InTopCommissions   bool
	TopCommissionIndex *big.Int
}, error) {
	return _NFTDistribution.Contract.UserInfo(&_NFTDistribution.CallOpts, arg0)
}

// VerifyProof is a free data retrieval call binding the contract method 0x54bc130a.
//
// Solidity: function verifyProof(address user, bytes32 proofRoot, bytes32[] proofPath) pure returns(bool)
func (_NFTDistribution *NFTDistributionCaller) VerifyProof(opts *bind.CallOpts, user common.Address, proofRoot [32]byte, proofPath [][32]byte) (bool, error) {
	var out []interface{}
	err := _NFTDistribution.contract.Call(opts, &out, "verifyProof", user, proofRoot, proofPath)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x54bc130a.
//
// Solidity: function verifyProof(address user, bytes32 proofRoot, bytes32[] proofPath) pure returns(bool)
func (_NFTDistribution *NFTDistributionSession) VerifyProof(user common.Address, proofRoot [32]byte, proofPath [][32]byte) (bool, error) {
	return _NFTDistribution.Contract.VerifyProof(&_NFTDistribution.CallOpts, user, proofRoot, proofPath)
}

// VerifyProof is a free data retrieval call binding the contract method 0x54bc130a.
//
// Solidity: function verifyProof(address user, bytes32 proofRoot, bytes32[] proofPath) pure returns(bool)
func (_NFTDistribution *NFTDistributionCallerSession) VerifyProof(user common.Address, proofRoot [32]byte, proofPath [][32]byte) (bool, error) {
	return _NFTDistribution.Contract.VerifyProof(&_NFTDistribution.CallOpts, user, proofRoot, proofPath)
}

// ClaimPreSaledNfts is a paid mutator transaction binding the contract method 0x114ce2c0.
//
// Solidity: function claimPreSaledNfts(uint256[] nftIds) returns()
func (_NFTDistribution *NFTDistributionTransactor) ClaimPreSaledNfts(opts *bind.TransactOpts, nftIds []*big.Int) (*types.Transaction, error) {
	return _NFTDistribution.contract.Transact(opts, "claimPreSaledNfts", nftIds)
}

// ClaimPreSaledNfts is a paid mutator transaction binding the contract method 0x114ce2c0.
//
// Solidity: function claimPreSaledNfts(uint256[] nftIds) returns()
func (_NFTDistribution *NFTDistributionSession) ClaimPreSaledNfts(nftIds []*big.Int) (*types.Transaction, error) {
	return _NFTDistribution.Contract.ClaimPreSaledNfts(&_NFTDistribution.TransactOpts, nftIds)
}

// ClaimPreSaledNfts is a paid mutator transaction binding the contract method 0x114ce2c0.
//
// Solidity: function claimPreSaledNfts(uint256[] nftIds) returns()
func (_NFTDistribution *NFTDistributionTransactorSession) ClaimPreSaledNfts(nftIds []*big.Int) (*types.Transaction, error) {
	return _NFTDistribution.Contract.ClaimPreSaledNfts(&_NFTDistribution.TransactOpts, nftIds)
}

// ClaimWithInviter is a paid mutator transaction binding the contract method 0xed38b5a4.
//
// Solidity: function claimWithInviter(address inviter) payable returns()
func (_NFTDistribution *NFTDistributionTransactor) ClaimWithInviter(opts *bind.TransactOpts, inviter common.Address) (*types.Transaction, error) {
	return _NFTDistribution.contract.Transact(opts, "claimWithInviter", inviter)
}

// ClaimWithInviter is a paid mutator transaction binding the contract method 0xed38b5a4.
//
// Solidity: function claimWithInviter(address inviter) payable returns()
func (_NFTDistribution *NFTDistributionSession) ClaimWithInviter(inviter common.Address) (*types.Transaction, error) {
	return _NFTDistribution.Contract.ClaimWithInviter(&_NFTDistribution.TransactOpts, inviter)
}

// ClaimWithInviter is a paid mutator transaction binding the contract method 0xed38b5a4.
//
// Solidity: function claimWithInviter(address inviter) payable returns()
func (_NFTDistribution *NFTDistributionTransactorSession) ClaimWithInviter(inviter common.Address) (*types.Transaction, error) {
	return _NFTDistribution.Contract.ClaimWithInviter(&_NFTDistribution.TransactOpts, inviter)
}

// ClaimWithProof is a paid mutator transaction binding the contract method 0x6bfe3436.
//
// Solidity: function claimWithProof(address inviter, bytes32 proofRoot, bytes32[] proofPath) payable returns()
func (_NFTDistribution *NFTDistributionTransactor) ClaimWithProof(opts *bind.TransactOpts, inviter common.Address, proofRoot [32]byte, proofPath [][32]byte) (*types.Transaction, error) {
	return _NFTDistribution.contract.Transact(opts, "claimWithProof", inviter, proofRoot, proofPath)
}

// ClaimWithProof is a paid mutator transaction binding the contract method 0x6bfe3436.
//
// Solidity: function claimWithProof(address inviter, bytes32 proofRoot, bytes32[] proofPath) payable returns()
func (_NFTDistribution *NFTDistributionSession) ClaimWithProof(inviter common.Address, proofRoot [32]byte, proofPath [][32]byte) (*types.Transaction, error) {
	return _NFTDistribution.Contract.ClaimWithProof(&_NFTDistribution.TransactOpts, inviter, proofRoot, proofPath)
}

// ClaimWithProof is a paid mutator transaction binding the contract method 0x6bfe3436.
//
// Solidity: function claimWithProof(address inviter, bytes32 proofRoot, bytes32[] proofPath) payable returns()
func (_NFTDistribution *NFTDistributionTransactorSession) ClaimWithProof(inviter common.Address, proofRoot [32]byte, proofPath [][32]byte) (*types.Transaction, error) {
	return _NFTDistribution.Contract.ClaimWithProof(&_NFTDistribution.TransactOpts, inviter, proofRoot, proofPath)
}

// ConfigPreSale is a paid mutator transaction binding the contract method 0x1262560f.
//
// Solidity: function configPreSale(uint256 endTime, uint256 discountRate) returns()
func (_NFTDistribution *NFTDistributionTransactor) ConfigPreSale(opts *bind.TransactOpts, endTime *big.Int, discountRate *big.Int) (*types.Transaction, error) {
	return _NFTDistribution.contract.Transact(opts, "configPreSale", endTime, discountRate)
}

// ConfigPreSale is a paid mutator transaction binding the contract method 0x1262560f.
//
// Solidity: function configPreSale(uint256 endTime, uint256 discountRate) returns()
func (_NFTDistribution *NFTDistributionSession) ConfigPreSale(endTime *big.Int, discountRate *big.Int) (*types.Transaction, error) {
	return _NFTDistribution.Contract.ConfigPreSale(&_NFTDistribution.TransactOpts, endTime, discountRate)
}

// ConfigPreSale is a paid mutator transaction binding the contract method 0x1262560f.
//
// Solidity: function configPreSale(uint256 endTime, uint256 discountRate) returns()
func (_NFTDistribution *NFTDistributionTransactorSession) ConfigPreSale(endTime *big.Int, discountRate *big.Int) (*types.Transaction, error) {
	return _NFTDistribution.Contract.ConfigPreSale(&_NFTDistribution.TransactOpts, endTime, discountRate)
}

// FlipSaleStatus is a paid mutator transaction binding the contract method 0xce03ec93.
//
// Solidity: function flipSaleStatus() returns()
func (_NFTDistribution *NFTDistributionTransactor) FlipSaleStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTDistribution.contract.Transact(opts, "flipSaleStatus")
}

// FlipSaleStatus is a paid mutator transaction binding the contract method 0xce03ec93.
//
// Solidity: function flipSaleStatus() returns()
func (_NFTDistribution *NFTDistributionSession) FlipSaleStatus() (*types.Transaction, error) {
	return _NFTDistribution.Contract.FlipSaleStatus(&_NFTDistribution.TransactOpts)
}

// FlipSaleStatus is a paid mutator transaction binding the contract method 0xce03ec93.
//
// Solidity: function flipSaleStatus() returns()
func (_NFTDistribution *NFTDistributionTransactorSession) FlipSaleStatus() (*types.Transaction, error) {
	return _NFTDistribution.Contract.FlipSaleStatus(&_NFTDistribution.TransactOpts)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_NFTDistribution *NFTDistributionTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _NFTDistribution.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_NFTDistribution *NFTDistributionSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _NFTDistribution.Contract.OnERC721Received(&_NFTDistribution.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_NFTDistribution *NFTDistributionTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _NFTDistribution.Contract.OnERC721Received(&_NFTDistribution.TransactOpts, arg0, arg1, arg2, arg3)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTDistribution *NFTDistributionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTDistribution.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTDistribution *NFTDistributionSession) RenounceOwnership() (*types.Transaction, error) {
	return _NFTDistribution.Contract.RenounceOwnership(&_NFTDistribution.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTDistribution *NFTDistributionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NFTDistribution.Contract.RenounceOwnership(&_NFTDistribution.TransactOpts)
}

// SetDefaultRateRule is a paid mutator transaction binding the contract method 0x57391e9d.
//
// Solidity: function setDefaultRateRule(address newRateRule) returns()
func (_NFTDistribution *NFTDistributionTransactor) SetDefaultRateRule(opts *bind.TransactOpts, newRateRule common.Address) (*types.Transaction, error) {
	return _NFTDistribution.contract.Transact(opts, "setDefaultRateRule", newRateRule)
}

// SetDefaultRateRule is a paid mutator transaction binding the contract method 0x57391e9d.
//
// Solidity: function setDefaultRateRule(address newRateRule) returns()
func (_NFTDistribution *NFTDistributionSession) SetDefaultRateRule(newRateRule common.Address) (*types.Transaction, error) {
	return _NFTDistribution.Contract.SetDefaultRateRule(&_NFTDistribution.TransactOpts, newRateRule)
}

// SetDefaultRateRule is a paid mutator transaction binding the contract method 0x57391e9d.
//
// Solidity: function setDefaultRateRule(address newRateRule) returns()
func (_NFTDistribution *NFTDistributionTransactorSession) SetDefaultRateRule(newRateRule common.Address) (*types.Transaction, error) {
	return _NFTDistribution.Contract.SetDefaultRateRule(&_NFTDistribution.TransactOpts, newRateRule)
}

// SetSalePrice is a paid mutator transaction binding the contract method 0xf492538b.
//
// Solidity: function setSalePrice(uint128 _salePrice) returns()
func (_NFTDistribution *NFTDistributionTransactor) SetSalePrice(opts *bind.TransactOpts, _salePrice *big.Int) (*types.Transaction, error) {
	return _NFTDistribution.contract.Transact(opts, "setSalePrice", _salePrice)
}

// SetSalePrice is a paid mutator transaction binding the contract method 0xf492538b.
//
// Solidity: function setSalePrice(uint128 _salePrice) returns()
func (_NFTDistribution *NFTDistributionSession) SetSalePrice(_salePrice *big.Int) (*types.Transaction, error) {
	return _NFTDistribution.Contract.SetSalePrice(&_NFTDistribution.TransactOpts, _salePrice)
}

// SetSalePrice is a paid mutator transaction binding the contract method 0xf492538b.
//
// Solidity: function setSalePrice(uint128 _salePrice) returns()
func (_NFTDistribution *NFTDistributionTransactorSession) SetSalePrice(_salePrice *big.Int) (*types.Transaction, error) {
	return _NFTDistribution.Contract.SetSalePrice(&_NFTDistribution.TransactOpts, _salePrice)
}

// SetSpecialRateRule is a paid mutator transaction binding the contract method 0x596b0363.
//
// Solidity: function setSpecialRateRule(bytes32 proofRoot, address newRateRule) returns()
func (_NFTDistribution *NFTDistributionTransactor) SetSpecialRateRule(opts *bind.TransactOpts, proofRoot [32]byte, newRateRule common.Address) (*types.Transaction, error) {
	return _NFTDistribution.contract.Transact(opts, "setSpecialRateRule", proofRoot, newRateRule)
}

// SetSpecialRateRule is a paid mutator transaction binding the contract method 0x596b0363.
//
// Solidity: function setSpecialRateRule(bytes32 proofRoot, address newRateRule) returns()
func (_NFTDistribution *NFTDistributionSession) SetSpecialRateRule(proofRoot [32]byte, newRateRule common.Address) (*types.Transaction, error) {
	return _NFTDistribution.Contract.SetSpecialRateRule(&_NFTDistribution.TransactOpts, proofRoot, newRateRule)
}

// SetSpecialRateRule is a paid mutator transaction binding the contract method 0x596b0363.
//
// Solidity: function setSpecialRateRule(bytes32 proofRoot, address newRateRule) returns()
func (_NFTDistribution *NFTDistributionTransactorSession) SetSpecialRateRule(proofRoot [32]byte, newRateRule common.Address) (*types.Transaction, error) {
	return _NFTDistribution.Contract.SetSpecialRateRule(&_NFTDistribution.TransactOpts, proofRoot, newRateRule)
}

// SetTokenScope is a paid mutator transaction binding the contract method 0x6663a375.
//
// Solidity: function setTokenScope(uint32 _fromTokenId, uint32 _toTokenId) returns()
func (_NFTDistribution *NFTDistributionTransactor) SetTokenScope(opts *bind.TransactOpts, _fromTokenId uint32, _toTokenId uint32) (*types.Transaction, error) {
	return _NFTDistribution.contract.Transact(opts, "setTokenScope", _fromTokenId, _toTokenId)
}

// SetTokenScope is a paid mutator transaction binding the contract method 0x6663a375.
//
// Solidity: function setTokenScope(uint32 _fromTokenId, uint32 _toTokenId) returns()
func (_NFTDistribution *NFTDistributionSession) SetTokenScope(_fromTokenId uint32, _toTokenId uint32) (*types.Transaction, error) {
	return _NFTDistribution.Contract.SetTokenScope(&_NFTDistribution.TransactOpts, _fromTokenId, _toTokenId)
}

// SetTokenScope is a paid mutator transaction binding the contract method 0x6663a375.
//
// Solidity: function setTokenScope(uint32 _fromTokenId, uint32 _toTokenId) returns()
func (_NFTDistribution *NFTDistributionTransactorSession) SetTokenScope(_fromTokenId uint32, _toTokenId uint32) (*types.Transaction, error) {
	return _NFTDistribution.Contract.SetTokenScope(&_NFTDistribution.TransactOpts, _fromTokenId, _toTokenId)
}

// SetTradeToken is a paid mutator transaction binding the contract method 0xd1c5f2c9.
//
// Solidity: function setTradeToken(address newTradeToken) returns()
func (_NFTDistribution *NFTDistributionTransactor) SetTradeToken(opts *bind.TransactOpts, newTradeToken common.Address) (*types.Transaction, error) {
	return _NFTDistribution.contract.Transact(opts, "setTradeToken", newTradeToken)
}

// SetTradeToken is a paid mutator transaction binding the contract method 0xd1c5f2c9.
//
// Solidity: function setTradeToken(address newTradeToken) returns()
func (_NFTDistribution *NFTDistributionSession) SetTradeToken(newTradeToken common.Address) (*types.Transaction, error) {
	return _NFTDistribution.Contract.SetTradeToken(&_NFTDistribution.TransactOpts, newTradeToken)
}

// SetTradeToken is a paid mutator transaction binding the contract method 0xd1c5f2c9.
//
// Solidity: function setTradeToken(address newTradeToken) returns()
func (_NFTDistribution *NFTDistributionTransactorSession) SetTradeToken(newTradeToken common.Address) (*types.Transaction, error) {
	return _NFTDistribution.Contract.SetTradeToken(&_NFTDistribution.TransactOpts, newTradeToken)
}

// SetUncommonNFTs is a paid mutator transaction binding the contract method 0x71f73dde.
//
// Solidity: function setUncommonNFTs(uint256[] nftIds, bool isUncommon) returns()
func (_NFTDistribution *NFTDistributionTransactor) SetUncommonNFTs(opts *bind.TransactOpts, nftIds []*big.Int, isUncommon bool) (*types.Transaction, error) {
	return _NFTDistribution.contract.Transact(opts, "setUncommonNFTs", nftIds, isUncommon)
}

// SetUncommonNFTs is a paid mutator transaction binding the contract method 0x71f73dde.
//
// Solidity: function setUncommonNFTs(uint256[] nftIds, bool isUncommon) returns()
func (_NFTDistribution *NFTDistributionSession) SetUncommonNFTs(nftIds []*big.Int, isUncommon bool) (*types.Transaction, error) {
	return _NFTDistribution.Contract.SetUncommonNFTs(&_NFTDistribution.TransactOpts, nftIds, isUncommon)
}

// SetUncommonNFTs is a paid mutator transaction binding the contract method 0x71f73dde.
//
// Solidity: function setUncommonNFTs(uint256[] nftIds, bool isUncommon) returns()
func (_NFTDistribution *NFTDistributionTransactorSession) SetUncommonNFTs(nftIds []*big.Int, isUncommon bool) (*types.Transaction, error) {
	return _NFTDistribution.Contract.SetUncommonNFTs(&_NFTDistribution.TransactOpts, nftIds, isUncommon)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTDistribution *NFTDistributionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NFTDistribution.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTDistribution *NFTDistributionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NFTDistribution.Contract.TransferOwnership(&_NFTDistribution.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTDistribution *NFTDistributionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NFTDistribution.Contract.TransferOwnership(&_NFTDistribution.TransactOpts, newOwner)
}

// UpdateFundMgr is a paid mutator transaction binding the contract method 0x9d73361c.
//
// Solidity: function updateFundMgr(address newFundMgr) returns()
func (_NFTDistribution *NFTDistributionTransactor) UpdateFundMgr(opts *bind.TransactOpts, newFundMgr common.Address) (*types.Transaction, error) {
	return _NFTDistribution.contract.Transact(opts, "updateFundMgr", newFundMgr)
}

// UpdateFundMgr is a paid mutator transaction binding the contract method 0x9d73361c.
//
// Solidity: function updateFundMgr(address newFundMgr) returns()
func (_NFTDistribution *NFTDistributionSession) UpdateFundMgr(newFundMgr common.Address) (*types.Transaction, error) {
	return _NFTDistribution.Contract.UpdateFundMgr(&_NFTDistribution.TransactOpts, newFundMgr)
}

// UpdateFundMgr is a paid mutator transaction binding the contract method 0x9d73361c.
//
// Solidity: function updateFundMgr(address newFundMgr) returns()
func (_NFTDistribution *NFTDistributionTransactorSession) UpdateFundMgr(newFundMgr common.Address) (*types.Transaction, error) {
	return _NFTDistribution.Contract.UpdateFundMgr(&_NFTDistribution.TransactOpts, newFundMgr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address token) returns()
func (_NFTDistribution *NFTDistributionTransactor) Withdraw(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _NFTDistribution.contract.Transact(opts, "withdraw", token)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address token) returns()
func (_NFTDistribution *NFTDistributionSession) Withdraw(token common.Address) (*types.Transaction, error) {
	return _NFTDistribution.Contract.Withdraw(&_NFTDistribution.TransactOpts, token)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address token) returns()
func (_NFTDistribution *NFTDistributionTransactorSession) Withdraw(token common.Address) (*types.Transaction, error) {
	return _NFTDistribution.Contract.Withdraw(&_NFTDistribution.TransactOpts, token)
}

// NFTDistributionClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the NFTDistribution contract.
type NFTDistributionClaimedIterator struct {
	Event *NFTDistributionClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTDistributionClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDistributionClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTDistributionClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTDistributionClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDistributionClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDistributionClaimed represents a Claimed event raised by the NFTDistribution contract.
type NFTDistributionClaimed struct {
	Claimer common.Address
	Inviter common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address indexed claimer, address indexed inviter, uint256 tokenId)
func (_NFTDistribution *NFTDistributionFilterer) FilterClaimed(opts *bind.FilterOpts, claimer []common.Address, inviter []common.Address) (*NFTDistributionClaimedIterator, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var inviterRule []interface{}
	for _, inviterItem := range inviter {
		inviterRule = append(inviterRule, inviterItem)
	}

	logs, sub, err := _NFTDistribution.contract.FilterLogs(opts, "Claimed", claimerRule, inviterRule)
	if err != nil {
		return nil, err
	}
	return &NFTDistributionClaimedIterator{contract: _NFTDistribution.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address indexed claimer, address indexed inviter, uint256 tokenId)
func (_NFTDistribution *NFTDistributionFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *NFTDistributionClaimed, claimer []common.Address, inviter []common.Address) (event.Subscription, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var inviterRule []interface{}
	for _, inviterItem := range inviter {
		inviterRule = append(inviterRule, inviterItem)
	}

	logs, sub, err := _NFTDistribution.contract.WatchLogs(opts, "Claimed", claimerRule, inviterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDistributionClaimed)
				if err := _NFTDistribution.contract.UnpackLog(event, "Claimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimed is a log parse operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address indexed claimer, address indexed inviter, uint256 tokenId)
func (_NFTDistribution *NFTDistributionFilterer) ParseClaimed(log types.Log) (*NFTDistributionClaimed, error) {
	event := new(NFTDistributionClaimed)
	if err := _NFTDistribution.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDistributionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NFTDistribution contract.
type NFTDistributionOwnershipTransferredIterator struct {
	Event *NFTDistributionOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTDistributionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDistributionOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTDistributionOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTDistributionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDistributionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDistributionOwnershipTransferred represents a OwnershipTransferred event raised by the NFTDistribution contract.
type NFTDistributionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFTDistribution *NFTDistributionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NFTDistributionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFTDistribution.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NFTDistributionOwnershipTransferredIterator{contract: _NFTDistribution.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFTDistribution *NFTDistributionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NFTDistributionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFTDistribution.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDistributionOwnershipTransferred)
				if err := _NFTDistribution.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFTDistribution *NFTDistributionFilterer) ParseOwnershipTransferred(log types.Log) (*NFTDistributionOwnershipTransferred, error) {
	event := new(NFTDistributionOwnershipTransferred)
	if err := _NFTDistribution.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDistributionPreSaleClaimedIterator is returned from FilterPreSaleClaimed and is used to iterate over the raw logs and unpacked data for PreSaleClaimed events raised by the NFTDistribution contract.
type NFTDistributionPreSaleClaimedIterator struct {
	Event *NFTDistributionPreSaleClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTDistributionPreSaleClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDistributionPreSaleClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTDistributionPreSaleClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTDistributionPreSaleClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDistributionPreSaleClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDistributionPreSaleClaimed represents a PreSaleClaimed event raised by the NFTDistribution contract.
type NFTDistributionPreSaleClaimed struct {
	Claimer common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPreSaleClaimed is a free log retrieval operation binding the contract event 0x8b40ac8053427a2f3a4193659def04afceb931b2c1456dd053a4b82259e91ece.
//
// Solidity: event PreSaleClaimed(address indexed claimer, uint256 tokenId)
func (_NFTDistribution *NFTDistributionFilterer) FilterPreSaleClaimed(opts *bind.FilterOpts, claimer []common.Address) (*NFTDistributionPreSaleClaimedIterator, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _NFTDistribution.contract.FilterLogs(opts, "PreSaleClaimed", claimerRule)
	if err != nil {
		return nil, err
	}
	return &NFTDistributionPreSaleClaimedIterator{contract: _NFTDistribution.contract, event: "PreSaleClaimed", logs: logs, sub: sub}, nil
}

// WatchPreSaleClaimed is a free log subscription operation binding the contract event 0x8b40ac8053427a2f3a4193659def04afceb931b2c1456dd053a4b82259e91ece.
//
// Solidity: event PreSaleClaimed(address indexed claimer, uint256 tokenId)
func (_NFTDistribution *NFTDistributionFilterer) WatchPreSaleClaimed(opts *bind.WatchOpts, sink chan<- *NFTDistributionPreSaleClaimed, claimer []common.Address) (event.Subscription, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _NFTDistribution.contract.WatchLogs(opts, "PreSaleClaimed", claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDistributionPreSaleClaimed)
				if err := _NFTDistribution.contract.UnpackLog(event, "PreSaleClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePreSaleClaimed is a log parse operation binding the contract event 0x8b40ac8053427a2f3a4193659def04afceb931b2c1456dd053a4b82259e91ece.
//
// Solidity: event PreSaleClaimed(address indexed claimer, uint256 tokenId)
func (_NFTDistribution *NFTDistributionFilterer) ParsePreSaleClaimed(log types.Log) (*NFTDistributionPreSaleClaimed, error) {
	event := new(NFTDistributionPreSaleClaimed)
	if err := _NFTDistribution.contract.UnpackLog(event, "PreSaleClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDistributionPreSaledIterator is returned from FilterPreSaled and is used to iterate over the raw logs and unpacked data for PreSaled events raised by the NFTDistribution contract.
type NFTDistributionPreSaledIterator struct {
	Event *NFTDistributionPreSaled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTDistributionPreSaledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDistributionPreSaled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTDistributionPreSaled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTDistributionPreSaledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDistributionPreSaledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDistributionPreSaled represents a PreSaled event raised by the NFTDistribution contract.
type NFTDistributionPreSaled struct {
	Claimer common.Address
	Inviter common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPreSaled is a free log retrieval operation binding the contract event 0xa753a91db684e7644c639c97262d6fff0b92266ae2994b0896d1e0d0ecd6549d.
//
// Solidity: event PreSaled(address indexed claimer, address indexed inviter, uint256 tokenId)
func (_NFTDistribution *NFTDistributionFilterer) FilterPreSaled(opts *bind.FilterOpts, claimer []common.Address, inviter []common.Address) (*NFTDistributionPreSaledIterator, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var inviterRule []interface{}
	for _, inviterItem := range inviter {
		inviterRule = append(inviterRule, inviterItem)
	}

	logs, sub, err := _NFTDistribution.contract.FilterLogs(opts, "PreSaled", claimerRule, inviterRule)
	if err != nil {
		return nil, err
	}
	return &NFTDistributionPreSaledIterator{contract: _NFTDistribution.contract, event: "PreSaled", logs: logs, sub: sub}, nil
}

// WatchPreSaled is a free log subscription operation binding the contract event 0xa753a91db684e7644c639c97262d6fff0b92266ae2994b0896d1e0d0ecd6549d.
//
// Solidity: event PreSaled(address indexed claimer, address indexed inviter, uint256 tokenId)
func (_NFTDistribution *NFTDistributionFilterer) WatchPreSaled(opts *bind.WatchOpts, sink chan<- *NFTDistributionPreSaled, claimer []common.Address, inviter []common.Address) (event.Subscription, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var inviterRule []interface{}
	for _, inviterItem := range inviter {
		inviterRule = append(inviterRule, inviterItem)
	}

	logs, sub, err := _NFTDistribution.contract.WatchLogs(opts, "PreSaled", claimerRule, inviterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDistributionPreSaled)
				if err := _NFTDistribution.contract.UnpackLog(event, "PreSaled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePreSaled is a log parse operation binding the contract event 0xa753a91db684e7644c639c97262d6fff0b92266ae2994b0896d1e0d0ecd6549d.
//
// Solidity: event PreSaled(address indexed claimer, address indexed inviter, uint256 tokenId)
func (_NFTDistribution *NFTDistributionFilterer) ParsePreSaled(log types.Log) (*NFTDistributionPreSaled, error) {
	event := new(NFTDistributionPreSaled)
	if err := _NFTDistribution.contract.UnpackLog(event, "PreSaled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTDistributionRolledOverIterator is returned from FilterRolledOver and is used to iterate over the raw logs and unpacked data for RolledOver events raised by the NFTDistribution contract.
type NFTDistributionRolledOverIterator struct {
	Event *NFTDistributionRolledOver // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NFTDistributionRolledOverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDistributionRolledOver)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NFTDistributionRolledOver)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NFTDistributionRolledOverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDistributionRolledOverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDistributionRolledOver represents a RolledOver event raised by the NFTDistribution contract.
type NFTDistributionRolledOver struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRolledOver is a free log retrieval operation binding the contract event 0xd2bce5b0f3cb0c9c70f1761d4ddcaf6236907e4b29bb4a7546b6a0c64be90cf8.
//
// Solidity: event RolledOver(bool status)
func (_NFTDistribution *NFTDistributionFilterer) FilterRolledOver(opts *bind.FilterOpts) (*NFTDistributionRolledOverIterator, error) {

	logs, sub, err := _NFTDistribution.contract.FilterLogs(opts, "RolledOver")
	if err != nil {
		return nil, err
	}
	return &NFTDistributionRolledOverIterator{contract: _NFTDistribution.contract, event: "RolledOver", logs: logs, sub: sub}, nil
}

// WatchRolledOver is a free log subscription operation binding the contract event 0xd2bce5b0f3cb0c9c70f1761d4ddcaf6236907e4b29bb4a7546b6a0c64be90cf8.
//
// Solidity: event RolledOver(bool status)
func (_NFTDistribution *NFTDistributionFilterer) WatchRolledOver(opts *bind.WatchOpts, sink chan<- *NFTDistributionRolledOver) (event.Subscription, error) {

	logs, sub, err := _NFTDistribution.contract.WatchLogs(opts, "RolledOver")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDistributionRolledOver)
				if err := _NFTDistribution.contract.UnpackLog(event, "RolledOver", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRolledOver is a log parse operation binding the contract event 0xd2bce5b0f3cb0c9c70f1761d4ddcaf6236907e4b29bb4a7546b6a0c64be90cf8.
//
// Solidity: event RolledOver(bool status)
func (_NFTDistribution *NFTDistributionFilterer) ParseRolledOver(log types.Log) (*NFTDistributionRolledOver, error) {
	event := new(NFTDistributionRolledOver)
	if err := _NFTDistribution.contract.UnpackLog(event, "RolledOver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
