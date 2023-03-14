package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/oraclemanager/types"
	rarimotypes "gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k msgServer) Vote(goCtx context.Context, msg *types.MsgVote) (*types.MsgVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	defer k.disableFee(ctx.GasMeter().GasConsumed(), ctx.GasMeter())

	// REQUIRES: validated index
	oracle, ok := k.GetOracle(ctx, msg.Index)
	if !ok {
		return nil, types.ErrOracleNotFound
	}

	if oracle.Status != types.OracleStatus_Active {
		return nil, types.ErrInactiveOracle
	}

	vote := rarimotypes.Vote{
		Index: &rarimotypes.VoteIndex{
			Operation: msg.Operation,
			Validator: msg.Index.Account,
		},
		Vote: msg.Vote,
	}

	firstUpdate, err := k.rarimo.CreateVote(ctx, vote)
	if err != nil {
		return nil, err
	}

	if firstUpdate {
		params := k.GetParams(ctx)
		k.AddToMonitorQueue(ctx, uint64(ctx.BlockHeight())+params.CheckOperationDelta, msg.Operation)
	}

	oracle.VotesCount = oracle.VotesCount + 1
	k.SetOracle(ctx, oracle)

	return &types.MsgVoteResponse{}, nil
}
