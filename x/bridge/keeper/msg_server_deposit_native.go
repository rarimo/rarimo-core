package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gitlab.com/rarimo/rarimo-core/x/bridge/types"
	operationorigin "gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/origin"
	rarimocoretypes "gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
	tokenmanagertypes "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k msgServer) DepositNative(goCtx context.Context, msg *types.MsgDepositNative) (*types.MsgDepositNativeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, _ := sdk.AccAddressFromBech32(msg.Creator)

	if err := k.chargeCommission(ctx, creatorAddr); err != nil {
		return nil, sdkerrors.Wrap(err, "failed to charge commission")
	}

	if err := k.bankKeeper.BurnTokens(ctx, creatorAddr, sdk.Coins{*msg.Amount}); err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to burn tokens for address (%s)", creatorAddr.String())
	}

	origin := operationorigin.NewDefaultOriginBuilder().
		SetTxHash(msg.Seed).
		SetOpId("").
		SetCurrentNetwork(ctx.ChainID()).
		Build().
		GetOrigin()

	err := k.rarimocoreKeeper.CreateTransferOperation(ctx, msg.Creator, &rarimocoretypes.Transfer{
		Origin:     hexutil.Encode(origin[:]),
		Tx:         msg.Seed,
		EventId:    "",
		Sender:     msg.Creator,
		Receiver:   msg.Receiver,
		Amount:     msg.Amount.Amount.String(),
		BundleData: msg.BundleData,
		BundleSalt: msg.BundleSalt,
		From: tokenmanagertypes.OnChainItemIndex{
			Chain: types.NetworkName,
		},
		To: msg.To,
	}, true)

	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to create transfer operation for address (%s)", creatorAddr.String())
	}

	return &types.MsgDepositNativeResponse{}, nil
}

func (k Keeper) chargeCommission(ctx sdk.Context, address sdk.AccAddress) error {
	network, ok := k.tokenmanagerKeepr.GetNetwork(ctx, types.NetworkName)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "native network not found")
	}

	feeparams := network.GetFeeParams()
	if feeparams == nil || len(feeparams.FeeTokens) != 1 {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "native fee token not found")
	}

	// picking first token, should be native
	feeToken := feeparams.FeeTokens[0]
	fee, _ := sdk.NewIntFromString(feeToken.Amount)
	return k.bankKeeper.SendCoinsFromAccountToModule(ctx, address, types.ModuleName, sdk.NewCoins(sdk.NewCoin(k.GetParams(ctx).WithdrawDenom, fee)))
}
