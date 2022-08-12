package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gitlab.com/rarify-protocol/rarimo-core/crypto"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyECDSA = []byte("keyECDSAParam")
	KeyEdDSA = []byte("keyEdDSAParam")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(keyECDSA, keyEdDSA string) Params {
	return Params{
		KeyECDSA: keyECDSA,
		KeyEdDSA: keyEdDSA,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams("", "")
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyECDSA, &p.KeyECDSA, validateKeyECDSA),
		paramtypes.NewParamSetPair(KeyEdDSA, &p.KeyEdDSA, validateKeyEdDSA),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateKeyECDSA(p.KeyECDSA); err != nil {
		return err
	}

	if err := validateKeyEdDSA(p.KeyEdDSA); err != nil {
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

func validateKeyEdDSA(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return crypto.ValidateEdDSAKey(v)
}
