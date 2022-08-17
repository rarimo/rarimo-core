package crypto

import (
	"crypto/sha256"

	"github.com/cbergoon/merkletree"
)

var _ merkletree.Content = HashContent{}

// HashContent implements the Content interface provided by merkletree and represents the content stored in the tree.
type HashContent struct {
	// Hash of the deposit tx
	TxHash string
	// Collection address on current chain
	TokenAddress string
	// TokenId
	TokenId string
	// Receiver address on target network
	Receiver      string
	TargetNetwork string
}

func (c HashContent) String() string {
	return c.TxHash + c.TokenAddress + c.TokenId + c.Receiver + c.TargetNetwork
}

func (c HashContent) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(c.String())); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

//Equals tests for equality of two Contents
func (c HashContent) Equals(other merkletree.Content) (bool, error) {
	if oc, ok := other.(HashContent); ok {
		return c.TxHash == oc.TxHash &&
			c.TokenAddress == oc.TokenAddress &&
			c.TokenId == oc.TokenId &&
			c.Receiver == oc.Receiver &&
			c.TokenAddress == oc.TargetNetwork, nil
	}
	return false, nil
}
