package rarimocore

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/rarimo/rarimo-core/x/rarimocore/keeper"
	"github.com/rarimo/rarimo-core/x/rarimocore/types"
)

func NewProposalHandler(k keeper.Keeper) govv1beta1.Handler {
	return func(ctx sdk.Context, content govv1beta1.Content) error {
		switch c := content.(type) {
		case *types.UnfreezeSignerPartyProposal:
			return k.UnfreezeSignerPartyProposal(ctx, c)
		case *types.ReshareKeysProposal:
			return k.ReshareKeysProposal(ctx, c)
		case *types.SlashProposal:
			return k.SlashProposal(ctx, c)
		case *types.DropPartiesProposal:
			return k.DropPartiesProposal(ctx, c)
		case *types.ArbitrarySigningProposal:
			return k.ArbitrarySigningProposal(ctx, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized proposal content type: %T", c)
		}
	}
}
