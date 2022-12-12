package operation

import (
	"bytes"

	"github.com/ethereum/go-ethereum/crypto"
	merkle "gitlab.com/rarimo/go-merkle"
)

// ContentData contains specific information for the certain content
type ContentData []byte

func NewEmptyContent() ContentData {
	return []byte{}
}

// OriginData contains specific information about network to be transferred from
type OriginData [32]byte

func NewEmptyOrigin() OriginData {
	return OriginData{}
}

// BundleData contains specific information about bundle operation
type BundleData []byte

func NewEmptyBundle() BundleData {
	return BundleData{}
}

// TransferContent implements the Content interface provided by go-merkle and represents the content stored in the tree.
type TransferContent struct {
	// Hash of the deposit tx info
	Origin        OriginData
	TargetNetwork string
	// Receiver address on target network
	Receiver []byte
	// Target bridge contract
	TargetContract []byte
	// Can contain any specific data for target chain to validate.
	Data ContentData
	// Bundle calls data
	Bundle BundleData
}

var _ merkle.Content = TransferContent{}

func (c TransferContent) CalculateHash() []byte {
	return crypto.Keccak256(c.Data, c.Bundle, c.Origin[:], []byte(c.TargetNetwork), c.Receiver, c.TargetContract)
}

// Equals tests for equality of two Contents
func (c TransferContent) Equals(other merkle.Content) bool {
	if oc, ok := other.(TransferContent); ok {
		return bytes.Equal(oc.CalculateHash(), c.CalculateHash())
	}

	if oc, ok := other.(*TransferContent); ok {
		return bytes.Equal(oc.CalculateHash(), c.CalculateHash())
	}

	return false
}
