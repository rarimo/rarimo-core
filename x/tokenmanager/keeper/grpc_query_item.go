package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ItemAll(c context.Context, req *types.QueryAllItemRequest) (*types.QueryAllItemResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var items []types.Item
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	itemStore := prefix.NewStore(store, types.KeyPrefix(types.ItemKeyPrefix))

	pageRes, err := query.Paginate(itemStore, req.Pagination, func(key []byte, value []byte) error {
		var item types.Item
		if err := k.cdc.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllItemResponse{Item: items, Pagination: pageRes}, nil
}

func (k Keeper) Item(c context.Context, req *types.QueryGetItemRequest) (*types.QueryGetItemResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetItem(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetItemResponse{Item: val}, nil
}

func (k Keeper) OnChainItemAll(c context.Context, req *types.QueryAllOnChainItemRequest) (*types.QueryAllOnChainItemResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var items []types.OnChainItem
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	itemStore := prefix.NewStore(store, types.KeyPrefix(types.OnChainItemKeyPrefix))

	pageRes, err := query.Paginate(itemStore, req.Pagination, func(key []byte, value []byte) error {
		var item types.OnChainItem
		if err := k.cdc.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOnChainItemResponse{Item: items, Pagination: pageRes}, nil
}

func (k Keeper) OnChainItem(c context.Context, req *types.QueryGetOnChainItemRequest) (*types.QueryGetOnChainItemResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetOnChainItem(
		ctx,
		&types.OnChainItemIndex{
			Chain:   req.Chain,
			Address: req.Address,
			TokenID: req.TokenID,
		},
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetOnChainItemResponse{Item: val}, nil
}

func (k Keeper) SeedAll(c context.Context, req *types.QueryAllSeedRequest) (*types.QueryAllSeedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var seeds []types.Seed
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	itemStore := prefix.NewStore(store, types.KeyPrefix(types.SeedKeyPrefix))

	pageRes, err := query.Paginate(itemStore, req.Pagination, func(key []byte, value []byte) error {
		var seed types.Seed
		if err := k.cdc.Unmarshal(value, &seed); err != nil {
			return err
		}

		seeds = append(seeds, seed)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSeedResponse{Seed: seeds, Pagination: pageRes}, nil
}

func (k Keeper) Seed(c context.Context, req *types.QueryGetSeedRequest) (*types.QueryGetSeedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetSeed(
		ctx,
		req.Seed,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetSeedResponse{Seed: val}, nil
}

func (k Keeper) ItemByOnChainItem(c context.Context, req *types.QueryGetItemByOnChainItemRequest) (*types.QueryGetItemByOnChainItemResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	onChainItem, found := k.GetOnChainItem(ctx, &types.OnChainItemIndex{Chain: req.Chain, Address: req.Address, TokenID: req.TokenID})
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	item, found := k.GetItem(ctx, onChainItem.Item)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetItemByOnChainItemResponse{Item: item}, nil
}

func (k Keeper) OnChainItemByOther(c context.Context, req *types.QueryGetOnChainItemByOtherRequest) (*types.QueryGetOnChainItemByOtherResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	onChainItem, found := k.GetOnChainItem(ctx, &types.OnChainItemIndex{Chain: req.Chain, Address: req.Address, TokenID: req.TokenID})
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	item, found := k.GetItem(ctx, onChainItem.Item)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	for _, index := range item.OnChain {
		if index.Chain == req.TargetChain {
			onChainItem, _ = k.GetOnChainItem(ctx, index)
			return &types.QueryGetOnChainItemByOtherResponse{Item: onChainItem}, nil
		}
	}

	return nil, status.Error(codes.NotFound, "not found")
}
