package types

import (
	"encoding/binary"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ binary.ByteOrder

const (
	// ConfirmationKeyPrefix is the prefix to retrieve all Confirmation
	ConfirmationKeyPrefix = "Confirmation/value/"
)

// ConfirmationKey returns the store key to retrieve a Confirmation from the index fields
func ConfirmationKey(
	root string,
) []byte {
	return append(hexutil.MustDecode(root), []byte("/")...)
}
