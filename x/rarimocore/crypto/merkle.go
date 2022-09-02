package crypto

import (
	"bytes"

	"github.com/ethereum/go-ethereum/crypto"
	merkle "gitlab.com/rarify-protocol/go-merkle"
)

var _ merkle.Content = HashContent{}

// HashContent implements the Content interface provided by go-merkle and represents the content stored in the tree.
type HashContent struct {
	// Hash of the deposit tx
	TxHash         string
	CurrentNetwork string
	EventId        string

	// Collection address on target chain
	TargetAddress []byte
	// TokenId on target chain
	TargetId []byte
	// Receiver address on target network
	Receiver      []byte
	TargetNetwork string
	// Memory representation of amount integer as a byte array in big-endian (with leading zeros if needed)
	// Use binary.BigEndian.PutUint64(amount, c.Amount)
	Amount    []byte
	ProgramId []byte
}

func (c HashContent) OriginHash() []byte {
	return crypto.Keccak256([]byte(c.TxHash), []byte(c.CurrentNetwork), []byte(c.EventId))
}

func (c HashContent) CalculateHash() []byte {
	return crypto.Keccak256(c.TargetAddress, c.TargetId, c.Amount, c.Receiver, crypto.Keccak256([]byte(c.TxHash), []byte(c.CurrentNetwork), []byte(c.EventId)), []byte(c.TargetNetwork), c.ProgramId)
}

//Equals tests for equality of two Contents
func (c HashContent) Equals(other merkle.Content) bool {
	if oc, ok := other.(HashContent); ok {
		return bytes.Equal(oc.CalculateHash(), c.CalculateHash())
	}
	return false
}
