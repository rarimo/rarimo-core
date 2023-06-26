package types

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func validateItem(i *Item) error {
	if i == nil {
		return fmt.Errorf("empty item")
	}

	if len(i.Index) == 0 {
		return fmt.Errorf("invalid item index")
	}

	if _, err := hexutil.Decode(i.Meta.Seed); i.Meta.Seed != "" && err != nil {
		return fmt.Errorf("invalid seed, can not decode %s", err.Error())
	}

	return nil
}
