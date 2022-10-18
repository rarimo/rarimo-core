package types

import (
	"encoding/binary"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ binary.ByteOrder

const (
	// OperationKeyPrefix is the prefix to retrieve all Operation
	OperationKeyPrefix = "Operation/value/"
)

// OperationKey returns the store key to retrieve an Operation from the index fields
func OperationKey(
	index string,
) []byte {
	return append(hexutil.MustDecode(index), []byte("/")...)
}
