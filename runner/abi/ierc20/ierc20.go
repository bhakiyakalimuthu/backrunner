// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ierc20

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

// Ierc20MetaData contains all meta data concerning the Ierc20 contract.
var Ierc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Ierc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use Ierc20MetaData.ABI instead.
var Ierc20ABI = Ierc20MetaData.ABI

// Ierc20 is an auto generated Go binding around an Ethereum contract.
type Ierc20 struct {
	Ierc20Caller     // Read-only binding to the contract
	Ierc20Transactor // Write-only binding to the contract
	Ierc20Filterer   // Log filterer for contract events
}

// Ierc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Ierc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ierc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Ierc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ierc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ierc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ierc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ierc20Session struct {
	Contract     *Ierc20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ierc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ierc20CallerSession struct {
	Contract *Ierc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Ierc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ierc20TransactorSession struct {
	Contract     *Ierc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ierc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Ierc20Raw struct {
	Contract *Ierc20 // Generic contract binding to access the raw methods on
}

// Ierc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ierc20CallerRaw struct {
	Contract *Ierc20Caller // Generic read-only contract binding to access the raw methods on
}

// Ierc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ierc20TransactorRaw struct {
	Contract *Ierc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIerc20 creates a new instance of Ierc20, bound to a specific deployed contract.
func NewIerc20(address common.Address, backend bind.ContractBackend) (*Ierc20, error) {
	contract, err := bindIerc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ierc20{Ierc20Caller: Ierc20Caller{contract: contract}, Ierc20Transactor: Ierc20Transactor{contract: contract}, Ierc20Filterer: Ierc20Filterer{contract: contract}}, nil
}

// NewIerc20Caller creates a new read-only instance of Ierc20, bound to a specific deployed contract.
func NewIerc20Caller(address common.Address, caller bind.ContractCaller) (*Ierc20Caller, error) {
	contract, err := bindIerc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ierc20Caller{contract: contract}, nil
}

// NewIerc20Transactor creates a new write-only instance of Ierc20, bound to a specific deployed contract.
func NewIerc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Ierc20Transactor, error) {
	contract, err := bindIerc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ierc20Transactor{contract: contract}, nil
}

// NewIerc20Filterer creates a new log filterer instance of Ierc20, bound to a specific deployed contract.
func NewIerc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Ierc20Filterer, error) {
	contract, err := bindIerc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ierc20Filterer{contract: contract}, nil
}

// bindIerc20 binds a generic wrapper to an already deployed contract.
func bindIerc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Ierc20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ierc20 *Ierc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ierc20.Contract.Ierc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ierc20 *Ierc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ierc20.Contract.Ierc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ierc20 *Ierc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ierc20.Contract.Ierc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ierc20 *Ierc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ierc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ierc20 *Ierc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ierc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ierc20 *Ierc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ierc20.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Ierc20 *Ierc20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ierc20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Ierc20 *Ierc20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _Ierc20.Contract.BalanceOf(&_Ierc20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Ierc20 *Ierc20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Ierc20.Contract.BalanceOf(&_Ierc20.CallOpts, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Ierc20 *Ierc20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ierc20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Ierc20 *Ierc20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ierc20.Contract.Transfer(&_Ierc20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Ierc20 *Ierc20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ierc20.Contract.Transfer(&_Ierc20.TransactOpts, to, amount)
}
