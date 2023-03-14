package keeper

import (
	"fmt"

	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/tendermint/tendermint/libs/log"
	oraclemanagermodulekeeper "gitlab.com/rarimo/rarimo-core/x/oraclemanager/keeper"
	tmkeeper "gitlab.com/rarimo/rarimo-core/x/tokenmanager/keeper"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace
		tm         *tmkeeper.Keeper
		staking    *stakingkeeper.Keeper
		gov        *govkeeper.Keeper
		oracle     *oraclemanagermodulekeeper.Keeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	tmkeeper *tmkeeper.Keeper,
	staking *stakingkeeper.Keeper,
	oracle *oraclemanagermodulekeeper.Keeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		paramstore: ps,
		memKey:     memKey,
		tm:         tmkeeper,
		staking:    staking,
		oracle:     oracle,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
