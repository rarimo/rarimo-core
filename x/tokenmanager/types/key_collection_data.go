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

func (i *CollectionDataIndex) Equal(o *CollectionDataIndex) bool {
	if o == nil || i == nil {
		return i == o
	}

	return i.Chain == o.Chain && i.Address == o.Address
}
