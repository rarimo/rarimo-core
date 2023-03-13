package keeper

import (
	"context"
	cosmostypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/rarimo/rarimo-core/x/bridge/types"
	operationorigin "gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/origin"
	rarimocoretypes "gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
	tokenmanagertypes "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k msgServer) DepositNative(goCtx context.Context, msg *types.MsgDepositNative) (*types.MsgDepositNativeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	err = k.bankKeeper.BurnTokens(ctx, creatorAddr, sdk.Coins{*msg.Amount})
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to burn tokens for address (%s)", creatorAddr.String())
	}

	origin := operationorigin.NewDefaultOriginBuilder().
		SetTxHash(msg.Seed).
		SetOpId("").
		SetCurrentNetwork(ctx.ChainID()).
		Build().
		GetOrigin()

	transfer := &rarimocoretypes.Transfer{
		Origin:     hexutil.Encode(origin[:]),
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
	}

	details, err := cosmostypes.NewAnyWithValue(transfer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error parsing details %s", err.Error())
	}

	k.rarimocoreKeeper.SetOperation(ctx, rarimocoretypes.Operation{
		Index:         hexutil.Encode(crypto.Keccak256([]byte(msg.Seed), []byte(ctx.ChainID()))),
		OperationType: rarimocoretypes.OpType_TRANSFER,
		Details:       details,
		Status:        rarimocoretypes.OpStatus_APPROVED,
		Creator:       msg.Creator,
		Timestamp:     uint64(ctx.BlockHeight()),
	})

	return &types.MsgDepositNativeResponse{}, nil
}
