// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package interact

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

// LiquidityManagerAddLiquidityParam is an auto generated low-level Go binding around an user-defined struct.
type LiquidityManagerAddLiquidityParam struct {
	Lid        *big.Int
	XLim       *big.Int
	YLim       *big.Int
	AmountXMin *big.Int
	AmountYMin *big.Int
	Deadline   *big.Int
}

// LiquidityManagerMintParam is an auto generated low-level Go binding around an user-defined struct.
type LiquidityManagerMintParam struct {
	Miner      common.Address
	TokenX     common.Address
	TokenY     common.Address
	Fee        *big.Int
	Pl         *big.Int
	Pr         *big.Int
	XLim       *big.Int
	YLim       *big.Int
	AmountXMin *big.Int
	AmountYMin *big.Int
	Deadline   *big.Int
}

// LiquiditymanagerMetaData contains all meta data concerning the Liquiditymanager contract.
var LiquiditymanagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"type\":\"address\",\"name\":\"factory\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"weth\",\"internalType\":\"address\"}]},{\"type\":\"event\",\"name\":\"AddLiquidity\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"nftId\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"address\",\"name\":\"pool\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"uint128\",\"name\":\"liquidityDelta\",\"internalType\":\"uint128\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amountX\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amountY\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"approved\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"operator\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"bool\",\"name\":\"approved\",\"internalType\":\"bool\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecLiquidity\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"nftId\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"address\",\"name\":\"pool\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"uint128\",\"name\":\"liquidityDelta\",\"internalType\":\"uint128\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amountX\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amountY\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"WETH9\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"uint128\",\"name\":\"liquidityDelta\",\"internalType\":\"uint128\"},{\"type\":\"uint256\",\"name\":\"amountX\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amountY\",\"internalType\":\"uint256\"}],\"name\":\"addLiquidity\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"addLiquidityParam\",\"internalType\":\"structLiquidityManager.AddLiquidityParam\",\"components\":[{\"type\":\"uint256\",\"name\":\"lid\",\"internalType\":\"uint256\"},{\"type\":\"uint128\",\"name\":\"xLim\",\"internalType\":\"uint128\"},{\"type\":\"uint128\",\"name\":\"yLim\",\"internalType\":\"uint128\"},{\"type\":\"uint128\",\"name\":\"amountXMin\",\"internalType\":\"uint128\"},{\"type\":\"uint128\",\"name\":\"amountYMin\",\"internalType\":\"uint128\"},{\"type\":\"uint256\",\"name\":\"deadline\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"approve\",\"inputs\":[{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"balanceOf\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"baseURI\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"success\",\"internalType\":\"bool\"}],\"name\":\"burn\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"lid\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"amountX\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amountY\",\"internalType\":\"uint256\"}],\"name\":\"collect\",\"inputs\":[{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"lid\",\"internalType\":\"uint256\"},{\"type\":\"uint128\",\"name\":\"amountXLim\",\"internalType\":\"uint128\"},{\"type\":\"uint128\",\"name\":\"amountYLim\",\"internalType\":\"uint128\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"createPool\",\"inputs\":[{\"type\":\"address\",\"name\":\"tokenX\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"tokenY\",\"internalType\":\"address\"},{\"type\":\"uint24\",\"name\":\"fee\",\"internalType\":\"uint24\"},{\"type\":\"int24\",\"name\":\"initialPoint\",\"internalType\":\"int24\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"amountX\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amountY\",\"internalType\":\"uint256\"}],\"name\":\"decLiquidity\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"lid\",\"internalType\":\"uint256\"},{\"type\":\"uint128\",\"name\":\"liquidDelta\",\"internalType\":\"uint128\"},{\"type\":\"uint256\",\"name\":\"amountXMin\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amountYMin\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"deadline\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"factory\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"getApproved\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isApprovedForAll\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"operator\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"int24\",\"name\":\"leftPt\",\"internalType\":\"int24\"},{\"type\":\"int24\",\"name\":\"rightPt\",\"internalType\":\"int24\"},{\"type\":\"uint128\",\"name\":\"liquidity\",\"internalType\":\"uint128\"},{\"type\":\"uint256\",\"name\":\"lastFeeScaleX_128\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"lastFeeScaleY_128\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"remainTokenX\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"remainTokenY\",\"internalType\":\"uint256\"},{\"type\":\"uint128\",\"name\":\"poolId\",\"internalType\":\"uint128\"}],\"name\":\"liquidities\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"liquidityNum\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"lid\",\"internalType\":\"uint256\"},{\"type\":\"uint128\",\"name\":\"liquidity\",\"internalType\":\"uint128\"},{\"type\":\"uint256\",\"name\":\"amountX\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amountY\",\"internalType\":\"uint256\"}],\"name\":\"mint\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"mintParam\",\"internalType\":\"structLiquidityManager.MintParam\",\"components\":[{\"type\":\"address\",\"name\":\"miner\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"tokenX\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"tokenY\",\"internalType\":\"address\"},{\"type\":\"uint24\",\"name\":\"fee\",\"internalType\":\"uint24\"},{\"type\":\"int24\",\"name\":\"pl\",\"internalType\":\"int24\"},{\"type\":\"int24\",\"name\":\"pr\",\"internalType\":\"int24\"},{\"type\":\"uint128\",\"name\":\"xLim\",\"internalType\":\"uint128\"},{\"type\":\"uint128\",\"name\":\"yLim\",\"internalType\":\"uint128\"},{\"type\":\"uint128\",\"name\":\"amountXMin\",\"internalType\":\"uint128\"},{\"type\":\"uint128\",\"name\":\"amountYMin\",\"internalType\":\"uint128\"},{\"type\":\"uint256\",\"name\":\"deadline\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"mintDepositCallback\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"x\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"y\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[{\"type\":\"bytes[]\",\"name\":\"results\",\"internalType\":\"bytes[]\"}],\"name\":\"multicall\",\"inputs\":[{\"type\":\"bytes[]\",\"name\":\"data\",\"internalType\":\"bytes[]\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"name\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"owner\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"ownerOf\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"pool\",\"inputs\":[{\"type\":\"address\",\"name\":\"tokenX\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"tokenY\",\"internalType\":\"address\"},{\"type\":\"uint24\",\"name\":\"fee\",\"internalType\":\"uint24\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint128\",\"name\":\"\",\"internalType\":\"uint128\"}],\"name\":\"poolIds\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"tokenX\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"tokenY\",\"internalType\":\"address\"},{\"type\":\"uint24\",\"name\":\"fee\",\"internalType\":\"uint24\"}],\"name\":\"poolMetas\",\"inputs\":[{\"type\":\"uint128\",\"name\":\"\",\"internalType\":\"uint128\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"refundETH\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"renounceOwnership\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"safeTransferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"safeTransferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setApprovalForAll\",\"inputs\":[{\"type\":\"address\",\"name\":\"operator\",\"internalType\":\"address\"},{\"type\":\"bool\",\"name\":\"approved\",\"internalType\":\"bool\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setBaseURI\",\"inputs\":[{\"type\":\"string\",\"name\":\"newBaseURI\",\"internalType\":\"string\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"supportsInterface\",\"inputs\":[{\"type\":\"bytes4\",\"name\":\"interfaceId\",\"internalType\":\"bytes4\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"sweepToken\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"minAmount\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"symbol\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"tokenByIndex\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"index\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"index\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"tokenURI\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"totalSupply\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"transferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"transferOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"unwrapWETH9\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"minAmount\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"recipient\",\"internalType\":\"address\"}]},{\"type\":\"receive\",\"stateMutability\":\"payable\"}]",
}

// LiquiditymanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquiditymanagerMetaData.ABI instead.
var LiquiditymanagerABI = LiquiditymanagerMetaData.ABI

// Liquiditymanager is an auto generated Go binding around an Ethereum contract.
type Liquiditymanager struct {
	LiquiditymanagerCaller     // Read-only binding to the contract
	LiquiditymanagerTransactor // Write-only binding to the contract
	LiquiditymanagerFilterer   // Log filterer for contract events
}

// LiquiditymanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquiditymanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquiditymanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquiditymanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquiditymanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquiditymanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquiditymanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquiditymanagerSession struct {
	Contract     *Liquiditymanager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LiquiditymanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquiditymanagerCallerSession struct {
	Contract *LiquiditymanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// LiquiditymanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquiditymanagerTransactorSession struct {
	Contract     *LiquiditymanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// LiquiditymanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquiditymanagerRaw struct {
	Contract *Liquiditymanager // Generic contract binding to access the raw methods on
}

// LiquiditymanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquiditymanagerCallerRaw struct {
	Contract *LiquiditymanagerCaller // Generic read-only contract binding to access the raw methods on
}

// LiquiditymanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquiditymanagerTransactorRaw struct {
	Contract *LiquiditymanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquiditymanager creates a new instance of Liquiditymanager, bound to a specific deployed contract.
func NewLiquiditymanager(address common.Address, backend bind.ContractBackend) (*Liquiditymanager, error) {
	contract, err := bindLiquiditymanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Liquiditymanager{LiquiditymanagerCaller: LiquiditymanagerCaller{contract: contract}, LiquiditymanagerTransactor: LiquiditymanagerTransactor{contract: contract}, LiquiditymanagerFilterer: LiquiditymanagerFilterer{contract: contract}}, nil
}

// NewLiquiditymanagerCaller creates a new read-only instance of Liquiditymanager, bound to a specific deployed contract.
func NewLiquiditymanagerCaller(address common.Address, caller bind.ContractCaller) (*LiquiditymanagerCaller, error) {
	contract, err := bindLiquiditymanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquiditymanagerCaller{contract: contract}, nil
}

// NewLiquiditymanagerTransactor creates a new write-only instance of Liquiditymanager, bound to a specific deployed contract.
func NewLiquiditymanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquiditymanagerTransactor, error) {
	contract, err := bindLiquiditymanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquiditymanagerTransactor{contract: contract}, nil
}

// NewLiquiditymanagerFilterer creates a new log filterer instance of Liquiditymanager, bound to a specific deployed contract.
func NewLiquiditymanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquiditymanagerFilterer, error) {
	contract, err := bindLiquiditymanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquiditymanagerFilterer{contract: contract}, nil
}

// bindLiquiditymanager binds a generic wrapper to an already deployed contract.
func bindLiquiditymanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LiquiditymanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Liquiditymanager *LiquiditymanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Liquiditymanager.Contract.LiquiditymanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Liquiditymanager *LiquiditymanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.LiquiditymanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Liquiditymanager *LiquiditymanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.LiquiditymanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Liquiditymanager *LiquiditymanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Liquiditymanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Liquiditymanager *LiquiditymanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Liquiditymanager *LiquiditymanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Liquiditymanager *LiquiditymanagerCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liquiditymanager.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Liquiditymanager *LiquiditymanagerSession) WETH9() (common.Address, error) {
	return _Liquiditymanager.Contract.WETH9(&_Liquiditymanager.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_Liquiditymanager *LiquiditymanagerCallerSession) WETH9() (common.Address, error) {
	return _Liquiditymanager.Contract.WETH9(&_Liquiditymanager.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Liquiditymanager *LiquiditymanagerCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Liquiditymanager.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Liquiditymanager *LiquiditymanagerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Liquiditymanager.Contract.BalanceOf(&_Liquiditymanager.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Liquiditymanager *LiquiditymanagerCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Liquiditymanager.Contract.BalanceOf(&_Liquiditymanager.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Liquiditymanager *LiquiditymanagerCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Liquiditymanager.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Liquiditymanager *LiquiditymanagerSession) BaseURI() (string, error) {
	return _Liquiditymanager.Contract.BaseURI(&_Liquiditymanager.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Liquiditymanager *LiquiditymanagerCallerSession) BaseURI() (string, error) {
	return _Liquiditymanager.Contract.BaseURI(&_Liquiditymanager.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Liquiditymanager *LiquiditymanagerCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liquiditymanager.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Liquiditymanager *LiquiditymanagerSession) Factory() (common.Address, error) {
	return _Liquiditymanager.Contract.Factory(&_Liquiditymanager.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Liquiditymanager *LiquiditymanagerCallerSession) Factory() (common.Address, error) {
	return _Liquiditymanager.Contract.Factory(&_Liquiditymanager.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Liquiditymanager *LiquiditymanagerCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Liquiditymanager.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Liquiditymanager *LiquiditymanagerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Liquiditymanager.Contract.GetApproved(&_Liquiditymanager.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Liquiditymanager *LiquiditymanagerCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Liquiditymanager.Contract.GetApproved(&_Liquiditymanager.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Liquiditymanager *LiquiditymanagerCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Liquiditymanager.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Liquiditymanager *LiquiditymanagerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Liquiditymanager.Contract.IsApprovedForAll(&_Liquiditymanager.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Liquiditymanager *LiquiditymanagerCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Liquiditymanager.Contract.IsApprovedForAll(&_Liquiditymanager.CallOpts, owner, operator)
}

// Liquidities is a free data retrieval call binding the contract method 0x0713051d.
//
// Solidity: function liquidities(uint256 ) view returns(int24 leftPt, int24 rightPt, uint128 liquidity, uint256 lastFeeScaleX_128, uint256 lastFeeScaleY_128, uint256 remainTokenX, uint256 remainTokenY, uint128 poolId)
func (_Liquiditymanager *LiquiditymanagerCaller) Liquidities(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LeftPt           *big.Int
	RightPt          *big.Int
	Liquidity        *big.Int
	LastFeeScaleX128 *big.Int
	LastFeeScaleY128 *big.Int
	RemainTokenX     *big.Int
	RemainTokenY     *big.Int
	PoolId           *big.Int
}, error) {
	var out []interface{}
	err := _Liquiditymanager.contract.Call(opts, &out, "liquidities", arg0)

	outstruct := new(struct {
		LeftPt           *big.Int
		RightPt          *big.Int
		Liquidity        *big.Int
		LastFeeScaleX128 *big.Int
		LastFeeScaleY128 *big.Int
		RemainTokenX     *big.Int
		RemainTokenY     *big.Int
		PoolId           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LeftPt = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RightPt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Liquidity = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LastFeeScaleX128 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastFeeScaleY128 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RemainTokenX = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.RemainTokenY = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.PoolId = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Liquidities is a free data retrieval call binding the contract method 0x0713051d.
//
// Solidity: function liquidities(uint256 ) view returns(int24 leftPt, int24 rightPt, uint128 liquidity, uint256 lastFeeScaleX_128, uint256 lastFeeScaleY_128, uint256 remainTokenX, uint256 remainTokenY, uint128 poolId)
func (_Liquiditymanager *LiquiditymanagerSession) Liquidities(arg0 *big.Int) (struct {
	LeftPt           *big.Int
	RightPt          *big.Int
	Liquidity        *big.Int
	LastFeeScaleX128 *big.Int
	LastFeeScaleY128 *big.Int
	RemainTokenX     *big.Int
	RemainTokenY     *big.Int
	PoolId           *big.Int
}, error) {
	return _Liquiditymanager.Contract.Liquidities(&_Liquiditymanager.CallOpts, arg0)
}

// Liquidities is a free data retrieval call binding the contract method 0x0713051d.
//
// Solidity: function liquidities(uint256 ) view returns(int24 leftPt, int24 rightPt, uint128 liquidity, uint256 lastFeeScaleX_128, uint256 lastFeeScaleY_128, uint256 remainTokenX, uint256 remainTokenY, uint128 poolId)
func (_Liquiditymanager *LiquiditymanagerCallerSession) Liquidities(arg0 *big.Int) (struct {
	LeftPt           *big.Int
	RightPt          *big.Int
	Liquidity        *big.Int
	LastFeeScaleX128 *big.Int
	LastFeeScaleY128 *big.Int
	RemainTokenX     *big.Int
	RemainTokenY     *big.Int
	PoolId           *big.Int
}, error) {
	return _Liquiditymanager.Contract.Liquidities(&_Liquiditymanager.CallOpts, arg0)
}

// LiquidityNum is a free data retrieval call binding the contract method 0xdca87bec.
//
// Solidity: function liquidityNum() view returns(uint256)
func (_Liquiditymanager *LiquiditymanagerCaller) LiquidityNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Liquiditymanager.contract.Call(opts, &out, "liquidityNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidityNum is a free data retrieval call binding the contract method 0xdca87bec.
//
// Solidity: function liquidityNum() view returns(uint256)
func (_Liquiditymanager *LiquiditymanagerSession) LiquidityNum() (*big.Int, error) {
	return _Liquiditymanager.Contract.LiquidityNum(&_Liquiditymanager.CallOpts)
}

// LiquidityNum is a free data retrieval call binding the contract method 0xdca87bec.
//
// Solidity: function liquidityNum() view returns(uint256)
func (_Liquiditymanager *LiquiditymanagerCallerSession) LiquidityNum() (*big.Int, error) {
	return _Liquiditymanager.Contract.LiquidityNum(&_Liquiditymanager.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Liquiditymanager *LiquiditymanagerCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Liquiditymanager.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Liquiditymanager *LiquiditymanagerSession) Name() (string, error) {
	return _Liquiditymanager.Contract.Name(&_Liquiditymanager.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Liquiditymanager *LiquiditymanagerCallerSession) Name() (string, error) {
	return _Liquiditymanager.Contract.Name(&_Liquiditymanager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Liquiditymanager *LiquiditymanagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liquiditymanager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Liquiditymanager *LiquiditymanagerSession) Owner() (common.Address, error) {
	return _Liquiditymanager.Contract.Owner(&_Liquiditymanager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Liquiditymanager *LiquiditymanagerCallerSession) Owner() (common.Address, error) {
	return _Liquiditymanager.Contract.Owner(&_Liquiditymanager.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Liquiditymanager *LiquiditymanagerCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Liquiditymanager.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Liquiditymanager *LiquiditymanagerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Liquiditymanager.Contract.OwnerOf(&_Liquiditymanager.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Liquiditymanager *LiquiditymanagerCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Liquiditymanager.Contract.OwnerOf(&_Liquiditymanager.CallOpts, tokenId)
}

// Pool is a free data retrieval call binding the contract method 0xbecbcc6a.
//
// Solidity: function pool(address tokenX, address tokenY, uint24 fee) view returns(address)
func (_Liquiditymanager *LiquiditymanagerCaller) Pool(opts *bind.CallOpts, tokenX common.Address, tokenY common.Address, fee *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Liquiditymanager.contract.Call(opts, &out, "pool", tokenX, tokenY, fee)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0xbecbcc6a.
//
// Solidity: function pool(address tokenX, address tokenY, uint24 fee) view returns(address)
func (_Liquiditymanager *LiquiditymanagerSession) Pool(tokenX common.Address, tokenY common.Address, fee *big.Int) (common.Address, error) {
	return _Liquiditymanager.Contract.Pool(&_Liquiditymanager.CallOpts, tokenX, tokenY, fee)
}

// Pool is a free data retrieval call binding the contract method 0xbecbcc6a.
//
// Solidity: function pool(address tokenX, address tokenY, uint24 fee) view returns(address)
func (_Liquiditymanager *LiquiditymanagerCallerSession) Pool(tokenX common.Address, tokenY common.Address, fee *big.Int) (common.Address, error) {
	return _Liquiditymanager.Contract.Pool(&_Liquiditymanager.CallOpts, tokenX, tokenY, fee)
}

// PoolIds is a free data retrieval call binding the contract method 0xd4175be2.
//
// Solidity: function poolIds(address ) view returns(uint128)
func (_Liquiditymanager *LiquiditymanagerCaller) PoolIds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Liquiditymanager.contract.Call(opts, &out, "poolIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolIds is a free data retrieval call binding the contract method 0xd4175be2.
//
// Solidity: function poolIds(address ) view returns(uint128)
func (_Liquiditymanager *LiquiditymanagerSession) PoolIds(arg0 common.Address) (*big.Int, error) {
	return _Liquiditymanager.Contract.PoolIds(&_Liquiditymanager.CallOpts, arg0)
}

// PoolIds is a free data retrieval call binding the contract method 0xd4175be2.
//
// Solidity: function poolIds(address ) view returns(uint128)
func (_Liquiditymanager *LiquiditymanagerCallerSession) PoolIds(arg0 common.Address) (*big.Int, error) {
	return _Liquiditymanager.Contract.PoolIds(&_Liquiditymanager.CallOpts, arg0)
}

// PoolMetas is a free data retrieval call binding the contract method 0xf655dbc1.
//
// Solidity: function poolMetas(uint128 ) view returns(address tokenX, address tokenY, uint24 fee)
func (_Liquiditymanager *LiquiditymanagerCaller) PoolMetas(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenX common.Address
	TokenY common.Address
	Fee    *big.Int
}, error) {
	var out []interface{}
	err := _Liquiditymanager.contract.Call(opts, &out, "poolMetas", arg0)

	outstruct := new(struct {
		TokenX common.Address
		TokenY common.Address
		Fee    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenX = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenY = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolMetas is a free data retrieval call binding the contract method 0xf655dbc1.
//
// Solidity: function poolMetas(uint128 ) view returns(address tokenX, address tokenY, uint24 fee)
func (_Liquiditymanager *LiquiditymanagerSession) PoolMetas(arg0 *big.Int) (struct {
	TokenX common.Address
	TokenY common.Address
	Fee    *big.Int
}, error) {
	return _Liquiditymanager.Contract.PoolMetas(&_Liquiditymanager.CallOpts, arg0)
}

// PoolMetas is a free data retrieval call binding the contract method 0xf655dbc1.
//
// Solidity: function poolMetas(uint128 ) view returns(address tokenX, address tokenY, uint24 fee)
func (_Liquiditymanager *LiquiditymanagerCallerSession) PoolMetas(arg0 *big.Int) (struct {
	TokenX common.Address
	TokenY common.Address
	Fee    *big.Int
}, error) {
	return _Liquiditymanager.Contract.PoolMetas(&_Liquiditymanager.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Liquiditymanager *LiquiditymanagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Liquiditymanager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Liquiditymanager *LiquiditymanagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Liquiditymanager.Contract.SupportsInterface(&_Liquiditymanager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Liquiditymanager *LiquiditymanagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Liquiditymanager.Contract.SupportsInterface(&_Liquiditymanager.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Liquiditymanager *LiquiditymanagerCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Liquiditymanager.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Liquiditymanager *LiquiditymanagerSession) Symbol() (string, error) {
	return _Liquiditymanager.Contract.Symbol(&_Liquiditymanager.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Liquiditymanager *LiquiditymanagerCallerSession) Symbol() (string, error) {
	return _Liquiditymanager.Contract.Symbol(&_Liquiditymanager.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Liquiditymanager *LiquiditymanagerCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Liquiditymanager.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Liquiditymanager *LiquiditymanagerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Liquiditymanager.Contract.TokenByIndex(&_Liquiditymanager.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Liquiditymanager *LiquiditymanagerCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Liquiditymanager.Contract.TokenByIndex(&_Liquiditymanager.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Liquiditymanager *LiquiditymanagerCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Liquiditymanager.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Liquiditymanager *LiquiditymanagerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Liquiditymanager.Contract.TokenOfOwnerByIndex(&_Liquiditymanager.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Liquiditymanager *LiquiditymanagerCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Liquiditymanager.Contract.TokenOfOwnerByIndex(&_Liquiditymanager.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Liquiditymanager *LiquiditymanagerCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Liquiditymanager.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Liquiditymanager *LiquiditymanagerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Liquiditymanager.Contract.TokenURI(&_Liquiditymanager.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Liquiditymanager *LiquiditymanagerCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Liquiditymanager.Contract.TokenURI(&_Liquiditymanager.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Liquiditymanager *LiquiditymanagerCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Liquiditymanager.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Liquiditymanager *LiquiditymanagerSession) TotalSupply() (*big.Int, error) {
	return _Liquiditymanager.Contract.TotalSupply(&_Liquiditymanager.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Liquiditymanager *LiquiditymanagerCallerSession) TotalSupply() (*big.Int, error) {
	return _Liquiditymanager.Contract.TotalSupply(&_Liquiditymanager.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xcbd89416.
//
// Solidity: function addLiquidity((uint256,uint128,uint128,uint128,uint128,uint256) addLiquidityParam) payable returns(uint128 liquidityDelta, uint256 amountX, uint256 amountY)
func (_Liquiditymanager *LiquiditymanagerTransactor) AddLiquidity(opts *bind.TransactOpts, addLiquidityParam LiquidityManagerAddLiquidityParam) (*types.Transaction, error) {
	return _Liquiditymanager.contract.Transact(opts, "addLiquidity", addLiquidityParam)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xcbd89416.
//
// Solidity: function addLiquidity((uint256,uint128,uint128,uint128,uint128,uint256) addLiquidityParam) payable returns(uint128 liquidityDelta, uint256 amountX, uint256 amountY)
func (_Liquiditymanager *LiquiditymanagerSession) AddLiquidity(addLiquidityParam LiquidityManagerAddLiquidityParam) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.AddLiquidity(&_Liquiditymanager.TransactOpts, addLiquidityParam)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xcbd89416.
//
// Solidity: function addLiquidity((uint256,uint128,uint128,uint128,uint128,uint256) addLiquidityParam) payable returns(uint128 liquidityDelta, uint256 amountX, uint256 amountY)
func (_Liquiditymanager *LiquiditymanagerTransactorSession) AddLiquidity(addLiquidityParam LiquidityManagerAddLiquidityParam) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.AddLiquidity(&_Liquiditymanager.TransactOpts, addLiquidityParam)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Liquiditymanager *LiquiditymanagerTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Liquiditymanager.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Liquiditymanager *LiquiditymanagerSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.Approve(&_Liquiditymanager.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Liquiditymanager *LiquiditymanagerTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.Approve(&_Liquiditymanager.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 lid) returns(bool success)
func (_Liquiditymanager *LiquiditymanagerTransactor) Burn(opts *bind.TransactOpts, lid *big.Int) (*types.Transaction, error) {
	return _Liquiditymanager.contract.Transact(opts, "burn", lid)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 lid) returns(bool success)
func (_Liquiditymanager *LiquiditymanagerSession) Burn(lid *big.Int) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.Burn(&_Liquiditymanager.TransactOpts, lid)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 lid) returns(bool success)
func (_Liquiditymanager *LiquiditymanagerTransactorSession) Burn(lid *big.Int) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.Burn(&_Liquiditymanager.TransactOpts, lid)
}

// Collect is a paid mutator transaction binding the contract method 0xa0e4eb3c.
//
// Solidity: function collect(address recipient, uint256 lid, uint128 amountXLim, uint128 amountYLim) payable returns(uint256 amountX, uint256 amountY)
func (_Liquiditymanager *LiquiditymanagerTransactor) Collect(opts *bind.TransactOpts, recipient common.Address, lid *big.Int, amountXLim *big.Int, amountYLim *big.Int) (*types.Transaction, error) {
	return _Liquiditymanager.contract.Transact(opts, "collect", recipient, lid, amountXLim, amountYLim)
}

// Collect is a paid mutator transaction binding the contract method 0xa0e4eb3c.
//
// Solidity: function collect(address recipient, uint256 lid, uint128 amountXLim, uint128 amountYLim) payable returns(uint256 amountX, uint256 amountY)
func (_Liquiditymanager *LiquiditymanagerSession) Collect(recipient common.Address, lid *big.Int, amountXLim *big.Int, amountYLim *big.Int) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.Collect(&_Liquiditymanager.TransactOpts, recipient, lid, amountXLim, amountYLim)
}

// Collect is a paid mutator transaction binding the contract method 0xa0e4eb3c.
//
// Solidity: function collect(address recipient, uint256 lid, uint128 amountXLim, uint128 amountYLim) payable returns(uint256 amountX, uint256 amountY)
func (_Liquiditymanager *LiquiditymanagerTransactorSession) Collect(recipient common.Address, lid *big.Int, amountXLim *big.Int, amountYLim *big.Int) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.Collect(&_Liquiditymanager.TransactOpts, recipient, lid, amountXLim, amountYLim)
}

// CreatePool is a paid mutator transaction binding the contract method 0xf425a3ce.
//
// Solidity: function createPool(address tokenX, address tokenY, uint24 fee, int24 initialPoint) returns(address)
func (_Liquiditymanager *LiquiditymanagerTransactor) CreatePool(opts *bind.TransactOpts, tokenX common.Address, tokenY common.Address, fee *big.Int, initialPoint *big.Int) (*types.Transaction, error) {
	return _Liquiditymanager.contract.Transact(opts, "createPool", tokenX, tokenY, fee, initialPoint)
}

// CreatePool is a paid mutator transaction binding the contract method 0xf425a3ce.
//
// Solidity: function createPool(address tokenX, address tokenY, uint24 fee, int24 initialPoint) returns(address)
func (_Liquiditymanager *LiquiditymanagerSession) CreatePool(tokenX common.Address, tokenY common.Address, fee *big.Int, initialPoint *big.Int) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.CreatePool(&_Liquiditymanager.TransactOpts, tokenX, tokenY, fee, initialPoint)
}

// CreatePool is a paid mutator transaction binding the contract method 0xf425a3ce.
//
// Solidity: function createPool(address tokenX, address tokenY, uint24 fee, int24 initialPoint) returns(address)
func (_Liquiditymanager *LiquiditymanagerTransactorSession) CreatePool(tokenX common.Address, tokenY common.Address, fee *big.Int, initialPoint *big.Int) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.CreatePool(&_Liquiditymanager.TransactOpts, tokenX, tokenY, fee, initialPoint)
}

// DecLiquidity is a paid mutator transaction binding the contract method 0x15feae51.
//
// Solidity: function decLiquidity(uint256 lid, uint128 liquidDelta, uint256 amountXMin, uint256 amountYMin, uint256 deadline) returns(uint256 amountX, uint256 amountY)
func (_Liquiditymanager *LiquiditymanagerTransactor) DecLiquidity(opts *bind.TransactOpts, lid *big.Int, liquidDelta *big.Int, amountXMin *big.Int, amountYMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Liquiditymanager.contract.Transact(opts, "decLiquidity", lid, liquidDelta, amountXMin, amountYMin, deadline)
}

// DecLiquidity is a paid mutator transaction binding the contract method 0x15feae51.
//
// Solidity: function decLiquidity(uint256 lid, uint128 liquidDelta, uint256 amountXMin, uint256 amountYMin, uint256 deadline) returns(uint256 amountX, uint256 amountY)
func (_Liquiditymanager *LiquiditymanagerSession) DecLiquidity(lid *big.Int, liquidDelta *big.Int, amountXMin *big.Int, amountYMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.DecLiquidity(&_Liquiditymanager.TransactOpts, lid, liquidDelta, amountXMin, amountYMin, deadline)
}

// DecLiquidity is a paid mutator transaction binding the contract method 0x15feae51.
//
// Solidity: function decLiquidity(uint256 lid, uint128 liquidDelta, uint256 amountXMin, uint256 amountYMin, uint256 deadline) returns(uint256 amountX, uint256 amountY)
func (_Liquiditymanager *LiquiditymanagerTransactorSession) DecLiquidity(lid *big.Int, liquidDelta *big.Int, amountXMin *big.Int, amountYMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.DecLiquidity(&_Liquiditymanager.TransactOpts, lid, liquidDelta, amountXMin, amountYMin, deadline)
}

// Mint is a paid mutator transaction binding the contract method 0x96f639ed.
//
// Solidity: function mint((address,address,address,uint24,int24,int24,uint128,uint128,uint128,uint128,uint256) mintParam) payable returns(uint256 lid, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_Liquiditymanager *LiquiditymanagerTransactor) Mint(opts *bind.TransactOpts, mintParam LiquidityManagerMintParam) (*types.Transaction, error) {
	return _Liquiditymanager.contract.Transact(opts, "mint", mintParam)
}

// Mint is a paid mutator transaction binding the contract method 0x96f639ed.
//
// Solidity: function mint((address,address,address,uint24,int24,int24,uint128,uint128,uint128,uint128,uint256) mintParam) payable returns(uint256 lid, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_Liquiditymanager *LiquiditymanagerSession) Mint(mintParam LiquidityManagerMintParam) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.Mint(&_Liquiditymanager.TransactOpts, mintParam)
}

// Mint is a paid mutator transaction binding the contract method 0x96f639ed.
//
// Solidity: function mint((address,address,address,uint24,int24,int24,uint128,uint128,uint128,uint128,uint256) mintParam) payable returns(uint256 lid, uint128 liquidity, uint256 amountX, uint256 amountY)
func (_Liquiditymanager *LiquiditymanagerTransactorSession) Mint(mintParam LiquidityManagerMintParam) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.Mint(&_Liquiditymanager.TransactOpts, mintParam)
}

// MintDepositCallback is a paid mutator transaction binding the contract method 0x84fe2b3d.
//
// Solidity: function mintDepositCallback(uint256 x, uint256 y, bytes data) returns()
func (_Liquiditymanager *LiquiditymanagerTransactor) MintDepositCallback(opts *bind.TransactOpts, x *big.Int, y *big.Int, data []byte) (*types.Transaction, error) {
	return _Liquiditymanager.contract.Transact(opts, "mintDepositCallback", x, y, data)
}

// MintDepositCallback is a paid mutator transaction binding the contract method 0x84fe2b3d.
//
// Solidity: function mintDepositCallback(uint256 x, uint256 y, bytes data) returns()
func (_Liquiditymanager *LiquiditymanagerSession) MintDepositCallback(x *big.Int, y *big.Int, data []byte) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.MintDepositCallback(&_Liquiditymanager.TransactOpts, x, y, data)
}

// MintDepositCallback is a paid mutator transaction binding the contract method 0x84fe2b3d.
//
// Solidity: function mintDepositCallback(uint256 x, uint256 y, bytes data) returns()
func (_Liquiditymanager *LiquiditymanagerTransactorSession) MintDepositCallback(x *big.Int, y *big.Int, data []byte) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.MintDepositCallback(&_Liquiditymanager.TransactOpts, x, y, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Liquiditymanager *LiquiditymanagerTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Liquiditymanager.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Liquiditymanager *LiquiditymanagerSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.Multicall(&_Liquiditymanager.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Liquiditymanager *LiquiditymanagerTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.Multicall(&_Liquiditymanager.TransactOpts, data)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_Liquiditymanager *LiquiditymanagerTransactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquiditymanager.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_Liquiditymanager *LiquiditymanagerSession) RefundETH() (*types.Transaction, error) {
	return _Liquiditymanager.Contract.RefundETH(&_Liquiditymanager.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_Liquiditymanager *LiquiditymanagerTransactorSession) RefundETH() (*types.Transaction, error) {
	return _Liquiditymanager.Contract.RefundETH(&_Liquiditymanager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Liquiditymanager *LiquiditymanagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquiditymanager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Liquiditymanager *LiquiditymanagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Liquiditymanager.Contract.RenounceOwnership(&_Liquiditymanager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Liquiditymanager *LiquiditymanagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Liquiditymanager.Contract.RenounceOwnership(&_Liquiditymanager.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Liquiditymanager *LiquiditymanagerTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Liquiditymanager.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Liquiditymanager *LiquiditymanagerSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.SafeTransferFrom(&_Liquiditymanager.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Liquiditymanager *LiquiditymanagerTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.SafeTransferFrom(&_Liquiditymanager.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Liquiditymanager *LiquiditymanagerTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Liquiditymanager.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Liquiditymanager *LiquiditymanagerSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.SafeTransferFrom0(&_Liquiditymanager.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Liquiditymanager *LiquiditymanagerTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.SafeTransferFrom0(&_Liquiditymanager.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Liquiditymanager *LiquiditymanagerTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Liquiditymanager.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Liquiditymanager *LiquiditymanagerSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.SetApprovalForAll(&_Liquiditymanager.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Liquiditymanager *LiquiditymanagerTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.SetApprovalForAll(&_Liquiditymanager.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Liquiditymanager *LiquiditymanagerTransactor) SetBaseURI(opts *bind.TransactOpts, newBaseURI string) (*types.Transaction, error) {
	return _Liquiditymanager.contract.Transact(opts, "setBaseURI", newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Liquiditymanager *LiquiditymanagerSession) SetBaseURI(newBaseURI string) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.SetBaseURI(&_Liquiditymanager.TransactOpts, newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string newBaseURI) returns()
func (_Liquiditymanager *LiquiditymanagerTransactorSession) SetBaseURI(newBaseURI string) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.SetBaseURI(&_Liquiditymanager.TransactOpts, newBaseURI)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 minAmount, address recipient) payable returns()
func (_Liquiditymanager *LiquiditymanagerTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Liquiditymanager.contract.Transact(opts, "sweepToken", token, minAmount, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 minAmount, address recipient) payable returns()
func (_Liquiditymanager *LiquiditymanagerSession) SweepToken(token common.Address, minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.SweepToken(&_Liquiditymanager.TransactOpts, token, minAmount, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 minAmount, address recipient) payable returns()
func (_Liquiditymanager *LiquiditymanagerTransactorSession) SweepToken(token common.Address, minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.SweepToken(&_Liquiditymanager.TransactOpts, token, minAmount, recipient)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Liquiditymanager *LiquiditymanagerTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Liquiditymanager.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Liquiditymanager *LiquiditymanagerSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.TransferFrom(&_Liquiditymanager.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Liquiditymanager *LiquiditymanagerTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.TransferFrom(&_Liquiditymanager.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Liquiditymanager *LiquiditymanagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Liquiditymanager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Liquiditymanager *LiquiditymanagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.TransferOwnership(&_Liquiditymanager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Liquiditymanager *LiquiditymanagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.TransferOwnership(&_Liquiditymanager.TransactOpts, newOwner)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 minAmount, address recipient) payable returns()
func (_Liquiditymanager *LiquiditymanagerTransactor) UnwrapWETH9(opts *bind.TransactOpts, minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Liquiditymanager.contract.Transact(opts, "unwrapWETH9", minAmount, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 minAmount, address recipient) payable returns()
func (_Liquiditymanager *LiquiditymanagerSession) UnwrapWETH9(minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.UnwrapWETH9(&_Liquiditymanager.TransactOpts, minAmount, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 minAmount, address recipient) payable returns()
func (_Liquiditymanager *LiquiditymanagerTransactorSession) UnwrapWETH9(minAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Liquiditymanager.Contract.UnwrapWETH9(&_Liquiditymanager.TransactOpts, minAmount, recipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Liquiditymanager *LiquiditymanagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquiditymanager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Liquiditymanager *LiquiditymanagerSession) Receive() (*types.Transaction, error) {
	return _Liquiditymanager.Contract.Receive(&_Liquiditymanager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Liquiditymanager *LiquiditymanagerTransactorSession) Receive() (*types.Transaction, error) {
	return _Liquiditymanager.Contract.Receive(&_Liquiditymanager.TransactOpts)
}

// LiquiditymanagerAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Liquiditymanager contract.
type LiquiditymanagerAddLiquidityIterator struct {
	Event *LiquiditymanagerAddLiquidity // Event containing the contract specifics and raw log

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
func (it *LiquiditymanagerAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquiditymanagerAddLiquidity)
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
		it.Event = new(LiquiditymanagerAddLiquidity)
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
func (it *LiquiditymanagerAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquiditymanagerAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquiditymanagerAddLiquidity represents a AddLiquidity event raised by the Liquiditymanager contract.
type LiquiditymanagerAddLiquidity struct {
	NftId          *big.Int
	Pool           common.Address
	LiquidityDelta *big.Int
	AmountX        *big.Int
	AmountY        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0xf565fdd70b3936f0ae8efc41c2e0822f9de5ecb4dc162b153b129ec4bb9cd93c.
//
// Solidity: event AddLiquidity(uint256 indexed nftId, address pool, uint128 liquidityDelta, uint256 amountX, uint256 amountY)
func (_Liquiditymanager *LiquiditymanagerFilterer) FilterAddLiquidity(opts *bind.FilterOpts, nftId []*big.Int) (*LiquiditymanagerAddLiquidityIterator, error) {

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Liquiditymanager.contract.FilterLogs(opts, "AddLiquidity", nftIdRule)
	if err != nil {
		return nil, err
	}
	return &LiquiditymanagerAddLiquidityIterator{contract: _Liquiditymanager.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0xf565fdd70b3936f0ae8efc41c2e0822f9de5ecb4dc162b153b129ec4bb9cd93c.
//
// Solidity: event AddLiquidity(uint256 indexed nftId, address pool, uint128 liquidityDelta, uint256 amountX, uint256 amountY)
func (_Liquiditymanager *LiquiditymanagerFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *LiquiditymanagerAddLiquidity, nftId []*big.Int) (event.Subscription, error) {

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Liquiditymanager.contract.WatchLogs(opts, "AddLiquidity", nftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquiditymanagerAddLiquidity)
				if err := _Liquiditymanager.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0xf565fdd70b3936f0ae8efc41c2e0822f9de5ecb4dc162b153b129ec4bb9cd93c.
//
// Solidity: event AddLiquidity(uint256 indexed nftId, address pool, uint128 liquidityDelta, uint256 amountX, uint256 amountY)
func (_Liquiditymanager *LiquiditymanagerFilterer) ParseAddLiquidity(log types.Log) (*LiquiditymanagerAddLiquidity, error) {
	event := new(LiquiditymanagerAddLiquidity)
	if err := _Liquiditymanager.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquiditymanagerApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Liquiditymanager contract.
type LiquiditymanagerApprovalIterator struct {
	Event *LiquiditymanagerApproval // Event containing the contract specifics and raw log

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
func (it *LiquiditymanagerApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquiditymanagerApproval)
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
		it.Event = new(LiquiditymanagerApproval)
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
func (it *LiquiditymanagerApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquiditymanagerApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquiditymanagerApproval represents a Approval event raised by the Liquiditymanager contract.
type LiquiditymanagerApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Liquiditymanager *LiquiditymanagerFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*LiquiditymanagerApprovalIterator, error) {

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

	logs, sub, err := _Liquiditymanager.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LiquiditymanagerApprovalIterator{contract: _Liquiditymanager.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Liquiditymanager *LiquiditymanagerFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LiquiditymanagerApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Liquiditymanager.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquiditymanagerApproval)
				if err := _Liquiditymanager.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Liquiditymanager *LiquiditymanagerFilterer) ParseApproval(log types.Log) (*LiquiditymanagerApproval, error) {
	event := new(LiquiditymanagerApproval)
	if err := _Liquiditymanager.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquiditymanagerApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Liquiditymanager contract.
type LiquiditymanagerApprovalForAllIterator struct {
	Event *LiquiditymanagerApprovalForAll // Event containing the contract specifics and raw log

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
func (it *LiquiditymanagerApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquiditymanagerApprovalForAll)
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
		it.Event = new(LiquiditymanagerApprovalForAll)
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
func (it *LiquiditymanagerApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquiditymanagerApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquiditymanagerApprovalForAll represents a ApprovalForAll event raised by the Liquiditymanager contract.
type LiquiditymanagerApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Liquiditymanager *LiquiditymanagerFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*LiquiditymanagerApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Liquiditymanager.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &LiquiditymanagerApprovalForAllIterator{contract: _Liquiditymanager.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Liquiditymanager *LiquiditymanagerFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LiquiditymanagerApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Liquiditymanager.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquiditymanagerApprovalForAll)
				if err := _Liquiditymanager.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Liquiditymanager *LiquiditymanagerFilterer) ParseApprovalForAll(log types.Log) (*LiquiditymanagerApprovalForAll, error) {
	event := new(LiquiditymanagerApprovalForAll)
	if err := _Liquiditymanager.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquiditymanagerDecLiquidityIterator is returned from FilterDecLiquidity and is used to iterate over the raw logs and unpacked data for DecLiquidity events raised by the Liquiditymanager contract.
type LiquiditymanagerDecLiquidityIterator struct {
	Event *LiquiditymanagerDecLiquidity // Event containing the contract specifics and raw log

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
func (it *LiquiditymanagerDecLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquiditymanagerDecLiquidity)
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
		it.Event = new(LiquiditymanagerDecLiquidity)
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
func (it *LiquiditymanagerDecLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquiditymanagerDecLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquiditymanagerDecLiquidity represents a DecLiquidity event raised by the Liquiditymanager contract.
type LiquiditymanagerDecLiquidity struct {
	NftId          *big.Int
	Pool           common.Address
	LiquidityDelta *big.Int
	AmountX        *big.Int
	AmountY        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDecLiquidity is a free log retrieval operation binding the contract event 0x24f4b91fa7871755148bc2a9e01f85d6fd73ec2a0e6bd9a5717c0d7f5be8c2c3.
//
// Solidity: event DecLiquidity(uint256 indexed nftId, address pool, uint128 liquidityDelta, uint256 amountX, uint256 amountY)
func (_Liquiditymanager *LiquiditymanagerFilterer) FilterDecLiquidity(opts *bind.FilterOpts, nftId []*big.Int) (*LiquiditymanagerDecLiquidityIterator, error) {

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Liquiditymanager.contract.FilterLogs(opts, "DecLiquidity", nftIdRule)
	if err != nil {
		return nil, err
	}
	return &LiquiditymanagerDecLiquidityIterator{contract: _Liquiditymanager.contract, event: "DecLiquidity", logs: logs, sub: sub}, nil
}

// WatchDecLiquidity is a free log subscription operation binding the contract event 0x24f4b91fa7871755148bc2a9e01f85d6fd73ec2a0e6bd9a5717c0d7f5be8c2c3.
//
// Solidity: event DecLiquidity(uint256 indexed nftId, address pool, uint128 liquidityDelta, uint256 amountX, uint256 amountY)
func (_Liquiditymanager *LiquiditymanagerFilterer) WatchDecLiquidity(opts *bind.WatchOpts, sink chan<- *LiquiditymanagerDecLiquidity, nftId []*big.Int) (event.Subscription, error) {

	var nftIdRule []interface{}
	for _, nftIdItem := range nftId {
		nftIdRule = append(nftIdRule, nftIdItem)
	}

	logs, sub, err := _Liquiditymanager.contract.WatchLogs(opts, "DecLiquidity", nftIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquiditymanagerDecLiquidity)
				if err := _Liquiditymanager.contract.UnpackLog(event, "DecLiquidity", log); err != nil {
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

// ParseDecLiquidity is a log parse operation binding the contract event 0x24f4b91fa7871755148bc2a9e01f85d6fd73ec2a0e6bd9a5717c0d7f5be8c2c3.
//
// Solidity: event DecLiquidity(uint256 indexed nftId, address pool, uint128 liquidityDelta, uint256 amountX, uint256 amountY)
func (_Liquiditymanager *LiquiditymanagerFilterer) ParseDecLiquidity(log types.Log) (*LiquiditymanagerDecLiquidity, error) {
	event := new(LiquiditymanagerDecLiquidity)
	if err := _Liquiditymanager.contract.UnpackLog(event, "DecLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquiditymanagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Liquiditymanager contract.
type LiquiditymanagerOwnershipTransferredIterator struct {
	Event *LiquiditymanagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LiquiditymanagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquiditymanagerOwnershipTransferred)
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
		it.Event = new(LiquiditymanagerOwnershipTransferred)
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
func (it *LiquiditymanagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquiditymanagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquiditymanagerOwnershipTransferred represents a OwnershipTransferred event raised by the Liquiditymanager contract.
type LiquiditymanagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Liquiditymanager *LiquiditymanagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LiquiditymanagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Liquiditymanager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LiquiditymanagerOwnershipTransferredIterator{contract: _Liquiditymanager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Liquiditymanager *LiquiditymanagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LiquiditymanagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Liquiditymanager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquiditymanagerOwnershipTransferred)
				if err := _Liquiditymanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Liquiditymanager *LiquiditymanagerFilterer) ParseOwnershipTransferred(log types.Log) (*LiquiditymanagerOwnershipTransferred, error) {
	event := new(LiquiditymanagerOwnershipTransferred)
	if err := _Liquiditymanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquiditymanagerTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Liquiditymanager contract.
type LiquiditymanagerTransferIterator struct {
	Event *LiquiditymanagerTransfer // Event containing the contract specifics and raw log

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
func (it *LiquiditymanagerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquiditymanagerTransfer)
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
		it.Event = new(LiquiditymanagerTransfer)
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
func (it *LiquiditymanagerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquiditymanagerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquiditymanagerTransfer represents a Transfer event raised by the Liquiditymanager contract.
type LiquiditymanagerTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Liquiditymanager *LiquiditymanagerFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*LiquiditymanagerTransferIterator, error) {

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

	logs, sub, err := _Liquiditymanager.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LiquiditymanagerTransferIterator{contract: _Liquiditymanager.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Liquiditymanager *LiquiditymanagerFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LiquiditymanagerTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Liquiditymanager.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquiditymanagerTransfer)
				if err := _Liquiditymanager.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Liquiditymanager *LiquiditymanagerFilterer) ParseTransfer(log types.Log) (*LiquiditymanagerTransfer, error) {
	event := new(LiquiditymanagerTransfer)
	if err := _Liquiditymanager.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
