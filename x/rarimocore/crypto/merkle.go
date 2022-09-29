package crypto

import (
	"bytes"

	"github.com/ethereum/go-ethereum/crypto"
	merkle "gitlab.com/rarify-protocol/go-merkle"
)

// ContentData contains specific information for the certain content
type ContentData []byte

func NewEmptyContent() ContentData {
	return []byte{}
}

// HashContent implements the Content interface provided by go-merkle and represents the content stored in the tree.
type HashContent struct {
	// Hash of the deposit tx
	TxHash         string
	OperationId    string
	CurrentNetwork string

	TargetNetwork string
	// Receiver address on target network
	Receiver []byte
	// Target bridge contract
	TargetContract []byte
	// Can contain any specific data for target chain to validate.
	Data ContentData
}

func (c HashContent) OriginHash() []byte {
	return crypto.Keccak256([]byte(c.TxHash), []byte(c.OperationId), []byte(c.CurrentNetwork))
}

var _ merkle.Content = HashContent{}

func (c HashContent) CalculateHash() []byte {
	return crypto.Keccak256(c.OriginHash(), []byte(c.TargetNetwork), c.Receiver, c.TargetContract, c.Data)
}

//Equals tests for equality of two Contents
func (c HashContent) Equals(other merkle.Content) bool {
	if oc, ok := other.(HashContent); ok {
		return bytes.Equal(oc.CalculateHash(), c.CalculateHash())
	}
	return false
}
