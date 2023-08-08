package operation

import (
	"bytes"

	eth "github.com/ethereum/go-ethereum/crypto"
	merkle "gitlab.com/rarimo/go-merkle"
)

// IdentityDefaultTransferContent implements the Content interface provided by go-merkle and represents the content stored in the tree.
type IdentityDefaultTransferContent struct {
	Contract                []byte
	GISTHash                []byte
	Id                      []byte
	StateHash               []byte
	StateCreatedAtTimestamp []byte
	StateCreatedAtBlock     []byte
	StateReplacedBy         []byte
	GISTReplacedBy          []byte
	GISTCreatedAtTimestamp  []byte
	GISTCreatedAtBlock      []byte
	ReplacedStateHash       []byte
	ReplacedGISTHash        []byte
}

var _ merkle.Content = IdentityDefaultTransferContent{}

func (c IdentityDefaultTransferContent) CalculateHash() []byte {
	return eth.Keccak256(
		c.Contract,
		c.Id,
		c.StateHash,
		c.StateReplacedBy,
		c.StateCreatedAtTimestamp,
		c.StateCreatedAtBlock,
		c.GISTHash,
		c.GISTReplacedBy,
		c.GISTCreatedAtTimestamp,
		c.GISTCreatedAtBlock,
		c.ReplacedStateHash,
		c.ReplacedGISTHash,
	)
}

// Equals tests for equality of two Contents
func (c IdentityDefaultTransferContent) Equals(other merkle.Content) bool {
	return bytes.Equal(other.CalculateHash(), c.CalculateHash())
}
