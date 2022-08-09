package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "gitlab.com/rarify-protocol/rarimo-core/testutil/keeper"
	"gitlab.com/rarify-protocol/rarimo-core/testutil/nullify"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestDepositQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.RarimocoreKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNDeposit(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetDepositRequest
		response *types.QueryGetDepositResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetDepositRequest{
				Tx: msgs[0].Tx,
			},
			response: &types.QueryGetDepositResponse{Deposit: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetDepositRequest{
				Tx: msgs[1].Tx,
			},
			response: &types.QueryGetDepositResponse{Deposit: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetDepositRequest{
				Tx: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Deposit(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestDepositQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.RarimocoreKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNDeposit(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllDepositRequest {
		return &types.QueryAllDepositRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.DepositAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Deposit), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Deposit),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.DepositAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Deposit), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Deposit),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.DepositAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Deposit),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.DepositAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
