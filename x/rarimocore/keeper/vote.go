package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

// SetVote set a specific deposit in the store from its index
func (k Keeper) SetVote(ctx sdk.Context, v types.Vote) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKeyPrefix))
	b := k.cdc.MustMarshal(&v)
	store.Set(types.VoteKey(v.Index), b)
}

// GetVote returns an Vote from its index
func (k Keeper) GetVote(
	ctx sdk.Context,
	index *types.VoteIndex,
) (val types.Vote, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKeyPrefix))

	b := store.Get(types.VoteKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVote removes an Vote from the store
func (k Keeper) RemoveVote(
	ctx sdk.Context,
	index *types.VoteIndex,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKeyPrefix))
	store.Delete(types.VoteKey(
		index,
	))
}

// GetAllVote returns all Vote
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

func (k Keeper) IterateVotes(ctx sdk.Context, operation string, f func(vote types.Vote) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(operation))

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var vote types.Vote
		k.cdc.MustUnmarshal(iterator.Value(), &vote)
		if f(vote) {
			break
		}
	}
}
