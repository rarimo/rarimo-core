package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	StateInfoKeyPrefix = "StateInfo/value/"
)

func StateInfoKey(
	s string,
) []byte {
	var key []byte
	key = append(key, []byte(s)...)
	key = append(key, []byte("/")...)

	return key
}
