package keeper

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/rarimo/rarimo-core/x/rootupdater/pkg/state"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/rarimo/rarimo-core/x/rootupdater/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
		rarimo     types.RarimocoreKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	rarimo types.RarimocoreKeeper,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		rarimo:     rarimo,
	}
}

// PostTxProcessing is used to listen EVM smart contract events,
// filter and process `StateTransited` events emitted by configured in module params contract address.
// Will be called by EVM module as hook.
func (k Keeper) PostTxProcessing(ctx sdk.Context, msg core.Message, receipt *ethtypes.Receipt) error {
	params := k.GetParams(ctx)

	stateV2, err := abi.JSON(strings.NewReader(state.PoseidonSMTABI))
	if err != nil {
		return err
	}

	// Validating message receiver address (should be our state smart contract)
	if msg.To() == nil || bytes.Compare(msg.To().Bytes(), hexutil.MustDecode(params.ContractAddress)) != 0 {
		return nil
	}

	// https://docs.evmos.org/protocol/modules/evm#posttxprocessing
	for _, log := range receipt.Logs {
		eventId := log.Topics[0]

		event, err := stateV2.EventByID(eventId)
		if err != nil {
			continue
		}

		if event.Name != params.EventName {
			continue
		}

		eventBody := state.PoseidonSMTRootUpdated{}
		if err := unpackLog(stateV2, &eventBody, event.Name, log); err != nil {
			return err
		}

		params.Root = hex.EncodeToString(eventBody.Root[:])

		k.Logger(ctx).Info(fmt.Sprintf("Received PostTxProcessing event in %s module: %v", types.ModuleName, eventBody))
		k.SetParams(ctx, params)
	}

	return nil
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// unpackLog copy-pasted from logic in generated s-c bindings.
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
