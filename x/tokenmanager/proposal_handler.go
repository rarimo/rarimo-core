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
		case *types.CreateTokenItemProposal:
			return k.HandleCreateTokenItemProposal(ctx, c)
		case *types.RemoveTokenItemProposal:
			return k.RemoveTokenItemProposal(ctx, c)
		case *types.SetNetworkProposal:
			return k.SetNetworkProposal(ctx, c)
		case *types.CreateCollectionProposal:
			return k.HandleCreateCollectionProposal(ctx, c)
		case *types.PutCollectionNetworkAddressProposal:
			return k.HandlePutCollectionNetworkAddressProposal(ctx, c)
		case *types.RemoveCollectionNetworkAddressProposal:
			return k.HandleRemoveCollectionNetworkAddressProposal(ctx, c)
		case *types.RemoveCollectionProposal:
			return k.HandleRemoveCollectionProposal(ctx, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized proposal content type: %T", c)
		}
	}
}
