package types

import "fmt"

func validateCollection(c *Collection) error {
	if c == nil {
		return fmt.Errorf("empty collection")
	}

	if len(c.Index) == 0 {
		return fmt.Errorf("invalid index")
	}

	if c.Meta == nil {
		return fmt.Errorf("empty meta")
	}

	return nil
}

func validateCollectionData(c *CollectionData) error {
	if c == nil {
		return fmt.Errorf("empty collection data")
	}

	return validateCollectionDataIndex(c.Index)
}

func validateCollectionDataIndex(c *CollectionDataIndex) error {
	if c == nil {
		return fmt.Errorf("empty index")
	}

	if len(c.Chain) == 0 {
		return fmt.Errorf("invalid chain")
	}

	if len(c.Address) == 0 {
		return fmt.Errorf("invalid address")
	}

	return nil
}
