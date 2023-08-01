package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
)

func (s StateInfo) CalculateHash() []byte {
	return crypto.Keccak256(
		hexutil.MustDecode(s.Index),
		hexutil.MustDecode(s.Hash),
		operation.To32Bytes(new(big.Int).SetUint64(s.CreatedAtTimestamp).Bytes()),
	)
}
