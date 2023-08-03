package operation

import (
	"bytes"

	eth "github.com/ethereum/go-ethereum/crypto"
	merkle "gitlab.com/rarimo/go-merkle"
)

// IdentityAggregatedTransferContent implements the Content interface provided by go-merkle and represents the content stored in the tree.
type IdentityAggregatedTransferContent struct {
	Contract      []byte
	Chain         string
	GISTHash      []byte
	StateRootHash []byte
	Timestamp     []byte
}

var _ merkle.Content = IdentityAggregatedTransferContent{}

func (c IdentityAggregatedTransferContent) CalculateHash() []byte {
	return eth.Keccak256(
		c.GISTHash,
		c.Timestamp,
		c.StateRootHash,
		c.Contract,
		[]byte(c.Chain),
	)
}

// Equals tests for equality of two Contents
func (c IdentityAggregatedTransferContent) Equals(other merkle.Content) bool {
	if oc, ok := other.(IdentityAggregatedTransferContent); ok {
		return bytes.Equal(oc.CalculateHash(), c.CalculateHash())
	}

	return false
}
