package types

import (
	"errors"
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyECDSA  = []byte("keyECDSAParam")
	Threshold = []byte("thresholdParam")
	Parties   = []byte("partiesParam")
	Steps     = []byte("strpsParam")
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
		paramtypes.NewParamSetPair(Threshold, &p.Threshold, validateThreshold),
		paramtypes.NewParamSetPair(Parties, &p.Parties, validateParties),
		paramtypes.NewParamSetPair(Steps, &p.Steps, validateSteps),
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

func validateThreshold(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return errors.New("threshold can not be zero")
	}

	return nil
}

func validateParties(i interface{}) error {
	v, ok := i.([]*Party)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if len(v) == 0 {
		return errors.New("there should be at least one party")
	}

	return nil
}

func validateSteps(i interface{}) error {
	v, ok := i.([]*Step)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if len(v) != 3 {
		return errors.New("should be 3 steps: proposing, accepting, signing")
	}

	return nil
}
