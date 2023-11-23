package operation

import (
	"bytes"

	eth "github.com/ethereum/go-ethereum/crypto"
	merkle "github.com/rarimo/go-merkle"
)

// IdentityStateTransferContent implements the Content interface provided by go-merkle and represents the content stored in the tree.
type IdentityStateTransferContent struct {
	Contract                []byte
	Id                      []byte
	StateHash               []byte
	StateCreatedAtTimestamp []byte
	StateCreatedAtBlock     []byte
	ReplacedStateHash       []byte
}

var _ merkle.Content = IdentityStateTransferContent{}

func (c IdentityStateTransferContent) CalculateHash() []byte {
	return eth.Keccak256(
		c.Contract,
		c.Id,
		c.StateHash,
		c.StateCreatedAtTimestamp,
		c.StateCreatedAtBlock,
		c.ReplacedStateHash,
	)
}

// Equals tests for equality of two Contents
func (c IdentityStateTransferContent) Equals(other merkle.Content) bool {
	return bytes.Equal(other.CalculateHash(), c.CalculateHash())
}
