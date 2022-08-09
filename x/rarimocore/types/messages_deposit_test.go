package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"gitlab.com/rarify-protocol/rarimo-core/testutil/sample"
)

func TestMsgCreateDeposit_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateDeposit
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateDeposit{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateDeposit{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateDeposit_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateDeposit
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateDeposit{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateDeposit{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteDeposit_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteDeposit
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteDeposit{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteDeposit{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
