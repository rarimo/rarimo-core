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

func (k Keeper) ConfirmationAll(c context.Context, req *types.QueryAllConfirmationRequest) (*types.QueryAllConfirmationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var confirmations []types.Confirmation
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	confirmationStore := prefix.NewStore(store, types.KeyPrefix(types.ConfirmationKeyPrefix))

	pageRes, err := query.Paginate(confirmationStore, req.Pagination, func(key []byte, value []byte) error {
		var confirmation types.Confirmation
		if err := k.cdc.Unmarshal(value, &confirmation); err != nil {
			return err
		}

		confirmations = append(confirmations, confirmation)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllConfirmationResponse{Confirmation: confirmations, Pagination: pageRes}, nil
}

func (k Keeper) Confirmation(c context.Context, req *types.QueryGetConfirmationRequest) (*types.QueryGetConfirmationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetConfirmation(
		ctx,
		req.Root,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetConfirmationResponse{Confirmation: val}, nil
}
