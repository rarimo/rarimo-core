package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyECDSA = []byte("keyECDSAParam")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(keyECDSA string) Params {
	return Params{
		KeyECDSA: keyECDSA,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams("")
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyECDSA, &p.KeyECDSA, validateKeyECDSA),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateKeyECDSA(p.KeyECDSA); err != nil {
		return err
	}
	return nil
}

func validateKeyECDSA(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return crypto.ValidateECDSAKey(v)
}
