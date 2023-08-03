// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ceresgo

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DataDaoABI is the input ABI used to generate the binding from.
const DataDaoABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DataDaoBin is the compiled bytecode used for deploying new contracts.
var DataDaoBin = "0x608060405234801561001057600080fd5b50610e9d806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806336568abe1161005b57806336568abe146100fe57806391d148541461011a578063a217fddf1461014a578063d547741f146101685761007d565b806301ffc9a714610082578063248a9ca3146100b25780632f2ff15d146100e2575b600080fd5b61009c600480360381019061009791906108e5565b610184565b6040516100a9919061092d565b60405180910390f35b6100cc60048036038101906100c7919061097e565b6101fe565b6040516100d991906109ba565b60405180910390f35b6100fc60048036038101906100f79190610a33565b61021d565b005b61011860048036038101906101139190610a33565b61023e565b005b610134600480360381019061012f9190610a33565b6102c1565b604051610141919061092d565b60405180910390f35b61015261032b565b60405161015f91906109ba565b60405180910390f35b610182600480360381019061017d9190610a33565b610332565b005b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806101f757506101f682610353565b5b9050919050565b6000806000838152602001908152602001600020600101549050919050565b610226826101fe565b61022f816103bd565b61023983836103d1565b505050565b6102466104b1565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146102b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102aa90610af6565b60405180910390fd5b6102bd82826104b9565b5050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6000801b81565b61033b826101fe565b610344816103bd565b61034e83836104b9565b505050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b6103ce816103c96104b1565b61059a565b50565b6103db82826102c1565b6104ad57600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506104526104b1565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600033905090565b6104c382826102c1565b1561059657600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555061053b6104b1565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b6105a482826102c1565b61061b576105b18161061f565b6105bf8360001c602061064c565b6040516020016105d0929190610c1f565b6040516020818303038152906040526040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106129190610ca3565b60405180910390fd5b5050565b60606106458273ffffffffffffffffffffffffffffffffffffffff16601460ff1661064c565b9050919050565b60606000600283600261065f9190610cfe565b6106699190610d40565b67ffffffffffffffff81111561068257610681610d74565b5b6040519080825280601f01601f1916602001820160405280156106b45781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000816000815181106106ec576106eb610da3565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f7800000000000000000000000000000000000000000000000000000000000000816001815181106107505761074f610da3565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600060018460026107909190610cfe565b61079a9190610d40565b90505b600181111561083a577f3031323334353637383961626364656600000000000000000000000000000000600f8616601081106107dc576107db610da3565b5b1a60f81b8282815181106107f3576107f2610da3565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c94508061083390610dd2565b905061079d565b506000841461087e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161087590610e47565b60405180910390fd5b8091505092915050565b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6108c28161088d565b81146108cd57600080fd5b50565b6000813590506108df816108b9565b92915050565b6000602082840312156108fb576108fa610888565b5b6000610909848285016108d0565b91505092915050565b60008115159050919050565b61092781610912565b82525050565b6000602082019050610942600083018461091e565b92915050565b6000819050919050565b61095b81610948565b811461096657600080fd5b50565b60008135905061097881610952565b92915050565b60006020828403121561099457610993610888565b5b60006109a284828501610969565b91505092915050565b6109b481610948565b82525050565b60006020820190506109cf60008301846109ab565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610a00826109d5565b9050919050565b610a10816109f5565b8114610a1b57600080fd5b50565b600081359050610a2d81610a07565b92915050565b60008060408385031215610a4a57610a49610888565b5b6000610a5885828601610969565b9250506020610a6985828601610a1e565b9150509250929050565b600082825260208201905092915050565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b6000610ae0602f83610a73565b9150610aeb82610a84565b604082019050919050565b60006020820190508181036000830152610b0f81610ad3565b9050919050565b600081905092915050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b6000610b57601783610b16565b9150610b6282610b21565b601782019050919050565b600081519050919050565b60005b83811015610b96578082015181840152602081019050610b7b565b60008484015250505050565b6000610bad82610b6d565b610bb78185610b16565b9350610bc7818560208601610b78565b80840191505092915050565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b6000610c09601183610b16565b9150610c1482610bd3565b601182019050919050565b6000610c2a82610b4a565b9150610c368285610ba2565b9150610c4182610bfc565b9150610c4d8284610ba2565b91508190509392505050565b6000601f19601f8301169050919050565b6000610c7582610b6d565b610c7f8185610a73565b9350610c8f818560208601610b78565b610c9881610c59565b840191505092915050565b60006020820190508181036000830152610cbd8184610c6a565b905092915050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610d0982610cc5565b9150610d1483610cc5565b9250828202610d2281610cc5565b91508282048414831517610d3957610d38610ccf565b5b5092915050565b6000610d4b82610cc5565b9150610d5683610cc5565b9250828201905080821115610d6e57610d6d610ccf565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000610ddd82610cc5565b915060008203610df057610def610ccf565b5b600182039050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b6000610e31602083610a73565b9150610e3c82610dfb565b602082019050919050565b60006020820190508181036000830152610e6081610e24565b905091905056fea2646970667358221220b5e73336ce7fd4f68bcbab29a13ca3ec5f469601155458c97d97de0018f6d24064736f6c63430008110033"

// DeployDataDao deploys a new Ethereum contract, binding an instance of DataDao to it.
func DeployDataDao(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DataDao, error) {
	parsed, err := abi.JSON(strings.NewReader(DataDaoABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DataDaoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DataDao{DataDaoCaller: DataDaoCaller{contract: contract}, DataDaoTransactor: DataDaoTransactor{contract: contract}, DataDaoFilterer: DataDaoFilterer{contract: contract}}, nil
}

// DataDao is an auto generated Go binding around an Ethereum contract.
type DataDao struct {
	DataDaoCaller     // Read-only binding to the contract
	DataDaoTransactor // Write-only binding to the contract
	DataDaoFilterer   // Log filterer for contract events
}

// DataDaoCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataDaoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataDaoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataDaoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataDaoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataDaoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataDaoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataDaoSession struct {
	Contract     *DataDao          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataDaoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataDaoCallerSession struct {
	Contract *DataDaoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DataDaoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataDaoTransactorSession struct {
	Contract     *DataDaoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DataDaoRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataDaoRaw struct {
	Contract *DataDao // Generic contract binding to access the raw methods on
}

// DataDaoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataDaoCallerRaw struct {
	Contract *DataDaoCaller // Generic read-only contract binding to access the raw methods on
}

// DataDaoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataDaoTransactorRaw struct {
	Contract *DataDaoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataDao creates a new instance of DataDao, bound to a specific deployed contract.
func NewDataDao(address common.Address, backend bind.ContractBackend) (*DataDao, error) {
	contract, err := bindDataDao(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataDao{DataDaoCaller: DataDaoCaller{contract: contract}, DataDaoTransactor: DataDaoTransactor{contract: contract}, DataDaoFilterer: DataDaoFilterer{contract: contract}}, nil
}

// NewDataDaoCaller creates a new read-only instance of DataDao, bound to a specific deployed contract.
func NewDataDaoCaller(address common.Address, caller bind.ContractCaller) (*DataDaoCaller, error) {
	contract, err := bindDataDao(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataDaoCaller{contract: contract}, nil
}

// NewDataDaoTransactor creates a new write-only instance of DataDao, bound to a specific deployed contract.
func NewDataDaoTransactor(address common.Address, transactor bind.ContractTransactor) (*DataDaoTransactor, error) {
	contract, err := bindDataDao(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataDaoTransactor{contract: contract}, nil
}

// NewDataDaoFilterer creates a new log filterer instance of DataDao, bound to a specific deployed contract.
func NewDataDaoFilterer(address common.Address, filterer bind.ContractFilterer) (*DataDaoFilterer, error) {
	contract, err := bindDataDao(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataDaoFilterer{contract: contract}, nil
}

// bindDataDao binds a generic wrapper to an already deployed contract.
func bindDataDao(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataDaoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataDao *DataDaoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataDao.Contract.DataDaoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataDao *DataDaoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataDao.Contract.DataDaoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataDao *DataDaoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataDao.Contract.DataDaoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataDao *DataDaoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataDao.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataDao *DataDaoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataDao.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataDao *DataDaoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataDao.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DataDao *DataDaoCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DataDao.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DataDao *DataDaoSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DataDao.Contract.DEFAULTADMINROLE(&_DataDao.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DataDao *DataDaoCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DataDao.Contract.DEFAULTADMINROLE(&_DataDao.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DataDao *DataDaoCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _DataDao.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DataDao *DataDaoSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DataDao.Contract.GetRoleAdmin(&_DataDao.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DataDao *DataDaoCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DataDao.Contract.GetRoleAdmin(&_DataDao.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DataDao *DataDaoCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _DataDao.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DataDao *DataDaoSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DataDao.Contract.HasRole(&_DataDao.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DataDao *DataDaoCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DataDao.Contract.HasRole(&_DataDao.CallOpts, role, account)
}