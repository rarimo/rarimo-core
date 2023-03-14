package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"gitlab.com/rarimo/rarimo-core/x/bridge/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) HashAll(goCtx context.Context, req *types.QueryAllHashRequest) (*types.QueryAllHashResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var hashes []types.Hash

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	hashStore := prefix.NewStore(store, types.KeyPrefix(types.HashKeyPrefix))

	pageRes, err := query.Paginate(hashStore, req.Pagination, func(key []byte, value []byte) error {
		var hash types.Hash
		if err := k.cdc.Unmarshal(value, &hash); err != nil {
			return err
		}

		hashes = append(hashes, hash)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllHashResponse{Hash: hashes, Pagination: pageRes}, nil
}

func (k Keeper) Hash(goCtx context.Context, req *types.QueryGetHashRequest) (*types.QueryGetHashResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetHash(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetHashResponse{Hash: val}, nil
}
