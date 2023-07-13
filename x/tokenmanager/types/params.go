package types

import (
	"fmt"
)

// NewParams creates a new Params instance
func NewParams(networks []*Network) Params {
	return Params{
		networks,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return Params{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return validateNetworkParams(p.Networks)
}

func validateNetworkParams(i interface{}) error {
	v, ok := i.([]*Network)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, n := range v {
		if err := validateNetwork(n); err != nil {
			return err
		}
	}

	return nil
}

func validateNetwork(n *Network) error {
	if n == nil {
		return fmt.Errorf("invalid params: empty")
	}

	if len(n.Name) == 0 {
		return fmt.Errorf("invalid network name")
	}

	return nil
}
