package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/types"
)

func (k msgServer) CreateInfo(goCtx context.Context, msg *types.MsgCreateInfo) (*types.MsgCreateInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetInfo(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var info = types.Info{
		Creator:      msg.Creator,
		Index:        msg.Index,
		Name:         msg.Name,
		Symbol:       msg.Symbol,
		Image:        msg.Image,
		Description:  msg.Description,
		AnimationUrl: msg.AnimationUrl,
		ExternalUrl:  msg.ExternalUrl,
		Attributes:   msg.Attributes,
		CurrentChain: msg.CurrentChain,
		Chains:       make(map[string]*types.ChainInfo),
	}

	k.SetInfo(
		ctx,
		info,
	)
	return &types.MsgCreateInfoResponse{}, nil
}

func (k msgServer) UpdateInfo(goCtx context.Context, msg *types.MsgUpdateInfo) (*types.MsgUpdateInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetInfo(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var info = types.Info{
		Creator:      msg.Creator,
		Index:        msg.Index,
		Name:         msg.Name,
		Symbol:       msg.Symbol,
		Image:        msg.Image,
		Description:  msg.Description,
		AnimationUrl: msg.AnimationUrl,
		ExternalUrl:  msg.ExternalUrl,
		Attributes:   msg.Attributes,
		CurrentChain: msg.CurrentChain,
		Chains:       valFound.Chains,
	}

	k.SetInfo(ctx, info)

	return &types.MsgUpdateInfoResponse{}, nil
}

func (k msgServer) DeleteInfo(goCtx context.Context, msg *types.MsgDeleteInfo) (*types.MsgDeleteInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetInfo(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveInfo(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteInfoResponse{}, nil
}

func (k msgServer) AddChain(goCtx context.Context, msg *types.MsgAddChain) (*types.MsgAddChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetInfo(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	valFound.Chains[msg.Chain] = &types.ChainInfo{
		TokenAddress: msg.TokenAddress,
		TokenId:      msg.TokenId,
	}

	k.SetInfo(ctx, valFound)

	return &types.MsgAddChainResponse{}, nil
}
