package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	eth "github.com/ethereum/go-ethereum/crypto"
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
		op, ok := k.GetOperation(ctx, index)
		if !ok {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("deposit %s not found", index))
		}

		if op.Signed {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("deposit %s is already signed", index))
		}
		operations = append(operations, op)

		content, err := k.getContent(ctx, op)
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
		err := k.applyOperation(ctx, op)
		if err != nil {
			return nil, err
		}
	}

	k.SetConfirmation(
		ctx,
		confirmation,
	)

	k.UpdateLastSignature(ctx, confirmation.SignatureECDSA)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeNewConfirmation,
		sdk.NewAttribute(types.AttributeKeyConfirmationId, msg.Root),
	))

	return &types.MsgCreateConfirmationResponse{}, nil
}

func (k msgServer) applyOperation(ctx sdk.Context, op types.Operation) error {
	switch op.OperationType {
	case types.OpType_TRANSFER:
		transfer, _ := pkg.GetTransfer(op)
		if err := k.applyTransfer(ctx, transfer); err != nil {
			return err
		}
	case types.OpType_CHANGE_PARTIES:
		change, _ := pkg.GetChangeParties(op)
		if err := k.applyChangeParties(ctx, change); err != nil {
			return err
		}
		ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeParamsUpdated,
			sdk.NewAttribute(types.AttributeKeyParamsUpdateType, types.ParamsUpdateType_CHANGE_SET.String()),
		))
	default:
		// Nothing to do
	}

	op.Signed = true
	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeOperationSigned,
		sdk.NewAttribute(types.AttributeKeyOperationId, op.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, op.OperationType.String()),
	))

	k.SetOperation(ctx, op)
	return nil
}

func (k msgServer) applyTransfer(ctx sdk.Context, _ *types.Transfer) error {
	if k.GetParams(ctx).IsUpdateRequired {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "can not apply transfer: parties update needed")
	}
	return nil
}

func (k msgServer) applyChangeParties(ctx sdk.Context, op *types.ChangeParties) error {
	params := k.GetParams(ctx)

	hash := hexutil.Encode(eth.Keccak256(hexutil.MustDecode(op.NewPublicKey)))
	if err := crypto.VerifyECDSA(op.Signature, hash, params.KeyECDSA); err != nil {
		return err
	}

	if len(params.Parties) != len(op.Parties) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid parties amount")
	}

	for i := range params.Parties {
		if op.Parties[i].Address != params.Parties[i].Address {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid parties")
		}

		params.Parties[i].PubKey = op.Parties[i].PubKey
	}

	params.KeyECDSA = op.NewPublicKey
	params.Threshold = uint64(((len(params.Parties) + 2) / 3) * 2)
	params.IsUpdateRequired = false
	k.SetParams(ctx, params)
	return nil
}

func (k msgServer) getContent(ctx sdk.Context, op types.Operation) (merkle.Content, error) {
	switch op.OperationType {
	case types.OpType_TRANSFER:
		transfer, err := pkg.GetTransfer(op)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "failed to unmarshal details")
		}

		return k.getTransferOperationContent(ctx, transfer)
	case types.OpType_CHANGE_PARTIES:
		change, err := pkg.GetChangeParties(op)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "failed to unmarshal details")
		}

		return pkg.GetChangePartiesContent(change)
	default:
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid operation")
	}
}

func (k msgServer) getTransferOperationContent(ctx sdk.Context, transfer *types.Transfer) (*operation.TransferContent, error) {
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
