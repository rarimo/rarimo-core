package types

import (
	"encoding/binary"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ binary.ByteOrder

const (
	// ChangeKeyECDSAKeyPrefix is the prefix to retrieve all ChangeKeyECDSA
	ChangeKeyECDSAKeyPrefix = "ChangeKeyECDSA/value/"
)

// ChangeKeyECDSAKey returns the store key to retrieve a ChangeKeyECDSA from the index fields
func ChangeKeyECDSAKey(
	newKey string,
) []byte {
	return append(hexutil.MustDecode(newKey), []byte("/")...)
}
