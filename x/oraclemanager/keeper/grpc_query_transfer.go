package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/oraclemanager/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Transfer(c context.Context, req *types.QueryGetTransferRequest) (*types.QueryGetTransferResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if err := req.Msg.ValidateBody(); err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)

	transfer, err := k.rarimo.GetTransfer(ctx, &req.Msg)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetTransferResponse{Transfer: *transfer}, nil
}
