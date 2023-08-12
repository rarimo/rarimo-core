package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"gitlab.com/rarimo/rarimo-core/x/identity/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

func (k Keeper) Node(c context.Context, req *types.QueryGetNodeRequest) (*types.QueryGetNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNode(
		ctx,
		req.Key,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNodeResponse{Node: val}, nil
}

func (k Keeper) NodeAll(c context.Context, req *types.QueryAllNodeRequest) (*types.QueryAllNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nodes []types.Node
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeKeyPrefix))

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

	return &types.QueryAllNodeResponse{Node: nodes, Pagination: pageRes}, nil
}

func (k Keeper) MerkleProof(c context.Context, req *types.QueryGetMerkleProofRequest) (*types.QueryGetMerkleProofResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	path := k.Path(ctx, req.Id)
	resp := &types.QueryGetMerkleProofResponse{Proof: make([]string, 0, len(path))}

	for i := len(path) - 1; i >= 0; i-- {
		if path[i] != EmptyHashStr {
			resp.Proof = append(resp.Proof, path[i])
		}
	}

	return resp, nil
}

func (k Keeper) State(c context.Context, req *types.QueryGetStateRequest) (*types.QueryGetStateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetStateInfo(
		ctx,
		req.Id,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStateResponse{State: val}, nil
}

func (k Keeper) StateAll(c context.Context, req *types.QueryAllStateRequest) (*types.QueryAllStateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	var states []types.StateInfo

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StateInfoKeyPrefix))

	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var state types.StateInfo
		if err := k.cdc.Unmarshal(value, &state); err != nil {
			return err
		}

		states = append(states, state)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStateResponse{State: states, Pagination: pageRes}, nil
}
