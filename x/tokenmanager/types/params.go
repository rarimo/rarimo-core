package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	Networks = []byte("networks")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(networks map[string]*ChainParams) Params {
	return Params{
		networks,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return Params{}
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(Networks, &p.Networks, func(value interface{}) error { return nil }),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}
