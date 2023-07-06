package keeper

import (
	cosmostypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	oracletypes "gitlab.com/rarimo/rarimo-core/x/oraclemanager/types"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/pkg"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k Keeper) CreateIdentityDefaultTransferOperation(ctx sdk.Context, creator string, transfer *types.IdentityDefaultTransfer) error {
	details, err := cosmostypes.NewAnyWithValue(transfer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error parsing details %s", err.Error())
	}

	content, err := pkg.GetIdentityDefaultTransferContent(transfer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error creating content %s", err.Error())
	}

	var operation = types.Operation{
		Index:         hexutil.Encode(content.CalculateHash()),
		OperationType: types.OpType_TRANSFER,
		Details:       details,
		Status:        types.OpStatus_INITIALIZED,
		Creator:       creator,
		Timestamp:     uint64(ctx.BlockTime().Unix()),
	}

	if op, ok := k.GetOperation(ctx, operation.Index); ok {
		if op.Status == types.OpStatus_INITIALIZED || op.Status == types.OpStatus_APPROVED {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "to change operation it should be unapproved or signed")
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
		sdk.NewAttribute(types.AttributeKeyOperationType, operation.OperationType.String()),
	))

	return nil
}

func (k Keeper) GetIdentityDefaultTransfer(_ sdk.Context, msg *oracletypes.MsgCreateIdentityDefaultTransferOp) (*types.IdentityDefaultTransfer, error) {
	return &types.IdentityDefaultTransfer{
		Contract:                 msg.Contract,
		Chain:                    msg.Chain,
		GISTHash:                 msg.GISTHash,
		Id:                       msg.Id,
		StateHash:                msg.StateHash,
		StateCreatedAtTimestamp:  msg.StateCreatedAtTimestamp,
		StateCreatedAtBlock:      msg.StateCreatedAtBlock,
		StateReplacedAtTimestamp: msg.StateReplacedAtTimestamp,
		StateReplacedAtBlock:     msg.StateReplacedAtBlock,
		StateReplacedBy:          msg.StateReplacedBy,
		GISTReplacedBy:           msg.GISTReplacedBy,
		GISTCreatedAtTimestamp:   msg.GISTCreatedAtTimestamp,
		GISTCreatedAtBlock:       msg.GISTCreatedAtBlock,
		GISTReplacedAtTimestamp:  msg.GISTReplacedAtTimestamp,
		GISTReplacedAtBlock:      msg.GISTReplacedAtBlock,
		ReplacedStateHash:        msg.ReplacedStateHash,
	}, nil
}
