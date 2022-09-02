package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func (k msgServer) CreateChangeKeyECDSA(goCtx context.Context, msg *types.MsgCreateChangeKeyECDSA) (*types.MsgCreateChangeKeyECDSAResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := crypto.VerifyECDSA(msg.Signature, msg.NewKey, k.GetKeyECDSA(ctx)); err != nil {
		return nil, err
	}

	// Check if the value already exists
	_, isFound := k.GetChangeKeyECDSA(
		ctx,
		msg.NewKey,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var changeKeyECDSA = types.ChangeKeyECDSA{
		Creator:   msg.Creator,
		NewKey:    msg.NewKey,
		Signature: msg.Signature,
	}

	k.SetChangeKeyECDSA(
		ctx,
		changeKeyECDSA,
	)

	k.UpdateKeyECDSA(
		ctx,
		msg.NewKey,
	)

	return &types.MsgCreateChangeKeyECDSAResponse{}, nil
}
