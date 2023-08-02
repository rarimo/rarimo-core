// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package state

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

// IStateGistProof is an auto generated low-level Go binding around an user-defined struct.
type IStateGistProof struct {
	Root         *big.Int
	Existence    bool
	Siblings     [64]*big.Int
	Index        *big.Int
	Value        *big.Int
	AuxExistence bool
	AuxIndex     *big.Int
	AuxValue     *big.Int
}

// IStateGistRootInfo is an auto generated low-level Go binding around an user-defined struct.
type IStateGistRootInfo struct {
	Root                *big.Int
	ReplacedByRoot      *big.Int
	CreatedAtTimestamp  *big.Int
	ReplacedAtTimestamp *big.Int
	CreatedAtBlock      *big.Int
	ReplacedAtBlock     *big.Int
}

// IStateStateInfo is an auto generated low-level Go binding around an user-defined struct.
type IStateStateInfo struct {
	Id                  *big.Int
	State               *big.Int
	ReplacedByState     *big.Int
	CreatedAtTimestamp  *big.Int
	ReplacedAtTimestamp *big.Int
	CreatedAtBlock      *big.Int
	ReplacedAtBlock     *big.Int
}

// StateV2HandlerMetaData contains all meta data concerning the StateV2Handler contract.
var StateV2HandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDefaultIdType\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getGISTProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[64]\",\"name\":\"siblings\",\"type\":\"uint256[64]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getGISTProofByBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[64]\",\"name\":\"siblings\",\"type\":\"uint256[64]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"getGISTProofByRoot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[64]\",\"name\":\"siblings\",\"type\":\"uint256[64]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getGISTProofByTime\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[64]\",\"name\":\"siblings\",\"type\":\"uint256[64]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGISTRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getGISTRootHistory\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistRootInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGISTRootHistoryLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"getGISTRootInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistRootInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getGISTRootInfoByBlock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistRootInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getGISTRootInfoByTime\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.GistRootInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getStateInfoById\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.StateInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"}],\"name\":\"getStateInfoByIdAndState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.StateInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getStateInfoHistoryById\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedByState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"replacedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIState.StateInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getStateInfoHistoryLengthById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"idExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStateTransitionVerifier\",\"name\":\"verifierContractAddr\",\"type\":\"address\"},{\"internalType\":\"bytes2\",\"name\":\"defaultIdType\",\"type\":\"bytes2\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes2\",\"name\":\"defaultIdType\",\"type\":\"bytes2\"}],\"name\":\"setDefaultIdType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newVerifierAddr\",\"type\":\"address\"}],\"name\":\"setVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"}],\"name\":\"stateExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newState\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOldStateGenesis\",\"type\":\"bool\"},{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"name\":\"transitState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldState\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newState\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOldStateGenesis\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"methodId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"methodParams\",\"type\":\"bytes\"}],\"name\":\"transitStateGeneric\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StateV2HandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use StateV2HandlerMetaData.ABI instead.
var StateV2HandlerABI = StateV2HandlerMetaData.ABI

// StateV2Handler is an auto generated Go binding around an Ethereum contract.
type StateV2Handler struct {
	StateV2HandlerCaller     // Read-only binding to the contract
	StateV2HandlerTransactor // Write-only binding to the contract
	StateV2HandlerFilterer   // Log filterer for contract events
}

// StateV2HandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StateV2HandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateV2HandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StateV2HandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateV2HandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateV2HandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateV2HandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateV2HandlerSession struct {
	Contract     *StateV2Handler   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateV2HandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateV2HandlerCallerSession struct {
	Contract *StateV2HandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StateV2HandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateV2HandlerTransactorSession struct {
	Contract     *StateV2HandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StateV2HandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StateV2HandlerRaw struct {
	Contract *StateV2Handler // Generic contract binding to access the raw methods on
}

// StateV2HandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateV2HandlerCallerRaw struct {
	Contract *StateV2HandlerCaller // Generic read-only contract binding to access the raw methods on
}

// StateV2HandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateV2HandlerTransactorRaw struct {
	Contract *StateV2HandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStateV2Handler creates a new instance of StateV2Handler, bound to a specific deployed contract.
func NewStateV2Handler(address common.Address, backend bind.ContractBackend) (*StateV2Handler, error) {
	contract, err := bindStateV2Handler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StateV2Handler{StateV2HandlerCaller: StateV2HandlerCaller{contract: contract}, StateV2HandlerTransactor: StateV2HandlerTransactor{contract: contract}, StateV2HandlerFilterer: StateV2HandlerFilterer{contract: contract}}, nil
}

// NewStateV2HandlerCaller creates a new read-only instance of StateV2Handler, bound to a specific deployed contract.
func NewStateV2HandlerCaller(address common.Address, caller bind.ContractCaller) (*StateV2HandlerCaller, error) {
	contract, err := bindStateV2Handler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateV2HandlerCaller{contract: contract}, nil
}

// NewStateV2HandlerTransactor creates a new write-only instance of StateV2Handler, bound to a specific deployed contract.
func NewStateV2HandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*StateV2HandlerTransactor, error) {
	contract, err := bindStateV2Handler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateV2HandlerTransactor{contract: contract}, nil
}

// NewStateV2HandlerFilterer creates a new log filterer instance of StateV2Handler, bound to a specific deployed contract.
func NewStateV2HandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*StateV2HandlerFilterer, error) {
	contract, err := bindStateV2Handler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateV2HandlerFilterer{contract: contract}, nil
}

// bindStateV2Handler binds a generic wrapper to an already deployed contract.
func bindStateV2Handler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StateV2HandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateV2Handler *StateV2HandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateV2Handler.Contract.StateV2HandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateV2Handler *StateV2HandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateV2Handler.Contract.StateV2HandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateV2Handler *StateV2HandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateV2Handler.Contract.StateV2HandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateV2Handler *StateV2HandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateV2Handler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateV2Handler *StateV2HandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateV2Handler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateV2Handler *StateV2HandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateV2Handler.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_StateV2Handler *StateV2HandlerCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StateV2Handler.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_StateV2Handler *StateV2HandlerSession) VERSION() (string, error) {
	return _StateV2Handler.Contract.VERSION(&_StateV2Handler.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_StateV2Handler *StateV2HandlerCallerSession) VERSION() (string, error) {
	return _StateV2Handler.Contract.VERSION(&_StateV2Handler.CallOpts)
}

// GetDefaultIdType is a free data retrieval call binding the contract method 0x133fe385.
//
// Solidity: function getDefaultIdType() view returns(bytes2)
func (_StateV2Handler *StateV2HandlerCaller) GetDefaultIdType(opts *bind.CallOpts) ([2]byte, error) {
	var out []interface{}
	err := _StateV2Handler.contract.Call(opts, &out, "getDefaultIdType")

	if err != nil {
		return *new([2]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([2]byte)).(*[2]byte)

	return out0, err

}

// GetDefaultIdType is a free data retrieval call binding the contract method 0x133fe385.
//
// Solidity: function getDefaultIdType() view returns(bytes2)
func (_StateV2Handler *StateV2HandlerSession) GetDefaultIdType() ([2]byte, error) {
	return _StateV2Handler.Contract.GetDefaultIdType(&_StateV2Handler.CallOpts)
}

// GetDefaultIdType is a free data retrieval call binding the contract method 0x133fe385.
//
// Solidity: function getDefaultIdType() view returns(bytes2)
func (_StateV2Handler *StateV2HandlerCallerSession) GetDefaultIdType() ([2]byte, error) {
	return _StateV2Handler.Contract.GetDefaultIdType(&_StateV2Handler.CallOpts)
}

// GetGISTProof is a free data retrieval call binding the contract method 0x3025bb8c.
//
// Solidity: function getGISTProof(uint256 id) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2Handler *StateV2HandlerCaller) GetGISTProof(opts *bind.CallOpts, id *big.Int) (IStateGistProof, error) {
	var out []interface{}
	err := _StateV2Handler.contract.Call(opts, &out, "getGISTProof", id)

	if err != nil {
		return *new(IStateGistProof), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistProof)).(*IStateGistProof)

	return out0, err

}

// GetGISTProof is a free data retrieval call binding the contract method 0x3025bb8c.
//
// Solidity: function getGISTProof(uint256 id) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2Handler *StateV2HandlerSession) GetGISTProof(id *big.Int) (IStateGistProof, error) {
	return _StateV2Handler.Contract.GetGISTProof(&_StateV2Handler.CallOpts, id)
}

// GetGISTProof is a free data retrieval call binding the contract method 0x3025bb8c.
//
// Solidity: function getGISTProof(uint256 id) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2Handler *StateV2HandlerCallerSession) GetGISTProof(id *big.Int) (IStateGistProof, error) {
	return _StateV2Handler.Contract.GetGISTProof(&_StateV2Handler.CallOpts, id)
}

// GetGISTProofByBlock is a free data retrieval call binding the contract method 0x046ff140.
//
// Solidity: function getGISTProofByBlock(uint256 id, uint256 blockNumber) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2Handler *StateV2HandlerCaller) GetGISTProofByBlock(opts *bind.CallOpts, id *big.Int, blockNumber *big.Int) (IStateGistProof, error) {
	var out []interface{}
	err := _StateV2Handler.contract.Call(opts, &out, "getGISTProofByBlock", id, blockNumber)

	if err != nil {
		return *new(IStateGistProof), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistProof)).(*IStateGistProof)

	return out0, err

}

// GetGISTProofByBlock is a free data retrieval call binding the contract method 0x046ff140.
//
// Solidity: function getGISTProofByBlock(uint256 id, uint256 blockNumber) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2Handler *StateV2HandlerSession) GetGISTProofByBlock(id *big.Int, blockNumber *big.Int) (IStateGistProof, error) {
	return _StateV2Handler.Contract.GetGISTProofByBlock(&_StateV2Handler.CallOpts, id, blockNumber)
}

// GetGISTProofByBlock is a free data retrieval call binding the contract method 0x046ff140.
//
// Solidity: function getGISTProofByBlock(uint256 id, uint256 blockNumber) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2Handler *StateV2HandlerCallerSession) GetGISTProofByBlock(id *big.Int, blockNumber *big.Int) (IStateGistProof, error) {
	return _StateV2Handler.Contract.GetGISTProofByBlock(&_StateV2Handler.CallOpts, id, blockNumber)
}

// GetGISTProofByRoot is a free data retrieval call binding the contract method 0xe12a36c0.
//
// Solidity: function getGISTProofByRoot(uint256 id, uint256 root) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2Handler *StateV2HandlerCaller) GetGISTProofByRoot(opts *bind.CallOpts, id *big.Int, root *big.Int) (IStateGistProof, error) {
	var out []interface{}
	err := _StateV2Handler.contract.Call(opts, &out, "getGISTProofByRoot", id, root)

	if err != nil {
		return *new(IStateGistProof), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistProof)).(*IStateGistProof)

	return out0, err

}

// GetGISTProofByRoot is a free data retrieval call binding the contract method 0xe12a36c0.
//
// Solidity: function getGISTProofByRoot(uint256 id, uint256 root) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2Handler *StateV2HandlerSession) GetGISTProofByRoot(id *big.Int, root *big.Int) (IStateGistProof, error) {
	return _StateV2Handler.Contract.GetGISTProofByRoot(&_StateV2Handler.CallOpts, id, root)
}

// GetGISTProofByRoot is a free data retrieval call binding the contract method 0xe12a36c0.
//
// Solidity: function getGISTProofByRoot(uint256 id, uint256 root) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2Handler *StateV2HandlerCallerSession) GetGISTProofByRoot(id *big.Int, root *big.Int) (IStateGistProof, error) {
	return _StateV2Handler.Contract.GetGISTProofByRoot(&_StateV2Handler.CallOpts, id, root)
}

// GetGISTProofByTime is a free data retrieval call binding the contract method 0xd51afebf.
//
// Solidity: function getGISTProofByTime(uint256 id, uint256 timestamp) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2Handler *StateV2HandlerCaller) GetGISTProofByTime(opts *bind.CallOpts, id *big.Int, timestamp *big.Int) (IStateGistProof, error) {
	var out []interface{}
	err := _StateV2Handler.contract.Call(opts, &out, "getGISTProofByTime", id, timestamp)

	if err != nil {
		return *new(IStateGistProof), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistProof)).(*IStateGistProof)

	return out0, err

}

// GetGISTProofByTime is a free data retrieval call binding the contract method 0xd51afebf.
//
// Solidity: function getGISTProofByTime(uint256 id, uint256 timestamp) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2Handler *StateV2HandlerSession) GetGISTProofByTime(id *big.Int, timestamp *big.Int) (IStateGistProof, error) {
	return _StateV2Handler.Contract.GetGISTProofByTime(&_StateV2Handler.CallOpts, id, timestamp)
}

// GetGISTProofByTime is a free data retrieval call binding the contract method 0xd51afebf.
//
// Solidity: function getGISTProofByTime(uint256 id, uint256 timestamp) view returns((uint256,bool,uint256[64],uint256,uint256,bool,uint256,uint256))
func (_StateV2Handler *StateV2HandlerCallerSession) GetGISTProofByTime(id *big.Int, timestamp *big.Int) (IStateGistProof, error) {
	return _StateV2Handler.Contract.GetGISTProofByTime(&_StateV2Handler.CallOpts, id, timestamp)
}

// GetGISTRoot is a free data retrieval call binding the contract method 0x2439e3a6.
//
// Solidity: function getGISTRoot() view returns(uint256)
func (_StateV2Handler *StateV2HandlerCaller) GetGISTRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateV2Handler.contract.Call(opts, &out, "getGISTRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGISTRoot is a free data retrieval call binding the contract method 0x2439e3a6.
//
// Solidity: function getGISTRoot() view returns(uint256)
func (_StateV2Handler *StateV2HandlerSession) GetGISTRoot() (*big.Int, error) {
	return _StateV2Handler.Contract.GetGISTRoot(&_StateV2Handler.CallOpts)
}

// GetGISTRoot is a free data retrieval call binding the contract method 0x2439e3a6.
//
// Solidity: function getGISTRoot() view returns(uint256)
func (_StateV2Handler *StateV2HandlerCallerSession) GetGISTRoot() (*big.Int, error) {
	return _StateV2Handler.Contract.GetGISTRoot(&_StateV2Handler.CallOpts)
}

// GetGISTRootHistory is a free data retrieval call binding the contract method 0x2f7670e4.
//
// Solidity: function getGISTRootHistory(uint256 start, uint256 length) view returns((uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_StateV2Handler *StateV2HandlerCaller) GetGISTRootHistory(opts *bind.CallOpts, start *big.Int, length *big.Int) ([]IStateGistRootInfo, error) {
	var out []interface{}
	err := _StateV2Handler.contract.Call(opts, &out, "getGISTRootHistory", start, length)

	if err != nil {
		return *new([]IStateGistRootInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStateGistRootInfo)).(*[]IStateGistRootInfo)

	return out0, err

}

// GetGISTRootHistory is a free data retrieval call binding the contract method 0x2f7670e4.
//
// Solidity: function getGISTRootHistory(uint256 start, uint256 length) view returns((uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_StateV2Handler *StateV2HandlerSession) GetGISTRootHistory(start *big.Int, length *big.Int) ([]IStateGistRootInfo, error) {
	return _StateV2Handler.Contract.GetGISTRootHistory(&_StateV2Handler.CallOpts, start, length)
}

// GetGISTRootHistory is a free data retrieval call binding the contract method 0x2f7670e4.
//
// Solidity: function getGISTRootHistory(uint256 start, uint256 length) view returns((uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_StateV2Handler *StateV2HandlerCallerSession) GetGISTRootHistory(start *big.Int, length *big.Int) ([]IStateGistRootInfo, error) {
	return _StateV2Handler.Contract.GetGISTRootHistory(&_StateV2Handler.CallOpts, start, length)
}

// GetGISTRootHistoryLength is a free data retrieval call binding the contract method 0xdccbd57a.
//
// Solidity: function getGISTRootHistoryLength() view returns(uint256)
func (_StateV2Handler *StateV2HandlerCaller) GetGISTRootHistoryLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateV2Handler.contract.Call(opts, &out, "getGISTRootHistoryLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGISTRootHistoryLength is a free data retrieval call binding the contract method 0xdccbd57a.
//
// Solidity: function getGISTRootHistoryLength() view returns(uint256)
func (_StateV2Handler *StateV2HandlerSession) GetGISTRootHistoryLength() (*big.Int, error) {
	return _StateV2Handler.Contract.GetGISTRootHistoryLength(&_StateV2Handler.CallOpts)
}

// GetGISTRootHistoryLength is a free data retrieval call binding the contract method 0xdccbd57a.
//
// Solidity: function getGISTRootHistoryLength() view returns(uint256)
func (_StateV2Handler *StateV2HandlerCallerSession) GetGISTRootHistoryLength() (*big.Int, error) {
	return _StateV2Handler.Contract.GetGISTRootHistoryLength(&_StateV2Handler.CallOpts)
}

// GetGISTRootInfo is a free data retrieval call binding the contract method 0x7c1a66de.
//
// Solidity: function getGISTRootInfo(uint256 root) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2Handler *StateV2HandlerCaller) GetGISTRootInfo(opts *bind.CallOpts, root *big.Int) (IStateGistRootInfo, error) {
	var out []interface{}
	err := _StateV2Handler.contract.Call(opts, &out, "getGISTRootInfo", root)

	if err != nil {
		return *new(IStateGistRootInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistRootInfo)).(*IStateGistRootInfo)

	return out0, err

}

// GetGISTRootInfo is a free data retrieval call binding the contract method 0x7c1a66de.
//
// Solidity: function getGISTRootInfo(uint256 root) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2Handler *StateV2HandlerSession) GetGISTRootInfo(root *big.Int) (IStateGistRootInfo, error) {
	return _StateV2Handler.Contract.GetGISTRootInfo(&_StateV2Handler.CallOpts, root)
}

// GetGISTRootInfo is a free data retrieval call binding the contract method 0x7c1a66de.
//
// Solidity: function getGISTRootInfo(uint256 root) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2Handler *StateV2HandlerCallerSession) GetGISTRootInfo(root *big.Int) (IStateGistRootInfo, error) {
	return _StateV2Handler.Contract.GetGISTRootInfo(&_StateV2Handler.CallOpts, root)
}

// GetGISTRootInfoByBlock is a free data retrieval call binding the contract method 0x5845e530.
//
// Solidity: function getGISTRootInfoByBlock(uint256 blockNumber) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2Handler *StateV2HandlerCaller) GetGISTRootInfoByBlock(opts *bind.CallOpts, blockNumber *big.Int) (IStateGistRootInfo, error) {
	var out []interface{}
	err := _StateV2Handler.contract.Call(opts, &out, "getGISTRootInfoByBlock", blockNumber)

	if err != nil {
		return *new(IStateGistRootInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistRootInfo)).(*IStateGistRootInfo)

	return out0, err

}

// GetGISTRootInfoByBlock is a free data retrieval call binding the contract method 0x5845e530.
//
// Solidity: function getGISTRootInfoByBlock(uint256 blockNumber) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2Handler *StateV2HandlerSession) GetGISTRootInfoByBlock(blockNumber *big.Int) (IStateGistRootInfo, error) {
	return _StateV2Handler.Contract.GetGISTRootInfoByBlock(&_StateV2Handler.CallOpts, blockNumber)
}

// GetGISTRootInfoByBlock is a free data retrieval call binding the contract method 0x5845e530.
//
// Solidity: function getGISTRootInfoByBlock(uint256 blockNumber) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2Handler *StateV2HandlerCallerSession) GetGISTRootInfoByBlock(blockNumber *big.Int) (IStateGistRootInfo, error) {
	return _StateV2Handler.Contract.GetGISTRootInfoByBlock(&_StateV2Handler.CallOpts, blockNumber)
}

// GetGISTRootInfoByTime is a free data retrieval call binding the contract method 0x0ef6e65b.
//
// Solidity: function getGISTRootInfoByTime(uint256 timestamp) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2Handler *StateV2HandlerCaller) GetGISTRootInfoByTime(opts *bind.CallOpts, timestamp *big.Int) (IStateGistRootInfo, error) {
	var out []interface{}
	err := _StateV2Handler.contract.Call(opts, &out, "getGISTRootInfoByTime", timestamp)

	if err != nil {
		return *new(IStateGistRootInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateGistRootInfo)).(*IStateGistRootInfo)

	return out0, err

}

// GetGISTRootInfoByTime is a free data retrieval call binding the contract method 0x0ef6e65b.
//
// Solidity: function getGISTRootInfoByTime(uint256 timestamp) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2Handler *StateV2HandlerSession) GetGISTRootInfoByTime(timestamp *big.Int) (IStateGistRootInfo, error) {
	return _StateV2Handler.Contract.GetGISTRootInfoByTime(&_StateV2Handler.CallOpts, timestamp)
}

// GetGISTRootInfoByTime is a free data retrieval call binding the contract method 0x0ef6e65b.
//
// Solidity: function getGISTRootInfoByTime(uint256 timestamp) view returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2Handler *StateV2HandlerCallerSession) GetGISTRootInfoByTime(timestamp *big.Int) (IStateGistRootInfo, error) {
	return _StateV2Handler.Contract.GetGISTRootInfoByTime(&_StateV2Handler.CallOpts, timestamp)
}

// GetStateInfoById is a free data retrieval call binding the contract method 0xb4bdea55.
//
// Solidity: function getStateInfoById(uint256 id) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2Handler *StateV2HandlerCaller) GetStateInfoById(opts *bind.CallOpts, id *big.Int) (IStateStateInfo, error) {
	var out []interface{}
	err := _StateV2Handler.contract.Call(opts, &out, "getStateInfoById", id)

	if err != nil {
		return *new(IStateStateInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateStateInfo)).(*IStateStateInfo)

	return out0, err

}

// GetStateInfoById is a free data retrieval call binding the contract method 0xb4bdea55.
//
// Solidity: function getStateInfoById(uint256 id) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2Handler *StateV2HandlerSession) GetStateInfoById(id *big.Int) (IStateStateInfo, error) {
	return _StateV2Handler.Contract.GetStateInfoById(&_StateV2Handler.CallOpts, id)
}

// GetStateInfoById is a free data retrieval call binding the contract method 0xb4bdea55.
//
// Solidity: function getStateInfoById(uint256 id) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2Handler *StateV2HandlerCallerSession) GetStateInfoById(id *big.Int) (IStateStateInfo, error) {
	return _StateV2Handler.Contract.GetStateInfoById(&_StateV2Handler.CallOpts, id)
}

// GetStateInfoByIdAndState is a free data retrieval call binding the contract method 0x53c87312.
//
// Solidity: function getStateInfoByIdAndState(uint256 id, uint256 state) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2Handler *StateV2HandlerCaller) GetStateInfoByIdAndState(opts *bind.CallOpts, id *big.Int, state *big.Int) (IStateStateInfo, error) {
	var out []interface{}
	err := _StateV2Handler.contract.Call(opts, &out, "getStateInfoByIdAndState", id, state)

	if err != nil {
		return *new(IStateStateInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IStateStateInfo)).(*IStateStateInfo)

	return out0, err

}

// GetStateInfoByIdAndState is a free data retrieval call binding the contract method 0x53c87312.
//
// Solidity: function getStateInfoByIdAndState(uint256 id, uint256 state) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2Handler *StateV2HandlerSession) GetStateInfoByIdAndState(id *big.Int, state *big.Int) (IStateStateInfo, error) {
	return _StateV2Handler.Contract.GetStateInfoByIdAndState(&_StateV2Handler.CallOpts, id, state)
}

// GetStateInfoByIdAndState is a free data retrieval call binding the contract method 0x53c87312.
//
// Solidity: function getStateInfoByIdAndState(uint256 id, uint256 state) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_StateV2Handler *StateV2HandlerCallerSession) GetStateInfoByIdAndState(id *big.Int, state *big.Int) (IStateStateInfo, error) {
	return _StateV2Handler.Contract.GetStateInfoByIdAndState(&_StateV2Handler.CallOpts, id, state)
}

// GetStateInfoHistoryById is a free data retrieval call binding the contract method 0xe99858fe.
//
// Solidity: function getStateInfoHistoryById(uint256 id, uint256 startIndex, uint256 length) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_StateV2Handler *StateV2HandlerCaller) GetStateInfoHistoryById(opts *bind.CallOpts, id *big.Int, startIndex *big.Int, length *big.Int) ([]IStateStateInfo, error) {
	var out []interface{}
	err := _StateV2Handler.contract.Call(opts, &out, "getStateInfoHistoryById", id, startIndex, length)

	if err != nil {
		return *new([]IStateStateInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStateStateInfo)).(*[]IStateStateInfo)

	return out0, err

}

// GetStateInfoHistoryById is a free data retrieval call binding the contract method 0xe99858fe.
//
// Solidity: function getStateInfoHistoryById(uint256 id, uint256 startIndex, uint256 length) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_StateV2Handler *StateV2HandlerSession) GetStateInfoHistoryById(id *big.Int, startIndex *big.Int, length *big.Int) ([]IStateStateInfo, error) {
	return _StateV2Handler.Contract.GetStateInfoHistoryById(&_StateV2Handler.CallOpts, id, startIndex, length)
}

// GetStateInfoHistoryById is a free data retrieval call binding the contract method 0xe99858fe.
//
// Solidity: function getStateInfoHistoryById(uint256 id, uint256 startIndex, uint256 length) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256)[])
func (_StateV2Handler *StateV2HandlerCallerSession) GetStateInfoHistoryById(id *big.Int, startIndex *big.Int, length *big.Int) ([]IStateStateInfo, error) {
	return _StateV2Handler.Contract.GetStateInfoHistoryById(&_StateV2Handler.CallOpts, id, startIndex, length)
}

// GetStateInfoHistoryLengthById is a free data retrieval call binding the contract method 0x676d5b5a.
//
// Solidity: function getStateInfoHistoryLengthById(uint256 id) view returns(uint256)
func (_StateV2Handler *StateV2HandlerCaller) GetStateInfoHistoryLengthById(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StateV2Handler.contract.Call(opts, &out, "getStateInfoHistoryLengthById", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStateInfoHistoryLengthById is a free data retrieval call binding the contract method 0x676d5b5a.
//
// Solidity: function getStateInfoHistoryLengthById(uint256 id) view returns(uint256)
func (_StateV2Handler *StateV2HandlerSession) GetStateInfoHistoryLengthById(id *big.Int) (*big.Int, error) {
	return _StateV2Handler.Contract.GetStateInfoHistoryLengthById(&_StateV2Handler.CallOpts, id)
}

// GetStateInfoHistoryLengthById is a free data retrieval call binding the contract method 0x676d5b5a.
//
// Solidity: function getStateInfoHistoryLengthById(uint256 id) view returns(uint256)
func (_StateV2Handler *StateV2HandlerCallerSession) GetStateInfoHistoryLengthById(id *big.Int) (*big.Int, error) {
	return _StateV2Handler.Contract.GetStateInfoHistoryLengthById(&_StateV2Handler.CallOpts, id)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_StateV2Handler *StateV2HandlerCaller) GetVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StateV2Handler.contract.Call(opts, &out, "getVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_StateV2Handler *StateV2HandlerSession) GetVerifier() (common.Address, error) {
	return _StateV2Handler.Contract.GetVerifier(&_StateV2Handler.CallOpts)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_StateV2Handler *StateV2HandlerCallerSession) GetVerifier() (common.Address, error) {
	return _StateV2Handler.Contract.GetVerifier(&_StateV2Handler.CallOpts)
}

// IdExists is a free data retrieval call binding the contract method 0x0b8a295a.
//
// Solidity: function idExists(uint256 id) view returns(bool)
func (_StateV2Handler *StateV2HandlerCaller) IdExists(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _StateV2Handler.contract.Call(opts, &out, "idExists", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IdExists is a free data retrieval call binding the contract method 0x0b8a295a.
//
// Solidity: function idExists(uint256 id) view returns(bool)
func (_StateV2Handler *StateV2HandlerSession) IdExists(id *big.Int) (bool, error) {
	return _StateV2Handler.Contract.IdExists(&_StateV2Handler.CallOpts, id)
}

// IdExists is a free data retrieval call binding the contract method 0x0b8a295a.
//
// Solidity: function idExists(uint256 id) view returns(bool)
func (_StateV2Handler *StateV2HandlerCallerSession) IdExists(id *big.Int) (bool, error) {
	return _StateV2Handler.Contract.IdExists(&_StateV2Handler.CallOpts, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StateV2Handler *StateV2HandlerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StateV2Handler.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StateV2Handler *StateV2HandlerSession) Owner() (common.Address, error) {
	return _StateV2Handler.Contract.Owner(&_StateV2Handler.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StateV2Handler *StateV2HandlerCallerSession) Owner() (common.Address, error) {
	return _StateV2Handler.Contract.Owner(&_StateV2Handler.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_StateV2Handler *StateV2HandlerCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StateV2Handler.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_StateV2Handler *StateV2HandlerSession) PendingOwner() (common.Address, error) {
	return _StateV2Handler.Contract.PendingOwner(&_StateV2Handler.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_StateV2Handler *StateV2HandlerCallerSession) PendingOwner() (common.Address, error) {
	return _StateV2Handler.Contract.PendingOwner(&_StateV2Handler.CallOpts)
}

// StateExists is a free data retrieval call binding the contract method 0x233a4d23.
//
// Solidity: function stateExists(uint256 id, uint256 state) view returns(bool)
func (_StateV2Handler *StateV2HandlerCaller) StateExists(opts *bind.CallOpts, id *big.Int, state *big.Int) (bool, error) {
	var out []interface{}
	err := _StateV2Handler.contract.Call(opts, &out, "stateExists", id, state)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StateExists is a free data retrieval call binding the contract method 0x233a4d23.
//
// Solidity: function stateExists(uint256 id, uint256 state) view returns(bool)
func (_StateV2Handler *StateV2HandlerSession) StateExists(id *big.Int, state *big.Int) (bool, error) {
	return _StateV2Handler.Contract.StateExists(&_StateV2Handler.CallOpts, id, state)
}

// StateExists is a free data retrieval call binding the contract method 0x233a4d23.
//
// Solidity: function stateExists(uint256 id, uint256 state) view returns(bool)
func (_StateV2Handler *StateV2HandlerCallerSession) StateExists(id *big.Int, state *big.Int) (bool, error) {
	return _StateV2Handler.Contract.StateExists(&_StateV2Handler.CallOpts, id, state)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_StateV2Handler *StateV2HandlerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateV2Handler.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_StateV2Handler *StateV2HandlerSession) AcceptOwnership() (*types.Transaction, error) {
	return _StateV2Handler.Contract.AcceptOwnership(&_StateV2Handler.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_StateV2Handler *StateV2HandlerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _StateV2Handler.Contract.AcceptOwnership(&_StateV2Handler.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x406e8b4c.
//
// Solidity: function initialize(address verifierContractAddr, bytes2 defaultIdType) returns()
func (_StateV2Handler *StateV2HandlerTransactor) Initialize(opts *bind.TransactOpts, verifierContractAddr common.Address, defaultIdType [2]byte) (*types.Transaction, error) {
	return _StateV2Handler.contract.Transact(opts, "initialize", verifierContractAddr, defaultIdType)
}

// Initialize is a paid mutator transaction binding the contract method 0x406e8b4c.
//
// Solidity: function initialize(address verifierContractAddr, bytes2 defaultIdType) returns()
func (_StateV2Handler *StateV2HandlerSession) Initialize(verifierContractAddr common.Address, defaultIdType [2]byte) (*types.Transaction, error) {
	return _StateV2Handler.Contract.Initialize(&_StateV2Handler.TransactOpts, verifierContractAddr, defaultIdType)
}

// Initialize is a paid mutator transaction binding the contract method 0x406e8b4c.
//
// Solidity: function initialize(address verifierContractAddr, bytes2 defaultIdType) returns()
func (_StateV2Handler *StateV2HandlerTransactorSession) Initialize(verifierContractAddr common.Address, defaultIdType [2]byte) (*types.Transaction, error) {
	return _StateV2Handler.Contract.Initialize(&_StateV2Handler.TransactOpts, verifierContractAddr, defaultIdType)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StateV2Handler *StateV2HandlerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateV2Handler.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StateV2Handler *StateV2HandlerSession) RenounceOwnership() (*types.Transaction, error) {
	return _StateV2Handler.Contract.RenounceOwnership(&_StateV2Handler.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StateV2Handler *StateV2HandlerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StateV2Handler.Contract.RenounceOwnership(&_StateV2Handler.TransactOpts)
}

// SetDefaultIdType is a paid mutator transaction binding the contract method 0xa7c57eeb.
//
// Solidity: function setDefaultIdType(bytes2 defaultIdType) returns()
func (_StateV2Handler *StateV2HandlerTransactor) SetDefaultIdType(opts *bind.TransactOpts, defaultIdType [2]byte) (*types.Transaction, error) {
	return _StateV2Handler.contract.Transact(opts, "setDefaultIdType", defaultIdType)
}

// SetDefaultIdType is a paid mutator transaction binding the contract method 0xa7c57eeb.
//
// Solidity: function setDefaultIdType(bytes2 defaultIdType) returns()
func (_StateV2Handler *StateV2HandlerSession) SetDefaultIdType(defaultIdType [2]byte) (*types.Transaction, error) {
	return _StateV2Handler.Contract.SetDefaultIdType(&_StateV2Handler.TransactOpts, defaultIdType)
}

// SetDefaultIdType is a paid mutator transaction binding the contract method 0xa7c57eeb.
//
// Solidity: function setDefaultIdType(bytes2 defaultIdType) returns()
func (_StateV2Handler *StateV2HandlerTransactorSession) SetDefaultIdType(defaultIdType [2]byte) (*types.Transaction, error) {
	return _StateV2Handler.Contract.SetDefaultIdType(&_StateV2Handler.TransactOpts, defaultIdType)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address newVerifierAddr) returns()
func (_StateV2Handler *StateV2HandlerTransactor) SetVerifier(opts *bind.TransactOpts, newVerifierAddr common.Address) (*types.Transaction, error) {
	return _StateV2Handler.contract.Transact(opts, "setVerifier", newVerifierAddr)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address newVerifierAddr) returns()
func (_StateV2Handler *StateV2HandlerSession) SetVerifier(newVerifierAddr common.Address) (*types.Transaction, error) {
	return _StateV2Handler.Contract.SetVerifier(&_StateV2Handler.TransactOpts, newVerifierAddr)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address newVerifierAddr) returns()
func (_StateV2Handler *StateV2HandlerTransactorSession) SetVerifier(newVerifierAddr common.Address) (*types.Transaction, error) {
	return _StateV2Handler.Contract.SetVerifier(&_StateV2Handler.TransactOpts, newVerifierAddr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StateV2Handler *StateV2HandlerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StateV2Handler.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StateV2Handler *StateV2HandlerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StateV2Handler.Contract.TransferOwnership(&_StateV2Handler.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StateV2Handler *StateV2HandlerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StateV2Handler.Contract.TransferOwnership(&_StateV2Handler.TransactOpts, newOwner)
}

// TransitState is a paid mutator transaction binding the contract method 0x28f88a65.
//
// Solidity: function transitState(uint256 id, uint256 oldState, uint256 newState, bool isOldStateGenesis, uint256[2] a, uint256[2][2] b, uint256[2] c) returns()
func (_StateV2Handler *StateV2HandlerTransactor) TransitState(opts *bind.TransactOpts, id *big.Int, oldState *big.Int, newState *big.Int, isOldStateGenesis bool, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (*types.Transaction, error) {
	return _StateV2Handler.contract.Transact(opts, "transitState", id, oldState, newState, isOldStateGenesis, a, b, c)
}

// TransitState is a paid mutator transaction binding the contract method 0x28f88a65.
//
// Solidity: function transitState(uint256 id, uint256 oldState, uint256 newState, bool isOldStateGenesis, uint256[2] a, uint256[2][2] b, uint256[2] c) returns()
func (_StateV2Handler *StateV2HandlerSession) TransitState(id *big.Int, oldState *big.Int, newState *big.Int, isOldStateGenesis bool, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (*types.Transaction, error) {
	return _StateV2Handler.Contract.TransitState(&_StateV2Handler.TransactOpts, id, oldState, newState, isOldStateGenesis, a, b, c)
}

// TransitState is a paid mutator transaction binding the contract method 0x28f88a65.
//
// Solidity: function transitState(uint256 id, uint256 oldState, uint256 newState, bool isOldStateGenesis, uint256[2] a, uint256[2][2] b, uint256[2] c) returns()
func (_StateV2Handler *StateV2HandlerTransactorSession) TransitState(id *big.Int, oldState *big.Int, newState *big.Int, isOldStateGenesis bool, a [2]*big.Int, b [2][2]*big.Int, c [2]*big.Int) (*types.Transaction, error) {
	return _StateV2Handler.Contract.TransitState(&_StateV2Handler.TransactOpts, id, oldState, newState, isOldStateGenesis, a, b, c)
}

// TransitStateGeneric is a paid mutator transaction binding the contract method 0x7d0dcba5.
//
// Solidity: function transitStateGeneric(uint256 id, uint256 oldState, uint256 newState, bool isOldStateGenesis, uint256 methodId, bytes methodParams) returns()
func (_StateV2Handler *StateV2HandlerTransactor) TransitStateGeneric(opts *bind.TransactOpts, id *big.Int, oldState *big.Int, newState *big.Int, isOldStateGenesis bool, methodId *big.Int, methodParams []byte) (*types.Transaction, error) {
	return _StateV2Handler.contract.Transact(opts, "transitStateGeneric", id, oldState, newState, isOldStateGenesis, methodId, methodParams)
}

// TransitStateGeneric is a paid mutator transaction binding the contract method 0x7d0dcba5.
//
// Solidity: function transitStateGeneric(uint256 id, uint256 oldState, uint256 newState, bool isOldStateGenesis, uint256 methodId, bytes methodParams) returns()
func (_StateV2Handler *StateV2HandlerSession) TransitStateGeneric(id *big.Int, oldState *big.Int, newState *big.Int, isOldStateGenesis bool, methodId *big.Int, methodParams []byte) (*types.Transaction, error) {
	return _StateV2Handler.Contract.TransitStateGeneric(&_StateV2Handler.TransactOpts, id, oldState, newState, isOldStateGenesis, methodId, methodParams)
}

// TransitStateGeneric is a paid mutator transaction binding the contract method 0x7d0dcba5.
//
// Solidity: function transitStateGeneric(uint256 id, uint256 oldState, uint256 newState, bool isOldStateGenesis, uint256 methodId, bytes methodParams) returns()
func (_StateV2Handler *StateV2HandlerTransactorSession) TransitStateGeneric(id *big.Int, oldState *big.Int, newState *big.Int, isOldStateGenesis bool, methodId *big.Int, methodParams []byte) (*types.Transaction, error) {
	return _StateV2Handler.Contract.TransitStateGeneric(&_StateV2Handler.TransactOpts, id, oldState, newState, isOldStateGenesis, methodId, methodParams)
}

// StateV2HandlerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StateV2Handler contract.
type StateV2HandlerInitializedIterator struct {
	Event *StateV2HandlerInitialized // Event containing the contract specifics and raw log

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
func (it *StateV2HandlerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateV2HandlerInitialized)
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
		it.Event = new(StateV2HandlerInitialized)
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
func (it *StateV2HandlerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateV2HandlerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateV2HandlerInitialized represents a Initialized event raised by the StateV2Handler contract.
type StateV2HandlerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StateV2Handler *StateV2HandlerFilterer) FilterInitialized(opts *bind.FilterOpts) (*StateV2HandlerInitializedIterator, error) {

	logs, sub, err := _StateV2Handler.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StateV2HandlerInitializedIterator{contract: _StateV2Handler.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StateV2Handler *StateV2HandlerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StateV2HandlerInitialized) (event.Subscription, error) {

	logs, sub, err := _StateV2Handler.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateV2HandlerInitialized)
				if err := _StateV2Handler.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StateV2Handler *StateV2HandlerFilterer) ParseInitialized(log types.Log) (*StateV2HandlerInitialized, error) {
	event := new(StateV2HandlerInitialized)
	if err := _StateV2Handler.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StateV2HandlerOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the StateV2Handler contract.
type StateV2HandlerOwnershipTransferStartedIterator struct {
	Event *StateV2HandlerOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *StateV2HandlerOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateV2HandlerOwnershipTransferStarted)
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
		it.Event = new(StateV2HandlerOwnershipTransferStarted)
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
func (it *StateV2HandlerOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateV2HandlerOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateV2HandlerOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the StateV2Handler contract.
type StateV2HandlerOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_StateV2Handler *StateV2HandlerFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StateV2HandlerOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StateV2Handler.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StateV2HandlerOwnershipTransferStartedIterator{contract: _StateV2Handler.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_StateV2Handler *StateV2HandlerFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *StateV2HandlerOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StateV2Handler.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateV2HandlerOwnershipTransferStarted)
				if err := _StateV2Handler.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_StateV2Handler *StateV2HandlerFilterer) ParseOwnershipTransferStarted(log types.Log) (*StateV2HandlerOwnershipTransferStarted, error) {
	event := new(StateV2HandlerOwnershipTransferStarted)
	if err := _StateV2Handler.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StateV2HandlerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StateV2Handler contract.
type StateV2HandlerOwnershipTransferredIterator struct {
	Event *StateV2HandlerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StateV2HandlerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateV2HandlerOwnershipTransferred)
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
		it.Event = new(StateV2HandlerOwnershipTransferred)
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
func (it *StateV2HandlerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateV2HandlerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateV2HandlerOwnershipTransferred represents a OwnershipTransferred event raised by the StateV2Handler contract.
type StateV2HandlerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StateV2Handler *StateV2HandlerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StateV2HandlerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StateV2Handler.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StateV2HandlerOwnershipTransferredIterator{contract: _StateV2Handler.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StateV2Handler *StateV2HandlerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StateV2HandlerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StateV2Handler.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateV2HandlerOwnershipTransferred)
				if err := _StateV2Handler.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StateV2Handler *StateV2HandlerFilterer) ParseOwnershipTransferred(log types.Log) (*StateV2HandlerOwnershipTransferred, error) {
	event := new(StateV2HandlerOwnershipTransferred)
	if err := _StateV2Handler.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
