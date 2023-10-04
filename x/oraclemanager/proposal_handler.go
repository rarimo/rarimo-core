package oraclemanager

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/rarimo/rarimo-core/x/oraclemanager/keeper"
	"github.com/rarimo/rarimo-core/x/oraclemanager/types"
)

func NewProposalHandler(k keeper.Keeper) govv1beta1.Handler {
	return func(ctx sdk.Context, content govv1beta1.Content) error {
		switch c := content.(type) {
		case *types.OracleUnfreezeProposal:
			return k.OracleUnfreezeProposal(ctx, c)
		case *types.ChangeParamsProposal:
			return k.ChangeParamsProposal(ctx, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized proposal content type: %T", c)
		}
	}
}
