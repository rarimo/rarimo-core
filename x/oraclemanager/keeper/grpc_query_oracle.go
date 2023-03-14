package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"gitlab.com/rarimo/rarimo-core/x/oraclemanager/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OracleAll(c context.Context, req *types.QueryGetAllOracleRequest) (*types.QueryGetAllOracleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	var oracles []types.Oracle

	store := ctx.KVStore(k.storeKey)
	confirmationStore := prefix.NewStore(store, types.KeyPrefix(types.OracleKeyPrefix))

	pageRes, err := query.Paginate(confirmationStore, req.Pagination, func(key []byte, value []byte) error {
		var oracle types.Oracle
		if err := k.cdc.Unmarshal(value, &oracle); err != nil {
			return err
		}

		oracles = append(oracles, oracle)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetAllOracleResponse{Oracle: oracles, Pagination: pageRes}, nil
}

func (k Keeper) Oracle(c context.Context, req *types.QueryGetOracleRequest) (*types.QueryGetOracleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetOracle(
		ctx,
		&types.OracleIndex{
			Chain:   req.Chain,
			Account: req.Address,
		},
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetOracleResponse{Oracle: val}, nil
}

func (k Keeper) OracleForChain(c context.Context, req *types.QueryGetOracleForChainRequest) (*types.QueryGetOracleForChainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	oracles := k.GetOracleForChain(ctx, req.Chain)

	return &types.QueryGetOracleForChainResponse{Oracle: oracles}, nil
}
