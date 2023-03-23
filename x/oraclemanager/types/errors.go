package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrOracleNotFound = sdkerrors.Register(ModuleName, 1100, "invalid oracle: does not exist")
	ErrInactiveOracle = sdkerrors.Register(ModuleName, 1101, "oracle is not active")
)
