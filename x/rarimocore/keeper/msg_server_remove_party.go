package keeper

import (
	"context"
	"math/big"

	cosmostypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	eth "github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operation"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func (k msgServer) CreateRemovePartyOperation(goCtx context.Context, msg *types.MsgCreateRemovePartyOp) (*types.MsgCreateRemovePartyOpResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	defer k.disableFee(ctx.GasMeter().GasConsumed(), ctx.GasMeter())

	if err := k.checkSenderIsAParty(ctx, msg.Creator); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "only parties can submit that transaction")
	}

	var removeOp = types.RemoveParty{
		CurrentSet: k.GetParams(ctx).Parties,
		PartyIndex: msg.Index,
	}

	details, err := cosmostypes.NewAnyWithValue(&removeOp)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error parsing details %s", err.Error())
	}

	var op = types.Operation{
		Index:         hexutil.Encode(eth.Keccak256(operation.GetPartiesHash(k.GetParams(ctx).Parties), big.NewInt(int64(removeOp.PartyIndex)).Bytes())),
		OperationType: types.OpType_REMOVE_PARTY,
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
		sdk.NewAttribute(types.AttributeKeyOperationType, types.OpType_REMOVE_PARTY.String()),
	))

	return &types.MsgCreateRemovePartyOpResponse{}, nil
}
