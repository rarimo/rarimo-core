package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	merkle "gitlab.com/rarimo/go-merkle"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OperationAll(c context.Context, req *types.QueryAllOperationRequest) (*types.QueryAllOperationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var operations []types.Operation
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	depositStore := prefix.NewStore(store, types.KeyPrefix(types.OperationKeyPrefix))

	pageRes, err := query.Paginate(depositStore, req.Pagination, func(key []byte, value []byte) error {
		var operation types.Operation
		if err := k.cdc.Unmarshal(value, &operation); err != nil {
			return err
		}

		operations = append(operations, operation)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOperationResponse{Operation: operations, Pagination: pageRes}, nil
}

func (k Keeper) Operation(c context.Context, req *types.QueryGetOperationRequest) (*types.QueryGetOperationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetOperation(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetOperationResponse{Operation: val}, nil
}

func (k Keeper) OperationProof(c context.Context, req *types.QueryGetOperationProofRequest) (*types.QueryGetOperationProofResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	operation, ok := k.GetOperation(ctx, req.Index)
	if !ok {
		return nil, status.Error(codes.NotFound, "not found")
	}

	if operation.Status != types.OpStatus_SIGNED {
		return nil, status.Error(codes.InvalidArgument, "operation is no signed")
	}

	confirmationId, ok := k.GetOperationConfirmationId(ctx, operation.Index)
	if !ok {
		return nil, status.Error(codes.NotFound, "confirmation id not found")
	}

	confirmation, ok := k.GetConfirmation(ctx, confirmationId)
	if !ok {
		return nil, status.Error(codes.NotFound, "confirmation not found")
	}

	contents := make([]merkle.Content, 0, len(confirmation.Indexes))

	for _, op := range confirmation.Indexes {
		operation, ok := k.GetOperation(ctx, op)
		if !ok {
			return nil, status.Error(codes.NotFound, "operation from set not found")
		}

		if operation.Status != types.OpStatus_SIGNED {
			return nil, status.Error(codes.InvalidArgument, "one of operations from set is no signed")
		}

		content, err := k.getContent(ctx, operation)
		if err != nil {
			return nil, sdkerrors.Wrap(err, "failed to get content")
		}

		contents = append(contents, content)
	}

	targetContent, err := k.getContent(ctx, operation)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to get target content")
	}

	t := merkle.NewTree(crypto.Keccak256, contents...)

	path, ok := t.Path(targetContent)
	if !ok {
		return nil, status.Error(codes.Internal, "can not find path")
	}

	resp := types.QueryGetOperationProofResponse{
		Path:      make([]string, 0, len(path)),
		Signature: confirmation.SignatureECDSA,
	}

	for _, p := range path {
		resp.Path = append(resp.Path, hexutil.Encode(p))
	}

	return &resp, nil
}
