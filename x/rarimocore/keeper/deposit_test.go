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

func createNDeposit(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Deposit {
	items := make([]types.Deposit, n)
	for i := range items {
		items[i].Tx = strconv.Itoa(i)

		keeper.SetDeposit(ctx, items[i])
	}
	return items
}

func TestDepositGet(t *testing.T) {
	keeper, ctx := keepertest.RarimocoreKeeper(t)
	items := createNDeposit(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDeposit(ctx,
			item.Tx,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDepositRemove(t *testing.T) {
	keeper, ctx := keepertest.RarimocoreKeeper(t)
	items := createNDeposit(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDeposit(ctx,
			item.Tx,
		)
		_, found := keeper.GetDeposit(ctx,
			item.Tx,
		)
		require.False(t, found)
	}
}

func TestDepositGetAll(t *testing.T) {
	keeper, ctx := keepertest.RarimocoreKeeper(t)
	items := createNDeposit(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDeposit(ctx)),
	)
}
