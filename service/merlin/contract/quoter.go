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

// QuoterMetaData contains all meta data concerning the Quoter contract.
var QuoterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicallNoRevert\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"successes\",\"type\":\"bool[]\"},{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"}],\"name\":\"swapAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"acquire\",\"type\":\"uint256\"},{\"internalType\":\"int24[]\",\"name\":\"pointAfterList\",\"type\":\"int24[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"desire\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"}],\"name\":\"swapDesire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"internalType\":\"int24[]\",\"name\":\"pointAfterList\",\"type\":\"int24[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"lowPt\",\"type\":\"int24\"}],\"name\":\"swapX2Y\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"finalPoint\",\"type\":\"int24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"}],\"name\":\"swapX2YCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint128\",\"name\":\"desireY\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"lowPt\",\"type\":\"int24\"}],\"name\":\"swapX2YDesireY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"finalPoint\",\"type\":\"int24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"highPt\",\"type\":\"int24\"}],\"name\":\"swapY2X\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountX\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"finalPoint\",\"type\":\"int24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"}],\"name\":\"swapY2XCallback\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint128\",\"name\":\"desireX\",\"type\":\"uint128\"},{\"internalType\":\"int24\",\"name\":\"highPt\",\"type\":\"int24\"}],\"name\":\"swapY2XDesireX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountY\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"finalPoint\",\"type\":\"int24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// QuoterABI is the input ABI used to generate the binding from.
// Deprecated: Use QuoterMetaData.ABI instead.
var QuoterABI = QuoterMetaData.ABI

// Quoter is an auto generated Go binding around an Ethereum contract.
type Quoter struct {
	QuoterCaller     // Read-only binding to the contract
	QuoterTransactor // Write-only binding to the contract
	QuoterFilterer   // Log filterer for contract events
}

// QuoterCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuoterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuoterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuoterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuoterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuoterSession struct {
	Contract     *Quoter           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuoterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuoterCallerSession struct {
	Contract *QuoterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// QuoterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuoterTransactorSession struct {
	Contract     *QuoterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuoterRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuoterRaw struct {
	Contract *Quoter // Generic contract binding to access the raw methods on
}

// QuoterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuoterCallerRaw struct {
	Contract *QuoterCaller // Generic read-only contract binding to access the raw methods on
}

// QuoterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuoterTransactorRaw struct {
	Contract *QuoterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuoter creates a new instance of Quoter, bound to a specific deployed contract.
func NewQuoter(address common.Address, backend bind.ContractBackend) (*Quoter, error) {
	contract, err := bindQuoter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Quoter{QuoterCaller: QuoterCaller{contract: contract}, QuoterTransactor: QuoterTransactor{contract: contract}, QuoterFilterer: QuoterFilterer{contract: contract}}, nil
}

// NewQuoterCaller creates a new read-only instance of Quoter, bound to a specific deployed contract.
func NewQuoterCaller(address common.Address, caller bind.ContractCaller) (*QuoterCaller, error) {
	contract, err := bindQuoter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuoterCaller{contract: contract}, nil
}

// NewQuoterTransactor creates a new write-only instance of Quoter, bound to a specific deployed contract.
func NewQuoterTransactor(address common.Address, transactor bind.ContractTransactor) (*QuoterTransactor, error) {
	contract, err := bindQuoter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuoterTransactor{contract: contract}, nil
}

// NewQuoterFilterer creates a new log filterer instance of Quoter, bound to a specific deployed contract.
func NewQuoterFilterer(address common.Address, filterer bind.ContractFilterer) (*QuoterFilterer, error) {
	contract, err := bindQuoter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuoterFilterer{contract: contract}, nil
}

// bindQuoter binds a generic wrapper to an already deployed contract.
func bindQuoter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QuoterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quoter *QuoterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Quoter.Contract.QuoterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quoter *QuoterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quoter.Contract.QuoterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quoter *QuoterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quoter.Contract.QuoterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quoter *QuoterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Quoter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quoter *QuoterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quoter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quoter *QuoterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quoter.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Quoter *QuoterCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Quoter *QuoterSession) WETH9() (common.Address, error) {
	return _Quoter.Contract.WETH9(&_Quoter.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Quoter *QuoterCallerSession) WETH9() (common.Address, error) {
	return _Quoter.Contract.WETH9(&_Quoter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Quoter *QuoterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Quoter *QuoterSession) Factory() (common.Address, error) {
	return _Quoter.Contract.Factory(&_Quoter.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Quoter *QuoterCallerSession) Factory() (common.Address, error) {
	return _Quoter.Contract.Factory(&_Quoter.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0xbecbcc6a.
//
// Solidity: function pool(address tokenX, address tokenY, uint24 fee) view returns(address)
func (_Quoter *QuoterCaller) Pool(opts *bind.CallOpts, tokenX common.Address, tokenY common.Address, fee *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "pool", tokenX, tokenY, fee)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0xbecbcc6a.
//
// Solidity: function pool(address tokenX, address tokenY, uint24 fee) view returns(address)
func (_Quoter *QuoterSession) Pool(tokenX common.Address, tokenY common.Address, fee *big.Int) (common.Address, error) {
	return _Quoter.Contract.Pool(&_Quoter.CallOpts, tokenX, tokenY, fee)
}

// Pool is a free data retrieval call binding the contract method 0xbecbcc6a.
//
// Solidity: function pool(address tokenX, address tokenY, uint24 fee) view returns(address)
func (_Quoter *QuoterCallerSession) Pool(tokenX common.Address, tokenY common.Address, fee *big.Int) (common.Address, error) {
	return _Quoter.Contract.Pool(&_Quoter.CallOpts, tokenX, tokenY, fee)
}

// SwapX2YCallback is a free data retrieval call binding the contract method 0x18780684.
//
// Solidity: function swapX2YCallback(uint256 x, uint256 y, bytes path) view returns()
func (_Quoter *QuoterCaller) SwapX2YCallback(opts *bind.CallOpts, x *big.Int, y *big.Int, path []byte) error {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "swapX2YCallback", x, y, path)

	if err != nil {
		return err
	}

	return err

}

// SwapX2YCallback is a free data retrieval call binding the contract method 0x18780684.
//
// Solidity: function swapX2YCallback(uint256 x, uint256 y, bytes path) view returns()
func (_Quoter *QuoterSession) SwapX2YCallback(x *big.Int, y *big.Int, path []byte) error {
	return _Quoter.Contract.SwapX2YCallback(&_Quoter.CallOpts, x, y, path)
}

// SwapX2YCallback is a free data retrieval call binding the contract method 0x18780684.
//
// Solidity: function swapX2YCallback(uint256 x, uint256 y, bytes path) view returns()
func (_Quoter *QuoterCallerSession) SwapX2YCallback(x *big.Int, y *big.Int, path []byte) error {
	return _Quoter.Contract.SwapX2YCallback(&_Quoter.CallOpts, x, y, path)
}

// SwapY2XCallback is a free data retrieval call binding the contract method 0xd3e1c284.
//
// Solidity: function swapY2XCallback(uint256 x, uint256 y, bytes path) view returns()
func (_Quoter *QuoterCaller) SwapY2XCallback(opts *bind.CallOpts, x *big.Int, y *big.Int, path []byte) error {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "swapY2XCallback", x, y, path)

	if err != nil {
		return err
	}

	return err

}

func (_Quoter *QuoterCaller) SwapX2YStatic(opts *bind.CallOpts, tokenX common.Address, tokenY common.Address, fee *big.Int, amount *big.Int, lowPt *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Quoter.contract.Call(opts, &out, "swapX2Y", tokenX, tokenY, fee, amount, lowPt)

	if err != nil {
		return nil, nil, err
	}
	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	return out0, out1, err
}

// SwapY2XCallback is a free data retrieval call binding the contract method 0xd3e1c284.
//
// Solidity: function swapY2XCallback(uint256 x, uint256 y, bytes path) view returns()
func (_Quoter *QuoterSession) SwapY2XCallback(x *big.Int, y *big.Int, path []byte) error {
	return _Quoter.Contract.SwapY2XCallback(&_Quoter.CallOpts, x, y, path)
}

// SwapY2XCallback is a free data retrieval call binding the contract method 0xd3e1c284.
//
// Solidity: function swapY2XCallback(uint256 x, uint256 y, bytes path) view returns()
func (_Quoter *QuoterCallerSession) SwapY2XCallback(x *big.Int, y *big.Int, path []byte) error {
	return _Quoter.Contract.SwapY2XCallback(&_Quoter.CallOpts, x, y, path)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Quoter *QuoterTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Quoter *QuoterSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Quoter.Contract.Multicall(&_Quoter.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Quoter *QuoterTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Quoter.Contract.Multicall(&_Quoter.TransactOpts, data)
}

// MulticallNoRevert is a paid mutator transaction binding the contract method 0xda61b46a.
//
// Solidity: function multicallNoRevert(bytes[] data) payable returns(bool[] successes, bytes[] results)
func (_Quoter *QuoterTransactor) MulticallNoRevert(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "multicallNoRevert", data)
}

// MulticallNoRevert is a paid mutator transaction binding the contract method 0xda61b46a.
//
// Solidity: function multicallNoRevert(bytes[] data) payable returns(bool[] successes, bytes[] results)
func (_Quoter *QuoterSession) MulticallNoRevert(data [][]byte) (*types.Transaction, error) {
	return _Quoter.Contract.MulticallNoRevert(&_Quoter.TransactOpts, data)
}

// MulticallNoRevert is a paid mutator transaction binding the contract method 0xda61b46a.
//
// Solidity: function multicallNoRevert(bytes[] data) payable returns(bool[] successes, bytes[] results)
func (_Quoter *QuoterTransactorSession) MulticallNoRevert(data [][]byte) (*types.Transaction, error) {
	return _Quoter.Contract.MulticallNoRevert(&_Quoter.TransactOpts, data)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_Quoter *QuoterTransactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_Quoter *QuoterSession) RefundETH() (*types.Transaction, error) {
	return _Quoter.Contract.RefundETH(&_Quoter.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_Quoter *QuoterTransactorSession) RefundETH() (*types.Transaction, error) {
	return _Quoter.Contract.RefundETH(&_Quoter.TransactOpts)
}

// SwapAmount is a paid mutator transaction binding the contract method 0x0980929e.
//
// Solidity: function swapAmount(uint128 amount, bytes path) returns(uint256 acquire, int24[] pointAfterList)
func (_Quoter *QuoterTransactor) SwapAmount(opts *bind.TransactOpts, amount *big.Int, path []byte) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "swapAmount", amount, path)
}

// SwapAmount is a paid mutator transaction binding the contract method 0x0980929e.
//
// Solidity: function swapAmount(uint128 amount, bytes path) returns(uint256 acquire, int24[] pointAfterList)
func (_Quoter *QuoterSession) SwapAmount(amount *big.Int, path []byte) (*types.Transaction, error) {
	return _Quoter.Contract.SwapAmount(&_Quoter.TransactOpts, amount, path)
}

// SwapAmount is a paid mutator transaction binding the contract method 0x0980929e.
//
// Solidity: function swapAmount(uint128 amount, bytes path) returns(uint256 acquire, int24[] pointAfterList)
func (_Quoter *QuoterTransactorSession) SwapAmount(amount *big.Int, path []byte) (*types.Transaction, error) {
	return _Quoter.Contract.SwapAmount(&_Quoter.TransactOpts, amount, path)
}

// SwapDesire is a paid mutator transaction binding the contract method 0x18ce0610.
//
// Solidity: function swapDesire(uint128 desire, bytes path) returns(uint256 cost, int24[] pointAfterList)
func (_Quoter *QuoterTransactor) SwapDesire(opts *bind.TransactOpts, desire *big.Int, path []byte) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "swapDesire", desire, path)
}

// SwapDesire is a paid mutator transaction binding the contract method 0x18ce0610.
//
// Solidity: function swapDesire(uint128 desire, bytes path) returns(uint256 cost, int24[] pointAfterList)
func (_Quoter *QuoterSession) SwapDesire(desire *big.Int, path []byte) (*types.Transaction, error) {
	return _Quoter.Contract.SwapDesire(&_Quoter.TransactOpts, desire, path)
}

// SwapDesire is a paid mutator transaction binding the contract method 0x18ce0610.
//
// Solidity: function swapDesire(uint128 desire, bytes path) returns(uint256 cost, int24[] pointAfterList)
func (_Quoter *QuoterTransactorSession) SwapDesire(desire *big.Int, path []byte) (*types.Transaction, error) {
	return _Quoter.Contract.SwapDesire(&_Quoter.TransactOpts, desire, path)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x8501721f.
//
// Solidity: function swapX2Y(address tokenX, address tokenY, uint24 fee, uint128 amount, int24 lowPt) returns(uint256 amountY, int24 finalPoint)
func (_Quoter *QuoterTransactor) SwapX2Y(opts *bind.TransactOpts, tokenX common.Address, tokenY common.Address, fee *big.Int, amount *big.Int, lowPt *big.Int) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "swapX2Y", tokenX, tokenY, fee, amount, lowPt)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x8501721f.
//
// Solidity: function swapX2Y(address tokenX, address tokenY, uint24 fee, uint128 amount, int24 lowPt) returns(uint256 amountY, int24 finalPoint)
func (_Quoter *QuoterSession) SwapX2Y(tokenX common.Address, tokenY common.Address, fee *big.Int, amount *big.Int, lowPt *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.SwapX2Y(&_Quoter.TransactOpts, tokenX, tokenY, fee, amount, lowPt)
}

func (_Quoter *QuoterSession) SwapX2YStatic(tokenX common.Address, tokenY common.Address, fee *big.Int, amount *big.Int, lowPt *big.Int) (*big.Int, *big.Int, error) {
	return _Quoter.Contract.SwapX2YStatic(&_Quoter.CallOpts, tokenX, tokenY, fee, amount, lowPt)
}

// SwapX2Y is a paid mutator transaction binding the contract method 0x8501721f.
//
// Solidity: function swapX2Y(address tokenX, address tokenY, uint24 fee, uint128 amount, int24 lowPt) returns(uint256 amountY, int24 finalPoint)
func (_Quoter *QuoterTransactorSession) SwapX2Y(tokenX common.Address, tokenY common.Address, fee *big.Int, amount *big.Int, lowPt *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.SwapX2Y(&_Quoter.TransactOpts, tokenX, tokenY, fee, amount, lowPt)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0xcdca428a.
//
// Solidity: function swapX2YDesireY(address tokenX, address tokenY, uint24 fee, uint128 desireY, int24 lowPt) returns(uint256 amountX, int24 finalPoint)
func (_Quoter *QuoterTransactor) SwapX2YDesireY(opts *bind.TransactOpts, tokenX common.Address, tokenY common.Address, fee *big.Int, desireY *big.Int, lowPt *big.Int) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "swapX2YDesireY", tokenX, tokenY, fee, desireY, lowPt)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0xcdca428a.
//
// Solidity: function swapX2YDesireY(address tokenX, address tokenY, uint24 fee, uint128 desireY, int24 lowPt) returns(uint256 amountX, int24 finalPoint)
func (_Quoter *QuoterSession) SwapX2YDesireY(tokenX common.Address, tokenY common.Address, fee *big.Int, desireY *big.Int, lowPt *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.SwapX2YDesireY(&_Quoter.TransactOpts, tokenX, tokenY, fee, desireY, lowPt)
}

// SwapX2YDesireY is a paid mutator transaction binding the contract method 0xcdca428a.
//
// Solidity: function swapX2YDesireY(address tokenX, address tokenY, uint24 fee, uint128 desireY, int24 lowPt) returns(uint256 amountX, int24 finalPoint)
func (_Quoter *QuoterTransactorSession) SwapX2YDesireY(tokenX common.Address, tokenY common.Address, fee *big.Int, desireY *big.Int, lowPt *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.SwapX2YDesireY(&_Quoter.TransactOpts, tokenX, tokenY, fee, desireY, lowPt)
}

// SwapY2X is a paid mutator transaction binding the contract method 0xc2ce9118.
//
// Solidity: function swapY2X(address tokenX, address tokenY, uint24 fee, uint128 amount, int24 highPt) returns(uint256 amountX, int24 finalPoint)
func (_Quoter *QuoterTransactor) SwapY2X(opts *bind.TransactOpts, tokenX common.Address, tokenY common.Address, fee *big.Int, amount *big.Int, highPt *big.Int) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "swapY2X", tokenX, tokenY, fee, amount, highPt)
}

// SwapY2X is a paid mutator transaction binding the contract method 0xc2ce9118.
//
// Solidity: function swapY2X(address tokenX, address tokenY, uint24 fee, uint128 amount, int24 highPt) returns(uint256 amountX, int24 finalPoint)
func (_Quoter *QuoterSession) SwapY2X(tokenX common.Address, tokenY common.Address, fee *big.Int, amount *big.Int, highPt *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.SwapY2X(&_Quoter.TransactOpts, tokenX, tokenY, fee, amount, highPt)
}

// SwapY2X is a paid mutator transaction binding the contract method 0xc2ce9118.
//
// Solidity: function swapY2X(address tokenX, address tokenY, uint24 fee, uint128 amount, int24 highPt) returns(uint256 amountX, int24 finalPoint)
func (_Quoter *QuoterTransactorSession) SwapY2X(tokenX common.Address, tokenY common.Address, fee *big.Int, amount *big.Int, highPt *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.SwapY2X(&_Quoter.TransactOpts, tokenX, tokenY, fee, amount, highPt)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0x0946629b.
//
// Solidity: function swapY2XDesireX(address tokenX, address tokenY, uint24 fee, uint128 desireX, int24 highPt) returns(uint256 amountY, int24 finalPoint)
func (_Quoter *QuoterTransactor) SwapY2XDesireX(opts *bind.TransactOpts, tokenX common.Address, tokenY common.Address, fee *big.Int, desireX *big.Int, highPt *big.Int) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "swapY2XDesireX", tokenX, tokenY, fee, desireX, highPt)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0x0946629b.
//
// Solidity: function swapY2XDesireX(address tokenX, address tokenY, uint24 fee, uint128 desireX, int24 highPt) returns(uint256 amountY, int24 finalPoint)
func (_Quoter *QuoterSession) SwapY2XDesireX(tokenX common.Address, tokenY common.Address, fee *big.Int, desireX *big.Int, highPt *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.SwapY2XDesireX(&_Quoter.TransactOpts, tokenX, tokenY, fee, desireX, highPt)
}

// SwapY2XDesireX is a paid mutator transaction binding the contract method 0x0946629b.
//
// Solidity: function swapY2XDesireX(address tokenX, address tokenY, uint24 fee, uint128 desireX, int24 highPt) returns(uint256 amountY, int24 finalPoint)
func (_Quoter *QuoterTransactorSession) SwapY2XDesireX(tokenX common.Address, tokenY common.Address, fee *big.Int, desireX *big.Int, highPt *big.Int) (*types.Transaction, error) {
	return _Quoter.Contract.SwapY2XDesireX(&_Quoter.TransactOpts, tokenX, tokenY, fee, desireX, highPt)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 minAmount, address recipient) payable returns()
func (_Quoter *QuoterTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "sweepToken", token, minAmount, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 minAmount, address recipient) payable returns()
func (_Quoter *QuoterSession) SweepToken(token common.Address, minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Quoter.Contract.SweepToken(&_Quoter.TransactOpts, token, minAmount, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 minAmount, address recipient) payable returns()
func (_Quoter *QuoterTransactorSession) SweepToken(token common.Address, minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Quoter.Contract.SweepToken(&_Quoter.TransactOpts, token, minAmount, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 minAmount, address recipient) payable returns()
func (_Quoter *QuoterTransactor) UnwrapWETH9(opts *bind.TransactOpts, minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Quoter.contract.Transact(opts, "unwrapWETH9", minAmount, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 minAmount, address recipient) payable returns()
func (_Quoter *QuoterSession) UnwrapWETH9(minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Quoter.Contract.UnwrapWETH9(&_Quoter.TransactOpts, minAmount, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 minAmount, address recipient) payable returns()
func (_Quoter *QuoterTransactorSession) UnwrapWETH9(minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Quoter.Contract.UnwrapWETH9(&_Quoter.TransactOpts, minAmount, recipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Quoter *QuoterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quoter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Quoter *QuoterSession) Receive() (*types.Transaction, error) {
	return _Quoter.Contract.Receive(&_Quoter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Quoter *QuoterTransactorSession) Receive() (*types.Transaction, error) {
	return _Quoter.Contract.Receive(&_Quoter.TransactOpts)
}
