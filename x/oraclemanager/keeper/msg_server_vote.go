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

	if err := k.rarimo.CreateVote(ctx, vote); err != nil {
		return nil, err
	}

	oracle.VotesCount = oracle.VotesCount + 1
	k.SetOracle(ctx, oracle)

	return &types.MsgVoteResponse{}, nil
}
