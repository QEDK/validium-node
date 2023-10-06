// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dabridgerouter

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

// DabridgerouterMetaData contains all meta data concerning the Dabridgerouter contract.
var DabridgerouterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"originAndNonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"DataRootReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_router\",\"type\":\"bytes32\"}],\"name\":\"enrollRemoteRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"handle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_xAppConnectionManager\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"availDomain\",\"type\":\"uint32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"remotes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"roots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_xAppConnectionManager\",\"type\":\"address\"}],\"name\":\"setXAppConnectionManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xAppConnectionManager\",\"outputs\":[{\"internalType\":\"contractXAppConnectionManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611dae806100206000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80638da5cb5b11610081578063c5e4c9f91161005b578063c5e4c9f914610288578063f2fde38b146102c7578063ffa1ad74146102fa576100c9565b80638da5cb5b14610194578063ab2dc3f51461019c578063b49c53a71461025f576100c9565b806341bdc8b5116100b257806341bdc8b514610134578063715018a61461016957806383bbb80614610171576100c9565b8063081dc681146100ce5780633339df9614610103575b600080fd5b6100f1600480360360208110156100e457600080fd5b503563ffffffff16610318565b60408051918252519081900360200190f35b61010b61032a565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6101676004803603602081101561014a57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610346565b005b610167610435565b6100f16004803603602081101561018757600080fd5b503563ffffffff1661054c565b61010b61055e565b610167600480360360808110156101b257600080fd5b63ffffffff8235811692602081013590911691604082013591908101906080810160608201356401000000008111156101ea57600080fd5b8201836020820111156101fc57600080fd5b8035906020019184600183028401116401000000008311171561021e57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061057a945050505050565b6101676004803603604081101561027557600080fd5b5063ffffffff81351690602001356107c6565b6101676004803603604081101561029e57600080fd5b50803573ffffffffffffffffffffffffffffffffffffffff16906020013563ffffffff1661091a565b610167600480360360208110156102dd57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610a69565b610302610c0b565b6040805160ff9092168252519081900360200190f35b60c96020526000908152604090205481565b60655473ffffffffffffffffffffffffffffffffffffffff1681565b61034e610c10565b73ffffffffffffffffffffffffffffffffffffffff1661036c61055e565b73ffffffffffffffffffffffffffffffffffffffff16146103ee57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b606580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b61043d610c10565b73ffffffffffffffffffffffffffffffffffffffff1661045b61055e565b73ffffffffffffffffffffffffffffffffffffffff16146104dd57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b60335460405160009173ffffffffffffffffffffffffffffffffffffffff16907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055565b60976020526000908152604090205481565b60335473ffffffffffffffffffffffffffffffffffffffff1690565b61058333610c14565b6105ee57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f217265706c696361000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b83826105fa8282610cbf565b61066557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f2172656d6f746520726f75746572000000000000000000000000000000000000604482015290519081900360640190fd5b60ca5463ffffffff8781169116146106de57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f2176616c696420646f6d61696e00000000000000000000000000000000000000604482015290519081900360640190fd5b60006107146106ed8583610ceb565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000016610d0f565b90506107417fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000008216610d6a565b1561075657610751878783610d87565b6107bd565b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f2176616c6964206d657373616765000000000000000000000000000000000000604482015290519081900360640190fd5b50505050505050565b6107ce610c10565b73ffffffffffffffffffffffffffffffffffffffff166107ec61055e565b73ffffffffffffffffffffffffffffffffffffffff161461086e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b610876610e1f565b63ffffffff168263ffffffff1614158015610896575063ffffffff821615155b61090157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f21646f6d61696e00000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b63ffffffff909116600090815260976020526040902055565b600054610100900460ff16806109335750610933610ebb565b80610941575060005460ff16155b610996576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e815260200180611d05602e913960400191505060405180910390fd5b600054610100900460ff161580156109fc57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b610a0583610ecc565b60ca80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff84161790558015610a6457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b505050565b610a71610c10565b73ffffffffffffffffffffffffffffffffffffffff16610a8f61055e565b73ffffffffffffffffffffffffffffffffffffffff1614610b1157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff8116610b7d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180611cbe6026913960400191505060405180910390fd5b60335460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600081565b3390565b606554604080517f5190bc5300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015291516000939290921691635190bc5391602480820192602092909190829003018186803b158015610c8b57600080fd5b505afa158015610c9f573d6000803e3d6000fd5b505050506040513d6020811015610cb557600080fd5b505190505b919050565b63ffffffff821660009081526097602052604081205482148015610ce257508115155b90505b92915050565b815160009060208401610d0664ffffffffff85168284611029565b95945050505050565b600080610d1b8361107f565b60ff166001811115610d2957fe5b9050610d63816001811115610d3a57fe5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000008516906110af565b9392505050565b6000610d778260016110d5565b8015610ce55750610ce5826110ff565b600080610d9383611145565b63ffffffff8216600090815260c96020526040902054919350915015610db557fe5b63ffffffff8216600081815260c960205260409020829055610dd786866111a7565b67ffffffffffffffff167f9da097fed283637232607265455165441cbe1a0e2869469e2eb118149fcc0931836040518082815260200191505060405180910390a35050505050565b606554604080517f8d3638f4000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff1691638d3638f4916004808301926020929190829003018186803b158015610e8a57600080fd5b505afa158015610e9e573d6000803e3d6000fd5b505050506040513d6020811015610eb457600080fd5b5051905090565b6000610ec6306111c1565b15905090565b600054610100900460ff1680610ee55750610ee5610ebb565b80610ef3575060005460ff16155b610f48576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e815260200180611d05602e913960400191505060405180910390fd5b600054610100900460ff16158015610fae57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b606580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8416179055610ff66111c7565b801561102557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b5050565b60008061103684846112eb565b9050604051811115611046575060005b80611074577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000915050610d63565b610d0685858561135d565b6000610ce57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000008316826001611370565b60d81b7affffffffffffffffffffffffffffffffffffffffffffffffffffff9091161790565b60008160018111156110e357fe5b6110ec84611391565b60018111156110f757fe5b149392505050565b60008061112d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000084166113cc565b6bffffffffffffffffffffffff166025149392505050565b6000806111737fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000084166113e0565b91506111a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000008416611411565b9050915091565b63ffffffff1660209190911b67ffffffff00000000161790565b3b151590565b600054610100900460ff16806111e057506111e0610ebb565b806111ee575060005460ff16155b611243576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e815260200180611d05602e913960400191505060405180910390fd5b600054610100900460ff161580156112a957600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b6112b1611442565b6112b9611554565b80156112e857600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b50565b81810182811015610ce557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f766572666c6f7720647572696e67206164646974696f6e2e00000000000000604482015290519081900360640190fd5b606092831b9190911790911b1760181b90565b60008160200360080260ff166113878585856116e4565b901c949350505050565b60006113be7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000831661188f565b60ff166001811115610ce557fe5b60181c6bffffffffffffffffffffffff1690565b6000610ce57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000831660016004611370565b6000610ce57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000008316600560206116e4565b600054610100900460ff168061145b575061145b610ebb565b80611469575060005460ff16155b6114be576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e815260200180611d05602e913960400191505060405180910390fd5b600054610100900460ff161580156112b957600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff9091166101001716600117905580156112e857600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905550565b600054610100900460ff168061156d575061156d610ebb565b8061157b575060005460ff16155b6115d0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e815260200180611d05602e913960400191505060405180910390fd5b600054610100900460ff1615801561163657600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b6000611640610c10565b603380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8316908117909155604051919250906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35080156112e857600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905550565b600060ff82166116f657506000610d63565b6116ff846113cc565b6bffffffffffffffffffffffff1661171a8460ff85166112eb565b11156117f95761175b61172c85611895565b6bffffffffffffffffffffffff16611743866113cc565b6bffffffffffffffffffffffff16858560ff166118a9565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156117be5781810151838201526020016117a6565b50505050905090810190601f1680156117eb5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b60208260ff161115611856576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603a815260200180611d33603a913960400191505060405180910390fd5b60088202600061186586611895565b6bffffffffffffffffffffffff169050600061188083611a04565b91909501511695945050505050565b60d81c90565b60781c6bffffffffffffffffffffffff1690565b606060006118b686611a4d565b91505060006118c486611a4d565b91505060006118d286611a4d565b91505060006118e086611a4d565b915050838383836040516020018080611d6d603591397fffffffffffff000000000000000000000000000000000000000000000000000060d087811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b84015286901b16604a8201526050016021611ce482397fffffffffffff000000000000000000000000000000000000000000000000000060d094851b811660218301527f2077697468206c656e677468203078000000000000000000000000000000000060278301529290931b9091166036830152507f2e00000000000000000000000000000000000000000000000000000000000000603c82015260408051601d818403018152603d90920190529b9a5050505050505050505050565b7f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9091011d90565b600080601f5b600f8160ff161115611ab55760ff600882021684901c611a7281611b21565b61ffff16841793508160ff16601014611a8d57601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611a53565b50600f5b60ff8160ff161015611b1b5760ff600882021684901c611ad881611b21565b61ffff16831792508160ff16600014611af357601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611ab9565b50915091565b6000611b3360048360ff16901c611b51565b60ff161760081b62ffff0016611b4882611b51565b60ff1617919050565b600060f08083179060ff82161415611b6d576030915050610cba565b8060ff1660f11415611b83576031915050610cba565b8060ff1660f21415611b99576032915050610cba565b8060ff1660f31415611baf576033915050610cba565b8060ff1660f41415611bc5576034915050610cba565b8060ff1660f51415611bdb576035915050610cba565b8060ff1660f61415611bf1576036915050610cba565b8060ff1660f71415611c07576037915050610cba565b8060ff1660f81415611c1d576038915050610cba565b8060ff1660f91415611c33576039915050610cba565b8060ff1660fa1415611c49576061915050610cba565b8060ff1660fb1415611c5f576062915050610cba565b8060ff1660fc1415611c75576063915050610cba565b8060ff1660fd1415611c8b576064915050610cba565b8060ff1660fe1415611ca1576065915050610cba565b8060ff1660ff1415611cb7576066915050610cba565b5091905056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573732e20417474656d7074656420746f20696e646578206174206f6666736574203078496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a656454797065644d656d566965772f696e646578202d20417474656d7074656420746f20696e646578206d6f7265207468616e20333220627974657354797065644d656d566965772f696e646578202d204f76657272616e2074686520766965772e20536c696365206973206174203078a164736f6c6343000706000a",
}

// DabridgerouterABI is the input ABI used to generate the binding from.
// Deprecated: Use DabridgerouterMetaData.ABI instead.
var DabridgerouterABI = DabridgerouterMetaData.ABI

// DabridgerouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DabridgerouterMetaData.Bin instead.
var DabridgerouterBin = DabridgerouterMetaData.Bin

// DeployDabridgerouter deploys a new Ethereum contract, binding an instance of Dabridgerouter to it.
func DeployDabridgerouter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Dabridgerouter, error) {
	parsed, err := DabridgerouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DabridgerouterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dabridgerouter{DabridgerouterCaller: DabridgerouterCaller{contract: contract}, DabridgerouterTransactor: DabridgerouterTransactor{contract: contract}, DabridgerouterFilterer: DabridgerouterFilterer{contract: contract}}, nil
}

// Dabridgerouter is an auto generated Go binding around an Ethereum contract.
type Dabridgerouter struct {
	DabridgerouterCaller     // Read-only binding to the contract
	DabridgerouterTransactor // Write-only binding to the contract
	DabridgerouterFilterer   // Log filterer for contract events
}

// DabridgerouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type DabridgerouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DabridgerouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DabridgerouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DabridgerouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DabridgerouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DabridgerouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DabridgerouterSession struct {
	Contract     *Dabridgerouter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DabridgerouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DabridgerouterCallerSession struct {
	Contract *DabridgerouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DabridgerouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DabridgerouterTransactorSession struct {
	Contract     *DabridgerouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DabridgerouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type DabridgerouterRaw struct {
	Contract *Dabridgerouter // Generic contract binding to access the raw methods on
}

// DabridgerouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DabridgerouterCallerRaw struct {
	Contract *DabridgerouterCaller // Generic read-only contract binding to access the raw methods on
}

// DabridgerouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DabridgerouterTransactorRaw struct {
	Contract *DabridgerouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDabridgerouter creates a new instance of Dabridgerouter, bound to a specific deployed contract.
func NewDabridgerouter(address common.Address, backend bind.ContractBackend) (*Dabridgerouter, error) {
	contract, err := bindDabridgerouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dabridgerouter{DabridgerouterCaller: DabridgerouterCaller{contract: contract}, DabridgerouterTransactor: DabridgerouterTransactor{contract: contract}, DabridgerouterFilterer: DabridgerouterFilterer{contract: contract}}, nil
}

// NewDabridgerouterCaller creates a new read-only instance of Dabridgerouter, bound to a specific deployed contract.
func NewDabridgerouterCaller(address common.Address, caller bind.ContractCaller) (*DabridgerouterCaller, error) {
	contract, err := bindDabridgerouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DabridgerouterCaller{contract: contract}, nil
}

// NewDabridgerouterTransactor creates a new write-only instance of Dabridgerouter, bound to a specific deployed contract.
func NewDabridgerouterTransactor(address common.Address, transactor bind.ContractTransactor) (*DabridgerouterTransactor, error) {
	contract, err := bindDabridgerouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DabridgerouterTransactor{contract: contract}, nil
}

// NewDabridgerouterFilterer creates a new log filterer instance of Dabridgerouter, bound to a specific deployed contract.
func NewDabridgerouterFilterer(address common.Address, filterer bind.ContractFilterer) (*DabridgerouterFilterer, error) {
	contract, err := bindDabridgerouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DabridgerouterFilterer{contract: contract}, nil
}

// bindDabridgerouter binds a generic wrapper to an already deployed contract.
func bindDabridgerouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DabridgerouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dabridgerouter *DabridgerouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dabridgerouter.Contract.DabridgerouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dabridgerouter *DabridgerouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dabridgerouter.Contract.DabridgerouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dabridgerouter *DabridgerouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dabridgerouter.Contract.DabridgerouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dabridgerouter *DabridgerouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dabridgerouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dabridgerouter *DabridgerouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dabridgerouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dabridgerouter *DabridgerouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dabridgerouter.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Dabridgerouter *DabridgerouterCaller) VERSION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Dabridgerouter.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Dabridgerouter *DabridgerouterSession) VERSION() (uint8, error) {
	return _Dabridgerouter.Contract.VERSION(&_Dabridgerouter.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Dabridgerouter *DabridgerouterCallerSession) VERSION() (uint8, error) {
	return _Dabridgerouter.Contract.VERSION(&_Dabridgerouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dabridgerouter *DabridgerouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dabridgerouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dabridgerouter *DabridgerouterSession) Owner() (common.Address, error) {
	return _Dabridgerouter.Contract.Owner(&_Dabridgerouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dabridgerouter *DabridgerouterCallerSession) Owner() (common.Address, error) {
	return _Dabridgerouter.Contract.Owner(&_Dabridgerouter.CallOpts)
}

// Remotes is a free data retrieval call binding the contract method 0x83bbb806.
//
// Solidity: function remotes(uint32 ) view returns(bytes32)
func (_Dabridgerouter *DabridgerouterCaller) Remotes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _Dabridgerouter.contract.Call(opts, &out, "remotes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Remotes is a free data retrieval call binding the contract method 0x83bbb806.
//
// Solidity: function remotes(uint32 ) view returns(bytes32)
func (_Dabridgerouter *DabridgerouterSession) Remotes(arg0 uint32) ([32]byte, error) {
	return _Dabridgerouter.Contract.Remotes(&_Dabridgerouter.CallOpts, arg0)
}

// Remotes is a free data retrieval call binding the contract method 0x83bbb806.
//
// Solidity: function remotes(uint32 ) view returns(bytes32)
func (_Dabridgerouter *DabridgerouterCallerSession) Remotes(arg0 uint32) ([32]byte, error) {
	return _Dabridgerouter.Contract.Remotes(&_Dabridgerouter.CallOpts, arg0)
}

// Roots is a free data retrieval call binding the contract method 0x081dc681.
//
// Solidity: function roots(uint32 ) view returns(bytes32)
func (_Dabridgerouter *DabridgerouterCaller) Roots(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _Dabridgerouter.contract.Call(opts, &out, "roots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Roots is a free data retrieval call binding the contract method 0x081dc681.
//
// Solidity: function roots(uint32 ) view returns(bytes32)
func (_Dabridgerouter *DabridgerouterSession) Roots(arg0 uint32) ([32]byte, error) {
	return _Dabridgerouter.Contract.Roots(&_Dabridgerouter.CallOpts, arg0)
}

// Roots is a free data retrieval call binding the contract method 0x081dc681.
//
// Solidity: function roots(uint32 ) view returns(bytes32)
func (_Dabridgerouter *DabridgerouterCallerSession) Roots(arg0 uint32) ([32]byte, error) {
	return _Dabridgerouter.Contract.Roots(&_Dabridgerouter.CallOpts, arg0)
}

// XAppConnectionManager is a free data retrieval call binding the contract method 0x3339df96.
//
// Solidity: function xAppConnectionManager() view returns(address)
func (_Dabridgerouter *DabridgerouterCaller) XAppConnectionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dabridgerouter.contract.Call(opts, &out, "xAppConnectionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XAppConnectionManager is a free data retrieval call binding the contract method 0x3339df96.
//
// Solidity: function xAppConnectionManager() view returns(address)
func (_Dabridgerouter *DabridgerouterSession) XAppConnectionManager() (common.Address, error) {
	return _Dabridgerouter.Contract.XAppConnectionManager(&_Dabridgerouter.CallOpts)
}

// XAppConnectionManager is a free data retrieval call binding the contract method 0x3339df96.
//
// Solidity: function xAppConnectionManager() view returns(address)
func (_Dabridgerouter *DabridgerouterCallerSession) XAppConnectionManager() (common.Address, error) {
	return _Dabridgerouter.Contract.XAppConnectionManager(&_Dabridgerouter.CallOpts)
}

// EnrollRemoteRouter is a paid mutator transaction binding the contract method 0xb49c53a7.
//
// Solidity: function enrollRemoteRouter(uint32 _domain, bytes32 _router) returns()
func (_Dabridgerouter *DabridgerouterTransactor) EnrollRemoteRouter(opts *bind.TransactOpts, _domain uint32, _router [32]byte) (*types.Transaction, error) {
	return _Dabridgerouter.contract.Transact(opts, "enrollRemoteRouter", _domain, _router)
}

// EnrollRemoteRouter is a paid mutator transaction binding the contract method 0xb49c53a7.
//
// Solidity: function enrollRemoteRouter(uint32 _domain, bytes32 _router) returns()
func (_Dabridgerouter *DabridgerouterSession) EnrollRemoteRouter(_domain uint32, _router [32]byte) (*types.Transaction, error) {
	return _Dabridgerouter.Contract.EnrollRemoteRouter(&_Dabridgerouter.TransactOpts, _domain, _router)
}

// EnrollRemoteRouter is a paid mutator transaction binding the contract method 0xb49c53a7.
//
// Solidity: function enrollRemoteRouter(uint32 _domain, bytes32 _router) returns()
func (_Dabridgerouter *DabridgerouterTransactorSession) EnrollRemoteRouter(_domain uint32, _router [32]byte) (*types.Transaction, error) {
	return _Dabridgerouter.Contract.EnrollRemoteRouter(&_Dabridgerouter.TransactOpts, _domain, _router)
}

// Handle is a paid mutator transaction binding the contract method 0xab2dc3f5.
//
// Solidity: function handle(uint32 _origin, uint32 _nonce, bytes32 _sender, bytes _message) returns()
func (_Dabridgerouter *DabridgerouterTransactor) Handle(opts *bind.TransactOpts, _origin uint32, _nonce uint32, _sender [32]byte, _message []byte) (*types.Transaction, error) {
	return _Dabridgerouter.contract.Transact(opts, "handle", _origin, _nonce, _sender, _message)
}

// Handle is a paid mutator transaction binding the contract method 0xab2dc3f5.
//
// Solidity: function handle(uint32 _origin, uint32 _nonce, bytes32 _sender, bytes _message) returns()
func (_Dabridgerouter *DabridgerouterSession) Handle(_origin uint32, _nonce uint32, _sender [32]byte, _message []byte) (*types.Transaction, error) {
	return _Dabridgerouter.Contract.Handle(&_Dabridgerouter.TransactOpts, _origin, _nonce, _sender, _message)
}

// Handle is a paid mutator transaction binding the contract method 0xab2dc3f5.
//
// Solidity: function handle(uint32 _origin, uint32 _nonce, bytes32 _sender, bytes _message) returns()
func (_Dabridgerouter *DabridgerouterTransactorSession) Handle(_origin uint32, _nonce uint32, _sender [32]byte, _message []byte) (*types.Transaction, error) {
	return _Dabridgerouter.Contract.Handle(&_Dabridgerouter.TransactOpts, _origin, _nonce, _sender, _message)
}

// Initialize is a paid mutator transaction binding the contract method 0xc5e4c9f9.
//
// Solidity: function initialize(address _xAppConnectionManager, uint32 availDomain) returns()
func (_Dabridgerouter *DabridgerouterTransactor) Initialize(opts *bind.TransactOpts, _xAppConnectionManager common.Address, availDomain uint32) (*types.Transaction, error) {
	return _Dabridgerouter.contract.Transact(opts, "initialize", _xAppConnectionManager, availDomain)
}

// Initialize is a paid mutator transaction binding the contract method 0xc5e4c9f9.
//
// Solidity: function initialize(address _xAppConnectionManager, uint32 availDomain) returns()
func (_Dabridgerouter *DabridgerouterSession) Initialize(_xAppConnectionManager common.Address, availDomain uint32) (*types.Transaction, error) {
	return _Dabridgerouter.Contract.Initialize(&_Dabridgerouter.TransactOpts, _xAppConnectionManager, availDomain)
}

// Initialize is a paid mutator transaction binding the contract method 0xc5e4c9f9.
//
// Solidity: function initialize(address _xAppConnectionManager, uint32 availDomain) returns()
func (_Dabridgerouter *DabridgerouterTransactorSession) Initialize(_xAppConnectionManager common.Address, availDomain uint32) (*types.Transaction, error) {
	return _Dabridgerouter.Contract.Initialize(&_Dabridgerouter.TransactOpts, _xAppConnectionManager, availDomain)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dabridgerouter *DabridgerouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dabridgerouter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dabridgerouter *DabridgerouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dabridgerouter.Contract.RenounceOwnership(&_Dabridgerouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dabridgerouter *DabridgerouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dabridgerouter.Contract.RenounceOwnership(&_Dabridgerouter.TransactOpts)
}

// SetXAppConnectionManager is a paid mutator transaction binding the contract method 0x41bdc8b5.
//
// Solidity: function setXAppConnectionManager(address _xAppConnectionManager) returns()
func (_Dabridgerouter *DabridgerouterTransactor) SetXAppConnectionManager(opts *bind.TransactOpts, _xAppConnectionManager common.Address) (*types.Transaction, error) {
	return _Dabridgerouter.contract.Transact(opts, "setXAppConnectionManager", _xAppConnectionManager)
}

// SetXAppConnectionManager is a paid mutator transaction binding the contract method 0x41bdc8b5.
//
// Solidity: function setXAppConnectionManager(address _xAppConnectionManager) returns()
func (_Dabridgerouter *DabridgerouterSession) SetXAppConnectionManager(_xAppConnectionManager common.Address) (*types.Transaction, error) {
	return _Dabridgerouter.Contract.SetXAppConnectionManager(&_Dabridgerouter.TransactOpts, _xAppConnectionManager)
}

// SetXAppConnectionManager is a paid mutator transaction binding the contract method 0x41bdc8b5.
//
// Solidity: function setXAppConnectionManager(address _xAppConnectionManager) returns()
func (_Dabridgerouter *DabridgerouterTransactorSession) SetXAppConnectionManager(_xAppConnectionManager common.Address) (*types.Transaction, error) {
	return _Dabridgerouter.Contract.SetXAppConnectionManager(&_Dabridgerouter.TransactOpts, _xAppConnectionManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dabridgerouter *DabridgerouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dabridgerouter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dabridgerouter *DabridgerouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dabridgerouter.Contract.TransferOwnership(&_Dabridgerouter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dabridgerouter *DabridgerouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dabridgerouter.Contract.TransferOwnership(&_Dabridgerouter.TransactOpts, newOwner)
}

// DabridgerouterDataRootReceivedIterator is returned from FilterDataRootReceived and is used to iterate over the raw logs and unpacked data for DataRootReceived events raised by the Dabridgerouter contract.
type DabridgerouterDataRootReceivedIterator struct {
	Event *DabridgerouterDataRootReceived // Event containing the contract specifics and raw log

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
func (it *DabridgerouterDataRootReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DabridgerouterDataRootReceived)
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
		it.Event = new(DabridgerouterDataRootReceived)
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
func (it *DabridgerouterDataRootReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DabridgerouterDataRootReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DabridgerouterDataRootReceived represents a DataRootReceived event raised by the Dabridgerouter contract.
type DabridgerouterDataRootReceived struct {
	OriginAndNonce uint64
	BlockNumber    uint32
	Root           [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDataRootReceived is a free log retrieval operation binding the contract event 0x9da097fed283637232607265455165441cbe1a0e2869469e2eb118149fcc0931.
//
// Solidity: event DataRootReceived(uint64 indexed originAndNonce, uint32 indexed blockNumber, bytes32 root)
func (_Dabridgerouter *DabridgerouterFilterer) FilterDataRootReceived(opts *bind.FilterOpts, originAndNonce []uint64, blockNumber []uint32) (*DabridgerouterDataRootReceivedIterator, error) {

	var originAndNonceRule []interface{}
	for _, originAndNonceItem := range originAndNonce {
		originAndNonceRule = append(originAndNonceRule, originAndNonceItem)
	}
	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}

	logs, sub, err := _Dabridgerouter.contract.FilterLogs(opts, "DataRootReceived", originAndNonceRule, blockNumberRule)
	if err != nil {
		return nil, err
	}
	return &DabridgerouterDataRootReceivedIterator{contract: _Dabridgerouter.contract, event: "DataRootReceived", logs: logs, sub: sub}, nil
}

// WatchDataRootReceived is a free log subscription operation binding the contract event 0x9da097fed283637232607265455165441cbe1a0e2869469e2eb118149fcc0931.
//
// Solidity: event DataRootReceived(uint64 indexed originAndNonce, uint32 indexed blockNumber, bytes32 root)
func (_Dabridgerouter *DabridgerouterFilterer) WatchDataRootReceived(opts *bind.WatchOpts, sink chan<- *DabridgerouterDataRootReceived, originAndNonce []uint64, blockNumber []uint32) (event.Subscription, error) {

	var originAndNonceRule []interface{}
	for _, originAndNonceItem := range originAndNonce {
		originAndNonceRule = append(originAndNonceRule, originAndNonceItem)
	}
	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}

	logs, sub, err := _Dabridgerouter.contract.WatchLogs(opts, "DataRootReceived", originAndNonceRule, blockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DabridgerouterDataRootReceived)
				if err := _Dabridgerouter.contract.UnpackLog(event, "DataRootReceived", log); err != nil {
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

// ParseDataRootReceived is a log parse operation binding the contract event 0x9da097fed283637232607265455165441cbe1a0e2869469e2eb118149fcc0931.
//
// Solidity: event DataRootReceived(uint64 indexed originAndNonce, uint32 indexed blockNumber, bytes32 root)
func (_Dabridgerouter *DabridgerouterFilterer) ParseDataRootReceived(log types.Log) (*DabridgerouterDataRootReceived, error) {
	event := new(DabridgerouterDataRootReceived)
	if err := _Dabridgerouter.contract.UnpackLog(event, "DataRootReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DabridgerouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dabridgerouter contract.
type DabridgerouterOwnershipTransferredIterator struct {
	Event *DabridgerouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DabridgerouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DabridgerouterOwnershipTransferred)
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
		it.Event = new(DabridgerouterOwnershipTransferred)
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
func (it *DabridgerouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DabridgerouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DabridgerouterOwnershipTransferred represents a OwnershipTransferred event raised by the Dabridgerouter contract.
type DabridgerouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dabridgerouter *DabridgerouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DabridgerouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dabridgerouter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DabridgerouterOwnershipTransferredIterator{contract: _Dabridgerouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dabridgerouter *DabridgerouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DabridgerouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dabridgerouter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DabridgerouterOwnershipTransferred)
				if err := _Dabridgerouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Dabridgerouter *DabridgerouterFilterer) ParseOwnershipTransferred(log types.Log) (*DabridgerouterOwnershipTransferred, error) {
	event := new(DabridgerouterOwnershipTransferred)
	if err := _Dabridgerouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
