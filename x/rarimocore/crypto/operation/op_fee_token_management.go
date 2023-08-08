package operation

import (
	"bytes"

	eth "github.com/ethereum/go-ethereum/crypto"
	merkle "gitlab.com/rarimo/go-merkle"
)

// FeeTokenManagementContent implements the Content interface provided by go-merkle and represents the content stored in the tree.
type FeeTokenManagementContent struct {
	Nonce         []byte
	TargetNetwork string
	// Receiver address on target network
	Receiver []byte
	// Target bridge contract
	TargetContract []byte
	// Can contain any specific data for target chain to validate.
	Data ContentData
}

var _ merkle.Content = FeeTokenManagementContent{}

func (f FeeTokenManagementContent) CalculateHash() []byte {
	return eth.Keccak256(f.Nonce, f.Receiver, f.TargetContract, []byte(f.TargetNetwork), f.Data)
}

// Equals tests for equality of two Contents
func (f FeeTokenManagementContent) Equals(other merkle.Content) bool {
	return bytes.Equal(other.CalculateHash(), c.CalculateHash())
}
