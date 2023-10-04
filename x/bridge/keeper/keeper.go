package keeper

import (
	"fmt"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/bridge/types"
)

type (
	Keeper struct {
		cdc               codec.BinaryCodec
		storeKey          storetypes.StoreKey
		memKey            storetypes.StoreKey
		rarimocoreKeeper  types.RarimocoreKeeper
		bankKeeper        types.BankKeeper
		tokenmanagerKeepr types.TokenmanagerKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	bankKeeper types.BankKeeper,
	rarimocoreKeeper types.RarimocoreKeeper,
	tokenmanagerKeepr types.TokenmanagerKeeper,
) *Keeper {
	return &Keeper{
		cdc:               cdc,
		storeKey:          storeKey,
		memKey:            memKey,
		bankKeeper:        bankKeeper,
		rarimocoreKeeper:  rarimocoreKeeper,
		tokenmanagerKeepr: tokenmanagerKeepr,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
