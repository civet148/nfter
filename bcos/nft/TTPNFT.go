// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nft

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/event"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AddressABI is the input ABI used to generate the binding from.
const AddressABI = "[]"

// AddressBin is the compiled bytecode used for deploying new contracts.
var AddressBin = "0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820c9c4158d25cf30e942a372eae35c482b8a41444de3374493d134cae232e763000029"

// DeployAddress deploys a new contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

func AsyncDeployAddress(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Address is an auto generated Go binding around a Solidity contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around a Solidity contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around a Solidity contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around a Solidity contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// ERC165ABI is the input ABI used to generate the binding from.
const ERC165ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ERC165FuncSigs maps the 4-byte function signature to its string representation.
var ERC165FuncSigs = map[string]string{
	"01ffc9a7": "supportsInterface(bytes4)",
}

// ERC165 is an auto generated Go binding around a Solidity contract.
type ERC165 struct {
	ERC165Caller     // Read-only binding to the contract
	ERC165Transactor // Write-only binding to the contract
	ERC165Filterer   // Log filterer for contract events
}

// ERC165Caller is an auto generated read-only Go binding around a Solidity contract.
type ERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Transactor is an auto generated write-only Go binding around a Solidity contract.
type ERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Filterer is an auto generated log filtering Go binding around a Solidity contract events.
type ERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Session is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ERC165Session struct {
	Contract     *ERC165           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165CallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ERC165CallerSession struct {
	Contract *ERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC165TransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ERC165TransactorSession struct {
	Contract     *ERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165Raw is an auto generated low-level Go binding around a Solidity contract.
type ERC165Raw struct {
	Contract *ERC165 // Generic contract binding to access the raw methods on
}

// ERC165CallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ERC165CallerRaw struct {
	Contract *ERC165Caller // Generic read-only contract binding to access the raw methods on
}

// ERC165TransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ERC165TransactorRaw struct {
	Contract *ERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165 creates a new instance of ERC165, bound to a specific deployed contract.
func NewERC165(address common.Address, backend bind.ContractBackend) (*ERC165, error) {
	contract, err := bindERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165{ERC165Caller: ERC165Caller{contract: contract}, ERC165Transactor: ERC165Transactor{contract: contract}, ERC165Filterer: ERC165Filterer{contract: contract}}, nil
}

// NewERC165Caller creates a new read-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Caller(address common.Address, caller bind.ContractCaller) (*ERC165Caller, error) {
	contract, err := bindERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Caller{contract: contract}, nil
}

// NewERC165Transactor creates a new write-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC165Transactor, error) {
	contract, err := bindERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Transactor{contract: contract}, nil
}

// NewERC165Filterer creates a new log filterer instance of ERC165, bound to a specific deployed contract.
func NewERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC165Filterer, error) {
	contract, err := bindERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165Filterer{contract: contract}, nil
}

// bindERC165 binds a generic wrapper to an already deployed contract.
func bindERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.ERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC165 *ERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC165.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC165 *ERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC165 *ERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// ERC721ABI is the input ABI used to generate the binding from.
const ERC721ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721FuncSigs maps the 4-byte function signature to its string representation.
var ERC721FuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC721Bin is the compiled bytecode used for deploying new contracts.
var ERC721Bin = "0x608060405234801561001057600080fd5b506100437f01ffc9a70000000000000000000000000000000000000000000000000000000064010000000061007a810204565b6100757f80ac58cd0000000000000000000000000000000000000000000000000000000064010000000061007a810204565b6100e6565b7fffffffff0000000000000000000000000000000000000000000000000000000080821614156100a957600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b6109a3806100f56000396000f3006080604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301ffc9a781146100a8578063081812fc146100f3578063095ea7b31461012757806323b872dd1461014d57806342842e0e146101775780636352211e146101a157806370a08231146101b9578063a22cb465146101ec578063b88d4fde14610212578063e985e9c514610281575b600080fd5b3480156100b457600080fd5b506100df7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19600435166102a8565b604080519115158252519081900360200190f35b3480156100ff57600080fd5b5061010b6004356102dc565b60408051600160a060020a039092168252519081900360200190f35b34801561013357600080fd5b5061014b600160a060020a036004351660243561030e565b005b34801561015957600080fd5b5061014b600160a060020a03600435811690602435166044356103c4565b34801561018357600080fd5b5061014b600160a060020a0360043581169060243516604435610452565b3480156101ad57600080fd5b5061010b600435610473565b3480156101c557600080fd5b506101da600160a060020a036004351661049d565b60408051918252519081900360200190f35b3480156101f857600080fd5b5061014b600160a060020a036004351660243515156104d0565b34801561021e57600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261014b94600160a060020a0381358116956024803590921695604435953695608494019181908401838280828437509497506105549650505050505050565b34801561028d57600080fd5b506100df600160a060020a036004358116906024351661057c565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191660009081526020819052604090205460ff1690565b60006102e7826105aa565b15156102f257600080fd5b50600090815260026020526040902054600160a060020a031690565b600061031982610473565b9050600160a060020a03838116908216141561033457600080fd5b33600160a060020a03821614806103505750610350813361057c565b151561035b57600080fd5b600082815260026020526040808220805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b6103ce33826105c7565b15156103d957600080fd5b600160a060020a03821615156103ee57600080fd5b6103f88382610626565b6104028382610697565b61040c828261072d565b8082600160a060020a031684600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b61046e8383836020604051908101604052806000815250610554565b505050565b600081815260016020526040812054600160a060020a031680151561049757600080fd5b92915050565b6000600160a060020a03821615156104b457600080fd5b50600160a060020a031660009081526003602052604090205490565b600160a060020a0382163314156104e657600080fd5b336000818152600460209081526040808320600160a060020a03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b61055f8484846103c4565b61056b848484846107bd565b151561057657600080fd5b50505050565b600160a060020a03918216600090815260046020908152604080832093909416825291909152205460ff1690565b600090815260016020526040902054600160a060020a0316151590565b6000806105d383610473565b905080600160a060020a031684600160a060020a0316148061060e575083600160a060020a0316610603846102dc565b600160a060020a0316145b8061061e575061061e818561057c565b949350505050565b81600160a060020a031661063982610473565b600160a060020a03161461064c57600080fd5b600081815260026020526040902054600160a060020a031615610693576000818152600260205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b5050565b81600160a060020a03166106aa82610473565b600160a060020a0316146106bd57600080fd5b600160a060020a0382166000908152600360205260409020546106e790600163ffffffff61093f16565b600160a060020a03909216600090815260036020908152604080832094909455918152600190915220805473ffffffffffffffffffffffffffffffffffffffff19169055565b600081815260016020526040902054600160a060020a03161561074f57600080fd5b6000818152600160208181526040808420805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038816908117909155845260039091529091205461079d91610956565b600160a060020a0390921660009081526003602052604090209190915550565b6000806107d285600160a060020a031661096f565b15156107e15760019150610936565b6040517f150b7a020000000000000000000000000000000000000000000000000000000081523360048201818152600160a060020a03898116602485015260448401889052608060648501908152875160848601528751918a169463150b7a0294938c938b938b93909160a490910190602085019080838360005b8381101561087457818101518382015260200161085c565b50505050905090810190601f1680156108a15780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b1580156108c357600080fd5b505af11580156108d7573d6000803e3d6000fd5b505050506040513d60208110156108ed57600080fd5b50517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1981167f150b7a020000000000000000000000000000000000000000000000000000000014925090505b50949350505050565b6000808383111561094f57600080fd5b5050900390565b60008282018381101561096857600080fd5b9392505050565b6000903b11905600a165627a7a72305820540528645edd9ae08faf4dc9fe25b638f15614fd1c85ca0e04577e2c603cc2b10029"

// DeployERC721 deploys a new contract, binding an instance of ERC721 to it.
func DeployERC721(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721{ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract}}, nil
}

func AsyncDeployERC721(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(ERC721Bin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// ERC721 is an auto generated Go binding around a Solidity contract.
type ERC721 struct {
	ERC721Caller     // Read-only binding to the contract
	ERC721Transactor // Write-only binding to the contract
	ERC721Filterer   // Log filterer for contract events
}

// ERC721Caller is an auto generated read-only Go binding around a Solidity contract.
type ERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Transactor is an auto generated write-only Go binding around a Solidity contract.
type ERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Filterer is an auto generated log filtering Go binding around a Solidity contract events.
type ERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Session is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ERC721Session struct {
	Contract     *ERC721           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721CallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ERC721CallerSession struct {
	Contract *ERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC721TransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ERC721TransactorSession struct {
	Contract     *ERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721Raw is an auto generated low-level Go binding around a Solidity contract.
type ERC721Raw struct {
	Contract *ERC721 // Generic contract binding to access the raw methods on
}

// ERC721CallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ERC721CallerRaw struct {
	Contract *ERC721Caller // Generic read-only contract binding to access the raw methods on
}

// ERC721TransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ERC721TransactorRaw struct {
	Contract *ERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721 creates a new instance of ERC721, bound to a specific deployed contract.
func NewERC721(address common.Address, backend bind.ContractBackend) (*ERC721, error) {
	contract, err := bindERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721{ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract}}, nil
}

// NewERC721Caller creates a new read-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Caller(address common.Address, caller bind.ContractCaller) (*ERC721Caller, error) {
	contract, err := bindERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Caller{contract: contract}, nil
}

// NewERC721Transactor creates a new write-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC721Transactor, error) {
	contract, err := bindERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Transactor{contract: contract}, nil
}

// NewERC721Filterer creates a new log filterer instance of ERC721, bound to a specific deployed contract.
func NewERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC721Filterer, error) {
	contract, err := bindERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721Filterer{contract: contract}, nil
}

// bindERC721 binds a generic wrapper to an already deployed contract.
func bindERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.ERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC721 *ERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC721 *ERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC721 *ERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_ERC721 *ERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_ERC721 *ERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetApproved(&_ERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_ERC721 *ERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetApproved(&_ERC721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_ERC721 *ERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_ERC721 *ERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_ERC721 *ERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_ERC721 *ERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_ERC721 *ERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_ERC721 *ERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC721 *ERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC721 *ERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC721 *ERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721.contract.Transact(opts, "approve", to, tokenId)
}

func (_ERC721 *ERC721Transactor) AsyncApprove(handler func(*types.Receipt, error), opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.AsyncTransact(opts, handler, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, to, tokenId)
}

func (_ERC721 *ERC721Session) AsyncApprove(handler func(*types.Receipt, error), to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.AsyncApprove(handler, &_ERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, to, tokenId)
}

func (_ERC721 *ERC721TransactorSession) AsyncApprove(handler func(*types.Receipt, error), to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.AsyncApprove(handler, &_ERC721.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

func (_ERC721 *ERC721Transactor) AsyncSafeTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.AsyncTransact(opts, handler, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

func (_ERC721 *ERC721Session) AsyncSafeTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.AsyncSafeTransferFrom(handler, &_ERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

func (_ERC721 *ERC721TransactorSession) AsyncSafeTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.AsyncSafeTransferFrom(handler, &_ERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721 *ERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, *types.Receipt, error) {
	return _ERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

func (_ERC721 *ERC721Transactor) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.contract.AsyncTransact(opts, handler, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721 *ERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, *types.Receipt, error) {
	return _ERC721.Contract.SafeTransferFrom0(&_ERC721.TransactOpts, from, to, tokenId, _data)
}

func (_ERC721 *ERC721Session) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.AsyncSafeTransferFrom0(handler, &_ERC721.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721 *ERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, *types.Receipt, error) {
	return _ERC721.Contract.SafeTransferFrom0(&_ERC721.TransactOpts, from, to, tokenId, _data)
}

func (_ERC721 *ERC721TransactorSession) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.AsyncSafeTransferFrom0(handler, &_ERC721.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_ERC721 *ERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, *types.Receipt, error) {
	return _ERC721.contract.Transact(opts, "setApprovalForAll", to, approved)
}

func (_ERC721 *ERC721Transactor) AsyncSetApprovalForAll(handler func(*types.Receipt, error), opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.contract.AsyncTransact(opts, handler, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_ERC721 *ERC721Session) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, *types.Receipt, error) {
	return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts, to, approved)
}

func (_ERC721 *ERC721Session) AsyncSetApprovalForAll(handler func(*types.Receipt, error), to common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.Contract.AsyncSetApprovalForAll(handler, &_ERC721.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_ERC721 *ERC721TransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, *types.Receipt, error) {
	return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts, to, approved)
}

func (_ERC721 *ERC721TransactorSession) AsyncSetApprovalForAll(handler func(*types.Receipt, error), to common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721.Contract.AsyncSetApprovalForAll(handler, &_ERC721.TransactOpts, to, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

func (_ERC721 *ERC721Transactor) AsyncTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.AsyncTransact(opts, handler, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

func (_ERC721 *ERC721Session) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.AsyncTransferFrom(handler, &_ERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721 *ERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, from, to, tokenId)
}

func (_ERC721 *ERC721TransactorSession) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.AsyncTransferFrom(handler, &_ERC721.TransactOpts, from, to, tokenId)
}

// ERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721 contract.
type ERC721ApprovalIterator struct {
	Event *ERC721Approval // Event containing the contract specifics and raw log

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
func (it *ERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Approval)
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
		it.Event = new(ERC721Approval)
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
func (it *ERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Approval represents a Approval event raised by the ERC721 contract.
type ERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalIterator{contract: _ERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Approval)
				if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) ParseApproval(log types.Log) (*ERC721Approval, error) {
	event := new(ERC721Approval)
	if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721 contract.
type ERC721ApprovalForAllIterator struct {
	Event *ERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721ApprovalForAll)
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
		it.Event = new(ERC721ApprovalForAll)
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
func (it *ERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721ApprovalForAll represents a ApprovalForAll event raised by the ERC721 contract.
type ERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721 *ERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalForAllIterator{contract: _ERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721 *ERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721ApprovalForAll)
				if err := _ERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721 *ERC721Filterer) ParseApprovalForAll(log types.Log) (*ERC721ApprovalForAll, error) {
	event := new(ERC721ApprovalForAll)
	if err := _ERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721 contract.
type ERC721TransferIterator struct {
	Event *ERC721Transfer // Event containing the contract specifics and raw log

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
func (it *ERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Transfer)
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
		it.Event = new(ERC721Transfer)
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
func (it *ERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Transfer represents a Transfer event raised by the ERC721 contract.
type ERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TransferIterator{contract: _ERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Transfer)
				if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721 *ERC721Filterer) ParseTransfer(log types.Log) (*ERC721Transfer, error) {
	event := new(ERC721Transfer)
	if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC721EnumerableABI is the input ABI used to generate the binding from.
const ERC721EnumerableABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721EnumerableFuncSigs maps the 4-byte function signature to its string representation.
var ERC721EnumerableFuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"4f6ccce7": "tokenByIndex(uint256)",
	"2f745c59": "tokenOfOwnerByIndex(address,uint256)",
	"18160ddd": "totalSupply()",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC721EnumerableBin is the compiled bytecode used for deploying new contracts.
var ERC721EnumerableBin = "0x608060405234801561001057600080fd5b506100437f01ffc9a7000000000000000000000000000000000000000000000000000000006401000000006100ac810204565b6100757f80ac58cd000000000000000000000000000000000000000000000000000000006401000000006100ac810204565b6100a77f780e9d63000000000000000000000000000000000000000000000000000000006401000000006100ac810204565b610118565b7fffffffff0000000000000000000000000000000000000000000000000000000080821614156100db57600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b610c2b806101276000396000f3006080604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301ffc9a781146100c9578063081812fc14610114578063095ea7b31461014857806318160ddd1461016e57806323b872dd146101955780632f745c59146101bf57806342842e0e146101e35780634f6ccce71461020d5780636352211e1461022557806370a082311461023d578063a22cb4651461025e578063b88d4fde14610284578063e985e9c5146102f3575b600080fd5b3480156100d557600080fd5b506101007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff196004351661031a565b604080519115158252519081900360200190f35b34801561012057600080fd5b5061012c60043561034e565b60408051600160a060020a039092168252519081900360200190f35b34801561015457600080fd5b5061016c600160a060020a0360043516602435610380565b005b34801561017a57600080fd5b50610183610436565b60408051918252519081900360200190f35b3480156101a157600080fd5b5061016c600160a060020a036004358116906024351660443561043d565b3480156101cb57600080fd5b50610183600160a060020a03600435166024356104cb565b3480156101ef57600080fd5b5061016c600160a060020a0360043581169060243516604435610518565b34801561021957600080fd5b50610183600435610539565b34801561023157600080fd5b5061012c60043561056e565b34801561024957600080fd5b50610183600160a060020a0360043516610598565b34801561026a57600080fd5b5061016c600160a060020a036004351660243515156105cb565b34801561029057600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261016c94600160a060020a03813581169560248035909216956044359536956084940191819084018382808284375094975061064f9650505050505050565b3480156102ff57600080fd5b50610100600160a060020a0360043581169060243516610677565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191660009081526020819052604090205460ff1690565b6000610359826106a5565b151561036457600080fd5b50600090815260026020526040902054600160a060020a031690565b600061038b8261056e565b9050600160a060020a0383811690821614156103a657600080fd5b33600160a060020a03821614806103c257506103c28133610677565b15156103cd57600080fd5b600082815260026020526040808220805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b6007545b90565b61044733826106c2565b151561045257600080fd5b600160a060020a038216151561046757600080fd5b6104718382610721565b61047b8382610792565b6104858282610899565b8082600160a060020a031684600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b60006104d683610598565b82106104e157600080fd5b600160a060020a038316600090815260056020526040902080548390811061050557fe5b9060005260206000200154905092915050565b610534838383602060405190810160405280600081525061064f565b505050565b6000610543610436565b821061054e57600080fd5b600780548390811061055c57fe5b90600052602060002001549050919050565b600081815260016020526040812054600160a060020a031680151561059257600080fd5b92915050565b6000600160a060020a03821615156105af57600080fd5b50600160a060020a031660009081526003602052604090205490565b600160a060020a0382163314156105e157600080fd5b336000818152600460209081526040808320600160a060020a03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b61065a84848461043d565b610666848484846108e2565b151561067157600080fd5b50505050565b600160a060020a03918216600090815260046020908152604080832093909416825291909152205460ff1690565b600090815260016020526040902054600160a060020a0316151590565b6000806106ce8361056e565b905080600160a060020a031684600160a060020a03161480610709575083600160a060020a03166106fe8461034e565b600160a060020a0316145b8061071957506107198185610677565b949350505050565b81600160a060020a03166107348261056e565b600160a060020a03161461074757600080fd5b600081815260026020526040902054600160a060020a03161561078e576000818152600260205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b5050565b60008060006107a18585610a64565b600084815260066020908152604080832054600160a060020a03891684526005909252909120549093506107dc90600163ffffffff610afa16565b600160a060020a03861660009081526005602052604090208054919350908390811061080457fe5b90600052602060002001549050806005600087600160a060020a0316600160a060020a031681526020019081526020016000208481548110151561084457fe5b6000918252602080832090910192909255600160a060020a038716815260059091526040902080549061087b906000198301610bc2565b50600093845260066020526040808520859055908452909220555050565b60006108a58383610b11565b50600160a060020a039091166000908152600560209081526040808320805460018101825590845282842081018590559383526006909152902055565b6000806108f785600160a060020a0316610ba1565b15156109065760019150610a5b565b6040517f150b7a020000000000000000000000000000000000000000000000000000000081523360048201818152600160a060020a03898116602485015260448401889052608060648501908152875160848601528751918a169463150b7a0294938c938b938b93909160a490910190602085019080838360005b83811015610999578181015183820152602001610981565b50505050905090810190601f1680156109c65780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b1580156109e857600080fd5b505af11580156109fc573d6000803e3d6000fd5b505050506040513d6020811015610a1257600080fd5b50517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1981167f150b7a020000000000000000000000000000000000000000000000000000000014925090505b50949350505050565b81600160a060020a0316610a778261056e565b600160a060020a031614610a8a57600080fd5b600160a060020a038216600090815260036020526040902054610ab490600163ffffffff610afa16565b600160a060020a03909216600090815260036020908152604080832094909455918152600190915220805473ffffffffffffffffffffffffffffffffffffffff19169055565b60008083831115610b0a57600080fd5b5050900390565b600081815260016020526040902054600160a060020a031615610b3357600080fd5b6000818152600160208181526040808420805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0388169081179091558452600390915290912054610b8191610ba9565b600160a060020a0390921660009081526003602052604090209190915550565b6000903b1190565b600082820183811015610bbb57600080fd5b9392505050565b8154818355818111156105345760008381526020902061053491810190830161043a91905b80821115610bfb5760008155600101610be7565b50905600a165627a7a72305820bf0cdbac3aaf39239b29492a0ac88e041fd47c9e7fbe6ebc45ab8bd0f9b1728c0029"

// DeployERC721Enumerable deploys a new contract, binding an instance of ERC721Enumerable to it.
func DeployERC721Enumerable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721Enumerable, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721EnumerableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721EnumerableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Enumerable{ERC721EnumerableCaller: ERC721EnumerableCaller{contract: contract}, ERC721EnumerableTransactor: ERC721EnumerableTransactor{contract: contract}, ERC721EnumerableFilterer: ERC721EnumerableFilterer{contract: contract}}, nil
}

func AsyncDeployERC721Enumerable(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721EnumerableABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(ERC721EnumerableBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// ERC721Enumerable is an auto generated Go binding around a Solidity contract.
type ERC721Enumerable struct {
	ERC721EnumerableCaller     // Read-only binding to the contract
	ERC721EnumerableTransactor // Write-only binding to the contract
	ERC721EnumerableFilterer   // Log filterer for contract events
}

// ERC721EnumerableCaller is an auto generated read-only Go binding around a Solidity contract.
type ERC721EnumerableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721EnumerableTransactor is an auto generated write-only Go binding around a Solidity contract.
type ERC721EnumerableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721EnumerableFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ERC721EnumerableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721EnumerableSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ERC721EnumerableSession struct {
	Contract     *ERC721Enumerable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721EnumerableCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ERC721EnumerableCallerSession struct {
	Contract *ERC721EnumerableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC721EnumerableTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ERC721EnumerableTransactorSession struct {
	Contract     *ERC721EnumerableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC721EnumerableRaw is an auto generated low-level Go binding around a Solidity contract.
type ERC721EnumerableRaw struct {
	Contract *ERC721Enumerable // Generic contract binding to access the raw methods on
}

// ERC721EnumerableCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ERC721EnumerableCallerRaw struct {
	Contract *ERC721EnumerableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721EnumerableTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ERC721EnumerableTransactorRaw struct {
	Contract *ERC721EnumerableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Enumerable creates a new instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721Enumerable(address common.Address, backend bind.ContractBackend) (*ERC721Enumerable, error) {
	contract, err := bindERC721Enumerable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Enumerable{ERC721EnumerableCaller: ERC721EnumerableCaller{contract: contract}, ERC721EnumerableTransactor: ERC721EnumerableTransactor{contract: contract}, ERC721EnumerableFilterer: ERC721EnumerableFilterer{contract: contract}}, nil
}

// NewERC721EnumerableCaller creates a new read-only instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721EnumerableCaller(address common.Address, caller bind.ContractCaller) (*ERC721EnumerableCaller, error) {
	contract, err := bindERC721Enumerable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableCaller{contract: contract}, nil
}

// NewERC721EnumerableTransactor creates a new write-only instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721EnumerableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721EnumerableTransactor, error) {
	contract, err := bindERC721Enumerable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableTransactor{contract: contract}, nil
}

// NewERC721EnumerableFilterer creates a new log filterer instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721EnumerableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721EnumerableFilterer, error) {
	contract, err := bindERC721Enumerable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableFilterer{contract: contract}, nil
}

// bindERC721Enumerable binds a generic wrapper to an already deployed contract.
func bindERC721Enumerable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721EnumerableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Enumerable *ERC721EnumerableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Enumerable.Contract.ERC721EnumerableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Enumerable *ERC721EnumerableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Enumerable.Contract.ERC721EnumerableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Enumerable *ERC721EnumerableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Enumerable.Contract.ERC721EnumerableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Enumerable *ERC721EnumerableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Enumerable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Enumerable *ERC721EnumerableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Enumerable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Enumerable *ERC721EnumerableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Enumerable.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721Enumerable.Contract.BalanceOf(&_ERC721Enumerable.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721Enumerable.Contract.BalanceOf(&_ERC721Enumerable.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_ERC721Enumerable *ERC721EnumerableCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_ERC721Enumerable *ERC721EnumerableSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.GetApproved(&_ERC721Enumerable.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.GetApproved(&_ERC721Enumerable.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_ERC721Enumerable *ERC721EnumerableCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_ERC721Enumerable *ERC721EnumerableSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721Enumerable.Contract.IsApprovedForAll(&_ERC721Enumerable.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721Enumerable.Contract.IsApprovedForAll(&_ERC721Enumerable.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_ERC721Enumerable *ERC721EnumerableCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_ERC721Enumerable *ERC721EnumerableSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.OwnerOf(&_ERC721Enumerable.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.OwnerOf(&_ERC721Enumerable.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC721Enumerable *ERC721EnumerableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC721Enumerable *ERC721EnumerableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721Enumerable.Contract.SupportsInterface(&_ERC721Enumerable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721Enumerable.Contract.SupportsInterface(&_ERC721Enumerable.CallOpts, interfaceId)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "tokenByIndex", index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenByIndex(&_ERC721Enumerable.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenByIndex(&_ERC721Enumerable.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "tokenOfOwnerByIndex", owner, index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenOfOwnerByIndex(&_ERC721Enumerable.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenOfOwnerByIndex(&_ERC721Enumerable.CallOpts, owner, index)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableSession) TotalSupply() (*big.Int, error) {
	return _ERC721Enumerable.Contract.TotalSupply(&_ERC721Enumerable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721Enumerable.Contract.TotalSupply(&_ERC721Enumerable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Enumerable.contract.Transact(opts, "approve", to, tokenId)
}

func (_ERC721Enumerable *ERC721EnumerableTransactor) AsyncApprove(handler func(*types.Receipt, error), opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.AsyncTransact(opts, handler, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Enumerable.Contract.Approve(&_ERC721Enumerable.TransactOpts, to, tokenId)
}

func (_ERC721Enumerable *ERC721EnumerableSession) AsyncApprove(handler func(*types.Receipt, error), to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.AsyncApprove(handler, &_ERC721Enumerable.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Enumerable.Contract.Approve(&_ERC721Enumerable.TransactOpts, to, tokenId)
}

func (_ERC721Enumerable *ERC721EnumerableTransactorSession) AsyncApprove(handler func(*types.Receipt, error), to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.AsyncApprove(handler, &_ERC721Enumerable.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Enumerable.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

func (_ERC721Enumerable *ERC721EnumerableTransactor) AsyncSafeTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.AsyncTransact(opts, handler, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Enumerable.Contract.SafeTransferFrom(&_ERC721Enumerable.TransactOpts, from, to, tokenId)
}

func (_ERC721Enumerable *ERC721EnumerableSession) AsyncSafeTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.AsyncSafeTransferFrom(handler, &_ERC721Enumerable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Enumerable.Contract.SafeTransferFrom(&_ERC721Enumerable.TransactOpts, from, to, tokenId)
}

func (_ERC721Enumerable *ERC721EnumerableTransactorSession) AsyncSafeTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.AsyncSafeTransferFrom(handler, &_ERC721Enumerable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Enumerable.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

func (_ERC721Enumerable *ERC721EnumerableTransactor) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.AsyncTransact(opts, handler, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Enumerable.Contract.SafeTransferFrom0(&_ERC721Enumerable.TransactOpts, from, to, tokenId, _data)
}

func (_ERC721Enumerable *ERC721EnumerableSession) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.AsyncSafeTransferFrom0(handler, &_ERC721Enumerable.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Enumerable.Contract.SafeTransferFrom0(&_ERC721Enumerable.TransactOpts, from, to, tokenId, _data)
}

func (_ERC721Enumerable *ERC721EnumerableTransactorSession) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.AsyncSafeTransferFrom0(handler, &_ERC721Enumerable.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Enumerable.contract.Transact(opts, "setApprovalForAll", to, approved)
}

func (_ERC721Enumerable *ERC721EnumerableTransactor) AsyncSetApprovalForAll(handler func(*types.Receipt, error), opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.AsyncTransact(opts, handler, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Enumerable.Contract.SetApprovalForAll(&_ERC721Enumerable.TransactOpts, to, approved)
}

func (_ERC721Enumerable *ERC721EnumerableSession) AsyncSetApprovalForAll(handler func(*types.Receipt, error), to common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.AsyncSetApprovalForAll(handler, &_ERC721Enumerable.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Enumerable.Contract.SetApprovalForAll(&_ERC721Enumerable.TransactOpts, to, approved)
}

func (_ERC721Enumerable *ERC721EnumerableTransactorSession) AsyncSetApprovalForAll(handler func(*types.Receipt, error), to common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.AsyncSetApprovalForAll(handler, &_ERC721Enumerable.TransactOpts, to, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Enumerable.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

func (_ERC721Enumerable *ERC721EnumerableTransactor) AsyncTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.AsyncTransact(opts, handler, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Enumerable.Contract.TransferFrom(&_ERC721Enumerable.TransactOpts, from, to, tokenId)
}

func (_ERC721Enumerable *ERC721EnumerableSession) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.AsyncTransferFrom(handler, &_ERC721Enumerable.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Enumerable.Contract.TransferFrom(&_ERC721Enumerable.TransactOpts, from, to, tokenId)
}

func (_ERC721Enumerable *ERC721EnumerableTransactorSession) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.AsyncTransferFrom(handler, &_ERC721Enumerable.TransactOpts, from, to, tokenId)
}

// ERC721EnumerableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Enumerable contract.
type ERC721EnumerableApprovalIterator struct {
	Event *ERC721EnumerableApproval // Event containing the contract specifics and raw log

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
func (it *ERC721EnumerableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721EnumerableApproval)
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
		it.Event = new(ERC721EnumerableApproval)
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
func (it *ERC721EnumerableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721EnumerableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721EnumerableApproval represents a Approval event raised by the ERC721Enumerable contract.
type ERC721EnumerableApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Enumerable *ERC721EnumerableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721EnumerableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableApprovalIterator{contract: _ERC721Enumerable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Enumerable *ERC721EnumerableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721EnumerableApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721EnumerableApproval)
				if err := _ERC721Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Enumerable *ERC721EnumerableFilterer) ParseApproval(log types.Log) (*ERC721EnumerableApproval, error) {
	event := new(ERC721EnumerableApproval)
	if err := _ERC721Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC721EnumerableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Enumerable contract.
type ERC721EnumerableApprovalForAllIterator struct {
	Event *ERC721EnumerableApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721EnumerableApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721EnumerableApprovalForAll)
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
		it.Event = new(ERC721EnumerableApprovalForAll)
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
func (it *ERC721EnumerableApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721EnumerableApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721EnumerableApprovalForAll represents a ApprovalForAll event raised by the ERC721Enumerable contract.
type ERC721EnumerableApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721Enumerable *ERC721EnumerableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721EnumerableApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableApprovalForAllIterator{contract: _ERC721Enumerable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721Enumerable *ERC721EnumerableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721EnumerableApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721EnumerableApprovalForAll)
				if err := _ERC721Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721Enumerable *ERC721EnumerableFilterer) ParseApprovalForAll(log types.Log) (*ERC721EnumerableApprovalForAll, error) {
	event := new(ERC721EnumerableApprovalForAll)
	if err := _ERC721Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC721EnumerableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Enumerable contract.
type ERC721EnumerableTransferIterator struct {
	Event *ERC721EnumerableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721EnumerableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721EnumerableTransfer)
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
		it.Event = new(ERC721EnumerableTransfer)
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
func (it *ERC721EnumerableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721EnumerableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721EnumerableTransfer represents a Transfer event raised by the ERC721Enumerable contract.
type ERC721EnumerableTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Enumerable *ERC721EnumerableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721EnumerableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableTransferIterator{contract: _ERC721Enumerable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Enumerable *ERC721EnumerableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721EnumerableTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721EnumerableTransfer)
				if err := _ERC721Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Enumerable *ERC721EnumerableFilterer) ParseTransfer(log types.Log) (*ERC721EnumerableTransfer, error) {
	event := new(ERC721EnumerableTransfer)
	if err := _ERC721Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC721FullABI is the input ABI used to generate the binding from.
const ERC721FullABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721FullFuncSigs maps the 4-byte function signature to its string representation.
var ERC721FullFuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"06fdde03": "name()",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"95d89b41": "symbol()",
	"4f6ccce7": "tokenByIndex(uint256)",
	"2f745c59": "tokenOfOwnerByIndex(address,uint256)",
	"c87b56dd": "tokenURI(uint256)",
	"18160ddd": "totalSupply()",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC721FullBin is the compiled bytecode used for deploying new contracts.
var ERC721FullBin = "0x60806040523480156200001157600080fd5b506040516200110c3803806200110c833981016040528051602082015190820191018181620000697f01ffc9a7000000000000000000000000000000000000000000000000000000006401000000006200013b810204565b6200009d7f80ac58cd000000000000000000000000000000000000000000000000000000006401000000006200013b810204565b620000d17f780e9d63000000000000000000000000000000000000000000000000000000006401000000006200013b810204565b8151620000e6906009906020850190620001a8565b508051620000fc90600a906020840190620001a8565b50620001317f5b5e139f000000000000000000000000000000000000000000000000000000006401000000006200013b810204565b505050506200024d565b7fffffffff0000000000000000000000000000000000000000000000000000000080821614156200016b57600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001eb57805160ff19168380011785556200021b565b828001600101855582156200021b579182015b828111156200021b578251825591602001919060010190620001fe565b50620002299291506200022d565b5090565b6200024a91905b8082111562000229576000815560010162000234565b90565b610eaf806200025d6000396000f3006080604052600436106100e55763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301ffc9a781146100ea57806306fdde0314610135578063081812fc146101bf578063095ea7b3146101f357806318160ddd1461021957806323b872dd146102405780632f745c591461026a57806342842e0e1461028e5780634f6ccce7146102b85780636352211e146102d057806370a08231146102e857806395d89b4114610309578063a22cb4651461031e578063b88d4fde14610344578063c87b56dd146103b3578063e985e9c5146103cb575b600080fd5b3480156100f657600080fd5b506101217bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19600435166103f2565b604080519115158252519081900360200190f35b34801561014157600080fd5b5061014a610426565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561018457818101518382015260200161016c565b50505050905090810190601f1680156101b15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101cb57600080fd5b506101d76004356104bd565b60408051600160a060020a039092168252519081900360200190f35b3480156101ff57600080fd5b50610217600160a060020a03600435166024356104ef565b005b34801561022557600080fd5b5061022e6105a5565b60408051918252519081900360200190f35b34801561024c57600080fd5b50610217600160a060020a03600435811690602435166044356105ab565b34801561027657600080fd5b5061022e600160a060020a0360043516602435610639565b34801561029a57600080fd5b50610217600160a060020a0360043581169060243516604435610686565b3480156102c457600080fd5b5061022e6004356106a7565b3480156102dc57600080fd5b506101d76004356106dc565b3480156102f457600080fd5b5061022e600160a060020a0360043516610706565b34801561031557600080fd5b5061014a610739565b34801561032a57600080fd5b50610217600160a060020a0360043516602435151561079a565b34801561035057600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261021794600160a060020a03813581169560248035909216956044359536956084940191819084018382808284375094975061081e9650505050505050565b3480156103bf57600080fd5b5061014a600435610846565b3480156103d757600080fd5b50610121600160a060020a03600435811690602435166108fb565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191660009081526020819052604090205460ff1690565b60098054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104b25780601f10610487576101008083540402835291602001916104b2565b820191906000526020600020905b81548152906001019060200180831161049557829003601f168201915b505050505090505b90565b60006104c882610929565b15156104d357600080fd5b50600090815260026020526040902054600160a060020a031690565b60006104fa826106dc565b9050600160a060020a03838116908216141561051557600080fd5b33600160a060020a0382161480610531575061053181336108fb565b151561053c57600080fd5b600082815260026020526040808220805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b60075490565b6105b53382610946565b15156105c057600080fd5b600160a060020a03821615156105d557600080fd5b6105df83826109a5565b6105e98382610a16565b6105f38282610b1d565b8082600160a060020a031684600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b600061064483610706565b821061064f57600080fd5b600160a060020a038316600090815260056020526040902080548390811061067357fe5b9060005260206000200154905092915050565b6106a2838383602060405190810160405280600081525061081e565b505050565b60006106b16105a5565b82106106bc57600080fd5b60078054839081106106ca57fe5b90600052602060002001549050919050565b600081815260016020526040812054600160a060020a031680151561070057600080fd5b92915050565b6000600160a060020a038216151561071d57600080fd5b50600160a060020a031660009081526003602052604090205490565b600a8054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104b25780601f10610487576101008083540402835291602001916104b2565b600160a060020a0382163314156107b057600080fd5b336000818152600460209081526040808320600160a060020a03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b6108298484846105ab565b61083584848484610b66565b151561084057600080fd5b50505050565b606061085182610929565b151561085c57600080fd5b6000828152600b602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156108ef5780601f106108c4576101008083540402835291602001916108ef565b820191906000526020600020905b8154815290600101906020018083116108d257829003601f168201915b50505050509050919050565b600160a060020a03918216600090815260046020908152604080832093909416825291909152205460ff1690565b600090815260016020526040902054600160a060020a0316151590565b600080610952836106dc565b905080600160a060020a031684600160a060020a0316148061098d575083600160a060020a0316610982846104bd565b600160a060020a0316145b8061099d575061099d81856108fb565b949350505050565b81600160a060020a03166109b8826106dc565b600160a060020a0316146109cb57600080fd5b600081815260026020526040902054600160a060020a031615610a12576000818152600260205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b5050565b6000806000610a258585610ce8565b600084815260066020908152604080832054600160a060020a0389168452600590925290912054909350610a6090600163ffffffff610d7e16565b600160a060020a038616600090815260056020526040902080549193509083908110610a8857fe5b90600052602060002001549050806005600087600160a060020a0316600160a060020a0316815260200190815260200160002084815481101515610ac857fe5b6000918252602080832090910192909255600160a060020a0387168152600590915260409020805490610aff906000198301610e46565b50600093845260066020526040808520859055908452909220555050565b6000610b298383610d95565b50600160a060020a039091166000908152600560209081526040808320805460018101825590845282842081018590559383526006909152902055565b600080610b7b85600160a060020a0316610e25565b1515610b8a5760019150610cdf565b6040517f150b7a020000000000000000000000000000000000000000000000000000000081523360048201818152600160a060020a03898116602485015260448401889052608060648501908152875160848601528751918a169463150b7a0294938c938b938b93909160a490910190602085019080838360005b83811015610c1d578181015183820152602001610c05565b50505050905090810190601f168015610c4a5780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015610c6c57600080fd5b505af1158015610c80573d6000803e3d6000fd5b505050506040513d6020811015610c9657600080fd5b50517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1981167f150b7a020000000000000000000000000000000000000000000000000000000014925090505b50949350505050565b81600160a060020a0316610cfb826106dc565b600160a060020a031614610d0e57600080fd5b600160a060020a038216600090815260036020526040902054610d3890600163ffffffff610d7e16565b600160a060020a03909216600090815260036020908152604080832094909455918152600190915220805473ffffffffffffffffffffffffffffffffffffffff19169055565b60008083831115610d8e57600080fd5b5050900390565b600081815260016020526040902054600160a060020a031615610db757600080fd5b6000818152600160208181526040808420805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0388169081179091558452600390915290912054610e0591610e2d565b600160a060020a0390921660009081526003602052604090209190915550565b6000903b1190565b600082820183811015610e3f57600080fd5b9392505050565b8154818355818111156106a2576000838152602090206106a29181019083016104ba91905b80821115610e7f5760008155600101610e6b565b50905600a165627a7a7230582070fb993b9dc395fa035206eae806cf3cf2decd5d23476fbb7c94edd484f030030029"

// DeployERC721Full deploys a new contract, binding an instance of ERC721Full to it.
func DeployERC721Full(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *ERC721Full, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721FullABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721FullBin), backend, name, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Full{ERC721FullCaller: ERC721FullCaller{contract: contract}, ERC721FullTransactor: ERC721FullTransactor{contract: contract}, ERC721FullFilterer: ERC721FullFilterer{contract: contract}}, nil
}

func AsyncDeployERC721Full(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend, name string, symbol string) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721FullABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(ERC721FullBin), backend, name, symbol)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// ERC721Full is an auto generated Go binding around a Solidity contract.
type ERC721Full struct {
	ERC721FullCaller     // Read-only binding to the contract
	ERC721FullTransactor // Write-only binding to the contract
	ERC721FullFilterer   // Log filterer for contract events
}

// ERC721FullCaller is an auto generated read-only Go binding around a Solidity contract.
type ERC721FullCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721FullTransactor is an auto generated write-only Go binding around a Solidity contract.
type ERC721FullTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721FullFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ERC721FullFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721FullSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ERC721FullSession struct {
	Contract     *ERC721Full       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721FullCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ERC721FullCallerSession struct {
	Contract *ERC721FullCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC721FullTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ERC721FullTransactorSession struct {
	Contract     *ERC721FullTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC721FullRaw is an auto generated low-level Go binding around a Solidity contract.
type ERC721FullRaw struct {
	Contract *ERC721Full // Generic contract binding to access the raw methods on
}

// ERC721FullCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ERC721FullCallerRaw struct {
	Contract *ERC721FullCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721FullTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ERC721FullTransactorRaw struct {
	Contract *ERC721FullTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Full creates a new instance of ERC721Full, bound to a specific deployed contract.
func NewERC721Full(address common.Address, backend bind.ContractBackend) (*ERC721Full, error) {
	contract, err := bindERC721Full(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Full{ERC721FullCaller: ERC721FullCaller{contract: contract}, ERC721FullTransactor: ERC721FullTransactor{contract: contract}, ERC721FullFilterer: ERC721FullFilterer{contract: contract}}, nil
}

// NewERC721FullCaller creates a new read-only instance of ERC721Full, bound to a specific deployed contract.
func NewERC721FullCaller(address common.Address, caller bind.ContractCaller) (*ERC721FullCaller, error) {
	contract, err := bindERC721Full(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721FullCaller{contract: contract}, nil
}

// NewERC721FullTransactor creates a new write-only instance of ERC721Full, bound to a specific deployed contract.
func NewERC721FullTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721FullTransactor, error) {
	contract, err := bindERC721Full(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721FullTransactor{contract: contract}, nil
}

// NewERC721FullFilterer creates a new log filterer instance of ERC721Full, bound to a specific deployed contract.
func NewERC721FullFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721FullFilterer, error) {
	contract, err := bindERC721Full(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721FullFilterer{contract: contract}, nil
}

// bindERC721Full binds a generic wrapper to an already deployed contract.
func bindERC721Full(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721FullABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Full *ERC721FullRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Full.Contract.ERC721FullCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Full *ERC721FullRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Full.Contract.ERC721FullTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Full *ERC721FullRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Full.Contract.ERC721FullTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Full *ERC721FullCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Full.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Full *ERC721FullTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Full.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Full *ERC721FullTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Full.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC721Full *ERC721FullCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Full.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC721Full *ERC721FullSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721Full.Contract.BalanceOf(&_ERC721Full.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC721Full *ERC721FullCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721Full.Contract.BalanceOf(&_ERC721Full.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_ERC721Full *ERC721FullCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Full.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_ERC721Full *ERC721FullSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721Full.Contract.GetApproved(&_ERC721Full.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_ERC721Full *ERC721FullCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721Full.Contract.GetApproved(&_ERC721Full.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_ERC721Full *ERC721FullCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Full.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_ERC721Full *ERC721FullSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721Full.Contract.IsApprovedForAll(&_ERC721Full.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_ERC721Full *ERC721FullCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721Full.Contract.IsApprovedForAll(&_ERC721Full.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC721Full *ERC721FullCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721Full.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC721Full *ERC721FullSession) Name() (string, error) {
	return _ERC721Full.Contract.Name(&_ERC721Full.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC721Full *ERC721FullCallerSession) Name() (string, error) {
	return _ERC721Full.Contract.Name(&_ERC721Full.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_ERC721Full *ERC721FullCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Full.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_ERC721Full *ERC721FullSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721Full.Contract.OwnerOf(&_ERC721Full.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_ERC721Full *ERC721FullCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721Full.Contract.OwnerOf(&_ERC721Full.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC721Full *ERC721FullCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Full.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC721Full *ERC721FullSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721Full.Contract.SupportsInterface(&_ERC721Full.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC721Full *ERC721FullCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721Full.Contract.SupportsInterface(&_ERC721Full.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC721Full *ERC721FullCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721Full.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC721Full *ERC721FullSession) Symbol() (string, error) {
	return _ERC721Full.Contract.Symbol(&_ERC721Full.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC721Full *ERC721FullCallerSession) Symbol() (string, error) {
	return _ERC721Full.Contract.Symbol(&_ERC721Full.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) constant returns(uint256)
func (_ERC721Full *ERC721FullCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Full.contract.Call(opts, out, "tokenByIndex", index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) constant returns(uint256)
func (_ERC721Full *ERC721FullSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ERC721Full.Contract.TokenByIndex(&_ERC721Full.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) constant returns(uint256)
func (_ERC721Full *ERC721FullCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ERC721Full.Contract.TokenByIndex(&_ERC721Full.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) constant returns(uint256)
func (_ERC721Full *ERC721FullCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Full.contract.Call(opts, out, "tokenOfOwnerByIndex", owner, index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) constant returns(uint256)
func (_ERC721Full *ERC721FullSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ERC721Full.Contract.TokenOfOwnerByIndex(&_ERC721Full.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) constant returns(uint256)
func (_ERC721Full *ERC721FullCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ERC721Full.Contract.TokenOfOwnerByIndex(&_ERC721Full.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) constant returns(string)
func (_ERC721Full *ERC721FullCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721Full.contract.Call(opts, out, "tokenURI", tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) constant returns(string)
func (_ERC721Full *ERC721FullSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721Full.Contract.TokenURI(&_ERC721Full.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) constant returns(string)
func (_ERC721Full *ERC721FullCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721Full.Contract.TokenURI(&_ERC721Full.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721Full *ERC721FullCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Full.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721Full *ERC721FullSession) TotalSupply() (*big.Int, error) {
	return _ERC721Full.Contract.TotalSupply(&_ERC721Full.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721Full *ERC721FullCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721Full.Contract.TotalSupply(&_ERC721Full.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Full *ERC721FullTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Full.contract.Transact(opts, "approve", to, tokenId)
}

func (_ERC721Full *ERC721FullTransactor) AsyncApprove(handler func(*types.Receipt, error), opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Full.contract.AsyncTransact(opts, handler, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Full *ERC721FullSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Full.Contract.Approve(&_ERC721Full.TransactOpts, to, tokenId)
}

func (_ERC721Full *ERC721FullSession) AsyncApprove(handler func(*types.Receipt, error), to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Full.Contract.AsyncApprove(handler, &_ERC721Full.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Full *ERC721FullTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Full.Contract.Approve(&_ERC721Full.TransactOpts, to, tokenId)
}

func (_ERC721Full *ERC721FullTransactorSession) AsyncApprove(handler func(*types.Receipt, error), to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Full.Contract.AsyncApprove(handler, &_ERC721Full.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Full *ERC721FullTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Full.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

func (_ERC721Full *ERC721FullTransactor) AsyncSafeTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Full.contract.AsyncTransact(opts, handler, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Full *ERC721FullSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Full.Contract.SafeTransferFrom(&_ERC721Full.TransactOpts, from, to, tokenId)
}

func (_ERC721Full *ERC721FullSession) AsyncSafeTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Full.Contract.AsyncSafeTransferFrom(handler, &_ERC721Full.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Full *ERC721FullTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Full.Contract.SafeTransferFrom(&_ERC721Full.TransactOpts, from, to, tokenId)
}

func (_ERC721Full *ERC721FullTransactorSession) AsyncSafeTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Full.Contract.AsyncSafeTransferFrom(handler, &_ERC721Full.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721Full *ERC721FullTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Full.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

func (_ERC721Full *ERC721FullTransactor) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Full.contract.AsyncTransact(opts, handler, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721Full *ERC721FullSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Full.Contract.SafeTransferFrom0(&_ERC721Full.TransactOpts, from, to, tokenId, _data)
}

func (_ERC721Full *ERC721FullSession) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Full.Contract.AsyncSafeTransferFrom0(handler, &_ERC721Full.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721Full *ERC721FullTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Full.Contract.SafeTransferFrom0(&_ERC721Full.TransactOpts, from, to, tokenId, _data)
}

func (_ERC721Full *ERC721FullTransactorSession) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Full.Contract.AsyncSafeTransferFrom0(handler, &_ERC721Full.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_ERC721Full *ERC721FullTransactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Full.contract.Transact(opts, "setApprovalForAll", to, approved)
}

func (_ERC721Full *ERC721FullTransactor) AsyncSetApprovalForAll(handler func(*types.Receipt, error), opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Full.contract.AsyncTransact(opts, handler, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_ERC721Full *ERC721FullSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Full.Contract.SetApprovalForAll(&_ERC721Full.TransactOpts, to, approved)
}

func (_ERC721Full *ERC721FullSession) AsyncSetApprovalForAll(handler func(*types.Receipt, error), to common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Full.Contract.AsyncSetApprovalForAll(handler, &_ERC721Full.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_ERC721Full *ERC721FullTransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Full.Contract.SetApprovalForAll(&_ERC721Full.TransactOpts, to, approved)
}

func (_ERC721Full *ERC721FullTransactorSession) AsyncSetApprovalForAll(handler func(*types.Receipt, error), to common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Full.Contract.AsyncSetApprovalForAll(handler, &_ERC721Full.TransactOpts, to, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Full *ERC721FullTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Full.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

func (_ERC721Full *ERC721FullTransactor) AsyncTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Full.contract.AsyncTransact(opts, handler, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Full *ERC721FullSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Full.Contract.TransferFrom(&_ERC721Full.TransactOpts, from, to, tokenId)
}

func (_ERC721Full *ERC721FullSession) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Full.Contract.AsyncTransferFrom(handler, &_ERC721Full.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Full *ERC721FullTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Full.Contract.TransferFrom(&_ERC721Full.TransactOpts, from, to, tokenId)
}

func (_ERC721Full *ERC721FullTransactorSession) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Full.Contract.AsyncTransferFrom(handler, &_ERC721Full.TransactOpts, from, to, tokenId)
}

// ERC721FullApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Full contract.
type ERC721FullApprovalIterator struct {
	Event *ERC721FullApproval // Event containing the contract specifics and raw log

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
func (it *ERC721FullApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721FullApproval)
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
		it.Event = new(ERC721FullApproval)
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
func (it *ERC721FullApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721FullApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721FullApproval represents a Approval event raised by the ERC721Full contract.
type ERC721FullApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Full *ERC721FullFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721FullApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Full.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721FullApprovalIterator{contract: _ERC721Full.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Full *ERC721FullFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721FullApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Full.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721FullApproval)
				if err := _ERC721Full.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Full *ERC721FullFilterer) ParseApproval(log types.Log) (*ERC721FullApproval, error) {
	event := new(ERC721FullApproval)
	if err := _ERC721Full.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC721FullApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Full contract.
type ERC721FullApprovalForAllIterator struct {
	Event *ERC721FullApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721FullApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721FullApprovalForAll)
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
		it.Event = new(ERC721FullApprovalForAll)
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
func (it *ERC721FullApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721FullApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721FullApprovalForAll represents a ApprovalForAll event raised by the ERC721Full contract.
type ERC721FullApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721Full *ERC721FullFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721FullApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721Full.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721FullApprovalForAllIterator{contract: _ERC721Full.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721Full *ERC721FullFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721FullApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721Full.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721FullApprovalForAll)
				if err := _ERC721Full.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721Full *ERC721FullFilterer) ParseApprovalForAll(log types.Log) (*ERC721FullApprovalForAll, error) {
	event := new(ERC721FullApprovalForAll)
	if err := _ERC721Full.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC721FullTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Full contract.
type ERC721FullTransferIterator struct {
	Event *ERC721FullTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721FullTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721FullTransfer)
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
		it.Event = new(ERC721FullTransfer)
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
func (it *ERC721FullTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721FullTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721FullTransfer represents a Transfer event raised by the ERC721Full contract.
type ERC721FullTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Full *ERC721FullFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721FullTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Full.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721FullTransferIterator{contract: _ERC721Full.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Full *ERC721FullFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721FullTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Full.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721FullTransfer)
				if err := _ERC721Full.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Full *ERC721FullFilterer) ParseTransfer(log types.Log) (*ERC721FullTransfer, error) {
	event := new(ERC721FullTransfer)
	if err := _ERC721Full.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC721MetadataABI is the input ABI used to generate the binding from.
const ERC721MetadataABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721MetadataFuncSigs maps the 4-byte function signature to its string representation.
var ERC721MetadataFuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"06fdde03": "name()",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"95d89b41": "symbol()",
	"c87b56dd": "tokenURI(uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC721MetadataBin is the compiled bytecode used for deploying new contracts.
var ERC721MetadataBin = "0x60806040523480156200001157600080fd5b5060405162000e4c38038062000e4c83398101604052805160208201519082019101620000677f01ffc9a70000000000000000000000000000000000000000000000000000000064010000000062000103810204565b6200009b7f80ac58cd0000000000000000000000000000000000000000000000000000000064010000000062000103810204565b8151620000b090600590602085019062000170565b508051620000c690600690602084019062000170565b50620000fb7f5b5e139f0000000000000000000000000000000000000000000000000000000064010000000062000103810204565b505062000215565b7fffffffff0000000000000000000000000000000000000000000000000000000080821614156200013357600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001b357805160ff1916838001178555620001e3565b82800160010185558215620001e3579182015b82811115620001e3578251825591602001919060010190620001c6565b50620001f1929150620001f5565b5090565b6200021291905b80821115620001f15760008155600101620001fc565b90565b610c2780620002256000396000f3006080604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301ffc9a781146100c957806306fdde0314610114578063081812fc1461019e578063095ea7b3146101d257806323b872dd146101f857806342842e0e146102225780636352211e1461024c57806370a082311461026457806395d89b4114610297578063a22cb465146102ac578063b88d4fde146102d2578063c87b56dd14610341578063e985e9c514610359575b600080fd5b3480156100d557600080fd5b506101007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1960043516610380565b604080519115158252519081900360200190f35b34801561012057600080fd5b506101296103b4565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561016357818101518382015260200161014b565b50505050905090810190601f1680156101905780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101aa57600080fd5b506101b660043561044a565b60408051600160a060020a039092168252519081900360200190f35b3480156101de57600080fd5b506101f6600160a060020a036004351660243561047c565b005b34801561020457600080fd5b506101f6600160a060020a0360043581169060243516604435610532565b34801561022e57600080fd5b506101f6600160a060020a03600435811690602435166044356105c0565b34801561025857600080fd5b506101b66004356105e1565b34801561027057600080fd5b50610285600160a060020a036004351661060b565b60408051918252519081900360200190f35b3480156102a357600080fd5b5061012961063e565b3480156102b857600080fd5b506101f6600160a060020a0360043516602435151561069f565b3480156102de57600080fd5b50604080516020601f6064356004818101359283018490048402850184019095528184526101f694600160a060020a0381358116956024803590921695604435953695608494019181908401838280828437509497506107239650505050505050565b34801561034d57600080fd5b5061012960043561074b565b34801561036557600080fd5b50610100600160a060020a0360043581169060243516610800565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191660009081526020819052604090205460ff1690565b60058054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104405780601f1061041557610100808354040283529160200191610440565b820191906000526020600020905b81548152906001019060200180831161042357829003601f168201915b5050505050905090565b60006104558261082e565b151561046057600080fd5b50600090815260026020526040902054600160a060020a031690565b6000610487826105e1565b9050600160a060020a0383811690821614156104a257600080fd5b33600160a060020a03821614806104be57506104be8133610800565b15156104c957600080fd5b600082815260026020526040808220805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b61053c338261084b565b151561054757600080fd5b600160a060020a038216151561055c57600080fd5b61056683826108aa565b610570838261091b565b61057a82826109b1565b8082600160a060020a031684600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b6105dc8383836020604051908101604052806000815250610723565b505050565b600081815260016020526040812054600160a060020a031680151561060557600080fd5b92915050565b6000600160a060020a038216151561062257600080fd5b50600160a060020a031660009081526003602052604090205490565b60068054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104405780601f1061041557610100808354040283529160200191610440565b600160a060020a0382163314156106b557600080fd5b336000818152600460209081526040808320600160a060020a03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b61072e848484610532565b61073a84848484610a41565b151561074557600080fd5b50505050565b60606107568261082e565b151561076157600080fd5b60008281526007602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156107f45780601f106107c9576101008083540402835291602001916107f4565b820191906000526020600020905b8154815290600101906020018083116107d757829003601f168201915b50505050509050919050565b600160a060020a03918216600090815260046020908152604080832093909416825291909152205460ff1690565b600090815260016020526040902054600160a060020a0316151590565b600080610857836105e1565b905080600160a060020a031684600160a060020a03161480610892575083600160a060020a03166108878461044a565b600160a060020a0316145b806108a257506108a28185610800565b949350505050565b81600160a060020a03166108bd826105e1565b600160a060020a0316146108d057600080fd5b600081815260026020526040902054600160a060020a031615610917576000818152600260205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b5050565b81600160a060020a031661092e826105e1565b600160a060020a03161461094157600080fd5b600160a060020a03821660009081526003602052604090205461096b90600163ffffffff610bc316565b600160a060020a03909216600090815260036020908152604080832094909455918152600190915220805473ffffffffffffffffffffffffffffffffffffffff19169055565b600081815260016020526040902054600160a060020a0316156109d357600080fd5b6000818152600160208181526040808420805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0388169081179091558452600390915290912054610a2191610bda565b600160a060020a0390921660009081526003602052604090209190915550565b600080610a5685600160a060020a0316610bf3565b1515610a655760019150610bba565b6040517f150b7a020000000000000000000000000000000000000000000000000000000081523360048201818152600160a060020a03898116602485015260448401889052608060648501908152875160848601528751918a169463150b7a0294938c938b938b93909160a490910190602085019080838360005b83811015610af8578181015183820152602001610ae0565b50505050905090810190601f168015610b255780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015610b4757600080fd5b505af1158015610b5b573d6000803e3d6000fd5b505050506040513d6020811015610b7157600080fd5b50517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1981167f150b7a020000000000000000000000000000000000000000000000000000000014925090505b50949350505050565b60008083831115610bd357600080fd5b5050900390565b600082820183811015610bec57600080fd5b9392505050565b6000903b11905600a165627a7a72305820e9e397418492d8706eb866639a2a3f3b9cb5e91cca1b0b1d35ebba84acd1a0c20029"

// DeployERC721Metadata deploys a new contract, binding an instance of ERC721Metadata to it.
func DeployERC721Metadata(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *ERC721Metadata, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721MetadataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721MetadataBin), backend, name, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Metadata{ERC721MetadataCaller: ERC721MetadataCaller{contract: contract}, ERC721MetadataTransactor: ERC721MetadataTransactor{contract: contract}, ERC721MetadataFilterer: ERC721MetadataFilterer{contract: contract}}, nil
}

func AsyncDeployERC721Metadata(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend, name string, symbol string) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721MetadataABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(ERC721MetadataBin), backend, name, symbol)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// ERC721Metadata is an auto generated Go binding around a Solidity contract.
type ERC721Metadata struct {
	ERC721MetadataCaller     // Read-only binding to the contract
	ERC721MetadataTransactor // Write-only binding to the contract
	ERC721MetadataFilterer   // Log filterer for contract events
}

// ERC721MetadataCaller is an auto generated read-only Go binding around a Solidity contract.
type ERC721MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MetadataTransactor is an auto generated write-only Go binding around a Solidity contract.
type ERC721MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MetadataFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ERC721MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MetadataSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ERC721MetadataSession struct {
	Contract     *ERC721Metadata   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721MetadataCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ERC721MetadataCallerSession struct {
	Contract *ERC721MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC721MetadataTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ERC721MetadataTransactorSession struct {
	Contract     *ERC721MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC721MetadataRaw is an auto generated low-level Go binding around a Solidity contract.
type ERC721MetadataRaw struct {
	Contract *ERC721Metadata // Generic contract binding to access the raw methods on
}

// ERC721MetadataCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ERC721MetadataCallerRaw struct {
	Contract *ERC721MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721MetadataTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ERC721MetadataTransactorRaw struct {
	Contract *ERC721MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Metadata creates a new instance of ERC721Metadata, bound to a specific deployed contract.
func NewERC721Metadata(address common.Address, backend bind.ContractBackend) (*ERC721Metadata, error) {
	contract, err := bindERC721Metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Metadata{ERC721MetadataCaller: ERC721MetadataCaller{contract: contract}, ERC721MetadataTransactor: ERC721MetadataTransactor{contract: contract}, ERC721MetadataFilterer: ERC721MetadataFilterer{contract: contract}}, nil
}

// NewERC721MetadataCaller creates a new read-only instance of ERC721Metadata, bound to a specific deployed contract.
func NewERC721MetadataCaller(address common.Address, caller bind.ContractCaller) (*ERC721MetadataCaller, error) {
	contract, err := bindERC721Metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataCaller{contract: contract}, nil
}

// NewERC721MetadataTransactor creates a new write-only instance of ERC721Metadata, bound to a specific deployed contract.
func NewERC721MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721MetadataTransactor, error) {
	contract, err := bindERC721Metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataTransactor{contract: contract}, nil
}

// NewERC721MetadataFilterer creates a new log filterer instance of ERC721Metadata, bound to a specific deployed contract.
func NewERC721MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721MetadataFilterer, error) {
	contract, err := bindERC721Metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataFilterer{contract: contract}, nil
}

// bindERC721Metadata binds a generic wrapper to an already deployed contract.
func bindERC721Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721MetadataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Metadata *ERC721MetadataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Metadata.Contract.ERC721MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Metadata *ERC721MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Metadata.Contract.ERC721MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Metadata *ERC721MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Metadata.Contract.ERC721MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Metadata *ERC721MetadataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Metadata *ERC721MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Metadata *ERC721MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Metadata.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC721Metadata *ERC721MetadataCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC721Metadata *ERC721MetadataSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721Metadata.Contract.BalanceOf(&_ERC721Metadata.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC721Metadata *ERC721MetadataCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721Metadata.Contract.BalanceOf(&_ERC721Metadata.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_ERC721Metadata *ERC721MetadataCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_ERC721Metadata *ERC721MetadataSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721Metadata.Contract.GetApproved(&_ERC721Metadata.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_ERC721Metadata *ERC721MetadataCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721Metadata.Contract.GetApproved(&_ERC721Metadata.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_ERC721Metadata *ERC721MetadataCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_ERC721Metadata *ERC721MetadataSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721Metadata.Contract.IsApprovedForAll(&_ERC721Metadata.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_ERC721Metadata *ERC721MetadataCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721Metadata.Contract.IsApprovedForAll(&_ERC721Metadata.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC721Metadata *ERC721MetadataCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC721Metadata *ERC721MetadataSession) Name() (string, error) {
	return _ERC721Metadata.Contract.Name(&_ERC721Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC721Metadata *ERC721MetadataCallerSession) Name() (string, error) {
	return _ERC721Metadata.Contract.Name(&_ERC721Metadata.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_ERC721Metadata *ERC721MetadataCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_ERC721Metadata *ERC721MetadataSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721Metadata.Contract.OwnerOf(&_ERC721Metadata.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_ERC721Metadata *ERC721MetadataCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721Metadata.Contract.OwnerOf(&_ERC721Metadata.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC721Metadata *ERC721MetadataCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC721Metadata *ERC721MetadataSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721Metadata.Contract.SupportsInterface(&_ERC721Metadata.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC721Metadata *ERC721MetadataCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721Metadata.Contract.SupportsInterface(&_ERC721Metadata.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC721Metadata *ERC721MetadataCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC721Metadata *ERC721MetadataSession) Symbol() (string, error) {
	return _ERC721Metadata.Contract.Symbol(&_ERC721Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC721Metadata *ERC721MetadataCallerSession) Symbol() (string, error) {
	return _ERC721Metadata.Contract.Symbol(&_ERC721Metadata.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) constant returns(string)
func (_ERC721Metadata *ERC721MetadataCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "tokenURI", tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) constant returns(string)
func (_ERC721Metadata *ERC721MetadataSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721Metadata.Contract.TokenURI(&_ERC721Metadata.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) constant returns(string)
func (_ERC721Metadata *ERC721MetadataCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721Metadata.Contract.TokenURI(&_ERC721Metadata.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Metadata *ERC721MetadataTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Metadata.contract.Transact(opts, "approve", to, tokenId)
}

func (_ERC721Metadata *ERC721MetadataTransactor) AsyncApprove(handler func(*types.Receipt, error), opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.contract.AsyncTransact(opts, handler, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Metadata *ERC721MetadataSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Metadata.Contract.Approve(&_ERC721Metadata.TransactOpts, to, tokenId)
}

func (_ERC721Metadata *ERC721MetadataSession) AsyncApprove(handler func(*types.Receipt, error), to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.AsyncApprove(handler, &_ERC721Metadata.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721Metadata *ERC721MetadataTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Metadata.Contract.Approve(&_ERC721Metadata.TransactOpts, to, tokenId)
}

func (_ERC721Metadata *ERC721MetadataTransactorSession) AsyncApprove(handler func(*types.Receipt, error), to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.AsyncApprove(handler, &_ERC721Metadata.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Metadata *ERC721MetadataTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Metadata.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

func (_ERC721Metadata *ERC721MetadataTransactor) AsyncSafeTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.contract.AsyncTransact(opts, handler, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Metadata *ERC721MetadataSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Metadata.Contract.SafeTransferFrom(&_ERC721Metadata.TransactOpts, from, to, tokenId)
}

func (_ERC721Metadata *ERC721MetadataSession) AsyncSafeTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.AsyncSafeTransferFrom(handler, &_ERC721Metadata.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Metadata *ERC721MetadataTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Metadata.Contract.SafeTransferFrom(&_ERC721Metadata.TransactOpts, from, to, tokenId)
}

func (_ERC721Metadata *ERC721MetadataTransactorSession) AsyncSafeTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.AsyncSafeTransferFrom(handler, &_ERC721Metadata.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721Metadata *ERC721MetadataTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Metadata.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

func (_ERC721Metadata *ERC721MetadataTransactor) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Metadata.contract.AsyncTransact(opts, handler, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721Metadata *ERC721MetadataSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Metadata.Contract.SafeTransferFrom0(&_ERC721Metadata.TransactOpts, from, to, tokenId, _data)
}

func (_ERC721Metadata *ERC721MetadataSession) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.AsyncSafeTransferFrom0(handler, &_ERC721Metadata.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721Metadata *ERC721MetadataTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Metadata.Contract.SafeTransferFrom0(&_ERC721Metadata.TransactOpts, from, to, tokenId, _data)
}

func (_ERC721Metadata *ERC721MetadataTransactorSession) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.AsyncSafeTransferFrom0(handler, &_ERC721Metadata.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_ERC721Metadata *ERC721MetadataTransactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Metadata.contract.Transact(opts, "setApprovalForAll", to, approved)
}

func (_ERC721Metadata *ERC721MetadataTransactor) AsyncSetApprovalForAll(handler func(*types.Receipt, error), opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Metadata.contract.AsyncTransact(opts, handler, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_ERC721Metadata *ERC721MetadataSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Metadata.Contract.SetApprovalForAll(&_ERC721Metadata.TransactOpts, to, approved)
}

func (_ERC721Metadata *ERC721MetadataSession) AsyncSetApprovalForAll(handler func(*types.Receipt, error), to common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.AsyncSetApprovalForAll(handler, &_ERC721Metadata.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_ERC721Metadata *ERC721MetadataTransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Metadata.Contract.SetApprovalForAll(&_ERC721Metadata.TransactOpts, to, approved)
}

func (_ERC721Metadata *ERC721MetadataTransactorSession) AsyncSetApprovalForAll(handler func(*types.Receipt, error), to common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.AsyncSetApprovalForAll(handler, &_ERC721Metadata.TransactOpts, to, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Metadata *ERC721MetadataTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Metadata.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

func (_ERC721Metadata *ERC721MetadataTransactor) AsyncTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.contract.AsyncTransact(opts, handler, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Metadata *ERC721MetadataSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Metadata.Contract.TransferFrom(&_ERC721Metadata.TransactOpts, from, to, tokenId)
}

func (_ERC721Metadata *ERC721MetadataSession) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.AsyncTransferFrom(handler, &_ERC721Metadata.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721Metadata *ERC721MetadataTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721Metadata.Contract.TransferFrom(&_ERC721Metadata.TransactOpts, from, to, tokenId)
}

func (_ERC721Metadata *ERC721MetadataTransactorSession) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.AsyncTransferFrom(handler, &_ERC721Metadata.TransactOpts, from, to, tokenId)
}

// ERC721MetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Metadata contract.
type ERC721MetadataApprovalIterator struct {
	Event *ERC721MetadataApproval // Event containing the contract specifics and raw log

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
func (it *ERC721MetadataApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MetadataApproval)
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
		it.Event = new(ERC721MetadataApproval)
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
func (it *ERC721MetadataApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MetadataApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MetadataApproval represents a Approval event raised by the ERC721Metadata contract.
type ERC721MetadataApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Metadata *ERC721MetadataFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721MetadataApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Metadata.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataApprovalIterator{contract: _ERC721Metadata.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Metadata *ERC721MetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721MetadataApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Metadata.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MetadataApproval)
				if err := _ERC721Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721Metadata *ERC721MetadataFilterer) ParseApproval(log types.Log) (*ERC721MetadataApproval, error) {
	event := new(ERC721MetadataApproval)
	if err := _ERC721Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC721MetadataApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Metadata contract.
type ERC721MetadataApprovalForAllIterator struct {
	Event *ERC721MetadataApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721MetadataApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MetadataApprovalForAll)
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
		it.Event = new(ERC721MetadataApprovalForAll)
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
func (it *ERC721MetadataApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MetadataApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MetadataApprovalForAll represents a ApprovalForAll event raised by the ERC721Metadata contract.
type ERC721MetadataApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721Metadata *ERC721MetadataFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721MetadataApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721Metadata.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataApprovalForAllIterator{contract: _ERC721Metadata.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721Metadata *ERC721MetadataFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721MetadataApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721Metadata.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MetadataApprovalForAll)
				if err := _ERC721Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721Metadata *ERC721MetadataFilterer) ParseApprovalForAll(log types.Log) (*ERC721MetadataApprovalForAll, error) {
	event := new(ERC721MetadataApprovalForAll)
	if err := _ERC721Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC721MetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Metadata contract.
type ERC721MetadataTransferIterator struct {
	Event *ERC721MetadataTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721MetadataTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MetadataTransfer)
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
		it.Event = new(ERC721MetadataTransfer)
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
func (it *ERC721MetadataTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MetadataTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MetadataTransfer represents a Transfer event raised by the ERC721Metadata contract.
type ERC721MetadataTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Metadata *ERC721MetadataFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721MetadataTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Metadata.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataTransferIterator{contract: _ERC721Metadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Metadata *ERC721MetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721MetadataTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721Metadata.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MetadataTransfer)
				if err := _ERC721Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721Metadata *ERC721MetadataFilterer) ParseTransfer(log types.Log) (*ERC721MetadataTransfer, error) {
	event := new(ERC721MetadataTransfer)
	if err := _ERC721Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC721MetadataMintableABI is the input ABI used to generate the binding from.
const ERC721MetadataMintableABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"mintWithTokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721MetadataMintableFuncSigs maps the 4-byte function signature to its string representation.
var ERC721MetadataMintableFuncSigs = map[string]string{
	"983b2d56": "addMinter(address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"aa271e1a": "isMinter(address)",
	"50bb4e7f": "mintWithTokenURI(address,uint256,string)",
	"06fdde03": "name()",
	"6352211e": "ownerOf(uint256)",
	"98650275": "renounceMinter()",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"95d89b41": "symbol()",
	"c87b56dd": "tokenURI(uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC721MetadataMintable is an auto generated Go binding around a Solidity contract.
type ERC721MetadataMintable struct {
	ERC721MetadataMintableCaller     // Read-only binding to the contract
	ERC721MetadataMintableTransactor // Write-only binding to the contract
	ERC721MetadataMintableFilterer   // Log filterer for contract events
}

// ERC721MetadataMintableCaller is an auto generated read-only Go binding around a Solidity contract.
type ERC721MetadataMintableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MetadataMintableTransactor is an auto generated write-only Go binding around a Solidity contract.
type ERC721MetadataMintableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MetadataMintableFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ERC721MetadataMintableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MetadataMintableSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ERC721MetadataMintableSession struct {
	Contract     *ERC721MetadataMintable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC721MetadataMintableCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ERC721MetadataMintableCallerSession struct {
	Contract *ERC721MetadataMintableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ERC721MetadataMintableTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ERC721MetadataMintableTransactorSession struct {
	Contract     *ERC721MetadataMintableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ERC721MetadataMintableRaw is an auto generated low-level Go binding around a Solidity contract.
type ERC721MetadataMintableRaw struct {
	Contract *ERC721MetadataMintable // Generic contract binding to access the raw methods on
}

// ERC721MetadataMintableCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ERC721MetadataMintableCallerRaw struct {
	Contract *ERC721MetadataMintableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721MetadataMintableTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ERC721MetadataMintableTransactorRaw struct {
	Contract *ERC721MetadataMintableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721MetadataMintable creates a new instance of ERC721MetadataMintable, bound to a specific deployed contract.
func NewERC721MetadataMintable(address common.Address, backend bind.ContractBackend) (*ERC721MetadataMintable, error) {
	contract, err := bindERC721MetadataMintable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataMintable{ERC721MetadataMintableCaller: ERC721MetadataMintableCaller{contract: contract}, ERC721MetadataMintableTransactor: ERC721MetadataMintableTransactor{contract: contract}, ERC721MetadataMintableFilterer: ERC721MetadataMintableFilterer{contract: contract}}, nil
}

// NewERC721MetadataMintableCaller creates a new read-only instance of ERC721MetadataMintable, bound to a specific deployed contract.
func NewERC721MetadataMintableCaller(address common.Address, caller bind.ContractCaller) (*ERC721MetadataMintableCaller, error) {
	contract, err := bindERC721MetadataMintable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataMintableCaller{contract: contract}, nil
}

// NewERC721MetadataMintableTransactor creates a new write-only instance of ERC721MetadataMintable, bound to a specific deployed contract.
func NewERC721MetadataMintableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721MetadataMintableTransactor, error) {
	contract, err := bindERC721MetadataMintable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataMintableTransactor{contract: contract}, nil
}

// NewERC721MetadataMintableFilterer creates a new log filterer instance of ERC721MetadataMintable, bound to a specific deployed contract.
func NewERC721MetadataMintableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721MetadataMintableFilterer, error) {
	contract, err := bindERC721MetadataMintable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataMintableFilterer{contract: contract}, nil
}

// bindERC721MetadataMintable binds a generic wrapper to an already deployed contract.
func bindERC721MetadataMintable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721MetadataMintableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721MetadataMintable *ERC721MetadataMintableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721MetadataMintable.Contract.ERC721MetadataMintableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721MetadataMintable *ERC721MetadataMintableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.Contract.ERC721MetadataMintableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721MetadataMintable *ERC721MetadataMintableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.Contract.ERC721MetadataMintableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721MetadataMintable *ERC721MetadataMintableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721MetadataMintable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721MetadataMintable *ERC721MetadataMintableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721MetadataMintable *ERC721MetadataMintableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC721MetadataMintable *ERC721MetadataMintableCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721MetadataMintable.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC721MetadataMintable *ERC721MetadataMintableSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721MetadataMintable.Contract.BalanceOf(&_ERC721MetadataMintable.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC721MetadataMintable *ERC721MetadataMintableCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC721MetadataMintable.Contract.BalanceOf(&_ERC721MetadataMintable.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_ERC721MetadataMintable *ERC721MetadataMintableCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721MetadataMintable.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_ERC721MetadataMintable *ERC721MetadataMintableSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721MetadataMintable.Contract.GetApproved(&_ERC721MetadataMintable.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_ERC721MetadataMintable *ERC721MetadataMintableCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ERC721MetadataMintable.Contract.GetApproved(&_ERC721MetadataMintable.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_ERC721MetadataMintable *ERC721MetadataMintableCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721MetadataMintable.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_ERC721MetadataMintable *ERC721MetadataMintableSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721MetadataMintable.Contract.IsApprovedForAll(&_ERC721MetadataMintable.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_ERC721MetadataMintable *ERC721MetadataMintableCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ERC721MetadataMintable.Contract.IsApprovedForAll(&_ERC721MetadataMintable.CallOpts, owner, operator)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_ERC721MetadataMintable *ERC721MetadataMintableCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721MetadataMintable.contract.Call(opts, out, "isMinter", account)
	return *ret0, err
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_ERC721MetadataMintable *ERC721MetadataMintableSession) IsMinter(account common.Address) (bool, error) {
	return _ERC721MetadataMintable.Contract.IsMinter(&_ERC721MetadataMintable.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_ERC721MetadataMintable *ERC721MetadataMintableCallerSession) IsMinter(account common.Address) (bool, error) {
	return _ERC721MetadataMintable.Contract.IsMinter(&_ERC721MetadataMintable.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC721MetadataMintable *ERC721MetadataMintableCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721MetadataMintable.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC721MetadataMintable *ERC721MetadataMintableSession) Name() (string, error) {
	return _ERC721MetadataMintable.Contract.Name(&_ERC721MetadataMintable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC721MetadataMintable *ERC721MetadataMintableCallerSession) Name() (string, error) {
	return _ERC721MetadataMintable.Contract.Name(&_ERC721MetadataMintable.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_ERC721MetadataMintable *ERC721MetadataMintableCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721MetadataMintable.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_ERC721MetadataMintable *ERC721MetadataMintableSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721MetadataMintable.Contract.OwnerOf(&_ERC721MetadataMintable.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_ERC721MetadataMintable *ERC721MetadataMintableCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ERC721MetadataMintable.Contract.OwnerOf(&_ERC721MetadataMintable.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC721MetadataMintable *ERC721MetadataMintableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721MetadataMintable.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC721MetadataMintable *ERC721MetadataMintableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721MetadataMintable.Contract.SupportsInterface(&_ERC721MetadataMintable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_ERC721MetadataMintable *ERC721MetadataMintableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC721MetadataMintable.Contract.SupportsInterface(&_ERC721MetadataMintable.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC721MetadataMintable *ERC721MetadataMintableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721MetadataMintable.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC721MetadataMintable *ERC721MetadataMintableSession) Symbol() (string, error) {
	return _ERC721MetadataMintable.Contract.Symbol(&_ERC721MetadataMintable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC721MetadataMintable *ERC721MetadataMintableCallerSession) Symbol() (string, error) {
	return _ERC721MetadataMintable.Contract.Symbol(&_ERC721MetadataMintable.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) constant returns(string)
func (_ERC721MetadataMintable *ERC721MetadataMintableCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721MetadataMintable.contract.Call(opts, out, "tokenURI", tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) constant returns(string)
func (_ERC721MetadataMintable *ERC721MetadataMintableSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721MetadataMintable.Contract.TokenURI(&_ERC721MetadataMintable.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) constant returns(string)
func (_ERC721MetadataMintable *ERC721MetadataMintableCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ERC721MetadataMintable.Contract.TokenURI(&_ERC721MetadataMintable.CallOpts, tokenId)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_ERC721MetadataMintable *ERC721MetadataMintableTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.contract.Transact(opts, "addMinter", account)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableTransactor) AsyncAddMinter(handler func(*types.Receipt, error), opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ERC721MetadataMintable.contract.AsyncTransact(opts, handler, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_ERC721MetadataMintable *ERC721MetadataMintableSession) AddMinter(account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.Contract.AddMinter(&_ERC721MetadataMintable.TransactOpts, account)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableSession) AsyncAddMinter(handler func(*types.Receipt, error), account common.Address) (*types.Transaction, error) {
	return _ERC721MetadataMintable.Contract.AsyncAddMinter(handler, &_ERC721MetadataMintable.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_ERC721MetadataMintable *ERC721MetadataMintableTransactorSession) AddMinter(account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.Contract.AddMinter(&_ERC721MetadataMintable.TransactOpts, account)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableTransactorSession) AsyncAddMinter(handler func(*types.Receipt, error), account common.Address) (*types.Transaction, error) {
	return _ERC721MetadataMintable.Contract.AsyncAddMinter(handler, &_ERC721MetadataMintable.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721MetadataMintable *ERC721MetadataMintableTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.contract.Transact(opts, "approve", to, tokenId)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableTransactor) AsyncApprove(handler func(*types.Receipt, error), opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721MetadataMintable.contract.AsyncTransact(opts, handler, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721MetadataMintable *ERC721MetadataMintableSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.Contract.Approve(&_ERC721MetadataMintable.TransactOpts, to, tokenId)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableSession) AsyncApprove(handler func(*types.Receipt, error), to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721MetadataMintable.Contract.AsyncApprove(handler, &_ERC721MetadataMintable.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ERC721MetadataMintable *ERC721MetadataMintableTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.Contract.Approve(&_ERC721MetadataMintable.TransactOpts, to, tokenId)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableTransactorSession) AsyncApprove(handler func(*types.Receipt, error), to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721MetadataMintable.Contract.AsyncApprove(handler, &_ERC721MetadataMintable.TransactOpts, to, tokenId)
}

// MintWithTokenURI is a paid mutator transaction binding the contract method 0x50bb4e7f.
//
// Solidity: function mintWithTokenURI(address to, uint256 tokenId, string tokenURI) returns(bool)
func (_ERC721MetadataMintable *ERC721MetadataMintableTransactor) MintWithTokenURI(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.contract.Transact(opts, "mintWithTokenURI", to, tokenId, tokenURI)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableTransactor) AsyncMintWithTokenURI(handler func(*types.Receipt, error), opts *bind.TransactOpts, to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _ERC721MetadataMintable.contract.AsyncTransact(opts, handler, "mintWithTokenURI", to, tokenId, tokenURI)
}

// MintWithTokenURI is a paid mutator transaction binding the contract method 0x50bb4e7f.
//
// Solidity: function mintWithTokenURI(address to, uint256 tokenId, string tokenURI) returns(bool)
func (_ERC721MetadataMintable *ERC721MetadataMintableSession) MintWithTokenURI(to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.Contract.MintWithTokenURI(&_ERC721MetadataMintable.TransactOpts, to, tokenId, tokenURI)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableSession) AsyncMintWithTokenURI(handler func(*types.Receipt, error), to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _ERC721MetadataMintable.Contract.AsyncMintWithTokenURI(handler, &_ERC721MetadataMintable.TransactOpts, to, tokenId, tokenURI)
}

// MintWithTokenURI is a paid mutator transaction binding the contract method 0x50bb4e7f.
//
// Solidity: function mintWithTokenURI(address to, uint256 tokenId, string tokenURI) returns(bool)
func (_ERC721MetadataMintable *ERC721MetadataMintableTransactorSession) MintWithTokenURI(to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.Contract.MintWithTokenURI(&_ERC721MetadataMintable.TransactOpts, to, tokenId, tokenURI)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableTransactorSession) AsyncMintWithTokenURI(handler func(*types.Receipt, error), to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _ERC721MetadataMintable.Contract.AsyncMintWithTokenURI(handler, &_ERC721MetadataMintable.TransactOpts, to, tokenId, tokenURI)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_ERC721MetadataMintable *ERC721MetadataMintableTransactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.contract.Transact(opts, "renounceMinter")
}

func (_ERC721MetadataMintable *ERC721MetadataMintableTransactor) AsyncRenounceMinter(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721MetadataMintable.contract.AsyncTransact(opts, handler, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_ERC721MetadataMintable *ERC721MetadataMintableSession) RenounceMinter() (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.Contract.RenounceMinter(&_ERC721MetadataMintable.TransactOpts)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableSession) AsyncRenounceMinter(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _ERC721MetadataMintable.Contract.AsyncRenounceMinter(handler, &_ERC721MetadataMintable.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_ERC721MetadataMintable *ERC721MetadataMintableTransactorSession) RenounceMinter() (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.Contract.RenounceMinter(&_ERC721MetadataMintable.TransactOpts)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableTransactorSession) AsyncRenounceMinter(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _ERC721MetadataMintable.Contract.AsyncRenounceMinter(handler, &_ERC721MetadataMintable.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721MetadataMintable *ERC721MetadataMintableTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableTransactor) AsyncSafeTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721MetadataMintable.contract.AsyncTransact(opts, handler, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721MetadataMintable *ERC721MetadataMintableSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.Contract.SafeTransferFrom(&_ERC721MetadataMintable.TransactOpts, from, to, tokenId)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableSession) AsyncSafeTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721MetadataMintable.Contract.AsyncSafeTransferFrom(handler, &_ERC721MetadataMintable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721MetadataMintable *ERC721MetadataMintableTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.Contract.SafeTransferFrom(&_ERC721MetadataMintable.TransactOpts, from, to, tokenId)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableTransactorSession) AsyncSafeTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721MetadataMintable.Contract.AsyncSafeTransferFrom(handler, &_ERC721MetadataMintable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721MetadataMintable *ERC721MetadataMintableTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableTransactor) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721MetadataMintable.contract.AsyncTransact(opts, handler, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721MetadataMintable *ERC721MetadataMintableSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.Contract.SafeTransferFrom0(&_ERC721MetadataMintable.TransactOpts, from, to, tokenId, _data)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableSession) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721MetadataMintable.Contract.AsyncSafeTransferFrom0(handler, &_ERC721MetadataMintable.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ERC721MetadataMintable *ERC721MetadataMintableTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.Contract.SafeTransferFrom0(&_ERC721MetadataMintable.TransactOpts, from, to, tokenId, _data)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableTransactorSession) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721MetadataMintable.Contract.AsyncSafeTransferFrom0(handler, &_ERC721MetadataMintable.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_ERC721MetadataMintable *ERC721MetadataMintableTransactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.contract.Transact(opts, "setApprovalForAll", to, approved)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableTransactor) AsyncSetApprovalForAll(handler func(*types.Receipt, error), opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721MetadataMintable.contract.AsyncTransact(opts, handler, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_ERC721MetadataMintable *ERC721MetadataMintableSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.Contract.SetApprovalForAll(&_ERC721MetadataMintable.TransactOpts, to, approved)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableSession) AsyncSetApprovalForAll(handler func(*types.Receipt, error), to common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721MetadataMintable.Contract.AsyncSetApprovalForAll(handler, &_ERC721MetadataMintable.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_ERC721MetadataMintable *ERC721MetadataMintableTransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.Contract.SetApprovalForAll(&_ERC721MetadataMintable.TransactOpts, to, approved)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableTransactorSession) AsyncSetApprovalForAll(handler func(*types.Receipt, error), to common.Address, approved bool) (*types.Transaction, error) {
	return _ERC721MetadataMintable.Contract.AsyncSetApprovalForAll(handler, &_ERC721MetadataMintable.TransactOpts, to, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721MetadataMintable *ERC721MetadataMintableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableTransactor) AsyncTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721MetadataMintable.contract.AsyncTransact(opts, handler, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721MetadataMintable *ERC721MetadataMintableSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.Contract.TransferFrom(&_ERC721MetadataMintable.TransactOpts, from, to, tokenId)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableSession) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721MetadataMintable.Contract.AsyncTransferFrom(handler, &_ERC721MetadataMintable.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ERC721MetadataMintable *ERC721MetadataMintableTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _ERC721MetadataMintable.Contract.TransferFrom(&_ERC721MetadataMintable.TransactOpts, from, to, tokenId)
}

func (_ERC721MetadataMintable *ERC721MetadataMintableTransactorSession) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721MetadataMintable.Contract.AsyncTransferFrom(handler, &_ERC721MetadataMintable.TransactOpts, from, to, tokenId)
}

// ERC721MetadataMintableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721MetadataMintable contract.
type ERC721MetadataMintableApprovalIterator struct {
	Event *ERC721MetadataMintableApproval // Event containing the contract specifics and raw log

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
func (it *ERC721MetadataMintableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MetadataMintableApproval)
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
		it.Event = new(ERC721MetadataMintableApproval)
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
func (it *ERC721MetadataMintableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MetadataMintableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MetadataMintableApproval represents a Approval event raised by the ERC721MetadataMintable contract.
type ERC721MetadataMintableApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721MetadataMintable *ERC721MetadataMintableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721MetadataMintableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721MetadataMintable.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataMintableApprovalIterator{contract: _ERC721MetadataMintable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721MetadataMintable *ERC721MetadataMintableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721MetadataMintableApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721MetadataMintable.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MetadataMintableApproval)
				if err := _ERC721MetadataMintable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ERC721MetadataMintable *ERC721MetadataMintableFilterer) ParseApproval(log types.Log) (*ERC721MetadataMintableApproval, error) {
	event := new(ERC721MetadataMintableApproval)
	if err := _ERC721MetadataMintable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC721MetadataMintableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721MetadataMintable contract.
type ERC721MetadataMintableApprovalForAllIterator struct {
	Event *ERC721MetadataMintableApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721MetadataMintableApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MetadataMintableApprovalForAll)
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
		it.Event = new(ERC721MetadataMintableApprovalForAll)
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
func (it *ERC721MetadataMintableApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MetadataMintableApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MetadataMintableApprovalForAll represents a ApprovalForAll event raised by the ERC721MetadataMintable contract.
type ERC721MetadataMintableApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721MetadataMintable *ERC721MetadataMintableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721MetadataMintableApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721MetadataMintable.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataMintableApprovalForAllIterator{contract: _ERC721MetadataMintable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721MetadataMintable *ERC721MetadataMintableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721MetadataMintableApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ERC721MetadataMintable.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MetadataMintableApprovalForAll)
				if err := _ERC721MetadataMintable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ERC721MetadataMintable *ERC721MetadataMintableFilterer) ParseApprovalForAll(log types.Log) (*ERC721MetadataMintableApprovalForAll, error) {
	event := new(ERC721MetadataMintableApprovalForAll)
	if err := _ERC721MetadataMintable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC721MetadataMintableMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the ERC721MetadataMintable contract.
type ERC721MetadataMintableMinterAddedIterator struct {
	Event *ERC721MetadataMintableMinterAdded // Event containing the contract specifics and raw log

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
func (it *ERC721MetadataMintableMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MetadataMintableMinterAdded)
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
		it.Event = new(ERC721MetadataMintableMinterAdded)
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
func (it *ERC721MetadataMintableMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MetadataMintableMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MetadataMintableMinterAdded represents a MinterAdded event raised by the ERC721MetadataMintable contract.
type ERC721MetadataMintableMinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_ERC721MetadataMintable *ERC721MetadataMintableFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*ERC721MetadataMintableMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERC721MetadataMintable.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataMintableMinterAddedIterator{contract: _ERC721MetadataMintable.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_ERC721MetadataMintable *ERC721MetadataMintableFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *ERC721MetadataMintableMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERC721MetadataMintable.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MetadataMintableMinterAdded)
				if err := _ERC721MetadataMintable.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// ParseMinterAdded is a log parse operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_ERC721MetadataMintable *ERC721MetadataMintableFilterer) ParseMinterAdded(log types.Log) (*ERC721MetadataMintableMinterAdded, error) {
	event := new(ERC721MetadataMintableMinterAdded)
	if err := _ERC721MetadataMintable.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC721MetadataMintableMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the ERC721MetadataMintable contract.
type ERC721MetadataMintableMinterRemovedIterator struct {
	Event *ERC721MetadataMintableMinterRemoved // Event containing the contract specifics and raw log

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
func (it *ERC721MetadataMintableMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MetadataMintableMinterRemoved)
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
		it.Event = new(ERC721MetadataMintableMinterRemoved)
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
func (it *ERC721MetadataMintableMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MetadataMintableMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MetadataMintableMinterRemoved represents a MinterRemoved event raised by the ERC721MetadataMintable contract.
type ERC721MetadataMintableMinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_ERC721MetadataMintable *ERC721MetadataMintableFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*ERC721MetadataMintableMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERC721MetadataMintable.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataMintableMinterRemovedIterator{contract: _ERC721MetadataMintable.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_ERC721MetadataMintable *ERC721MetadataMintableFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *ERC721MetadataMintableMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ERC721MetadataMintable.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MetadataMintableMinterRemoved)
				if err := _ERC721MetadataMintable.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_ERC721MetadataMintable *ERC721MetadataMintableFilterer) ParseMinterRemoved(log types.Log) (*ERC721MetadataMintableMinterRemoved, error) {
	event := new(ERC721MetadataMintableMinterRemoved)
	if err := _ERC721MetadataMintable.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC721MetadataMintableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721MetadataMintable contract.
type ERC721MetadataMintableTransferIterator struct {
	Event *ERC721MetadataMintableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721MetadataMintableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MetadataMintableTransfer)
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
		it.Event = new(ERC721MetadataMintableTransfer)
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
func (it *ERC721MetadataMintableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MetadataMintableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MetadataMintableTransfer represents a Transfer event raised by the ERC721MetadataMintable contract.
type ERC721MetadataMintableTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721MetadataMintable *ERC721MetadataMintableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721MetadataMintableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721MetadataMintable.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataMintableTransferIterator{contract: _ERC721MetadataMintable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721MetadataMintable *ERC721MetadataMintableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721MetadataMintableTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721MetadataMintable.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MetadataMintableTransfer)
				if err := _ERC721MetadataMintable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ERC721MetadataMintable *ERC721MetadataMintableFilterer) ParseTransfer(log types.Log) (*ERC721MetadataMintableTransfer, error) {
	event := new(ERC721MetadataMintableTransfer)
	if err := _ERC721MetadataMintable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC165ABI is the input ABI used to generate the binding from.
const IERC165ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IERC165FuncSigs maps the 4-byte function signature to its string representation.
var IERC165FuncSigs = map[string]string{
	"01ffc9a7": "supportsInterface(bytes4)",
}

// IERC165 is an auto generated Go binding around a Solidity contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around a Solidity contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around a Solidity contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around a Solidity contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around a Solidity contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IERC165.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// IERC721ABI is the input ABI used to generate the binding from.
const IERC721ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// IERC721FuncSigs maps the 4-byte function signature to its string representation.
var IERC721FuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC721 is an auto generated Go binding around a Solidity contract.
type IERC721 struct {
	IERC721Caller     // Read-only binding to the contract
	IERC721Transactor // Write-only binding to the contract
	IERC721Filterer   // Log filterer for contract events
}

// IERC721Caller is an auto generated read-only Go binding around a Solidity contract.
type IERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Transactor is an auto generated write-only Go binding around a Solidity contract.
type IERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Filterer is an auto generated log filtering Go binding around a Solidity contract events.
type IERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Session is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type IERC721Session struct {
	Contract     *IERC721          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721CallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type IERC721CallerSession struct {
	Contract *IERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC721TransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type IERC721TransactorSession struct {
	Contract     *IERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC721Raw is an auto generated low-level Go binding around a Solidity contract.
type IERC721Raw struct {
	Contract *IERC721 // Generic contract binding to access the raw methods on
}

// IERC721CallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type IERC721CallerRaw struct {
	Contract *IERC721Caller // Generic read-only contract binding to access the raw methods on
}

// IERC721TransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type IERC721TransactorRaw struct {
	Contract *IERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721 creates a new instance of IERC721, bound to a specific deployed contract.
func NewIERC721(address common.Address, backend bind.ContractBackend) (*IERC721, error) {
	contract, err := bindIERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721{IERC721Caller: IERC721Caller{contract: contract}, IERC721Transactor: IERC721Transactor{contract: contract}, IERC721Filterer: IERC721Filterer{contract: contract}}, nil
}

// NewIERC721Caller creates a new read-only instance of IERC721, bound to a specific deployed contract.
func NewIERC721Caller(address common.Address, caller bind.ContractCaller) (*IERC721Caller, error) {
	contract, err := bindIERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721Caller{contract: contract}, nil
}

// NewIERC721Transactor creates a new write-only instance of IERC721, bound to a specific deployed contract.
func NewIERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC721Transactor, error) {
	contract, err := bindIERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721Transactor{contract: contract}, nil
}

// NewIERC721Filterer creates a new log filterer instance of IERC721, bound to a specific deployed contract.
func NewIERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC721Filterer, error) {
	contract, err := bindIERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721Filterer{contract: contract}, nil
}

// bindIERC721 binds a generic wrapper to an already deployed contract.
func bindIERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721 *IERC721Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC721.Contract.IERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721 *IERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _IERC721.Contract.IERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721 *IERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _IERC721.Contract.IERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721 *IERC721CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721 *IERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _IERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721 *IERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _IERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_IERC721 *IERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC721.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_IERC721 *IERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_IERC721 *IERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address operator)
func (_IERC721 *IERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IERC721.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address operator)
func (_IERC721 *IERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.GetApproved(&_IERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address operator)
func (_IERC721 *IERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.GetApproved(&_IERC721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_IERC721 *IERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IERC721.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_IERC721 *IERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_IERC721 *IERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address owner)
func (_IERC721 *IERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IERC721.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address owner)
func (_IERC721 *IERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address owner)
func (_IERC721 *IERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC721 *IERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IERC721.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC721 *IERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721.Contract.SupportsInterface(&_IERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC721 *IERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721.Contract.SupportsInterface(&_IERC721.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721.contract.Transact(opts, "approve", to, tokenId)
}

func (_IERC721 *IERC721Transactor) AsyncApprove(handler func(*types.Receipt, error), opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.AsyncTransact(opts, handler, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721.Contract.Approve(&_IERC721.TransactOpts, to, tokenId)
}

func (_IERC721 *IERC721Session) AsyncApprove(handler func(*types.Receipt, error), to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.AsyncApprove(handler, &_IERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721.Contract.Approve(&_IERC721.TransactOpts, to, tokenId)
}

func (_IERC721 *IERC721TransactorSession) AsyncApprove(handler func(*types.Receipt, error), to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.AsyncApprove(handler, &_IERC721.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

func (_IERC721 *IERC721Transactor) AsyncSafeTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.AsyncTransact(opts, handler, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

func (_IERC721 *IERC721Session) AsyncSafeTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.AsyncSafeTransferFrom(handler, &_IERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

func (_IERC721 *IERC721TransactorSession) AsyncSafeTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.AsyncSafeTransferFrom(handler, &_IERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, *types.Receipt, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

func (_IERC721 *IERC721Transactor) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.contract.AsyncTransact(opts, handler, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, *types.Receipt, error) {
	return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts, from, to, tokenId, data)
}

func (_IERC721 *IERC721Session) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.AsyncSafeTransferFrom0(handler, &_IERC721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, *types.Receipt, error) {
	return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts, from, to, tokenId, data)
}

func (_IERC721 *IERC721TransactorSession) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.AsyncSafeTransferFrom0(handler, &_IERC721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, *types.Receipt, error) {
	return _IERC721.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

func (_IERC721 *IERC721Transactor) AsyncSetApprovalForAll(handler func(*types.Receipt, error), opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.contract.AsyncTransact(opts, handler, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721Session) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, *types.Receipt, error) {
	return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts, operator, _approved)
}

func (_IERC721 *IERC721Session) AsyncSetApprovalForAll(handler func(*types.Receipt, error), operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.Contract.AsyncSetApprovalForAll(handler, &_IERC721.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721TransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, *types.Receipt, error) {
	return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts, operator, _approved)
}

func (_IERC721 *IERC721TransactorSession) AsyncSetApprovalForAll(handler func(*types.Receipt, error), operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.Contract.AsyncSetApprovalForAll(handler, &_IERC721.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

func (_IERC721 *IERC721Transactor) AsyncTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.AsyncTransact(opts, handler, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

func (_IERC721 *IERC721Session) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.AsyncTransferFrom(handler, &_IERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

func (_IERC721 *IERC721TransactorSession) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.AsyncTransferFrom(handler, &_IERC721.TransactOpts, from, to, tokenId)
}

// IERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721 contract.
type IERC721ApprovalIterator struct {
	Event *IERC721Approval // Event containing the contract specifics and raw log

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
func (it *IERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721Approval)
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
		it.Event = new(IERC721Approval)
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
func (it *IERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721Approval represents a Approval event raised by the IERC721 contract.
type IERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721ApprovalIterator{contract: _IERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721Approval)
				if err := _IERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) ParseApproval(log types.Log) (*IERC721Approval, error) {
	event := new(IERC721Approval)
	if err := _IERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721 contract.
type IERC721ApprovalForAllIterator struct {
	Event *IERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721ApprovalForAll)
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
		it.Event = new(IERC721ApprovalForAll)
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
func (it *IERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721ApprovalForAll represents a ApprovalForAll event raised by the IERC721 contract.
type IERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC721ApprovalForAllIterator{contract: _IERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721ApprovalForAll)
				if err := _IERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) ParseApprovalForAll(log types.Log) (*IERC721ApprovalForAll, error) {
	event := new(IERC721ApprovalForAll)
	if err := _IERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721 contract.
type IERC721TransferIterator struct {
	Event *IERC721Transfer // Event containing the contract specifics and raw log

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
func (it *IERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721Transfer)
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
		it.Event = new(IERC721Transfer)
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
func (it *IERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721Transfer represents a Transfer event raised by the IERC721 contract.
type IERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721TransferIterator{contract: _IERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721Transfer)
				if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) ParseTransfer(log types.Log) (*IERC721Transfer, error) {
	event := new(IERC721Transfer)
	if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC721EnumerableABI is the input ABI used to generate the binding from.
const IERC721EnumerableABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// IERC721EnumerableFuncSigs maps the 4-byte function signature to its string representation.
var IERC721EnumerableFuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"4f6ccce7": "tokenByIndex(uint256)",
	"2f745c59": "tokenOfOwnerByIndex(address,uint256)",
	"18160ddd": "totalSupply()",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC721Enumerable is an auto generated Go binding around a Solidity contract.
type IERC721Enumerable struct {
	IERC721EnumerableCaller     // Read-only binding to the contract
	IERC721EnumerableTransactor // Write-only binding to the contract
	IERC721EnumerableFilterer   // Log filterer for contract events
}

// IERC721EnumerableCaller is an auto generated read-only Go binding around a Solidity contract.
type IERC721EnumerableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721EnumerableTransactor is an auto generated write-only Go binding around a Solidity contract.
type IERC721EnumerableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721EnumerableFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type IERC721EnumerableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721EnumerableSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type IERC721EnumerableSession struct {
	Contract     *IERC721Enumerable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC721EnumerableCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type IERC721EnumerableCallerSession struct {
	Contract *IERC721EnumerableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IERC721EnumerableTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type IERC721EnumerableTransactorSession struct {
	Contract     *IERC721EnumerableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IERC721EnumerableRaw is an auto generated low-level Go binding around a Solidity contract.
type IERC721EnumerableRaw struct {
	Contract *IERC721Enumerable // Generic contract binding to access the raw methods on
}

// IERC721EnumerableCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type IERC721EnumerableCallerRaw struct {
	Contract *IERC721EnumerableCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721EnumerableTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type IERC721EnumerableTransactorRaw struct {
	Contract *IERC721EnumerableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Enumerable creates a new instance of IERC721Enumerable, bound to a specific deployed contract.
func NewIERC721Enumerable(address common.Address, backend bind.ContractBackend) (*IERC721Enumerable, error) {
	contract, err := bindIERC721Enumerable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Enumerable{IERC721EnumerableCaller: IERC721EnumerableCaller{contract: contract}, IERC721EnumerableTransactor: IERC721EnumerableTransactor{contract: contract}, IERC721EnumerableFilterer: IERC721EnumerableFilterer{contract: contract}}, nil
}

// NewIERC721EnumerableCaller creates a new read-only instance of IERC721Enumerable, bound to a specific deployed contract.
func NewIERC721EnumerableCaller(address common.Address, caller bind.ContractCaller) (*IERC721EnumerableCaller, error) {
	contract, err := bindIERC721Enumerable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableCaller{contract: contract}, nil
}

// NewIERC721EnumerableTransactor creates a new write-only instance of IERC721Enumerable, bound to a specific deployed contract.
func NewIERC721EnumerableTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721EnumerableTransactor, error) {
	contract, err := bindIERC721Enumerable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableTransactor{contract: contract}, nil
}

// NewIERC721EnumerableFilterer creates a new log filterer instance of IERC721Enumerable, bound to a specific deployed contract.
func NewIERC721EnumerableFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721EnumerableFilterer, error) {
	contract, err := bindIERC721Enumerable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableFilterer{contract: contract}, nil
}

// bindIERC721Enumerable binds a generic wrapper to an already deployed contract.
func bindIERC721Enumerable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721EnumerableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Enumerable *IERC721EnumerableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC721Enumerable.Contract.IERC721EnumerableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Enumerable *IERC721EnumerableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Enumerable.Contract.IERC721EnumerableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Enumerable *IERC721EnumerableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Enumerable.Contract.IERC721EnumerableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Enumerable *IERC721EnumerableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC721Enumerable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Enumerable *IERC721EnumerableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Enumerable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Enumerable *IERC721EnumerableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Enumerable.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_IERC721Enumerable *IERC721EnumerableCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC721Enumerable.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_IERC721Enumerable *IERC721EnumerableSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721Enumerable.Contract.BalanceOf(&_IERC721Enumerable.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721Enumerable.Contract.BalanceOf(&_IERC721Enumerable.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address operator)
func (_IERC721Enumerable *IERC721EnumerableCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IERC721Enumerable.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address operator)
func (_IERC721Enumerable *IERC721EnumerableSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721Enumerable.Contract.GetApproved(&_IERC721Enumerable.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address operator)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721Enumerable.Contract.GetApproved(&_IERC721Enumerable.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_IERC721Enumerable *IERC721EnumerableCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IERC721Enumerable.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_IERC721Enumerable *IERC721EnumerableSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721Enumerable.Contract.IsApprovedForAll(&_IERC721Enumerable.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721Enumerable.Contract.IsApprovedForAll(&_IERC721Enumerable.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address owner)
func (_IERC721Enumerable *IERC721EnumerableCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IERC721Enumerable.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address owner)
func (_IERC721Enumerable *IERC721EnumerableSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721Enumerable.Contract.OwnerOf(&_IERC721Enumerable.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address owner)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721Enumerable.Contract.OwnerOf(&_IERC721Enumerable.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC721Enumerable *IERC721EnumerableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IERC721Enumerable.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC721Enumerable *IERC721EnumerableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721Enumerable.Contract.SupportsInterface(&_IERC721Enumerable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721Enumerable.Contract.SupportsInterface(&_IERC721Enumerable.CallOpts, interfaceId)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) constant returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC721Enumerable.contract.Call(opts, out, "tokenByIndex", index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) constant returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _IERC721Enumerable.Contract.TokenByIndex(&_IERC721Enumerable.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) constant returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _IERC721Enumerable.Contract.TokenByIndex(&_IERC721Enumerable.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) constant returns(uint256 tokenId)
func (_IERC721Enumerable *IERC721EnumerableCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC721Enumerable.contract.Call(opts, out, "tokenOfOwnerByIndex", owner, index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) constant returns(uint256 tokenId)
func (_IERC721Enumerable *IERC721EnumerableSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _IERC721Enumerable.Contract.TokenOfOwnerByIndex(&_IERC721Enumerable.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) constant returns(uint256 tokenId)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _IERC721Enumerable.Contract.TokenOfOwnerByIndex(&_IERC721Enumerable.CallOpts, owner, index)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC721Enumerable.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableSession) TotalSupply() (*big.Int, error) {
	return _IERC721Enumerable.Contract.TotalSupply(&_IERC721Enumerable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC721Enumerable *IERC721EnumerableCallerSession) TotalSupply() (*big.Int, error) {
	return _IERC721Enumerable.Contract.TotalSupply(&_IERC721Enumerable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Enumerable.contract.Transact(opts, "approve", to, tokenId)
}

func (_IERC721Enumerable *IERC721EnumerableTransactor) AsyncApprove(handler func(*types.Receipt, error), opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.contract.AsyncTransact(opts, handler, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Enumerable.Contract.Approve(&_IERC721Enumerable.TransactOpts, to, tokenId)
}

func (_IERC721Enumerable *IERC721EnumerableSession) AsyncApprove(handler func(*types.Receipt, error), to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.AsyncApprove(handler, &_IERC721Enumerable.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Enumerable.Contract.Approve(&_IERC721Enumerable.TransactOpts, to, tokenId)
}

func (_IERC721Enumerable *IERC721EnumerableTransactorSession) AsyncApprove(handler func(*types.Receipt, error), to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.AsyncApprove(handler, &_IERC721Enumerable.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Enumerable.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

func (_IERC721Enumerable *IERC721EnumerableTransactor) AsyncSafeTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.contract.AsyncTransact(opts, handler, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Enumerable.Contract.SafeTransferFrom(&_IERC721Enumerable.TransactOpts, from, to, tokenId)
}

func (_IERC721Enumerable *IERC721EnumerableSession) AsyncSafeTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.AsyncSafeTransferFrom(handler, &_IERC721Enumerable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Enumerable.Contract.SafeTransferFrom(&_IERC721Enumerable.TransactOpts, from, to, tokenId)
}

func (_IERC721Enumerable *IERC721EnumerableTransactorSession) AsyncSafeTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.AsyncSafeTransferFrom(handler, &_IERC721Enumerable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Enumerable.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

func (_IERC721Enumerable *IERC721EnumerableTransactor) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Enumerable.contract.AsyncTransact(opts, handler, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Enumerable *IERC721EnumerableSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Enumerable.Contract.SafeTransferFrom0(&_IERC721Enumerable.TransactOpts, from, to, tokenId, data)
}

func (_IERC721Enumerable *IERC721EnumerableSession) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.AsyncSafeTransferFrom0(handler, &_IERC721Enumerable.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Enumerable.Contract.SafeTransferFrom0(&_IERC721Enumerable.TransactOpts, from, to, tokenId, data)
}

func (_IERC721Enumerable *IERC721EnumerableTransactorSession) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.AsyncSafeTransferFrom0(handler, &_IERC721Enumerable.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Enumerable.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

func (_IERC721Enumerable *IERC721EnumerableTransactor) AsyncSetApprovalForAll(handler func(*types.Receipt, error), opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Enumerable.contract.AsyncTransact(opts, handler, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Enumerable *IERC721EnumerableSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Enumerable.Contract.SetApprovalForAll(&_IERC721Enumerable.TransactOpts, operator, _approved)
}

func (_IERC721Enumerable *IERC721EnumerableSession) AsyncSetApprovalForAll(handler func(*types.Receipt, error), operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.AsyncSetApprovalForAll(handler, &_IERC721Enumerable.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Enumerable.Contract.SetApprovalForAll(&_IERC721Enumerable.TransactOpts, operator, _approved)
}

func (_IERC721Enumerable *IERC721EnumerableTransactorSession) AsyncSetApprovalForAll(handler func(*types.Receipt, error), operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.AsyncSetApprovalForAll(handler, &_IERC721Enumerable.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Enumerable.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

func (_IERC721Enumerable *IERC721EnumerableTransactor) AsyncTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.contract.AsyncTransact(opts, handler, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Enumerable.Contract.TransferFrom(&_IERC721Enumerable.TransactOpts, from, to, tokenId)
}

func (_IERC721Enumerable *IERC721EnumerableSession) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.AsyncTransferFrom(handler, &_IERC721Enumerable.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Enumerable *IERC721EnumerableTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Enumerable.Contract.TransferFrom(&_IERC721Enumerable.TransactOpts, from, to, tokenId)
}

func (_IERC721Enumerable *IERC721EnumerableTransactorSession) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Enumerable.Contract.AsyncTransferFrom(handler, &_IERC721Enumerable.TransactOpts, from, to, tokenId)
}

// IERC721EnumerableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721Enumerable contract.
type IERC721EnumerableApprovalIterator struct {
	Event *IERC721EnumerableApproval // Event containing the contract specifics and raw log

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
func (it *IERC721EnumerableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721EnumerableApproval)
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
		it.Event = new(IERC721EnumerableApproval)
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
func (it *IERC721EnumerableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721EnumerableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721EnumerableApproval represents a Approval event raised by the IERC721Enumerable contract.
type IERC721EnumerableApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Enumerable *IERC721EnumerableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721EnumerableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Enumerable.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableApprovalIterator{contract: _IERC721Enumerable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Enumerable *IERC721EnumerableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721EnumerableApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Enumerable.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721EnumerableApproval)
				if err := _IERC721Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Enumerable *IERC721EnumerableFilterer) ParseApproval(log types.Log) (*IERC721EnumerableApproval, error) {
	event := new(IERC721EnumerableApproval)
	if err := _IERC721Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC721EnumerableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721Enumerable contract.
type IERC721EnumerableApprovalForAllIterator struct {
	Event *IERC721EnumerableApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC721EnumerableApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721EnumerableApprovalForAll)
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
		it.Event = new(IERC721EnumerableApprovalForAll)
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
func (it *IERC721EnumerableApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721EnumerableApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721EnumerableApprovalForAll represents a ApprovalForAll event raised by the IERC721Enumerable contract.
type IERC721EnumerableApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Enumerable *IERC721EnumerableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721EnumerableApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721Enumerable.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableApprovalForAllIterator{contract: _IERC721Enumerable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Enumerable *IERC721EnumerableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721EnumerableApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721Enumerable.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721EnumerableApprovalForAll)
				if err := _IERC721Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Enumerable *IERC721EnumerableFilterer) ParseApprovalForAll(log types.Log) (*IERC721EnumerableApprovalForAll, error) {
	event := new(IERC721EnumerableApprovalForAll)
	if err := _IERC721Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC721EnumerableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721Enumerable contract.
type IERC721EnumerableTransferIterator struct {
	Event *IERC721EnumerableTransfer // Event containing the contract specifics and raw log

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
func (it *IERC721EnumerableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721EnumerableTransfer)
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
		it.Event = new(IERC721EnumerableTransfer)
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
func (it *IERC721EnumerableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721EnumerableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721EnumerableTransfer represents a Transfer event raised by the IERC721Enumerable contract.
type IERC721EnumerableTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Enumerable *IERC721EnumerableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721EnumerableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Enumerable.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721EnumerableTransferIterator{contract: _IERC721Enumerable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Enumerable *IERC721EnumerableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721EnumerableTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Enumerable.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721EnumerableTransfer)
				if err := _IERC721Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Enumerable *IERC721EnumerableFilterer) ParseTransfer(log types.Log) (*IERC721EnumerableTransfer, error) {
	event := new(IERC721EnumerableTransfer)
	if err := _IERC721Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC721MetadataABI is the input ABI used to generate the binding from.
const IERC721MetadataABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// IERC721MetadataFuncSigs maps the 4-byte function signature to its string representation.
var IERC721MetadataFuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"06fdde03": "name()",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"95d89b41": "symbol()",
	"c87b56dd": "tokenURI(uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC721Metadata is an auto generated Go binding around a Solidity contract.
type IERC721Metadata struct {
	IERC721MetadataCaller     // Read-only binding to the contract
	IERC721MetadataTransactor // Write-only binding to the contract
	IERC721MetadataFilterer   // Log filterer for contract events
}

// IERC721MetadataCaller is an auto generated read-only Go binding around a Solidity contract.
type IERC721MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721MetadataTransactor is an auto generated write-only Go binding around a Solidity contract.
type IERC721MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721MetadataFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type IERC721MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721MetadataSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type IERC721MetadataSession struct {
	Contract     *IERC721Metadata  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721MetadataCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type IERC721MetadataCallerSession struct {
	Contract *IERC721MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IERC721MetadataTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type IERC721MetadataTransactorSession struct {
	Contract     *IERC721MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IERC721MetadataRaw is an auto generated low-level Go binding around a Solidity contract.
type IERC721MetadataRaw struct {
	Contract *IERC721Metadata // Generic contract binding to access the raw methods on
}

// IERC721MetadataCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type IERC721MetadataCallerRaw struct {
	Contract *IERC721MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721MetadataTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type IERC721MetadataTransactorRaw struct {
	Contract *IERC721MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Metadata creates a new instance of IERC721Metadata, bound to a specific deployed contract.
func NewIERC721Metadata(address common.Address, backend bind.ContractBackend) (*IERC721Metadata, error) {
	contract, err := bindIERC721Metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Metadata{IERC721MetadataCaller: IERC721MetadataCaller{contract: contract}, IERC721MetadataTransactor: IERC721MetadataTransactor{contract: contract}, IERC721MetadataFilterer: IERC721MetadataFilterer{contract: contract}}, nil
}

// NewIERC721MetadataCaller creates a new read-only instance of IERC721Metadata, bound to a specific deployed contract.
func NewIERC721MetadataCaller(address common.Address, caller bind.ContractCaller) (*IERC721MetadataCaller, error) {
	contract, err := bindIERC721Metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataCaller{contract: contract}, nil
}

// NewIERC721MetadataTransactor creates a new write-only instance of IERC721Metadata, bound to a specific deployed contract.
func NewIERC721MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721MetadataTransactor, error) {
	contract, err := bindIERC721Metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataTransactor{contract: contract}, nil
}

// NewIERC721MetadataFilterer creates a new log filterer instance of IERC721Metadata, bound to a specific deployed contract.
func NewIERC721MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721MetadataFilterer, error) {
	contract, err := bindIERC721Metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataFilterer{contract: contract}, nil
}

// bindIERC721Metadata binds a generic wrapper to an already deployed contract.
func bindIERC721Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721MetadataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Metadata *IERC721MetadataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC721Metadata.Contract.IERC721MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Metadata *IERC721MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Metadata.Contract.IERC721MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Metadata *IERC721MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Metadata.Contract.IERC721MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Metadata *IERC721MetadataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC721Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Metadata *IERC721MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Metadata *IERC721MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Metadata.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_IERC721Metadata *IERC721MetadataCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC721Metadata.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_IERC721Metadata *IERC721MetadataSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721Metadata.Contract.BalanceOf(&_IERC721Metadata.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_IERC721Metadata *IERC721MetadataCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721Metadata.Contract.BalanceOf(&_IERC721Metadata.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address operator)
func (_IERC721Metadata *IERC721MetadataCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IERC721Metadata.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address operator)
func (_IERC721Metadata *IERC721MetadataSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721Metadata.Contract.GetApproved(&_IERC721Metadata.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address operator)
func (_IERC721Metadata *IERC721MetadataCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721Metadata.Contract.GetApproved(&_IERC721Metadata.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_IERC721Metadata *IERC721MetadataCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IERC721Metadata.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_IERC721Metadata *IERC721MetadataSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721Metadata.Contract.IsApprovedForAll(&_IERC721Metadata.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_IERC721Metadata *IERC721MetadataCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721Metadata.Contract.IsApprovedForAll(&_IERC721Metadata.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_IERC721Metadata *IERC721MetadataCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IERC721Metadata.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_IERC721Metadata *IERC721MetadataSession) Name() (string, error) {
	return _IERC721Metadata.Contract.Name(&_IERC721Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_IERC721Metadata *IERC721MetadataCallerSession) Name() (string, error) {
	return _IERC721Metadata.Contract.Name(&_IERC721Metadata.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address owner)
func (_IERC721Metadata *IERC721MetadataCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IERC721Metadata.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address owner)
func (_IERC721Metadata *IERC721MetadataSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721Metadata.Contract.OwnerOf(&_IERC721Metadata.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address owner)
func (_IERC721Metadata *IERC721MetadataCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721Metadata.Contract.OwnerOf(&_IERC721Metadata.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC721Metadata *IERC721MetadataCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IERC721Metadata.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC721Metadata *IERC721MetadataSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721Metadata.Contract.SupportsInterface(&_IERC721Metadata.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC721Metadata *IERC721MetadataCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721Metadata.Contract.SupportsInterface(&_IERC721Metadata.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_IERC721Metadata *IERC721MetadataCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IERC721Metadata.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_IERC721Metadata *IERC721MetadataSession) Symbol() (string, error) {
	return _IERC721Metadata.Contract.Symbol(&_IERC721Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_IERC721Metadata *IERC721MetadataCallerSession) Symbol() (string, error) {
	return _IERC721Metadata.Contract.Symbol(&_IERC721Metadata.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) constant returns(string)
func (_IERC721Metadata *IERC721MetadataCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IERC721Metadata.contract.Call(opts, out, "tokenURI", tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) constant returns(string)
func (_IERC721Metadata *IERC721MetadataSession) TokenURI(tokenId *big.Int) (string, error) {
	return _IERC721Metadata.Contract.TokenURI(&_IERC721Metadata.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) constant returns(string)
func (_IERC721Metadata *IERC721MetadataCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _IERC721Metadata.Contract.TokenURI(&_IERC721Metadata.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Metadata.contract.Transact(opts, "approve", to, tokenId)
}

func (_IERC721Metadata *IERC721MetadataTransactor) AsyncApprove(handler func(*types.Receipt, error), opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.contract.AsyncTransact(opts, handler, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Metadata.Contract.Approve(&_IERC721Metadata.TransactOpts, to, tokenId)
}

func (_IERC721Metadata *IERC721MetadataSession) AsyncApprove(handler func(*types.Receipt, error), to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.AsyncApprove(handler, &_IERC721Metadata.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Metadata.Contract.Approve(&_IERC721Metadata.TransactOpts, to, tokenId)
}

func (_IERC721Metadata *IERC721MetadataTransactorSession) AsyncApprove(handler func(*types.Receipt, error), to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.AsyncApprove(handler, &_IERC721Metadata.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Metadata.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

func (_IERC721Metadata *IERC721MetadataTransactor) AsyncSafeTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.contract.AsyncTransact(opts, handler, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Metadata.Contract.SafeTransferFrom(&_IERC721Metadata.TransactOpts, from, to, tokenId)
}

func (_IERC721Metadata *IERC721MetadataSession) AsyncSafeTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.AsyncSafeTransferFrom(handler, &_IERC721Metadata.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Metadata.Contract.SafeTransferFrom(&_IERC721Metadata.TransactOpts, from, to, tokenId)
}

func (_IERC721Metadata *IERC721MetadataTransactorSession) AsyncSafeTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.AsyncSafeTransferFrom(handler, &_IERC721Metadata.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Metadata.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

func (_IERC721Metadata *IERC721MetadataTransactor) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Metadata.contract.AsyncTransact(opts, handler, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Metadata *IERC721MetadataSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Metadata.Contract.SafeTransferFrom0(&_IERC721Metadata.TransactOpts, from, to, tokenId, data)
}

func (_IERC721Metadata *IERC721MetadataSession) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.AsyncSafeTransferFrom0(handler, &_IERC721Metadata.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Metadata.Contract.SafeTransferFrom0(&_IERC721Metadata.TransactOpts, from, to, tokenId, data)
}

func (_IERC721Metadata *IERC721MetadataTransactorSession) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.AsyncSafeTransferFrom0(handler, &_IERC721Metadata.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Metadata.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

func (_IERC721Metadata *IERC721MetadataTransactor) AsyncSetApprovalForAll(handler func(*types.Receipt, error), opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Metadata.contract.AsyncTransact(opts, handler, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Metadata *IERC721MetadataSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Metadata.Contract.SetApprovalForAll(&_IERC721Metadata.TransactOpts, operator, _approved)
}

func (_IERC721Metadata *IERC721MetadataSession) AsyncSetApprovalForAll(handler func(*types.Receipt, error), operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.AsyncSetApprovalForAll(handler, &_IERC721Metadata.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Metadata.Contract.SetApprovalForAll(&_IERC721Metadata.TransactOpts, operator, _approved)
}

func (_IERC721Metadata *IERC721MetadataTransactorSession) AsyncSetApprovalForAll(handler func(*types.Receipt, error), operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.AsyncSetApprovalForAll(handler, &_IERC721Metadata.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Metadata.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

func (_IERC721Metadata *IERC721MetadataTransactor) AsyncTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.contract.AsyncTransact(opts, handler, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Metadata.Contract.TransferFrom(&_IERC721Metadata.TransactOpts, from, to, tokenId)
}

func (_IERC721Metadata *IERC721MetadataSession) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.AsyncTransferFrom(handler, &_IERC721Metadata.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721Metadata *IERC721MetadataTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Metadata.Contract.TransferFrom(&_IERC721Metadata.TransactOpts, from, to, tokenId)
}

func (_IERC721Metadata *IERC721MetadataTransactorSession) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721Metadata.Contract.AsyncTransferFrom(handler, &_IERC721Metadata.TransactOpts, from, to, tokenId)
}

// IERC721MetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721Metadata contract.
type IERC721MetadataApprovalIterator struct {
	Event *IERC721MetadataApproval // Event containing the contract specifics and raw log

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
func (it *IERC721MetadataApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721MetadataApproval)
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
		it.Event = new(IERC721MetadataApproval)
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
func (it *IERC721MetadataApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721MetadataApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721MetadataApproval represents a Approval event raised by the IERC721Metadata contract.
type IERC721MetadataApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721MetadataApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Metadata.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataApprovalIterator{contract: _IERC721Metadata.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721MetadataApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Metadata.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721MetadataApproval)
				if err := _IERC721Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) ParseApproval(log types.Log) (*IERC721MetadataApproval, error) {
	event := new(IERC721MetadataApproval)
	if err := _IERC721Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC721MetadataApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721Metadata contract.
type IERC721MetadataApprovalForAllIterator struct {
	Event *IERC721MetadataApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC721MetadataApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721MetadataApprovalForAll)
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
		it.Event = new(IERC721MetadataApprovalForAll)
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
func (it *IERC721MetadataApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721MetadataApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721MetadataApprovalForAll represents a ApprovalForAll event raised by the IERC721Metadata contract.
type IERC721MetadataApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Metadata *IERC721MetadataFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721MetadataApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721Metadata.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataApprovalForAllIterator{contract: _IERC721Metadata.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Metadata *IERC721MetadataFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721MetadataApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721Metadata.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721MetadataApprovalForAll)
				if err := _IERC721Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721Metadata *IERC721MetadataFilterer) ParseApprovalForAll(log types.Log) (*IERC721MetadataApprovalForAll, error) {
	event := new(IERC721MetadataApprovalForAll)
	if err := _IERC721Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC721MetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721Metadata contract.
type IERC721MetadataTransferIterator struct {
	Event *IERC721MetadataTransfer // Event containing the contract specifics and raw log

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
func (it *IERC721MetadataTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721MetadataTransfer)
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
		it.Event = new(IERC721MetadataTransfer)
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
func (it *IERC721MetadataTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721MetadataTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721MetadataTransfer represents a Transfer event raised by the IERC721Metadata contract.
type IERC721MetadataTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721MetadataTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Metadata.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721MetadataTransferIterator{contract: _IERC721Metadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721MetadataTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721Metadata.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721MetadataTransfer)
				if err := _IERC721Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721Metadata *IERC721MetadataFilterer) ParseTransfer(log types.Log) (*IERC721MetadataTransfer, error) {
	event := new(IERC721MetadataTransfer)
	if err := _IERC721Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC721ReceiverABI is the input ABI used to generate the binding from.
const IERC721ReceiverABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC721ReceiverFuncSigs maps the 4-byte function signature to its string representation.
var IERC721ReceiverFuncSigs = map[string]string{
	"150b7a02": "onERC721Received(address,address,uint256,bytes)",
}

// IERC721Receiver is an auto generated Go binding around a Solidity contract.
type IERC721Receiver struct {
	IERC721ReceiverCaller     // Read-only binding to the contract
	IERC721ReceiverTransactor // Write-only binding to the contract
	IERC721ReceiverFilterer   // Log filterer for contract events
}

// IERC721ReceiverCaller is an auto generated read-only Go binding around a Solidity contract.
type IERC721ReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ReceiverTransactor is an auto generated write-only Go binding around a Solidity contract.
type IERC721ReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ReceiverFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type IERC721ReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721ReceiverSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type IERC721ReceiverSession struct {
	Contract     *IERC721Receiver  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721ReceiverCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type IERC721ReceiverCallerSession struct {
	Contract *IERC721ReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IERC721ReceiverTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type IERC721ReceiverTransactorSession struct {
	Contract     *IERC721ReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IERC721ReceiverRaw is an auto generated low-level Go binding around a Solidity contract.
type IERC721ReceiverRaw struct {
	Contract *IERC721Receiver // Generic contract binding to access the raw methods on
}

// IERC721ReceiverCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type IERC721ReceiverCallerRaw struct {
	Contract *IERC721ReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IERC721ReceiverTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type IERC721ReceiverTransactorRaw struct {
	Contract *IERC721ReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721Receiver creates a new instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721Receiver(address common.Address, backend bind.ContractBackend) (*IERC721Receiver, error) {
	contract, err := bindIERC721Receiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721Receiver{IERC721ReceiverCaller: IERC721ReceiverCaller{contract: contract}, IERC721ReceiverTransactor: IERC721ReceiverTransactor{contract: contract}, IERC721ReceiverFilterer: IERC721ReceiverFilterer{contract: contract}}, nil
}

// NewIERC721ReceiverCaller creates a new read-only instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721ReceiverCaller(address common.Address, caller bind.ContractCaller) (*IERC721ReceiverCaller, error) {
	contract, err := bindIERC721Receiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721ReceiverCaller{contract: contract}, nil
}

// NewIERC721ReceiverTransactor creates a new write-only instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721ReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721ReceiverTransactor, error) {
	contract, err := bindIERC721Receiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721ReceiverTransactor{contract: contract}, nil
}

// NewIERC721ReceiverFilterer creates a new log filterer instance of IERC721Receiver, bound to a specific deployed contract.
func NewIERC721ReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721ReceiverFilterer, error) {
	contract, err := bindIERC721Receiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721ReceiverFilterer{contract: contract}, nil
}

// bindIERC721Receiver binds a generic wrapper to an already deployed contract.
func bindIERC721Receiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721ReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Receiver *IERC721ReceiverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC721Receiver.Contract.IERC721ReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Receiver *IERC721ReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Receiver.Contract.IERC721ReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Receiver *IERC721ReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Receiver.Contract.IERC721ReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721Receiver *IERC721ReceiverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC721Receiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721Receiver *IERC721ReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Receiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721Receiver *IERC721ReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Receiver.Contract.contract.Transact(opts, method, params...)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IERC721Receiver *IERC721ReceiverTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Receiver.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

func (_IERC721Receiver *IERC721ReceiverTransactor) AsyncOnERC721Received(handler func(*types.Receipt, error), opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Receiver.contract.AsyncTransact(opts, handler, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IERC721Receiver *IERC721ReceiverSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Receiver.Contract.OnERC721Received(&_IERC721Receiver.TransactOpts, operator, from, tokenId, data)
}

func (_IERC721Receiver *IERC721ReceiverSession) AsyncOnERC721Received(handler func(*types.Receipt, error), operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.AsyncOnERC721Received(handler, &_IERC721Receiver.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_IERC721Receiver *IERC721ReceiverTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, *types.Receipt, error) {
	return _IERC721Receiver.Contract.OnERC721Received(&_IERC721Receiver.TransactOpts, operator, from, tokenId, data)
}

func (_IERC721Receiver *IERC721ReceiverTransactorSession) AsyncOnERC721Received(handler func(*types.Receipt, error), operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721Receiver.Contract.AsyncOnERC721Received(handler, &_IERC721Receiver.TransactOpts, operator, from, tokenId, data)
}

// MinterRoleABI is the input ABI used to generate the binding from.
const MinterRoleABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"}]"

// MinterRoleFuncSigs maps the 4-byte function signature to its string representation.
var MinterRoleFuncSigs = map[string]string{
	"983b2d56": "addMinter(address)",
	"aa271e1a": "isMinter(address)",
	"98650275": "renounceMinter()",
}

// MinterRole is an auto generated Go binding around a Solidity contract.
type MinterRole struct {
	MinterRoleCaller     // Read-only binding to the contract
	MinterRoleTransactor // Write-only binding to the contract
	MinterRoleFilterer   // Log filterer for contract events
}

// MinterRoleCaller is an auto generated read-only Go binding around a Solidity contract.
type MinterRoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterRoleTransactor is an auto generated write-only Go binding around a Solidity contract.
type MinterRoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterRoleFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type MinterRoleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterRoleSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type MinterRoleSession struct {
	Contract     *MinterRole       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MinterRoleCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type MinterRoleCallerSession struct {
	Contract *MinterRoleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MinterRoleTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type MinterRoleTransactorSession struct {
	Contract     *MinterRoleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MinterRoleRaw is an auto generated low-level Go binding around a Solidity contract.
type MinterRoleRaw struct {
	Contract *MinterRole // Generic contract binding to access the raw methods on
}

// MinterRoleCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type MinterRoleCallerRaw struct {
	Contract *MinterRoleCaller // Generic read-only contract binding to access the raw methods on
}

// MinterRoleTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type MinterRoleTransactorRaw struct {
	Contract *MinterRoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMinterRole creates a new instance of MinterRole, bound to a specific deployed contract.
func NewMinterRole(address common.Address, backend bind.ContractBackend) (*MinterRole, error) {
	contract, err := bindMinterRole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MinterRole{MinterRoleCaller: MinterRoleCaller{contract: contract}, MinterRoleTransactor: MinterRoleTransactor{contract: contract}, MinterRoleFilterer: MinterRoleFilterer{contract: contract}}, nil
}

// NewMinterRoleCaller creates a new read-only instance of MinterRole, bound to a specific deployed contract.
func NewMinterRoleCaller(address common.Address, caller bind.ContractCaller) (*MinterRoleCaller, error) {
	contract, err := bindMinterRole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MinterRoleCaller{contract: contract}, nil
}

// NewMinterRoleTransactor creates a new write-only instance of MinterRole, bound to a specific deployed contract.
func NewMinterRoleTransactor(address common.Address, transactor bind.ContractTransactor) (*MinterRoleTransactor, error) {
	contract, err := bindMinterRole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MinterRoleTransactor{contract: contract}, nil
}

// NewMinterRoleFilterer creates a new log filterer instance of MinterRole, bound to a specific deployed contract.
func NewMinterRoleFilterer(address common.Address, filterer bind.ContractFilterer) (*MinterRoleFilterer, error) {
	contract, err := bindMinterRole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MinterRoleFilterer{contract: contract}, nil
}

// bindMinterRole binds a generic wrapper to an already deployed contract.
func bindMinterRole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MinterRoleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinterRole *MinterRoleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MinterRole.Contract.MinterRoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinterRole *MinterRoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _MinterRole.Contract.MinterRoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinterRole *MinterRoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _MinterRole.Contract.MinterRoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinterRole *MinterRoleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MinterRole.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinterRole *MinterRoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _MinterRole.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinterRole *MinterRoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _MinterRole.Contract.contract.Transact(opts, method, params...)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_MinterRole *MinterRoleCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MinterRole.contract.Call(opts, out, "isMinter", account)
	return *ret0, err
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_MinterRole *MinterRoleSession) IsMinter(account common.Address) (bool, error) {
	return _MinterRole.Contract.IsMinter(&_MinterRole.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_MinterRole *MinterRoleCallerSession) IsMinter(account common.Address) (bool, error) {
	return _MinterRole.Contract.IsMinter(&_MinterRole.CallOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_MinterRole *MinterRoleTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _MinterRole.contract.Transact(opts, "addMinter", account)
}

func (_MinterRole *MinterRoleTransactor) AsyncAddMinter(handler func(*types.Receipt, error), opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _MinterRole.contract.AsyncTransact(opts, handler, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_MinterRole *MinterRoleSession) AddMinter(account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _MinterRole.Contract.AddMinter(&_MinterRole.TransactOpts, account)
}

func (_MinterRole *MinterRoleSession) AsyncAddMinter(handler func(*types.Receipt, error), account common.Address) (*types.Transaction, error) {
	return _MinterRole.Contract.AsyncAddMinter(handler, &_MinterRole.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_MinterRole *MinterRoleTransactorSession) AddMinter(account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _MinterRole.Contract.AddMinter(&_MinterRole.TransactOpts, account)
}

func (_MinterRole *MinterRoleTransactorSession) AsyncAddMinter(handler func(*types.Receipt, error), account common.Address) (*types.Transaction, error) {
	return _MinterRole.Contract.AsyncAddMinter(handler, &_MinterRole.TransactOpts, account)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_MinterRole *MinterRoleTransactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _MinterRole.contract.Transact(opts, "renounceMinter")
}

func (_MinterRole *MinterRoleTransactor) AsyncRenounceMinter(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterRole.contract.AsyncTransact(opts, handler, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_MinterRole *MinterRoleSession) RenounceMinter() (*types.Transaction, *types.Receipt, error) {
	return _MinterRole.Contract.RenounceMinter(&_MinterRole.TransactOpts)
}

func (_MinterRole *MinterRoleSession) AsyncRenounceMinter(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _MinterRole.Contract.AsyncRenounceMinter(handler, &_MinterRole.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_MinterRole *MinterRoleTransactorSession) RenounceMinter() (*types.Transaction, *types.Receipt, error) {
	return _MinterRole.Contract.RenounceMinter(&_MinterRole.TransactOpts)
}

func (_MinterRole *MinterRoleTransactorSession) AsyncRenounceMinter(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _MinterRole.Contract.AsyncRenounceMinter(handler, &_MinterRole.TransactOpts)
}

// MinterRoleMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the MinterRole contract.
type MinterRoleMinterAddedIterator struct {
	Event *MinterRoleMinterAdded // Event containing the contract specifics and raw log

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
func (it *MinterRoleMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterRoleMinterAdded)
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
		it.Event = new(MinterRoleMinterAdded)
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
func (it *MinterRoleMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterRoleMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterRoleMinterAdded represents a MinterAdded event raised by the MinterRole contract.
type MinterRoleMinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_MinterRole *MinterRoleFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*MinterRoleMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MinterRole.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &MinterRoleMinterAddedIterator{contract: _MinterRole.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_MinterRole *MinterRoleFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *MinterRoleMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MinterRole.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterRoleMinterAdded)
				if err := _MinterRole.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// ParseMinterAdded is a log parse operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_MinterRole *MinterRoleFilterer) ParseMinterAdded(log types.Log) (*MinterRoleMinterAdded, error) {
	event := new(MinterRoleMinterAdded)
	if err := _MinterRole.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MinterRoleMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the MinterRole contract.
type MinterRoleMinterRemovedIterator struct {
	Event *MinterRoleMinterRemoved // Event containing the contract specifics and raw log

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
func (it *MinterRoleMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterRoleMinterRemoved)
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
		it.Event = new(MinterRoleMinterRemoved)
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
func (it *MinterRoleMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterRoleMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterRoleMinterRemoved represents a MinterRemoved event raised by the MinterRole contract.
type MinterRoleMinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_MinterRole *MinterRoleFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*MinterRoleMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MinterRole.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &MinterRoleMinterRemovedIterator{contract: _MinterRole.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_MinterRole *MinterRoleFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *MinterRoleMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MinterRole.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterRoleMinterRemoved)
				if err := _MinterRole.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_MinterRole *MinterRoleFilterer) ParseMinterRemoved(log types.Log) (*MinterRoleMinterRemoved, error) {
	event := new(MinterRoleMinterRemoved)
	if err := _MinterRole.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RolesABI is the input ABI used to generate the binding from.
const RolesABI = "[]"

// RolesBin is the compiled bytecode used for deploying new contracts.
var RolesBin = "0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820a975f8d5082078fdcf359b5892ef0701e90ed511058625db1f6c0f34f225d3500029"

// DeployRoles deploys a new contract, binding an instance of Roles to it.
func DeployRoles(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Roles, error) {
	parsed, err := abi.JSON(strings.NewReader(RolesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RolesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Roles{RolesCaller: RolesCaller{contract: contract}, RolesTransactor: RolesTransactor{contract: contract}, RolesFilterer: RolesFilterer{contract: contract}}, nil
}

func AsyncDeployRoles(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(RolesABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(RolesBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Roles is an auto generated Go binding around a Solidity contract.
type Roles struct {
	RolesCaller     // Read-only binding to the contract
	RolesTransactor // Write-only binding to the contract
	RolesFilterer   // Log filterer for contract events
}

// RolesCaller is an auto generated read-only Go binding around a Solidity contract.
type RolesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesTransactor is an auto generated write-only Go binding around a Solidity contract.
type RolesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type RolesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type RolesSession struct {
	Contract     *Roles            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RolesCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type RolesCallerSession struct {
	Contract *RolesCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RolesTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type RolesTransactorSession struct {
	Contract     *RolesTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RolesRaw is an auto generated low-level Go binding around a Solidity contract.
type RolesRaw struct {
	Contract *Roles // Generic contract binding to access the raw methods on
}

// RolesCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type RolesCallerRaw struct {
	Contract *RolesCaller // Generic read-only contract binding to access the raw methods on
}

// RolesTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type RolesTransactorRaw struct {
	Contract *RolesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoles creates a new instance of Roles, bound to a specific deployed contract.
func NewRoles(address common.Address, backend bind.ContractBackend) (*Roles, error) {
	contract, err := bindRoles(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Roles{RolesCaller: RolesCaller{contract: contract}, RolesTransactor: RolesTransactor{contract: contract}, RolesFilterer: RolesFilterer{contract: contract}}, nil
}

// NewRolesCaller creates a new read-only instance of Roles, bound to a specific deployed contract.
func NewRolesCaller(address common.Address, caller bind.ContractCaller) (*RolesCaller, error) {
	contract, err := bindRoles(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RolesCaller{contract: contract}, nil
}

// NewRolesTransactor creates a new write-only instance of Roles, bound to a specific deployed contract.
func NewRolesTransactor(address common.Address, transactor bind.ContractTransactor) (*RolesTransactor, error) {
	contract, err := bindRoles(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RolesTransactor{contract: contract}, nil
}

// NewRolesFilterer creates a new log filterer instance of Roles, bound to a specific deployed contract.
func NewRolesFilterer(address common.Address, filterer bind.ContractFilterer) (*RolesFilterer, error) {
	contract, err := bindRoles(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RolesFilterer{contract: contract}, nil
}

// bindRoles binds a generic wrapper to an already deployed contract.
func bindRoles(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RolesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Roles *RolesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Roles.Contract.RolesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Roles *RolesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Roles.Contract.RolesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Roles *RolesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Roles.Contract.RolesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Roles *RolesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Roles.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Roles *RolesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Roles.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Roles *RolesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Roles.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820d47a7dbb112d1330fd10222df28f29c0f78ea9d6e4b40656afe05a7b14973a630029"

// DeploySafeMath deploys a new contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

func AsyncDeploySafeMath(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// SafeMath is an auto generated Go binding around a Solidity contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around a Solidity contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around a Solidity contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around a Solidity contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// TTPNftABI is the input ABI used to generate the binding from.
const TTPNftABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"mintWithTokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// TTPNftFuncSigs maps the 4-byte function signature to its string representation.
var TTPNftFuncSigs = map[string]string{
	"983b2d56": "addMinter(address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"aa271e1a": "isMinter(address)",
	"50bb4e7f": "mintWithTokenURI(address,uint256,string)",
	"06fdde03": "name()",
	"6352211e": "ownerOf(uint256)",
	"98650275": "renounceMinter()",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"95d89b41": "symbol()",
	"4f6ccce7": "tokenByIndex(uint256)",
	"2f745c59": "tokenOfOwnerByIndex(address,uint256)",
	"c87b56dd": "tokenURI(uint256)",
	"18160ddd": "totalSupply()",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// TTPNftBin is the compiled bytecode used for deploying new contracts.
var TTPNftBin = "0x60806040523480156200001157600080fd5b50604080518082018252600681527f5454504e465400000000000000000000000000000000000000000000000000006020808301919091528251808401909352600383527f4e4654000000000000000000000000000000000000000000000000000000000090830152908181620000b17f01ffc9a7000000000000000000000000000000000000000000000000000000006401000000006200019d810204565b620000e57f80ac58cd000000000000000000000000000000000000000000000000000000006401000000006200019d810204565b620001197f780e9d63000000000000000000000000000000000000000000000000000000006401000000006200019d810204565b81516200012e906009906020850190620002ef565b5080516200014490600a906020840190620002ef565b50620001797f5b5e139f000000000000000000000000000000000000000000000000000000006401000000006200019d810204565b5050505062000197336200020a640100000000026401000000009004565b62000394565b7fffffffff000000000000000000000000000000000000000000000000000000008082161415620001cd57600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b62000225600c82640100000000620011276200025c82021704565b604051600160a060020a038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b600160a060020a03811615156200027257600080fd5b620002878282640100000000620002b7810204565b156200029257600080fd5b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b6000600160a060020a0382161515620002cf57600080fd5b50600160a060020a03166000908152602091909152604090205460ff1690565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200033257805160ff191683800117855562000362565b8280016001018555821562000362579182015b828111156200036257825182559160200191906001019062000345565b506200037092915062000374565b5090565b6200039191905b808211156200037057600081556001016200037b565b90565b6112ca80620003a46000396000f3006080604052600436106101115763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301ffc9a7811461011657806306fdde0314610161578063081812fc146101eb578063095ea7b31461021f57806318160ddd1461024557806323b872dd1461026c5780632f745c591461029657806342842e0e146102ba5780634f6ccce7146102e457806350bb4e7f146102fc5780636352211e1461036557806370a082311461037d57806395d89b411461039e578063983b2d56146103b357806398650275146103d4578063a22cb465146103e9578063aa271e1a1461040f578063b88d4fde14610430578063c87b56dd1461049f578063e985e9c5146104b7575b600080fd5b34801561012257600080fd5b5061014d7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19600435166104de565b604080519115158252519081900360200190f35b34801561016d57600080fd5b50610176610512565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101b0578181015183820152602001610198565b50505050905090810190601f1680156101dd5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101f757600080fd5b506102036004356105a9565b60408051600160a060020a039092168252519081900360200190f35b34801561022b57600080fd5b50610243600160a060020a03600435166024356105db565b005b34801561025157600080fd5b5061025a610691565b60408051918252519081900360200190f35b34801561027857600080fd5b50610243600160a060020a0360043581169060243516604435610697565b3480156102a257600080fd5b5061025a600160a060020a0360043516602435610725565b3480156102c657600080fd5b50610243600160a060020a0360043581169060243516604435610772565b3480156102f057600080fd5b5061025a600435610793565b34801561030857600080fd5b50604080516020600460443581810135601f810184900484028501840190955284845261014d948235600160a060020a03169460248035953695946064949201919081908401838280828437509497506107c89650505050505050565b34801561037157600080fd5b506102036004356107fc565b34801561038957600080fd5b5061025a600160a060020a0360043516610826565b3480156103aa57600080fd5b50610176610859565b3480156103bf57600080fd5b50610243600160a060020a03600435166108ba565b3480156103e057600080fd5b506102436108da565b3480156103f557600080fd5b50610243600160a060020a036004351660243515156108e5565b34801561041b57600080fd5b5061014d600160a060020a0360043516610969565b34801561043c57600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261024394600160a060020a03813581169560248035909216956044359536956084940191819084018382808284375094975061097c9650505050505050565b3480156104ab57600080fd5b506101766004356109a4565b3480156104c357600080fd5b5061014d600160a060020a0360043581169060243516610a59565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191660009081526020819052604090205460ff1690565b60098054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561059e5780601f106105735761010080835404028352916020019161059e565b820191906000526020600020905b81548152906001019060200180831161058157829003601f168201915b505050505090505b90565b60006105b482610a87565b15156105bf57600080fd5b50600090815260026020526040902054600160a060020a031690565b60006105e6826107fc565b9050600160a060020a03838116908216141561060157600080fd5b33600160a060020a038216148061061d575061061d8133610a59565b151561062857600080fd5b600082815260026020526040808220805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b60075490565b6106a13382610aa4565b15156106ac57600080fd5b600160a060020a03821615156106c157600080fd5b6106cb8382610b03565b6106d58382610b74565b6106df8282610c7b565b8082600160a060020a031684600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b600061073083610826565b821061073b57600080fd5b600160a060020a038316600090815260056020526040902080548390811061075f57fe5b9060005260206000200154905092915050565b61078e838383602060405190810160405280600081525061097c565b505050565b600061079d610691565b82106107a857600080fd5b60078054839081106107b657fe5b90600052602060002001549050919050565b60006107d333610969565b15156107de57600080fd5b6107e88484610cc4565b6107f28383610d13565b5060019392505050565b600081815260016020526040812054600160a060020a031680151561082057600080fd5b92915050565b6000600160a060020a038216151561083d57600080fd5b50600160a060020a031660009081526003602052604090205490565b600a8054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561059e5780601f106105735761010080835404028352916020019161059e565b6108c333610969565b15156108ce57600080fd5b6108d781610d46565b50565b6108e333610d8e565b565b600160a060020a0382163314156108fb57600080fd5b336000818152600460209081526040808320600160a060020a03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b6000610820600c8363ffffffff610dd616565b610987848484610697565b61099384848484610e0d565b151561099e57600080fd5b50505050565b60606109af82610a87565b15156109ba57600080fd5b6000828152600b602090815260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015610a4d5780601f10610a2257610100808354040283529160200191610a4d565b820191906000526020600020905b815481529060010190602001808311610a3057829003601f168201915b50505050509050919050565b600160a060020a03918216600090815260046020908152604080832093909416825291909152205460ff1690565b600090815260016020526040902054600160a060020a0316151590565b600080610ab0836107fc565b905080600160a060020a031684600160a060020a03161480610aeb575083600160a060020a0316610ae0846105a9565b600160a060020a0316145b80610afb5750610afb8185610a59565b949350505050565b81600160a060020a0316610b16826107fc565b600160a060020a031614610b2957600080fd5b600081815260026020526040902054600160a060020a031615610b70576000818152600260205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b5050565b6000806000610b838585610f8f565b600084815260066020908152604080832054600160a060020a0389168452600590925290912054909350610bbe90600163ffffffff61102516565b600160a060020a038616600090815260056020526040902080549193509083908110610be657fe5b90600052602060002001549050806005600087600160a060020a0316600160a060020a0316815260200190815260200160002084815481101515610c2657fe5b6000918252602080832090910192909255600160a060020a0387168152600590915260409020805490610c5d9060001983016111e2565b50600093845260066020526040808520859055908452909220555050565b6000610c87838361103c565b50600160a060020a039091166000908152600560209081526040808320805460018101825590845282842081018590559383526006909152902055565b610cce82826110cc565b600780546000838152600860205260408120829055600182018355919091527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688015550565b610d1c82610a87565b1515610d2757600080fd5b6000828152600b60209081526040909120825161078e92840190611206565b610d57600c8263ffffffff61112716565b604051600160a060020a038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b610d9f600c8263ffffffff61117516565b604051600160a060020a038216907fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669290600090a250565b6000600160a060020a0382161515610ded57600080fd5b50600160a060020a03166000908152602091909152604090205460ff1690565b600080610e2285600160a060020a03166111c1565b1515610e315760019150610f86565b6040517f150b7a020000000000000000000000000000000000000000000000000000000081523360048201818152600160a060020a03898116602485015260448401889052608060648501908152875160848601528751918a169463150b7a0294938c938b938b93909160a490910190602085019080838360005b83811015610ec4578181015183820152602001610eac565b50505050905090810190601f168015610ef15780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015610f1357600080fd5b505af1158015610f27573d6000803e3d6000fd5b505050506040513d6020811015610f3d57600080fd5b50517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1981167f150b7a020000000000000000000000000000000000000000000000000000000014925090505b50949350505050565b81600160a060020a0316610fa2826107fc565b600160a060020a031614610fb557600080fd5b600160a060020a038216600090815260036020526040902054610fdf90600163ffffffff61102516565b600160a060020a03909216600090815260036020908152604080832094909455918152600190915220805473ffffffffffffffffffffffffffffffffffffffff19169055565b6000808383111561103557600080fd5b5050900390565b600081815260016020526040902054600160a060020a03161561105e57600080fd5b6000818152600160208181526040808420805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03881690811790915584526003909152909120546110ac916111c9565b600160a060020a0390921660009081526003602052604090209190915550565b600160a060020a03821615156110e157600080fd5b6110eb8282610c7b565b6040518190600160a060020a038416906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b600160a060020a038116151561113c57600080fd5b6111468282610dd6565b1561115057600080fd5b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b600160a060020a038116151561118a57600080fd5b6111948282610dd6565b151561119f57600080fd5b600160a060020a0316600090815260209190915260409020805460ff19169055565b6000903b1190565b6000828201838110156111db57600080fd5b9392505050565b81548183558181111561078e5760008381526020902061078e918101908301611284565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061124757805160ff1916838001178555611274565b82800160010185558215611274579182015b82811115611274578251825591602001919060010190611259565b50611280929150611284565b5090565b6105a691905b80821115611280576000815560010161128a5600a165627a7a72305820d60fc5c1065ee9202d7e155d9c9e43f1415e2a2bf3c066515d365c96b5eda5110029"

// DeployTTPNft deploys a new contract, binding an instance of TTPNft to it.
func DeployTTPNft(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TTPNft, error) {
	parsed, err := abi.JSON(strings.NewReader(TTPNftABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TTPNftBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TTPNft{TTPNftCaller: TTPNftCaller{contract: contract}, TTPNftTransactor: TTPNftTransactor{contract: contract}, TTPNftFilterer: TTPNftFilterer{contract: contract}}, nil
}

func AsyncDeployTTPNft(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(TTPNftABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(TTPNftBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// TTPNft is an auto generated Go binding around a Solidity contract.
type TTPNft struct {
	TTPNftCaller     // Read-only binding to the contract
	TTPNftTransactor // Write-only binding to the contract
	TTPNftFilterer   // Log filterer for contract events
}

// TTPNftCaller is an auto generated read-only Go binding around a Solidity contract.
type TTPNftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TTPNftTransactor is an auto generated write-only Go binding around a Solidity contract.
type TTPNftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TTPNftFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type TTPNftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TTPNftSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type TTPNftSession struct {
	Contract     *TTPNft           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TTPNftCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type TTPNftCallerSession struct {
	Contract *TTPNftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TTPNftTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type TTPNftTransactorSession struct {
	Contract     *TTPNftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TTPNftRaw is an auto generated low-level Go binding around a Solidity contract.
type TTPNftRaw struct {
	Contract *TTPNft // Generic contract binding to access the raw methods on
}

// TTPNftCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type TTPNftCallerRaw struct {
	Contract *TTPNftCaller // Generic read-only contract binding to access the raw methods on
}

// TTPNftTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type TTPNftTransactorRaw struct {
	Contract *TTPNftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTTPNft creates a new instance of TTPNft, bound to a specific deployed contract.
func NewTTPNft(address common.Address, backend bind.ContractBackend) (*TTPNft, error) {
	contract, err := bindTTPNft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TTPNft{TTPNftCaller: TTPNftCaller{contract: contract}, TTPNftTransactor: TTPNftTransactor{contract: contract}, TTPNftFilterer: TTPNftFilterer{contract: contract}}, nil
}

// NewTTPNftCaller creates a new read-only instance of TTPNft, bound to a specific deployed contract.
func NewTTPNftCaller(address common.Address, caller bind.ContractCaller) (*TTPNftCaller, error) {
	contract, err := bindTTPNft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TTPNftCaller{contract: contract}, nil
}

// NewTTPNftTransactor creates a new write-only instance of TTPNft, bound to a specific deployed contract.
func NewTTPNftTransactor(address common.Address, transactor bind.ContractTransactor) (*TTPNftTransactor, error) {
	contract, err := bindTTPNft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TTPNftTransactor{contract: contract}, nil
}

// NewTTPNftFilterer creates a new log filterer instance of TTPNft, bound to a specific deployed contract.
func NewTTPNftFilterer(address common.Address, filterer bind.ContractFilterer) (*TTPNftFilterer, error) {
	contract, err := bindTTPNft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TTPNftFilterer{contract: contract}, nil
}

// bindTTPNft binds a generic wrapper to an already deployed contract.
func bindTTPNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TTPNftABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TTPNft *TTPNftRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TTPNft.Contract.TTPNftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TTPNft *TTPNftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.Contract.TTPNftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TTPNft *TTPNftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.Contract.TTPNftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TTPNft *TTPNftCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TTPNft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TTPNft *TTPNftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TTPNft *TTPNftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_TTPNft *TTPNftCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TTPNft.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_TTPNft *TTPNftSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TTPNft.Contract.BalanceOf(&_TTPNft.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_TTPNft *TTPNftCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TTPNft.Contract.BalanceOf(&_TTPNft.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_TTPNft *TTPNftCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TTPNft.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_TTPNft *TTPNftSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _TTPNft.Contract.GetApproved(&_TTPNft.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address)
func (_TTPNft *TTPNftCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _TTPNft.Contract.GetApproved(&_TTPNft.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_TTPNft *TTPNftCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TTPNft.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_TTPNft *TTPNftSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _TTPNft.Contract.IsApprovedForAll(&_TTPNft.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_TTPNft *TTPNftCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _TTPNft.Contract.IsApprovedForAll(&_TTPNft.CallOpts, owner, operator)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_TTPNft *TTPNftCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TTPNft.contract.Call(opts, out, "isMinter", account)
	return *ret0, err
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_TTPNft *TTPNftSession) IsMinter(account common.Address) (bool, error) {
	return _TTPNft.Contract.IsMinter(&_TTPNft.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_TTPNft *TTPNftCallerSession) IsMinter(account common.Address) (bool, error) {
	return _TTPNft.Contract.IsMinter(&_TTPNft.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TTPNft *TTPNftCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TTPNft.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TTPNft *TTPNftSession) Name() (string, error) {
	return _TTPNft.Contract.Name(&_TTPNft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TTPNft *TTPNftCallerSession) Name() (string, error) {
	return _TTPNft.Contract.Name(&_TTPNft.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_TTPNft *TTPNftCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TTPNft.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_TTPNft *TTPNftSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _TTPNft.Contract.OwnerOf(&_TTPNft.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address)
func (_TTPNft *TTPNftCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _TTPNft.Contract.OwnerOf(&_TTPNft.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_TTPNft *TTPNftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TTPNft.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_TTPNft *TTPNftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TTPNft.Contract.SupportsInterface(&_TTPNft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_TTPNft *TTPNftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TTPNft.Contract.SupportsInterface(&_TTPNft.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TTPNft *TTPNftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TTPNft.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TTPNft *TTPNftSession) Symbol() (string, error) {
	return _TTPNft.Contract.Symbol(&_TTPNft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TTPNft *TTPNftCallerSession) Symbol() (string, error) {
	return _TTPNft.Contract.Symbol(&_TTPNft.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) constant returns(uint256)
func (_TTPNft *TTPNftCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TTPNft.contract.Call(opts, out, "tokenByIndex", index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) constant returns(uint256)
func (_TTPNft *TTPNftSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _TTPNft.Contract.TokenByIndex(&_TTPNft.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) constant returns(uint256)
func (_TTPNft *TTPNftCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _TTPNft.Contract.TokenByIndex(&_TTPNft.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) constant returns(uint256)
func (_TTPNft *TTPNftCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TTPNft.contract.Call(opts, out, "tokenOfOwnerByIndex", owner, index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) constant returns(uint256)
func (_TTPNft *TTPNftSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _TTPNft.Contract.TokenOfOwnerByIndex(&_TTPNft.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) constant returns(uint256)
func (_TTPNft *TTPNftCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _TTPNft.Contract.TokenOfOwnerByIndex(&_TTPNft.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) constant returns(string)
func (_TTPNft *TTPNftCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TTPNft.contract.Call(opts, out, "tokenURI", tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) constant returns(string)
func (_TTPNft *TTPNftSession) TokenURI(tokenId *big.Int) (string, error) {
	return _TTPNft.Contract.TokenURI(&_TTPNft.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) constant returns(string)
func (_TTPNft *TTPNftCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _TTPNft.Contract.TokenURI(&_TTPNft.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TTPNft *TTPNftCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TTPNft.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TTPNft *TTPNftSession) TotalSupply() (*big.Int, error) {
	return _TTPNft.Contract.TotalSupply(&_TTPNft.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TTPNft *TTPNftCallerSession) TotalSupply() (*big.Int, error) {
	return _TTPNft.Contract.TotalSupply(&_TTPNft.CallOpts)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_TTPNft *TTPNftTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.contract.Transact(opts, "addMinter", account)
}

func (_TTPNft *TTPNftTransactor) AsyncAddMinter(handler func(*types.Receipt, error), opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _TTPNft.contract.AsyncTransact(opts, handler, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_TTPNft *TTPNftSession) AddMinter(account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.Contract.AddMinter(&_TTPNft.TransactOpts, account)
}

func (_TTPNft *TTPNftSession) AsyncAddMinter(handler func(*types.Receipt, error), account common.Address) (*types.Transaction, error) {
	return _TTPNft.Contract.AsyncAddMinter(handler, &_TTPNft.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_TTPNft *TTPNftTransactorSession) AddMinter(account common.Address) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.Contract.AddMinter(&_TTPNft.TransactOpts, account)
}

func (_TTPNft *TTPNftTransactorSession) AsyncAddMinter(handler func(*types.Receipt, error), account common.Address) (*types.Transaction, error) {
	return _TTPNft.Contract.AsyncAddMinter(handler, &_TTPNft.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TTPNft *TTPNftTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.contract.Transact(opts, "approve", to, tokenId)
}

func (_TTPNft *TTPNftTransactor) AsyncApprove(handler func(*types.Receipt, error), opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TTPNft.contract.AsyncTransact(opts, handler, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TTPNft *TTPNftSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.Contract.Approve(&_TTPNft.TransactOpts, to, tokenId)
}

func (_TTPNft *TTPNftSession) AsyncApprove(handler func(*types.Receipt, error), to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TTPNft.Contract.AsyncApprove(handler, &_TTPNft.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TTPNft *TTPNftTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.Contract.Approve(&_TTPNft.TransactOpts, to, tokenId)
}

func (_TTPNft *TTPNftTransactorSession) AsyncApprove(handler func(*types.Receipt, error), to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TTPNft.Contract.AsyncApprove(handler, &_TTPNft.TransactOpts, to, tokenId)
}

// MintWithTokenURI is a paid mutator transaction binding the contract method 0x50bb4e7f.
//
// Solidity: function mintWithTokenURI(address to, uint256 tokenId, string tokenURI) returns(bool)
func (_TTPNft *TTPNftTransactor) MintWithTokenURI(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.contract.Transact(opts, "mintWithTokenURI", to, tokenId, tokenURI)
}

func (_TTPNft *TTPNftTransactor) AsyncMintWithTokenURI(handler func(*types.Receipt, error), opts *bind.TransactOpts, to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _TTPNft.contract.AsyncTransact(opts, handler, "mintWithTokenURI", to, tokenId, tokenURI)
}

// MintWithTokenURI is a paid mutator transaction binding the contract method 0x50bb4e7f.
//
// Solidity: function mintWithTokenURI(address to, uint256 tokenId, string tokenURI) returns(bool)
func (_TTPNft *TTPNftSession) MintWithTokenURI(to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.Contract.MintWithTokenURI(&_TTPNft.TransactOpts, to, tokenId, tokenURI)
}

func (_TTPNft *TTPNftSession) AsyncMintWithTokenURI(handler func(*types.Receipt, error), to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _TTPNft.Contract.AsyncMintWithTokenURI(handler, &_TTPNft.TransactOpts, to, tokenId, tokenURI)
}

// MintWithTokenURI is a paid mutator transaction binding the contract method 0x50bb4e7f.
//
// Solidity: function mintWithTokenURI(address to, uint256 tokenId, string tokenURI) returns(bool)
func (_TTPNft *TTPNftTransactorSession) MintWithTokenURI(to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.Contract.MintWithTokenURI(&_TTPNft.TransactOpts, to, tokenId, tokenURI)
}

func (_TTPNft *TTPNftTransactorSession) AsyncMintWithTokenURI(handler func(*types.Receipt, error), to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _TTPNft.Contract.AsyncMintWithTokenURI(handler, &_TTPNft.TransactOpts, to, tokenId, tokenURI)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_TTPNft *TTPNftTransactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.contract.Transact(opts, "renounceMinter")
}

func (_TTPNft *TTPNftTransactor) AsyncRenounceMinter(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TTPNft.contract.AsyncTransact(opts, handler, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_TTPNft *TTPNftSession) RenounceMinter() (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.Contract.RenounceMinter(&_TTPNft.TransactOpts)
}

func (_TTPNft *TTPNftSession) AsyncRenounceMinter(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _TTPNft.Contract.AsyncRenounceMinter(handler, &_TTPNft.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_TTPNft *TTPNftTransactorSession) RenounceMinter() (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.Contract.RenounceMinter(&_TTPNft.TransactOpts)
}

func (_TTPNft *TTPNftTransactorSession) AsyncRenounceMinter(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _TTPNft.Contract.AsyncRenounceMinter(handler, &_TTPNft.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TTPNft *TTPNftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

func (_TTPNft *TTPNftTransactor) AsyncSafeTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TTPNft.contract.AsyncTransact(opts, handler, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TTPNft *TTPNftSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.Contract.SafeTransferFrom(&_TTPNft.TransactOpts, from, to, tokenId)
}

func (_TTPNft *TTPNftSession) AsyncSafeTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TTPNft.Contract.AsyncSafeTransferFrom(handler, &_TTPNft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TTPNft *TTPNftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.Contract.SafeTransferFrom(&_TTPNft.TransactOpts, from, to, tokenId)
}

func (_TTPNft *TTPNftTransactorSession) AsyncSafeTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TTPNft.Contract.AsyncSafeTransferFrom(handler, &_TTPNft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_TTPNft *TTPNftTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

func (_TTPNft *TTPNftTransactor) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _TTPNft.contract.AsyncTransact(opts, handler, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_TTPNft *TTPNftSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.Contract.SafeTransferFrom0(&_TTPNft.TransactOpts, from, to, tokenId, _data)
}

func (_TTPNft *TTPNftSession) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _TTPNft.Contract.AsyncSafeTransferFrom0(handler, &_TTPNft.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_TTPNft *TTPNftTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.Contract.SafeTransferFrom0(&_TTPNft.TransactOpts, from, to, tokenId, _data)
}

func (_TTPNft *TTPNftTransactorSession) AsyncSafeTransferFrom0(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _TTPNft.Contract.AsyncSafeTransferFrom0(handler, &_TTPNft.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_TTPNft *TTPNftTransactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.contract.Transact(opts, "setApprovalForAll", to, approved)
}

func (_TTPNft *TTPNftTransactor) AsyncSetApprovalForAll(handler func(*types.Receipt, error), opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _TTPNft.contract.AsyncTransact(opts, handler, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_TTPNft *TTPNftSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.Contract.SetApprovalForAll(&_TTPNft.TransactOpts, to, approved)
}

func (_TTPNft *TTPNftSession) AsyncSetApprovalForAll(handler func(*types.Receipt, error), to common.Address, approved bool) (*types.Transaction, error) {
	return _TTPNft.Contract.AsyncSetApprovalForAll(handler, &_TTPNft.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_TTPNft *TTPNftTransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.Contract.SetApprovalForAll(&_TTPNft.TransactOpts, to, approved)
}

func (_TTPNft *TTPNftTransactorSession) AsyncSetApprovalForAll(handler func(*types.Receipt, error), to common.Address, approved bool) (*types.Transaction, error) {
	return _TTPNft.Contract.AsyncSetApprovalForAll(handler, &_TTPNft.TransactOpts, to, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TTPNft *TTPNftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

func (_TTPNft *TTPNftTransactor) AsyncTransferFrom(handler func(*types.Receipt, error), opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TTPNft.contract.AsyncTransact(opts, handler, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TTPNft *TTPNftSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.Contract.TransferFrom(&_TTPNft.TransactOpts, from, to, tokenId)
}

func (_TTPNft *TTPNftSession) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TTPNft.Contract.AsyncTransferFrom(handler, &_TTPNft.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TTPNft *TTPNftTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _TTPNft.Contract.TransferFrom(&_TTPNft.TransactOpts, from, to, tokenId)
}

func (_TTPNft *TTPNftTransactorSession) AsyncTransferFrom(handler func(*types.Receipt, error), from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TTPNft.Contract.AsyncTransferFrom(handler, &_TTPNft.TransactOpts, from, to, tokenId)
}

// TTPNftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TTPNft contract.
type TTPNftApprovalIterator struct {
	Event *TTPNftApproval // Event containing the contract specifics and raw log

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
func (it *TTPNftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TTPNftApproval)
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
		it.Event = new(TTPNftApproval)
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
func (it *TTPNftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TTPNftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TTPNftApproval represents a Approval event raised by the TTPNft contract.
type TTPNftApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TTPNft *TTPNftFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*TTPNftApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TTPNft.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TTPNftApprovalIterator{contract: _TTPNft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TTPNft *TTPNftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TTPNftApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TTPNft.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TTPNftApproval)
				if err := _TTPNft.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TTPNft *TTPNftFilterer) ParseApproval(log types.Log) (*TTPNftApproval, error) {
	event := new(TTPNftApproval)
	if err := _TTPNft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TTPNftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the TTPNft contract.
type TTPNftApprovalForAllIterator struct {
	Event *TTPNftApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TTPNftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TTPNftApprovalForAll)
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
		it.Event = new(TTPNftApprovalForAll)
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
func (it *TTPNftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TTPNftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TTPNftApprovalForAll represents a ApprovalForAll event raised by the TTPNft contract.
type TTPNftApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TTPNft *TTPNftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*TTPNftApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TTPNft.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TTPNftApprovalForAllIterator{contract: _TTPNft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TTPNft *TTPNftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TTPNftApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TTPNft.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TTPNftApprovalForAll)
				if err := _TTPNft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TTPNft *TTPNftFilterer) ParseApprovalForAll(log types.Log) (*TTPNftApprovalForAll, error) {
	event := new(TTPNftApprovalForAll)
	if err := _TTPNft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TTPNftMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the TTPNft contract.
type TTPNftMinterAddedIterator struct {
	Event *TTPNftMinterAdded // Event containing the contract specifics and raw log

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
func (it *TTPNftMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TTPNftMinterAdded)
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
		it.Event = new(TTPNftMinterAdded)
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
func (it *TTPNftMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TTPNftMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TTPNftMinterAdded represents a MinterAdded event raised by the TTPNft contract.
type TTPNftMinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_TTPNft *TTPNftFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*TTPNftMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TTPNft.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &TTPNftMinterAddedIterator{contract: _TTPNft.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_TTPNft *TTPNftFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *TTPNftMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TTPNft.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TTPNftMinterAdded)
				if err := _TTPNft.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// ParseMinterAdded is a log parse operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_TTPNft *TTPNftFilterer) ParseMinterAdded(log types.Log) (*TTPNftMinterAdded, error) {
	event := new(TTPNftMinterAdded)
	if err := _TTPNft.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TTPNftMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the TTPNft contract.
type TTPNftMinterRemovedIterator struct {
	Event *TTPNftMinterRemoved // Event containing the contract specifics and raw log

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
func (it *TTPNftMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TTPNftMinterRemoved)
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
		it.Event = new(TTPNftMinterRemoved)
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
func (it *TTPNftMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TTPNftMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TTPNftMinterRemoved represents a MinterRemoved event raised by the TTPNft contract.
type TTPNftMinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_TTPNft *TTPNftFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*TTPNftMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TTPNft.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &TTPNftMinterRemovedIterator{contract: _TTPNft.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_TTPNft *TTPNftFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *TTPNftMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TTPNft.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TTPNftMinterRemoved)
				if err := _TTPNft.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_TTPNft *TTPNftFilterer) ParseMinterRemoved(log types.Log) (*TTPNftMinterRemoved, error) {
	event := new(TTPNftMinterRemoved)
	if err := _TTPNft.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TTPNftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TTPNft contract.
type TTPNftTransferIterator struct {
	Event *TTPNftTransfer // Event containing the contract specifics and raw log

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
func (it *TTPNftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TTPNftTransfer)
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
		it.Event = new(TTPNftTransfer)
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
func (it *TTPNftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TTPNftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TTPNftTransfer represents a Transfer event raised by the TTPNft contract.
type TTPNftTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TTPNft *TTPNftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*TTPNftTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TTPNft.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TTPNftTransferIterator{contract: _TTPNft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TTPNft *TTPNftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TTPNftTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TTPNft.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TTPNftTransfer)
				if err := _TTPNft.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TTPNft *TTPNftFilterer) ParseTransfer(log types.Log) (*TTPNftTransfer, error) {
	event := new(TTPNftTransfer)
	if err := _TTPNft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
