package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
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
