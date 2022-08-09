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

func TestDepositMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.RarimocoreKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateDeposit{Creator: creator,
			Tx: strconv.Itoa(i),
		}
		_, err := srv.CreateDeposit(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetDeposit(ctx,
			expected.Tx,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestDepositMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateDeposit
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateDeposit{Creator: creator,
				Tx: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateDeposit{Creator: "B",
				Tx: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateDeposit{Creator: creator,
				Tx: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RarimocoreKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateDeposit{Creator: creator,
				Tx: strconv.Itoa(0),
			}
			_, err := srv.CreateDeposit(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateDeposit(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetDeposit(ctx,
					expected.Tx,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestDepositMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteDeposit
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteDeposit{Creator: creator,
				Tx: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteDeposit{Creator: "B",
				Tx: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteDeposit{Creator: creator,
				Tx: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RarimocoreKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateDeposit(wctx, &types.MsgCreateDeposit{Creator: creator,
				Tx: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteDeposit(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetDeposit(ctx,
					tc.request.Tx,
				)
				require.False(t, found)
			}
		})
	}
}
