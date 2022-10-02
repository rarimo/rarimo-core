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

	if _, ok := k.GetParams(ctx).Networks[msg.CurrentChain]; !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "network not found")
	}

	var info = types.Info{
		Creator: msg.Creator,
		Index:   msg.Index,
		Chains: map[string]*types.ChainInfo{
			msg.CurrentChain: &types.ChainInfo{
				TokenAddress: msg.CurrentAddress,
				TokenId:      msg.CurrentId,
				TokenType:    msg.CurrentType,
			},
		},
	}

	k.SetInfo(
		ctx,
		info,
	)

	var item = types.Item{
		TokenAddress: msg.CurrentAddress,
		TokenId:      msg.CurrentId,
		Chain:        msg.CurrentChain,
		Index:        msg.Index,
		Name:         msg.CurrentName,
		Symbol:       msg.CurrentSymbol,
		Uri:          msg.CurrentURI,
		Decimals:     msg.CurrentDecimals,
		Seed:         msg.CurrentSeed,
	}

	k.SetItem(ctx, item)

	return &types.MsgCreateInfoResponse{}, nil
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

	info, ok := k.GetInfo(
		ctx,
		msg.Index,
	)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	if _, ok := k.GetItem(ctx, msg.TokenAddress, msg.TokenId, msg.ChainName); !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != info.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if _, ok := k.GetParams(ctx).Networks[msg.ChainName]; !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "network not found")
	}

	info.Chains[msg.ChainName] = &types.ChainInfo{
		TokenAddress: msg.TokenAddress,
		TokenId:      msg.TokenId,
		TokenType:    msg.TokenType,
	}

	k.SetInfo(ctx, info)

	k.SetItem(ctx, types.Item{
		TokenAddress: msg.TokenAddress,
		TokenId:      msg.TokenId,
		Index:        info.Index,
		Chain:        msg.ChainName,
		Name:         msg.Name,
		Symbol:       msg.Symbol,
		Uri:          msg.Uri,
		Decimals:     msg.Decimals,
		Seed:         msg.Seed,
	})

	return &types.MsgAddChainResponse{}, nil
}
