// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package utils

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

// ExecutorPluginExecutionRequest is an auto generated low-level Go binding around an user-defined struct.
type ExecutorPluginExecutionRequest struct {
	Exec               TypesExecutable
	Account            common.Address
	Executor           common.Address
	ExecutorSignature  []byte
	ValidatorSignature []byte
}

// TypesExecutable is an auto generated low-level Go binding around an user-defined struct.
type TypesExecutable struct {
	CallType uint8
	Target   common.Address
	Value    *big.Int
	Data     []byte
}

// ExecutorpluginMetaData contains all meta data concerning the Executorplugin contract.
var ExecutorpluginMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddressProvider\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ModuleExecutionFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"NotGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnableToParseOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"addressProvider\",\"outputs\":[{\"internalType\":\"contractAddressProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressProviderTarget\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumTypes.CallType\",\"name\":\"callType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.Executable\",\"name\":\"exec\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"executorSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorSignature\",\"type\":\"bytes\"}],\"internalType\":\"structExecutorPlugin.ExecutionRequest\",\"name\":\"execRequest\",\"type\":\"tuple\"}],\"name\":\"executeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"executorNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executorRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"policyRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"walletRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ExecutorpluginABI is the input ABI used to generate the binding from.
// Deprecated: Use ExecutorpluginMetaData.ABI instead.
var ExecutorpluginABI = ExecutorpluginMetaData.ABI

// Executorplugin is an auto generated Go binding around an Ethereum contract.
type Executorplugin struct {
	ExecutorpluginCaller     // Read-only binding to the contract
	ExecutorpluginTransactor // Write-only binding to the contract
	ExecutorpluginFilterer   // Log filterer for contract events
}

// ExecutorpluginCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecutorpluginCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorpluginTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecutorpluginTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorpluginFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecutorpluginFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorpluginSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecutorpluginSession struct {
	Contract     *Executorplugin   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExecutorpluginCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecutorpluginCallerSession struct {
	Contract *ExecutorpluginCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ExecutorpluginTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecutorpluginTransactorSession struct {
	Contract     *ExecutorpluginTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ExecutorpluginRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecutorpluginRaw struct {
	Contract *Executorplugin // Generic contract binding to access the raw methods on
}

// ExecutorpluginCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecutorpluginCallerRaw struct {
	Contract *ExecutorpluginCaller // Generic read-only contract binding to access the raw methods on
}

// ExecutorpluginTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecutorpluginTransactorRaw struct {
	Contract *ExecutorpluginTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExecutorplugin creates a new instance of Executorplugin, bound to a specific deployed contract.
func NewExecutorplugin(address common.Address, backend bind.ContractBackend) (*Executorplugin, error) {
	contract, err := bindExecutorplugin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Executorplugin{ExecutorpluginCaller: ExecutorpluginCaller{contract: contract}, ExecutorpluginTransactor: ExecutorpluginTransactor{contract: contract}, ExecutorpluginFilterer: ExecutorpluginFilterer{contract: contract}}, nil
}

// NewExecutorpluginCaller creates a new read-only instance of Executorplugin, bound to a specific deployed contract.
func NewExecutorpluginCaller(address common.Address, caller bind.ContractCaller) (*ExecutorpluginCaller, error) {
	contract, err := bindExecutorplugin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorpluginCaller{contract: contract}, nil
}

// NewExecutorpluginTransactor creates a new write-only instance of Executorplugin, bound to a specific deployed contract.
func NewExecutorpluginTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutorpluginTransactor, error) {
	contract, err := bindExecutorplugin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorpluginTransactor{contract: contract}, nil
}

// NewExecutorpluginFilterer creates a new log filterer instance of Executorplugin, bound to a specific deployed contract.
func NewExecutorpluginFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutorpluginFilterer, error) {
	contract, err := bindExecutorplugin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutorpluginFilterer{contract: contract}, nil
}

// bindExecutorplugin binds a generic wrapper to an already deployed contract.
func bindExecutorplugin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExecutorpluginMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Executorplugin *ExecutorpluginRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Executorplugin.Contract.ExecutorpluginCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Executorplugin *ExecutorpluginRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executorplugin.Contract.ExecutorpluginTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Executorplugin *ExecutorpluginRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Executorplugin.Contract.ExecutorpluginTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Executorplugin *ExecutorpluginCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Executorplugin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Executorplugin *ExecutorpluginTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executorplugin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Executorplugin *ExecutorpluginTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Executorplugin.Contract.contract.Transact(opts, method, params...)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Executorplugin *ExecutorpluginCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Executorplugin.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Executorplugin *ExecutorpluginSession) AddressProvider() (common.Address, error) {
	return _Executorplugin.Contract.AddressProvider(&_Executorplugin.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Executorplugin *ExecutorpluginCallerSession) AddressProvider() (common.Address, error) {
	return _Executorplugin.Contract.AddressProvider(&_Executorplugin.CallOpts)
}

// AddressProviderTarget is a free data retrieval call binding the contract method 0x21b1e480.
//
// Solidity: function addressProviderTarget() view returns(address)
func (_Executorplugin *ExecutorpluginCaller) AddressProviderTarget(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Executorplugin.contract.Call(opts, &out, "addressProviderTarget")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProviderTarget is a free data retrieval call binding the contract method 0x21b1e480.
//
// Solidity: function addressProviderTarget() view returns(address)
func (_Executorplugin *ExecutorpluginSession) AddressProviderTarget() (common.Address, error) {
	return _Executorplugin.Contract.AddressProviderTarget(&_Executorplugin.CallOpts)
}

// AddressProviderTarget is a free data retrieval call binding the contract method 0x21b1e480.
//
// Solidity: function addressProviderTarget() view returns(address)
func (_Executorplugin *ExecutorpluginCallerSession) AddressProviderTarget() (common.Address, error) {
	return _Executorplugin.Contract.AddressProviderTarget(&_Executorplugin.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Executorplugin *ExecutorpluginCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Executorplugin.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Executorplugin *ExecutorpluginSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Executorplugin.Contract.Eip712Domain(&_Executorplugin.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Executorplugin *ExecutorpluginCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Executorplugin.Contract.Eip712Domain(&_Executorplugin.CallOpts)
}

// ExecutorNonce is a free data retrieval call binding the contract method 0x4611b4c7.
//
// Solidity: function executorNonce(address account, address executor) view returns(uint256 nonce)
func (_Executorplugin *ExecutorpluginCaller) ExecutorNonce(opts *bind.CallOpts, account common.Address, executor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Executorplugin.contract.Call(opts, &out, "executorNonce", account, executor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExecutorNonce is a free data retrieval call binding the contract method 0x4611b4c7.
//
// Solidity: function executorNonce(address account, address executor) view returns(uint256 nonce)
func (_Executorplugin *ExecutorpluginSession) ExecutorNonce(account common.Address, executor common.Address) (*big.Int, error) {
	return _Executorplugin.Contract.ExecutorNonce(&_Executorplugin.CallOpts, account, executor)
}

// ExecutorNonce is a free data retrieval call binding the contract method 0x4611b4c7.
//
// Solidity: function executorNonce(address account, address executor) view returns(uint256 nonce)
func (_Executorplugin *ExecutorpluginCallerSession) ExecutorNonce(account common.Address, executor common.Address) (*big.Int, error) {
	return _Executorplugin.Contract.ExecutorNonce(&_Executorplugin.CallOpts, account, executor)
}

// ExecutorRegistry is a free data retrieval call binding the contract method 0xb1cebbe0.
//
// Solidity: function executorRegistry() view returns(address)
func (_Executorplugin *ExecutorpluginCaller) ExecutorRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Executorplugin.contract.Call(opts, &out, "executorRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutorRegistry is a free data retrieval call binding the contract method 0xb1cebbe0.
//
// Solidity: function executorRegistry() view returns(address)
func (_Executorplugin *ExecutorpluginSession) ExecutorRegistry() (common.Address, error) {
	return _Executorplugin.Contract.ExecutorRegistry(&_Executorplugin.CallOpts)
}

// ExecutorRegistry is a free data retrieval call binding the contract method 0xb1cebbe0.
//
// Solidity: function executorRegistry() view returns(address)
func (_Executorplugin *ExecutorpluginCallerSession) ExecutorRegistry() (common.Address, error) {
	return _Executorplugin.Contract.ExecutorRegistry(&_Executorplugin.CallOpts)
}

// PolicyRegistry is a free data retrieval call binding the contract method 0x1c4dd7d0.
//
// Solidity: function policyRegistry() view returns(address)
func (_Executorplugin *ExecutorpluginCaller) PolicyRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Executorplugin.contract.Call(opts, &out, "policyRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolicyRegistry is a free data retrieval call binding the contract method 0x1c4dd7d0.
//
// Solidity: function policyRegistry() view returns(address)
func (_Executorplugin *ExecutorpluginSession) PolicyRegistry() (common.Address, error) {
	return _Executorplugin.Contract.PolicyRegistry(&_Executorplugin.CallOpts)
}

// PolicyRegistry is a free data retrieval call binding the contract method 0x1c4dd7d0.
//
// Solidity: function policyRegistry() view returns(address)
func (_Executorplugin *ExecutorpluginCallerSession) PolicyRegistry() (common.Address, error) {
	return _Executorplugin.Contract.PolicyRegistry(&_Executorplugin.CallOpts)
}

// WalletRegistry is a free data retrieval call binding the contract method 0xab7aa6ad.
//
// Solidity: function walletRegistry() view returns(address)
func (_Executorplugin *ExecutorpluginCaller) WalletRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Executorplugin.contract.Call(opts, &out, "walletRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WalletRegistry is a free data retrieval call binding the contract method 0xab7aa6ad.
//
// Solidity: function walletRegistry() view returns(address)
func (_Executorplugin *ExecutorpluginSession) WalletRegistry() (common.Address, error) {
	return _Executorplugin.Contract.WalletRegistry(&_Executorplugin.CallOpts)
}

// WalletRegistry is a free data retrieval call binding the contract method 0xab7aa6ad.
//
// Solidity: function walletRegistry() view returns(address)
func (_Executorplugin *ExecutorpluginCallerSession) WalletRegistry() (common.Address, error) {
	return _Executorplugin.Contract.WalletRegistry(&_Executorplugin.CallOpts)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x2bf4762b.
//
// Solidity: function executeTransaction(((uint8,address,uint256,bytes),address,address,bytes,bytes) execRequest) returns(bytes)
func (_Executorplugin *ExecutorpluginTransactor) ExecuteTransaction(opts *bind.TransactOpts, execRequest ExecutorPluginExecutionRequest) (*types.Transaction, error) {
	return _Executorplugin.contract.Transact(opts, "executeTransaction", execRequest)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x2bf4762b.
//
// Solidity: function executeTransaction(((uint8,address,uint256,bytes),address,address,bytes,bytes) execRequest) returns(bytes)
func (_Executorplugin *ExecutorpluginSession) ExecuteTransaction(execRequest ExecutorPluginExecutionRequest) (*types.Transaction, error) {
	return _Executorplugin.Contract.ExecuteTransaction(&_Executorplugin.TransactOpts, execRequest)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x2bf4762b.
//
// Solidity: function executeTransaction(((uint8,address,uint256,bytes),address,address,bytes,bytes) execRequest) returns(bytes)
func (_Executorplugin *ExecutorpluginTransactorSession) ExecuteTransaction(execRequest ExecutorPluginExecutionRequest) (*types.Transaction, error) {
	return _Executorplugin.Contract.ExecuteTransaction(&_Executorplugin.TransactOpts, execRequest)
}
