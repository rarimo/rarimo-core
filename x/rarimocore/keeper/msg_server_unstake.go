package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k msgServer) Unstake(goCtx context.Context, msg *types.MsgUnstake) (*types.MsgUnstakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)

	if err := k.checkIsAnActiveParty(ctx, msg.Account); err != nil {
		return nil, err
	}

	party := getPartyByAccount(msg.Account, params.Parties)

	if party == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid party account: not found")
	}

	if msg.Creator != party.Account || msg.Creator != party.Delegator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "invalid party account: not authorized")
	}

	if party.Status != types.PartyStatus_Active {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid party account: frozen or slashed party cannot unstake")
	}

	receiver := party.Account
	if party.Delegator != "" {
		receiver = party.Delegator
	}

	if err := k.unstake(ctx, receiver); err != nil {
		return nil, err
	}

	parties := make([]*types.Party, 0)

	for _, p := range params.Parties {
		if p.Account != msg.Account {
			parties = append(parties, p)
		}
	}

	params.Parties = parties
	params.IsUpdateRequired = true
	k.SetParams(ctx, params)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUnstaked,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyPartyAccount, msg.Account),
		),
	})

	return &types.MsgUnstakeResponse{}, nil
}

func (k Keeper) unstake(ctx sdk.Context, receiver string) error {
	receiverAddr, _ := sdk.AccAddressFromBech32(receiver)

	stakeCoin, err := k.getStakeCoin(ctx)
	if err != nil {
		return sdkerrors.Wrap(err, "failed to get stake coin")
	}

	if err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.NewCoins(stakeCoin)); err != nil {
		return sdkerrors.Wrapf(err, "failed to send coins from module to account, account: %s", receiver)
	}

	return nil
}
