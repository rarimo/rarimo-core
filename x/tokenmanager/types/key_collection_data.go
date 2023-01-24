package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	// CollectionDataKeyPrefix is the prefix to retrieve all CollectionData
	CollectionDataKeyPrefix = "CollectionData/value/"
)

func CollectionDataKey(index *CollectionDataIndex) []byte {
	var key []byte

	key = append(key, []byte(index.Chain)...)
	key = append(key, []byte("/")...)

	key = append(key, []byte(index.Address)...)
	key = append(key, []byte("/")...)

	return key
}
