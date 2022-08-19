package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// InfoKeyPrefix is the prefix to retrieve all Info
	InfoKeyPrefix = "Info/value/"
)

// InfoKey returns the store key to retrieve a Info from the index fields
func InfoKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
