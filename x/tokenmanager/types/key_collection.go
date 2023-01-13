package types

import (
	"encoding/binary"
	"fmt"
)

var _ binary.ByteOrder

const (
	// CollectionKeyPrefix is the prefix to retrieve all Collection
	CollectionKeyPrefix  = "Collection/value/"
	CollectionAddrPrefix = "Collection/value/addr/"
)

func CollectionKey(index string) []byte {
	return append([]byte(index), '/')
}

// address is hex-encoded (should be as it might have been done hexutil package)
func CollectionNetworkKey(network, address string) []byte {
	return append([]byte(fmt.Sprintf("%s:%s", network, address)), '/')
}
