package tokenmanager

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/keeper"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func NewProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.SetTokenInfoProposal:
			return k.SetTokenInfoProposal(ctx, c)
		case *types.SetTokenItemProposal:
			return k.SetTokenItemProposal(ctx, c)
		case *types.RemoveTokenItemProposal:
			return k.RemoveTokenItemProposal(ctx, c)
		case *types.RemoveTokenInfoProposal:
			return k.RemoveTokenInfoProposal(ctx, c)
		case *types.SetNetworkProposal:
			return k.SetNetworkProposal(ctx, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized proposal content type: %T", c)
		}
	}
}
