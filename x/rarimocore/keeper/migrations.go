package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v3 "github.com/rarimo/rarimo-core/x/rarimocore/migrations/v3"
)

type Migrator struct {
	keeper Keeper
}

func NewMigrator(keeper Keeper) Migrator {
	return Migrator{
		keeper: keeper,
	}
}

func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	if err := v3.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc); err != nil {
		return err
	}

	return nil
}
