package types

import "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "vestingmint"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_vestingmint"

	// ParamsKey defines the store key for module params entry
	ParamsKey = "params"

	ActiveVestingsKeyPrefixSize = 1
	BigEndianU64Size            = 8
)

var (
	ActiveVestingsKeyPrefix = []byte{0x001}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func ActiveVestingsQueueKey(block uint64, index uint64) []byte {
	return append(ActiveVestingsQueueBlockKey(block), types.Uint64ToBigEndian(index)...)
}

func ActiveVestingsQueueBlockKey(block uint64) []byte {
	return append(ActiveVestingsKeyPrefix, types.Uint64ToBigEndian(block)...)
}

func GetVestingIndexFromQueueKey(key []byte) uint64 {
	return types.BigEndianToUint64(key[ActiveVestingsKeyPrefixSize+BigEndianU64Size:])
}
