package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k Keeper) UnfreezeSignerPartyProposal(ctx sdk.Context, proposal *types.UnfreezeSignerPartyProposal) error {
	params := k.GetParams(ctx)

	frozenSignerParty := getPartyByAccount(proposal.Account, params.Parties)
	if frozenSignerParty == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid party account: not found")
	}

	if frozenSignerParty.Status != types.PartyStatus_Frozen {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid party status: %s, party not frozen", frozenSignerParty.Status)
	}

	// Unfreezing resets violations and sets up status as Inactive. Keys resharing is required.
	frozenSignerParty.Status = types.PartyStatus_Inactive
	frozenSignerParty.ViolationsCount = 0
	frozenSignerParty.FreezeEndBlock = 0
	frozenSignerParty.ReportedSessions = []string{}
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

func (k Keeper) SlashProposal(ctx sdk.Context, proposal *types.SlashProposal) error {
	params := k.GetParams(ctx)

	for _, party := range params.Parties {
		if party.Account == proposal.Party {
			party.Status = types.PartyStatus_Slashed
		}
	}

	params.IsUpdateRequired = true
	k.SetParams(ctx, params)
	return nil
}

func (k Keeper) DropPartiesProposal(ctx sdk.Context, _ *types.DropPartiesProposal) error {
	params := k.GetParams(ctx)

	// Unstaking for all unslashed parties
	for _, party := range params.Parties {
		if party.Status != types.PartyStatus_Slashed {
			receiver := party.Account
			if party.Delegator != "" {
				receiver = party.Delegator
			}

			if err := k.unstake(ctx, receiver); err != nil {
				return err
			}
		}
	}

	params.IsUpdateRequired = true
	params.KeyECDSA = ""
	params.Parties = make([]*types.Party, 0, 0)
	params.Threshold = 0

	k.SetParams(ctx, params)
	return nil
}
