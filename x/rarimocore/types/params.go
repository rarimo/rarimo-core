package types

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	ECDSAPublicKeySize = 65
)

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

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateKeyECDSA(p.KeyECDSA); err != nil {
		return err
	}

	if err := validateStakeAmount(p.StakeAmount); err != nil {
		return err
	}

	if err := validateFreezeBlocksPeriod(p.FreezeBlocksPeriod); err != nil {
		return err
	}

	if err := validateMaxViolationsCount(p.MaxViolationsCount); err != nil {
		return err
	}

	if err := validateStakeDenom(p.StakeDenom); err != nil {
		return err
	}

	return validateLastSignature(p.LastSignature)
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

func validateLastSignature(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	_, err := hexutil.Decode(v)
	return err
}

func validateStakeAmount(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	amount, ok := big.NewInt(0).SetString(v, 10)
	if !ok {
		return fmt.Errorf("invalid stake amount")
	}

	if amount.Cmp(big.NewInt(0)) <= 0 {
		return fmt.Errorf("invalid stake amount")
	}

	return nil
}

func validateStakeDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if len(v) == 0 {
		return fmt.Errorf("invalid stake denom")
	}

	return nil
}

func validateMaxViolationsCount(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max violations count can not be zero")
	}

	return nil
}

func validateFreezeBlocksPeriod(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("freeze blocks periood can not be zero")
	}

	return nil
}
