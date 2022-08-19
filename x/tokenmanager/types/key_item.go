package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ItemKeyPrefix is the prefix to retrieve all Item
	ItemKeyPrefix = "Item/value/"
)

// ItemKey returns the store key to retrieve a Item from the index fields
func ItemKey(
	tokenAddress string,
	tokenId string,
) []byte {
	var key []byte

	tokenAddressBytes := []byte(tokenAddress)
	key = append(key, tokenAddressBytes...)
	key = append(key, []byte("/")...)

	tokenIdBytes := []byte(tokenId)
	key = append(key, tokenIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
