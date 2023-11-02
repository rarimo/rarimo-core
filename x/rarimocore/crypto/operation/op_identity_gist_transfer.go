package operation

import (
	"bytes"

	eth "github.com/ethereum/go-ethereum/crypto"
	merkle "github.com/rarimo/go-merkle"
)

// IdentityGISTTransferContent implements the Content interface provided by go-merkle and represents the content stored in the tree.
type IdentityGISTTransferContent struct {
	Contract               []byte
	GISTHash               []byte
	GISTReplacedBy         []byte
	GISTCreatedAtTimestamp []byte
	GISTCreatedAtBlock     []byte
	ReplacedGISTHash       []byte
}

var _ merkle.Content = IdentityGISTTransferContent{}

func (c IdentityGISTTransferContent) CalculateHash() []byte {
	return eth.Keccak256(
		c.Contract,
		c.GISTHash,
		c.GISTReplacedBy,
		c.GISTCreatedAtTimestamp,
		c.GISTCreatedAtBlock,
		c.ReplacedGISTHash,
	)
}

// Equals tests for equality of two Contents
func (c IdentityGISTTransferContent) Equals(other merkle.Content) bool {
	return bytes.Equal(other.CalculateHash(), c.CalculateHash())
}
