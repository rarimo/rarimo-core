package keeper

import (
	"context"

	cosmostypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/pkg"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k msgServer) CreateChangePartiesOperation(goCtx context.Context, msg *types.MsgCreateChangePartiesOp) (*types.MsgCreateChangePartiesOpResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	defer k.disableFee(ctx.GasMeter().GasConsumed(), ctx.GasMeter())

	if err := k.checkIsAnActiveParty(ctx, msg.Creator); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "only active parties can submit that transaction")
	}

	var changeOp = &types.ChangeParties{
		Parties:      msg.NewSet,
		Signature:    msg.Signature,
		NewPublicKey: msg.NewPublicKey,
	}

	details, err := cosmostypes.NewAnyWithValue(changeOp)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error parsing details %s", err.Error())
	}

	content, _ := pkg.GetChangePartiesContent(changeOp)

	var op = types.Operation{
		Index:         hexutil.Encode(content.CalculateHash()),
		OperationType: types.OpType_CHANGE_PARTIES,
		Details:       details,
		Status:        types.OpStatus_APPROVED, // Auto approve
		Creator:       msg.Creator,
		Timestamp:     uint64(ctx.BlockHeight()),
	}

	k.SetOperation(
		ctx,
		op,
	)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeNewOperation,
		sdk.NewAttribute(types.AttributeKeyOperationId, op.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, types.OpType_CHANGE_PARTIES.String()),
	))

	return &types.MsgCreateChangePartiesOpResponse{}, nil
}
