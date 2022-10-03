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
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operations"
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

func (k *Keeper) transferOperationContent(ctx sdk.Context, op types.Transfer) (crypto.HashContent, error) {
	info, ok := k.tm.GetInfo(ctx, op.TokenIndex)
	if !ok {
		return crypto.HashContent{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "token info not found")
	}

	item, ok := k.tm.GetItem(ctx, info.Chains[op.ToChain].TokenAddress, info.Chains[op.ToChain].TokenId, op.ToChain)
	if !ok {
		return crypto.HashContent{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "token item not found")
	}

	chainParams, ok := k.tm.GetParams(ctx).Networks[op.ToChain]
	if !ok {
		return crypto.HashContent{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("requested network not found: %s", op.ToChain))
	}

	var operation operations.Operation

	switch item.TokenType {
	case tokentypes.Type_METAPLEX_FT | tokentypes.Type_METAPLEX_NFT:
		operation = operations.NewTransferFullMetaWithBundleOperation(
			info.Chains[op.ToChain].TokenAddress,
			info.Chains[op.ToChain].TokenId,
			op.Amount, item.Name, item.Symbol, item.Uri, uint8(item.Decimals),
			"", "",
		)
	default:
		operation = operations.NewTransferWithBundleOperation(
			info.Chains[op.ToChain].TokenAddress,
			info.Chains[op.ToChain].TokenId,
			op.Amount, item.Uri,
			op.BundleData, op.BundleSalt,
		)
	}

	return crypto.HashContent{
		Origin:         origin.NewDefaultOrigin(op.Tx, op.EventId, op.FromChain).GetOrigin(),
		TargetNetwork:  op.ToChain,
		Receiver:       hexutil.MustDecode(op.Receiver),
		Data:           operation.GetContent(),
		TargetContract: hexutil.MustDecode(chainParams.Contract),
	}, nil
}
