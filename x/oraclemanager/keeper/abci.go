package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/oraclemanager/types"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/pkg"
	rarimotypes "github.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k Keeper) EndBlocker(ctx sdk.Context) {
	k.IterateOverMonitorQueue(ctx, uint64(ctx.BlockHeight()), func(operation rarimotypes.Operation) (stop bool) {
		k.RemoveFromMonitorQueue(ctx, uint64(ctx.BlockHeight()), operation.Index)

		monitoringOperationTypes := map[rarimotypes.OpType]struct{}{
			rarimotypes.OpType_TRANSFER:                    {},
			rarimotypes.OpType_IDENTITY_GIST_TRANSFER:      {},
			rarimotypes.OpType_IDENTITY_DEFAULT_TRANSFER:   {},
			rarimotypes.OpType_WORLDCOIN_IDENTITY_TRANSFER: {},
		}

		if _, ok := monitoringOperationTypes[operation.OperationType]; !ok {
			return false
		}

		chain, err := getSourceChain(operation)
		if err != nil {
			return false
		}

		results := make(map[string]*rarimotypes.VoteType)
		for _, oracle := range k.GetOracleForChain(ctx, chain) {
			if oracle.Status == types.OracleStatus_Active {
				results[oracle.Index.Account] = nil
			}
		}

		k.rarimo.IterateVotes(ctx, operation.Index, func(vote rarimotypes.Vote) (stop bool) {
			if _, ok := results[vote.Index.Validator]; !ok {
				return false
			}

			results[vote.Index.Validator] = &vote.Vote
			return false
		})

		switch operation.Status {
		case rarimotypes.OpStatus_APPROVED:
			for account, vote := range results {
				if vote == nil {
					k.NoteMissed(ctx, &types.OracleIndex{Chain: chain, Account: account})
					continue
				}

				if *vote != rarimotypes.VoteType_YES {
					k.NoteViolation(ctx, &types.OracleIndex{Chain: chain, Account: account})
				}
			}
		case rarimotypes.OpStatus_NOT_APPROVED:
			for account, vote := range results {
				if vote == nil {
					k.NoteMissed(ctx, &types.OracleIndex{Chain: chain, Account: account})
					continue
				}

				if *vote != rarimotypes.VoteType_NO {
					k.NoteViolation(ctx, &types.OracleIndex{Chain: chain, Account: account})
				}
			}

			k.NoteViolation(ctx, &types.OracleIndex{Chain: chain, Account: operation.Creator})
		}

		return false
	})

	k.IterateOverFreezedQueue(ctx, uint64(ctx.BlockHeight()), func(oracle types.Oracle) (stop bool) {
		k.RemoveFromFreezedQueue(ctx, uint64(ctx.BlockHeight()), oracle.Index)
		if oracle.Status != types.OracleStatus_Freezed {
			return false
		}

		oracle.Status = types.OracleStatus_Slashed
		k.SetOracle(ctx, oracle)

		ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeOracleSlashed,
			sdk.NewAttribute(types.AttributeKeyChain, oracle.Index.Chain),
			sdk.NewAttribute(types.AttributeKeyAccount, oracle.Index.Account),
		))

		return false
	})
}

func (k Keeper) NoteViolation(ctx sdk.Context, index *types.OracleIndex) {
	params := k.GetParams(ctx)
	if oracle, ok := k.GetOracle(ctx, index); ok {

		oracle.ViolationsCount = oracle.ViolationsCount + 1
		if oracle.ViolationsCount == params.MaxViolationsCount {
			oracle.Status = types.OracleStatus_Freezed
			oracle.FreezeEndBlock = uint64(ctx.BlockHeight()) + params.SlashedFreezeBlocks

			k.AddToFreezedQueue(ctx, uint64(ctx.BlockHeight()), index)

			ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeOracleFreezed,
				sdk.NewAttribute(types.AttributeKeyChain, index.Chain),
				sdk.NewAttribute(types.AttributeKeyAccount, index.Account),
			))
		}
		k.SetOracle(ctx, oracle)
	}
}

func (k Keeper) NoteMissed(ctx sdk.Context, index *types.OracleIndex) {
	params := k.GetParams(ctx)
	if oracle, ok := k.GetOracle(ctx, index); ok {
		oracle.MissedCount = oracle.MissedCount + 1
		if oracle.MissedCount == params.MaxMissedCount {
			oracle.Status = types.OracleStatus_Jailed

			ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeOracleJailed,
				sdk.NewAttribute(types.AttributeKeyChain, index.Chain),
				sdk.NewAttribute(types.AttributeKeyAccount, index.Account),
			))
		}

		k.SetOracle(ctx, oracle)
	}
}

func getSourceChain(op rarimotypes.Operation) (string, error) {
	switch op.OperationType {
	case rarimotypes.OpType_TRANSFER:
		transfer, err := pkg.GetTransfer(op)
		if err != nil {
			return "", err
		}

		return transfer.From.Chain, nil

	case rarimotypes.OpType_IDENTITY_DEFAULT_TRANSFER:
		transfer, err := pkg.GetIdentityDefaultTransfer(op)
		if err != nil {
			return "", err
		}

		return transfer.Chain, nil

	case rarimotypes.OpType_IDENTITY_GIST_TRANSFER:
		transfer, err := pkg.GetIdentityGISTTransfer(op)
		if err != nil {
			return "", err
		}

		return transfer.Chain, nil

	case rarimotypes.OpType_WORLDCOIN_IDENTITY_TRANSFER:
		transfer, err := pkg.GetWorldCoinIdentityTransfer(op)
		if err != nil {
			return "", err
		}
		return transfer.Chain, nil
	}

	return "", nil
}
