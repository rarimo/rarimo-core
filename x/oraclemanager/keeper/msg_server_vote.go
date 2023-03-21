package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/oraclemanager/types"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/pkg"
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

	voteCreated, err := k.rarimo.CreateVote(ctx, vote)
	if err != nil {
		return nil, err
	}

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
	switch operation.OperationType {
	case rarimotypes.OpType_TRANSFER:
		return k.collectTransferVotes(ctx, operation)
	default:
		// Nothing to do
	}
	return nil
}

func (k Keeper) collectTransferVotes(ctx sdk.Context, operation rarimotypes.Operation) error {
	transfer, _ := pkg.GetTransfer(operation) // error handled before
	yesResult := sdk.ZeroInt()
	noResult := sdk.ZeroInt()
	totalPowerForChain := sdk.ZeroInt()

	for _, oracle := range k.GetOracleForChain(ctx, transfer.From.Chain) {
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
	percentVoting := totalVotingPower.ToDec().Quo(totalPowerForChain.ToDec())
	if percentVoting.LT(quorum) {
		return nil
	}

	if operation.Status == rarimotypes.OpStatus_INITIALIZED {
		// Firstly moving from initialized to approved/unapproved status
		k.AddToMonitorQueue(ctx, uint64(ctx.BlockHeight())+params.CheckOperationDelta, operation.Index)
	}

	if yesResult.ToDec().Quo(totalVotingPower.ToDec()).GT(threshold) {
		return k.rarimo.ApproveOperation(ctx, operation)
	}

	return k.rarimo.UnapproveOperation(ctx, operation)
}
