package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/types"
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
		req.TokenAddress,
		req.TokenId,
		req.Chain,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetItemResponse{Item: val}, nil
}
