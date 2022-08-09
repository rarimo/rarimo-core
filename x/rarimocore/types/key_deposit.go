package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DepositKeyPrefix is the prefix to retrieve all Deposit
	DepositKeyPrefix = "Deposit/value/"
)

// DepositKey returns the store key to retrieve a Deposit from the index fields
func DepositKey(
	tx string,
) []byte {
	var key []byte

	txBytes := []byte(tx)
	key = append(key, txBytes...)
	key = append(key, []byte("/")...)

	return key
}
