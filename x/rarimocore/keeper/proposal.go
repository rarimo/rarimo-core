package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k Keeper) AddSignerParty(ctx sdk.Context, proposal *types.AddSignerPartyProposal) error {
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

func (k Keeper) RemoveSignerParty(ctx sdk.Context, proposal *types.RemoveSignerPartyProposal) error {
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

func (k Keeper) ReshareKeys(ctx sdk.Context, _ *types.ReshareKeysProposal) error {
	params := k.GetParams(ctx)
	params.IsUpdateRequired = true
	k.SetParams(ctx, params)
	return nil
}
