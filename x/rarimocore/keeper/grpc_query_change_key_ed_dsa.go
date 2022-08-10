package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ChangeKeyEdDSAAll(c context.Context, req *types.QueryAllChangeKeyEdDSARequest) (*types.QueryAllChangeKeyEdDSAResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var changeKeyEdDSAs []types.ChangeKeyEdDSA
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	changeKeyEdDSAStore := prefix.NewStore(store, types.KeyPrefix(types.ChangeKeyEdDSAKeyPrefix))

	pageRes, err := query.Paginate(changeKeyEdDSAStore, req.Pagination, func(key []byte, value []byte) error {
		var changeKeyEdDSA types.ChangeKeyEdDSA
		if err := k.cdc.Unmarshal(value, &changeKeyEdDSA); err != nil {
			return err
		}

		changeKeyEdDSAs = append(changeKeyEdDSAs, changeKeyEdDSA)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllChangeKeyEdDSAResponse{ChangeKeyEdDSA: changeKeyEdDSAs, Pagination: pageRes}, nil
}

func (k Keeper) ChangeKeyEdDSA(c context.Context, req *types.QueryGetChangeKeyEdDSARequest) (*types.QueryGetChangeKeyEdDSAResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetChangeKeyEdDSA(
		ctx,
		req.NewKey,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetChangeKeyEdDSAResponse{ChangeKeyEdDSA: val}, nil
}
