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
	chain string,
) []byte {
	var key []byte

	key = append(key, []byte(tokenAddress)...)
	key = append(key, []byte("/")...)

	key = append(key, []byte(tokenId)...)
	key = append(key, []byte("/")...)

	key = append(key, []byte(chain)...)
	key = append(key, []byte("/")...)

	return key
}

func ItemKeyToIndex(key []byte) string {
	return hexutil.Encode(key)
}

func ItemIndex(
	tokenAddress string,
	tokenId string,
	chain string,
) string {
	return ItemKeyToIndex(ItemKey(tokenAddress, tokenId, chain))
}
