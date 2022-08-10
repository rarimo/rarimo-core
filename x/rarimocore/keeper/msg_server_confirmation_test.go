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

func TestConfirmationMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.RarimocoreKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateConfirmation{Creator: creator,
			Height: strconv.Itoa(i),
		}
		_, err := srv.CreateConfirmation(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetConfirmation(ctx,
			expected.Height,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestConfirmationMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateConfirmation
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateConfirmation{Creator: creator,
				Height: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateConfirmation{Creator: "B",
				Height: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateConfirmation{Creator: creator,
				Height: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RarimocoreKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateConfirmation{Creator: creator,
				Height: strconv.Itoa(0),
			}
			_, err := srv.CreateConfirmation(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateConfirmation(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetConfirmation(ctx,
					expected.Height,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestConfirmationMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteConfirmation
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteConfirmation{Creator: creator,
				Height: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteConfirmation{Creator: "B",
				Height: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteConfirmation{Creator: creator,
				Height: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RarimocoreKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateConfirmation(wctx, &types.MsgCreateConfirmation{Creator: creator,
				Height: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteConfirmation(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetConfirmation(ctx,
					tc.request.Height,
				)
				require.False(t, found)
			}
		})
	}
}
