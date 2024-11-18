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

// SwapSwapAmountParams is an auto generated low-level Go binding around an user-defined struct.
type SwapSwapAmountParams struct {
	Path        []byte
	Recipient   common.Address
	Amount      *big.Int
	MinAcquired *big.Int
	Deadline    *big.Int
}

// SwapSwapDesireParams is an auto generated low-level Go binding around an user-defined struct.
type SwapSwapDesireParams struct {
	Path      []byte
	Recipient common.Address
	Desire    *big.Int
	MaxPayed  *big.Int
	Deadline  *big.Int
}

// SwapSwapParams is an auto generated low-level Go binding around an user-defined struct.
type SwapSwapParams struct {
	TokenX      common.Address
	TokenY      common.Address
	Fee         *big.Int
	BoundaryPt  *big.Int
	Recipient   common.Address
	Amount      *big.Int
	MaxPayed    *big.Int
	MinAcquired *big.Int
	Deadline    *big.Int
}

// RouterMetaData contains all meta data concerning the Router contract.
var RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"minAcquired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structSwap.SwapAmountParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"swapAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acquire\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"desire\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxPayed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structSwap.SwapDesireParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"swapDesire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"acquire\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"boundaryPt\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxPayed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAcquired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structSwap.SwapParams\",\"name\":\"swapParams\",\"type\":\"tuple\"}],\"name\":\"swapX2Y\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swapX2YCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"boundaryPt\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxPayed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAcquired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structSwap.SwapParams\",\"name\":\"swapParams\",\"type\":\"tuple\"}],\"name\":\"swapX2YDesireY\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"boundaryPt\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxPayed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAcquired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structSwap.SwapParams\",\"name\":\"swapParams\",\"type\":\"tuple\"}],\"name\":\"swapY2X\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swapY2XCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"boundaryPt\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"maxPayed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAcquired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structSwap.SwapParams\",\"name\":\"swapParams\",\"type\":\"tuple\"}],\"name\":\"swapY2XDesireX\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterMetaData.ABI instead.
var RouterABI = RouterMetaData.ABI

// Router is an auto generated Go binding around an Ethereum contract.
type Router struct {
	RouterCaller     // Read-only binding to the contract
	RouterTransactor // Write-only binding to the contract
	RouterFilterer   // Log filterer for contract events
}

// RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterSession struct {
	Contract     *Router           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterCallerSession struct {
	Contract *RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterTransactorSession struct {
	Contract     *RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterRaw struct {
	Contract *Router // Generic contract binding to access the raw methods on
}

// RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterCallerRaw struct {
	Contract *RouterCaller // Generic read-only contract binding to access the raw methods on
}

// RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterTransactorRaw struct {
	Contract *RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter creates a new instance of Router, bound to a specific deployed contract.
func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// NewRouterCaller creates a new read-only instance of Router, bound to a specific deployed contract.
func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

// NewRouterTransactor creates a new write-only instance of Router, bound to a specific deployed contract.
func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

// NewRouterFilterer creates a new log filterer instance of Router, bound to a specific deployed contract.
func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

// bindRouter binds a generic wrapper to an already deployed contract.
func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Router *RouterCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Router *RouterSession) WETH9() (common.Address, error) {
	return _Router.Contract.WETH9(&_Router.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Router *RouterCallerSession) WETH9() (common.Address, error) {
	return _Router.Contract.WETH9(&_Router.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Router *RouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Router *RouterSession) Factory() (common.Address, error) {
	return _Router.Contract.Factory(&_Router.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Router *RouterCallerSession) Factory() (common.Address, error) {
	return _Router.Contract.Factory(&_Router.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0xbecbcc6a.
//
// Solidity: function pool(address tokenX, address tokenY, uint24 fee) view returns(address)
func (_Router *RouterCaller) Pool(opts *bind.CallOpts, tokenX common.Address, tokenY common.Address, fee *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "pool", tokenX, tokenY, fee)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0xbecbcc6a.
//
// Solidity: function pool(address tokenX, address tokenY, uint24 fee) view returns(address)
func (_Router *RouterSession) Pool(tokenX common.Address, tokenY common.Address, fee *big.Int) (common.Address, error) {
	return _Router.Contract.Pool(&_Router.CallOpts, tokenX, tokenY, fee)
}

// Pool is a free data retrieval call binding the contract method 0xbecbcc6a.
//
// Solidity: function pool(address tokenX, address tokenY, uint24 fee) view returns(address)
func (_Router *RouterCallerSession) Pool(tokenX common.Address, tokenY common.Address, fee *big.Int) (common.Address, error) {
	return _Router.Contract.Pool(&_Router.CallOpts, tokenX, tokenY, fee)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Router *RouterTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Router *RouterSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Router.Contract.Multicall(&_Router.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Router *RouterTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Router.Contract.Multicall(&_Router.TransactOpts, data)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_Router *RouterTransactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_Router *RouterSession) RefundETH() (*types.Transaction, error) {
	return _Router.Contract.RefundETH(&_Router.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_Router *RouterTransactorSession) RefundETH() (*types.Transaction, error) {
	return _Router.Contract.RefundETH(&_Router.TransactOpts)
}

// SwapAmount is a paid mutator transaction binding the contract method 0x75ceafe6.
//
// Solidity: function swapAmount((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_Router *RouterTransactor) SwapAmount(opts *bind.TransactOpts, params SwapSwapAmountParams) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapAmount", params)
}

// SwapAmount is a paid mutator transaction binding the contract method 0x75ceafe6.
//
// Solidity: function swapAmount((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_Router *RouterSession) SwapAmount(params SwapSwapAmountParams) (*types.Transaction, error) {
	return _Router.Contract.SwapAmount(&_Router.TransactOpts, params)
}

// SwapAmount is a paid mutator transaction binding the contract method 0x75ceafe6.
//
// Solidity: function swapAmount((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_Router *RouterTransactorSession) SwapAmount(params SwapSwapAmountParams) (*types.Transaction, error) {
	return _Router.Contract.SwapAmount(&_Router.TransactOpts, params)
}

// SwapDesire is a paid mutator transaction binding the contract method 0x115ff67e.
//
// Solidity: function swapDesire((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_Router *RouterTransactor) SwapDesire(opts *bind.TransactOpts, params SwapSwapDesireParams) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapDesire", params)
}

// SwapDesire is a paid mutator transaction binding the contract method 0x115ff67e.
//
// Solidity: function swapDesire((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_Router *RouterSession) SwapDesire(params SwapSwapDesireParams) (*types.Transaction, error) {
	return _Router.Contract.SwapDesire(&_Router.TransactOpts, params)
}

// SwapDesire is a paid mutator transaction binding the contract method 0x115ff67e.
//
// Solidity: function swapDesire((bytes,address,uint128,uint256,uint256) params) payable returns(uint256 cost, uint256 acquire)
func (_Router *RouterTransactorSession) SwapDesire(params SwapSwapDesireParams) (*types.Transaction, error) {
	return _Router.Contract.SwapDesire(&_Router.TransactOpts, params)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x46edd9c8.
//
// Solidity: function swapX2Y((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Router *RouterTransactor) SwapX2Y(opts *bind.TransactOpts, swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapX2Y", swapParams)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x46edd9c8.
//
// Solidity: function swapX2Y((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Router *RouterSession) SwapX2Y(swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Router.Contract.SwapX2Y(&_Router.TransactOpts, swapParams)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x46edd9c8.
//
// Solidity: function swapX2Y((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Router *RouterTransactorSession) SwapX2Y(swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Router.Contract.SwapX2Y(&_Router.TransactOpts, swapParams)
}

// SwapX2YCallback is a paid mutator transaction binding the contract method 0x18780684.
//
// Solidity: function swapX2YCallback(uint256 x, uint256 y, bytes data) returns()
func (_Router *RouterTransactor) SwapX2YCallback(opts *bind.TransactOpts, x *big.Int, y *big.Int, data []byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapX2YCallback", x, y, data)
}

// SwapX2YCallback is a paid mutator transaction binding the contract method 0x18780684.
//
// Solidity: function swapX2YCallback(uint256 x, uint256 y, bytes data) returns()
func (_Router *RouterSession) SwapX2YCallback(x *big.Int, y *big.Int, data []byte) (*types.Transaction, error) {
	return _Router.Contract.SwapX2YCallback(&_Router.TransactOpts, x, y, data)
}

// SwapX2YCallback is a paid mutator transaction binding the contract method 0x18780684.
//
// Solidity: function swapX2YCallback(uint256 x, uint256 y, bytes data) returns()
func (_Router *RouterTransactorSession) SwapX2YCallback(x *big.Int, y *big.Int, data []byte) (*types.Transaction, error) {
	return _Router.Contract.SwapX2YCallback(&_Router.TransactOpts, x, y, data)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0xf3da61a9.
//
// Solidity: function swapX2YDesireY((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Router *RouterTransactor) SwapX2YDesireY(opts *bind.TransactOpts, swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapX2YDesireY", swapParams)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0xf3da61a9.
//
// Solidity: function swapX2YDesireY((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Router *RouterSession) SwapX2YDesireY(swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Router.Contract.SwapX2YDesireY(&_Router.TransactOpts, swapParams)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0xf3da61a9.
//
// Solidity: function swapX2YDesireY((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Router *RouterTransactorSession) SwapX2YDesireY(swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Router.Contract.SwapX2YDesireY(&_Router.TransactOpts, swapParams)
}

// SwapY2X is a paid mutator transaction binding the contract method 0x247ec02c.
//
// Solidity: function swapY2X((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Router *RouterTransactor) SwapY2X(opts *bind.TransactOpts, swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapY2X", swapParams)
}

// SwapY2X is a paid mutator transaction binding the contract method 0x247ec02c.
//
// Solidity: function swapY2X((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Router *RouterSession) SwapY2X(swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Router.Contract.SwapY2X(&_Router.TransactOpts, swapParams)
}

// SwapY2X is a paid mutator transaction binding the contract method 0x247ec02c.
//
// Solidity: function swapY2X((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Router *RouterTransactorSession) SwapY2X(swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Router.Contract.SwapY2X(&_Router.TransactOpts, swapParams)
}

// SwapY2XCallback is a paid mutator transaction binding the contract method 0xd3e1c284.
//
// Solidity: function swapY2XCallback(uint256 x, uint256 y, bytes data) returns()
func (_Router *RouterTransactor) SwapY2XCallback(opts *bind.TransactOpts, x *big.Int, y *big.Int, data []byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapY2XCallback", x, y, data)
}

// SwapY2XCallback is a paid mutator transaction binding the contract method 0xd3e1c284.
//
// Solidity: function swapY2XCallback(uint256 x, uint256 y, bytes data) returns()
func (_Router *RouterSession) SwapY2XCallback(x *big.Int, y *big.Int, data []byte) (*types.Transaction, error) {
	return _Router.Contract.SwapY2XCallback(&_Router.TransactOpts, x, y, data)
}

// SwapY2XCallback is a paid mutator transaction binding the contract method 0xd3e1c284.
//
// Solidity: function swapY2XCallback(uint256 x, uint256 y, bytes data) returns()
func (_Router *RouterTransactorSession) SwapY2XCallback(x *big.Int, y *big.Int, data []byte) (*types.Transaction, error) {
	return _Router.Contract.SwapY2XCallback(&_Router.TransactOpts, x, y, data)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0x826377f6.
//
// Solidity: function swapY2XDesireX((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Router *RouterTransactor) SwapY2XDesireX(opts *bind.TransactOpts, swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapY2XDesireX", swapParams)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0x826377f6.
//
// Solidity: function swapY2XDesireX((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Router *RouterSession) SwapY2XDesireX(swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Router.Contract.SwapY2XDesireX(&_Router.TransactOpts, swapParams)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0x826377f6.
//
// Solidity: function swapY2XDesireX((address,address,uint24,int24,address,uint128,uint256,uint256,uint256) swapParams) payable returns()
func (_Router *RouterTransactorSession) SwapY2XDesireX(swapParams SwapSwapParams) (*types.Transaction, error) {
	return _Router.Contract.SwapY2XDesireX(&_Router.TransactOpts, swapParams)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 minAmount, address recipient) payable returns()
func (_Router *RouterTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "sweepToken", token, minAmount, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 minAmount, address recipient) payable returns()
func (_Router *RouterSession) SweepToken(token common.Address, minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Router.Contract.SweepToken(&_Router.TransactOpts, token, minAmount, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 minAmount, address recipient) payable returns()
func (_Router *RouterTransactorSession) SweepToken(token common.Address, minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Router.Contract.SweepToken(&_Router.TransactOpts, token, minAmount, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 minAmount, address recipient) payable returns()
func (_Router *RouterTransactor) UnwrapWETH9(opts *bind.TransactOpts, minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "unwrapWETH9", minAmount, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 minAmount, address recipient) payable returns()
func (_Router *RouterSession) UnwrapWETH9(minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Router.Contract.UnwrapWETH9(&_Router.TransactOpts, minAmount, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 minAmount, address recipient) payable returns()
func (_Router *RouterTransactorSession) UnwrapWETH9(minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Router.Contract.UnwrapWETH9(&_Router.TransactOpts, minAmount, recipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterSession) Receive() (*types.Transaction, error) {
	return _Router.Contract.Receive(&_Router.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterTransactorSession) Receive() (*types.Transaction, error) {
	return _Router.Contract.Receive(&_Router.TransactOpts)
}
