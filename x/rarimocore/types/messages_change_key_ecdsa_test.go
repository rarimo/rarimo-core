package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"gitlab.com/rarify-protocol/rarimo-core/testutil/sample"
)

func TestMsgCreateChangeKeyECDSA_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateChangeKeyECDSA
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateChangeKeyECDSA{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateChangeKeyECDSA{
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

func TestMsgUpdateChangeKeyECDSA_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateChangeKeyECDSA
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateChangeKeyECDSA{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateChangeKeyECDSA{
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

func TestMsgDeleteChangeKeyECDSA_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteChangeKeyECDSA
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteChangeKeyECDSA{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteChangeKeyECDSA{
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
