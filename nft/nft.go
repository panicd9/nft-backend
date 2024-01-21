// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nft

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

// NftMetaData contains all meta data concerning the Nft contract.
var NftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Nft__AlreadyBlacklisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Nft__AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Nft__Blacklisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Nft__InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Nft__NotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Nft__TokenDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addedBy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"blacklistedAddress\",\"type\":\"address\"}],\"name\":\"Blacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MINTING_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REGISTRATION_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"addToBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"blacklistedAddresses\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a0604052600a6080523480156200001657600080fd5b5060408051808201825260038082526213999d60ea1b6020808401829052845180860190955291845290830152339160006200005383826200019f565b5060016200006282826200019f565b5050506001600160a01b0381166200009457604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b6200009f81620000a6565b506200026b565b600680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200012357607f821691505b6020821081036200014457634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200019a576000816000526020600020601f850160051c81016020861015620001755750805b601f850160051c820191505b81811015620001965782815560010162000181565b5050505b505050565b81516001600160401b03811115620001bb57620001bb620000f8565b620001d381620001cc84546200010e565b846200014a565b602080601f8311600181146200020b5760008415620001f25750858301515b600019600386901b1c1916600185901b17855562000196565b600085815260208120601f198616915b828110156200023c578886015182559484019460019091019084016200021b565b50858210156200025b5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6080516115646200028e60003960008181610333015261087401526115646000f3fe60806040526004361061014b5760003560e01c8063715018a6116100b6578063b2dd5c071161006f578063b2dd5c071461039d578063b88d4fde146103cd578063c87b56dd146103ed578063d5749d421461040d578063e985e9c51461043d578063f2fde38b1461045d57600080fd5b8063715018a6146102ee5780638da5cb5b14610303578063902d55a51461032157806395d89b4114610355578063a0712d681461036a578063a22cb4651461037d57600080fd5b80632663e0b5116101085780632663e0b51461022957806342842e0e1461025257806344337ea1146102725780636352211e1461029257806363d9d166146102b257806370a08231146102ce57600080fd5b806301ffc9a71461015057806306fdde0314610185578063081812fc146101a7578063095ea7b3146101df5780631aa3a0081461020157806323b872dd14610209575b600080fd5b34801561015c57600080fd5b5061017061016b36600461119e565b61047d565b60405190151581526020015b60405180910390f35b34801561019157600080fd5b5061019a6104cf565b60405161017c919061120b565b3480156101b357600080fd5b506101c76101c236600461121e565b610561565b6040516001600160a01b03909116815260200161017c565b3480156101eb57600080fd5b506101ff6101fa366004611253565b61058a565b005b6101ff610599565b34801561021557600080fd5b506101ff61022436600461127d565b61063d565b34801561023557600080fd5b50610244662386f26fc1000081565b60405190815260200161017c565b34801561025e57600080fd5b506101ff61026d36600461127d565b6106c8565b34801561027e57600080fd5b506101ff61028d3660046112b9565b6106e8565b34801561029e57600080fd5b506101c76102ad36600461121e565b610797565b3480156102be57600080fd5b506102446702c68af0bb14000081565b3480156102da57600080fd5b506102446102e93660046112b9565b6107a2565b3480156102fa57600080fd5b506101ff6107ea565b34801561030f57600080fd5b506006546001600160a01b03166101c7565b34801561032d57600080fd5b506102447f000000000000000000000000000000000000000000000000000000000000000081565b34801561036157600080fd5b5061019a6107fe565b6101ff61037836600461121e565b61080d565b34801561038957600080fd5b506101ff6103983660046112d4565b61091b565b3480156103a957600080fd5b506101706103b83660046112b9565b60086020526000908152604090205460ff1681565b3480156103d957600080fd5b506101ff6103e8366004611326565b610926565b3480156103f957600080fd5b5061019a61040836600461121e565b61093d565b34801561041957600080fd5b506101706104283660046112b9565b60076020526000908152604090205460ff1681565b34801561044957600080fd5b50610170610458366004611402565b6109a5565b34801561046957600080fd5b506101ff6104783660046112b9565b6109d3565b60006001600160e01b031982166380ac58cd60e01b14806104ae57506001600160e01b03198216635b5e139f60e01b145b806104c957506301ffc9a760e01b6001600160e01b03198316145b92915050565b6060600080546104de90611435565b80601f016020809104026020016040519081016040528092919081815260200182805461050a90611435565b80156105575780601f1061052c57610100808354040283529160200191610557565b820191906000526020600020905b81548152906001019060200180831161053a57829003601f168201915b5050505050905090565b600061056c82610a11565b506000828152600460205260409020546001600160a01b03166104c9565b610595828233610a4a565b5050565b3360009081526008602052604090205460ff16156105d15760405163a2dcdd8560e01b81523360048201526024015b60405180910390fd5b662386f26fc100003410156105fb5760405163e5b9458d60e01b81523460048201526024016105c8565b33600081815260086020526040808220805460ff19166001179055517f2d3734a8e47ac8316e500ac231c90a6e1848ca2285f40d07eaa52005e4b3a0e99190a2565b6001600160a01b03821661066757604051633250574960e11b8152600060048201526024016105c8565b6000610674838333610a57565b9050836001600160a01b0316816001600160a01b0316146106c2576040516364283d7b60e01b81526001600160a01b03808616600483015260248201849052821660448201526064016105c8565b50505050565b6106e383838360405180602001604052806000815250610926565b505050565b6106f0610b50565b6001600160a01b03811660009081526007602052604090205460ff1615610735576040516302327fb360e51b81526001600160a01b03821660048201526024016105c8565b6001600160a01b0381166000818152600760205260409020805460ff1916600117905561075f3390565b6001600160a01b03167fd36871fdf6981136f3ac0564927005901eda06f7a9dff1e8b2a1d7846b8ebb5060405160405180910390a350565b60006104c982610a11565b60006001600160a01b0382166107ce576040516322718ad960e21b8152600060048201526024016105c8565b506001600160a01b031660009081526003602052604090205490565b6107f2610b50565b6107fc6000610b7d565b565b6060600180546104de90611435565b3360009081526008602052604090205460ff1661083f57604051636a6aec4760e11b81523360048201526024016105c8565b3360009081526007602052604090205460ff16156108725760405163dba40efb60e01b81523360048201526024016105c8565b7f00000000000000000000000000000000000000000000000000000000000000008111156108b657604051635310fce960e11b8152600481018290526024016105c8565b6702c68af0bb1400003410156108e15760405163e5b9458d60e01b81523460048201526024016105c8565b6108eb3382610bcf565b604051819033907f30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe90600090a350565b610595338383610be9565b61093184848461063d565b6106c284848484610c88565b606061094882610a11565b506000610953610db1565b90506000815111610973576040518060200160405280600081525061099e565b8061097d84610dd1565b60405160200161098e92919061146f565b6040516020818303038152906040525b9392505050565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b6109db610b50565b6001600160a01b038116610a0557604051631e4fbdf760e01b8152600060048201526024016105c8565b610a0e81610b7d565b50565b6000818152600260205260408120546001600160a01b0316806104c957604051637e27328960e01b8152600481018490526024016105c8565b6106e38383836001610e64565b6000828152600260205260408120546001600160a01b0390811690831615610a8457610a84818486610f6a565b6001600160a01b03811615610ac257610aa1600085600080610e64565b6001600160a01b038116600090815260036020526040902080546000190190555b6001600160a01b03851615610af1576001600160a01b0385166000908152600360205260409020805460010190555b60008481526002602052604080822080546001600160a01b0319166001600160a01b0389811691821790925591518793918516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4949350505050565b6006546001600160a01b031633146107fc5760405163118cdaa760e01b81523360048201526024016105c8565b600680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b610595828260405180602001604052806000815250610fce565b6001600160a01b038216610c1b57604051630b61174360e31b81526001600160a01b03831660048201526024016105c8565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b0383163b156106c257604051630a85bd0160e11b81526001600160a01b0384169063150b7a0290610cca90339088908790879060040161149e565b6020604051808303816000875af1925050508015610d05575060408051601f3d908101601f19168201909252610d02918101906114db565b60015b610d6e573d808015610d33576040519150601f19603f3d011682016040523d82523d6000602084013e610d38565b606091505b508051600003610d6657604051633250574960e11b81526001600160a01b03851660048201526024016105c8565b805181602001fd5b6001600160e01b03198116630a85bd0160e11b14610daa57604051633250574960e11b81526001600160a01b03851660048201526024016105c8565b5050505050565b60606040518060600160405280603681526020016114f960369139905090565b60606000610dde83610fe5565b600101905060008167ffffffffffffffff811115610dfe57610dfe611310565b6040519080825280601f01601f191660200182016040528015610e28576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a8504945084610e3257509392505050565b8080610e7857506001600160a01b03821615155b15610f3a576000610e8884610a11565b90506001600160a01b03831615801590610eb45750826001600160a01b0316816001600160a01b031614155b8015610ec75750610ec581846109a5565b155b15610ef05760405163a9fbf51f60e01b81526001600160a01b03841660048201526024016105c8565b8115610f385783856001600160a01b0316826001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b5050600090815260046020526040902080546001600160a01b0319166001600160a01b0392909216919091179055565b610f758383836110bd565b6106e3576001600160a01b038316610fa357604051637e27328960e01b8152600481018290526024016105c8565b60405163177e802f60e01b81526001600160a01b0383166004820152602481018290526044016105c8565b610fd88383611123565b6106e36000848484610c88565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b83106110245772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef81000000008310611050576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061106e57662386f26fc10000830492506010015b6305f5e1008310611086576305f5e100830492506008015b612710831061109a57612710830492506004015b606483106110ac576064830492506002015b600a83106104c95760010192915050565b60006001600160a01b0383161580159061111b5750826001600160a01b0316846001600160a01b031614806110f757506110f784846109a5565b8061111b57506000828152600460205260409020546001600160a01b038481169116145b949350505050565b6001600160a01b03821661114d57604051633250574960e11b8152600060048201526024016105c8565b600061115b83836000610a57565b90506001600160a01b038116156106e3576040516339e3563760e11b8152600060048201526024016105c8565b6001600160e01b031981168114610a0e57600080fd5b6000602082840312156111b057600080fd5b813561099e81611188565b60005b838110156111d65781810151838201526020016111be565b50506000910152565b600081518084526111f78160208601602086016111bb565b601f01601f19169290920160200192915050565b60208152600061099e60208301846111df565b60006020828403121561123057600080fd5b5035919050565b80356001600160a01b038116811461124e57600080fd5b919050565b6000806040838503121561126657600080fd5b61126f83611237565b946020939093013593505050565b60008060006060848603121561129257600080fd5b61129b84611237565b92506112a960208501611237565b9150604084013590509250925092565b6000602082840312156112cb57600080fd5b61099e82611237565b600080604083850312156112e757600080fd5b6112f083611237565b91506020830135801515811461130557600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b6000806000806080858703121561133c57600080fd5b61134585611237565b935061135360208601611237565b925060408501359150606085013567ffffffffffffffff8082111561137757600080fd5b818701915087601f83011261138b57600080fd5b81358181111561139d5761139d611310565b604051601f8201601f19908116603f011681019083821181831017156113c5576113c5611310565b816040528281528a60208487010111156113de57600080fd5b82602086016020830137600060208483010152809550505050505092959194509250565b6000806040838503121561141557600080fd5b61141e83611237565b915061142c60208401611237565b90509250929050565b600181811c9082168061144957607f821691505b60208210810361146957634e487b7160e01b600052602260045260246000fd5b50919050565b600083516114818184602088016111bb565b8351908301906114958183602088016111bb565b01949350505050565b6001600160a01b03858116825284166020820152604081018390526080606082018190526000906114d1908301846111df565b9695505050505050565b6000602082840312156114ed57600080fd5b815161099e8161118856fe697066733a2f2f516d624877354b614d4c784d59646a436f476d6f384c47355775725165527044426765697a386f78464438796e6f2fa2646970667358221220cf984879931b6bfb98f57abeb48e3fefa5f23537935800ca58f482a5386d23b364736f6c63430008170033",
}

// NftABI is the input ABI used to generate the binding from.
// Deprecated: Use NftMetaData.ABI instead.
var NftABI = NftMetaData.ABI

// NftBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NftMetaData.Bin instead.
var NftBin = NftMetaData.Bin

// DeployNft deploys a new Ethereum contract, binding an instance of Nft to it.
func DeployNft(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Nft, error) {
	parsed, err := NftMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NftBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Nft{NftCaller: NftCaller{contract: contract}, NftTransactor: NftTransactor{contract: contract}, NftFilterer: NftFilterer{contract: contract}}, nil
}

// Nft is an auto generated Go binding around an Ethereum contract.
type Nft struct {
	NftCaller     // Read-only binding to the contract
	NftTransactor // Write-only binding to the contract
	NftFilterer   // Log filterer for contract events
}

// NftCaller is an auto generated read-only Go binding around an Ethereum contract.
type NftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NftSession struct {
	Contract     *Nft              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NftCallerSession struct {
	Contract *NftCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NftTransactorSession struct {
	Contract     *NftTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NftRaw is an auto generated low-level Go binding around an Ethereum contract.
type NftRaw struct {
	Contract *Nft // Generic contract binding to access the raw methods on
}

// NftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NftCallerRaw struct {
	Contract *NftCaller // Generic read-only contract binding to access the raw methods on
}

// NftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NftTransactorRaw struct {
	Contract *NftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNft creates a new instance of Nft, bound to a specific deployed contract.
func NewNft(address common.Address, backend bind.ContractBackend) (*Nft, error) {
	contract, err := bindNft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nft{NftCaller: NftCaller{contract: contract}, NftTransactor: NftTransactor{contract: contract}, NftFilterer: NftFilterer{contract: contract}}, nil
}

// NewNftCaller creates a new read-only instance of Nft, bound to a specific deployed contract.
func NewNftCaller(address common.Address, caller bind.ContractCaller) (*NftCaller, error) {
	contract, err := bindNft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NftCaller{contract: contract}, nil
}

// NewNftTransactor creates a new write-only instance of Nft, bound to a specific deployed contract.
func NewNftTransactor(address common.Address, transactor bind.ContractTransactor) (*NftTransactor, error) {
	contract, err := bindNft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NftTransactor{contract: contract}, nil
}

// NewNftFilterer creates a new log filterer instance of Nft, bound to a specific deployed contract.
func NewNftFilterer(address common.Address, filterer bind.ContractFilterer) (*NftFilterer, error) {
	contract, err := bindNft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NftFilterer{contract: contract}, nil
}

// bindNft binds a generic wrapper to an already deployed contract.
func bindNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nft *NftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nft.Contract.NftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nft *NftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nft.Contract.NftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nft *NftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nft.Contract.NftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nft *NftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nft *NftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nft *NftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nft.Contract.contract.Transact(opts, method, params...)
}

// MINTINGPRICE is a free data retrieval call binding the contract method 0x63d9d166.
//
// Solidity: function MINTING_PRICE() view returns(uint256)
func (_Nft *NftCaller) MINTINGPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "MINTING_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTINGPRICE is a free data retrieval call binding the contract method 0x63d9d166.
//
// Solidity: function MINTING_PRICE() view returns(uint256)
func (_Nft *NftSession) MINTINGPRICE() (*big.Int, error) {
	return _Nft.Contract.MINTINGPRICE(&_Nft.CallOpts)
}

// MINTINGPRICE is a free data retrieval call binding the contract method 0x63d9d166.
//
// Solidity: function MINTING_PRICE() view returns(uint256)
func (_Nft *NftCallerSession) MINTINGPRICE() (*big.Int, error) {
	return _Nft.Contract.MINTINGPRICE(&_Nft.CallOpts)
}

// REGISTRATIONPRICE is a free data retrieval call binding the contract method 0x2663e0b5.
//
// Solidity: function REGISTRATION_PRICE() view returns(uint256)
func (_Nft *NftCaller) REGISTRATIONPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "REGISTRATION_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REGISTRATIONPRICE is a free data retrieval call binding the contract method 0x2663e0b5.
//
// Solidity: function REGISTRATION_PRICE() view returns(uint256)
func (_Nft *NftSession) REGISTRATIONPRICE() (*big.Int, error) {
	return _Nft.Contract.REGISTRATIONPRICE(&_Nft.CallOpts)
}

// REGISTRATIONPRICE is a free data retrieval call binding the contract method 0x2663e0b5.
//
// Solidity: function REGISTRATION_PRICE() view returns(uint256)
func (_Nft *NftCallerSession) REGISTRATIONPRICE() (*big.Int, error) {
	return _Nft.Contract.REGISTRATIONPRICE(&_Nft.CallOpts)
}

// TOTALSUPPLY is a free data retrieval call binding the contract method 0x902d55a5.
//
// Solidity: function TOTAL_SUPPLY() view returns(uint256)
func (_Nft *NftCaller) TOTALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "TOTAL_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALSUPPLY is a free data retrieval call binding the contract method 0x902d55a5.
//
// Solidity: function TOTAL_SUPPLY() view returns(uint256)
func (_Nft *NftSession) TOTALSUPPLY() (*big.Int, error) {
	return _Nft.Contract.TOTALSUPPLY(&_Nft.CallOpts)
}

// TOTALSUPPLY is a free data retrieval call binding the contract method 0x902d55a5.
//
// Solidity: function TOTAL_SUPPLY() view returns(uint256)
func (_Nft *NftCallerSession) TOTALSUPPLY() (*big.Int, error) {
	return _Nft.Contract.TOTALSUPPLY(&_Nft.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Nft *NftCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Nft *NftSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Nft.Contract.BalanceOf(&_Nft.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Nft *NftCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Nft.Contract.BalanceOf(&_Nft.CallOpts, owner)
}

// BlacklistedAddresses is a free data retrieval call binding the contract method 0xd5749d42.
//
// Solidity: function blacklistedAddresses(address ) view returns(bool)
func (_Nft *NftCaller) BlacklistedAddresses(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "blacklistedAddresses", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlacklistedAddresses is a free data retrieval call binding the contract method 0xd5749d42.
//
// Solidity: function blacklistedAddresses(address ) view returns(bool)
func (_Nft *NftSession) BlacklistedAddresses(arg0 common.Address) (bool, error) {
	return _Nft.Contract.BlacklistedAddresses(&_Nft.CallOpts, arg0)
}

// BlacklistedAddresses is a free data retrieval call binding the contract method 0xd5749d42.
//
// Solidity: function blacklistedAddresses(address ) view returns(bool)
func (_Nft *NftCallerSession) BlacklistedAddresses(arg0 common.Address) (bool, error) {
	return _Nft.Contract.BlacklistedAddresses(&_Nft.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Nft *NftCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Nft *NftSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Nft.Contract.GetApproved(&_Nft.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Nft *NftCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Nft.Contract.GetApproved(&_Nft.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Nft *NftCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Nft *NftSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Nft.Contract.IsApprovedForAll(&_Nft.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Nft *NftCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Nft.Contract.IsApprovedForAll(&_Nft.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nft *NftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nft *NftSession) Name() (string, error) {
	return _Nft.Contract.Name(&_Nft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nft *NftCallerSession) Name() (string, error) {
	return _Nft.Contract.Name(&_Nft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nft *NftCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nft *NftSession) Owner() (common.Address, error) {
	return _Nft.Contract.Owner(&_Nft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nft *NftCallerSession) Owner() (common.Address, error) {
	return _Nft.Contract.Owner(&_Nft.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Nft *NftCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Nft *NftSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Nft.Contract.OwnerOf(&_Nft.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Nft *NftCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Nft.Contract.OwnerOf(&_Nft.CallOpts, tokenId)
}

// Registered is a free data retrieval call binding the contract method 0xb2dd5c07.
//
// Solidity: function registered(address ) view returns(bool)
func (_Nft *NftCaller) Registered(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "registered", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Registered is a free data retrieval call binding the contract method 0xb2dd5c07.
//
// Solidity: function registered(address ) view returns(bool)
func (_Nft *NftSession) Registered(arg0 common.Address) (bool, error) {
	return _Nft.Contract.Registered(&_Nft.CallOpts, arg0)
}

// Registered is a free data retrieval call binding the contract method 0xb2dd5c07.
//
// Solidity: function registered(address ) view returns(bool)
func (_Nft *NftCallerSession) Registered(arg0 common.Address) (bool, error) {
	return _Nft.Contract.Registered(&_Nft.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Nft *NftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Nft *NftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Nft.Contract.SupportsInterface(&_Nft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Nft *NftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Nft.Contract.SupportsInterface(&_Nft.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nft *NftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nft *NftSession) Symbol() (string, error) {
	return _Nft.Contract.Symbol(&_Nft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nft *NftCallerSession) Symbol() (string, error) {
	return _Nft.Contract.Symbol(&_Nft.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Nft *NftCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Nft.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Nft *NftSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Nft.Contract.TokenURI(&_Nft.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Nft *NftCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Nft.Contract.TokenURI(&_Nft.CallOpts, tokenId)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x44337ea1.
//
// Solidity: function addToBlacklist(address user) returns()
func (_Nft *NftTransactor) AddToBlacklist(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "addToBlacklist", user)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x44337ea1.
//
// Solidity: function addToBlacklist(address user) returns()
func (_Nft *NftSession) AddToBlacklist(user common.Address) (*types.Transaction, error) {
	return _Nft.Contract.AddToBlacklist(&_Nft.TransactOpts, user)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x44337ea1.
//
// Solidity: function addToBlacklist(address user) returns()
func (_Nft *NftTransactorSession) AddToBlacklist(user common.Address) (*types.Transaction, error) {
	return _Nft.Contract.AddToBlacklist(&_Nft.TransactOpts, user)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Nft *NftTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Nft *NftSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.Approve(&_Nft.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Nft *NftTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.Approve(&_Nft.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 tokenId) payable returns()
func (_Nft *NftTransactor) Mint(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "mint", tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 tokenId) payable returns()
func (_Nft *NftSession) Mint(tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.Mint(&_Nft.TransactOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 tokenId) payable returns()
func (_Nft *NftTransactorSession) Mint(tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.Mint(&_Nft.TransactOpts, tokenId)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() payable returns()
func (_Nft *NftTransactor) Register(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "register")
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() payable returns()
func (_Nft *NftSession) Register() (*types.Transaction, error) {
	return _Nft.Contract.Register(&_Nft.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() payable returns()
func (_Nft *NftTransactorSession) Register() (*types.Transaction, error) {
	return _Nft.Contract.Register(&_Nft.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nft *NftTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nft *NftSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nft.Contract.RenounceOwnership(&_Nft.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nft *NftTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nft.Contract.RenounceOwnership(&_Nft.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Nft *NftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Nft *NftSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SafeTransferFrom(&_Nft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Nft *NftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.SafeTransferFrom(&_Nft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Nft *NftTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Nft *NftSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Nft.Contract.SafeTransferFrom0(&_Nft.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Nft *NftTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Nft.Contract.SafeTransferFrom0(&_Nft.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Nft *NftTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Nft *NftSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Nft.Contract.SetApprovalForAll(&_Nft.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Nft *NftTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Nft.Contract.SetApprovalForAll(&_Nft.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Nft *NftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Nft *NftSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.TransferFrom(&_Nft.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Nft *NftTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Nft.Contract.TransferFrom(&_Nft.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nft *NftTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Nft.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nft *NftSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nft.Contract.TransferOwnership(&_Nft.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nft *NftTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nft.Contract.TransferOwnership(&_Nft.TransactOpts, newOwner)
}

// NftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Nft contract.
type NftApprovalIterator struct {
	Event *NftApproval // Event containing the contract specifics and raw log

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
func (it *NftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftApproval)
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
		it.Event = new(NftApproval)
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
func (it *NftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftApproval represents a Approval event raised by the Nft contract.
type NftApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Nft *NftFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*NftApprovalIterator, error) {

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

	logs, sub, err := _Nft.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftApprovalIterator{contract: _Nft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Nft *NftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NftApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Nft.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftApproval)
				if err := _Nft.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Nft *NftFilterer) ParseApproval(log types.Log) (*NftApproval, error) {
	event := new(NftApproval)
	if err := _Nft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Nft contract.
type NftApprovalForAllIterator struct {
	Event *NftApprovalForAll // Event containing the contract specifics and raw log

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
func (it *NftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftApprovalForAll)
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
		it.Event = new(NftApprovalForAll)
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
func (it *NftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftApprovalForAll represents a ApprovalForAll event raised by the Nft contract.
type NftApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Nft *NftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*NftApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &NftApprovalForAllIterator{contract: _Nft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Nft *NftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *NftApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftApprovalForAll)
				if err := _Nft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Nft *NftFilterer) ParseApprovalForAll(log types.Log) (*NftApprovalForAll, error) {
	event := new(NftApprovalForAll)
	if err := _Nft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftBlacklistedIterator is returned from FilterBlacklisted and is used to iterate over the raw logs and unpacked data for Blacklisted events raised by the Nft contract.
type NftBlacklistedIterator struct {
	Event *NftBlacklisted // Event containing the contract specifics and raw log

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
func (it *NftBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftBlacklisted)
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
		it.Event = new(NftBlacklisted)
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
func (it *NftBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftBlacklisted represents a Blacklisted event raised by the Nft contract.
type NftBlacklisted struct {
	AddedBy            common.Address
	BlacklistedAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBlacklisted is a free log retrieval operation binding the contract event 0xd36871fdf6981136f3ac0564927005901eda06f7a9dff1e8b2a1d7846b8ebb50.
//
// Solidity: event Blacklisted(address indexed addedBy, address indexed blacklistedAddress)
func (_Nft *NftFilterer) FilterBlacklisted(opts *bind.FilterOpts, addedBy []common.Address, blacklistedAddress []common.Address) (*NftBlacklistedIterator, error) {

	var addedByRule []interface{}
	for _, addedByItem := range addedBy {
		addedByRule = append(addedByRule, addedByItem)
	}
	var blacklistedAddressRule []interface{}
	for _, blacklistedAddressItem := range blacklistedAddress {
		blacklistedAddressRule = append(blacklistedAddressRule, blacklistedAddressItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "Blacklisted", addedByRule, blacklistedAddressRule)
	if err != nil {
		return nil, err
	}
	return &NftBlacklistedIterator{contract: _Nft.contract, event: "Blacklisted", logs: logs, sub: sub}, nil
}

// WatchBlacklisted is a free log subscription operation binding the contract event 0xd36871fdf6981136f3ac0564927005901eda06f7a9dff1e8b2a1d7846b8ebb50.
//
// Solidity: event Blacklisted(address indexed addedBy, address indexed blacklistedAddress)
func (_Nft *NftFilterer) WatchBlacklisted(opts *bind.WatchOpts, sink chan<- *NftBlacklisted, addedBy []common.Address, blacklistedAddress []common.Address) (event.Subscription, error) {

	var addedByRule []interface{}
	for _, addedByItem := range addedBy {
		addedByRule = append(addedByRule, addedByItem)
	}
	var blacklistedAddressRule []interface{}
	for _, blacklistedAddressItem := range blacklistedAddress {
		blacklistedAddressRule = append(blacklistedAddressRule, blacklistedAddressItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "Blacklisted", addedByRule, blacklistedAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftBlacklisted)
				if err := _Nft.contract.UnpackLog(event, "Blacklisted", log); err != nil {
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

// ParseBlacklisted is a log parse operation binding the contract event 0xd36871fdf6981136f3ac0564927005901eda06f7a9dff1e8b2a1d7846b8ebb50.
//
// Solidity: event Blacklisted(address indexed addedBy, address indexed blacklistedAddress)
func (_Nft *NftFilterer) ParseBlacklisted(log types.Log) (*NftBlacklisted, error) {
	event := new(NftBlacklisted)
	if err := _Nft.contract.UnpackLog(event, "Blacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the Nft contract.
type NftMintedIterator struct {
	Event *NftMinted // Event containing the contract specifics and raw log

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
func (it *NftMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftMinted)
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
		it.Event = new(NftMinted)
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
func (it *NftMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftMinted represents a Minted event raised by the Nft contract.
type NftMinted struct {
	User    common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address indexed user, uint256 indexed tokenId)
func (_Nft *NftFilterer) FilterMinted(opts *bind.FilterOpts, user []common.Address, tokenId []*big.Int) (*NftMintedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "Minted", userRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftMintedIterator{contract: _Nft.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address indexed user, uint256 indexed tokenId)
func (_Nft *NftFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *NftMinted, user []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "Minted", userRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftMinted)
				if err := _Nft.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address indexed user, uint256 indexed tokenId)
func (_Nft *NftFilterer) ParseMinted(log types.Log) (*NftMinted, error) {
	event := new(NftMinted)
	if err := _Nft.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Nft contract.
type NftOwnershipTransferredIterator struct {
	Event *NftOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NftOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftOwnershipTransferred)
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
		it.Event = new(NftOwnershipTransferred)
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
func (it *NftOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftOwnershipTransferred represents a OwnershipTransferred event raised by the Nft contract.
type NftOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nft *NftFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NftOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NftOwnershipTransferredIterator{contract: _Nft.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nft *NftFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NftOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftOwnershipTransferred)
				if err := _Nft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Nft *NftFilterer) ParseOwnershipTransferred(log types.Log) (*NftOwnershipTransferred, error) {
	event := new(NftOwnershipTransferred)
	if err := _Nft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the Nft contract.
type NftRegisteredIterator struct {
	Event *NftRegistered // Event containing the contract specifics and raw log

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
func (it *NftRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftRegistered)
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
		it.Event = new(NftRegistered)
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
func (it *NftRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftRegistered represents a Registered event raised by the Nft contract.
type NftRegistered struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x2d3734a8e47ac8316e500ac231c90a6e1848ca2285f40d07eaa52005e4b3a0e9.
//
// Solidity: event Registered(address indexed user)
func (_Nft *NftFilterer) FilterRegistered(opts *bind.FilterOpts, user []common.Address) (*NftRegisteredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "Registered", userRule)
	if err != nil {
		return nil, err
	}
	return &NftRegisteredIterator{contract: _Nft.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x2d3734a8e47ac8316e500ac231c90a6e1848ca2285f40d07eaa52005e4b3a0e9.
//
// Solidity: event Registered(address indexed user)
func (_Nft *NftFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *NftRegistered, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "Registered", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftRegistered)
				if err := _Nft.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0x2d3734a8e47ac8316e500ac231c90a6e1848ca2285f40d07eaa52005e4b3a0e9.
//
// Solidity: event Registered(address indexed user)
func (_Nft *NftFilterer) ParseRegistered(log types.Log) (*NftRegistered, error) {
	event := new(NftRegistered)
	if err := _Nft.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Nft contract.
type NftTransferIterator struct {
	Event *NftTransfer // Event containing the contract specifics and raw log

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
func (it *NftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftTransfer)
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
		it.Event = new(NftTransfer)
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
func (it *NftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftTransfer represents a Transfer event raised by the Nft contract.
type NftTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Nft *NftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*NftTransferIterator, error) {

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

	logs, sub, err := _Nft.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NftTransferIterator{contract: _Nft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Nft *NftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NftTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Nft.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftTransfer)
				if err := _Nft.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Nft *NftFilterer) ParseTransfer(log types.Log) (*NftTransfer, error) {
	event := new(NftTransfer)
	if err := _Nft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
