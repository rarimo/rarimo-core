package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateVoteQuorum(p.VoteQuorum); err != nil {
		return err
	}

	if err := validateKeyMinOracleStake(p.VoteQuorum); err != nil {
		return err
	}

	return validateVoteThreshold(p.VoteThreshold)
}

func validateKeyMinOracleStake(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if _, ok := sdk.NewIntFromString(v); !ok {
		return fmt.Errorf("invalid amount")
	}

	return nil
}

func validateVoteQuorum(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	_, err := sdk.NewDecFromStr(v)
	return err
}

func validateVoteThreshold(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	_, err := sdk.NewDecFromStr(v)
	return err
}
