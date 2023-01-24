package keeper

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	tokentypes "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"

	cosmostypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k msgServer) CreateTransferOperation(goCtx context.Context, msg *types.MsgCreateTransferOp) (*types.MsgCreateTransferOpResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.checkCreatorIsValidator(ctx, msg.Creator); err != nil {
		return nil, err
	}

	// Index is HASH(tx, event, chain, blockhash)
	index := hexutil.Encode(crypto.Keccak256([]byte(msg.Tx), []byte(msg.EventId), []byte(msg.From.Chain), big.NewInt(ctx.BlockHeight()).Bytes()))

	if _, ok := k.GetOperation(ctx, index); ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	collection, ok := k.tm.GetCollection(ctx, msg.Item.Collection)
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "collection not found by index [%d]", msg.Item.Collection)
	}

	var currentData, targetData tokentypes.CollectionData
	for _, data := range collection.Data {
		if data.Chain == msg.From.Chain {
			currentData, _ = k.tm.GetCollectionData(ctx, data) // always exists
		}

		if data.Chain == msg.To.Chain {
			targetData, _ = k.tm.GetCollectionData(ctx, data) // always exists
		}
	}

	// Check if not found
	if currentData.Collection == "" || targetData.Collection == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "collection data not found")
	}

	var transferOp = types.Transfer{
		Origin:     index,
		Tx:         msg.Tx,
		EventId:    msg.EventId,
		Receiver:   msg.Receiver,
		Amount:     castAmount(msg.Amount, uint8(currentData.Decimals), uint8(targetData.Decimals)),
		BundleData: msg.BundleData,
		BundleSalt: msg.BundleSalt,
		From:       msg.From,
		To:         msg.To,
		Item:       msg.Item,
		Meta:       msg.Meta,
	}

	details, err := cosmostypes.NewAnyWithValue(&transferOp)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error parsing details %s", err.Error())
	}

	var operation = types.Operation{
		Index:         index,
		OperationType: types.OpType_TRANSFER,
		Details:       details,
		Signed:        false,
		Creator:       msg.Creator,
		Timestamp:     uint64(ctx.BlockHeight()),
	}

	k.SetOperation(
		ctx,
		operation,
	)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeNewOperation,
		sdk.NewAttribute(types.AttributeKeyOperationId, operation.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, types.OpType_TRANSFER.String()),
	))
	return &types.MsgCreateTransferOpResponse{}, nil
}

func castAmount(currentAmount string, currentDecimals uint8, targetDecimals uint8) string {
	if currentDecimals == targetDecimals {
		return currentAmount
	}

	value, _ := new(big.Int).SetString(currentAmount, 10)

	if currentDecimals < targetDecimals {
		for i := uint8(0); i < targetDecimals-currentDecimals; i++ {
			value.Mul(value, new(big.Int).SetInt64(10))
		}

		return value.String()
	}

	for i := uint8(0); i < currentDecimals-targetDecimals; i++ {
		value.Div(value, new(big.Int).SetInt64(10))
	}

	return value.String()
}
