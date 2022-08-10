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

func TestChangeKeyEdDSAQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.RarimocoreKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNChangeKeyEdDSA(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetChangeKeyEdDSARequest
		response *types.QueryGetChangeKeyEdDSAResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetChangeKeyEdDSARequest{
				NewKey: msgs[0].NewKey,
			},
			response: &types.QueryGetChangeKeyEdDSAResponse{ChangeKeyEdDSA: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetChangeKeyEdDSARequest{
				NewKey: msgs[1].NewKey,
			},
			response: &types.QueryGetChangeKeyEdDSAResponse{ChangeKeyEdDSA: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetChangeKeyEdDSARequest{
				NewKey: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ChangeKeyEdDSA(wctx, tc.request)
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

func TestChangeKeyEdDSAQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.RarimocoreKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNChangeKeyEdDSA(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllChangeKeyEdDSARequest {
		return &types.QueryAllChangeKeyEdDSARequest{
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
			resp, err := keeper.ChangeKeyEdDSAAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ChangeKeyEdDSA), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ChangeKeyEdDSA),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ChangeKeyEdDSAAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ChangeKeyEdDSA), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ChangeKeyEdDSA),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ChangeKeyEdDSAAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ChangeKeyEdDSA),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ChangeKeyEdDSAAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
