package keeper

import (
	"bytes"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/rarimo/rarimo-core/ethermint/utils"
	"github.com/rarimo/rarimo-core/x/rootupdater/pkg/state"
	"github.com/tendermint/tendermint/libs/log"
	"strings"

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
// filter and process `RootUpdated` events emitted by configured in module params contract address.
// Will be called by EVM module as hook.
func (k Keeper) PostTxProcessing(ctx sdk.Context, msg core.Message, receipt *ethtypes.Receipt) error {
	params := k.GetParams(ctx)

	k.Logger(ctx).Info("PostTxProcessing", "msg", msg, "receipt", receipt)

	stateV2, err := abi.JSON(strings.NewReader(state.PoseidonSMTABI))
	if err != nil {
		k.Logger(ctx).Error("failed to marshal poseidon smart abi", "error", err)
		return nil
	}

	contractAddress, err := hexutil.Decode(params.ContractAddress)
	if err != nil {
		// If return an error here, the whole EVM module won't work
		k.Logger(ctx).Error("failed to decode contract address")
		return nil
	}

	// https://docs.evmos.org/protocol/modules/evm#posttxprocessing

	if len(receipt.Logs) == 0 {
		k.Logger(ctx).Error("logs is empty")
	}

	for _, log := range receipt.Logs {
		// Validating message receiver address (should be our state smart contract)
		if log.Address.Bytes() == nil || bytes.Compare(log.Address.Bytes(), contractAddress) != 0 {
			k.Logger(ctx).Info(fmt.Sprintf("inappropriate contract address: expected %s got %s", params.ContractAddress, log.Address.String()))
			continue
		}

		eventId := log.Topics[0]
		event, err := stateV2.EventByID(eventId)
		if err != nil {
			k.Logger(ctx).Error("failed to get event by ID")
			continue
		}

		if event.Name != params.EventName {
			k.Logger(ctx).Info(fmt.Sprintf("unmatched event: got %s, expected %s", event.Name, params.EventName))
			continue
		}

		eventBody := state.PoseidonSMTRootUpdated{}
		if err := utils.UnpackLog(stateV2, &eventBody, event.Name, log); err != nil {
			k.Logger(ctx).Error("failed to unpack event body")
			continue
		}

		params.Root = hexutil.Encode(eventBody.Root[:])
		params.RootTimestamp = ctx.BlockTime().Unix()
		params.BlockHeight = log.BlockNumber

		k.Logger(ctx).Info(fmt.Sprintf("Received PostTxProcessing event in %s module: %v", types.ModuleName, eventBody))
		k.SetParams(ctx, params)
	}

	return nil
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
