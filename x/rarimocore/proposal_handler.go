package rarimocore

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/keeper"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func NewProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.AddSignerPartyProposal:
			return k.AddSignerPartyProposal(ctx, c)
		case *types.RemoveSignerPartyProposal:
			return k.RemoveSignerPartyProposal(ctx, c)
		case *types.ReshareKeysProposal:
			return k.ReshareKeysProposal(ctx, c)
		case *types.ChangeThresholdProposal:
			return k.ChangeThresholdProposal(ctx, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized proposal content type: %T", c)
		}
	}
}
