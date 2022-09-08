package types

import (
	"encoding/binary"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ binary.ByteOrder

const (
	// DepositKeyPrefix is the prefix to retrieve all Deposit
	DepositKeyPrefix = "Deposit/value/"
)

// DepositKey returns the store key to retrieve a Deposit from the index fields
func DepositKey(
	index string,
) []byte {
	return append(hexutil.MustDecode(index), []byte("/")...)
}
