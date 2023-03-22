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

	var party *types.Party

	for _, p := range params.Parties {
		if p.Account == msg.Account {
			party = p
		}
	}

	if party == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid party account: not found")
	}

	if party.Status == types.PartyStatus_Frozen || party.Status == types.PartyStatus_Slashed {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid party account: frozen or slashed party cannot unstake")
	}

	receiver := party.Account

	if party.Delegator != "" {
		receiver = party.Delegator
	}

	stakeCoin, err := k.getStakeCoin(ctx)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to get stake coin")
	}

	receiverAddr, _ := sdk.AccAddressFromBech32(receiver)

	if err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiverAddr, sdk.NewCoins(stakeCoin)); err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to send coins from module to account, account: %s", receiver)
	}

	parties := make([]*types.Party, 0)

	for _, p := range params.Parties {
		if p.Account != msg.Account {
			parties = append(parties, p)
		}
	}

	params.Parties = parties
	k.UpdateParams(ctx, params)

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
