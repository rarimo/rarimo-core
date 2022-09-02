package types

import (
	"encoding/binary"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ binary.ByteOrder

const (
	// ItemKeyPrefix is the prefix to retrieve all Item
	ItemKeyPrefix = "Item/value/"
)

// ItemKey returns the store key to retrieve a Item from the index fields
func ItemKey(
	tokenAddress string,
	tokenId string,
) []byte {
	var key []byte

	key = append(key, hexutil.MustDecode(tokenAddress)...)
	key = append(key, []byte("/")...)

	key = append(key, hexutil.MustDecode(tokenId)...)
	key = append(key, []byte("/")...)

	return key
}
