package operation

import (
	"bytes"

	eth "github.com/ethereum/go-ethereum/crypto"
	merkle "github.com/rarimo/go-merkle"
)

type PassportRootUpdateContent struct {
	ContractAddress []byte
	Root            []byte
	RootTimestamp   []byte
}

const hashPassportRootUpdatePrefix = "Rarimo passport root"

var _ merkle.Content = PassportRootUpdateContent{}

func (c PassportRootUpdateContent) CalculateHash() []byte {
	return eth.Keccak256([]byte(hashPassportRootUpdatePrefix), c.ContractAddress, c.Root, c.RootTimestamp)
}

func (c PassportRootUpdateContent) Equals(other merkle.Content) bool {
	return bytes.Equal(other.CalculateHash(), c.CalculateHash())
}
