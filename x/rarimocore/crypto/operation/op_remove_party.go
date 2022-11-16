package operation

import (
	"bytes"
	"math/big"

	eth "github.com/ethereum/go-ethereum/crypto"
	merkle "gitlab.com/rarify-protocol/go-merkle"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

// RemovePartyContent implements the Content interface provided by go-merkle and represents the content stored in the tree.
type RemovePartyContent struct {
	Set   []*types.Party
	Index uint32
}

var _ merkle.Content = RemovePartyContent{}

func (c RemovePartyContent) CalculateHash() []byte {
	return eth.Keccak256(GetPartiesHash(c.Set), big.NewInt(int64(c.Index)).Bytes())
}

// Equals tests for equality of two Contents
func (c RemovePartyContent) Equals(other merkle.Content) bool {
	if oc, ok := other.(RemovePartyContent); ok {
		return bytes.Equal(oc.CalculateHash(), c.CalculateHash())
	}

	if oc, ok := other.(*RemovePartyContent); ok {
		return bytes.Equal(oc.CalculateHash(), c.CalculateHash())
	}

	return false
}

func GetPartiesHash(set []*types.Party) []byte {
	var data []byte
	for _, p := range set {
		data = append(data, []byte(p.PubKey)...)
		data = append(data, []byte(p.Address)...)
		data = append(data, []byte(p.Account)...)
	}
	return eth.Keccak256(data)
}
