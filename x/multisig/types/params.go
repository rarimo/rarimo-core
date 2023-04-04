package types

import (
	"fmt"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var ErrInvalidPeriod = fmt.Errorf("period cannot be 0")

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	ParamGroupSequence    = []byte("GroupSequence")
	ParamProposalSequence = []byte("ProposalSequence")
	ParamPrunePeriod      = []byte("PrunePeriod")
	ParamVotingPeriod     = []byte("VotingPeriod")
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
		paramtypes.NewParamSetPair(ParamGroupSequence, &p.GroupSequence, validateSequence),
		paramtypes.NewParamSetPair(ParamProposalSequence, &p.ProposalSequence, validateSequence),
		paramtypes.NewParamSetPair(ParamPrunePeriod, &p.PrunePeriod, validatePeriod),
		paramtypes.NewParamSetPair(ParamVotingPeriod, &p.VotingPeriod, validatePeriod),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateSequence(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validatePeriod(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return ErrInvalidPeriod
	}

	return nil
}
