package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	// OracleKeyPrefix is the prefix to retrieve all Oracle
	OracleKeyPrefix = "Oracle/value/"
)

// OracleKey returns the store key to retrieve an Oracle from the index fields
func OracleKey(
	index *OracleIndex,
) []byte {
	var key []byte

	key = append(key, []byte(index.Chain)...)
	key = append(key, []byte("/")...)

	key = append(key, []byte(index.Account)...)
	key = append(key, []byte("/")...)

	return key
}
