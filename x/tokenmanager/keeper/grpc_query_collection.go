package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k Keeper) Collection(ctx context.Context, r *types.QueryGetCollectionRequest) (*types.QueryGetCollectionResponse, error) {
	if r == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	val, ok := k.GetCollection(sdkCtx, r.Index)
	if !ok {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCollectionResponse{Collection: val}, nil
}

func (k Keeper) CollectionAll(ctx context.Context, r *types.QueryAllCollectionRequest) (*types.QueryAllCollectionResponse, error) {
	if r == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	var collections []types.Collection

	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionKeyPrefix))

	pageRes, err := query.Paginate(store, r.Pagination, func(key []byte, value []byte) error {
		var collection types.Collection
		if err := k.cdc.Unmarshal(value, &collection); err != nil {
			return err
		}

		collections = append(collections, collection)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCollectionResponse{Collection: collections, Pagination: pageRes}, nil
}

func (k Keeper) CollectionData(ctx context.Context, r *types.QueryGetCollectionDataRequest) (*types.QueryGetCollectionDataResponse, error) {
	if r == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	val, ok := k.GetCollectionData(sdkCtx, &types.CollectionDataIndex{Chain: r.Chain, Address: r.Address})
	if !ok {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCollectionDataResponse{Data: val}, nil
}

func (k Keeper) CollectionDataAll(ctx context.Context, r *types.QueryAllCollectionDataRequest) (*types.QueryAllCollectionDataResponse, error) {
	if r == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	var data []types.CollectionData

	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionDataKeyPrefix))

	pageRes, err := query.Paginate(store, r.Pagination, func(key []byte, value []byte) error {
		var d types.CollectionData
		if err := k.cdc.Unmarshal(value, &d); err != nil {
			return err
		}

		data = append(data, d)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCollectionDataResponse{Data: data, Pagination: pageRes}, nil
}

func (k Keeper) CollectionByCollectionData(c context.Context, req *types.QueryGetCollectionByCollectionDataRequest) (*types.QueryGetCollectionByCollectionDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	data, found := k.GetCollectionData(ctx, &types.CollectionDataIndex{Chain: req.Chain, Address: req.Address})
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	collection, found := k.GetCollection(ctx, data.Collection)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCollectionByCollectionDataResponse{Collection: collection}, nil
}
