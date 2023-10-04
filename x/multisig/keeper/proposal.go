package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pkg/errors"
	"github.com/rarimo/rarimo-core/x/multisig/types"
)

// SetProposal set a specific proposal in the store from its id
func (k Keeper) SetProposal(ctx sdk.Context, proposal types.Proposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposalKeyPrefix))
	b := k.cdc.MustMarshal(&proposal)
	store.Set(types.ProposalKey(proposal.Id), b)
}

// GetProposal returns a proposal from its id
func (k Keeper) GetProposal(ctx sdk.Context, id uint64) (val types.Proposal, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposalKeyPrefix))

	b := store.Get(types.ProposalKey(id))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveProposal removes a proposal from the store
func (k Keeper) RemoveProposal(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposalKeyPrefix))
	store.Delete(types.ProposalKey(id))
}

// GetAllProposal returns all proposals
func (k Keeper) GetAllProposal(ctx sdk.Context) (list []types.Proposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Proposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) IterateProposals(ctx sdk.Context, f func(report types.Proposal) (stop bool)) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProposalKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var proposal types.Proposal
		k.cdc.MustUnmarshal(iterator.Value(), &proposal)
		if f(proposal) {
			break
		}
	}
}

func (k Keeper) Tally(ctx sdk.Context, proposalId uint64) *types.TallyResult {
	tally := types.TallyResult{}
	votes := k.GetAllVoteByProposalId(ctx, proposalId)

	for _, vote := range votes {
		switch vote.Option {
		case types.VoteOption_YES:
			tally.YesCount++
		case types.VoteOption_NO:
			tally.NoCount++
		}
	}

	return &tally
}

func (k Keeper) ExecuteProposal(ctx sdk.Context, proposal types.Proposal) (status types.ProposalStatus, err error) {
	// Caching context so that we don't update the store in case of failure.
	cacheCtx, flush := ctx.CacheContext()
	status = types.ProposalStatus_EXECUTED

	if err = k.doExecuteMsgs(cacheCtx, proposal); err != nil {
		status = types.ProposalStatus_FAILED
	}
	if status == types.ProposalStatus_EXECUTED {
		flush()
	}

	return
}

// doExecuteMsgs routes the messages to the registered handlers. Messages are limited to those that require no authZ or
// by the account of group only. Otherwise, this gives access to other peoples accounts as the sdk middlewares are bypassed
func (k Keeper) doExecuteMsgs(ctx sdk.Context, proposal types.Proposal) error {
	msgs, err := proposal.GetMsgs()
	if err != nil {
		return err
	}

	if err = ensureMsgAuthZ(msgs, proposal.Group); err != nil {
		return err
	}

	var events sdk.Events

	for i, msg := range msgs {
		handler := k.router.Handler(msg)
		if handler == nil {
			return sdkerrors.Wrapf(sdkerrors.ErrLogic, "no message handler found for %q", sdk.MsgTypeURL(msg))
		}
		r, err := handler(ctx, msg)
		if err != nil {
			return errors.Wrapf(err, "message %s at position %d", sdk.MsgTypeURL(msg), i)
		}
		// Handler should always return non-nil sdk.Result.
		if r == nil {
			return fmt.Errorf("got nil sdk.Result for message %q at position %d", msg, i)
		}
		events = append(events, r.GetEvents()...)
	}

	ctx.EventManager().EmitEvents(events)
	return nil
}
