package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"gitlab.com/rarify-protocol/rarimo-core/testutil/sample"
)

func TestMsgCreateConfirmation_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateConfirmation
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateConfirmation{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateConfirmation{
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

func TestMsgUpdateConfirmation_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateConfirmation
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateConfirmation{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateConfirmation{
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

func TestMsgDeleteConfirmation_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteConfirmation
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteConfirmation{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteConfirmation{
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
