package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ItemKeyPrefix is the prefix to retrieve all Item
	ItemKeyPrefix = "Item/value/"
)

// ItemKey returns the store key to retrieve a Item from the index fields
func ItemKey(
	tokenAddress []byte,
	tokenId []byte,
) []byte {
	var key []byte

	key = append(key, tokenAddress...)
	key = append(key, []byte("/")...)

	key = append(key, tokenId...)
	key = append(key, []byte("/")...)

	return key
}
