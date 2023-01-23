package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	// ItemKeyPrefix is the prefix to retrieve all Item
	ItemKeyPrefix = "Item/value/"
)

// ItemKey returns the store key to retrieve an Item from the index fields
func ItemKey(
	index *ItemIndex,
) []byte {
	var key []byte

	key = append(key, []byte(index.Collection)...)
	key = append(key, []byte("/")...)

	key = append(key, []byte(index.Name)...)
	key = append(key, []byte("/")...)

	key = append(key, []byte(index.Symbol)...)
	key = append(key, []byte("/")...)

	key = append(key, []byte(index.Uri)...)
	key = append(key, []byte("/")...)

	return key
}

func (i *ItemIndex) Equal(o *ItemIndex) bool {
	if o == nil || i == nil {
		return i == o
	}

	return i.Collection == o.Collection && i.Name == o.Name && i.Symbol == o.Symbol && i.Uri == o.Uri
}
