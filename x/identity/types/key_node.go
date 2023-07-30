package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	NodeKeyPrefix = "Node/value/"
)

func NodeKey(
	s string,
) []byte {
	var key []byte
	key = append(key, []byte(s)...)
	key = append(key, []byte("/")...)
	return key
}
