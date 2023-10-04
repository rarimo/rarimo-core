package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
)

func (s StateInfo) CalculateHash() []byte {
	return crypto.Keccak256(
		operation.To32Bytes(hexutil.MustDecode(s.Index)),
		operation.To32Bytes(hexutil.MustDecode(s.Hash)),
		operation.To32Bytes(new(big.Int).SetUint64(s.CreatedAtTimestamp).Bytes()),
	)
}
