package operation

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common/hexutil"
	eth "github.com/ethereum/go-ethereum/crypto"
	merkle "gitlab.com/rarify-protocol/go-merkle"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

// ChangePartiesContent implements the Content interface provided by go-merkle and represents the content stored in the tree.
type ChangePartiesContent struct {
	NewSet    []*types.Party
	NewKey    string
	Signature string
}

var _ merkle.Content = ChangePartiesContent{}

func (c ChangePartiesContent) CalculateHash() []byte {
	return eth.Keccak256(crypto.GetPartiesHash(c.NewSet), hexutil.MustDecode(c.NewKey), hexutil.MustDecode(c.Signature))
}

// Equals tests for equality of two Contents
func (c ChangePartiesContent) Equals(other merkle.Content) bool {
	if oc, ok := other.(ChangePartiesContent); ok {
		return bytes.Equal(oc.CalculateHash(), c.CalculateHash())
	}

	return false
}
