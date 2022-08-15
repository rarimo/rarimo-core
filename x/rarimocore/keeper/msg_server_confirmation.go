package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarify-protocol/rarimo-core/crypto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func (k msgServer) CreateConfirmation(goCtx context.Context, msg *types.MsgCreateConfirmation) (*types.MsgCreateConfirmationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := crypto.VerifyECDSA(msg.SigECDSA, msg.Root, k.GetKeyECDSA(ctx)); err != nil {
		return nil, err
	}

	if err := crypto.VerifyEdDSA(msg.SigEdDSA, msg.Root, k.GetKeyEdDSA(ctx)); err != nil {
		return nil, err
	}

	// Check if the value already exists
	_, isFound := k.GetConfirmation(
		ctx,
		msg.Height,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	for _, hash := range msg.Hashes {
		_, ok := k.GetDeposit(ctx, hash)
		if !ok {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "deposit "+hash+" not found")
		}
	}

	var confirmation = types.Confirmation{
		Creator:  msg.Creator,
		Height:   msg.Height,
		Root:     msg.Root,
		Hashes:   msg.Hashes,
		SigECDSA: msg.SigECDSA,
		SigEdDSA: msg.SigEdDSA,
	}

	k.SetConfirmation(
		ctx,
		confirmation,
	)

	return &types.MsgCreateConfirmationResponse{}, nil
}

func (k msgServer) UpdateConfirmation(goCtx context.Context, msg *types.MsgUpdateConfirmation) (*types.MsgUpdateConfirmationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetConfirmation(
		ctx,
		msg.Height,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var confirmation = types.Confirmation{
		Creator:  msg.Creator,
		Height:   msg.Height,
		Root:     msg.Root,
		Hashes:   msg.Hashes,
		SigECDSA: msg.SigECDSA,
		SigEdDSA: msg.SigEdDSA,
	}

	k.SetConfirmation(ctx, confirmation)

	return &types.MsgUpdateConfirmationResponse{}, nil
}

func (k msgServer) DeleteConfirmation(goCtx context.Context, msg *types.MsgDeleteConfirmation) (*types.MsgDeleteConfirmationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetConfirmation(
		ctx,
		msg.Height,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveConfirmation(
		ctx,
		msg.Height,
	)

	return &types.MsgDeleteConfirmationResponse{}, nil
}
