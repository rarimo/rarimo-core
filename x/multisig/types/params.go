package types

import (
	"fmt"
)

var ErrInvalidPeriod = fmt.Errorf("period cannot be 0")

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
	if err := validatePeriod(p.VotingPeriod); err != nil {
		return err
	}

	return validatePeriod(p.PrunePeriod)
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
