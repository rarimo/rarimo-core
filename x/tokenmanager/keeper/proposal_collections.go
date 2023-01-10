package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k Keeper) HandleCreateCollectionProposal(
	ctx sdk.Context,
	proposal *types.CreateCollectionProposal,
) error {
	collection := k.GetCollectionInfo(ctx, proposal.Index)
	if collection != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("collection with index %s already exists", proposal.Index))
	}

	meta := proposal.Metadata
	if meta == nil {
		meta = &types.CollectionMetadata{}
	}

	k.PutCollectionInfo(ctx, types.CollectionInfo{
		Index: proposal.Index,
		ChainParams: map[string]*types.CollectionChainParams{
			proposal.Network: {
				Address:  proposal.Address,
				Decimals: proposal.Decimals,
			},
		},
		Items:     make(map[string]string),
		TokenType: proposal.TokenType,
		Metadata:  meta,
	})

	return nil
}

func (k Keeper) HandlePutCollectionNetworkAddressProposal(
	ctx sdk.Context,
	proposal *types.PutCollectionNetworkAddressProposal,
) error {
	collection := k.GetCollectionInfo(ctx, proposal.Index)
	if collection == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("collection with index %s not found", proposal.Index))
	}

	collection.ChainParams[proposal.Network] = &types.CollectionChainParams{
		Address:  proposal.Address,
		Decimals: proposal.Decimals,
	}

	k.PutCollectionInfo(ctx, *collection)

	return nil
}

func (k Keeper) HandleRemoveCollectionNetworkAddressProposal(
	ctx sdk.Context,
	proposal *types.RemoveCollectionNetworkAddressProposal,
) error {
	collection := k.GetCollectionInfo(ctx, proposal.Index)
	if collection == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("collection with index %s not found", proposal.Index))
	}

	delete(collection.ChainParams, proposal.Network)

	k.PutCollectionInfo(ctx, *collection)

	return nil
}

func (k Keeper) HandleRemoveCollectionProposal(
	ctx sdk.Context,
	proposal *types.RemoveCollectionProposal,
) error {
	collection := k.GetCollectionInfo(ctx, proposal.Index)
	if collection == nil {
		return nil
	}

	k.RemoveCollectionInfo(ctx, proposal.Index)
	return nil
}
