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

func createNConfirmation(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Confirmation {
	items := make([]types.Confirmation, n)
	for i := range items {
		items[i].Height = strconv.Itoa(i)

		keeper.SetConfirmation(ctx, items[i])
	}
	return items
}

func TestConfirmationGet(t *testing.T) {
	keeper, ctx := keepertest.RarimocoreKeeper(t)
	items := createNConfirmation(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetConfirmation(ctx,
			item.Height,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestConfirmationRemove(t *testing.T) {
	keeper, ctx := keepertest.RarimocoreKeeper(t)
	items := createNConfirmation(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveConfirmation(ctx,
			item.Height,
		)
		_, found := keeper.GetConfirmation(ctx,
			item.Height,
		)
		require.False(t, found)
	}
}

func TestConfirmationGetAll(t *testing.T) {
	keeper, ctx := keepertest.RarimocoreKeeper(t)
	items := createNConfirmation(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllConfirmation(ctx)),
	)
}
