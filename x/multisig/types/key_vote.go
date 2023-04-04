package types

import (
	"bytes"
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	// VoteKeyPrefix is the prefix to retrieve all Votes
	VoteKeyPrefix = "VoteKey/value/"
)

// VoteKey returns the store key to retrieve a vote from the index fields
func VoteKey(
	proposalId uint64,
	voter string,
) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, proposalId)
	return bytes.Join([][]byte{buf, []byte(voter)}, []byte("/"))
}
