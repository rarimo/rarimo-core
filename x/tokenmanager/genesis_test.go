package tokenmanager_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "gitlab.com/rarify-protocol/rarimo-core/testutil/keeper"
	"gitlab.com/rarify-protocol/rarimo-core/testutil/nullify"
	"gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager"
	"gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ItemList: []types.Item{
			{
				TokenAddress: "0",
				TokenId:      "0",
			},
			{
				TokenAddress: "1",
				TokenId:      "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TokenmanagerKeeper(t)
	tokenmanager.InitGenesis(ctx, *k, genesisState)
	got := tokenmanager.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ItemList, got.ItemList)
	// this line is used by starport scaffolding # genesis/test/assert
}
