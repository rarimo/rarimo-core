package operation

import (
	"bytes"

	eth "github.com/ethereum/go-ethereum/crypto"
	merkle "gitlab.com/rarimo/go-merkle"
)

// IdentityDefaultTransferContent implements the Content interface provided by go-merkle and represents the content stored in the tree.
type IdentityDefaultTransferContent struct {
	Contract                 []byte
	GISTHash                 []byte
	Id                       []byte
	StateHash                []byte
	StateCreatedAtTimestamp  []byte
	StateCreatedAtBlock      []byte
	StateReplacedAtTimestamp []byte
	StateReplacedAtBlock     []byte
	StateReplacedBy          []byte
	GISTReplacedBy           []byte
	GISTCreatedAtTimestamp   []byte
	GISTCreatedAtBlock       []byte
	GISTReplacedAtTimestamp  []byte
	GISTReplacedAtBlock      []byte
	ReplacedStateHash        []byte
}

var _ merkle.Content = IdentityDefaultTransferContent{}

func (c IdentityDefaultTransferContent) CalculateHash() []byte {
	return eth.Keccak256(
		c.Contract,
		c.Id,
		c.StateHash,
		c.StateReplacedBy,
		c.StateCreatedAtTimestamp,
		c.StateReplacedAtTimestamp,
		c.StateCreatedAtBlock,
		c.StateReplacedAtBlock,
		c.GISTHash,
		c.GISTReplacedBy,
		c.GISTCreatedAtTimestamp,
		c.GISTReplacedAtTimestamp,
		c.GISTCreatedAtBlock,
		c.GISTReplacedAtBlock,
		c.ReplacedStateHash,
	)
}

// Equals tests for equality of two Contents
func (c IdentityDefaultTransferContent) Equals(other merkle.Content) bool {
	if oc, ok := other.(IdentityDefaultTransferContent); ok {
		return bytes.Equal(oc.CalculateHash(), c.CalculateHash())
	}

	return false
}
