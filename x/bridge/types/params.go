package types

import (
	"fmt"
)

var ErrEmptyWithdrawDenom = fmt.Errorf("withdraw denom cannot be empty")

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
	return validateWithdrawDenom(p.WithdrawDenom)
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
