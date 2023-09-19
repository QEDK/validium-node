// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mockdabridgerouter

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

// MockdabridgerouterMetaData contains all meta data concerning the Mockdabridgerouter contract.
var MockdabridgerouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"roots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506101168061001d5f395ff3fe6080604052348015600e575f80fd5b50600436106030575f3560e01c8063081dc681146034578063967326c1146062575b5f80fd5b6050603f366004609e565b5f6020819052908152604090205481565b60405190815260200160405180910390f35b6085606d36600460bb565b63ffffffff9091165f90815260208190526040902055565b005b803563ffffffff811681146099575f80fd5b919050565b5f6020828403121560ad575f80fd5b60b4826087565b9392505050565b5f806040838503121560cb575f80fd5b60d2836087565b94602093909301359350505056fea26469706673582212200b9ddaa5be40fe9cd57198b5c33cf11c9e7201ecf9d0b6ab7ba1e1f2066a270864736f6c63430008140033",
}

// MockdabridgerouterABI is the input ABI used to generate the binding from.
// Deprecated: Use MockdabridgerouterMetaData.ABI instead.
var MockdabridgerouterABI = MockdabridgerouterMetaData.ABI

// MockdabridgerouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockdabridgerouterMetaData.Bin instead.
var MockdabridgerouterBin = MockdabridgerouterMetaData.Bin

// DeployMockdabridgerouter deploys a new Ethereum contract, binding an instance of Mockdabridgerouter to it.
func DeployMockdabridgerouter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Mockdabridgerouter, error) {
	parsed, err := MockdabridgerouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockdabridgerouterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mockdabridgerouter{MockdabridgerouterCaller: MockdabridgerouterCaller{contract: contract}, MockdabridgerouterTransactor: MockdabridgerouterTransactor{contract: contract}, MockdabridgerouterFilterer: MockdabridgerouterFilterer{contract: contract}}, nil
}

// Mockdabridgerouter is an auto generated Go binding around an Ethereum contract.
type Mockdabridgerouter struct {
	MockdabridgerouterCaller     // Read-only binding to the contract
	MockdabridgerouterTransactor // Write-only binding to the contract
	MockdabridgerouterFilterer   // Log filterer for contract events
}

// MockdabridgerouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockdabridgerouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockdabridgerouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockdabridgerouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockdabridgerouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockdabridgerouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockdabridgerouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockdabridgerouterSession struct {
	Contract     *Mockdabridgerouter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MockdabridgerouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockdabridgerouterCallerSession struct {
	Contract *MockdabridgerouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MockdabridgerouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockdabridgerouterTransactorSession struct {
	Contract     *MockdabridgerouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MockdabridgerouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockdabridgerouterRaw struct {
	Contract *Mockdabridgerouter // Generic contract binding to access the raw methods on
}

// MockdabridgerouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockdabridgerouterCallerRaw struct {
	Contract *MockdabridgerouterCaller // Generic read-only contract binding to access the raw methods on
}

// MockdabridgerouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockdabridgerouterTransactorRaw struct {
	Contract *MockdabridgerouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockdabridgerouter creates a new instance of Mockdabridgerouter, bound to a specific deployed contract.
func NewMockdabridgerouter(address common.Address, backend bind.ContractBackend) (*Mockdabridgerouter, error) {
	contract, err := bindMockdabridgerouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mockdabridgerouter{MockdabridgerouterCaller: MockdabridgerouterCaller{contract: contract}, MockdabridgerouterTransactor: MockdabridgerouterTransactor{contract: contract}, MockdabridgerouterFilterer: MockdabridgerouterFilterer{contract: contract}}, nil
}

// NewMockdabridgerouterCaller creates a new read-only instance of Mockdabridgerouter, bound to a specific deployed contract.
func NewMockdabridgerouterCaller(address common.Address, caller bind.ContractCaller) (*MockdabridgerouterCaller, error) {
	contract, err := bindMockdabridgerouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockdabridgerouterCaller{contract: contract}, nil
}

// NewMockdabridgerouterTransactor creates a new write-only instance of Mockdabridgerouter, bound to a specific deployed contract.
func NewMockdabridgerouterTransactor(address common.Address, transactor bind.ContractTransactor) (*MockdabridgerouterTransactor, error) {
	contract, err := bindMockdabridgerouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockdabridgerouterTransactor{contract: contract}, nil
}

// NewMockdabridgerouterFilterer creates a new log filterer instance of Mockdabridgerouter, bound to a specific deployed contract.
func NewMockdabridgerouterFilterer(address common.Address, filterer bind.ContractFilterer) (*MockdabridgerouterFilterer, error) {
	contract, err := bindMockdabridgerouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockdabridgerouterFilterer{contract: contract}, nil
}

// bindMockdabridgerouter binds a generic wrapper to an already deployed contract.
func bindMockdabridgerouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockdabridgerouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mockdabridgerouter *MockdabridgerouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mockdabridgerouter.Contract.MockdabridgerouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mockdabridgerouter *MockdabridgerouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mockdabridgerouter.Contract.MockdabridgerouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mockdabridgerouter *MockdabridgerouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mockdabridgerouter.Contract.MockdabridgerouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mockdabridgerouter *MockdabridgerouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mockdabridgerouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mockdabridgerouter *MockdabridgerouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mockdabridgerouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mockdabridgerouter *MockdabridgerouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mockdabridgerouter.Contract.contract.Transact(opts, method, params...)
}

// Roots is a free data retrieval call binding the contract method 0x081dc681.
//
// Solidity: function roots(uint32 ) view returns(bytes32)
func (_Mockdabridgerouter *MockdabridgerouterCaller) Roots(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _Mockdabridgerouter.contract.Call(opts, &out, "roots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Roots is a free data retrieval call binding the contract method 0x081dc681.
//
// Solidity: function roots(uint32 ) view returns(bytes32)
func (_Mockdabridgerouter *MockdabridgerouterSession) Roots(arg0 uint32) ([32]byte, error) {
	return _Mockdabridgerouter.Contract.Roots(&_Mockdabridgerouter.CallOpts, arg0)
}

// Roots is a free data retrieval call binding the contract method 0x081dc681.
//
// Solidity: function roots(uint32 ) view returns(bytes32)
func (_Mockdabridgerouter *MockdabridgerouterCallerSession) Roots(arg0 uint32) ([32]byte, error) {
	return _Mockdabridgerouter.Contract.Roots(&_Mockdabridgerouter.CallOpts, arg0)
}

// Set is a paid mutator transaction binding the contract method 0x967326c1.
//
// Solidity: function set(uint32 blockNumber, bytes32 root) returns()
func (_Mockdabridgerouter *MockdabridgerouterTransactor) Set(opts *bind.TransactOpts, blockNumber uint32, root [32]byte) (*types.Transaction, error) {
	return _Mockdabridgerouter.contract.Transact(opts, "set", blockNumber, root)
}

// Set is a paid mutator transaction binding the contract method 0x967326c1.
//
// Solidity: function set(uint32 blockNumber, bytes32 root) returns()
func (_Mockdabridgerouter *MockdabridgerouterSession) Set(blockNumber uint32, root [32]byte) (*types.Transaction, error) {
	return _Mockdabridgerouter.Contract.Set(&_Mockdabridgerouter.TransactOpts, blockNumber, root)
}

// Set is a paid mutator transaction binding the contract method 0x967326c1.
//
// Solidity: function set(uint32 blockNumber, bytes32 root) returns()
func (_Mockdabridgerouter *MockdabridgerouterTransactorSession) Set(blockNumber uint32, root [32]byte) (*types.Transaction, error) {
	return _Mockdabridgerouter.Contract.Set(&_Mockdabridgerouter.TransactOpts, blockNumber, root)
}
