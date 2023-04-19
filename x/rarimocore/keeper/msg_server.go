package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
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

func (k msgServer) checkCreatorIsValidator(ctx sdk.Context, creator string) error {
	sender, _ := sdk.AccAddressFromBech32(creator)

	if k.staking.Validator(ctx, sdk.ValAddress(sender.Bytes())) != nil {
		return nil
	}

	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "sender is not a validator")
}

func (k *Keeper) checkIsAnActiveParty(ctx sdk.Context, sender string) error {
	for _, party := range k.GetParams(ctx).Parties {
		if party.Account != sender {
			continue
		}

		if party.Status == types.PartyStatus_Active {
			return nil
		}
	}

	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "not an active party")
}

func (k *Keeper) checkIsValidParty(ctx sdk.Context, sender string) error {
	for _, party := range k.GetParams(ctx).Parties {
		if party.Account != sender {
			continue
		}

		if party.Status == types.PartyStatus_Active || party.Status == types.PartyStatus_Inactive {
			return nil
		}
	}

	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "not an active party")
}
