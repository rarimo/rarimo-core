package types

import (
	"encoding/binary"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ binary.ByteOrder

const (
	// HashKeyPrefix is the prefix to retrieve all Hashes
	HashKeyPrefix = "HashKey/value/"
)

// HashKey returns the store key to retrieve a  from the index fields
func HashKey(
	root string,
) []byte {
	return append(hexutil.MustDecode(root), []byte("/")...)
}
