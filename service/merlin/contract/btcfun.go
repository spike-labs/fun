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
	_ = abi.ConvertType
)

// SetsKeyValue is an auto generated low-level Go binding around an user-defined struct.
type SetsKeyValue struct {
	Key   [32]byte
	Value *big.Int
}

// BtcfunMetaData contains all meta data concerning the Btcfun contract.
var BtcfunMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"volume\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"offered\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalOffered\",\"type\":\"uint256\"}],\"name\":\"Completed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quota\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"CreatePool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"offered\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"name\":\"Offer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refunded\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"currency\",\"type\":\"address\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"airdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"airdropAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"amounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"checkAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimedOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"collect\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quota\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pre\",\"type\":\"uint256\"}],\"name\":\"createPool\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"currencies\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"expiries\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"keys\",\"type\":\"bytes32[]\"}],\"name\":\"gets\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"}],\"name\":\"initSetable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"offer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"offeredOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"offerorN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"offerors\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"quotas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"refundAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"refundedOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structSets.KeyValue[]\",\"name\":\"s\",\"type\":\"tuple[]\"}],\"name\":\"sets_\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"starts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supplies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalOffered\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BtcfunABI is the input ABI used to generate the binding from.
// Deprecated: Use BtcfunMetaData.ABI instead.
var BtcfunABI = BtcfunMetaData.ABI

// Btcfun is an auto generated Go binding around an Ethereum contract.
type Btcfun struct {
	BtcfunCaller     // Read-only binding to the contract
	BtcfunTransactor // Write-only binding to the contract
	BtcfunFilterer   // Log filterer for contract events
}

// BtcfunCaller is an auto generated read-only Go binding around an Ethereum contract.
type BtcfunCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BtcfunTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BtcfunTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BtcfunFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BtcfunFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BtcfunSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BtcfunSession struct {
	Contract     *Btcfun           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BtcfunCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BtcfunCallerSession struct {
	Contract *BtcfunCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BtcfunTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BtcfunTransactorSession struct {
	Contract     *BtcfunTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BtcfunRaw is an auto generated low-level Go binding around an Ethereum contract.
type BtcfunRaw struct {
	Contract *Btcfun // Generic contract binding to access the raw methods on
}

// BtcfunCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BtcfunCallerRaw struct {
	Contract *BtcfunCaller // Generic read-only contract binding to access the raw methods on
}

// BtcfunTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BtcfunTransactorRaw struct {
	Contract *BtcfunTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBtcfun creates a new instance of Btcfun, bound to a specific deployed contract.
func NewBtcfun(address common.Address, backend bind.ContractBackend) (*Btcfun, error) {
	contract, err := bindBtcfun(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Btcfun{BtcfunCaller: BtcfunCaller{contract: contract}, BtcfunTransactor: BtcfunTransactor{contract: contract}, BtcfunFilterer: BtcfunFilterer{contract: contract}}, nil
}

// NewBtcfunCaller creates a new read-only instance of Btcfun, bound to a specific deployed contract.
func NewBtcfunCaller(address common.Address, caller bind.ContractCaller) (*BtcfunCaller, error) {
	contract, err := bindBtcfun(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BtcfunCaller{contract: contract}, nil
}

// NewBtcfunTransactor creates a new write-only instance of Btcfun, bound to a specific deployed contract.
func NewBtcfunTransactor(address common.Address, transactor bind.ContractTransactor) (*BtcfunTransactor, error) {
	contract, err := bindBtcfun(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BtcfunTransactor{contract: contract}, nil
}

// NewBtcfunFilterer creates a new log filterer instance of Btcfun, bound to a specific deployed contract.
func NewBtcfunFilterer(address common.Address, filterer bind.ContractFilterer) (*BtcfunFilterer, error) {
	contract, err := bindBtcfun(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BtcfunFilterer{contract: contract}, nil
}

// bindBtcfun binds a generic wrapper to an already deployed contract.
func bindBtcfun(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BtcfunMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Btcfun *BtcfunRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Btcfun.Contract.BtcfunCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Btcfun *BtcfunRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Btcfun.Contract.BtcfunTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Btcfun *BtcfunRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Btcfun.Contract.BtcfunTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Btcfun *BtcfunCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Btcfun.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Btcfun *BtcfunTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Btcfun.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Btcfun *BtcfunTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Btcfun.Contract.contract.Transact(opts, method, params...)
}

// Amounts is a free data retrieval call binding the contract method 0x55a3b2c1.
//
// Solidity: function amounts(address ) view returns(uint256)
func (_Btcfun *BtcfunCaller) Amounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Btcfun.contract.Call(opts, &out, "amounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Amounts is a free data retrieval call binding the contract method 0x55a3b2c1.
//
// Solidity: function amounts(address ) view returns(uint256)
func (_Btcfun *BtcfunSession) Amounts(arg0 common.Address) (*big.Int, error) {
	return _Btcfun.Contract.Amounts(&_Btcfun.CallOpts, arg0)
}

// Amounts is a free data retrieval call binding the contract method 0x55a3b2c1.
//
// Solidity: function amounts(address ) view returns(uint256)
func (_Btcfun *BtcfunCallerSession) Amounts(arg0 common.Address) (*big.Int, error) {
	return _Btcfun.Contract.Amounts(&_Btcfun.CallOpts, arg0)
}

// CheckAmount is a free data retrieval call binding the contract method 0xd6b5983d.
//
// Solidity: function checkAmount(address token, uint256 amount) view returns(uint256)
func (_Btcfun *BtcfunCaller) CheckAmount(opts *bind.CallOpts, token common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Btcfun.contract.Call(opts, &out, "checkAmount", token, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckAmount is a free data retrieval call binding the contract method 0xd6b5983d.
//
// Solidity: function checkAmount(address token, uint256 amount) view returns(uint256)
func (_Btcfun *BtcfunSession) CheckAmount(token common.Address, amount *big.Int) (*big.Int, error) {
	return _Btcfun.Contract.CheckAmount(&_Btcfun.CallOpts, token, amount)
}

// CheckAmount is a free data retrieval call binding the contract method 0xd6b5983d.
//
// Solidity: function checkAmount(address token, uint256 amount) view returns(uint256)
func (_Btcfun *BtcfunCallerSession) CheckAmount(token common.Address, amount *big.Int) (*big.Int, error) {
	return _Btcfun.Contract.CheckAmount(&_Btcfun.CallOpts, token, amount)
}

// ClaimedOf is a free data retrieval call binding the contract method 0x9cf00275.
//
// Solidity: function claimedOf(address , address ) view returns(uint256)
func (_Btcfun *BtcfunCaller) ClaimedOf(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Btcfun.contract.Call(opts, &out, "claimedOf", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedOf is a free data retrieval call binding the contract method 0x9cf00275.
//
// Solidity: function claimedOf(address , address ) view returns(uint256)
func (_Btcfun *BtcfunSession) ClaimedOf(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Btcfun.Contract.ClaimedOf(&_Btcfun.CallOpts, arg0, arg1)
}

// ClaimedOf is a free data retrieval call binding the contract method 0x9cf00275.
//
// Solidity: function claimedOf(address , address ) view returns(uint256)
func (_Btcfun *BtcfunCallerSession) ClaimedOf(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Btcfun.Contract.ClaimedOf(&_Btcfun.CallOpts, arg0, arg1)
}

// Currencies is a free data retrieval call binding the contract method 0x6036cba3.
//
// Solidity: function currencies(address ) view returns(address)
func (_Btcfun *BtcfunCaller) Currencies(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Btcfun.contract.Call(opts, &out, "currencies", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Currencies is a free data retrieval call binding the contract method 0x6036cba3.
//
// Solidity: function currencies(address ) view returns(address)
func (_Btcfun *BtcfunSession) Currencies(arg0 common.Address) (common.Address, error) {
	return _Btcfun.Contract.Currencies(&_Btcfun.CallOpts, arg0)
}

// Currencies is a free data retrieval call binding the contract method 0x6036cba3.
//
// Solidity: function currencies(address ) view returns(address)
func (_Btcfun *BtcfunCallerSession) Currencies(arg0 common.Address) (common.Address, error) {
	return _Btcfun.Contract.Currencies(&_Btcfun.CallOpts, arg0)
}

// Expiries is a free data retrieval call binding the contract method 0x1fd6a5dc.
//
// Solidity: function expiries(address ) view returns(uint256)
func (_Btcfun *BtcfunCaller) Expiries(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Btcfun.contract.Call(opts, &out, "expiries", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Expiries is a free data retrieval call binding the contract method 0x1fd6a5dc.
//
// Solidity: function expiries(address ) view returns(uint256)
func (_Btcfun *BtcfunSession) Expiries(arg0 common.Address) (*big.Int, error) {
	return _Btcfun.Contract.Expiries(&_Btcfun.CallOpts, arg0)
}

// Expiries is a free data retrieval call binding the contract method 0x1fd6a5dc.
//
// Solidity: function expiries(address ) view returns(uint256)
func (_Btcfun *BtcfunCallerSession) Expiries(arg0 common.Address) (*big.Int, error) {
	return _Btcfun.Contract.Expiries(&_Btcfun.CallOpts, arg0)
}

// Gets is a free data retrieval call binding the contract method 0x8da7bd3e.
//
// Solidity: function gets(bytes32[] keys) view returns(uint256[] values)
func (_Btcfun *BtcfunCaller) Gets(opts *bind.CallOpts, keys [][32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _Btcfun.contract.Call(opts, &out, "gets", keys)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// Gets is a free data retrieval call binding the contract method 0x8da7bd3e.
//
// Solidity: function gets(bytes32[] keys) view returns(uint256[] values)
func (_Btcfun *BtcfunSession) Gets(keys [][32]byte) ([]*big.Int, error) {
	return _Btcfun.Contract.Gets(&_Btcfun.CallOpts, keys)
}

// Gets is a free data retrieval call binding the contract method 0x8da7bd3e.
//
// Solidity: function gets(bytes32[] keys) view returns(uint256[] values)
func (_Btcfun *BtcfunCallerSession) Gets(keys [][32]byte) ([]*big.Int, error) {
	return _Btcfun.Contract.Gets(&_Btcfun.CallOpts, keys)
}

// IsGovernor is a free data retrieval call binding the contract method 0xc7af3352.
//
// Solidity: function isGovernor() view returns(bool)
func (_Btcfun *BtcfunCaller) IsGovernor(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Btcfun.contract.Call(opts, &out, "isGovernor")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGovernor is a free data retrieval call binding the contract method 0xc7af3352.
//
// Solidity: function isGovernor() view returns(bool)
func (_Btcfun *BtcfunSession) IsGovernor() (bool, error) {
	return _Btcfun.Contract.IsGovernor(&_Btcfun.CallOpts)
}

// IsGovernor is a free data retrieval call binding the contract method 0xc7af3352.
//
// Solidity: function isGovernor() view returns(bool)
func (_Btcfun *BtcfunCallerSession) IsGovernor() (bool, error) {
	return _Btcfun.Contract.IsGovernor(&_Btcfun.CallOpts)
}

// OfferedOf is a free data retrieval call binding the contract method 0x3863300b.
//
// Solidity: function offeredOf(address , address ) view returns(uint256)
func (_Btcfun *BtcfunCaller) OfferedOf(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Btcfun.contract.Call(opts, &out, "offeredOf", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OfferedOf is a free data retrieval call binding the contract method 0x3863300b.
//
// Solidity: function offeredOf(address , address ) view returns(uint256)
func (_Btcfun *BtcfunSession) OfferedOf(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Btcfun.Contract.OfferedOf(&_Btcfun.CallOpts, arg0, arg1)
}

// OfferedOf is a free data retrieval call binding the contract method 0x3863300b.
//
// Solidity: function offeredOf(address , address ) view returns(uint256)
func (_Btcfun *BtcfunCallerSession) OfferedOf(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Btcfun.Contract.OfferedOf(&_Btcfun.CallOpts, arg0, arg1)
}

// OfferorN is a free data retrieval call binding the contract method 0x1685aee8.
//
// Solidity: function offerorN(address token) view returns(uint256)
func (_Btcfun *BtcfunCaller) OfferorN(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Btcfun.contract.Call(opts, &out, "offerorN", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OfferorN is a free data retrieval call binding the contract method 0x1685aee8.
//
// Solidity: function offerorN(address token) view returns(uint256)
func (_Btcfun *BtcfunSession) OfferorN(token common.Address) (*big.Int, error) {
	return _Btcfun.Contract.OfferorN(&_Btcfun.CallOpts, token)
}

// OfferorN is a free data retrieval call binding the contract method 0x1685aee8.
//
// Solidity: function offerorN(address token) view returns(uint256)
func (_Btcfun *BtcfunCallerSession) OfferorN(token common.Address) (*big.Int, error) {
	return _Btcfun.Contract.OfferorN(&_Btcfun.CallOpts, token)
}

// Offerors is a free data retrieval call binding the contract method 0xa8c58265.
//
// Solidity: function offerors(address , uint256 ) view returns(address)
func (_Btcfun *BtcfunCaller) Offerors(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Btcfun.contract.Call(opts, &out, "offerors", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Offerors is a free data retrieval call binding the contract method 0xa8c58265.
//
// Solidity: function offerors(address , uint256 ) view returns(address)
func (_Btcfun *BtcfunSession) Offerors(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Btcfun.Contract.Offerors(&_Btcfun.CallOpts, arg0, arg1)
}

// Offerors is a free data retrieval call binding the contract method 0xa8c58265.
//
// Solidity: function offerors(address , uint256 ) view returns(address)
func (_Btcfun *BtcfunCallerSession) Offerors(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Btcfun.Contract.Offerors(&_Btcfun.CallOpts, arg0, arg1)
}

// Quotas is a free data retrieval call binding the contract method 0xc33342e9.
//
// Solidity: function quotas(address ) view returns(uint256)
func (_Btcfun *BtcfunCaller) Quotas(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Btcfun.contract.Call(opts, &out, "quotas", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quotas is a free data retrieval call binding the contract method 0xc33342e9.
//
// Solidity: function quotas(address ) view returns(uint256)
func (_Btcfun *BtcfunSession) Quotas(arg0 common.Address) (*big.Int, error) {
	return _Btcfun.Contract.Quotas(&_Btcfun.CallOpts, arg0)
}

// Quotas is a free data retrieval call binding the contract method 0xc33342e9.
//
// Solidity: function quotas(address ) view returns(uint256)
func (_Btcfun *BtcfunCallerSession) Quotas(arg0 common.Address) (*big.Int, error) {
	return _Btcfun.Contract.Quotas(&_Btcfun.CallOpts, arg0)
}

// RefundedOf is a free data retrieval call binding the contract method 0xb6e6efbf.
//
// Solidity: function refundedOf(address , address ) view returns(uint256)
func (_Btcfun *BtcfunCaller) RefundedOf(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Btcfun.contract.Call(opts, &out, "refundedOf", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RefundedOf is a free data retrieval call binding the contract method 0xb6e6efbf.
//
// Solidity: function refundedOf(address , address ) view returns(uint256)
func (_Btcfun *BtcfunSession) RefundedOf(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Btcfun.Contract.RefundedOf(&_Btcfun.CallOpts, arg0, arg1)
}

// RefundedOf is a free data retrieval call binding the contract method 0xb6e6efbf.
//
// Solidity: function refundedOf(address , address ) view returns(uint256)
func (_Btcfun *BtcfunCallerSession) RefundedOf(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Btcfun.Contract.RefundedOf(&_Btcfun.CallOpts, arg0, arg1)
}

// Starts is a free data retrieval call binding the contract method 0xb6c238b5.
//
// Solidity: function starts(address ) view returns(uint256)
func (_Btcfun *BtcfunCaller) Starts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Btcfun.contract.Call(opts, &out, "starts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Starts is a free data retrieval call binding the contract method 0xb6c238b5.
//
// Solidity: function starts(address ) view returns(uint256)
func (_Btcfun *BtcfunSession) Starts(arg0 common.Address) (*big.Int, error) {
	return _Btcfun.Contract.Starts(&_Btcfun.CallOpts, arg0)
}

// Starts is a free data retrieval call binding the contract method 0xb6c238b5.
//
// Solidity: function starts(address ) view returns(uint256)
func (_Btcfun *BtcfunCallerSession) Starts(arg0 common.Address) (*big.Int, error) {
	return _Btcfun.Contract.Starts(&_Btcfun.CallOpts, arg0)
}

// Supplies is a free data retrieval call binding the contract method 0x274cee31.
//
// Solidity: function supplies(address ) view returns(uint256)
func (_Btcfun *BtcfunCaller) Supplies(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Btcfun.contract.Call(opts, &out, "supplies", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Supplies is a free data retrieval call binding the contract method 0x274cee31.
//
// Solidity: function supplies(address ) view returns(uint256)
func (_Btcfun *BtcfunSession) Supplies(arg0 common.Address) (*big.Int, error) {
	return _Btcfun.Contract.Supplies(&_Btcfun.CallOpts, arg0)
}

// Supplies is a free data retrieval call binding the contract method 0x274cee31.
//
// Solidity: function supplies(address ) view returns(uint256)
func (_Btcfun *BtcfunCallerSession) Supplies(arg0 common.Address) (*big.Int, error) {
	return _Btcfun.Contract.Supplies(&_Btcfun.CallOpts, arg0)
}

// TokenIds is a free data retrieval call binding the contract method 0xfc97a303.
//
// Solidity: function tokenIds(address ) view returns(uint256)
func (_Btcfun *BtcfunCaller) TokenIds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Btcfun.contract.Call(opts, &out, "tokenIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIds is a free data retrieval call binding the contract method 0xfc97a303.
//
// Solidity: function tokenIds(address ) view returns(uint256)
func (_Btcfun *BtcfunSession) TokenIds(arg0 common.Address) (*big.Int, error) {
	return _Btcfun.Contract.TokenIds(&_Btcfun.CallOpts, arg0)
}

// TokenIds is a free data retrieval call binding the contract method 0xfc97a303.
//
// Solidity: function tokenIds(address ) view returns(uint256)
func (_Btcfun *BtcfunCallerSession) TokenIds(arg0 common.Address) (*big.Int, error) {
	return _Btcfun.Contract.TokenIds(&_Btcfun.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x04c2320b.
//
// Solidity: function tokens(string ) view returns(address)
func (_Btcfun *BtcfunCaller) Tokens(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _Btcfun.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x04c2320b.
//
// Solidity: function tokens(string ) view returns(address)
func (_Btcfun *BtcfunSession) Tokens(arg0 string) (common.Address, error) {
	return _Btcfun.Contract.Tokens(&_Btcfun.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x04c2320b.
//
// Solidity: function tokens(string ) view returns(address)
func (_Btcfun *BtcfunCallerSession) Tokens(arg0 string) (common.Address, error) {
	return _Btcfun.Contract.Tokens(&_Btcfun.CallOpts, arg0)
}

// TotalOffered is a free data retrieval call binding the contract method 0x38d8afd3.
//
// Solidity: function totalOffered(address ) view returns(uint256)
func (_Btcfun *BtcfunCaller) TotalOffered(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Btcfun.contract.Call(opts, &out, "totalOffered", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalOffered is a free data retrieval call binding the contract method 0x38d8afd3.
//
// Solidity: function totalOffered(address ) view returns(uint256)
func (_Btcfun *BtcfunSession) TotalOffered(arg0 common.Address) (*big.Int, error) {
	return _Btcfun.Contract.TotalOffered(&_Btcfun.CallOpts, arg0)
}

// TotalOffered is a free data retrieval call binding the contract method 0x38d8afd3.
//
// Solidity: function totalOffered(address ) view returns(uint256)
func (_Btcfun *BtcfunCallerSession) TotalOffered(arg0 common.Address) (*big.Int, error) {
	return _Btcfun.Contract.TotalOffered(&_Btcfun.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Btcfun *BtcfunCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Btcfun.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Btcfun *BtcfunSession) Version() (string, error) {
	return _Btcfun.Contract.Version(&_Btcfun.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Btcfun *BtcfunCallerSession) Version() (string, error) {
	return _Btcfun.Contract.Version(&_Btcfun.CallOpts)
}

// Airdrop is a paid mutator transaction binding the contract method 0xe1bc2967.
//
// Solidity: function airdrop(address token, uint256 offset, uint256 count) returns()
func (_Btcfun *BtcfunTransactor) Airdrop(opts *bind.TransactOpts, token common.Address, offset *big.Int, count *big.Int) (*types.Transaction, error) {
	return _Btcfun.contract.Transact(opts, "airdrop", token, offset, count)
}

// Airdrop is a paid mutator transaction binding the contract method 0xe1bc2967.
//
// Solidity: function airdrop(address token, uint256 offset, uint256 count) returns()
func (_Btcfun *BtcfunSession) Airdrop(token common.Address, offset *big.Int, count *big.Int) (*types.Transaction, error) {
	return _Btcfun.Contract.Airdrop(&_Btcfun.TransactOpts, token, offset, count)
}

// Airdrop is a paid mutator transaction binding the contract method 0xe1bc2967.
//
// Solidity: function airdrop(address token, uint256 offset, uint256 count) returns()
func (_Btcfun *BtcfunTransactorSession) Airdrop(token common.Address, offset *big.Int, count *big.Int) (*types.Transaction, error) {
	return _Btcfun.Contract.Airdrop(&_Btcfun.TransactOpts, token, offset, count)
}

// AirdropAccount is a paid mutator transaction binding the contract method 0x1313327a.
//
// Solidity: function airdropAccount(address token, address sender) returns()
func (_Btcfun *BtcfunTransactor) AirdropAccount(opts *bind.TransactOpts, token common.Address, sender common.Address) (*types.Transaction, error) {
	return _Btcfun.contract.Transact(opts, "airdropAccount", token, sender)
}

// AirdropAccount is a paid mutator transaction binding the contract method 0x1313327a.
//
// Solidity: function airdropAccount(address token, address sender) returns()
func (_Btcfun *BtcfunSession) AirdropAccount(token common.Address, sender common.Address) (*types.Transaction, error) {
	return _Btcfun.Contract.AirdropAccount(&_Btcfun.TransactOpts, token, sender)
}

// AirdropAccount is a paid mutator transaction binding the contract method 0x1313327a.
//
// Solidity: function airdropAccount(address token, address sender) returns()
func (_Btcfun *BtcfunTransactorSession) AirdropAccount(token common.Address, sender common.Address) (*types.Transaction, error) {
	return _Btcfun.Contract.AirdropAccount(&_Btcfun.TransactOpts, token, sender)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address token) returns()
func (_Btcfun *BtcfunTransactor) Collect(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Btcfun.contract.Transact(opts, "collect", token)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address token) returns()
func (_Btcfun *BtcfunSession) Collect(token common.Address) (*types.Transaction, error) {
	return _Btcfun.Contract.Collect(&_Btcfun.TransactOpts, token)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(address token) returns()
func (_Btcfun *BtcfunTransactorSession) Collect(token common.Address) (*types.Transaction, error) {
	return _Btcfun.Contract.Collect(&_Btcfun.TransactOpts, token)
}

// CreatePool is a paid mutator transaction binding the contract method 0xb6493e6b.
//
// Solidity: function createPool(address token, uint256 supply, address currency, uint256 amount, uint256 quota, uint256 start, uint256 expiry, uint256 pre) payable returns()
func (_Btcfun *BtcfunTransactor) CreatePool(opts *bind.TransactOpts, token common.Address, supply *big.Int, currency common.Address, amount *big.Int, quota *big.Int, start *big.Int, expiry *big.Int, pre *big.Int) (*types.Transaction, error) {
	return _Btcfun.contract.Transact(opts, "createPool", token, supply, currency, amount, quota, start, expiry, pre)
}

// CreatePool is a paid mutator transaction binding the contract method 0xb6493e6b.
//
// Solidity: function createPool(address token, uint256 supply, address currency, uint256 amount, uint256 quota, uint256 start, uint256 expiry, uint256 pre) payable returns()
func (_Btcfun *BtcfunSession) CreatePool(token common.Address, supply *big.Int, currency common.Address, amount *big.Int, quota *big.Int, start *big.Int, expiry *big.Int, pre *big.Int) (*types.Transaction, error) {
	return _Btcfun.Contract.CreatePool(&_Btcfun.TransactOpts, token, supply, currency, amount, quota, start, expiry, pre)
}

// CreatePool is a paid mutator transaction binding the contract method 0xb6493e6b.
//
// Solidity: function createPool(address token, uint256 supply, address currency, uint256 amount, uint256 quota, uint256 start, uint256 expiry, uint256 pre) payable returns()
func (_Btcfun *BtcfunTransactorSession) CreatePool(token common.Address, supply *big.Int, currency common.Address, amount *big.Int, quota *big.Int, start *big.Int, expiry *big.Int, pre *big.Int) (*types.Transaction, error) {
	return _Btcfun.Contract.CreatePool(&_Btcfun.TransactOpts, token, supply, currency, amount, quota, start, expiry, pre)
}

// InitSetable is a paid mutator transaction binding the contract method 0xa1b4e794.
//
// Solidity: function initSetable(address msgSender) returns()
func (_Btcfun *BtcfunTransactor) InitSetable(opts *bind.TransactOpts, msgSender common.Address) (*types.Transaction, error) {
	return _Btcfun.contract.Transact(opts, "initSetable", msgSender)
}

// InitSetable is a paid mutator transaction binding the contract method 0xa1b4e794.
//
// Solidity: function initSetable(address msgSender) returns()
func (_Btcfun *BtcfunSession) InitSetable(msgSender common.Address) (*types.Transaction, error) {
	return _Btcfun.Contract.InitSetable(&_Btcfun.TransactOpts, msgSender)
}

// InitSetable is a paid mutator transaction binding the contract method 0xa1b4e794.
//
// Solidity: function initSetable(address msgSender) returns()
func (_Btcfun *BtcfunTransactorSession) InitSetable(msgSender common.Address) (*types.Transaction, error) {
	return _Btcfun.Contract.InitSetable(&_Btcfun.TransactOpts, msgSender)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address msgSender) returns()
func (_Btcfun *BtcfunTransactor) Initialize(opts *bind.TransactOpts, msgSender common.Address) (*types.Transaction, error) {
	return _Btcfun.contract.Transact(opts, "initialize", msgSender)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address msgSender) returns()
func (_Btcfun *BtcfunSession) Initialize(msgSender common.Address) (*types.Transaction, error) {
	return _Btcfun.Contract.Initialize(&_Btcfun.TransactOpts, msgSender)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address msgSender) returns()
func (_Btcfun *BtcfunTransactorSession) Initialize(msgSender common.Address) (*types.Transaction, error) {
	return _Btcfun.Contract.Initialize(&_Btcfun.TransactOpts, msgSender)
}

// Offer is a paid mutator transaction binding the contract method 0x6e978426.
//
// Solidity: function offer(address token, uint256 amount, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Btcfun *BtcfunTransactor) Offer(opts *bind.TransactOpts, token common.Address, amount *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Btcfun.contract.Transact(opts, "offer", token, amount, expiry, v, r, s)
}

// Offer is a paid mutator transaction binding the contract method 0x6e978426.
//
// Solidity: function offer(address token, uint256 amount, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Btcfun *BtcfunSession) Offer(token common.Address, amount *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Btcfun.Contract.Offer(&_Btcfun.TransactOpts, token, amount, expiry, v, r, s)
}

// Offer is a paid mutator transaction binding the contract method 0x6e978426.
//
// Solidity: function offer(address token, uint256 amount, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Btcfun *BtcfunTransactorSession) Offer(token common.Address, amount *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Btcfun.Contract.Offer(&_Btcfun.TransactOpts, token, amount, expiry, v, r, s)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Btcfun *BtcfunTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Btcfun.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Btcfun *BtcfunSession) Pause() (*types.Transaction, error) {
	return _Btcfun.Contract.Pause(&_Btcfun.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Btcfun *BtcfunTransactorSession) Pause() (*types.Transaction, error) {
	return _Btcfun.Contract.Pause(&_Btcfun.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x18fd8903.
//
// Solidity: function refund(address token, uint256 offset, uint256 count) returns()
func (_Btcfun *BtcfunTransactor) Refund(opts *bind.TransactOpts, token common.Address, offset *big.Int, count *big.Int) (*types.Transaction, error) {
	return _Btcfun.contract.Transact(opts, "refund", token, offset, count)
}

// Refund is a paid mutator transaction binding the contract method 0x18fd8903.
//
// Solidity: function refund(address token, uint256 offset, uint256 count) returns()
func (_Btcfun *BtcfunSession) Refund(token common.Address, offset *big.Int, count *big.Int) (*types.Transaction, error) {
	return _Btcfun.Contract.Refund(&_Btcfun.TransactOpts, token, offset, count)
}

// Refund is a paid mutator transaction binding the contract method 0x18fd8903.
//
// Solidity: function refund(address token, uint256 offset, uint256 count) returns()
func (_Btcfun *BtcfunTransactorSession) Refund(token common.Address, offset *big.Int, count *big.Int) (*types.Transaction, error) {
	return _Btcfun.Contract.Refund(&_Btcfun.TransactOpts, token, offset, count)
}

// RefundAccount is a paid mutator transaction binding the contract method 0x1b41c23d.
//
// Solidity: function refundAccount(address token, address sender) returns()
func (_Btcfun *BtcfunTransactor) RefundAccount(opts *bind.TransactOpts, token common.Address, sender common.Address) (*types.Transaction, error) {
	return _Btcfun.contract.Transact(opts, "refundAccount", token, sender)
}

// RefundAccount is a paid mutator transaction binding the contract method 0x1b41c23d.
//
// Solidity: function refundAccount(address token, address sender) returns()
func (_Btcfun *BtcfunSession) RefundAccount(token common.Address, sender common.Address) (*types.Transaction, error) {
	return _Btcfun.Contract.RefundAccount(&_Btcfun.TransactOpts, token, sender)
}

// RefundAccount is a paid mutator transaction binding the contract method 0x1b41c23d.
//
// Solidity: function refundAccount(address token, address sender) returns()
func (_Btcfun *BtcfunTransactorSession) RefundAccount(token common.Address, sender common.Address) (*types.Transaction, error) {
	return _Btcfun.Contract.RefundAccount(&_Btcfun.TransactOpts, token, sender)
}

// Sets is a paid mutator transaction binding the contract method 0xdc87add9.
//
// Solidity: function sets_((bytes32,uint256)[] s) returns()
func (_Btcfun *BtcfunTransactor) Sets(opts *bind.TransactOpts, s []SetsKeyValue) (*types.Transaction, error) {
	return _Btcfun.contract.Transact(opts, "sets_", s)
}

// Sets is a paid mutator transaction binding the contract method 0xdc87add9.
//
// Solidity: function sets_((bytes32,uint256)[] s) returns()
func (_Btcfun *BtcfunSession) Sets(s []SetsKeyValue) (*types.Transaction, error) {
	return _Btcfun.Contract.Sets(&_Btcfun.TransactOpts, s)
}

// Sets is a paid mutator transaction binding the contract method 0xdc87add9.
//
// Solidity: function sets_((bytes32,uint256)[] s) returns()
func (_Btcfun *BtcfunTransactorSession) Sets(s []SetsKeyValue) (*types.Transaction, error) {
	return _Btcfun.Contract.Sets(&_Btcfun.TransactOpts, s)
}

// Unlock is a paid mutator transaction binding the contract method 0x2f6c493c.
//
// Solidity: function unlock(address token) returns()
func (_Btcfun *BtcfunTransactor) Unlock(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Btcfun.contract.Transact(opts, "unlock", token)
}

// Unlock is a paid mutator transaction binding the contract method 0x2f6c493c.
//
// Solidity: function unlock(address token) returns()
func (_Btcfun *BtcfunSession) Unlock(token common.Address) (*types.Transaction, error) {
	return _Btcfun.Contract.Unlock(&_Btcfun.TransactOpts, token)
}

// Unlock is a paid mutator transaction binding the contract method 0x2f6c493c.
//
// Solidity: function unlock(address token) returns()
func (_Btcfun *BtcfunTransactorSession) Unlock(token common.Address) (*types.Transaction, error) {
	return _Btcfun.Contract.Unlock(&_Btcfun.TransactOpts, token)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Btcfun *BtcfunTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Btcfun.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Btcfun *BtcfunSession) Unpause() (*types.Transaction, error) {
	return _Btcfun.Contract.Unpause(&_Btcfun.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Btcfun *BtcfunTransactorSession) Unpause() (*types.Transaction, error) {
	return _Btcfun.Contract.Unpause(&_Btcfun.TransactOpts)
}

// BtcfunClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Btcfun contract.
type BtcfunClaimIterator struct {
	Event *BtcfunClaim // Event containing the contract specifics and raw log

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
func (it *BtcfunClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcfunClaim)
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
		it.Event = new(BtcfunClaim)
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
func (it *BtcfunClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcfunClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcfunClaim represents a Claim event raised by the Btcfun contract.
type BtcfunClaim struct {
	Sender   common.Address
	Token    common.Address
	To       common.Address
	Volume   *big.Int
	Currency common.Address
	Offered  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x2ac7591f882fa86fbfbc1dd1b39f0dde42bc4d3c0d5015f553c498f843c495a8.
//
// Solidity: event Claim(address sender, address indexed token, address indexed to, uint256 volume, address indexed currency, uint256 offered)
func (_Btcfun *BtcfunFilterer) FilterClaim(opts *bind.FilterOpts, token []common.Address, to []common.Address, currency []common.Address) (*BtcfunClaimIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	logs, sub, err := _Btcfun.contract.FilterLogs(opts, "Claim", tokenRule, toRule, currencyRule)
	if err != nil {
		return nil, err
	}
	return &BtcfunClaimIterator{contract: _Btcfun.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x2ac7591f882fa86fbfbc1dd1b39f0dde42bc4d3c0d5015f553c498f843c495a8.
//
// Solidity: event Claim(address sender, address indexed token, address indexed to, uint256 volume, address indexed currency, uint256 offered)
func (_Btcfun *BtcfunFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *BtcfunClaim, token []common.Address, to []common.Address, currency []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	logs, sub, err := _Btcfun.contract.WatchLogs(opts, "Claim", tokenRule, toRule, currencyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcfunClaim)
				if err := _Btcfun.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x2ac7591f882fa86fbfbc1dd1b39f0dde42bc4d3c0d5015f553c498f843c495a8.
//
// Solidity: event Claim(address sender, address indexed token, address indexed to, uint256 volume, address indexed currency, uint256 offered)
func (_Btcfun *BtcfunFilterer) ParseClaim(log types.Log) (*BtcfunClaim, error) {
	event := new(BtcfunClaim)
	if err := _Btcfun.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtcfunCompletedIterator is returned from FilterCompleted and is used to iterate over the raw logs and unpacked data for Completed events raised by the Btcfun contract.
type BtcfunCompletedIterator struct {
	Event *BtcfunCompleted // Event containing the contract specifics and raw log

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
func (it *BtcfunCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcfunCompleted)
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
		it.Event = new(BtcfunCompleted)
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
func (it *BtcfunCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcfunCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcfunCompleted represents a Completed event raised by the Btcfun contract.
type BtcfunCompleted struct {
	Token        common.Address
	Currency     common.Address
	TotalOffered *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompleted is a free log retrieval operation binding the contract event 0xcdb9f4e58f1d1744a76d4003dc0c940c04eb85e7c04b3ce13f230594318c995d.
//
// Solidity: event Completed(address indexed token, address indexed currency, uint256 totalOffered)
func (_Btcfun *BtcfunFilterer) FilterCompleted(opts *bind.FilterOpts, token []common.Address, currency []common.Address) (*BtcfunCompletedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	logs, sub, err := _Btcfun.contract.FilterLogs(opts, "Completed", tokenRule, currencyRule)
	if err != nil {
		return nil, err
	}
	return &BtcfunCompletedIterator{contract: _Btcfun.contract, event: "Completed", logs: logs, sub: sub}, nil
}

// WatchCompleted is a free log subscription operation binding the contract event 0xcdb9f4e58f1d1744a76d4003dc0c940c04eb85e7c04b3ce13f230594318c995d.
//
// Solidity: event Completed(address indexed token, address indexed currency, uint256 totalOffered)
func (_Btcfun *BtcfunFilterer) WatchCompleted(opts *bind.WatchOpts, sink chan<- *BtcfunCompleted, token []common.Address, currency []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	logs, sub, err := _Btcfun.contract.WatchLogs(opts, "Completed", tokenRule, currencyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcfunCompleted)
				if err := _Btcfun.contract.UnpackLog(event, "Completed", log); err != nil {
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

// ParseCompleted is a log parse operation binding the contract event 0xcdb9f4e58f1d1744a76d4003dc0c940c04eb85e7c04b3ce13f230594318c995d.
//
// Solidity: event Completed(address indexed token, address indexed currency, uint256 totalOffered)
func (_Btcfun *BtcfunFilterer) ParseCompleted(log types.Log) (*BtcfunCompleted, error) {
	event := new(BtcfunCompleted)
	if err := _Btcfun.contract.UnpackLog(event, "Completed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtcfunCreatePoolIterator is returned from FilterCreatePool and is used to iterate over the raw logs and unpacked data for CreatePool events raised by the Btcfun contract.
type BtcfunCreatePoolIterator struct {
	Event *BtcfunCreatePool // Event containing the contract specifics and raw log

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
func (it *BtcfunCreatePoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcfunCreatePool)
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
		it.Event = new(BtcfunCreatePool)
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
func (it *BtcfunCreatePoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcfunCreatePoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcfunCreatePool represents a CreatePool event raised by the Btcfun contract.
type BtcfunCreatePool struct {
	Name     string
	Supply   *big.Int
	Token    common.Address
	Currency common.Address
	Amount   *big.Int
	Quota    *big.Int
	Start    *big.Int
	Expiry   *big.Int
	Pool     common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCreatePool is a free log retrieval operation binding the contract event 0xd57e9547ce0fa1b6d62bf34ca193cf6b77dd72e8b063c9ee12094ed8d04ac90c.
//
// Solidity: event CreatePool(string name, uint256 supply, address indexed token, address indexed currency, uint256 amount, uint256 quota, uint256 start, uint256 expiry, address indexed pool)
func (_Btcfun *BtcfunFilterer) FilterCreatePool(opts *bind.FilterOpts, token []common.Address, currency []common.Address, pool []common.Address) (*BtcfunCreatePoolIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Btcfun.contract.FilterLogs(opts, "CreatePool", tokenRule, currencyRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &BtcfunCreatePoolIterator{contract: _Btcfun.contract, event: "CreatePool", logs: logs, sub: sub}, nil
}

// WatchCreatePool is a free log subscription operation binding the contract event 0xd57e9547ce0fa1b6d62bf34ca193cf6b77dd72e8b063c9ee12094ed8d04ac90c.
//
// Solidity: event CreatePool(string name, uint256 supply, address indexed token, address indexed currency, uint256 amount, uint256 quota, uint256 start, uint256 expiry, address indexed pool)
func (_Btcfun *BtcfunFilterer) WatchCreatePool(opts *bind.WatchOpts, sink chan<- *BtcfunCreatePool, token []common.Address, currency []common.Address, pool []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Btcfun.contract.WatchLogs(opts, "CreatePool", tokenRule, currencyRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcfunCreatePool)
				if err := _Btcfun.contract.UnpackLog(event, "CreatePool", log); err != nil {
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

// ParseCreatePool is a log parse operation binding the contract event 0xd57e9547ce0fa1b6d62bf34ca193cf6b77dd72e8b063c9ee12094ed8d04ac90c.
//
// Solidity: event CreatePool(string name, uint256 supply, address indexed token, address indexed currency, uint256 amount, uint256 quota, uint256 start, uint256 expiry, address indexed pool)
func (_Btcfun *BtcfunFilterer) ParseCreatePool(log types.Log) (*BtcfunCreatePool, error) {
	event := new(BtcfunCreatePool)
	if err := _Btcfun.contract.UnpackLog(event, "CreatePool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtcfunInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Btcfun contract.
type BtcfunInitializedIterator struct {
	Event *BtcfunInitialized // Event containing the contract specifics and raw log

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
func (it *BtcfunInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcfunInitialized)
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
		it.Event = new(BtcfunInitialized)
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
func (it *BtcfunInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcfunInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcfunInitialized represents a Initialized event raised by the Btcfun contract.
type BtcfunInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Btcfun *BtcfunFilterer) FilterInitialized(opts *bind.FilterOpts) (*BtcfunInitializedIterator, error) {

	logs, sub, err := _Btcfun.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BtcfunInitializedIterator{contract: _Btcfun.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Btcfun *BtcfunFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BtcfunInitialized) (event.Subscription, error) {

	logs, sub, err := _Btcfun.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcfunInitialized)
				if err := _Btcfun.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Btcfun *BtcfunFilterer) ParseInitialized(log types.Log) (*BtcfunInitialized, error) {
	event := new(BtcfunInitialized)
	if err := _Btcfun.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtcfunOfferIterator is returned from FilterOffer and is used to iterate over the raw logs and unpacked data for Offer events raised by the Btcfun contract.
type BtcfunOfferIterator struct {
	Event *BtcfunOffer // Event containing the contract specifics and raw log

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
func (it *BtcfunOfferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcfunOffer)
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
		it.Event = new(BtcfunOffer)
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
func (it *BtcfunOfferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcfunOfferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcfunOffer represents a Offer event raised by the Btcfun contract.
type BtcfunOffer struct {
	Token   common.Address
	Sender  common.Address
	Amount  *big.Int
	Offered *big.Int
	Total   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOffer is a free log retrieval operation binding the contract event 0x259a476765a4377de3336cff6c0bda61d38942cd0b9ccc7a2d2f3c1c465d0515.
//
// Solidity: event Offer(address indexed token, address indexed sender, uint256 amount, uint256 offered, uint256 total)
func (_Btcfun *BtcfunFilterer) FilterOffer(opts *bind.FilterOpts, token []common.Address, sender []common.Address) (*BtcfunOfferIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Btcfun.contract.FilterLogs(opts, "Offer", tokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BtcfunOfferIterator{contract: _Btcfun.contract, event: "Offer", logs: logs, sub: sub}, nil
}

// WatchOffer is a free log subscription operation binding the contract event 0x259a476765a4377de3336cff6c0bda61d38942cd0b9ccc7a2d2f3c1c465d0515.
//
// Solidity: event Offer(address indexed token, address indexed sender, uint256 amount, uint256 offered, uint256 total)
func (_Btcfun *BtcfunFilterer) WatchOffer(opts *bind.WatchOpts, sink chan<- *BtcfunOffer, token []common.Address, sender []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Btcfun.contract.WatchLogs(opts, "Offer", tokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcfunOffer)
				if err := _Btcfun.contract.UnpackLog(event, "Offer", log); err != nil {
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

// ParseOffer is a log parse operation binding the contract event 0x259a476765a4377de3336cff6c0bda61d38942cd0b9ccc7a2d2f3c1c465d0515.
//
// Solidity: event Offer(address indexed token, address indexed sender, uint256 amount, uint256 offered, uint256 total)
func (_Btcfun *BtcfunFilterer) ParseOffer(log types.Log) (*BtcfunOffer, error) {
	event := new(BtcfunOffer)
	if err := _Btcfun.contract.UnpackLog(event, "Offer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtcfunPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Btcfun contract.
type BtcfunPauseIterator struct {
	Event *BtcfunPause // Event containing the contract specifics and raw log

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
func (it *BtcfunPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcfunPause)
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
		it.Event = new(BtcfunPause)
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
func (it *BtcfunPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcfunPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcfunPause represents a Pause event raised by the Btcfun contract.
type BtcfunPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Btcfun *BtcfunFilterer) FilterPause(opts *bind.FilterOpts) (*BtcfunPauseIterator, error) {

	logs, sub, err := _Btcfun.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &BtcfunPauseIterator{contract: _Btcfun.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Btcfun *BtcfunFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *BtcfunPause) (event.Subscription, error) {

	logs, sub, err := _Btcfun.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcfunPause)
				if err := _Btcfun.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Btcfun *BtcfunFilterer) ParsePause(log types.Log) (*BtcfunPause, error) {
	event := new(BtcfunPause)
	if err := _Btcfun.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtcfunRefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the Btcfun contract.
type BtcfunRefundIterator struct {
	Event *BtcfunRefund // Event containing the contract specifics and raw log

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
func (it *BtcfunRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcfunRefund)
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
		it.Event = new(BtcfunRefund)
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
func (it *BtcfunRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcfunRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcfunRefund represents a Refund event raised by the Btcfun contract.
type BtcfunRefund struct {
	Token    common.Address
	Sender   common.Address
	Refunded *big.Int
	Currency common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0x4dc6fb46a7873cfa850a800f3408620fc2d8f32112069d21c1ac97720cc62344.
//
// Solidity: event Refund(address indexed token, address indexed sender, uint256 refunded, address indexed currency)
func (_Btcfun *BtcfunFilterer) FilterRefund(opts *bind.FilterOpts, token []common.Address, sender []common.Address, currency []common.Address) (*BtcfunRefundIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	logs, sub, err := _Btcfun.contract.FilterLogs(opts, "Refund", tokenRule, senderRule, currencyRule)
	if err != nil {
		return nil, err
	}
	return &BtcfunRefundIterator{contract: _Btcfun.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0x4dc6fb46a7873cfa850a800f3408620fc2d8f32112069d21c1ac97720cc62344.
//
// Solidity: event Refund(address indexed token, address indexed sender, uint256 refunded, address indexed currency)
func (_Btcfun *BtcfunFilterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *BtcfunRefund, token []common.Address, sender []common.Address, currency []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}

	logs, sub, err := _Btcfun.contract.WatchLogs(opts, "Refund", tokenRule, senderRule, currencyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcfunRefund)
				if err := _Btcfun.contract.UnpackLog(event, "Refund", log); err != nil {
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

// ParseRefund is a log parse operation binding the contract event 0x4dc6fb46a7873cfa850a800f3408620fc2d8f32112069d21c1ac97720cc62344.
//
// Solidity: event Refund(address indexed token, address indexed sender, uint256 refunded, address indexed currency)
func (_Btcfun *BtcfunFilterer) ParseRefund(log types.Log) (*BtcfunRefund, error) {
	event := new(BtcfunRefund)
	if err := _Btcfun.contract.UnpackLog(event, "Refund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BtcfunUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Btcfun contract.
type BtcfunUnpauseIterator struct {
	Event *BtcfunUnpause // Event containing the contract specifics and raw log

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
func (it *BtcfunUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BtcfunUnpause)
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
		it.Event = new(BtcfunUnpause)
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
func (it *BtcfunUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BtcfunUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BtcfunUnpause represents a Unpause event raised by the Btcfun contract.
type BtcfunUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Btcfun *BtcfunFilterer) FilterUnpause(opts *bind.FilterOpts) (*BtcfunUnpauseIterator, error) {

	logs, sub, err := _Btcfun.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &BtcfunUnpauseIterator{contract: _Btcfun.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Btcfun *BtcfunFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *BtcfunUnpause) (event.Subscription, error) {

	logs, sub, err := _Btcfun.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BtcfunUnpause)
				if err := _Btcfun.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Btcfun *BtcfunFilterer) ParseUnpause(log types.Log) (*BtcfunUnpause, error) {
	event := new(BtcfunUnpause)
	if err := _Btcfun.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
