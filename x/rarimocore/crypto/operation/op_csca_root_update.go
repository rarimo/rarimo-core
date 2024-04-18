package operation

import (
	"bytes"

	eth "github.com/ethereum/go-ethereum/crypto"
	merkle "github.com/rarimo/go-merkle"
)

type CSCARootUpdateContent struct {
	Root      []byte
	Timestamp []byte
}

const hashPrefix = "Rarimo CSCA root"

var _ merkle.Content = CSCARootUpdateContent{}

func (c CSCARootUpdateContent) CalculateHash() []byte {
	return eth.Keccak256([]byte(hashPrefix), c.Root, c.Timestamp)
}

func (c CSCARootUpdateContent) Equals(other merkle.Content) bool {
	return bytes.Equal(other.CalculateHash(), c.CalculateHash())
}
