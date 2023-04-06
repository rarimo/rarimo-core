package types

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs *GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

var _ types.UnpackInterfacesMessage = &GenesisState{}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (gs *GenesisState) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	for _, p := range gs.ProposalList {
		err := p.UnpackInterfaces(unpacker)
		if err != nil {
			return err
		}
	}
	return nil
}
