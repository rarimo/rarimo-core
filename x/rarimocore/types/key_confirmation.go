package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ConfirmationKeyPrefix is the prefix to retrieve all Confirmation
	ConfirmationKeyPrefix = "Confirmation/value/"
)

// ConfirmationKey returns the store key to retrieve a Confirmation from the index fields
func ConfirmationKey(
	height string,
) []byte {
	var key []byte

	heightBytes := []byte(height)
	key = append(key, heightBytes...)
	key = append(key, []byte("/")...)

	return key
}
