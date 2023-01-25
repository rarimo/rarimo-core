package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	// ItemKeyPrefix is the prefix to retrieve all Item
	ItemKeyPrefix        = "Item/value/"
	OnChainItemKeyPrefix = "OnChainItem/value/"
)

// ItemKey returns the store key to retrieve an Item from the index fields
func ItemKey(
	index string,
) []byte {
	var key []byte

	key = append(key, []byte(index)...)
	key = append(key, []byte("/")...)

	return key
}

func OnChainItemKey(index *OnChainItemIndex) []byte {
	var key []byte

	key = append(key, []byte(index.Chain)...)
	key = append(key, []byte("/")...)

	key = append(key, []byte(index.Address)...)
	key = append(key, []byte("/")...)

	key = append(key, []byte(index.TokenID)...)
	key = append(key, []byte("/")...)

	return key
}
