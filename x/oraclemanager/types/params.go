package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	ParamKeyMinOracleStake      = []byte("KeyMinOracleStake")
	ParamKeyCheckOperationDelta = []byte("KeyCheckOperationDelta")
	ParamKeyMaxViolationsCount  = []byte("KeyMaxViolationsCount")
	ParamKeyMaxMissedCount      = []byte("KeyMaxMissedCount")
	ParamKeySlashedFreezeBlocks = []byte("KeySlashedFreezeBlocks")
	ParamKeyMinOraclesCount     = []byte("KeyMinOraclesCount")
	ParamKeyStakeDenom          = []byte("KeyStakeDenom")
	ParamVoteQuorum             = []byte("VoteQuorum")
	ParamVoteThreshold          = []byte("VoteThreshold")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamKeyMinOracleStake, &p.MinOracleStake, validateKeyMinOracleStake),
		paramtypes.NewParamSetPair(ParamKeyCheckOperationDelta, &p.CheckOperationDelta, validateBlocks),
		paramtypes.NewParamSetPair(ParamKeyMaxViolationsCount, &p.MaxViolationsCount, validateCount),
		paramtypes.NewParamSetPair(ParamKeyMaxMissedCount, &p.MaxMissedCount, validateCount),
		paramtypes.NewParamSetPair(ParamKeySlashedFreezeBlocks, &p.SlashedFreezeBlocks, validateBlocks),
		paramtypes.NewParamSetPair(ParamKeyMinOraclesCount, &p.MinOraclesCount, validateCount),
		paramtypes.NewParamSetPair(ParamKeyStakeDenom, &p.StakeDenom, validateDenom),
		paramtypes.NewParamSetPair(ParamVoteQuorum, &p.VoteQuorum, validateVoteQuorum),
		paramtypes.NewParamSetPair(ParamVoteThreshold, &p.VoteThreshold, validateVoteThreshold),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
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

func validateBlocks(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateCount(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid count")
	}

	return nil
}

func validateDenom(i interface{}) error {
	_, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
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
