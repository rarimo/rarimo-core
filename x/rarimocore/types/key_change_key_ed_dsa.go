package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ChangeKeyEdDSAKeyPrefix is the prefix to retrieve all ChangeKeyEdDSA
	ChangeKeyEdDSAKeyPrefix = "ChangeKeyEdDSA/value/"
)

// ChangeKeyEdDSAKey returns the store key to retrieve a ChangeKeyEdDSA from the index fields
func ChangeKeyEdDSAKey(
	newKey string,
) []byte {
	var key []byte

	newKeyBytes := []byte(newKey)
	key = append(key, newKeyBytes...)
	key = append(key, []byte("/")...)

	return key
}
