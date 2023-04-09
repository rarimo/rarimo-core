package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k Keeper) UnfreezeSignerPartyProposal(ctx sdk.Context, proposal *types.UnfreezeSignerPartyProposal) error {
	params := k.GetParams(ctx)

	var frozenSignerParty *types.Party

	for _, party := range params.Parties {
		if party.Account != proposal.Account {
			continue
		}

		frozenSignerParty = party
	}

	if frozenSignerParty == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid party account: not found")
	}

	if frozenSignerParty.Status != types.PartyStatus_Frozen {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid party status: %s, party not frozen", frozenSignerParty.Status)
	}

	frozenSignerParty.Status = types.PartyStatus_Inactive
	frozenSignerParty.ViolationsCount = 0
	frozenSignerParty.FreezeEndBlock = 0
	params.IsUpdateRequired = true

	k.SetParams(ctx, params)

	return nil
}

func (k Keeper) ReshareKeysProposal(ctx sdk.Context, _ *types.ReshareKeysProposal) error {
	params := k.GetParams(ctx)
	params.IsUpdateRequired = true
	k.SetParams(ctx, params)
	return nil
}
