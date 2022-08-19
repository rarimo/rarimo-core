package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "gitlab.com/rarify-protocol/rarimo-core/testutil/keeper"
	"gitlab.com/rarify-protocol/rarimo-core/testutil/nullify"
	"gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/keeper"
	"gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNItem(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Item {
	items := make([]types.Item, n)
	for i := range items {
		items[i].TokenAddress = strconv.Itoa(i)
		items[i].TokenId = strconv.Itoa(i)

		keeper.SetItem(ctx, items[i])
	}
	return items
}

func TestItemGet(t *testing.T) {
	keeper, ctx := keepertest.TokenmanagerKeeper(t)
	items := createNItem(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetItem(ctx,
			item.TokenAddress,
			item.TokenId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestItemRemove(t *testing.T) {
	keeper, ctx := keepertest.TokenmanagerKeeper(t)
	items := createNItem(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveItem(ctx,
			item.TokenAddress,
			item.TokenId,
		)
		_, found := keeper.GetItem(ctx,
			item.TokenAddress,
			item.TokenId,
		)
		require.False(t, found)
	}
}

func TestItemGetAll(t *testing.T) {
	keeper, ctx := keepertest.TokenmanagerKeeper(t)
	items := createNItem(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllItem(ctx)),
	)
}
