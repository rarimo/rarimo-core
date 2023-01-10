package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Collections: []CollectionInfo{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	collectionsCheck := make(map[string]struct{})

	for _, collection := range gs.Collections {
		if _, ok := collectionsCheck[collection.Index]; ok {
			return fmt.Errorf("duplicated collection index: %s", collection.Index)
		}
		collectionsCheck[collection.Index] = struct{}{}
	}

	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
