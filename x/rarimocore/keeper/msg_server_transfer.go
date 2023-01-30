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
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/origin"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k msgServer) CreateTransferOperation(goCtx context.Context, msg *types.MsgCreateTransferOp) (*types.MsgCreateTransferOpResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	defer k.disableFee(ctx.GasMeter().GasConsumed(), ctx.GasMeter())

	if err := k.checkCreatorIsValidator(ctx, msg.Creator); err != nil {
		return nil, err
	}

	// Index is HASH(tx, event, chain)
	index := hexutil.Encode(crypto.Keccak256([]byte(msg.Tx), []byte(msg.EventId), []byte(msg.From.Chain)))

	transferOp, err := k.GetTransfer(ctx, msg)
	if err != nil {
		return nil, err
	}

	details, err := cosmostypes.NewAnyWithValue(transferOp)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error parsing details %s", err.Error())
	}

	var operation = types.Operation{
		Index:         index,
		OperationType: types.OpType_TRANSFER,
		Details:       details,
		Status:        types.OpStatus_INITIALIZED,
		Creator:       msg.Creator,
		Timestamp:     uint64(ctx.BlockHeight()),
	}

	if op, ok := k.GetOperation(ctx, index); ok {
		if op.Status == types.OpStatus_INITIALIZED || op.Status == types.OpStatus_APPROVED {
			// To change operation it should be unapproved or signed
			return &types.MsgCreateTransferOpResponse{}, nil
		}

		// Otherwise - clear votes
		k.IterateVotes(ctx, op.Index, func(vote types.Vote) (stop bool) {
			k.RemoveVote(ctx, vote.Index)
			return false
		})
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

func (k Keeper) GetTransfer(ctx sdk.Context, msg *types.MsgCreateTransferOp) (*types.Transfer, error) {
	origin := origin.NewDefaultOriginBuilder().
		SetTxHash(msg.Tx).
		SetOpId(msg.EventId).
		SetCurrentNetwork(msg.From.Chain).
		Build().
		GetOrigin()

	currentData, ok := k.tm.GetCollectionData(ctx, &tokentypes.CollectionDataIndex{Chain: msg.From.Chain, Address: msg.From.Address})
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "collection data not found")
	}

	targetData, ok := k.tm.GetCollectionData(ctx, &tokentypes.CollectionDataIndex{Chain: msg.To.Chain, Address: msg.To.Address})
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "collection data not found")
	}

	if _, ok = k.tm.GetOnChainItem(ctx, msg.From); !ok && msg.Meta == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "metadata should be provided")
	}

	return &types.Transfer{
		Origin:     hexutil.Encode(origin[:]),
		Tx:         msg.Tx,
		EventId:    msg.EventId,
		Sender:     msg.Sender,
		Receiver:   msg.Receiver,
		Amount:     castAmount(msg.Amount, uint8(currentData.Decimals), uint8(targetData.Decimals)),
		BundleData: msg.BundleData,
		BundleSalt: msg.BundleSalt,
		From:       msg.From,
		To:         msg.To,
		Meta:       msg.Meta,
	}, nil
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
