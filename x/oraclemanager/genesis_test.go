package oraclemanager_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "gitlab.com/rarimo/rarimo-core/testutil/keeper"
	"gitlab.com/rarimo/rarimo-core/testutil/nullify"
	"gitlab.com/rarimo/rarimo-core/x/oraclemanager"
	"gitlab.com/rarimo/rarimo-core/x/oraclemanager/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OraclemanagerKeeper(t)
	oraclemanager.InitGenesis(ctx, *k, genesisState)
	got := oraclemanager.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
