package operation

import (
	"bytes"

	"github.com/ethereum/go-ethereum/crypto"
	merkle "gitlab.com/rarimo/go-merkle"
)

// ChangeKeyContent implements the Content interface provided by go-merkle and represents the content stored in the tree.
type ChangeKeyContent struct {
	Key []byte
}

var _ merkle.Content = ChangeKeyContent{}

func (c ChangeKeyContent) CalculateHash() []byte {
	return crypto.Keccak256(c.Key)
}

// Equals tests for equality of two Contents
func (c ChangeKeyContent) Equals(other merkle.Content) bool {
	if oc, ok := other.(ChangeKeyContent); ok {
		return bytes.Equal(oc.CalculateHash(), c.CalculateHash())
	}

	if oc, ok := other.(*ChangeKeyContent); ok {
		return bytes.Equal(oc.CalculateHash(), c.CalculateHash())
	}

	return false
}
