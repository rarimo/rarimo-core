package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/tendermint/tendermint/libs/log"
	evmtypes "gitlab.com/rarimo/rarimo-core/x/evm/types"
	"gitlab.com/rarimo/rarimo-core/x/identity/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey
		memKey   storetypes.StoreKey
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
	}
}

var _ evmtypes.EvmHooks = Keeper{}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) UpdateIdentity(ctx sdk.Context, id string, hash string) {
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
		CreatedAtBlock:     uint64(ctx.BlockHeight()),
	}

	k.SetStateInfo(ctx, state)

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
	//params := k.GetParams(ctx)

	// TODO
	// https://docs.evmos.org/protocol/modules/evm#posttxprocessing

	return nil
}
