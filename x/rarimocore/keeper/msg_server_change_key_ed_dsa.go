package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarify-protocol/rarimo-core/crypto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func (k msgServer) CreateChangeKeyEdDSA(goCtx context.Context, msg *types.MsgCreateChangeKeyEdDSA) (*types.MsgCreateChangeKeyEdDSAResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//TODO add public key
	if err := crypto.VerifyEdDSA(msg.Signature, msg.NewKey, nil); err != nil {
		return nil, err
	}

	// Check if the value already exists
	_, isFound := k.GetChangeKeyEdDSA(
		ctx,
		msg.NewKey,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var changeKeyEdDSA = types.ChangeKeyEdDSA{
		Creator:   msg.Creator,
		NewKey:    msg.NewKey,
		Signature: msg.Signature,
	}

	k.SetChangeKeyEdDSA(
		ctx,
		changeKeyEdDSA,
	)
	return &types.MsgCreateChangeKeyEdDSAResponse{}, nil
}

func (k msgServer) UpdateChangeKeyEdDSA(goCtx context.Context, msg *types.MsgUpdateChangeKeyEdDSA) (*types.MsgUpdateChangeKeyEdDSAResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetChangeKeyEdDSA(
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

	var changeKeyEdDSA = types.ChangeKeyEdDSA{
		Creator:   msg.Creator,
		NewKey:    msg.NewKey,
		Signature: msg.Signature,
	}

	k.SetChangeKeyEdDSA(ctx, changeKeyEdDSA)

	return &types.MsgUpdateChangeKeyEdDSAResponse{}, nil
}

func (k msgServer) DeleteChangeKeyEdDSA(goCtx context.Context, msg *types.MsgDeleteChangeKeyEdDSA) (*types.MsgDeleteChangeKeyEdDSAResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetChangeKeyEdDSA(
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

	k.RemoveChangeKeyEdDSA(
		ctx,
		msg.NewKey,
	)

	return &types.MsgDeleteChangeKeyEdDSAResponse{}, nil
}
