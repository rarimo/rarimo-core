package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"gitlab.com/rarify-protocol/rarimo-core/testutil/sample"
)

func TestMsgCreateChangeKeyEdDSA_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateChangeKeyEdDSA
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateChangeKeyEdDSA{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateChangeKeyEdDSA{
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

func TestMsgUpdateChangeKeyEdDSA_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateChangeKeyEdDSA
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateChangeKeyEdDSA{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateChangeKeyEdDSA{
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

func TestMsgDeleteChangeKeyEdDSA_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteChangeKeyEdDSA
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteChangeKeyEdDSA{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteChangeKeyEdDSA{
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
