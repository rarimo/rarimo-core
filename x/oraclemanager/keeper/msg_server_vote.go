package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/oraclemanager/types"
	rarimotypes "github.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k msgServer) Vote(goCtx context.Context, msg *types.MsgVote) (*types.MsgVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	defer k.disableFee(ctx.GasMeter().GasConsumed(), ctx.GasMeter())

	// REQUIRES: validated index
	oracle, ok := k.GetOracle(ctx, msg.Index)
	if !ok {
		return nil, types.ErrOracleNotFound
	}

	// Only active oracle can vote
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

	// Votes are stored in rarimocore module.
	voteCreated, err := k.rarimo.CreateVote(ctx, vote)
	if err != nil {
		return nil, err
	}

	// If vote was successfully created - then we have to recalculate votes to check if results can be evaluated.
	if voteCreated {
		if err := k.collectVotes(ctx, msg.Operation); err != nil {
			return nil, err
		}
	}

	oracle.VotesCount = oracle.VotesCount + 1
	k.SetOracle(ctx, oracle)
	return &types.MsgVoteResponse{}, nil
}

func (k Keeper) collectVotes(ctx sdk.Context, index string) error {
	operation, _ := k.rarimo.GetOperation(ctx, index)
	// Votes colelcting can be different for different types of operation.
	switch operation.OperationType {
	case rarimotypes.OpType_TRANSFER, rarimotypes.OpType_IDENTITY_DEFAULT_TRANSFER, rarimotypes.OpType_IDENTITY_GIST_TRANSFER:
		return k.collectOperationVotes(ctx, operation)
	default:
		// Nothing to do
	}
	return nil
}

func (k Keeper) collectOperationVotes(ctx sdk.Context, operation rarimotypes.Operation) error {
	chain, err := getSourceChain(operation)
	if err != nil {
		return err
	}

	// Calculating results depending on voting power. Oracle voting power is equal to amount of staked tokens.

	yesResult := sdk.ZeroInt()
	noResult := sdk.ZeroInt()
	totalPowerForChain := sdk.ZeroInt()

	for _, oracle := range k.GetOracleForChain(ctx, chain) {
		if oracle.Status != types.OracleStatus_Active {
			continue
		}

		oracleVotingPower, _ := sdk.NewIntFromString(oracle.Stake)
		totalPowerForChain = totalPowerForChain.Add(oracleVotingPower)

		vote, ok := k.rarimo.GetVote(ctx, &rarimotypes.VoteIndex{Validator: oracle.Index.Account, Operation: operation.Index})
		if ok {
			switch vote.Vote {
			case rarimotypes.VoteType_YES:
				yesResult = yesResult.Add(oracleVotingPower)
			case rarimotypes.VoteType_NO:
				noResult = noResult.Add(oracleVotingPower)
			}
		}
	}

	totalVotingPower := yesResult.Add(noResult)

	params := k.GetParams(ctx)
	quorum, _ := sdk.NewDecFromStr(params.VoteQuorum)
	threshold, _ := sdk.NewDecFromStr(params.VoteThreshold)

	// If there is not enough quorum of votes, finish the flow
	percentVoting := sdk.NewDecFromInt(totalVotingPower).Quo(sdk.NewDecFromInt(totalPowerForChain))
	if percentVoting.LT(quorum) {
		return nil
	}

	if operation.Status == rarimotypes.OpStatus_INITIALIZED {
		// Firstly moving from initialized to approved/unapproved status
		k.AddToMonitorQueue(ctx, uint64(ctx.BlockHeight())+params.CheckOperationDelta, operation.Index)
	}

	// After successful quorum check - check the threshold for yes votes.
	if sdk.NewDecFromInt(yesResult).Quo(sdk.NewDecFromInt(totalVotingPower)).GT(threshold) {
		return k.rarimo.ApproveOperation(ctx, operation)
	}

	return k.rarimo.UnapproveOperation(ctx, operation)
}
