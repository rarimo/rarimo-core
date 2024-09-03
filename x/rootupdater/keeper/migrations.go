package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v2 "github.com/rarimo/rarimo-core/x/rootupdater/migrations/v2"
)

type Migrator struct {
	keeper Keeper
}

func NewMigrator(keeper Keeper) Migrator {
	return Migrator{
		keeper: keeper,
	}
}

func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	if err := v2.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc); err != nil {
		return err
	}

	return nil
}
