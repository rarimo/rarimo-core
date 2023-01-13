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

func (k Keeper) CollectionByNetworkParams(ctx context.Context,
	r *types.QueryGetCollectionByNetworkParamsRequest,
) (*types.QueryGetCollectionByNetworkParamsResponse, error) {
	if r == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	val := k.GetCollectionInfoByNetworkParams(sdkCtx, r.Network, r.Address)
	if val == nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCollectionByNetworkParamsResponse{Collection: *val}, nil
}

func (k Keeper) Collection(ctx context.Context, r *types.QueryGetCollectionRequest) (*types.QueryGetCollectionResponse, error) {
	if r == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	val := k.GetCollectionInfo(sdkCtx, r.Index)
	if val == nil {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCollectionResponse{Collection: *val}, nil
}

func (k Keeper) CollectionAll(ctx context.Context, r *types.QueryAllCollectionRequest) (*types.QueryAllCollectionResponse, error) {
	if r == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	var collections []types.CollectionInfo

	store := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.KeyPrefix(types.CollectionKeyPrefix))

	pageRes, err := query.Paginate(store, r.Pagination, func(key []byte, value []byte) error {
		var collection types.CollectionInfo
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
