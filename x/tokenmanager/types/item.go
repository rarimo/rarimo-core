package types

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func validateItem(i *Item) error {
	if i == nil {
		return fmt.Errorf("empty item")
	}

	if err := validateItemIndex(i.Index); err != nil {
		return err
	}

	if i.Meta == nil {
		return fmt.Errorf("empty item metadata")
	}

	if len(i.Meta.ImageUri) == 0 || len(i.Meta.ImageHash) == 0 || len(i.Meta.Seed) == 0 {
		return fmt.Errorf("invalid item metadata")
	}

	if _, err := hexutil.Decode(i.Meta.Seed); err != nil {
		return fmt.Errorf("invalid seed, can not decode %s", err.Error())
	}

	return nil
}

func validateItemIndex(i *ItemIndex) error {
	if i == nil {
		return fmt.Errorf("empty item index")
	}

	if len(i.Collection) == 0 {
		return fmt.Errorf("invalid collection")
	}

	return nil
}
