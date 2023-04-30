package keeper

import (
	"math/big"

	cosmostypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
	tokentypes "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

const SystemCreator = "system"

func (k Keeper) CreateAddFeeTokenOperation(ctx sdk.Context, token tokentypes.FeeToken, chain string) error {
	network, ok := k.tm.GetNetwork(ctx, chain)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "network not found")
	}

	op := &types.FeeTokenManagement{
		OpType:   types.FeeTokenManagementType_ADD_FEE_TOKEN,
		Token:    token,
		Chain:    chain,
		Receiver: network.Contract,
	}

	return k.CreateFeeTokenManagementOperation(ctx, op)
}

func (k Keeper) CreateRemoveFeeTokenOperation(ctx sdk.Context, token tokentypes.FeeToken, chain string) error {
	network, ok := k.tm.GetNetwork(ctx, chain)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "network not found")
	}

	op := &types.FeeTokenManagement{
		OpType:   types.FeeTokenManagementType_REMOVE_FEE_TOKEN,
		Token:    token,
		Chain:    chain,
		Receiver: network.Contract,
	}

	return k.CreateFeeTokenManagementOperation(ctx, op)
}

func (k Keeper) CreateUpdateFeeTokenOperation(ctx sdk.Context, token tokentypes.FeeToken, chain string) error {
	network, ok := k.tm.GetNetwork(ctx, chain)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "network not found")
	}

	op := &types.FeeTokenManagement{
		OpType:   types.FeeTokenManagementType_UPDATE_FEE_TOKEN,
		Token:    token,
		Chain:    chain,
		Receiver: network.Contract,
	}

	return k.CreateFeeTokenManagementOperation(ctx, op)
}

func (k Keeper) CreateWithdrawFeeOperation(ctx sdk.Context, token tokentypes.FeeToken, chain string, receiver string) error {
	op := &types.FeeTokenManagement{
		OpType:   types.FeeTokenManagementType_WITHDRAW_FEE_TOKEN,
		Token:    token,
		Chain:    chain,
		Receiver: receiver,
	}

	return k.CreateFeeTokenManagementOperation(ctx, op)
}

func (k Keeper) CreateFeeTokenManagementOperation(ctx sdk.Context, op *types.FeeTokenManagement) error {
	// Index is HASH(block height, chain, fee token contract, fee amount)
	index := hexutil.Encode(crypto.Keccak256(big.NewInt(ctx.BlockHeight()).Bytes(), []byte(op.Chain), []byte(op.Token.Contract), []byte(op.Token.Amount)))

	details, err := cosmostypes.NewAnyWithValue(op)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error parsing details %s", err.Error())
	}

	var operation = types.Operation{
		Index:         index,
		OperationType: types.OpType_FEE_TOKEN_MANAGEMENT,
		Details:       details,
		Status:        types.OpStatus_INITIALIZED,
		Creator:       SystemCreator,
		Timestamp:     uint64(ctx.BlockTime().Unix()),
	}

	if op, ok := k.GetOperation(ctx, index); ok {
		if op.Status != types.OpStatus_SIGNED {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "to change operation it should be signed")
		}
	}

	k.SetOperation(
		ctx,
		operation,
	)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeNewOperation,
		sdk.NewAttribute(types.AttributeKeyOperationId, operation.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, types.OpType_FEE_TOKEN_MANAGEMENT.String()),
	))

	if err := k.ApproveOperation(ctx, operation); err != nil {
		return sdkerrors.Wrap(err, "failed to auto-approve operation")
	}

	return nil
}
