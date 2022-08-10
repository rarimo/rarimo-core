package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "gitlab.com/rarify-protocol/rarimo-core/testutil/keeper"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/keeper"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestChangeKeyEdDSAMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.RarimocoreKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateChangeKeyEdDSA{Creator: creator,
			NewKey: strconv.Itoa(i),
		}
		_, err := srv.CreateChangeKeyEdDSA(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetChangeKeyEdDSA(ctx,
			expected.NewKey,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestChangeKeyEdDSAMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateChangeKeyEdDSA
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateChangeKeyEdDSA{Creator: creator,
				NewKey: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateChangeKeyEdDSA{Creator: "B",
				NewKey: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateChangeKeyEdDSA{Creator: creator,
				NewKey: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RarimocoreKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateChangeKeyEdDSA{Creator: creator,
				NewKey: strconv.Itoa(0),
			}
			_, err := srv.CreateChangeKeyEdDSA(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateChangeKeyEdDSA(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetChangeKeyEdDSA(ctx,
					expected.NewKey,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestChangeKeyEdDSAMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteChangeKeyEdDSA
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteChangeKeyEdDSA{Creator: creator,
				NewKey: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteChangeKeyEdDSA{Creator: "B",
				NewKey: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteChangeKeyEdDSA{Creator: creator,
				NewKey: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RarimocoreKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateChangeKeyEdDSA(wctx, &types.MsgCreateChangeKeyEdDSA{Creator: creator,
				NewKey: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteChangeKeyEdDSA(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetChangeKeyEdDSA(ctx,
					tc.request.NewKey,
				)
				require.False(t, found)
			}
		})
	}
}
