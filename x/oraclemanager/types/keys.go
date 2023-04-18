package types

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/types/kv"
)

const (
	// ModuleName defines the module name
	ModuleName = "oraclemanager"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oraclemanager"

	// ParamsKey defines the store key for module params entry
	ParamsKey = "params"

	lenUint64 = 8
)

var (
	MonitorOperationQueueKeyPrefix = []byte{0x001}
	FreezedOracleQueueKeyPrefix    = []byte{0x002}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func MonitorOperationQueueKey(op string, block uint64) []byte {
	return append(MonitorOperationQueueByBlockKey(block), KeyPrefix(op)...)
}

func FreezedOracleQueueKey(index *OracleIndex, block uint64) []byte {
	key := append(FreezedOracleQueueByBlockKey(block), strKey(index.Chain)...)
	return append(key, []byte(index.Account)...)
}

func MonitorOperationQueueByBlockKey(endBlock uint64) []byte {
	return append(MonitorOperationQueueKeyPrefix, getBlockKey(endBlock)...)
}

func FreezedOracleQueueByBlockKey(endBlock uint64) []byte {
	return append(FreezedOracleQueueKeyPrefix, getBlockKey(endBlock)...)
}

func SplitMonitorOperationQueueKey(key []byte) (string, uint64) {
	index, block := splitKeyWithBlock(key)
	return string(index), block
}

func SplitFreezedOracleQueueKey(key []byte) (*OracleIndex, uint64) {
	index, block := splitKeyWithBlock(key)
	kv.AssertKeyAtLeastLength(index, 1)

	chainLen := index[0]
	return &OracleIndex{Chain: string(index[1 : 1+chainLen]), Account: string(index[1+chainLen:])}, block
}

func getBlockKey(i uint64) []byte {
	b := make([]byte, lenUint64)
	binary.LittleEndian.PutUint64(b, i)
	return b
}

func splitKeyWithBlock(key []byte) (index []byte, endBlock uint64) {
	kv.AssertKeyAtLeastLength(key[1:], lenUint64)

	endBlock = binary.LittleEndian.Uint64(key[1 : 1+lenUint64])
	index = key[1+lenUint64:]
	return
}

func strKey(s string) []byte {
	if len(s) > 255 {
		panic("to long string")
	}

	return append([]byte{byte(len(s))}, []byte(s)...)
}
