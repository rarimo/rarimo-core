package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	rarimocoretypes "gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
	tokenmanagertypes "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

type TokenmanagerKeeper interface {
	GetParams(ctx sdk.Context) (params tokenmanagertypes.Params)
}

type RarimocoreKeeper interface {
	SetOperation(ctx sdk.Context, op rarimocoretypes.Operation)
	GetOperation(ctx sdk.Context, index string) (val rarimocoretypes.Operation, found bool)
	// Methods imported from rarimocore should be defined here
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	MintTokens(ctx sdk.Context, address sdk.AccAddress, amounts sdk.Coins) error
	BurnTokens(ctx sdk.Context, address sdk.AccAddress, amounts sdk.Coins) error
	// Methods imported from bank should be defined here
}
