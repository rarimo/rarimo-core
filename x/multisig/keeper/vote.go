package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/multisig/types"
)

// SetVote set a specific vote in the store from its proposal id and voter
func (k Keeper) SetVote(ctx sdk.Context, vote types.Vote) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKeyPrefix))
	b := k.cdc.MustMarshal(&vote)
	store.Set(types.VoteKey(vote.ProposalId, vote.Voter), b)
}

// GetVote returns a vote from its proposal id and voter
func (k Keeper) GetVote(ctx sdk.Context, proposalId uint64, voter string) (val types.Vote, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKeyPrefix))

	b := store.Get(types.VoteKey(proposalId, voter))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVote removes a vote from the store
func (k Keeper) RemoveVote(ctx sdk.Context, proposalId uint64, voter string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKeyPrefix))
	store.Delete(types.VoteKey(proposalId, voter))
}

// GetAllVote returns all votes
func (k Keeper) GetAllVote(ctx sdk.Context) (list []types.Vote) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Vote
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllVoteByProposalId returns all votes by proposal id
func (k Keeper) GetAllVoteByProposalId(ctx sdk.Context, proposalId uint64) (list []types.Vote) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, types.ProposalKey(proposalId))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Vote
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
