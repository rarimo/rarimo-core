package operation

import (
	"bytes"

	eth "github.com/ethereum/go-ethereum/crypto"
	merkle "github.com/rarimo/go-merkle"
)

type RootUpdateContent struct {
	ContractAddress []byte
	Root            []byte
}

const hashRootUpdatePrefix = "Rarimo root"

var _ merkle.Content = RootUpdateContent{}

func (c RootUpdateContent) CalculateHash() []byte {
	return eth.Keccak256([]byte(hashRootUpdatePrefix), c.ContractAddress, c.Root)
}

func (c RootUpdateContent) Equals(other merkle.Content) bool {
	return bytes.Equal(other.CalculateHash(), c.CalculateHash())
}
