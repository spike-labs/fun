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

// BindBoxMetaData contains all meta data concerning the BindBox contract.
var BindBoxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_fundMgr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"NotWhitelisted\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"RolledOver\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"changeMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTokenId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flipSaleStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fromTokenId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundMgr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAllowedToClaim\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"saleActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"salePrice\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_salePrice\",\"type\":\"uint128\"}],\"name\":\"setSalePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_fromTokenId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_toTokenId\",\"type\":\"uint32\"}],\"name\":\"setTokenScope\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toTokenId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFundMgr\",\"type\":\"address\"}],\"name\":\"updateFundMgr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BindBoxABI is the input ABI used to generate the binding from.
// Deprecated: Use BindBoxMetaData.ABI instead.
var BindBoxABI = BindBoxMetaData.ABI

// BindBox is an auto generated Go binding around an Ethereum contract.
type BindBox struct {
	BindBoxCaller     // Read-only binding to the contract
	BindBoxTransactor // Write-only binding to the contract
	BindBoxFilterer   // Log filterer for contract events
}

// BindBoxCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindBoxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindBoxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindBoxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindBoxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindBoxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindBoxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindBoxSession struct {
	Contract     *BindBox          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindBoxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindBoxCallerSession struct {
	Contract *BindBoxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BindBoxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindBoxTransactorSession struct {
	Contract     *BindBoxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BindBoxRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindBoxRaw struct {
	Contract *BindBox // Generic contract binding to access the raw methods on
}

// BindBoxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindBoxCallerRaw struct {
	Contract *BindBoxCaller // Generic read-only contract binding to access the raw methods on
}

// BindBoxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindBoxTransactorRaw struct {
	Contract *BindBoxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindBox creates a new instance of BindBox, bound to a specific deployed contract.
func NewBindBox(address common.Address, backend bind.ContractBackend) (*BindBox, error) {
	contract, err := bindBindBox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BindBox{BindBoxCaller: BindBoxCaller{contract: contract}, BindBoxTransactor: BindBoxTransactor{contract: contract}, BindBoxFilterer: BindBoxFilterer{contract: contract}}, nil
}

// NewBindBoxCaller creates a new read-only instance of BindBox, bound to a specific deployed contract.
func NewBindBoxCaller(address common.Address, caller bind.ContractCaller) (*BindBoxCaller, error) {
	contract, err := bindBindBox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindBoxCaller{contract: contract}, nil
}

// NewBindBoxTransactor creates a new write-only instance of BindBox, bound to a specific deployed contract.
func NewBindBoxTransactor(address common.Address, transactor bind.ContractTransactor) (*BindBoxTransactor, error) {
	contract, err := bindBindBox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindBoxTransactor{contract: contract}, nil
}

// NewBindBoxFilterer creates a new log filterer instance of BindBox, bound to a specific deployed contract.
func NewBindBoxFilterer(address common.Address, filterer bind.ContractFilterer) (*BindBoxFilterer, error) {
	contract, err := bindBindBox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindBoxFilterer{contract: contract}, nil
}

// bindBindBox binds a generic wrapper to an already deployed contract.
func bindBindBox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BindBoxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindBox *BindBoxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindBox.Contract.BindBoxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindBox *BindBoxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindBox.Contract.BindBoxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindBox *BindBoxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindBox.Contract.BindBoxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindBox *BindBoxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindBox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindBox *BindBoxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindBox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindBox *BindBoxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindBox.Contract.contract.Transact(opts, method, params...)
}

// Claimed is a free data retrieval call binding the contract method 0xc884ef83.
//
// Solidity: function claimed(address ) view returns(uint256)
func (_BindBox *BindBoxCaller) Claimed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BindBox.contract.Call(opts, &out, "claimed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimed is a free data retrieval call binding the contract method 0xc884ef83.
//
// Solidity: function claimed(address ) view returns(uint256)
func (_BindBox *BindBoxSession) Claimed(arg0 common.Address) (*big.Int, error) {
	return _BindBox.Contract.Claimed(&_BindBox.CallOpts, arg0)
}

// Claimed is a free data retrieval call binding the contract method 0xc884ef83.
//
// Solidity: function claimed(address ) view returns(uint256)
func (_BindBox *BindBoxCallerSession) Claimed(arg0 common.Address) (*big.Int, error) {
	return _BindBox.Contract.Claimed(&_BindBox.CallOpts, arg0)
}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint32)
func (_BindBox *BindBoxCaller) CurrentTokenId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BindBox.contract.Call(opts, &out, "currentTokenId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint32)
func (_BindBox *BindBoxSession) CurrentTokenId() (uint32, error) {
	return _BindBox.Contract.CurrentTokenId(&_BindBox.CallOpts)
}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint32)
func (_BindBox *BindBoxCallerSession) CurrentTokenId() (uint32, error) {
	return _BindBox.Contract.CurrentTokenId(&_BindBox.CallOpts)
}

// FromTokenId is a free data retrieval call binding the contract method 0x112e0dee.
//
// Solidity: function fromTokenId() view returns(uint32)
func (_BindBox *BindBoxCaller) FromTokenId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BindBox.contract.Call(opts, &out, "fromTokenId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FromTokenId is a free data retrieval call binding the contract method 0x112e0dee.
//
// Solidity: function fromTokenId() view returns(uint32)
func (_BindBox *BindBoxSession) FromTokenId() (uint32, error) {
	return _BindBox.Contract.FromTokenId(&_BindBox.CallOpts)
}

// FromTokenId is a free data retrieval call binding the contract method 0x112e0dee.
//
// Solidity: function fromTokenId() view returns(uint32)
func (_BindBox *BindBoxCallerSession) FromTokenId() (uint32, error) {
	return _BindBox.Contract.FromTokenId(&_BindBox.CallOpts)
}

// FundMgr is a free data retrieval call binding the contract method 0x4b8539ea.
//
// Solidity: function fundMgr() view returns(address)
func (_BindBox *BindBoxCaller) FundMgr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindBox.contract.Call(opts, &out, "fundMgr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FundMgr is a free data retrieval call binding the contract method 0x4b8539ea.
//
// Solidity: function fundMgr() view returns(address)
func (_BindBox *BindBoxSession) FundMgr() (common.Address, error) {
	return _BindBox.Contract.FundMgr(&_BindBox.CallOpts)
}

// FundMgr is a free data retrieval call binding the contract method 0x4b8539ea.
//
// Solidity: function fundMgr() view returns(address)
func (_BindBox *BindBoxCallerSession) FundMgr() (common.Address, error) {
	return _BindBox.Contract.FundMgr(&_BindBox.CallOpts)
}

// MaxAllowedToClaim is a free data retrieval call binding the contract method 0xa56897f5.
//
// Solidity: function maxAllowedToClaim() view returns(uint32)
func (_BindBox *BindBoxCaller) MaxAllowedToClaim(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BindBox.contract.Call(opts, &out, "maxAllowedToClaim")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxAllowedToClaim is a free data retrieval call binding the contract method 0xa56897f5.
//
// Solidity: function maxAllowedToClaim() view returns(uint32)
func (_BindBox *BindBoxSession) MaxAllowedToClaim() (uint32, error) {
	return _BindBox.Contract.MaxAllowedToClaim(&_BindBox.CallOpts)
}

// MaxAllowedToClaim is a free data retrieval call binding the contract method 0xa56897f5.
//
// Solidity: function maxAllowedToClaim() view returns(uint32)
func (_BindBox *BindBoxCallerSession) MaxAllowedToClaim() (uint32, error) {
	return _BindBox.Contract.MaxAllowedToClaim(&_BindBox.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_BindBox *BindBoxCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BindBox.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_BindBox *BindBoxSession) MerkleRoot() ([32]byte, error) {
	return _BindBox.Contract.MerkleRoot(&_BindBox.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_BindBox *BindBoxCallerSession) MerkleRoot() ([32]byte, error) {
	return _BindBox.Contract.MerkleRoot(&_BindBox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BindBox *BindBoxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindBox.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BindBox *BindBoxSession) Owner() (common.Address, error) {
	return _BindBox.Contract.Owner(&_BindBox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BindBox *BindBoxCallerSession) Owner() (common.Address, error) {
	return _BindBox.Contract.Owner(&_BindBox.CallOpts)
}

// SaleActive is a free data retrieval call binding the contract method 0x68428a1b.
//
// Solidity: function saleActive() view returns(bool)
func (_BindBox *BindBoxCaller) SaleActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BindBox.contract.Call(opts, &out, "saleActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SaleActive is a free data retrieval call binding the contract method 0x68428a1b.
//
// Solidity: function saleActive() view returns(bool)
func (_BindBox *BindBoxSession) SaleActive() (bool, error) {
	return _BindBox.Contract.SaleActive(&_BindBox.CallOpts)
}

// SaleActive is a free data retrieval call binding the contract method 0x68428a1b.
//
// Solidity: function saleActive() view returns(bool)
func (_BindBox *BindBoxCallerSession) SaleActive() (bool, error) {
	return _BindBox.Contract.SaleActive(&_BindBox.CallOpts)
}

// SalePrice is a free data retrieval call binding the contract method 0xf51f96dd.
//
// Solidity: function salePrice() view returns(uint128)
func (_BindBox *BindBoxCaller) SalePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BindBox.contract.Call(opts, &out, "salePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SalePrice is a free data retrieval call binding the contract method 0xf51f96dd.
//
// Solidity: function salePrice() view returns(uint128)
func (_BindBox *BindBoxSession) SalePrice() (*big.Int, error) {
	return _BindBox.Contract.SalePrice(&_BindBox.CallOpts)
}

// SalePrice is a free data retrieval call binding the contract method 0xf51f96dd.
//
// Solidity: function salePrice() view returns(uint128)
func (_BindBox *BindBoxCallerSession) SalePrice() (*big.Int, error) {
	return _BindBox.Contract.SalePrice(&_BindBox.CallOpts)
}

// ToTokenId is a free data retrieval call binding the contract method 0x980f3700.
//
// Solidity: function toTokenId() view returns(uint32)
func (_BindBox *BindBoxCaller) ToTokenId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BindBox.contract.Call(opts, &out, "toTokenId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ToTokenId is a free data retrieval call binding the contract method 0x980f3700.
//
// Solidity: function toTokenId() view returns(uint32)
func (_BindBox *BindBoxSession) ToTokenId() (uint32, error) {
	return _BindBox.Contract.ToTokenId(&_BindBox.CallOpts)
}

// ToTokenId is a free data retrieval call binding the contract method 0x980f3700.
//
// Solidity: function toTokenId() view returns(uint32)
func (_BindBox *BindBoxCallerSession) ToTokenId() (uint32, error) {
	return _BindBox.Contract.ToTokenId(&_BindBox.CallOpts)
}

// VerifyProof is a free data retrieval call binding the contract method 0xe3486434.
//
// Solidity: function verifyProof(address user, bytes32[] proof) view returns(bool)
func (_BindBox *BindBoxCaller) VerifyProof(opts *bind.CallOpts, user common.Address, proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _BindBox.contract.Call(opts, &out, "verifyProof", user, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0xe3486434.
//
// Solidity: function verifyProof(address user, bytes32[] proof) view returns(bool)
func (_BindBox *BindBoxSession) VerifyProof(user common.Address, proof [][32]byte) (bool, error) {
	return _BindBox.Contract.VerifyProof(&_BindBox.CallOpts, user, proof)
}

// VerifyProof is a free data retrieval call binding the contract method 0xe3486434.
//
// Solidity: function verifyProof(address user, bytes32[] proof) view returns(bool)
func (_BindBox *BindBoxCallerSession) VerifyProof(user common.Address, proof [][32]byte) (bool, error) {
	return _BindBox.Contract.VerifyProof(&_BindBox.CallOpts, user, proof)
}

// ChangeMerkleRoot is a paid mutator transaction binding the contract method 0xebcea3db.
//
// Solidity: function changeMerkleRoot(bytes32 _merkleRoot) returns()
func (_BindBox *BindBoxTransactor) ChangeMerkleRoot(opts *bind.TransactOpts, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _BindBox.contract.Transact(opts, "changeMerkleRoot", _merkleRoot)
}

// ChangeMerkleRoot is a paid mutator transaction binding the contract method 0xebcea3db.
//
// Solidity: function changeMerkleRoot(bytes32 _merkleRoot) returns()
func (_BindBox *BindBoxSession) ChangeMerkleRoot(_merkleRoot [32]byte) (*types.Transaction, error) {
	return _BindBox.Contract.ChangeMerkleRoot(&_BindBox.TransactOpts, _merkleRoot)
}

// ChangeMerkleRoot is a paid mutator transaction binding the contract method 0xebcea3db.
//
// Solidity: function changeMerkleRoot(bytes32 _merkleRoot) returns()
func (_BindBox *BindBoxTransactorSession) ChangeMerkleRoot(_merkleRoot [32]byte) (*types.Transaction, error) {
	return _BindBox.Contract.ChangeMerkleRoot(&_BindBox.TransactOpts, _merkleRoot)
}

// Claim is a paid mutator transaction binding the contract method 0xd7aada81.
//
// Solidity: function claim(address token, bytes32[] proof) payable returns()
func (_BindBox *BindBoxTransactor) Claim(opts *bind.TransactOpts, token common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _BindBox.contract.Transact(opts, "claim", token, proof)
}

// Claim is a paid mutator transaction binding the contract method 0xd7aada81.
//
// Solidity: function claim(address token, bytes32[] proof) payable returns()
func (_BindBox *BindBoxSession) Claim(token common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _BindBox.Contract.Claim(&_BindBox.TransactOpts, token, proof)
}

// Claim is a paid mutator transaction binding the contract method 0xd7aada81.
//
// Solidity: function claim(address token, bytes32[] proof) payable returns()
func (_BindBox *BindBoxTransactorSession) Claim(token common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _BindBox.Contract.Claim(&_BindBox.TransactOpts, token, proof)
}

// FlipSaleStatus is a paid mutator transaction binding the contract method 0xce03ec93.
//
// Solidity: function flipSaleStatus() returns()
func (_BindBox *BindBoxTransactor) FlipSaleStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindBox.contract.Transact(opts, "flipSaleStatus")
}

// FlipSaleStatus is a paid mutator transaction binding the contract method 0xce03ec93.
//
// Solidity: function flipSaleStatus() returns()
func (_BindBox *BindBoxSession) FlipSaleStatus() (*types.Transaction, error) {
	return _BindBox.Contract.FlipSaleStatus(&_BindBox.TransactOpts)
}

// FlipSaleStatus is a paid mutator transaction binding the contract method 0xce03ec93.
//
// Solidity: function flipSaleStatus() returns()
func (_BindBox *BindBoxTransactorSession) FlipSaleStatus() (*types.Transaction, error) {
	return _BindBox.Contract.FlipSaleStatus(&_BindBox.TransactOpts)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_BindBox *BindBoxTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _BindBox.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_BindBox *BindBoxSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _BindBox.Contract.OnERC721Received(&_BindBox.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_BindBox *BindBoxTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _BindBox.Contract.OnERC721Received(&_BindBox.TransactOpts, arg0, arg1, arg2, arg3)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BindBox *BindBoxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindBox.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BindBox *BindBoxSession) RenounceOwnership() (*types.Transaction, error) {
	return _BindBox.Contract.RenounceOwnership(&_BindBox.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BindBox *BindBoxTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BindBox.Contract.RenounceOwnership(&_BindBox.TransactOpts)
}

// SetSalePrice is a paid mutator transaction binding the contract method 0xf492538b.
//
// Solidity: function setSalePrice(uint128 _salePrice) returns()
func (_BindBox *BindBoxTransactor) SetSalePrice(opts *bind.TransactOpts, _salePrice *big.Int) (*types.Transaction, error) {
	return _BindBox.contract.Transact(opts, "setSalePrice", _salePrice)
}

// SetSalePrice is a paid mutator transaction binding the contract method 0xf492538b.
//
// Solidity: function setSalePrice(uint128 _salePrice) returns()
func (_BindBox *BindBoxSession) SetSalePrice(_salePrice *big.Int) (*types.Transaction, error) {
	return _BindBox.Contract.SetSalePrice(&_BindBox.TransactOpts, _salePrice)
}

// SetSalePrice is a paid mutator transaction binding the contract method 0xf492538b.
//
// Solidity: function setSalePrice(uint128 _salePrice) returns()
func (_BindBox *BindBoxTransactorSession) SetSalePrice(_salePrice *big.Int) (*types.Transaction, error) {
	return _BindBox.Contract.SetSalePrice(&_BindBox.TransactOpts, _salePrice)
}

// SetTokenScope is a paid mutator transaction binding the contract method 0x6663a375.
//
// Solidity: function setTokenScope(uint32 _fromTokenId, uint32 _toTokenId) returns()
func (_BindBox *BindBoxTransactor) SetTokenScope(opts *bind.TransactOpts, _fromTokenId uint32, _toTokenId uint32) (*types.Transaction, error) {
	return _BindBox.contract.Transact(opts, "setTokenScope", _fromTokenId, _toTokenId)
}

// SetTokenScope is a paid mutator transaction binding the contract method 0x6663a375.
//
// Solidity: function setTokenScope(uint32 _fromTokenId, uint32 _toTokenId) returns()
func (_BindBox *BindBoxSession) SetTokenScope(_fromTokenId uint32, _toTokenId uint32) (*types.Transaction, error) {
	return _BindBox.Contract.SetTokenScope(&_BindBox.TransactOpts, _fromTokenId, _toTokenId)
}

// SetTokenScope is a paid mutator transaction binding the contract method 0x6663a375.
//
// Solidity: function setTokenScope(uint32 _fromTokenId, uint32 _toTokenId) returns()
func (_BindBox *BindBoxTransactorSession) SetTokenScope(_fromTokenId uint32, _toTokenId uint32) (*types.Transaction, error) {
	return _BindBox.Contract.SetTokenScope(&_BindBox.TransactOpts, _fromTokenId, _toTokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BindBox *BindBoxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BindBox.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BindBox *BindBoxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BindBox.Contract.TransferOwnership(&_BindBox.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BindBox *BindBoxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BindBox.Contract.TransferOwnership(&_BindBox.TransactOpts, newOwner)
}

// UpdateFundMgr is a paid mutator transaction binding the contract method 0x9d73361c.
//
// Solidity: function updateFundMgr(address newFundMgr) returns()
func (_BindBox *BindBoxTransactor) UpdateFundMgr(opts *bind.TransactOpts, newFundMgr common.Address) (*types.Transaction, error) {
	return _BindBox.contract.Transact(opts, "updateFundMgr", newFundMgr)
}

// UpdateFundMgr is a paid mutator transaction binding the contract method 0x9d73361c.
//
// Solidity: function updateFundMgr(address newFundMgr) returns()
func (_BindBox *BindBoxSession) UpdateFundMgr(newFundMgr common.Address) (*types.Transaction, error) {
	return _BindBox.Contract.UpdateFundMgr(&_BindBox.TransactOpts, newFundMgr)
}

// UpdateFundMgr is a paid mutator transaction binding the contract method 0x9d73361c.
//
// Solidity: function updateFundMgr(address newFundMgr) returns()
func (_BindBox *BindBoxTransactorSession) UpdateFundMgr(newFundMgr common.Address) (*types.Transaction, error) {
	return _BindBox.Contract.UpdateFundMgr(&_BindBox.TransactOpts, newFundMgr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BindBox *BindBoxTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindBox.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BindBox *BindBoxSession) Withdraw() (*types.Transaction, error) {
	return _BindBox.Contract.Withdraw(&_BindBox.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_BindBox *BindBoxTransactorSession) Withdraw() (*types.Transaction, error) {
	return _BindBox.Contract.Withdraw(&_BindBox.TransactOpts)
}

// BindBoxClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the BindBox contract.
type BindBoxClaimedIterator struct {
	Event *BindBoxClaimed // Event containing the contract specifics and raw log

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
func (it *BindBoxClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindBoxClaimed)
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
		it.Event = new(BindBoxClaimed)
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
func (it *BindBoxClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindBoxClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindBoxClaimed represents a Claimed event raised by the BindBox contract.
type BindBoxClaimed struct {
	Claimer common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed claimer, uint256 indexed tokenId)
func (_BindBox *BindBoxFilterer) FilterClaimed(opts *bind.FilterOpts, claimer []common.Address, tokenId []*big.Int) (*BindBoxClaimedIterator, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BindBox.contract.FilterLogs(opts, "Claimed", claimerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BindBoxClaimedIterator{contract: _BindBox.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed claimer, uint256 indexed tokenId)
func (_BindBox *BindBoxFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *BindBoxClaimed, claimer []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BindBox.contract.WatchLogs(opts, "Claimed", claimerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindBoxClaimed)
				if err := _BindBox.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address indexed claimer, uint256 indexed tokenId)
func (_BindBox *BindBoxFilterer) ParseClaimed(log types.Log) (*BindBoxClaimed, error) {
	event := new(BindBoxClaimed)
	if err := _BindBox.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindBoxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BindBox contract.
type BindBoxOwnershipTransferredIterator struct {
	Event *BindBoxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BindBoxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindBoxOwnershipTransferred)
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
		it.Event = new(BindBoxOwnershipTransferred)
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
func (it *BindBoxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindBoxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindBoxOwnershipTransferred represents a OwnershipTransferred event raised by the BindBox contract.
type BindBoxOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BindBox *BindBoxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BindBoxOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BindBox.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BindBoxOwnershipTransferredIterator{contract: _BindBox.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BindBox *BindBoxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BindBoxOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BindBox.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindBoxOwnershipTransferred)
				if err := _BindBox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BindBox *BindBoxFilterer) ParseOwnershipTransferred(log types.Log) (*BindBoxOwnershipTransferred, error) {
	event := new(BindBoxOwnershipTransferred)
	if err := _BindBox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindBoxRolledOverIterator is returned from FilterRolledOver and is used to iterate over the raw logs and unpacked data for RolledOver events raised by the BindBox contract.
type BindBoxRolledOverIterator struct {
	Event *BindBoxRolledOver // Event containing the contract specifics and raw log

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
func (it *BindBoxRolledOverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindBoxRolledOver)
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
		it.Event = new(BindBoxRolledOver)
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
func (it *BindBoxRolledOverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindBoxRolledOverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindBoxRolledOver represents a RolledOver event raised by the BindBox contract.
type BindBoxRolledOver struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRolledOver is a free log retrieval operation binding the contract event 0xd2bce5b0f3cb0c9c70f1761d4ddcaf6236907e4b29bb4a7546b6a0c64be90cf8.
//
// Solidity: event RolledOver(bool status)
func (_BindBox *BindBoxFilterer) FilterRolledOver(opts *bind.FilterOpts) (*BindBoxRolledOverIterator, error) {

	logs, sub, err := _BindBox.contract.FilterLogs(opts, "RolledOver")
	if err != nil {
		return nil, err
	}
	return &BindBoxRolledOverIterator{contract: _BindBox.contract, event: "RolledOver", logs: logs, sub: sub}, nil
}

// WatchRolledOver is a free log subscription operation binding the contract event 0xd2bce5b0f3cb0c9c70f1761d4ddcaf6236907e4b29bb4a7546b6a0c64be90cf8.
//
// Solidity: event RolledOver(bool status)
func (_BindBox *BindBoxFilterer) WatchRolledOver(opts *bind.WatchOpts, sink chan<- *BindBoxRolledOver) (event.Subscription, error) {

	logs, sub, err := _BindBox.contract.WatchLogs(opts, "RolledOver")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindBoxRolledOver)
				if err := _BindBox.contract.UnpackLog(event, "RolledOver", log); err != nil {
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
func (_BindBox *BindBoxFilterer) ParseRolledOver(log types.Log) (*BindBoxRolledOver, error) {
	event := new(BindBoxRolledOver)
	if err := _BindBox.contract.UnpackLog(event, "RolledOver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
