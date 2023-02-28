package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	// VoteKeyPrefix is the prefix to retrieve all Vote
	VoteKeyPrefix = "Vote/value/"
)

// VoteKey returns the store key to retrieve an Vote from the index fields
func VoteKey(
	index *VoteIndex,
) []byte {
	var key []byte

	key = append(key, []byte(index.Operation)...)
	key = append(key, []byte("/")...)

	key = append(key, []byte(index.Validator)...)
	key = append(key, []byte("/")...)

	return key
}
