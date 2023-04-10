package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	// GroupKeyPrefix is the prefix to retrieve all Groups
	GroupKeyPrefix = "GroupKey/value/"
)

// GroupKey returns the store key to retrieve a group from the index fields
func GroupKey(
	account string,
) []byte {
	return append([]byte(account), []byte("/")...)
}
