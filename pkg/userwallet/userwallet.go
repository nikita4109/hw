// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package userwallet

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

// UserwalletMetaData contains all meta data concerning the Userwallet contract.
var UserwalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"assets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"asset\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"buyAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"asset\",\"type\":\"string\"}],\"name\":\"getAssetBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"asset\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sellAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UserwalletABI is the input ABI used to generate the binding from.
// Deprecated: Use UserwalletMetaData.ABI instead.
var UserwalletABI = UserwalletMetaData.ABI

// Userwallet is an auto generated Go binding around an Ethereum contract.
type Userwallet struct {
	UserwalletCaller     // Read-only binding to the contract
	UserwalletTransactor // Write-only binding to the contract
	UserwalletFilterer   // Log filterer for contract events
}

// UserwalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserwalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserwalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserwalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserwalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserwalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserwalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserwalletSession struct {
	Contract     *Userwallet       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserwalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserwalletCallerSession struct {
	Contract *UserwalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// UserwalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserwalletTransactorSession struct {
	Contract     *UserwalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// UserwalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserwalletRaw struct {
	Contract *Userwallet // Generic contract binding to access the raw methods on
}

// UserwalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserwalletCallerRaw struct {
	Contract *UserwalletCaller // Generic read-only contract binding to access the raw methods on
}

// UserwalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserwalletTransactorRaw struct {
	Contract *UserwalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserwallet creates a new instance of Userwallet, bound to a specific deployed contract.
func NewUserwallet(address common.Address, backend bind.ContractBackend) (*Userwallet, error) {
	contract, err := bindUserwallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Userwallet{UserwalletCaller: UserwalletCaller{contract: contract}, UserwalletTransactor: UserwalletTransactor{contract: contract}, UserwalletFilterer: UserwalletFilterer{contract: contract}}, nil
}

// NewUserwalletCaller creates a new read-only instance of Userwallet, bound to a specific deployed contract.
func NewUserwalletCaller(address common.Address, caller bind.ContractCaller) (*UserwalletCaller, error) {
	contract, err := bindUserwallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserwalletCaller{contract: contract}, nil
}

// NewUserwalletTransactor creates a new write-only instance of Userwallet, bound to a specific deployed contract.
func NewUserwalletTransactor(address common.Address, transactor bind.ContractTransactor) (*UserwalletTransactor, error) {
	contract, err := bindUserwallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserwalletTransactor{contract: contract}, nil
}

// NewUserwalletFilterer creates a new log filterer instance of Userwallet, bound to a specific deployed contract.
func NewUserwalletFilterer(address common.Address, filterer bind.ContractFilterer) (*UserwalletFilterer, error) {
	contract, err := bindUserwallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserwalletFilterer{contract: contract}, nil
}

// bindUserwallet binds a generic wrapper to an already deployed contract.
func bindUserwallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UserwalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Userwallet *UserwalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Userwallet.Contract.UserwalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Userwallet *UserwalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Userwallet.Contract.UserwalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Userwallet *UserwalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Userwallet.Contract.UserwalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Userwallet *UserwalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Userwallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Userwallet *UserwalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Userwallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Userwallet *UserwalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Userwallet.Contract.contract.Transact(opts, method, params...)
}

// Assets is a free data retrieval call binding the contract method 0x85936228.
//
// Solidity: function assets(string ) view returns(uint256)
func (_Userwallet *UserwalletCaller) Assets(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Userwallet.contract.Call(opts, &out, "assets", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Assets is a free data retrieval call binding the contract method 0x85936228.
//
// Solidity: function assets(string ) view returns(uint256)
func (_Userwallet *UserwalletSession) Assets(arg0 string) (*big.Int, error) {
	return _Userwallet.Contract.Assets(&_Userwallet.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0x85936228.
//
// Solidity: function assets(string ) view returns(uint256)
func (_Userwallet *UserwalletCallerSession) Assets(arg0 string) (*big.Int, error) {
	return _Userwallet.Contract.Assets(&_Userwallet.CallOpts, arg0)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Userwallet *UserwalletCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Userwallet.contract.Call(opts, &out, "balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Userwallet *UserwalletSession) Balance() (*big.Int, error) {
	return _Userwallet.Contract.Balance(&_Userwallet.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Userwallet *UserwalletCallerSession) Balance() (*big.Int, error) {
	return _Userwallet.Contract.Balance(&_Userwallet.CallOpts)
}

// GetAssetBalance is a free data retrieval call binding the contract method 0xe2e51284.
//
// Solidity: function getAssetBalance(string asset) view returns(uint256)
func (_Userwallet *UserwalletCaller) GetAssetBalance(opts *bind.CallOpts, asset string) (*big.Int, error) {
	var out []interface{}
	err := _Userwallet.contract.Call(opts, &out, "getAssetBalance", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAssetBalance is a free data retrieval call binding the contract method 0xe2e51284.
//
// Solidity: function getAssetBalance(string asset) view returns(uint256)
func (_Userwallet *UserwalletSession) GetAssetBalance(asset string) (*big.Int, error) {
	return _Userwallet.Contract.GetAssetBalance(&_Userwallet.CallOpts, asset)
}

// GetAssetBalance is a free data retrieval call binding the contract method 0xe2e51284.
//
// Solidity: function getAssetBalance(string asset) view returns(uint256)
func (_Userwallet *UserwalletCallerSession) GetAssetBalance(asset string) (*big.Int, error) {
	return _Userwallet.Contract.GetAssetBalance(&_Userwallet.CallOpts, asset)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Userwallet *UserwalletCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Userwallet.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Userwallet *UserwalletSession) GetBalance() (*big.Int, error) {
	return _Userwallet.Contract.GetBalance(&_Userwallet.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Userwallet *UserwalletCallerSession) GetBalance() (*big.Int, error) {
	return _Userwallet.Contract.GetBalance(&_Userwallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Userwallet *UserwalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Userwallet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Userwallet *UserwalletSession) Owner() (common.Address, error) {
	return _Userwallet.Contract.Owner(&_Userwallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Userwallet *UserwalletCallerSession) Owner() (common.Address, error) {
	return _Userwallet.Contract.Owner(&_Userwallet.CallOpts)
}

// BuyAsset is a paid mutator transaction binding the contract method 0x71f473bf.
//
// Solidity: function buyAsset(string asset, uint256 amount) returns()
func (_Userwallet *UserwalletTransactor) BuyAsset(opts *bind.TransactOpts, asset string, amount *big.Int) (*types.Transaction, error) {
	return _Userwallet.contract.Transact(opts, "buyAsset", asset, amount)
}

// BuyAsset is a paid mutator transaction binding the contract method 0x71f473bf.
//
// Solidity: function buyAsset(string asset, uint256 amount) returns()
func (_Userwallet *UserwalletSession) BuyAsset(asset string, amount *big.Int) (*types.Transaction, error) {
	return _Userwallet.Contract.BuyAsset(&_Userwallet.TransactOpts, asset, amount)
}

// BuyAsset is a paid mutator transaction binding the contract method 0x71f473bf.
//
// Solidity: function buyAsset(string asset, uint256 amount) returns()
func (_Userwallet *UserwalletTransactorSession) BuyAsset(asset string, amount *big.Int) (*types.Transaction, error) {
	return _Userwallet.Contract.BuyAsset(&_Userwallet.TransactOpts, asset, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Userwallet *UserwalletTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Userwallet.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Userwallet *UserwalletSession) Deposit() (*types.Transaction, error) {
	return _Userwallet.Contract.Deposit(&_Userwallet.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Userwallet *UserwalletTransactorSession) Deposit() (*types.Transaction, error) {
	return _Userwallet.Contract.Deposit(&_Userwallet.TransactOpts)
}

// SellAsset is a paid mutator transaction binding the contract method 0xb32c26e8.
//
// Solidity: function sellAsset(string asset, uint256 amount) returns()
func (_Userwallet *UserwalletTransactor) SellAsset(opts *bind.TransactOpts, asset string, amount *big.Int) (*types.Transaction, error) {
	return _Userwallet.contract.Transact(opts, "sellAsset", asset, amount)
}

// SellAsset is a paid mutator transaction binding the contract method 0xb32c26e8.
//
// Solidity: function sellAsset(string asset, uint256 amount) returns()
func (_Userwallet *UserwalletSession) SellAsset(asset string, amount *big.Int) (*types.Transaction, error) {
	return _Userwallet.Contract.SellAsset(&_Userwallet.TransactOpts, asset, amount)
}

// SellAsset is a paid mutator transaction binding the contract method 0xb32c26e8.
//
// Solidity: function sellAsset(string asset, uint256 amount) returns()
func (_Userwallet *UserwalletTransactorSession) SellAsset(asset string, amount *big.Int) (*types.Transaction, error) {
	return _Userwallet.Contract.SellAsset(&_Userwallet.TransactOpts, asset, amount)
}
