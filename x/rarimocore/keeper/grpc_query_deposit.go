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

func (k Keeper) DepositAll(c context.Context, req *types.QueryAllDepositRequest) (*types.QueryAllDepositResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var deposits []types.Deposit
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	depositStore := prefix.NewStore(store, types.KeyPrefix(types.DepositKeyPrefix))

	pageRes, err := query.Paginate(depositStore, req.Pagination, func(key []byte, value []byte) error {
		var deposit types.Deposit
		if err := k.cdc.Unmarshal(value, &deposit); err != nil {
			return err
		}

		deposits = append(deposits, deposit)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDepositResponse{Deposit: deposits, Pagination: pageRes}, nil
}

func (k Keeper) Deposit(c context.Context, req *types.QueryGetDepositRequest) (*types.QueryGetDepositResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetDeposit(
		ctx,
		req.Tx,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDepositResponse{Deposit: val}, nil
}
