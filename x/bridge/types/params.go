package types

import (
	"fmt"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var ErrEmptyWithdrawDenom = fmt.Errorf("withdraw denom cannot be empty")

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	ParamWithdrawDenom = []byte("WithdrawDenom")
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
		paramtypes.NewParamSetPair(ParamWithdrawDenom, &p.WithdrawDenom, validateWithdrawDenom),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateWithdrawDenom(p.WithdrawDenom); err != nil {
		return err
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateWithdrawDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == "" {
		return ErrEmptyWithdrawDenom
	}

	return nil
}
