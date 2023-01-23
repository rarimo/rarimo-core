package types

import (
	"fmt"

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
func NewParams(networks []*NetworkParams) Params {
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
		paramtypes.NewParamSetPair(Networks, &p.Networks, validateNetworkParams),
	}
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
	if len(n.Name) == 0 {
		return fmt.Errorf("invalida network name")
	}

	if len(n.Types) == 0 {
		return fmt.Errorf("invalida network types")
	}

	if len(n.Contract) == 0 {
		return fmt.Errorf("invalida network contract address")
	}

	return nil
}
