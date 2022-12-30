package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k Keeper) AddSignerPartyProposal(ctx sdk.Context, proposal *types.AddSignerPartyProposal) error {
	params := k.GetParams(ctx)
	params.Parties = append(params.Parties, &types.Party{
		PubKey:   proposal.TrialPublicKey,
		Address:  proposal.Address,
		Account:  proposal.Account,
		Verified: false,
	})
	params.Threshold = uint64(crypto.GetThreshold(len(params.Parties)))
	params.IsUpdateRequired = true
	k.SetParams(ctx, params)
	return nil
}

func (k Keeper) RemoveSignerPartyProposal(ctx sdk.Context, proposal *types.RemoveSignerPartyProposal) error {
	params := k.GetParams(ctx)
	parties := make([]*types.Party, 0, len(params.Parties)-1)
	for _, p := range params.Parties {
		if p.Account != proposal.Account {
			parties = append(parties, p)
		}
	}

	params.Parties = parties
	params.Threshold = uint64(crypto.GetThreshold(len(params.Parties)))
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

func (k Keeper) ChangeThresholdProposal(ctx sdk.Context, proposal *types.ChangeThresholdProposal) error {
	params := k.GetParams(ctx)
	if proposal.Threshold < uint32(len(params.Parties)) {
		params.Threshold = uint64(proposal.Threshold)
		params.IsUpdateRequired = true
		k.SetParams(ctx, params)
	}
	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid threshold: must be less the N")
}
