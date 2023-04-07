package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k Keeper) HandleCreateCollectionProposal(ctx sdk.Context, proposal *types.CreateCollectionProposal) error {
	if _, ok := k.GetCollection(ctx, proposal.Index); ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("collection with index %s already exists", proposal.Index))
	}

	col := types.Collection{
		Index: proposal.Index,
		Meta:  proposal.Metadata,
		Data:  make([]*types.CollectionDataIndex, 0, len(proposal.Data)),
	}

	for _, data := range proposal.Data {
		col.Data = append(col.Data, data.Index)
	}

	k.SetCollection(ctx, col)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeCollectionCreated,
		sdk.NewAttribute(types.AttributeKeyCollectionIndex, proposal.Index),
	))

	for _, data := range proposal.Data {
		k.SetCollectionData(ctx, *data)
		ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeCollectionDataCreated,
			sdk.NewAttribute(types.AttributeKeyCollectionIndex, proposal.Index),
			sdk.NewAttribute(types.AttributeKeyCollectionDataChain, data.Index.Chain),
		))
	}

	for _, item := range proposal.Item {
		k.SetItem(ctx, *item)
		// TODO should we check item.Collection == proposal.Index?
		ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeItemCreated,
			sdk.NewAttribute(types.AttributeKeyItemIndex, item.Index),
		))
	}

	for _, item := range proposal.OnChainItem {
		k.SetOnChainItem(ctx, *item)
		ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeOnChainItemCreated,
			sdk.NewAttribute(types.AttributeKeyItemIndex, item.Item),
			sdk.NewAttribute(types.AttributeKeyOnChainItemChain, item.Index.Chain),
		))
	}

	return nil
}

func (k Keeper) HandleAddCollectionDataProposal(ctx sdk.Context, proposal *types.AddCollectionDataProposal) error {
	for _, data := range proposal.Data {
		col, ok := k.GetCollection(ctx, data.Collection)
		if !ok {
			return sdkerrors.Wrap(sdkerrors.ErrNotFound, "not found")
		}

		col.Data = append(col.Data, data.Index)
		k.SetCollection(ctx, col)

		if _, ok := k.GetCollectionData(ctx, data.Index); ok {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("data with index %s already exists", data.Index))
		}

		k.SetCollectionData(ctx, *data)
		ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeCollectionDataCreated,
			sdk.NewAttribute(types.AttributeKeyCollectionIndex, data.Collection),
			sdk.NewAttribute(types.AttributeKeyCollectionDataChain, data.Index.Chain),
		))
	}

	return nil
}

func (k Keeper) HandleUpdateCollectionDataProposal(ctx sdk.Context, proposal *types.UpdateCollectionDataProposal) error {
	for _, data := range proposal.Data {
		if _, ok := k.GetCollectionData(ctx, data.Index); !ok {
			return sdkerrors.Wrap(sdkerrors.ErrNotFound, "not found")
		}

		k.SetCollectionData(ctx, *data)

		ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeCollectionDataUpdated,
			sdk.NewAttribute(types.AttributeKeyCollectionIndex, data.Collection),
			sdk.NewAttribute(types.AttributeKeyCollectionDataChain, data.Index.Chain),
		))
	}

	return nil
}

func (k Keeper) HandleRemoveCollectionDataProposal(ctx sdk.Context, proposal *types.RemoveCollectionDataProposal) error {
	for _, index := range proposal.Index {
		data, ok := k.GetCollectionData(ctx, index)
		if !ok {
			return sdkerrors.Wrap(sdkerrors.ErrNotFound, "not found")
		}

		col, ok := k.GetCollection(ctx, data.Collection)
		if !ok {
			return sdkerrors.Wrap(sdkerrors.ErrNotFound, "not found")
		}

		// removing data index from collection
		for i, colDataIndex := range col.Data {
			if proto.Equal(colDataIndex, index) {
				if i == len(col.Data)-1 {
					col.Data = col.Data[:i]
					break
				}

				col.Data = append(col.Data[:i], col.Data[i+1:]...)
			}
		}

		k.SetCollection(ctx, col)
		k.RemoveCollectionData(ctx, index)

		ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeCollectionDataRemoved,
			sdk.NewAttribute(types.AttributeKeyCollectionIndex, data.Collection),
			sdk.NewAttribute(types.AttributeKeyCollectionDataChain, index.Chain),
		))
	}

	return nil
}

func (k Keeper) HandleRemoveCollectionProposal(ctx sdk.Context, proposal *types.RemoveCollectionProposal) error {
	col, ok := k.GetCollection(ctx, proposal.Index)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "not found")
	}

	for _, index := range col.Data {
		k.RemoveCollectionData(ctx, index)

		ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeCollectionDataRemoved,
			sdk.NewAttribute(types.AttributeKeyCollectionIndex, proposal.Index),
			sdk.NewAttribute(types.AttributeKeyCollectionDataChain, index.Chain),
		))
	}

	k.RemoveCollection(ctx, proposal.Index)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeCollectionRemoved,
		sdk.NewAttribute(types.AttributeKeyCollectionIndex, proposal.Index),
	))

	return nil
}
