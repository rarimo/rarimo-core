package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/bridge/types"
	rarimocoretypes "gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
	tokenmanagertypes "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k msgServer) DepositNative(goCtx context.Context, msg *types.MsgDepositNative) (*types.MsgDepositNativeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, _ := sdk.AccAddressFromBech32(msg.Creator)

	err := k.bankKeeper.BurnTokens(ctx, creatorAddr, sdk.Coins{*msg.Amount})
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to burn tokens for address (%s)", creatorAddr.String())
	}

	// TODO: Fix setting approved operation
	_, err = k.rarimocoreKeeper.CreateTransferOperation(goCtx, &rarimocoretypes.MsgCreateTransferOp{
		Creator:    msg.Creator,
		Tx:         msg.Seed,
		EventId:    "",
		Sender:     msg.Creator,
		Receiver:   msg.Receiver,
		Amount:     msg.Amount.Amount.String(),
		BundleData: msg.BundleData,
		BundleSalt: msg.BundleSalt,
		From: &tokenmanagertypes.OnChainItemIndex{
			Chain: ctx.ChainID(),
		},
		To: msg.To,
	})
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to create transfer operation for address (%s)", creatorAddr.String())
	}

	return &types.MsgDepositNativeResponse{}, nil
}
