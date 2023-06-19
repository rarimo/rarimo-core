package types

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/types"
)

var _ binary.ByteOrder

const (
	VestingKeyPrefix = "Vesting/value/"
)

func VestingKey(
	index uint64,
) []byte {
	var key []byte

	key = append(key, types.Uint64ToBigEndian(index)...)
	key = append(key, []byte("/")...)

	return key
}
