package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k Keeper) HandleCreateTokenItemProposal(ctx sdk.Context, proposal *types.CreateTokenItemProposal) error {
	collection := k.GetCollectionInfo(ctx, proposal.Item.Collection)
	if collection == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("collection with index %s not found", proposal.Item.Collection))
	}

	if collection.Items == nil {
		collection.Items = make(map[string]string)
	}

	keys := make([]string, 0, len(proposal.Item.ChainParams))

	for chain, params := range proposal.Item.ChainParams {
		collectionChainParams, ok := collection.ChainParams[chain]
		if !ok {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("collection %s does not support network %s", collection.Index, chain))
		}
		itemKey := types.ItemKey(collectionChainParams.Address, params.TokenID, chain)

		keys = append(keys, string(itemKey))

		collection.Items[string(itemKey)] = proposal.Item.Index
	}

	k.PutItem(ctx, *proposal.Item, keys...)

	k.PutCollectionInfo(ctx, *collection)

	return nil
}

func (k Keeper) RemoveTokenItemProposal(ctx sdk.Context, proposal *types.RemoveTokenItemProposal) error {
	item, ok := k.GetItem(ctx, proposal.TokenAddress, proposal.TokenId, proposal.Network)
	if !ok {
		return nil // TODO return error?
	}

	itemKey := types.ItemKey(proposal.TokenAddress, proposal.TokenId, proposal.Network)

	collection := k.GetCollectionInfo(ctx, item.Collection)
	if collection == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("collection with index %s not found", item.Collection))
	}

	delete(collection.Items, string(itemKey))

	k.PutCollectionInfo(ctx, *collection)

	k.RemoveItem(ctx, proposal.TokenAddress, proposal.TokenId, proposal.Network)
	return nil
}

func (k Keeper) SetNetworkProposal(ctx sdk.Context, proposal *types.SetNetworkProposal) error {
	params := k.GetParams(ctx)
	for i, network := range params.Networks {
		if network.Name == proposal.NetworkParams.Name {
			params.Networks[i] = proposal.NetworkParams
			k.SetParams(ctx, params)
			return nil
		}
	}

	params.Networks = append(params.Networks, proposal.NetworkParams)
	k.SetParams(ctx, params)
	return nil
}
