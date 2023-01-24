package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k msgServer) Vote(goCtx context.Context, msg *types.MsgVote) (*types.MsgVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.checkCreatorIsValidator(ctx, msg.Creator); err != nil {
		return nil, err
	}

	operation, ok := k.GetOperation(ctx, msg.Operation)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "operation not found")
	}

	k.SetVote(ctx, types.Vote{
		Index: &types.VoteIndex{
			Operation: msg.Operation,
			Validator: msg.Creator,
		},
		Vote: msg.Vote,
	})

	if operation.Approved {
		return &types.MsgVoteResponse{}, nil
	}

	type ValidatorVoteInfo struct {
		BondedTokens sdk.Int // Power of a Validator
		Vote         types.VoteType
	}

	validators := make(map[string]*ValidatorVoteInfo)

	// Filling validators map with current validators
	k.staking.IterateBondedValidatorsByPower(ctx, func(index int64, validator stakingtypes.ValidatorI) (stop bool) {
		addr := sdk.ValAddress(validator.GetOperator().Bytes())
		validators[addr.String()] = &ValidatorVoteInfo{
			BondedTokens: validator.GetBondedTokens(),
		}

		return false
	})

	// Setting votes in validators map
	k.IterateVotes(ctx, msg.Operation, func(vote types.Vote) (stop bool) {
		voter := sdk.MustAccAddressFromBech32(vote.Index.Validator)
		if val, ok := validators[voter.String()]; ok {
			val.Vote = vote.Vote
		}
		return false
	})

	totalVotingPower := sdk.ZeroDec()
	yesResult := sdk.ZeroDec()

	for _, validator := range validators {
		votingPower := validator.BondedTokens.ToDec()
		if validator.Vote == types.VoteType_YES {
			yesResult = yesResult.Add(votingPower)
		}
		totalVotingPower = totalVotingPower.Add(votingPower)
	}

	tallyParams := k.gov.GetTallyParams(ctx)
	if yesResult.Quo(totalVotingPower).GT(tallyParams.Threshold) {
		operation.Approved = true
		k.SetOperation(ctx, operation)

		ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeOperationApproved,
			sdk.NewAttribute(types.AttributeKeyOperationId, operation.Index),
			sdk.NewAttribute(types.AttributeKeyOperationType, operation.OperationType.String()),
		))
	}

	return &types.MsgVoteResponse{}, nil
}
