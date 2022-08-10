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

func TestChangeKeyECDSAQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.RarimocoreKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNChangeKeyECDSA(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetChangeKeyECDSARequest
		response *types.QueryGetChangeKeyECDSAResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetChangeKeyECDSARequest{
				NewKey: msgs[0].NewKey,
			},
			response: &types.QueryGetChangeKeyECDSAResponse{ChangeKeyECDSA: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetChangeKeyECDSARequest{
				NewKey: msgs[1].NewKey,
			},
			response: &types.QueryGetChangeKeyECDSAResponse{ChangeKeyECDSA: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetChangeKeyECDSARequest{
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
			response, err := keeper.ChangeKeyECDSA(wctx, tc.request)
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

func TestChangeKeyECDSAQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.RarimocoreKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNChangeKeyECDSA(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllChangeKeyECDSARequest {
		return &types.QueryAllChangeKeyECDSARequest{
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
			resp, err := keeper.ChangeKeyECDSAAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ChangeKeyECDSA), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ChangeKeyECDSA),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ChangeKeyECDSAAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ChangeKeyECDSA), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ChangeKeyECDSA),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ChangeKeyECDSAAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ChangeKeyECDSA),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ChangeKeyECDSAAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
