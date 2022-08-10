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

func createNChangeKeyECDSA(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ChangeKeyECDSA {
	items := make([]types.ChangeKeyECDSA, n)
	for i := range items {
		items[i].NewKey = strconv.Itoa(i)

		keeper.SetChangeKeyECDSA(ctx, items[i])
	}
	return items
}

func TestChangeKeyECDSAGet(t *testing.T) {
	keeper, ctx := keepertest.RarimocoreKeeper(t)
	items := createNChangeKeyECDSA(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetChangeKeyECDSA(ctx,
			item.NewKey,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestChangeKeyECDSARemove(t *testing.T) {
	keeper, ctx := keepertest.RarimocoreKeeper(t)
	items := createNChangeKeyECDSA(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveChangeKeyECDSA(ctx,
			item.NewKey,
		)
		_, found := keeper.GetChangeKeyECDSA(ctx,
			item.NewKey,
		)
		require.False(t, found)
	}
}

func TestChangeKeyECDSAGetAll(t *testing.T) {
	keeper, ctx := keepertest.RarimocoreKeeper(t)
	items := createNChangeKeyECDSA(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllChangeKeyECDSA(ctx)),
	)
}
