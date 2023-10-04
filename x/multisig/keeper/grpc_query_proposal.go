package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/rarimo/rarimo-core/x/multisig/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ProposalAll(goCtx context.Context, req *types.QueryAllProposalRequest) (*types.QueryAllProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var proposals []types.Proposal

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	proposalStore := prefix.NewStore(store, types.KeyPrefix(types.ProposalKeyPrefix))

	pageRes, err := query.Paginate(proposalStore, req.Pagination, func(key []byte, value []byte) error {
		var proposal types.Proposal
		if err := k.cdc.Unmarshal(value, &proposal); err != nil {
			return err
		}

		proposals = append(proposals, proposal)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProposalResponse{Proposal: proposals, Pagination: pageRes}, nil
}

func (k Keeper) Proposal(goCtx context.Context, req *types.QueryGetProposalRequest) (*types.QueryGetProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetProposal(ctx, req.ProposalId)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetProposalResponse{Proposal: val}, nil
}
