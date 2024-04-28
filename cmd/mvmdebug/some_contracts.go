package main

import (
	"encoding/hex"
	"fmt"
	"github.com/a1146910248/mixchain/mvm/common"
	"github.com/a1146910248/mixchain/mvm/state"
)

var normalAddress, _ = hex.DecodeString("123456abc")
var hellWorldcontractAddress, _ = hex.DecodeString("12321214314")
var baseContractAddress, _ = hex.DecodeString("dc6ab9f81e16baa35c697e817b79f2284a978b99")
var normalAccount = common.BytesToAddress(normalAddress)
var helloWorldcontactAccont = common.BytesToAddress(hellWorldcontractAddress)
var baseContractAccont = common.BytesToAddress(baseContractAddress)

// 基本账户字节码
var baseCodeStr = "6080604052348015600e575f80fd5b506101f28061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c80632b225f29146100385780638da5cb5b14610056575b5f80fd5b610040610074565b60405161004d9190610144565b60405180910390f35b61005e6100b1565b60405161006b91906101a3565b60405180910390f35b60606040518060400160405280601081526020017f42617365436f6e747261637456302e3100000000000000000000000000000000815250905090565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610116826100d4565b61012081856100de565b93506101308185602086016100ee565b610139816100fc565b840191505092915050565b5f6020820190508181035f83015261015c818461010c565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61018d82610164565b9050919050565b61019d81610183565b82525050565b5f6020820190506101b65f830184610194565b9291505056fea2646970667358221220ddd67cfcb983d0de85f57df6f1c4eba145dc7529e6a6fffa2befb1eaa71bfb0664736f6c63430008190033"

// hellworld 账户字节码
var hellCodeStr = "608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b5780636057361d14610059575b600080fd5b610043610075565b60405161005091906100a1565b60405180910390f35b610073600480360381019061006e91906100ed565b61007e565b005b60008054905090565b8060008190555050565b6000819050919050565b61009b81610088565b82525050565b60006020820190506100b66000830184610092565b92915050565b600080fd5b6100ca81610088565b81146100d557600080fd5b50565b6000813590506100e7816100c1565b92915050565b600060208284031215610103576101026100bc565b5b6000610111848285016100d8565b9150509291505056fea264697066735822122051c67b4bae92a1ce8f1aaf8ed71e1f3a889144a9551cc31ae452fbc75e33440664736f6c63430008190033"

var helloCode, _ = hex.DecodeString(hellCodeStr)
var baseCode, _ = hex.DecodeString(baseCodeStr)

func updateContract() {
	// 加载账户State
	stateDb, err := state.TryLoadFromDisk()
	if err != nil {
		panic(err)
	}
	stateDb.SetCode(helloWorldcontactAccont, helloCode)
	stateDb.SetCode(baseContractAccont, baseCode)
	fmt.Println(stateDb.Commit())
}

var baseContractABIJson = `[
	{
		"inputs": [],
		"name": "CurrentVersion",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "owner",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	}
]`
var storeContractABIJson = `[
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "num",
				"type": "uint256"
			}
		],
		"name": "store",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "retrieve",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	}
]`
var hellWorldContractABIJson = `[
	{
		"inputs": [],
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "address",
				"name": "",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"name": "Triggle",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "contractAddr",
				"type": "address"
			}
		],
		"name": "getVersion",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "getbalance",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "onlytest",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "tmp",
				"type": "uint256"
			}
		],
		"name": "setBalance",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`
