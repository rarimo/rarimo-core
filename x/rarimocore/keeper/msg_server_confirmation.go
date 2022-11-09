package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	merkle "gitlab.com/rarify-protocol/go-merkle"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operation"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/pkg"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func (k msgServer) CreateConfirmation(goCtx context.Context, msg *types.MsgCreateConfirmation) (*types.MsgCreateConfirmationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := crypto.VerifyECDSA(msg.SignatureECDSA, msg.Root, k.GetKeyECDSA(ctx)); err != nil {
		return nil, err
	}

	if err := k.checkSenderIsAParty(ctx, msg.Creator); err != nil {
		return nil, err
	}

	if _, isFound := k.GetConfirmation(ctx, msg.Root); isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	operations := make([]types.Operation, 0, len(msg.Indexes))
	contents := make([]merkle.Content, 0, len(msg.Indexes))

	for _, index := range msg.Indexes {
		operation, ok := k.GetOperation(ctx, index)
		if !ok {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("deposit %s not found", index))
		}

		if operation.Signed {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("deposit %s is already signed", index))
		}

		operations = append(operations, operation)

		content, err := k.getContent(ctx, operation)
		if err != nil {
			return nil, err
		}

		contents = append(contents, content)
	}

	if err := crypto.VerifyMerkleRoot(contents, msg.Root); err != nil {
		return nil, err
	}

	var confirmation = types.Confirmation{
		Creator:        msg.Creator,
		Root:           msg.Root,
		Indexes:        msg.Indexes,
		SignatureECDSA: msg.SignatureECDSA,
	}

	for _, op := range operations {
		op.Signed = true
		k.SetOperation(ctx, op)
	}

	k.SetConfirmation(
		ctx,
		confirmation,
	)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeNewConfirmation,
		sdk.NewAttribute(types.AttributeKeyConfirmationId, msg.Root),
	))

	return &types.MsgCreateConfirmationResponse{}, nil
}

func (k *Keeper) checkSenderIsAParty(ctx sdk.Context, sender string) error {
	for _, party := range k.GetParams(ctx).Parties {
		if party.Account == sender {
			return nil
		}
	}

	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "sender is not a party")
}

func (k *Keeper) applyOperation(ctx sdk.Context, op types.Operation) {
	switch op.OperationType {
	case types.OpType_TRANSFER:
		// Nothing to do
	case types.OpType_CHANGE_KEY:
		// Change params
		change, _ := pkg.GetChangeKey(op)
		params := k.GetParams(ctx)
		params.KeyECDSA = change.NewKey
		params.Parties = change.Parties
		params.Threshold = change.Threshold
		k.SetParams(ctx, params)
	default:
		// Nothing to do
	}

	op.Signed = true
	k.SetOperation(ctx, op)
}

func (k *Keeper) getContent(ctx sdk.Context, op types.Operation) (merkle.Content, error) {
	switch op.OperationType {
	case types.OpType_TRANSFER:
		transfer, err := pkg.GetTransfer(op)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "failed to unmarshal details")
		}

		return k.getTransferOperationContent(ctx, transfer)
	case types.OpType_CHANGE_KEY:
		change, err := pkg.GetChangeKey(op)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "failed to unmarshal details")
		}

		return pkg.GetChangeKeyContent(change)
	default:
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid operation")
	}
}

func (k *Keeper) getTransferOperationContent(ctx sdk.Context, transfer *types.Transfer) (*operation.TransferContent, error) {
	item, ok := k.tm.GetItemByChain(ctx, transfer.TokenIndex, transfer.ToChain)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "token item not found")
	}

	chainParams, ok := k.tm.GetParams(ctx).Networks[transfer.ToChain]
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("requested network not found: %s", transfer.ToChain))
	}

	return pkg.GetTransferContent(&item, chainParams, transfer)
}
