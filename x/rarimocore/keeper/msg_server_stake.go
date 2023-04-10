package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k msgServer) Stake(goCtx context.Context, msg *types.MsgStake) (*types.MsgStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)

	for _, party := range params.Parties {
		if party.Account == msg.Account {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "account already set")
		}
	}

	if err := k.stake(ctx, msg.Creator); err != nil {
		return nil, err
	}

	party := &types.Party{
		PubKey:  msg.TrialPublicKey,
		Address: msg.Address,
		Account: msg.Account,
		Status:  types.PartyStatus_Inactive,
	}

	if msg.Creator != msg.Account {
		party.Delegator = msg.Creator
	}

	params.Parties = append(params.Parties, party)
	params.IsUpdateRequired = true
	k.SetParams(ctx, params)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeStaked,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyPartyAccount, msg.Account),
		),
	})

	return &types.MsgStakeResponse{}, nil
}

func (k Keeper) stake(ctx sdk.Context, sender string) error {
	params := k.GetParams(ctx)

	senderAddr, _ := sdk.AccAddressFromBech32(sender)

	stakeCoin, err := k.getStakeCoin(ctx)
	if err != nil {
		return sdkerrors.Wrap(err, "failed to get stake coin")
	}

	balance := k.bank.GetBalance(ctx, senderAddr, params.StakeDenom)

	if balance.IsLT(stakeCoin) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "insufficient funds")
	}

	if err = k.bank.SendCoinsFromAccountToModule(ctx, senderAddr, types.ModuleName, sdk.NewCoins(stakeCoin)); err != nil {
		return sdkerrors.Wrapf(err, "failed to send coins from account to module, account: %s", sender)
	}

	return nil
}

func (k Keeper) getStakeCoin(ctx sdk.Context) (sdk.Coin, error) {
	params := k.GetParams(ctx)

	amount, ok := sdk.NewIntFromString(params.StakeAmount)
	if !ok {
		return sdk.Coin{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid stake amount")
	}

	return sdk.NewCoin(params.StakeDenom, amount), nil
}
