package types

const (
	// ModuleName defines the module name
	ModuleName = "cscalist"
	// StoreKey defines the primary module store key
	StoreKey = ModuleName
	// RouterKey is the message route for slashing
	RouterKey = ModuleName
	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cscalist"
	// ParamsKey defines the store key for module params entry
	ParamsKey = "params"
	// NodeKeyPrefix defines the prefix for the Merkle tree's node entry
	NodeKeyPrefix = "Node/value/"
)
