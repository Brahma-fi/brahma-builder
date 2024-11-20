package entity

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const (
	BaseChainID           = 8453
	SafeMultiSendCallOnly = "0x40a2accbd92bca938b02010e17a5b8929b49130d"
)

var (
	SafeMultiSendABI, _ = abi.JSON(strings.NewReader(`
		[
		  {
			"inputs": [
			  {
				"internalType": "bytes",
				"name": "transactions",
				"type": "bytes"
			  }
			],
			"name": "multiSend",
			"outputs": [],
			"stateMutability": "payable",
			"type": "function"
		  }
		]
	`))
)
