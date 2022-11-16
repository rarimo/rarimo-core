package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (msgServer) disableFee(initial sdk.Gas, meter sdk.GasMeter) {
	meter.RefundGas(meter.GasConsumed()-initial, "Gas disabled for that message")
}

func (k msgServer) checkCreatorIsValidator(ctx sdk.Context, creator string) bool {
	addr, _ := sdk.ValAddressFromBech32(creator)
	return k.staking.Validator(ctx, addr) != nil
}

func (k *Keeper) checkSenderIsAParty(ctx sdk.Context, sender string) error {
	for _, party := range k.GetParams(ctx).Parties {
		if party.Account == sender {
			return nil
		}
	}

	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "sender is not a party")
}
