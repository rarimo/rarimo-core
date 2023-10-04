package operation

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common/hexutil"
	eth "github.com/ethereum/go-ethereum/crypto"
	merkle "github.com/rarimo/go-merkle"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto"
	"github.com/rarimo/rarimo-core/x/rarimocore/types"
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
	return bytes.Equal(other.CalculateHash(), c.CalculateHash())
}
