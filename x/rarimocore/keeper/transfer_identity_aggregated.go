package keeper

import (
	cosmostypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/pkg"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k Keeper) CreateIdentityAggregatedTransferOperation(ctx sdk.Context, creator string, transfer *types.IdentityAggregatedTransfer) (string, error) {
	details, err := cosmostypes.NewAnyWithValue(transfer)
	if err != nil {
		return "", sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error parsing details %s", err.Error())
	}

	content, err := pkg.GetIdentityAggregatedTransferContent(transfer)
	if err != nil {
		return "", sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error creating content %s", err.Error())
	}

	var operation = types.Operation{
		Index:         hexutil.Encode(content.CalculateHash()),
		OperationType: types.OpType_IDENTITY_AGGREGATED_TRANSFER,
		Details:       details,
		Status:        types.OpStatus_INITIALIZED,
		Creator:       creator,
		Timestamp:     uint64(ctx.BlockTime().Unix()),
	}

	k.SetOperation(
		ctx,
		operation,
	)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeNewOperation,
		sdk.NewAttribute(types.AttributeKeyOperationId, operation.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, operation.OperationType.String()),
	))

	if err := k.ApproveOperation(ctx, operation); err != nil {
		return "", sdkerrors.Wrap(err, "failed to auto-approve operation")
	}

	return operation.Index, nil
}
