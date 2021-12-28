// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"val1\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val2\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"val3\",\"type\":\"bool\"}],\"name\":\"ExampleEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"val1\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"val2\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"val3\",\"type\":\"bool\"}],\"name\":\"exampleFunction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// ExampleFunction is a paid mutator transaction binding the contract method 0xa3cba26d.
//
// Solidity: function exampleFunction(string val1, uint256 val2, bool val3) returns()
func (_Main *MainTransactor) ExampleFunction(opts *bind.TransactOpts, val1 string, val2 *big.Int, val3 bool) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "exampleFunction", val1, val2, val3)
}

// ExampleFunction is a paid mutator transaction binding the contract method 0xa3cba26d.
//
// Solidity: function exampleFunction(string val1, uint256 val2, bool val3) returns()
func (_Main *MainSession) ExampleFunction(val1 string, val2 *big.Int, val3 bool) (*types.Transaction, error) {
	return _Main.Contract.ExampleFunction(&_Main.TransactOpts, val1, val2, val3)
}

// ExampleFunction is a paid mutator transaction binding the contract method 0xa3cba26d.
//
// Solidity: function exampleFunction(string val1, uint256 val2, bool val3) returns()
func (_Main *MainTransactorSession) ExampleFunction(val1 string, val2 *big.Int, val3 bool) (*types.Transaction, error) {
	return _Main.Contract.ExampleFunction(&_Main.TransactOpts, val1, val2, val3)
}

// MainExampleEventIterator is returned from FilterExampleEvent and is used to iterate over the raw logs and unpacked data for ExampleEvent events raised by the Main contract.
type MainExampleEventIterator struct {
	Event *MainExampleEvent // Event containing the contract specifics and raw log

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
func (it *MainExampleEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainExampleEvent)
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
		it.Event = new(MainExampleEvent)
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
func (it *MainExampleEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainExampleEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainExampleEvent represents a ExampleEvent event raised by the Main contract.
type MainExampleEvent struct {
	Val1 string
	Val2 *big.Int
	Val3 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterExampleEvent is a free log retrieval operation binding the contract event 0x31a41810e5816fde8b8185333f120aa1f74d973d9939679a7a0ad63ca3925806.
//
// Solidity: event ExampleEvent(string val1, uint256 val2, bool val3)
func (_Main *MainFilterer) FilterExampleEvent(opts *bind.FilterOpts) (*MainExampleEventIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "ExampleEvent")
	if err != nil {
		return nil, err
	}
	return &MainExampleEventIterator{contract: _Main.contract, event: "ExampleEvent", logs: logs, sub: sub}, nil
}

// WatchExampleEvent is a free log subscription operation binding the contract event 0x31a41810e5816fde8b8185333f120aa1f74d973d9939679a7a0ad63ca3925806.
//
// Solidity: event ExampleEvent(string val1, uint256 val2, bool val3)
func (_Main *MainFilterer) WatchExampleEvent(opts *bind.WatchOpts, sink chan<- *MainExampleEvent) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "ExampleEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainExampleEvent)
				if err := _Main.contract.UnpackLog(event, "ExampleEvent", log); err != nil {
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

// ParseExampleEvent is a log parse operation binding the contract event 0x31a41810e5816fde8b8185333f120aa1f74d973d9939679a7a0ad63ca3925806.
//
// Solidity: event ExampleEvent(string val1, uint256 val2, bool val3)
func (_Main *MainFilterer) ParseExampleEvent(log types.Log) (*MainExampleEvent, error) {
	event := new(MainExampleEvent)
	if err := _Main.contract.UnpackLog(event, "ExampleEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
