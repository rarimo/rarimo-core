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

// SparseMerkleTreeNode is an auto generated low-level Go binding around an user-defined struct.
type SparseMerkleTreeNode struct {
	NodeType   uint8
	ChildLeft  uint64
	ChildRight uint64
	NodeHash   [32]byte
	Key        [32]byte
	Value      [32]byte
}

// SparseMerkleTreeProof is an auto generated low-level Go binding around an user-defined struct.
type SparseMerkleTreeProof struct {
	Root         [32]byte
	Siblings     [][32]byte
	Existence    bool
	Key          [32]byte
	Value        [32]byte
	AuxExistence bool
	AuxKey       [32]byte
	AuxValue     [32]byte
}

// PoseidonSMTMetaData contains all meta data concerning the PoseidonSMT contract.
var PoseidonSMTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"RootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAGIC_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"P\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROOT_VALIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"chainName_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"stateKeeper_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"treeHeight_\",\"type\":\"uint256\"}],\"name\":\"__PoseidonSMT_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyOfElement_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"element_\",\"type\":\"bytes32\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"newSignerPubKey_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"changeSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key_\",\"type\":\"bytes32\"}],\"name\":\"getNodeByKey\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSparseMerkleTree.NodeType\",\"name\":\"nodeType\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"childLeft\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"childRight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structSparseMerkleTree.Node\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"methodId_\",\"type\":\"uint8\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key_\",\"type\":\"bytes32\"}],\"name\":\"getProof\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"auxKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auxValue\",\"type\":\"bytes32\"}],\"internalType\":\"structSparseMerkleTree.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"}],\"name\":\"isRootLatest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"}],\"name\":\"isRootValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyOfElement_\",\"type\":\"bytes32\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateKeeper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyOfElement_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newElement_\",\"type\":\"bytes32\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCallWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof_\",\"type\":\"bytes\"}],\"name\":\"upgradeToWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PoseidonSMTABI is the input ABI used to generate the binding from.
// Deprecated: Use PoseidonSMTMetaData.ABI instead.
var PoseidonSMTABI = PoseidonSMTMetaData.ABI

// PoseidonSMT is an auto generated Go binding around an Ethereum contract.
type PoseidonSMT struct {
	PoseidonSMTCaller     // Read-only binding to the contract
	PoseidonSMTTransactor // Write-only binding to the contract
	PoseidonSMTFilterer   // Log filterer for contract events
}

// PoseidonSMTCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoseidonSMTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoseidonSMTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoseidonSMTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoseidonSMTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoseidonSMTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoseidonSMTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoseidonSMTSession struct {
	Contract     *PoseidonSMT      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoseidonSMTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoseidonSMTCallerSession struct {
	Contract *PoseidonSMTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PoseidonSMTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoseidonSMTTransactorSession struct {
	Contract     *PoseidonSMTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PoseidonSMTRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoseidonSMTRaw struct {
	Contract *PoseidonSMT // Generic contract binding to access the raw methods on
}

// PoseidonSMTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoseidonSMTCallerRaw struct {
	Contract *PoseidonSMTCaller // Generic read-only contract binding to access the raw methods on
}

// PoseidonSMTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoseidonSMTTransactorRaw struct {
	Contract *PoseidonSMTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoseidonSMT creates a new instance of PoseidonSMT, bound to a specific deployed contract.
func NewPoseidonSMT(address common.Address, backend bind.ContractBackend) (*PoseidonSMT, error) {
	contract, err := bindPoseidonSMT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoseidonSMT{PoseidonSMTCaller: PoseidonSMTCaller{contract: contract}, PoseidonSMTTransactor: PoseidonSMTTransactor{contract: contract}, PoseidonSMTFilterer: PoseidonSMTFilterer{contract: contract}}, nil
}

// NewPoseidonSMTCaller creates a new read-only instance of PoseidonSMT, bound to a specific deployed contract.
func NewPoseidonSMTCaller(address common.Address, caller bind.ContractCaller) (*PoseidonSMTCaller, error) {
	contract, err := bindPoseidonSMT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoseidonSMTCaller{contract: contract}, nil
}

// NewPoseidonSMTTransactor creates a new write-only instance of PoseidonSMT, bound to a specific deployed contract.
func NewPoseidonSMTTransactor(address common.Address, transactor bind.ContractTransactor) (*PoseidonSMTTransactor, error) {
	contract, err := bindPoseidonSMT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoseidonSMTTransactor{contract: contract}, nil
}

// NewPoseidonSMTFilterer creates a new log filterer instance of PoseidonSMT, bound to a specific deployed contract.
func NewPoseidonSMTFilterer(address common.Address, filterer bind.ContractFilterer) (*PoseidonSMTFilterer, error) {
	contract, err := bindPoseidonSMT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoseidonSMTFilterer{contract: contract}, nil
}

// bindPoseidonSMT binds a generic wrapper to an already deployed contract.
func bindPoseidonSMT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoseidonSMTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoseidonSMT *PoseidonSMTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoseidonSMT.Contract.PoseidonSMTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoseidonSMT *PoseidonSMTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.PoseidonSMTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoseidonSMT *PoseidonSMTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.PoseidonSMTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoseidonSMT *PoseidonSMTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoseidonSMT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoseidonSMT *PoseidonSMTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoseidonSMT *PoseidonSMTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.contract.Transact(opts, method, params...)
}

// MAGICID is a free data retrieval call binding the contract method 0xdf95574a.
//
// Solidity: function MAGIC_ID() view returns(uint8)
func (_PoseidonSMT *PoseidonSMTCaller) MAGICID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PoseidonSMT.contract.Call(opts, &out, "MAGIC_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAGICID is a free data retrieval call binding the contract method 0xdf95574a.
//
// Solidity: function MAGIC_ID() view returns(uint8)
func (_PoseidonSMT *PoseidonSMTSession) MAGICID() (uint8, error) {
	return _PoseidonSMT.Contract.MAGICID(&_PoseidonSMT.CallOpts)
}

// MAGICID is a free data retrieval call binding the contract method 0xdf95574a.
//
// Solidity: function MAGIC_ID() view returns(uint8)
func (_PoseidonSMT *PoseidonSMTCallerSession) MAGICID() (uint8, error) {
	return _PoseidonSMT.Contract.MAGICID(&_PoseidonSMT.CallOpts)
}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_PoseidonSMT *PoseidonSMTCaller) P(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoseidonSMT.contract.Call(opts, &out, "P")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_PoseidonSMT *PoseidonSMTSession) P() (*big.Int, error) {
	return _PoseidonSMT.Contract.P(&_PoseidonSMT.CallOpts)
}

// P is a free data retrieval call binding the contract method 0x8b8fbd92.
//
// Solidity: function P() view returns(uint256)
func (_PoseidonSMT *PoseidonSMTCallerSession) P() (*big.Int, error) {
	return _PoseidonSMT.Contract.P(&_PoseidonSMT.CallOpts)
}

// ROOTVALIDITY is a free data retrieval call binding the contract method 0xcffe9676.
//
// Solidity: function ROOT_VALIDITY() view returns(uint256)
func (_PoseidonSMT *PoseidonSMTCaller) ROOTVALIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoseidonSMT.contract.Call(opts, &out, "ROOT_VALIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ROOTVALIDITY is a free data retrieval call binding the contract method 0xcffe9676.
//
// Solidity: function ROOT_VALIDITY() view returns(uint256)
func (_PoseidonSMT *PoseidonSMTSession) ROOTVALIDITY() (*big.Int, error) {
	return _PoseidonSMT.Contract.ROOTVALIDITY(&_PoseidonSMT.CallOpts)
}

// ROOTVALIDITY is a free data retrieval call binding the contract method 0xcffe9676.
//
// Solidity: function ROOT_VALIDITY() view returns(uint256)
func (_PoseidonSMT *PoseidonSMTCallerSession) ROOTVALIDITY() (*big.Int, error) {
	return _PoseidonSMT.Contract.ROOTVALIDITY(&_PoseidonSMT.CallOpts)
}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_PoseidonSMT *PoseidonSMTCaller) ChainName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PoseidonSMT.contract.Call(opts, &out, "chainName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_PoseidonSMT *PoseidonSMTSession) ChainName() (string, error) {
	return _PoseidonSMT.Contract.ChainName(&_PoseidonSMT.CallOpts)
}

// ChainName is a free data retrieval call binding the contract method 0x1c93b03a.
//
// Solidity: function chainName() view returns(string)
func (_PoseidonSMT *PoseidonSMTCallerSession) ChainName() (string, error) {
	return _PoseidonSMT.Contract.ChainName(&_PoseidonSMT.CallOpts)
}

// GetNodeByKey is a free data retrieval call binding the contract method 0x083a8580.
//
// Solidity: function getNodeByKey(bytes32 key_) view returns((uint8,uint64,uint64,bytes32,bytes32,bytes32))
func (_PoseidonSMT *PoseidonSMTCaller) GetNodeByKey(opts *bind.CallOpts, key_ [32]byte) (SparseMerkleTreeNode, error) {
	var out []interface{}
	err := _PoseidonSMT.contract.Call(opts, &out, "getNodeByKey", key_)

	if err != nil {
		return *new(SparseMerkleTreeNode), err
	}

	out0 := *abi.ConvertType(out[0], new(SparseMerkleTreeNode)).(*SparseMerkleTreeNode)

	return out0, err

}

// GetNodeByKey is a free data retrieval call binding the contract method 0x083a8580.
//
// Solidity: function getNodeByKey(bytes32 key_) view returns((uint8,uint64,uint64,bytes32,bytes32,bytes32))
func (_PoseidonSMT *PoseidonSMTSession) GetNodeByKey(key_ [32]byte) (SparseMerkleTreeNode, error) {
	return _PoseidonSMT.Contract.GetNodeByKey(&_PoseidonSMT.CallOpts, key_)
}

// GetNodeByKey is a free data retrieval call binding the contract method 0x083a8580.
//
// Solidity: function getNodeByKey(bytes32 key_) view returns((uint8,uint64,uint64,bytes32,bytes32,bytes32))
func (_PoseidonSMT *PoseidonSMTCallerSession) GetNodeByKey(key_ [32]byte) (SparseMerkleTreeNode, error) {
	return _PoseidonSMT.Contract.GetNodeByKey(&_PoseidonSMT.CallOpts, key_)
}

// GetNonce is a free data retrieval call binding the contract method 0xf4fc6341.
//
// Solidity: function getNonce(uint8 methodId_) view returns(uint256)
func (_PoseidonSMT *PoseidonSMTCaller) GetNonce(opts *bind.CallOpts, methodId_ uint8) (*big.Int, error) {
	var out []interface{}
	err := _PoseidonSMT.contract.Call(opts, &out, "getNonce", methodId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xf4fc6341.
//
// Solidity: function getNonce(uint8 methodId_) view returns(uint256)
func (_PoseidonSMT *PoseidonSMTSession) GetNonce(methodId_ uint8) (*big.Int, error) {
	return _PoseidonSMT.Contract.GetNonce(&_PoseidonSMT.CallOpts, methodId_)
}

// GetNonce is a free data retrieval call binding the contract method 0xf4fc6341.
//
// Solidity: function getNonce(uint8 methodId_) view returns(uint256)
func (_PoseidonSMT *PoseidonSMTCallerSession) GetNonce(methodId_ uint8) (*big.Int, error) {
	return _PoseidonSMT.Contract.GetNonce(&_PoseidonSMT.CallOpts, methodId_)
}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 key_) view returns((bytes32,bytes32[],bool,bytes32,bytes32,bool,bytes32,bytes32))
func (_PoseidonSMT *PoseidonSMTCaller) GetProof(opts *bind.CallOpts, key_ [32]byte) (SparseMerkleTreeProof, error) {
	var out []interface{}
	err := _PoseidonSMT.contract.Call(opts, &out, "getProof", key_)

	if err != nil {
		return *new(SparseMerkleTreeProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SparseMerkleTreeProof)).(*SparseMerkleTreeProof)

	return out0, err

}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 key_) view returns((bytes32,bytes32[],bool,bytes32,bytes32,bool,bytes32,bytes32))
func (_PoseidonSMT *PoseidonSMTSession) GetProof(key_ [32]byte) (SparseMerkleTreeProof, error) {
	return _PoseidonSMT.Contract.GetProof(&_PoseidonSMT.CallOpts, key_)
}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 key_) view returns((bytes32,bytes32[],bool,bytes32,bytes32,bool,bytes32,bytes32))
func (_PoseidonSMT *PoseidonSMTCallerSession) GetProof(key_ [32]byte) (SparseMerkleTreeProof, error) {
	return _PoseidonSMT.Contract.GetProof(&_PoseidonSMT.CallOpts, key_)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_PoseidonSMT *PoseidonSMTCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoseidonSMT.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_PoseidonSMT *PoseidonSMTSession) GetRoot() ([32]byte, error) {
	return _PoseidonSMT.Contract.GetRoot(&_PoseidonSMT.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_PoseidonSMT *PoseidonSMTCallerSession) GetRoot() ([32]byte, error) {
	return _PoseidonSMT.Contract.GetRoot(&_PoseidonSMT.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_PoseidonSMT *PoseidonSMTCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoseidonSMT.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_PoseidonSMT *PoseidonSMTSession) Implementation() (common.Address, error) {
	return _PoseidonSMT.Contract.Implementation(&_PoseidonSMT.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_PoseidonSMT *PoseidonSMTCallerSession) Implementation() (common.Address, error) {
	return _PoseidonSMT.Contract.Implementation(&_PoseidonSMT.CallOpts)
}

// IsRootLatest is a free data retrieval call binding the contract method 0x8492307f.
//
// Solidity: function isRootLatest(bytes32 root_) view returns(bool)
func (_PoseidonSMT *PoseidonSMTCaller) IsRootLatest(opts *bind.CallOpts, root_ [32]byte) (bool, error) {
	var out []interface{}
	err := _PoseidonSMT.contract.Call(opts, &out, "isRootLatest", root_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRootLatest is a free data retrieval call binding the contract method 0x8492307f.
//
// Solidity: function isRootLatest(bytes32 root_) view returns(bool)
func (_PoseidonSMT *PoseidonSMTSession) IsRootLatest(root_ [32]byte) (bool, error) {
	return _PoseidonSMT.Contract.IsRootLatest(&_PoseidonSMT.CallOpts, root_)
}

// IsRootLatest is a free data retrieval call binding the contract method 0x8492307f.
//
// Solidity: function isRootLatest(bytes32 root_) view returns(bool)
func (_PoseidonSMT *PoseidonSMTCallerSession) IsRootLatest(root_ [32]byte) (bool, error) {
	return _PoseidonSMT.Contract.IsRootLatest(&_PoseidonSMT.CallOpts, root_)
}

// IsRootValid is a free data retrieval call binding the contract method 0x30ef41b4.
//
// Solidity: function isRootValid(bytes32 root_) view returns(bool)
func (_PoseidonSMT *PoseidonSMTCaller) IsRootValid(opts *bind.CallOpts, root_ [32]byte) (bool, error) {
	var out []interface{}
	err := _PoseidonSMT.contract.Call(opts, &out, "isRootValid", root_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRootValid is a free data retrieval call binding the contract method 0x30ef41b4.
//
// Solidity: function isRootValid(bytes32 root_) view returns(bool)
func (_PoseidonSMT *PoseidonSMTSession) IsRootValid(root_ [32]byte) (bool, error) {
	return _PoseidonSMT.Contract.IsRootValid(&_PoseidonSMT.CallOpts, root_)
}

// IsRootValid is a free data retrieval call binding the contract method 0x30ef41b4.
//
// Solidity: function isRootValid(bytes32 root_) view returns(bool)
func (_PoseidonSMT *PoseidonSMTCallerSession) IsRootValid(root_ [32]byte) (bool, error) {
	return _PoseidonSMT.Contract.IsRootValid(&_PoseidonSMT.CallOpts, root_)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PoseidonSMT *PoseidonSMTCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PoseidonSMT.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PoseidonSMT *PoseidonSMTSession) ProxiableUUID() ([32]byte, error) {
	return _PoseidonSMT.Contract.ProxiableUUID(&_PoseidonSMT.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_PoseidonSMT *PoseidonSMTCallerSession) ProxiableUUID() ([32]byte, error) {
	return _PoseidonSMT.Contract.ProxiableUUID(&_PoseidonSMT.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_PoseidonSMT *PoseidonSMTCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoseidonSMT.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_PoseidonSMT *PoseidonSMTSession) Signer() (common.Address, error) {
	return _PoseidonSMT.Contract.Signer(&_PoseidonSMT.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_PoseidonSMT *PoseidonSMTCallerSession) Signer() (common.Address, error) {
	return _PoseidonSMT.Contract.Signer(&_PoseidonSMT.CallOpts)
}

// StateKeeper is a free data retrieval call binding the contract method 0xc3e6f8f9.
//
// Solidity: function stateKeeper() view returns(address)
func (_PoseidonSMT *PoseidonSMTCaller) StateKeeper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoseidonSMT.contract.Call(opts, &out, "stateKeeper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StateKeeper is a free data retrieval call binding the contract method 0xc3e6f8f9.
//
// Solidity: function stateKeeper() view returns(address)
func (_PoseidonSMT *PoseidonSMTSession) StateKeeper() (common.Address, error) {
	return _PoseidonSMT.Contract.StateKeeper(&_PoseidonSMT.CallOpts)
}

// StateKeeper is a free data retrieval call binding the contract method 0xc3e6f8f9.
//
// Solidity: function stateKeeper() view returns(address)
func (_PoseidonSMT *PoseidonSMTCallerSession) StateKeeper() (common.Address, error) {
	return _PoseidonSMT.Contract.StateKeeper(&_PoseidonSMT.CallOpts)
}

// PoseidonSMTInit is a paid mutator transaction binding the contract method 0x897f0f79.
//
// Solidity: function __PoseidonSMT_init(address signer_, string chainName_, address stateKeeper_, uint256 treeHeight_) returns()
func (_PoseidonSMT *PoseidonSMTTransactor) PoseidonSMTInit(opts *bind.TransactOpts, signer_ common.Address, chainName_ string, stateKeeper_ common.Address, treeHeight_ *big.Int) (*types.Transaction, error) {
	return _PoseidonSMT.contract.Transact(opts, "__PoseidonSMT_init", signer_, chainName_, stateKeeper_, treeHeight_)
}

// PoseidonSMTInit is a paid mutator transaction binding the contract method 0x897f0f79.
//
// Solidity: function __PoseidonSMT_init(address signer_, string chainName_, address stateKeeper_, uint256 treeHeight_) returns()
func (_PoseidonSMT *PoseidonSMTSession) PoseidonSMTInit(signer_ common.Address, chainName_ string, stateKeeper_ common.Address, treeHeight_ *big.Int) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.PoseidonSMTInit(&_PoseidonSMT.TransactOpts, signer_, chainName_, stateKeeper_, treeHeight_)
}

// PoseidonSMTInit is a paid mutator transaction binding the contract method 0x897f0f79.
//
// Solidity: function __PoseidonSMT_init(address signer_, string chainName_, address stateKeeper_, uint256 treeHeight_) returns()
func (_PoseidonSMT *PoseidonSMTTransactorSession) PoseidonSMTInit(signer_ common.Address, chainName_ string, stateKeeper_ common.Address, treeHeight_ *big.Int) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.PoseidonSMTInit(&_PoseidonSMT.TransactOpts, signer_, chainName_, stateKeeper_, treeHeight_)
}

// Add is a paid mutator transaction binding the contract method 0xd1de592a.
//
// Solidity: function add(bytes32 keyOfElement_, bytes32 element_) returns()
func (_PoseidonSMT *PoseidonSMTTransactor) Add(opts *bind.TransactOpts, keyOfElement_ [32]byte, element_ [32]byte) (*types.Transaction, error) {
	return _PoseidonSMT.contract.Transact(opts, "add", keyOfElement_, element_)
}

// Add is a paid mutator transaction binding the contract method 0xd1de592a.
//
// Solidity: function add(bytes32 keyOfElement_, bytes32 element_) returns()
func (_PoseidonSMT *PoseidonSMTSession) Add(keyOfElement_ [32]byte, element_ [32]byte) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.Add(&_PoseidonSMT.TransactOpts, keyOfElement_, element_)
}

// Add is a paid mutator transaction binding the contract method 0xd1de592a.
//
// Solidity: function add(bytes32 keyOfElement_, bytes32 element_) returns()
func (_PoseidonSMT *PoseidonSMTTransactorSession) Add(keyOfElement_ [32]byte, element_ [32]byte) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.Add(&_PoseidonSMT.TransactOpts, keyOfElement_, element_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_PoseidonSMT *PoseidonSMTTransactor) ChangeSigner(opts *bind.TransactOpts, newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _PoseidonSMT.contract.Transact(opts, "changeSigner", newSignerPubKey_, signature_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_PoseidonSMT *PoseidonSMTSession) ChangeSigner(newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.ChangeSigner(&_PoseidonSMT.TransactOpts, newSignerPubKey_, signature_)
}

// ChangeSigner is a paid mutator transaction binding the contract method 0x497f6959.
//
// Solidity: function changeSigner(bytes newSignerPubKey_, bytes signature_) returns()
func (_PoseidonSMT *PoseidonSMTTransactorSession) ChangeSigner(newSignerPubKey_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.ChangeSigner(&_PoseidonSMT.TransactOpts, newSignerPubKey_, signature_)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 keyOfElement_) returns()
func (_PoseidonSMT *PoseidonSMTTransactor) Remove(opts *bind.TransactOpts, keyOfElement_ [32]byte) (*types.Transaction, error) {
	return _PoseidonSMT.contract.Transact(opts, "remove", keyOfElement_)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 keyOfElement_) returns()
func (_PoseidonSMT *PoseidonSMTSession) Remove(keyOfElement_ [32]byte) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.Remove(&_PoseidonSMT.TransactOpts, keyOfElement_)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 keyOfElement_) returns()
func (_PoseidonSMT *PoseidonSMTTransactorSession) Remove(keyOfElement_ [32]byte) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.Remove(&_PoseidonSMT.TransactOpts, keyOfElement_)
}

// Update is a paid mutator transaction binding the contract method 0x13f57c3e.
//
// Solidity: function update(bytes32 keyOfElement_, bytes32 newElement_) returns()
func (_PoseidonSMT *PoseidonSMTTransactor) Update(opts *bind.TransactOpts, keyOfElement_ [32]byte, newElement_ [32]byte) (*types.Transaction, error) {
	return _PoseidonSMT.contract.Transact(opts, "update", keyOfElement_, newElement_)
}

// Update is a paid mutator transaction binding the contract method 0x13f57c3e.
//
// Solidity: function update(bytes32 keyOfElement_, bytes32 newElement_) returns()
func (_PoseidonSMT *PoseidonSMTSession) Update(keyOfElement_ [32]byte, newElement_ [32]byte) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.Update(&_PoseidonSMT.TransactOpts, keyOfElement_, newElement_)
}

// Update is a paid mutator transaction binding the contract method 0x13f57c3e.
//
// Solidity: function update(bytes32 keyOfElement_, bytes32 newElement_) returns()
func (_PoseidonSMT *PoseidonSMTTransactorSession) Update(keyOfElement_ [32]byte, newElement_ [32]byte) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.Update(&_PoseidonSMT.TransactOpts, keyOfElement_, newElement_)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PoseidonSMT *PoseidonSMTTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _PoseidonSMT.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PoseidonSMT *PoseidonSMTSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.UpgradeTo(&_PoseidonSMT.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PoseidonSMT *PoseidonSMTTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.UpgradeTo(&_PoseidonSMT.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PoseidonSMT *PoseidonSMTTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PoseidonSMT.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PoseidonSMT *PoseidonSMTSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.UpgradeToAndCall(&_PoseidonSMT.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PoseidonSMT *PoseidonSMTTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.UpgradeToAndCall(&_PoseidonSMT.TransactOpts, newImplementation, data)
}

// UpgradeToAndCallWithProof is a paid mutator transaction binding the contract method 0xbf2c6db7.
//
// Solidity: function upgradeToAndCallWithProof(address newImplementation_, bytes proof_, bytes data_) returns()
func (_PoseidonSMT *PoseidonSMTTransactor) UpgradeToAndCallWithProof(opts *bind.TransactOpts, newImplementation_ common.Address, proof_ []byte, data_ []byte) (*types.Transaction, error) {
	return _PoseidonSMT.contract.Transact(opts, "upgradeToAndCallWithProof", newImplementation_, proof_, data_)
}

// UpgradeToAndCallWithProof is a paid mutator transaction binding the contract method 0xbf2c6db7.
//
// Solidity: function upgradeToAndCallWithProof(address newImplementation_, bytes proof_, bytes data_) returns()
func (_PoseidonSMT *PoseidonSMTSession) UpgradeToAndCallWithProof(newImplementation_ common.Address, proof_ []byte, data_ []byte) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.UpgradeToAndCallWithProof(&_PoseidonSMT.TransactOpts, newImplementation_, proof_, data_)
}

// UpgradeToAndCallWithProof is a paid mutator transaction binding the contract method 0xbf2c6db7.
//
// Solidity: function upgradeToAndCallWithProof(address newImplementation_, bytes proof_, bytes data_) returns()
func (_PoseidonSMT *PoseidonSMTTransactorSession) UpgradeToAndCallWithProof(newImplementation_ common.Address, proof_ []byte, data_ []byte) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.UpgradeToAndCallWithProof(&_PoseidonSMT.TransactOpts, newImplementation_, proof_, data_)
}

// UpgradeToWithProof is a paid mutator transaction binding the contract method 0x628543ab.
//
// Solidity: function upgradeToWithProof(address newImplementation_, bytes proof_) returns()
func (_PoseidonSMT *PoseidonSMTTransactor) UpgradeToWithProof(opts *bind.TransactOpts, newImplementation_ common.Address, proof_ []byte) (*types.Transaction, error) {
	return _PoseidonSMT.contract.Transact(opts, "upgradeToWithProof", newImplementation_, proof_)
}

// UpgradeToWithProof is a paid mutator transaction binding the contract method 0x628543ab.
//
// Solidity: function upgradeToWithProof(address newImplementation_, bytes proof_) returns()
func (_PoseidonSMT *PoseidonSMTSession) UpgradeToWithProof(newImplementation_ common.Address, proof_ []byte) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.UpgradeToWithProof(&_PoseidonSMT.TransactOpts, newImplementation_, proof_)
}

// UpgradeToWithProof is a paid mutator transaction binding the contract method 0x628543ab.
//
// Solidity: function upgradeToWithProof(address newImplementation_, bytes proof_) returns()
func (_PoseidonSMT *PoseidonSMTTransactorSession) UpgradeToWithProof(newImplementation_ common.Address, proof_ []byte) (*types.Transaction, error) {
	return _PoseidonSMT.Contract.UpgradeToWithProof(&_PoseidonSMT.TransactOpts, newImplementation_, proof_)
}

// PoseidonSMTAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the PoseidonSMT contract.
type PoseidonSMTAdminChangedIterator struct {
	Event *PoseidonSMTAdminChanged // Event containing the contract specifics and raw log

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
func (it *PoseidonSMTAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoseidonSMTAdminChanged)
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
		it.Event = new(PoseidonSMTAdminChanged)
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
func (it *PoseidonSMTAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoseidonSMTAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoseidonSMTAdminChanged represents a AdminChanged event raised by the PoseidonSMT contract.
type PoseidonSMTAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PoseidonSMT *PoseidonSMTFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*PoseidonSMTAdminChangedIterator, error) {

	logs, sub, err := _PoseidonSMT.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &PoseidonSMTAdminChangedIterator{contract: _PoseidonSMT.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PoseidonSMT *PoseidonSMTFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *PoseidonSMTAdminChanged) (event.Subscription, error) {

	logs, sub, err := _PoseidonSMT.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoseidonSMTAdminChanged)
				if err := _PoseidonSMT.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PoseidonSMT *PoseidonSMTFilterer) ParseAdminChanged(log types.Log) (*PoseidonSMTAdminChanged, error) {
	event := new(PoseidonSMTAdminChanged)
	if err := _PoseidonSMT.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoseidonSMTBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the PoseidonSMT contract.
type PoseidonSMTBeaconUpgradedIterator struct {
	Event *PoseidonSMTBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *PoseidonSMTBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoseidonSMTBeaconUpgraded)
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
		it.Event = new(PoseidonSMTBeaconUpgraded)
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
func (it *PoseidonSMTBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoseidonSMTBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoseidonSMTBeaconUpgraded represents a BeaconUpgraded event raised by the PoseidonSMT contract.
type PoseidonSMTBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_PoseidonSMT *PoseidonSMTFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*PoseidonSMTBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _PoseidonSMT.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &PoseidonSMTBeaconUpgradedIterator{contract: _PoseidonSMT.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_PoseidonSMT *PoseidonSMTFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *PoseidonSMTBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _PoseidonSMT.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoseidonSMTBeaconUpgraded)
				if err := _PoseidonSMT.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_PoseidonSMT *PoseidonSMTFilterer) ParseBeaconUpgraded(log types.Log) (*PoseidonSMTBeaconUpgraded, error) {
	event := new(PoseidonSMTBeaconUpgraded)
	if err := _PoseidonSMT.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoseidonSMTInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PoseidonSMT contract.
type PoseidonSMTInitializedIterator struct {
	Event *PoseidonSMTInitialized // Event containing the contract specifics and raw log

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
func (it *PoseidonSMTInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoseidonSMTInitialized)
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
		it.Event = new(PoseidonSMTInitialized)
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
func (it *PoseidonSMTInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoseidonSMTInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoseidonSMTInitialized represents a Initialized event raised by the PoseidonSMT contract.
type PoseidonSMTInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PoseidonSMT *PoseidonSMTFilterer) FilterInitialized(opts *bind.FilterOpts) (*PoseidonSMTInitializedIterator, error) {

	logs, sub, err := _PoseidonSMT.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PoseidonSMTInitializedIterator{contract: _PoseidonSMT.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PoseidonSMT *PoseidonSMTFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PoseidonSMTInitialized) (event.Subscription, error) {

	logs, sub, err := _PoseidonSMT.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoseidonSMTInitialized)
				if err := _PoseidonSMT.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PoseidonSMT *PoseidonSMTFilterer) ParseInitialized(log types.Log) (*PoseidonSMTInitialized, error) {
	event := new(PoseidonSMTInitialized)
	if err := _PoseidonSMT.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoseidonSMTRootUpdatedIterator is returned from FilterRootUpdated and is used to iterate over the raw logs and unpacked data for RootUpdated events raised by the PoseidonSMT contract.
type PoseidonSMTRootUpdatedIterator struct {
	Event *PoseidonSMTRootUpdated // Event containing the contract specifics and raw log

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
func (it *PoseidonSMTRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoseidonSMTRootUpdated)
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
		it.Event = new(PoseidonSMTRootUpdated)
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
func (it *PoseidonSMTRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoseidonSMTRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoseidonSMTRootUpdated represents a RootUpdated event raised by the PoseidonSMT contract.
type PoseidonSMTRootUpdated struct {
	Root [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRootUpdated is a free log retrieval operation binding the contract event 0x2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa4.
//
// Solidity: event RootUpdated(bytes32 root)
func (_PoseidonSMT *PoseidonSMTFilterer) FilterRootUpdated(opts *bind.FilterOpts) (*PoseidonSMTRootUpdatedIterator, error) {

	logs, sub, err := _PoseidonSMT.contract.FilterLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return &PoseidonSMTRootUpdatedIterator{contract: _PoseidonSMT.contract, event: "RootUpdated", logs: logs, sub: sub}, nil
}

// WatchRootUpdated is a free log subscription operation binding the contract event 0x2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa4.
//
// Solidity: event RootUpdated(bytes32 root)
func (_PoseidonSMT *PoseidonSMTFilterer) WatchRootUpdated(opts *bind.WatchOpts, sink chan<- *PoseidonSMTRootUpdated) (event.Subscription, error) {

	logs, sub, err := _PoseidonSMT.contract.WatchLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoseidonSMTRootUpdated)
				if err := _PoseidonSMT.contract.UnpackLog(event, "RootUpdated", log); err != nil {
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

// ParseRootUpdated is a log parse operation binding the contract event 0x2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa4.
//
// Solidity: event RootUpdated(bytes32 root)
func (_PoseidonSMT *PoseidonSMTFilterer) ParseRootUpdated(log types.Log) (*PoseidonSMTRootUpdated, error) {
	event := new(PoseidonSMTRootUpdated)
	if err := _PoseidonSMT.contract.UnpackLog(event, "RootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoseidonSMTUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the PoseidonSMT contract.
type PoseidonSMTUpgradedIterator struct {
	Event *PoseidonSMTUpgraded // Event containing the contract specifics and raw log

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
func (it *PoseidonSMTUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoseidonSMTUpgraded)
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
		it.Event = new(PoseidonSMTUpgraded)
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
func (it *PoseidonSMTUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoseidonSMTUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoseidonSMTUpgraded represents a Upgraded event raised by the PoseidonSMT contract.
type PoseidonSMTUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PoseidonSMT *PoseidonSMTFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*PoseidonSMTUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PoseidonSMT.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &PoseidonSMTUpgradedIterator{contract: _PoseidonSMT.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PoseidonSMT *PoseidonSMTFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PoseidonSMTUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PoseidonSMT.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoseidonSMTUpgraded)
				if err := _PoseidonSMT.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PoseidonSMT *PoseidonSMTFilterer) ParseUpgraded(log types.Log) (*PoseidonSMTUpgraded, error) {
	event := new(PoseidonSMTUpgraded)
	if err := _PoseidonSMT.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
