package cscalist

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/rarimo/rarimo-core/x/cscalist/keeper"
	"github.com/rarimo/rarimo-core/x/cscalist/types"
)

func NewProposalHandler(k keeper.Keeper) govv1beta1.Handler {
	return func(ctx sdk.Context, content govv1beta1.Content) error {
		if err := content.ValidateBasic(); err != nil {
			return err
		}

		switch c := content.(type) {
		case *types.EditCSCAListProposal:
			return k.EditCSCAListProposal(ctx, c)
		case *types.ReplaceCSCAListProposal:
			return k.ReplaceCSCAListProposal(ctx, c)
		default:
			return errors.Wrapf(types.ErrUnknownRequest, "unrecognized proposal content type: %T", content)
		}
	}
}
