package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/rarimocore module sentinel errors
var (
	ErrInvalidKey = sdkerrors.Register(ModuleName, 1100, "The public key is invalid in some way")
)
