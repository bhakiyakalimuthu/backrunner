// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package executor

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

// ExecutorMetaData contains all meta data concerning the Executor contract.
var ExecutorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sushiPair\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_uniPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"sushiZeroForOne\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"uniZeroForOne\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"buyAmount\",\"type\":\"int256\"},{\"internalType\":\"uint160\",\"name\":\"uniSqrtPriceLimitX96\",\"type\":\"uint160\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600080546001600160a01b031990811633179091556001805490911673c02aaa39b223fe8d0a0e5c4f27ead9083c756cc217905561055e806100546000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806328a070251461003b578063a809879514610045575b600080fd5b61004361006a565b005b61005861005336600461043a565b610155565b60405190815260200160405180910390f35b6001546000546040516370a0823160e01b81523060048201526001600160a01b039283169263a9059cbb92169083906370a0823190602401602060405180830381865afa1580156100bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100e391906104c7565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af115801561012e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061015291906104e0565b50565b600080546001600160a01b031633146101a35760405162461bcd60e51b815260206004820152600c60248201526b1d5b985d5d1a1bdc9a5e995960a21b604482015260640160405180910390fd5b604051630251596160e31b81523060048201528415156024820152604481018490526001600160a01b03838116606483015260a06084830152600060a483015287918a918a919082169063128acb089060c40160408051808303816000875af1158015610214573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102389190610504565b5050816001600160a01b031663022c0d9f896102555760006102bd565b6040516370a0823160e01b81523060048201526001600160a01b038616906370a0823190602401602060405180830381865afa158015610299573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102bd91906104c7565b8a61032f576040516370a0823160e01b81523060048201526001600160a01b038716906370a0823190602401602060405180830381865afa158015610306573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061032a91906104c7565b610332565b60005b6040516001600160e01b031960e085901b16815260048101929092526024820152306044820152608060648201526000608482015260a401600060405180830381600087803b15801561038457600080fd5b505af1158015610398573d6000803e3d6000fd5b50506001546040516370a0823160e01b81523060048201526001600160a01b0390911692506370a082319150602401602060405180830381865afa1580156103e4573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061040891906104c7565b9b9a5050505050505050505050565b6001600160a01b038116811461015257600080fd5b801515811461015257600080fd5b600080600080600080600060e0888a03121561045557600080fd5b873561046081610417565b9650602088013561047081610417565b9550604088013561048081610417565b945060608801356104908161042c565b935060808801356104a08161042c565b925060a0880135915060c08801356104b781610417565b8091505092959891949750929550565b6000602082840312156104d957600080fd5b5051919050565b6000602082840312156104f257600080fd5b81516104fd8161042c565b9392505050565b6000806040838503121561051757600080fd5b50508051602090910151909290915056fea264697066735822122082eff0790fd9dceec37662c7ab8a17149de904c241c0482b22556c137a5e8abc64736f6c63430008120033",
}

// ExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use ExecutorMetaData.ABI instead.
var ExecutorABI = ExecutorMetaData.ABI

// ExecutorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExecutorMetaData.Bin instead.
var ExecutorBin = ExecutorMetaData.Bin

// DeployExecutor deploys a new Ethereum contract, binding an instance of Executor to it.
func DeployExecutor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Executor, error) {
	parsed, err := ExecutorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExecutorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Executor{ExecutorCaller: ExecutorCaller{contract: contract}, ExecutorTransactor: ExecutorTransactor{contract: contract}, ExecutorFilterer: ExecutorFilterer{contract: contract}}, nil
}

// Executor is an auto generated Go binding around an Ethereum contract.
type Executor struct {
	ExecutorCaller     // Read-only binding to the contract
	ExecutorTransactor // Write-only binding to the contract
	ExecutorFilterer   // Log filterer for contract events
}

// ExecutorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecutorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecutorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecutorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecutorSession struct {
	Contract     *Executor         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExecutorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecutorCallerSession struct {
	Contract *ExecutorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExecutorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecutorTransactorSession struct {
	Contract     *ExecutorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExecutorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecutorRaw struct {
	Contract *Executor // Generic contract binding to access the raw methods on
}

// ExecutorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecutorCallerRaw struct {
	Contract *ExecutorCaller // Generic read-only contract binding to access the raw methods on
}

// ExecutorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecutorTransactorRaw struct {
	Contract *ExecutorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExecutor creates a new instance of Executor, bound to a specific deployed contract.
func NewExecutor(address common.Address, backend bind.ContractBackend) (*Executor, error) {
	contract, err := bindExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Executor{ExecutorCaller: ExecutorCaller{contract: contract}, ExecutorTransactor: ExecutorTransactor{contract: contract}, ExecutorFilterer: ExecutorFilterer{contract: contract}}, nil
}

// NewExecutorCaller creates a new read-only instance of Executor, bound to a specific deployed contract.
func NewExecutorCaller(address common.Address, caller bind.ContractCaller) (*ExecutorCaller, error) {
	contract, err := bindExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorCaller{contract: contract}, nil
}

// NewExecutorTransactor creates a new write-only instance of Executor, bound to a specific deployed contract.
func NewExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutorTransactor, error) {
	contract, err := bindExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorTransactor{contract: contract}, nil
}

// NewExecutorFilterer creates a new log filterer instance of Executor, bound to a specific deployed contract.
func NewExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutorFilterer, error) {
	contract, err := bindExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutorFilterer{contract: contract}, nil
}

// bindExecutor binds a generic wrapper to an already deployed contract.
func bindExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExecutorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Executor *ExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Executor.Contract.ExecutorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Executor *ExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executor.Contract.ExecutorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Executor *ExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Executor.Contract.ExecutorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Executor *ExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Executor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Executor *ExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Executor *ExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Executor.Contract.contract.Transact(opts, method, params...)
}

// Execute is a paid mutator transaction binding the contract method 0xa8098795.
//
// Solidity: function execute(address _sushiPair, address _uniPool, address _token, bool sushiZeroForOne, bool uniZeroForOne, int256 buyAmount, uint160 uniSqrtPriceLimitX96) returns(uint256)
func (_Executor *ExecutorTransactor) Execute(opts *bind.TransactOpts, _sushiPair common.Address, _uniPool common.Address, _token common.Address, sushiZeroForOne bool, uniZeroForOne bool, buyAmount *big.Int, uniSqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "execute", _sushiPair, _uniPool, _token, sushiZeroForOne, uniZeroForOne, buyAmount, uniSqrtPriceLimitX96)
}

// Execute is a paid mutator transaction binding the contract method 0xa8098795.
//
// Solidity: function execute(address _sushiPair, address _uniPool, address _token, bool sushiZeroForOne, bool uniZeroForOne, int256 buyAmount, uint160 uniSqrtPriceLimitX96) returns(uint256)
func (_Executor *ExecutorSession) Execute(_sushiPair common.Address, _uniPool common.Address, _token common.Address, sushiZeroForOne bool, uniZeroForOne bool, buyAmount *big.Int, uniSqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Executor.Contract.Execute(&_Executor.TransactOpts, _sushiPair, _uniPool, _token, sushiZeroForOne, uniZeroForOne, buyAmount, uniSqrtPriceLimitX96)
}

// Execute is a paid mutator transaction binding the contract method 0xa8098795.
//
// Solidity: function execute(address _sushiPair, address _uniPool, address _token, bool sushiZeroForOne, bool uniZeroForOne, int256 buyAmount, uint160 uniSqrtPriceLimitX96) returns(uint256)
func (_Executor *ExecutorTransactorSession) Execute(_sushiPair common.Address, _uniPool common.Address, _token common.Address, sushiZeroForOne bool, uniZeroForOne bool, buyAmount *big.Int, uniSqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Executor.Contract.Execute(&_Executor.TransactOpts, _sushiPair, _uniPool, _token, sushiZeroForOne, uniZeroForOne, buyAmount, uniSqrtPriceLimitX96)
}

// Liquidate is a paid mutator transaction binding the contract method 0x28a07025.
//
// Solidity: function liquidate() returns()
func (_Executor *ExecutorTransactor) Liquidate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "liquidate")
}

// Liquidate is a paid mutator transaction binding the contract method 0x28a07025.
//
// Solidity: function liquidate() returns()
func (_Executor *ExecutorSession) Liquidate() (*types.Transaction, error) {
	return _Executor.Contract.Liquidate(&_Executor.TransactOpts)
}

// Liquidate is a paid mutator transaction binding the contract method 0x28a07025.
//
// Solidity: function liquidate() returns()
func (_Executor *ExecutorTransactorSession) Liquidate() (*types.Transaction, error) {
	return _Executor.Contract.Liquidate(&_Executor.TransactOpts)
}
