package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "gitlab.com/rarify-protocol/rarimo-core/testutil/keeper"
	"gitlab.com/rarify-protocol/rarimo-core/testutil/nullify"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/keeper"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNChangeKeyEdDSA(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ChangeKeyEdDSA {
	items := make([]types.ChangeKeyEdDSA, n)
	for i := range items {
		items[i].NewKey = strconv.Itoa(i)

		keeper.SetChangeKeyEdDSA(ctx, items[i])
	}
	return items
}

func TestChangeKeyEdDSAGet(t *testing.T) {
	keeper, ctx := keepertest.RarimocoreKeeper(t)
	items := createNChangeKeyEdDSA(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetChangeKeyEdDSA(ctx,
			item.NewKey,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestChangeKeyEdDSARemove(t *testing.T) {
	keeper, ctx := keepertest.RarimocoreKeeper(t)
	items := createNChangeKeyEdDSA(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveChangeKeyEdDSA(ctx,
			item.NewKey,
		)
		_, found := keeper.GetChangeKeyEdDSA(ctx,
			item.NewKey,
		)
		require.False(t, found)
	}
}

func TestChangeKeyEdDSAGetAll(t *testing.T) {
	keeper, ctx := keepertest.RarimocoreKeeper(t)
	items := createNChangeKeyEdDSA(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllChangeKeyEdDSA(ctx)),
	)
}
