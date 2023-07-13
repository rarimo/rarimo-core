package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	rarimocoretypes "gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
	tokenmanagertypes "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

type RarimocoreKeeper interface {
	CreateTransferOperation(ctx sdk.Context, creator string, transfer *rarimocoretypes.Transfer, approved bool) error
	GetOperation(ctx sdk.Context, index string) (val rarimocoretypes.Operation, found bool)
	// Methods imported from rarimocore should be defined here
}

type TokenmanagerKeeper interface {
	GetNetwork(ctx sdk.Context, name string) (tokenmanagertypes.Network, bool)
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
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	// Methods imported from bank should be defined here
}
