package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ChangeKeyECDSAKeyPrefix is the prefix to retrieve all ChangeKeyECDSA
	ChangeKeyECDSAKeyPrefix = "ChangeKeyECDSA/value/"
)

// ChangeKeyECDSAKey returns the store key to retrieve a ChangeKeyECDSA from the index fields
func ChangeKeyECDSAKey(
	newKey string,
) []byte {
	var key []byte

	newKeyBytes := []byte(newKey)
	key = append(key, newKeyBytes...)
	key = append(key, []byte("/")...)

	return key
}
