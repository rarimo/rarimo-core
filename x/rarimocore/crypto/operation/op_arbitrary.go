package operation

import (
	"bytes"

	eth "github.com/ethereum/go-ethereum/crypto"
	merkle "github.com/rarimo/go-merkle"
)

// ArbitraryContent implements the Content interface provided by go-merkle and represents the content stored in the tree.
type ArbitraryContent struct {
	Data []byte
}

var _ merkle.Content = ArbitraryContent{}

func (c ArbitraryContent) CalculateHash() []byte {
	return eth.Keccak256(c.Data)
}

// Equals tests for equality of two Contents
func (c ArbitraryContent) Equals(other merkle.Content) bool {
	return bytes.Equal(other.CalculateHash(), c.CalculateHash())
}
