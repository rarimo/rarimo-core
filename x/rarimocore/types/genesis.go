package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DepositList:        []Deposit{},
		ConfirmationList:   []Confirmation{},
		ChangeKeyECDSAList: []ChangeKeyECDSA{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in deposit
	depositIndexMap := make(map[string]struct{})

	for _, elem := range gs.DepositList {
		index := string(DepositKey(elem.Index))
		if _, ok := depositIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for deposit")
		}
		depositIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in confirmation
	confirmationIndexMap := make(map[string]struct{})

	for _, elem := range gs.ConfirmationList {
		index := string(ConfirmationKey(elem.Root))
		if _, ok := confirmationIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for confirmation")
		}
		confirmationIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in changeKeyECDSA
	changeKeyECDSAIndexMap := make(map[string]struct{})

	for _, elem := range gs.ChangeKeyECDSAList {
		index := string(ChangeKeyECDSAKey(elem.NewKey))
		if _, ok := changeKeyECDSAIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for changeKeyECDSA")
		}
		changeKeyECDSAIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
