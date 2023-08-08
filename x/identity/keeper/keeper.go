package keeper

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/tendermint/tendermint/libs/log"
	evmtypes "gitlab.com/rarimo/rarimo-core/x/evm/types"
	"gitlab.com/rarimo/rarimo-core/x/identity/pkg/state"
	"gitlab.com/rarimo/rarimo-core/x/identity/types"
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

	state, ok := k.GetStateInfo(ctx, id)
	if ok {
		currentKey := hexutil.Encode(state.CalculateHash())
		treap.Remove(ctx, currentKey)
	}

	state = types.StateInfo{
		Index:              id,
		Hash:               hash,
		CreatedAtTimestamp: uint64(ctx.BlockTime().UTC().Unix()),
	}

	k.SetStateInfo(ctx, state)
	k.SetGIST(ctx, gist)
	k.AddToWaitingList(ctx, id)

	key := hexutil.Encode(state.CalculateHash())
	priority := lcg.Next(ctx)

	treap.Insert(ctx, key, priority)
	return
}

func (k Keeper) Path(ctx sdk.Context, id string) []string {
	treap := Treap{k}
	state, ok := k.GetStateInfo(ctx, id)
	if !ok {
		return nil
	}

	return treap.MerklePath(ctx, hexutil.Encode(state.CalculateHash()))
}

func (k Keeper) PostTxProcessing(ctx sdk.Context, msg core.Message, receipt *ethtypes.Receipt) error {
	params := k.GetParams(ctx)

	const eventName = "StateTransited"

	stateV2, err := abi.JSON(strings.NewReader(state.StateABI))
	if err != nil {
		return err
	}

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
		if err := unpackLog(stateV2, &eventBody, event.Name, log); err != nil {
			return err
		}

		k.Logger(ctx).Info(fmt.Sprintf("Received PostTxProcessing event in %s module: %v", types.ModuleName, eventBody))
		k.UpdateIdentity(ctx, hexutil.Encode(eventBody.GistRoot.Bytes()), hexutil.Encode(eventBody.Id.Bytes()), hexutil.Encode(eventBody.State.Bytes()))
	}

	return nil
}

func unpackLog(contractAbi abi.ABI, out interface{}, event string, log *ethtypes.Log) error {
	if log.Topics[0] != contractAbi.Events[event].ID {
		return fmt.Errorf("event signature mismatch")
	}

	if len(log.Data) > 0 {
		if err := contractAbi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range contractAbi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return abi.ParseTopics(out, indexed, log.Topics[1:])
}
