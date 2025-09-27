// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package voting

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

// Candidate is an auto generated low-level Go binding around an user-defined struct.
type Candidate struct {
	Id     *big.Int
	Name   string
	Votes  *big.Int
	Voters []common.Address
}

// Pool is an auto generated low-level Go binding around an user-defined struct.
type Pool struct {
	Id          *big.Int
	Name        string
	Description string
	CreatedAt   *big.Int
	ClosedAt    *big.Int
	TotalVotes  *big.Int
	CreatedBy   common.Address
}

// VotingMetaData contains all meta data concerning the Voting contract.
var VotingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"moderator\",\"type\":\"address\"}],\"name\":\"ModeratorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"moderator\",\"type\":\"address\"}],\"name\":\"ModeratorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"closedAt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"cadidates\",\"type\":\"string[]\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cadidateId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PoolVoted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"moderator\",\"type\":\"address\"}],\"name\":\"addModerator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_cadidates\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"closedAt\",\"type\":\"uint256\"}],\"name\":\"createPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"getCandidates\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"voters\",\"type\":\"address[]\"}],\"internalType\":\"structCandidate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"getPool\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVotes\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"createdBy\",\"type\":\"address\"}],\"internalType\":\"structPool\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolList\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVotes\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"createdBy\",\"type\":\"address\"}],\"internalType\":\"structPool[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"moderator\",\"type\":\"address\"}],\"name\":\"getRole\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"moderator\",\"type\":\"address\"}],\"name\":\"isModerator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"isVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"moderator\",\"type\":\"address\"}],\"name\":\"removeModerator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VotingABI is the input ABI used to generate the binding from.
// Deprecated: Use VotingMetaData.ABI instead.
var VotingABI = VotingMetaData.ABI

// Voting is an auto generated Go binding around an Ethereum contract.
type Voting struct {
	VotingCaller     // Read-only binding to the contract
	VotingTransactor // Write-only binding to the contract
	VotingFilterer   // Log filterer for contract events
}

// VotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingSession struct {
	Contract     *Voting           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingCallerSession struct {
	Contract *VotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingTransactorSession struct {
	Contract     *VotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingRaw struct {
	Contract *Voting // Generic contract binding to access the raw methods on
}

// VotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingCallerRaw struct {
	Contract *VotingCaller // Generic read-only contract binding to access the raw methods on
}

// VotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingTransactorRaw struct {
	Contract *VotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoting creates a new instance of Voting, bound to a specific deployed contract.
func NewVoting(address common.Address, backend bind.ContractBackend) (*Voting, error) {
	contract, err := bindVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// NewVotingCaller creates a new read-only instance of Voting, bound to a specific deployed contract.
func NewVotingCaller(address common.Address, caller bind.ContractCaller) (*VotingCaller, error) {
	contract, err := bindVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotingCaller{contract: contract}, nil
}

// NewVotingTransactor creates a new write-only instance of Voting, bound to a specific deployed contract.
func NewVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingTransactor, error) {
	contract, err := bindVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotingTransactor{contract: contract}, nil
}

// NewVotingFilterer creates a new log filterer instance of Voting, bound to a specific deployed contract.
func NewVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*VotingFilterer, error) {
	contract, err := bindVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotingFilterer{contract: contract}, nil
}

// bindVoting binds a generic wrapper to an already deployed contract.
func bindVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VotingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.VotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transact(opts, method, params...)
}

// GetCandidates is a free data retrieval call binding the contract method 0x3e39a7a5.
//
// Solidity: function getCandidates(uint256 poolId) view returns((uint256,string,uint256,address[])[])
func (_Voting *VotingCaller) GetCandidates(opts *bind.CallOpts, poolId *big.Int) ([]Candidate, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getCandidates", poolId)

	if err != nil {
		return *new([]Candidate), err
	}

	out0 := *abi.ConvertType(out[0], new([]Candidate)).(*[]Candidate)

	return out0, err

}

// GetCandidates is a free data retrieval call binding the contract method 0x3e39a7a5.
//
// Solidity: function getCandidates(uint256 poolId) view returns((uint256,string,uint256,address[])[])
func (_Voting *VotingSession) GetCandidates(poolId *big.Int) ([]Candidate, error) {
	return _Voting.Contract.GetCandidates(&_Voting.CallOpts, poolId)
}

// GetCandidates is a free data retrieval call binding the contract method 0x3e39a7a5.
//
// Solidity: function getCandidates(uint256 poolId) view returns((uint256,string,uint256,address[])[])
func (_Voting *VotingCallerSession) GetCandidates(poolId *big.Int) ([]Candidate, error) {
	return _Voting.Contract.GetCandidates(&_Voting.CallOpts, poolId)
}

// GetPool is a free data retrieval call binding the contract method 0x068bcd8d.
//
// Solidity: function getPool(uint256 poolId) view returns((uint256,string,string,uint256,uint256,uint256,address))
func (_Voting *VotingCaller) GetPool(opts *bind.CallOpts, poolId *big.Int) (Pool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getPool", poolId)

	if err != nil {
		return *new(Pool), err
	}

	out0 := *abi.ConvertType(out[0], new(Pool)).(*Pool)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x068bcd8d.
//
// Solidity: function getPool(uint256 poolId) view returns((uint256,string,string,uint256,uint256,uint256,address))
func (_Voting *VotingSession) GetPool(poolId *big.Int) (Pool, error) {
	return _Voting.Contract.GetPool(&_Voting.CallOpts, poolId)
}

// GetPool is a free data retrieval call binding the contract method 0x068bcd8d.
//
// Solidity: function getPool(uint256 poolId) view returns((uint256,string,string,uint256,uint256,uint256,address))
func (_Voting *VotingCallerSession) GetPool(poolId *big.Int) (Pool, error) {
	return _Voting.Contract.GetPool(&_Voting.CallOpts, poolId)
}

// GetPoolList is a free data retrieval call binding the contract method 0xd41dcbea.
//
// Solidity: function getPoolList() view returns((uint256,string,string,uint256,uint256,uint256,address)[])
func (_Voting *VotingCaller) GetPoolList(opts *bind.CallOpts) ([]Pool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getPoolList")

	if err != nil {
		return *new([]Pool), err
	}

	out0 := *abi.ConvertType(out[0], new([]Pool)).(*[]Pool)

	return out0, err

}

// GetPoolList is a free data retrieval call binding the contract method 0xd41dcbea.
//
// Solidity: function getPoolList() view returns((uint256,string,string,uint256,uint256,uint256,address)[])
func (_Voting *VotingSession) GetPoolList() ([]Pool, error) {
	return _Voting.Contract.GetPoolList(&_Voting.CallOpts)
}

// GetPoolList is a free data retrieval call binding the contract method 0xd41dcbea.
//
// Solidity: function getPoolList() view returns((uint256,string,string,uint256,uint256,uint256,address)[])
func (_Voting *VotingCallerSession) GetPoolList() ([]Pool, error) {
	return _Voting.Contract.GetPoolList(&_Voting.CallOpts)
}

// GetRole is a free data retrieval call binding the contract method 0x44276733.
//
// Solidity: function getRole(address moderator) view returns(uint64)
func (_Voting *VotingCaller) GetRole(opts *bind.CallOpts, moderator common.Address) (uint64, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getRole", moderator)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetRole is a free data retrieval call binding the contract method 0x44276733.
//
// Solidity: function getRole(address moderator) view returns(uint64)
func (_Voting *VotingSession) GetRole(moderator common.Address) (uint64, error) {
	return _Voting.Contract.GetRole(&_Voting.CallOpts, moderator)
}

// GetRole is a free data retrieval call binding the contract method 0x44276733.
//
// Solidity: function getRole(address moderator) view returns(uint64)
func (_Voting *VotingCallerSession) GetRole(moderator common.Address) (uint64, error) {
	return _Voting.Contract.GetRole(&_Voting.CallOpts, moderator)
}

// IsModerator is a free data retrieval call binding the contract method 0xfa6f3936.
//
// Solidity: function isModerator(address moderator) view returns(bool)
func (_Voting *VotingCaller) IsModerator(opts *bind.CallOpts, moderator common.Address) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "isModerator", moderator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModerator is a free data retrieval call binding the contract method 0xfa6f3936.
//
// Solidity: function isModerator(address moderator) view returns(bool)
func (_Voting *VotingSession) IsModerator(moderator common.Address) (bool, error) {
	return _Voting.Contract.IsModerator(&_Voting.CallOpts, moderator)
}

// IsModerator is a free data retrieval call binding the contract method 0xfa6f3936.
//
// Solidity: function isModerator(address moderator) view returns(bool)
func (_Voting *VotingCallerSession) IsModerator(moderator common.Address) (bool, error) {
	return _Voting.Contract.IsModerator(&_Voting.CallOpts, moderator)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Voting *VotingCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Voting *VotingSession) IsOwner() (bool, error) {
	return _Voting.Contract.IsOwner(&_Voting.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Voting *VotingCallerSession) IsOwner() (bool, error) {
	return _Voting.Contract.IsOwner(&_Voting.CallOpts)
}

// IsVoted is a free data retrieval call binding the contract method 0x49c1a581.
//
// Solidity: function isVoted(uint256 poolId, address voter) view returns(bool)
func (_Voting *VotingCaller) IsVoted(opts *bind.CallOpts, poolId *big.Int, voter common.Address) (bool, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "isVoted", poolId, voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVoted is a free data retrieval call binding the contract method 0x49c1a581.
//
// Solidity: function isVoted(uint256 poolId, address voter) view returns(bool)
func (_Voting *VotingSession) IsVoted(poolId *big.Int, voter common.Address) (bool, error) {
	return _Voting.Contract.IsVoted(&_Voting.CallOpts, poolId, voter)
}

// IsVoted is a free data retrieval call binding the contract method 0x49c1a581.
//
// Solidity: function isVoted(uint256 poolId, address voter) view returns(bool)
func (_Voting *VotingCallerSession) IsVoted(poolId *big.Int, voter common.Address) (bool, error) {
	return _Voting.Contract.IsVoted(&_Voting.CallOpts, poolId, voter)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Voting *VotingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Voting *VotingSession) Owner() (common.Address, error) {
	return _Voting.Contract.Owner(&_Voting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Voting *VotingCallerSession) Owner() (common.Address, error) {
	return _Voting.Contract.Owner(&_Voting.CallOpts)
}

// AddModerator is a paid mutator transaction binding the contract method 0xb532e4cb.
//
// Solidity: function addModerator(address moderator) returns()
func (_Voting *VotingTransactor) AddModerator(opts *bind.TransactOpts, moderator common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "addModerator", moderator)
}

// AddModerator is a paid mutator transaction binding the contract method 0xb532e4cb.
//
// Solidity: function addModerator(address moderator) returns()
func (_Voting *VotingSession) AddModerator(moderator common.Address) (*types.Transaction, error) {
	return _Voting.Contract.AddModerator(&_Voting.TransactOpts, moderator)
}

// AddModerator is a paid mutator transaction binding the contract method 0xb532e4cb.
//
// Solidity: function addModerator(address moderator) returns()
func (_Voting *VotingTransactorSession) AddModerator(moderator common.Address) (*types.Transaction, error) {
	return _Voting.Contract.AddModerator(&_Voting.TransactOpts, moderator)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Voting *VotingTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Voting *VotingSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Voting.Contract.ChangeOwner(&_Voting.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Voting *VotingTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Voting.Contract.ChangeOwner(&_Voting.TransactOpts, newOwner)
}

// CreatePool is a paid mutator transaction binding the contract method 0xd802310e.
//
// Solidity: function createPool(string name, string description, string[] _cadidates, uint256 closedAt) returns()
func (_Voting *VotingTransactor) CreatePool(opts *bind.TransactOpts, name string, description string, _cadidates []string, closedAt *big.Int) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "createPool", name, description, _cadidates, closedAt)
}

// CreatePool is a paid mutator transaction binding the contract method 0xd802310e.
//
// Solidity: function createPool(string name, string description, string[] _cadidates, uint256 closedAt) returns()
func (_Voting *VotingSession) CreatePool(name string, description string, _cadidates []string, closedAt *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.CreatePool(&_Voting.TransactOpts, name, description, _cadidates, closedAt)
}

// CreatePool is a paid mutator transaction binding the contract method 0xd802310e.
//
// Solidity: function createPool(string name, string description, string[] _cadidates, uint256 closedAt) returns()
func (_Voting *VotingTransactorSession) CreatePool(name string, description string, _cadidates []string, closedAt *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.CreatePool(&_Voting.TransactOpts, name, description, _cadidates, closedAt)
}

// RemoveModerator is a paid mutator transaction binding the contract method 0x869d785f.
//
// Solidity: function removeModerator(address moderator) returns()
func (_Voting *VotingTransactor) RemoveModerator(opts *bind.TransactOpts, moderator common.Address) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "removeModerator", moderator)
}

// RemoveModerator is a paid mutator transaction binding the contract method 0x869d785f.
//
// Solidity: function removeModerator(address moderator) returns()
func (_Voting *VotingSession) RemoveModerator(moderator common.Address) (*types.Transaction, error) {
	return _Voting.Contract.RemoveModerator(&_Voting.TransactOpts, moderator)
}

// RemoveModerator is a paid mutator transaction binding the contract method 0x869d785f.
//
// Solidity: function removeModerator(address moderator) returns()
func (_Voting *VotingTransactorSession) RemoveModerator(moderator common.Address) (*types.Transaction, error) {
	return _Voting.Contract.RemoveModerator(&_Voting.TransactOpts, moderator)
}

// Vote is a paid mutator transaction binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 poolId, uint256 candidateId) returns()
func (_Voting *VotingTransactor) Vote(opts *bind.TransactOpts, poolId *big.Int, candidateId *big.Int) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "vote", poolId, candidateId)
}

// Vote is a paid mutator transaction binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 poolId, uint256 candidateId) returns()
func (_Voting *VotingSession) Vote(poolId *big.Int, candidateId *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, poolId, candidateId)
}

// Vote is a paid mutator transaction binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 poolId, uint256 candidateId) returns()
func (_Voting *VotingTransactorSession) Vote(poolId *big.Int, candidateId *big.Int) (*types.Transaction, error) {
	return _Voting.Contract.Vote(&_Voting.TransactOpts, poolId, candidateId)
}

// VotingModeratorAddedIterator is returned from FilterModeratorAdded and is used to iterate over the raw logs and unpacked data for ModeratorAdded events raised by the Voting contract.
type VotingModeratorAddedIterator struct {
	Event *VotingModeratorAdded // Event containing the contract specifics and raw log

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
func (it *VotingModeratorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingModeratorAdded)
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
		it.Event = new(VotingModeratorAdded)
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
func (it *VotingModeratorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingModeratorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingModeratorAdded represents a ModeratorAdded event raised by the Voting contract.
type VotingModeratorAdded struct {
	Moderator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterModeratorAdded is a free log retrieval operation binding the contract event 0xd378ad41c1a753fd1ba9ec0fcd7970526c175b68545b4a02d6d15e7606fe3596.
//
// Solidity: event ModeratorAdded(address moderator)
func (_Voting *VotingFilterer) FilterModeratorAdded(opts *bind.FilterOpts) (*VotingModeratorAddedIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "ModeratorAdded")
	if err != nil {
		return nil, err
	}
	return &VotingModeratorAddedIterator{contract: _Voting.contract, event: "ModeratorAdded", logs: logs, sub: sub}, nil
}

// WatchModeratorAdded is a free log subscription operation binding the contract event 0xd378ad41c1a753fd1ba9ec0fcd7970526c175b68545b4a02d6d15e7606fe3596.
//
// Solidity: event ModeratorAdded(address moderator)
func (_Voting *VotingFilterer) WatchModeratorAdded(opts *bind.WatchOpts, sink chan<- *VotingModeratorAdded) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "ModeratorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingModeratorAdded)
				if err := _Voting.contract.UnpackLog(event, "ModeratorAdded", log); err != nil {
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

// ParseModeratorAdded is a log parse operation binding the contract event 0xd378ad41c1a753fd1ba9ec0fcd7970526c175b68545b4a02d6d15e7606fe3596.
//
// Solidity: event ModeratorAdded(address moderator)
func (_Voting *VotingFilterer) ParseModeratorAdded(log types.Log) (*VotingModeratorAdded, error) {
	event := new(VotingModeratorAdded)
	if err := _Voting.contract.UnpackLog(event, "ModeratorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingModeratorRemovedIterator is returned from FilterModeratorRemoved and is used to iterate over the raw logs and unpacked data for ModeratorRemoved events raised by the Voting contract.
type VotingModeratorRemovedIterator struct {
	Event *VotingModeratorRemoved // Event containing the contract specifics and raw log

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
func (it *VotingModeratorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingModeratorRemoved)
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
		it.Event = new(VotingModeratorRemoved)
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
func (it *VotingModeratorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingModeratorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingModeratorRemoved represents a ModeratorRemoved event raised by the Voting contract.
type VotingModeratorRemoved struct {
	Moderator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterModeratorRemoved is a free log retrieval operation binding the contract event 0x7a9f1e23d5426b34819d173153d59084cc3578d5a129b80bb27df683ac6b2278.
//
// Solidity: event ModeratorRemoved(address moderator)
func (_Voting *VotingFilterer) FilterModeratorRemoved(opts *bind.FilterOpts) (*VotingModeratorRemovedIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "ModeratorRemoved")
	if err != nil {
		return nil, err
	}
	return &VotingModeratorRemovedIterator{contract: _Voting.contract, event: "ModeratorRemoved", logs: logs, sub: sub}, nil
}

// WatchModeratorRemoved is a free log subscription operation binding the contract event 0x7a9f1e23d5426b34819d173153d59084cc3578d5a129b80bb27df683ac6b2278.
//
// Solidity: event ModeratorRemoved(address moderator)
func (_Voting *VotingFilterer) WatchModeratorRemoved(opts *bind.WatchOpts, sink chan<- *VotingModeratorRemoved) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "ModeratorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingModeratorRemoved)
				if err := _Voting.contract.UnpackLog(event, "ModeratorRemoved", log); err != nil {
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

// ParseModeratorRemoved is a log parse operation binding the contract event 0x7a9f1e23d5426b34819d173153d59084cc3578d5a129b80bb27df683ac6b2278.
//
// Solidity: event ModeratorRemoved(address moderator)
func (_Voting *VotingFilterer) ParseModeratorRemoved(log types.Log) (*VotingModeratorRemoved, error) {
	event := new(VotingModeratorRemoved)
	if err := _Voting.contract.UnpackLog(event, "ModeratorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the Voting contract.
type VotingPoolCreatedIterator struct {
	Event *VotingPoolCreated // Event containing the contract specifics and raw log

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
func (it *VotingPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingPoolCreated)
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
		it.Event = new(VotingPoolCreated)
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
func (it *VotingPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingPoolCreated represents a PoolCreated event raised by the Voting contract.
type VotingPoolCreated struct {
	Id          *big.Int
	Name        string
	Description string
	ClosedAt    *big.Int
	Timestamp   *big.Int
	Cadidates   []string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x1d9ed257ba18f079fa73a2be0a204f1883bf6fbea149574f90e7853dda2d13e7.
//
// Solidity: event PoolCreated(uint256 id, string name, string description, uint256 closedAt, uint256 timestamp, string[] cadidates)
func (_Voting *VotingFilterer) FilterPoolCreated(opts *bind.FilterOpts) (*VotingPoolCreatedIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "PoolCreated")
	if err != nil {
		return nil, err
	}
	return &VotingPoolCreatedIterator{contract: _Voting.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x1d9ed257ba18f079fa73a2be0a204f1883bf6fbea149574f90e7853dda2d13e7.
//
// Solidity: event PoolCreated(uint256 id, string name, string description, uint256 closedAt, uint256 timestamp, string[] cadidates)
func (_Voting *VotingFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *VotingPoolCreated) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "PoolCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingPoolCreated)
				if err := _Voting.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0x1d9ed257ba18f079fa73a2be0a204f1883bf6fbea149574f90e7853dda2d13e7.
//
// Solidity: event PoolCreated(uint256 id, string name, string description, uint256 closedAt, uint256 timestamp, string[] cadidates)
func (_Voting *VotingFilterer) ParsePoolCreated(log types.Log) (*VotingPoolCreated, error) {
	event := new(VotingPoolCreated)
	if err := _Voting.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VotingPoolVotedIterator is returned from FilterPoolVoted and is used to iterate over the raw logs and unpacked data for PoolVoted events raised by the Voting contract.
type VotingPoolVotedIterator struct {
	Event *VotingPoolVoted // Event containing the contract specifics and raw log

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
func (it *VotingPoolVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VotingPoolVoted)
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
		it.Event = new(VotingPoolVoted)
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
func (it *VotingPoolVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VotingPoolVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VotingPoolVoted represents a PoolVoted event raised by the Voting contract.
type VotingPoolVoted struct {
	Id         *big.Int
	CadidateId *big.Int
	Name       string
	Voter      common.Address
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPoolVoted is a free log retrieval operation binding the contract event 0x27e4e9f06dee6980758f4150556c677abc1c6ba9832c4be48ce08e0e850e12f4.
//
// Solidity: event PoolVoted(uint256 id, uint256 cadidateId, string name, address voter, uint256 timestamp)
func (_Voting *VotingFilterer) FilterPoolVoted(opts *bind.FilterOpts) (*VotingPoolVotedIterator, error) {

	logs, sub, err := _Voting.contract.FilterLogs(opts, "PoolVoted")
	if err != nil {
		return nil, err
	}
	return &VotingPoolVotedIterator{contract: _Voting.contract, event: "PoolVoted", logs: logs, sub: sub}, nil
}

// WatchPoolVoted is a free log subscription operation binding the contract event 0x27e4e9f06dee6980758f4150556c677abc1c6ba9832c4be48ce08e0e850e12f4.
//
// Solidity: event PoolVoted(uint256 id, uint256 cadidateId, string name, address voter, uint256 timestamp)
func (_Voting *VotingFilterer) WatchPoolVoted(opts *bind.WatchOpts, sink chan<- *VotingPoolVoted) (event.Subscription, error) {

	logs, sub, err := _Voting.contract.WatchLogs(opts, "PoolVoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VotingPoolVoted)
				if err := _Voting.contract.UnpackLog(event, "PoolVoted", log); err != nil {
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

// ParsePoolVoted is a log parse operation binding the contract event 0x27e4e9f06dee6980758f4150556c677abc1c6ba9832c4be48ce08e0e850e12f4.
//
// Solidity: event PoolVoted(uint256 id, uint256 cadidateId, string name, address voter, uint256 timestamp)
func (_Voting *VotingFilterer) ParsePoolVoted(log types.Log) (*VotingPoolVoted, error) {
	event := new(VotingPoolVoted)
	if err := _Voting.contract.UnpackLog(event, "PoolVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
