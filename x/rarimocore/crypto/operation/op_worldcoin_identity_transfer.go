package operation

import (
	"bytes"

	eth "github.com/ethereum/go-ethereum/crypto"
	merkle "github.com/rarimo/go-merkle"
)

// WorldCoinIdentityTransferContent implements the Content interface provided by go-merkle and represents the content stored in the tree.
type WorldCoinIdentityTransferContent struct {
	Contract []byte
	// Previous (replaced) identity state
	PrevState []byte
	// Worldcoin identity state
	State []byte
	// Worldcoin timestamp
	Timestamp []byte
}

var _ merkle.Content = WorldCoinIdentityTransferContent{}

func (c WorldCoinIdentityTransferContent) CalculateHash() []byte {
	return eth.Keccak256(c.Contract, c.PrevState, c.State, c.Timestamp)
}

// Equals tests for equality of two Contents
func (c WorldCoinIdentityTransferContent) Equals(other merkle.Content) bool {
	return bytes.Equal(other.CalculateHash(), c.CalculateHash())
}
