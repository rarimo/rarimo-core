package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarify-protocol/rarimo-core/crypto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func (k msgServer) CreateChangeKeyECDSA(goCtx context.Context, msg *types.MsgCreateChangeKeyECDSA) (*types.MsgCreateChangeKeyECDSAResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//TODO add public key
	if err := crypto.VerifyECDSA(msg.Signature, msg.NewKey, nil); err != nil {
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
	return &types.MsgCreateChangeKeyECDSAResponse{}, nil
}

func (k msgServer) UpdateChangeKeyECDSA(goCtx context.Context, msg *types.MsgUpdateChangeKeyECDSA) (*types.MsgUpdateChangeKeyECDSAResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetChangeKeyECDSA(
		ctx,
		msg.NewKey,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var changeKeyECDSA = types.ChangeKeyECDSA{
		Creator:   msg.Creator,
		NewKey:    msg.NewKey,
		Signature: msg.Signature,
	}

	k.SetChangeKeyECDSA(ctx, changeKeyECDSA)

	return &types.MsgUpdateChangeKeyECDSAResponse{}, nil
}

func (k msgServer) DeleteChangeKeyECDSA(goCtx context.Context, msg *types.MsgDeleteChangeKeyECDSA) (*types.MsgDeleteChangeKeyECDSAResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetChangeKeyECDSA(
		ctx,
		msg.NewKey,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveChangeKeyECDSA(
		ctx,
		msg.NewKey,
	)

	return &types.MsgDeleteChangeKeyECDSAResponse{}, nil
}
