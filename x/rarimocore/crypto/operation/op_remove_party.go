package operation

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common/hexutil"
	eth "github.com/ethereum/go-ethereum/crypto"
	merkle "gitlab.com/rarify-protocol/go-merkle"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

// ChangePartiesContent implements the Content interface provided by go-merkle and represents the content stored in the tree.
type ChangePartiesContent struct {
	NewSet    []*types.Party
	Signature string
}

var _ merkle.Content = ChangePartiesContent{}

func (c ChangePartiesContent) CalculateHash() []byte {
	return eth.Keccak256(GetPartiesHash(c.NewSet), hexutil.MustDecode(c.Signature))
}

// Equals tests for equality of two Contents
func (c ChangePartiesContent) Equals(other merkle.Content) bool {
	if oc, ok := other.(ChangePartiesContent); ok {
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
