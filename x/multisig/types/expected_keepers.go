package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetModuleAddress(moduleName string) sdk.AccAddress
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	NewAccount(ctx sdk.Context, acc types.AccountI) types.AccountI
	SetAccount(ctx sdk.Context, acc types.AccountI)
	// Methods imported from account should be defined here
}
