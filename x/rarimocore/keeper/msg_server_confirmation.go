package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gogo/protobuf/proto"
	merkle "gitlab.com/rarify-protocol/go-merkle"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/bundle"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operation"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/origin"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
	tokentypes "gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/types"
)

func (k msgServer) CreateConfirmation(goCtx context.Context, msg *types.MsgCreateConfirmation) (*types.MsgCreateConfirmationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := crypto.VerifyECDSA(msg.SignatureECDSA, msg.Root, k.GetKeyECDSA(ctx)); err != nil {
		return nil, err
	}

	_, isFound := k.GetConfirmation(
		ctx,
		msg.Root,
	)
	if isFound {
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

	return &types.MsgCreateConfirmationResponse{}, nil
}

func (k *Keeper) getContent(ctx sdk.Context, op types.Operation) (crypto.HashContent, error) {
	switch op.OperationType {
	case types.OpType_TRANSFER:
		transfer := types.Transfer{}
		if err := proto.Unmarshal(op.Details.Value, &transfer); err != nil {
			return crypto.HashContent{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "failed to unmarshal details")
		}
		return k.transferOperationContent(ctx, transfer)
	default:
		return crypto.HashContent{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "undefined details")
	}
}

func (k *Keeper) transferOperationContent(ctx sdk.Context, transfer types.Transfer) (crypto.HashContent, error) {
	info, ok := k.tm.GetInfo(ctx, transfer.TokenIndex)
	if !ok {
		return crypto.HashContent{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "token info not found")
	}

	item, ok := k.tm.GetItem(ctx, info.Chains[transfer.ToChain].TokenAddress, info.Chains[transfer.ToChain].TokenId, transfer.ToChain)
	if !ok {
		return crypto.HashContent{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "token item not found")
	}

	chainParams, ok := k.tm.GetParams(ctx).Networks[transfer.ToChain]
	if !ok {
		return crypto.HashContent{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("requested network not found: %s", transfer.ToChain))
	}

	var op operation.Operation

	switch item.TokenType {
	case tokentypes.Type_METAPLEX_FT | tokentypes.Type_METAPLEX_NFT:
		op = operation.NewTransferFullMetaOperation(
			info.Chains[transfer.ToChain].TokenAddress,
			info.Chains[transfer.ToChain].TokenId,
			transfer.Amount, item.Name, item.Symbol, item.Uri, uint8(item.Decimals),
		)
	default:
		op = operation.NewTransferOperation(
			info.Chains[transfer.ToChain].TokenAddress,
			info.Chains[transfer.ToChain].TokenId,
			transfer.Amount, item.Uri,
		)
	}

	return crypto.HashContent{
		Origin:         origin.NewDefaultOrigin(transfer.Tx, transfer.EventId, transfer.FromChain).GetOrigin(),
		TargetNetwork:  transfer.ToChain,
		Receiver:       hexutil.MustDecode(transfer.Receiver),
		Data:           op.GetContent(),
		TargetContract: hexutil.MustDecode(chainParams.Contract),
		Bundle:         bundle.NewDefaultBundle(transfer.BundleSalt, transfer.BundleData).GetBundle(),
	}, nil
}
