package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/oraclemanager/types"
)

func (k msgServer) Stake(goCtx context.Context, msg *types.MsgStake) (*types.MsgStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// REQUIRES: validated index
	oracle, ok := k.GetOracle(ctx, msg.Index)
	if ok && oracle.Status != types.OracleStatus_Inactive && oracle.Status != types.OracleStatus_Active {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "can not stake for oracle: invalid status in existing entry")
	}

	// For recalculating of stake summary. Will be default values if oracle does not exist yet.
	existingDelegations := make([]types.Delegation, 0, 1)
	existingStake := sdk.ZeroInt()

	if ok {
		existingStake, _ = sdk.NewIntFromString(oracle.Stake)
		existingDelegations = oracle.Delegations
	}

	params := k.GetParams(ctx)

	// REQUIRES: validated amount
	amount, _ := sdk.NewIntFromString(msg.Amount)

	moduleAccount := k.ak.GetModuleAccount(ctx, types.ModuleName)
	if moduleAccount == nil {
		panic(sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "module account %s does not exist", types.ModuleName))
	}

	if msg.From == "" {
		msg.From = msg.Index.Account
	}

	// REQUIRES: validated account
	payerAddr, _ := sdk.AccAddressFromBech32(msg.From)

	// Transferring tokens to module account
	if err := k.bank.SendCoins(ctx, payerAddr, moduleAccount.GetAddress(), sdk.NewCoins(sdk.NewCoin(params.StakeDenom, amount))); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error withdrawing coins: %s", err.Error())
	}

	// Calculation of total stake and check it stake is enough to be active.

	totalStake := existingStake.Add(amount)
	resultStatus := types.OracleStatus_Active

	// REQUIRES: validated amount in params
	requiredAmount, _ := sdk.NewIntFromString(params.MinOracleStake)
	if totalStake.LT(requiredAmount) {
		resultStatus = types.OracleStatus_Inactive
	}

	// Iterating over existing delegators to check if wee need to increment the existing or add new ones to the list.
	delegatorExist := false
	for i, d := range existingDelegations {
		if d.Delegator == msg.From {
			delegatorExist = true
			existingDelegation, _ := sdk.NewIntFromString(d.Amount)
			existingDelegations[i].Amount = existingDelegation.Add(amount).String()
			break
		}
	}

	if !delegatorExist {
		existingDelegations = append(existingDelegations, types.Delegation{
			Delegator: msg.From,
			Amount:    amount.String(),
		})
	}

	k.SetOracle(ctx, types.Oracle{
		Index:       msg.Index,
		Status:      resultStatus,
		Stake:       totalStake.String(),
		Delegations: existingDelegations,
	})

	if resultStatus == types.OracleStatus_Active {
		ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeOracleActivated,
			sdk.NewAttribute(types.AttributeKeyChain, msg.Index.Chain),
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Index.Account),
		))
	}

	return &types.MsgStakeResponse{}, nil
}
