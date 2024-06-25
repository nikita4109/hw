// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletfactory

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

// WalletfactoryMetaData contains all meta data concerning the Walletfactory contract.
var WalletfactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"createWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userWallets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WalletfactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletfactoryMetaData.ABI instead.
var WalletfactoryABI = WalletfactoryMetaData.ABI

// Walletfactory is an auto generated Go binding around an Ethereum contract.
type Walletfactory struct {
	WalletfactoryCaller     // Read-only binding to the contract
	WalletfactoryTransactor // Write-only binding to the contract
	WalletfactoryFilterer   // Log filterer for contract events
}

// WalletfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletfactorySession struct {
	Contract     *Walletfactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletfactoryCallerSession struct {
	Contract *WalletfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// WalletfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletfactoryTransactorSession struct {
	Contract     *WalletfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// WalletfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletfactoryRaw struct {
	Contract *Walletfactory // Generic contract binding to access the raw methods on
}

// WalletfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletfactoryCallerRaw struct {
	Contract *WalletfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// WalletfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletfactoryTransactorRaw struct {
	Contract *WalletfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletfactory creates a new instance of Walletfactory, bound to a specific deployed contract.
func NewWalletfactory(address common.Address, backend bind.ContractBackend) (*Walletfactory, error) {
	contract, err := bindWalletfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Walletfactory{WalletfactoryCaller: WalletfactoryCaller{contract: contract}, WalletfactoryTransactor: WalletfactoryTransactor{contract: contract}, WalletfactoryFilterer: WalletfactoryFilterer{contract: contract}}, nil
}

// NewWalletfactoryCaller creates a new read-only instance of Walletfactory, bound to a specific deployed contract.
func NewWalletfactoryCaller(address common.Address, caller bind.ContractCaller) (*WalletfactoryCaller, error) {
	contract, err := bindWalletfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletfactoryCaller{contract: contract}, nil
}

// NewWalletfactoryTransactor creates a new write-only instance of Walletfactory, bound to a specific deployed contract.
func NewWalletfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletfactoryTransactor, error) {
	contract, err := bindWalletfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletfactoryTransactor{contract: contract}, nil
}

// NewWalletfactoryFilterer creates a new log filterer instance of Walletfactory, bound to a specific deployed contract.
func NewWalletfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletfactoryFilterer, error) {
	contract, err := bindWalletfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletfactoryFilterer{contract: contract}, nil
}

// bindWalletfactory binds a generic wrapper to an already deployed contract.
func bindWalletfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletfactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Walletfactory *WalletfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Walletfactory.Contract.WalletfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Walletfactory *WalletfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Walletfactory.Contract.WalletfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Walletfactory *WalletfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Walletfactory.Contract.WalletfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Walletfactory *WalletfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Walletfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Walletfactory *WalletfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Walletfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Walletfactory *WalletfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Walletfactory.Contract.contract.Transact(opts, method, params...)
}

// GetWallet is a free data retrieval call binding the contract method 0x04d0a647.
//
// Solidity: function getWallet(address user) view returns(address)
func (_Walletfactory *WalletfactoryCaller) GetWallet(opts *bind.CallOpts, user common.Address) (common.Address, error) {
	var out []interface{}
	err := _Walletfactory.contract.Call(opts, &out, "getWallet", user)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWallet is a free data retrieval call binding the contract method 0x04d0a647.
//
// Solidity: function getWallet(address user) view returns(address)
func (_Walletfactory *WalletfactorySession) GetWallet(user common.Address) (common.Address, error) {
	return _Walletfactory.Contract.GetWallet(&_Walletfactory.CallOpts, user)
}

// GetWallet is a free data retrieval call binding the contract method 0x04d0a647.
//
// Solidity: function getWallet(address user) view returns(address)
func (_Walletfactory *WalletfactoryCallerSession) GetWallet(user common.Address) (common.Address, error) {
	return _Walletfactory.Contract.GetWallet(&_Walletfactory.CallOpts, user)
}

// UserWallets is a free data retrieval call binding the contract method 0x63e6ffdd.
//
// Solidity: function userWallets(address ) view returns(address)
func (_Walletfactory *WalletfactoryCaller) UserWallets(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Walletfactory.contract.Call(opts, &out, "userWallets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserWallets is a free data retrieval call binding the contract method 0x63e6ffdd.
//
// Solidity: function userWallets(address ) view returns(address)
func (_Walletfactory *WalletfactorySession) UserWallets(arg0 common.Address) (common.Address, error) {
	return _Walletfactory.Contract.UserWallets(&_Walletfactory.CallOpts, arg0)
}

// UserWallets is a free data retrieval call binding the contract method 0x63e6ffdd.
//
// Solidity: function userWallets(address ) view returns(address)
func (_Walletfactory *WalletfactoryCallerSession) UserWallets(arg0 common.Address) (common.Address, error) {
	return _Walletfactory.Contract.UserWallets(&_Walletfactory.CallOpts, arg0)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x11ebbf24.
//
// Solidity: function createWallet() returns()
func (_Walletfactory *WalletfactoryTransactor) CreateWallet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Walletfactory.contract.Transact(opts, "createWallet")
}

// CreateWallet is a paid mutator transaction binding the contract method 0x11ebbf24.
//
// Solidity: function createWallet() returns()
func (_Walletfactory *WalletfactorySession) CreateWallet() (*types.Transaction, error) {
	return _Walletfactory.Contract.CreateWallet(&_Walletfactory.TransactOpts)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x11ebbf24.
//
// Solidity: function createWallet() returns()
func (_Walletfactory *WalletfactoryTransactorSession) CreateWallet() (*types.Transaction, error) {
	return _Walletfactory.Contract.CreateWallet(&_Walletfactory.TransactOpts)
}
