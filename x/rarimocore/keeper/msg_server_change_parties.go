package keeper

import (
	"context"

	cosmostypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	eth "github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operation"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func (k msgServer) CreateChangePartiesOperation(goCtx context.Context, msg *types.MsgCreateChangePartiesOp) (*types.MsgCreateChangePartiesOpResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	defer k.disableFee(ctx.GasMeter().GasConsumed(), ctx.GasMeter())

	if err := k.checkSenderIsAParty(ctx, msg.Creator); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "only parties can submit that transaction")
	}

	var changeOp = types.ChangeParties{
		CurrentSet: k.GetParams(ctx).Parties,
		NewSet:     msg.NewSet,
	}

	changeType, err := getChangeType(changeOp.CurrentSet, changeOp.NewSet)
	if err != nil {
		return nil, err
	}

	changeOp.Type = changeType

	details, err := cosmostypes.NewAnyWithValue(&changeOp)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error parsing details %s", err.Error())
	}

	var op = types.Operation{
		Index:         getIndex(changeOp.CurrentSet, changeOp.NewSet),
		OperationType: types.OpType_CHANGE_PARTIES,
		Details:       details,
		Signed:        false,
		Creator:       msg.Creator,
		Timestamp:     ctx.BlockTime().Unix(),
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

func getIndex(old []*types.Party, new []*types.Party) string {
	return hexutil.Encode(eth.Keccak256(operation.GetPartiesHash(old), operation.GetPartiesHash(new)))
}

func getChangeType(old []*types.Party, new []*types.Party) (types.ChangeType, error) {
	intersection := make(map[string]struct{})

	// working in n^2 possible cause small set
	for _, pOld := range old {
		for _, pNew := range new {
			if pOld == pNew {
				if _, ok := intersection[pOld.Address]; ok {
					return 0, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid set: contains duplicates")
				}

				intersection[pOld.Address] = struct{}{}
			}
		}
	}

	if len(old) != len(intersection) && len(old)+1 != len(intersection) {
		return 0, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid set: cannot define change - invalid intersection size")
	}

	if len(old) == len(intersection) {
		for _, pNew := range new {
			if _, ok := intersection[pNew.Address]; !ok {
				return types.ChangeType_ADD, nil
			}
		}
	}

	if len(old)+1 == len(intersection) {
		for _, pOld := range old {
			if _, ok := intersection[pOld.Address]; !ok {
				return types.ChangeType_REMOVE, nil
			}
		}
	}

	return 0, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid set: cannot define change")
}
