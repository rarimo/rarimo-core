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

func (k Keeper) ChangeKeyECDSAAll(c context.Context, req *types.QueryAllChangeKeyECDSARequest) (*types.QueryAllChangeKeyECDSAResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var changeKeyECDSAs []types.ChangeKeyECDSA
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	changeKeyECDSAStore := prefix.NewStore(store, types.KeyPrefix(types.ChangeKeyECDSAKeyPrefix))

	pageRes, err := query.Paginate(changeKeyECDSAStore, req.Pagination, func(key []byte, value []byte) error {
		var changeKeyECDSA types.ChangeKeyECDSA
		if err := k.cdc.Unmarshal(value, &changeKeyECDSA); err != nil {
			return err
		}

		changeKeyECDSAs = append(changeKeyECDSAs, changeKeyECDSA)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllChangeKeyECDSAResponse{ChangeKeyECDSA: changeKeyECDSAs, Pagination: pageRes}, nil
}

func (k Keeper) ChangeKeyECDSA(c context.Context, req *types.QueryGetChangeKeyECDSARequest) (*types.QueryGetChangeKeyECDSAResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetChangeKeyECDSA(
		ctx,
		req.NewKey,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetChangeKeyECDSAResponse{ChangeKeyECDSA: val}, nil
}
