package types

import (
	"errors"
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	ECDSAPublicKeySize = 65
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	ParamKeyECDSA         = []byte("KeyECDSA")
	ParamThreshold        = []byte("Threshold")
	ParamParties          = []byte("Parties")
	ParamIsUpdateRequired = []byte("IsUpdateRequired")
	ParamLastSignature    = []byte("LastSignature")
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
	return Params{}
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamKeyECDSA, &p.KeyECDSA, validateKeyECDSA),
		paramtypes.NewParamSetPair(ParamThreshold, &p.Threshold, validateThreshold),
		paramtypes.NewParamSetPair(ParamParties, &p.Parties, validateParties),
		paramtypes.NewParamSetPair(ParamIsUpdateRequired, &p.IsUpdateRequired, validateIsUpdateRequired),
		paramtypes.NewParamSetPair(ParamLastSignature, &p.LastSignature, validateLastSignature),
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

	// not initialized
	if len(v) == 0 {
		return nil
	}

	keyBytes, err := hexutil.Decode(v)
	if err != nil {
		return fmt.Errorf("invalid ECDSA key format %v", err)
	}

	if len(keyBytes) != ECDSAPublicKeySize {
		return fmt.Errorf("invalid ECDSA key len")
	}

	return nil
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

func validateIsUpdateRequired(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}

func validateLastSignature(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	_, err := hexutil.Decode(v)
	return err
}
