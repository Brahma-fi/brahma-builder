package entity

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	ChainIDMainnet = 1
	ChainIDBase    = 8453
)

type Transaction struct {
	Target common.Address `json:"to"`
	Val    *big.Int       `json:"value"`
	Data   string         `json:"data"`
}

func (t *Transaction) From() common.Address {
	return common.HexToAddress("")
}

func (t *Transaction) To() common.Address {
	return t.Target
}
func (t *Transaction) CallData() string {
	return t.Data
}
func (t *Transaction) Value() *big.Int {
	return t.Val
}
func (t *Transaction) Operation() uint8 {
	// Call Only
	return 0
}

var (
	ZeroAddress  = common.HexToAddress("0x0000000000000000000000000000000000000000")
	ZeroAddressE = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")
)

// ToZeroAddress converts types.ZeroAddressE to types.ZeroAddress address
// else returns the input address
func ToZeroAddress(address common.Address) common.Address {
	if address == ZeroAddressE {
		return ZeroAddress
	}
	return address
}
