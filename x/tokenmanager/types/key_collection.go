package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	// CollectionKeyPrefix is the prefix to retrieve all Collection
	CollectionKeyPrefix = "Collection/value/"
)

func CollectionKey(index string) []byte {
	return append([]byte(index), '/')
}
