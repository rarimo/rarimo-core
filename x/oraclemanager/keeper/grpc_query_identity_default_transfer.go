package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/oraclemanager/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) IdentityDefaultTransfer(c context.Context, req *types.QueryGetIdentityDefaultTransferRequest) (*types.QueryGetIdentityDefaultTransferResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if err := req.Msg.ValidateBody(); err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)

	transfer, err := k.rarimo.GetIdentityDefaultTransfer(ctx, &req.Msg)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetIdentityDefaultTransferResponse{Transfer: *transfer}, nil
}
