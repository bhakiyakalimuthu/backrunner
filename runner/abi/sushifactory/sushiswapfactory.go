// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sushiswapfactory

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

// SushiswapfactoryMetaData contains all meta data concerning the Sushiswapfactory contract.
var SushiswapfactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairCodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_migrator\",\"type\":\"address\"}],\"name\":\"setMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SushiswapfactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use SushiswapfactoryMetaData.ABI instead.
var SushiswapfactoryABI = SushiswapfactoryMetaData.ABI

// Sushiswapfactory is an auto generated Go binding around an Ethereum contract.
type Sushiswapfactory struct {
	SushiswapfactoryCaller     // Read-only binding to the contract
	SushiswapfactoryTransactor // Write-only binding to the contract
	SushiswapfactoryFilterer   // Log filterer for contract events
}

// SushiswapfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SushiswapfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushiswapfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SushiswapfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushiswapfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SushiswapfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushiswapfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SushiswapfactorySession struct {
	Contract     *Sushiswapfactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SushiswapfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SushiswapfactoryCallerSession struct {
	Contract *SushiswapfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SushiswapfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SushiswapfactoryTransactorSession struct {
	Contract     *SushiswapfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SushiswapfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SushiswapfactoryRaw struct {
	Contract *Sushiswapfactory // Generic contract binding to access the raw methods on
}

// SushiswapfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SushiswapfactoryCallerRaw struct {
	Contract *SushiswapfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// SushiswapfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SushiswapfactoryTransactorRaw struct {
	Contract *SushiswapfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSushiswapfactory creates a new instance of Sushiswapfactory, bound to a specific deployed contract.
func NewSushiswapfactory(address common.Address, backend bind.ContractBackend) (*Sushiswapfactory, error) {
	contract, err := bindSushiswapfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sushiswapfactory{SushiswapfactoryCaller: SushiswapfactoryCaller{contract: contract}, SushiswapfactoryTransactor: SushiswapfactoryTransactor{contract: contract}, SushiswapfactoryFilterer: SushiswapfactoryFilterer{contract: contract}}, nil
}

// NewSushiswapfactoryCaller creates a new read-only instance of Sushiswapfactory, bound to a specific deployed contract.
func NewSushiswapfactoryCaller(address common.Address, caller bind.ContractCaller) (*SushiswapfactoryCaller, error) {
	contract, err := bindSushiswapfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SushiswapfactoryCaller{contract: contract}, nil
}

// NewSushiswapfactoryTransactor creates a new write-only instance of Sushiswapfactory, bound to a specific deployed contract.
func NewSushiswapfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*SushiswapfactoryTransactor, error) {
	contract, err := bindSushiswapfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SushiswapfactoryTransactor{contract: contract}, nil
}

// NewSushiswapfactoryFilterer creates a new log filterer instance of Sushiswapfactory, bound to a specific deployed contract.
func NewSushiswapfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*SushiswapfactoryFilterer, error) {
	contract, err := bindSushiswapfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SushiswapfactoryFilterer{contract: contract}, nil
}

// bindSushiswapfactory binds a generic wrapper to an already deployed contract.
func bindSushiswapfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SushiswapfactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sushiswapfactory *SushiswapfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sushiswapfactory.Contract.SushiswapfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sushiswapfactory *SushiswapfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sushiswapfactory.Contract.SushiswapfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sushiswapfactory *SushiswapfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sushiswapfactory.Contract.SushiswapfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sushiswapfactory *SushiswapfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sushiswapfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sushiswapfactory *SushiswapfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sushiswapfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sushiswapfactory *SushiswapfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sushiswapfactory.Contract.contract.Transact(opts, method, params...)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Sushiswapfactory *SushiswapfactoryCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Sushiswapfactory.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Sushiswapfactory *SushiswapfactorySession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _Sushiswapfactory.Contract.AllPairs(&_Sushiswapfactory.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Sushiswapfactory *SushiswapfactoryCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _Sushiswapfactory.Contract.AllPairs(&_Sushiswapfactory.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Sushiswapfactory *SushiswapfactoryCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sushiswapfactory.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Sushiswapfactory *SushiswapfactorySession) AllPairsLength() (*big.Int, error) {
	return _Sushiswapfactory.Contract.AllPairsLength(&_Sushiswapfactory.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Sushiswapfactory *SushiswapfactoryCallerSession) AllPairsLength() (*big.Int, error) {
	return _Sushiswapfactory.Contract.AllPairsLength(&_Sushiswapfactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Sushiswapfactory *SushiswapfactoryCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sushiswapfactory.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Sushiswapfactory *SushiswapfactorySession) FeeTo() (common.Address, error) {
	return _Sushiswapfactory.Contract.FeeTo(&_Sushiswapfactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Sushiswapfactory *SushiswapfactoryCallerSession) FeeTo() (common.Address, error) {
	return _Sushiswapfactory.Contract.FeeTo(&_Sushiswapfactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Sushiswapfactory *SushiswapfactoryCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sushiswapfactory.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Sushiswapfactory *SushiswapfactorySession) FeeToSetter() (common.Address, error) {
	return _Sushiswapfactory.Contract.FeeToSetter(&_Sushiswapfactory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Sushiswapfactory *SushiswapfactoryCallerSession) FeeToSetter() (common.Address, error) {
	return _Sushiswapfactory.Contract.FeeToSetter(&_Sushiswapfactory.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_Sushiswapfactory *SushiswapfactoryCaller) GetPair(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Sushiswapfactory.contract.Call(opts, &out, "getPair", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_Sushiswapfactory *SushiswapfactorySession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _Sushiswapfactory.Contract.GetPair(&_Sushiswapfactory.CallOpts, arg0, arg1)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_Sushiswapfactory *SushiswapfactoryCallerSession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _Sushiswapfactory.Contract.GetPair(&_Sushiswapfactory.CallOpts, arg0, arg1)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_Sushiswapfactory *SushiswapfactoryCaller) Migrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sushiswapfactory.contract.Call(opts, &out, "migrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_Sushiswapfactory *SushiswapfactorySession) Migrator() (common.Address, error) {
	return _Sushiswapfactory.Contract.Migrator(&_Sushiswapfactory.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_Sushiswapfactory *SushiswapfactoryCallerSession) Migrator() (common.Address, error) {
	return _Sushiswapfactory.Contract.Migrator(&_Sushiswapfactory.CallOpts)
}

// PairCodeHash is a free data retrieval call binding the contract method 0x9aab9248.
//
// Solidity: function pairCodeHash() pure returns(bytes32)
func (_Sushiswapfactory *SushiswapfactoryCaller) PairCodeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Sushiswapfactory.contract.Call(opts, &out, "pairCodeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PairCodeHash is a free data retrieval call binding the contract method 0x9aab9248.
//
// Solidity: function pairCodeHash() pure returns(bytes32)
func (_Sushiswapfactory *SushiswapfactorySession) PairCodeHash() ([32]byte, error) {
	return _Sushiswapfactory.Contract.PairCodeHash(&_Sushiswapfactory.CallOpts)
}

// PairCodeHash is a free data retrieval call binding the contract method 0x9aab9248.
//
// Solidity: function pairCodeHash() pure returns(bytes32)
func (_Sushiswapfactory *SushiswapfactoryCallerSession) PairCodeHash() ([32]byte, error) {
	return _Sushiswapfactory.Contract.PairCodeHash(&_Sushiswapfactory.CallOpts)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Sushiswapfactory *SushiswapfactoryTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Sushiswapfactory.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Sushiswapfactory *SushiswapfactorySession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Sushiswapfactory.Contract.CreatePair(&_Sushiswapfactory.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Sushiswapfactory *SushiswapfactoryTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Sushiswapfactory.Contract.CreatePair(&_Sushiswapfactory.TransactOpts, tokenA, tokenB)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Sushiswapfactory *SushiswapfactoryTransactor) SetFeeTo(opts *bind.TransactOpts, _feeTo common.Address) (*types.Transaction, error) {
	return _Sushiswapfactory.contract.Transact(opts, "setFeeTo", _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Sushiswapfactory *SushiswapfactorySession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _Sushiswapfactory.Contract.SetFeeTo(&_Sushiswapfactory.TransactOpts, _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Sushiswapfactory *SushiswapfactoryTransactorSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _Sushiswapfactory.Contract.SetFeeTo(&_Sushiswapfactory.TransactOpts, _feeTo)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_Sushiswapfactory *SushiswapfactoryTransactor) SetFeeToSetter(opts *bind.TransactOpts, _feeToSetter common.Address) (*types.Transaction, error) {
	return _Sushiswapfactory.contract.Transact(opts, "setFeeToSetter", _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_Sushiswapfactory *SushiswapfactorySession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _Sushiswapfactory.Contract.SetFeeToSetter(&_Sushiswapfactory.TransactOpts, _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_Sushiswapfactory *SushiswapfactoryTransactorSession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _Sushiswapfactory.Contract.SetFeeToSetter(&_Sushiswapfactory.TransactOpts, _feeToSetter)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_Sushiswapfactory *SushiswapfactoryTransactor) SetMigrator(opts *bind.TransactOpts, _migrator common.Address) (*types.Transaction, error) {
	return _Sushiswapfactory.contract.Transact(opts, "setMigrator", _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_Sushiswapfactory *SushiswapfactorySession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _Sushiswapfactory.Contract.SetMigrator(&_Sushiswapfactory.TransactOpts, _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_Sushiswapfactory *SushiswapfactoryTransactorSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _Sushiswapfactory.Contract.SetMigrator(&_Sushiswapfactory.TransactOpts, _migrator)
}

// SushiswapfactoryPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the Sushiswapfactory contract.
type SushiswapfactoryPairCreatedIterator struct {
	Event *SushiswapfactoryPairCreated // Event containing the contract specifics and raw log

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
func (it *SushiswapfactoryPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SushiswapfactoryPairCreated)
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
		it.Event = new(SushiswapfactoryPairCreated)
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
func (it *SushiswapfactoryPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SushiswapfactoryPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SushiswapfactoryPairCreated represents a PairCreated event raised by the Sushiswapfactory contract.
type SushiswapfactoryPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Sushiswapfactory *SushiswapfactoryFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*SushiswapfactoryPairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Sushiswapfactory.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &SushiswapfactoryPairCreatedIterator{contract: _Sushiswapfactory.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Sushiswapfactory *SushiswapfactoryFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *SushiswapfactoryPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Sushiswapfactory.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SushiswapfactoryPairCreated)
				if err := _Sushiswapfactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Sushiswapfactory *SushiswapfactoryFilterer) ParsePairCreated(log types.Log) (*SushiswapfactoryPairCreated, error) {
	event := new(SushiswapfactoryPairCreated)
	if err := _Sushiswapfactory.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
