package types

import (
	"fmt"
)

// NewParams creates a new Params instance
func NewParams(networks []*NetworkParams) Params {
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
	v, ok := i.([]*NetworkParams)
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

func validateNetwork(n *NetworkParams) error {
	if n == nil {
		return fmt.Errorf("invalid params: empty")
	}

	if len(n.Name) == 0 {
		return fmt.Errorf("invalid network name")
	}

	if len(n.Contract) == 0 {
		return fmt.Errorf("invalid network contract address")
	}

	if len(n.Fee.FeeContract) == 0 {
		return fmt.Errorf("invalid fee contract address")
	}

	if len(n.Fee.FeeTokens) == 0 {
		return fmt.Errorf("invalid fee token list: empty")
	}

	return nil
}
