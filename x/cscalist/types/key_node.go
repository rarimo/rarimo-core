package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	// NodeKeyPrefix defines the prefix for the Merkle tree's node entry
	NodeKeyPrefix = "Node/value/"
)

func NodeKey(s string) []byte {
	return []byte(s + "/")
}
