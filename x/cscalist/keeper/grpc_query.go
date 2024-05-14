package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/rarimo/rarimo-core/x/cscalist/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(goCtx context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

func (k Keeper) Tree(c context.Context, req *types.QueryTreeRequest) (*types.QueryTreeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// prevent extreme memory usage on bigger numbers, while the storage shouldn't
	// contain more than that
	if req.Pagination.Limit > 600 {
		req.Pagination.Limit = 600
	}

	var (
		ctx   = sdk.UnwrapSDKContext(c)
		nodes = make([]types.Node, 0, req.Pagination.Limit)
		store = prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NodeKeyPrefix))
	)

	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var node types.Node
		if err := k.cdc.Unmarshal(value, &node); err != nil {
			return err
		}
		nodes = append(nodes, node)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryTreeResponse{Tree: nodes, Pagination: pageRes}, nil
}
