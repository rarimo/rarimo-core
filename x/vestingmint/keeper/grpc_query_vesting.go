package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/rarimo/rarimo-core/x/vestingmint/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Vesting(c context.Context, req *types.QueryGetVestingRequest) (*types.QueryGetVestingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVesting(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetVestingResponse{Vesting: val}, nil
}

func (k Keeper) VestingAll(c context.Context, req *types.QueryAllVestingRequest) (*types.QueryAllVestingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	var vestings []types.Vesting

	store := ctx.KVStore(k.storeKey)
	confirmationStore := prefix.NewStore(store, types.KeyPrefix(types.VestingKeyPrefix))

	pageRes, err := query.Paginate(confirmationStore, req.Pagination, func(key []byte, value []byte) error {
		var v types.Vesting
		if err := k.cdc.Unmarshal(value, &v); err != nil {
			return err
		}

		vestings = append(vestings, v)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVestingResponse{Vesting: vestings, Pagination: pageRes}, nil
}
