package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	// ProposalKeyPrefix is the prefix to retrieve all Proposals
	ProposalKeyPrefix = "ProposalKey/value/"
)

// ProposalKey returns the store key to retrieve a proposal from the index fields
func ProposalKey(
	proposalId uint64,
) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, proposalId)
	return append(buf, []byte("/")...)
}
