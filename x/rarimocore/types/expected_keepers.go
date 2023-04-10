package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	tokenmanagertypes "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	// Methods imported from bank should be defined here
}

// StakingKeeper defines the expected interface needed to retrieve validator account.
type StakingKeeper interface {
	Validator(ctx sdk.Context, address sdk.ValAddress) stakingtypes.ValidatorI
}

// TokenmanagerKeeper defines the expected interface needed to retrieve token and collection data.
type TokenmanagerKeeper interface {
	GetCollectionData(ctx sdk.Context, index *tokenmanagertypes.CollectionDataIndex) (tokenmanagertypes.CollectionData, bool)
	GetCollection(ctx sdk.Context, index string) (tokenmanagertypes.Collection, bool)
	GetOnChainItem(ctx sdk.Context, index *tokenmanagertypes.OnChainItemIndex) (tokenmanagertypes.OnChainItem, bool)
	SetOnChainItem(ctx sdk.Context, item tokenmanagertypes.OnChainItem)
	GetItem(ctx sdk.Context, index string) (tokenmanagertypes.Item, bool)
	SetItem(ctx sdk.Context, item tokenmanagertypes.Item)
	GetNetwork(ctx sdk.Context, name string) (tokenmanagertypes.NetworkParams, bool)
	GetSeed(ctx sdk.Context, seed string) (tokenmanagertypes.Seed, bool)
	SetSeed(ctx sdk.Context, seed tokenmanagertypes.Seed)
}
