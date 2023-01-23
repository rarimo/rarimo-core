package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k Keeper) HandleCreateCollectionProposal(ctx sdk.Context, proposal *types.CreateCollectionProposal) error {
	if _, ok := k.GetCollection(ctx, proposal.Index); ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("collection with index %s already exists", proposal.Index))
	}

	k.PutCollection(ctx, types.Collection{
		Index: proposal.Index,
		Meta:  proposal.Metadata,
	})

	for _, data := range proposal.Data {
		k.PutCollectionData(ctx, *data)
	}

	return nil
}

func (k Keeper) HandleAddCollectionDataProposal(ctx sdk.Context, proposal *types.AddCollectionDataProposal) error {
	for _, data := range proposal.Data {
		if _, ok := k.GetCollection(ctx, data.Index.Collection); !ok {
			return sdkerrors.Wrap(sdkerrors.ErrNotFound, "not found")
		}

		if _, ok := k.GetCollectionData(ctx, data.Index); ok {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("data with index %s already exists", data.Index))
		}

		k.PutCollectionData(ctx, *data)
	}

	return nil
}

func (k Keeper) HandleUpdateCollectionDataProposal(ctx sdk.Context, proposal *types.UpdateCollectionDataProposal) error {
	for _, data := range proposal.Data {
		if _, ok := k.GetCollectionData(ctx, data.Index); !ok {
			return sdkerrors.Wrap(sdkerrors.ErrNotFound, "not found")
		}

		k.PutCollectionData(ctx, *data)
	}

	return nil
}

func (k Keeper) HandleRemoveCollectionDataProposal(ctx sdk.Context, proposal *types.RemoveCollectionDataProposal) error {
	for _, index := range proposal.Index {
		k.RemoveCollectionData(ctx, index)
	}

	return nil
}

func (k Keeper) HandleRemoveCollectionProposal(ctx sdk.Context, proposal *types.RemoveCollectionProposal) error {
	k.RemoveCollection(ctx, proposal.Index)
	return nil
}
