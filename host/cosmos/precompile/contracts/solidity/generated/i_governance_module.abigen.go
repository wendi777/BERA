// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generated

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

// Coin is an auto generated low-level Go binding around an user-defined struct.
type Coin struct {
	Amount *big.Int
	Denom  string
}

// GovernanceModuleMetaData contains all meta data concerning the GovernanceModule contract.
var GovernanceModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"cancelProposal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCoin[]\",\"name\":\"initialDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"expedited\",\"type\":\"bool\"}],\"name\":\"submitProposal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GovernanceModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernanceModuleMetaData.ABI instead.
var GovernanceModuleABI = GovernanceModuleMetaData.ABI

// GovernanceModule is an auto generated Go binding around an Ethereum contract.
type GovernanceModule struct {
	GovernanceModuleCaller     // Read-only binding to the contract
	GovernanceModuleTransactor // Write-only binding to the contract
	GovernanceModuleFilterer   // Log filterer for contract events
}

// GovernanceModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernanceModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernanceModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernanceModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernanceModuleSession struct {
	Contract     *GovernanceModule // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernanceModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernanceModuleCallerSession struct {
	Contract *GovernanceModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// GovernanceModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernanceModuleTransactorSession struct {
	Contract     *GovernanceModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// GovernanceModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernanceModuleRaw struct {
	Contract *GovernanceModule // Generic contract binding to access the raw methods on
}

// GovernanceModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernanceModuleCallerRaw struct {
	Contract *GovernanceModuleCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernanceModuleTransactorRaw struct {
	Contract *GovernanceModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernanceModule creates a new instance of GovernanceModule, bound to a specific deployed contract.
func NewGovernanceModule(address common.Address, backend bind.ContractBackend) (*GovernanceModule, error) {
	contract, err := bindGovernanceModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GovernanceModule{GovernanceModuleCaller: GovernanceModuleCaller{contract: contract}, GovernanceModuleTransactor: GovernanceModuleTransactor{contract: contract}, GovernanceModuleFilterer: GovernanceModuleFilterer{contract: contract}}, nil
}

// NewGovernanceModuleCaller creates a new read-only instance of GovernanceModule, bound to a specific deployed contract.
func NewGovernanceModuleCaller(address common.Address, caller bind.ContractCaller) (*GovernanceModuleCaller, error) {
	contract, err := bindGovernanceModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceModuleCaller{contract: contract}, nil
}

// NewGovernanceModuleTransactor creates a new write-only instance of GovernanceModule, bound to a specific deployed contract.
func NewGovernanceModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceModuleTransactor, error) {
	contract, err := bindGovernanceModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceModuleTransactor{contract: contract}, nil
}

// NewGovernanceModuleFilterer creates a new log filterer instance of GovernanceModule, bound to a specific deployed contract.
func NewGovernanceModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceModuleFilterer, error) {
	contract, err := bindGovernanceModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceModuleFilterer{contract: contract}, nil
}

// bindGovernanceModule binds a generic wrapper to an already deployed contract.
func bindGovernanceModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GovernanceModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceModule *GovernanceModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovernanceModule.Contract.GovernanceModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceModule *GovernanceModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceModule.Contract.GovernanceModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceModule *GovernanceModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceModule.Contract.GovernanceModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceModule *GovernanceModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovernanceModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceModule *GovernanceModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceModule *GovernanceModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceModule.Contract.contract.Transact(opts, method, params...)
}

// CancelProposal is a paid mutator transaction binding the contract method 0xe0a8f6f5.
//
// Solidity: function cancelProposal(uint256 proposalId) returns(uint64, uint64)
func (_GovernanceModule *GovernanceModuleTransactor) CancelProposal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _GovernanceModule.contract.Transact(opts, "cancelProposal", proposalId)
}

// CancelProposal is a paid mutator transaction binding the contract method 0xe0a8f6f5.
//
// Solidity: function cancelProposal(uint256 proposalId) returns(uint64, uint64)
func (_GovernanceModule *GovernanceModuleSession) CancelProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _GovernanceModule.Contract.CancelProposal(&_GovernanceModule.TransactOpts, proposalId)
}

// CancelProposal is a paid mutator transaction binding the contract method 0xe0a8f6f5.
//
// Solidity: function cancelProposal(uint256 proposalId) returns(uint64, uint64)
func (_GovernanceModule *GovernanceModuleTransactorSession) CancelProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _GovernanceModule.Contract.CancelProposal(&_GovernanceModule.TransactOpts, proposalId)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x4259573d.
//
// Solidity: function submitProposal(bytes message, (uint256,string)[] initialDeposit, string metadata, string title, string summary, bool expedited) returns(uint64)
func (_GovernanceModule *GovernanceModuleTransactor) SubmitProposal(opts *bind.TransactOpts, message []byte, initialDeposit []Coin, metadata string, title string, summary string, expedited bool) (*types.Transaction, error) {
	return _GovernanceModule.contract.Transact(opts, "submitProposal", message, initialDeposit, metadata, title, summary, expedited)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x4259573d.
//
// Solidity: function submitProposal(bytes message, (uint256,string)[] initialDeposit, string metadata, string title, string summary, bool expedited) returns(uint64)
func (_GovernanceModule *GovernanceModuleSession) SubmitProposal(message []byte, initialDeposit []Coin, metadata string, title string, summary string, expedited bool) (*types.Transaction, error) {
	return _GovernanceModule.Contract.SubmitProposal(&_GovernanceModule.TransactOpts, message, initialDeposit, metadata, title, summary, expedited)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x4259573d.
//
// Solidity: function submitProposal(bytes message, (uint256,string)[] initialDeposit, string metadata, string title, string summary, bool expedited) returns(uint64)
func (_GovernanceModule *GovernanceModuleTransactorSession) SubmitProposal(message []byte, initialDeposit []Coin, metadata string, title string, summary string, expedited bool) (*types.Transaction, error) {
	return _GovernanceModule.Contract.SubmitProposal(&_GovernanceModule.TransactOpts, message, initialDeposit, metadata, title, summary, expedited)
}
