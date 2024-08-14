package keeper

import (
	"bytes"
	"fmt"
	"github.com/rarimo/rarimo-core/ethermint/utils"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	evmtypes "github.com/rarimo/rarimo-core/x/evm/types"
	"github.com/rarimo/rarimo-core/x/identity/pkg/state"
	"github.com/rarimo/rarimo-core/x/identity/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey
		memKey   storetypes.StoreKey
		rarimo   types.RarimocoreKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	rarimo types.RarimocoreKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		rarimo:   rarimo,
	}
}

var _ evmtypes.EvmHooks = Keeper{}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) UpdateIdentity(ctx sdk.Context, gist string, id string, hash string) {
	treap := Treap{k}
	lcg := LCG{k}

	// Getting the state by index
	state, ok := k.GetStateInfo(ctx, id)
	if ok {
		currentKey := hexutil.Encode(state.CalculateHash())
		treap.Remove(ctx, currentKey)
	}

	state = types.StateInfo{
		Index:              id,
		Hash:               hash,
		CreatedAtTimestamp: uint64(ctx.BlockTime().UTC().Unix()),
		CreatedAtBlock:     uint64(ctx.BlockHeight()),
	}

	// Saving updated state, updating GIST and changed states list.
	k.SetStateInfo(ctx, state)
	k.SetGIST(ctx, gist)
	k.AddToWaitingList(ctx, id)

	// Inserting into Merkle tree. Insert method already manages how to act if state already exists.

	key := hexutil.Encode(state.CalculateHash())
	priority := lcg.Next(ctx)

	treap.Insert(ctx, key, priority)
	return
}

// Path returns the Merkle path for the requested state hash by state index.
func (k Keeper) Path(ctx sdk.Context, id string) []string {
	treap := Treap{k}
	state, ok := k.GetStateInfo(ctx, id)
	if !ok {
		return nil
	}

	return treap.MerklePath(ctx, hexutil.Encode(state.CalculateHash()))
}

// PostTxProcessing is used to listen EVM smart contract events,
// filter and process `StateTransited` events emitted by configured in module params contract address.
// Will be called by EVM module as hook.
func (k Keeper) PostTxProcessing(ctx sdk.Context, msg core.Message, receipt *ethtypes.Receipt) error {
	params := k.GetParams(ctx)

	const eventName = "StateTransited"

	stateV2, err := abi.JSON(strings.NewReader(state.StateABI))
	if err != nil {
		return err
	}

	// Validating message receiver address (should be our state smart contract)
	if msg.To() == nil || bytes.Compare(msg.To().Bytes(), hexutil.MustDecode(params.IdentityContractAddress)) != 0 {
		return nil
	}

	// https://docs.evmos.org/protocol/modules/evm#posttxprocessing
	for _, log := range receipt.Logs {
		eventId := log.Topics[0]

		event, err := stateV2.EventByID(eventId)
		if err != nil {
			continue
		}

		if event.Name != eventName {
			continue
		}

		eventBody := state.StateStateTransited{}
		if err := utils.UnpackLog(stateV2, &eventBody, event.Name, log); err != nil {
			return err
		}

		k.Logger(ctx).Info(fmt.Sprintf("Received PostTxProcessing event in %s module: %v", types.ModuleName, eventBody))

		// Updating identity Merkle tree
		k.UpdateIdentity(ctx, hexutil.Encode(eventBody.GistRoot.Bytes()), hexutil.Encode(eventBody.Id.Bytes()), hexutil.Encode(eventBody.State.Bytes()))
	}

	return nil
}
