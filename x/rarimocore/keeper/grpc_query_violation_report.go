package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/rarimo/rarimo-core/x/rarimocore/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ViolationReportAll(c context.Context, req *types.QueryAllViolationReportRequest) (*types.QueryAllViolationReportResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var reports []types.ViolationReport
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	confirmationStore := prefix.NewStore(store, types.KeyPrefix(types.ViolationReportKeyPrefix))

	pageRes, err := query.Paginate(confirmationStore, req.Pagination, func(key []byte, value []byte) error {
		var report types.ViolationReport
		if err := k.cdc.Unmarshal(value, &report); err != nil {
			return err
		}

		reports = append(reports, report)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllViolationReportResponse{ViolationReport: reports, Pagination: pageRes}, nil
}

func (k Keeper) ViolationReport(c context.Context, req *types.QueryGetViolationReportRequest) (*types.QueryGetViolationReportResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetViolationReport(
		ctx,
		types.NewViolationReportIndex(req.SessionId, req.Offender, req.Sender, req.ViolationType),
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetViolationReportResponse{ViolationReport: val}, nil
}
