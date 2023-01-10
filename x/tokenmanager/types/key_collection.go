package types

import (
	"encoding/binary"
	"fmt"
)

var _ binary.ByteOrder

const (
	// CollectionKeyPrefix is the prefix to retrieve all Collection
	CollectionKeyPrefix = "Collection/value/"
	// CollectionChainParamsKeyPrefix is the prefix to save collection chain params (index, decimals)
	CollectionChainParamsKeyPrefix = "Collection/index/"
)

func CollectionKey(index string) []byte {
	return append([]byte(index), '/')
}

// address is hex-encoded (should be as it might have been done hexutil package)
func CollectionIndexKey(network, address string) []byte {
	return append([]byte(fmt.Sprintf("%s:%s", network, address)), '/')
}
