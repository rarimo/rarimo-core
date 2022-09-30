package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	merkle "gitlab.com/rarify-protocol/go-merkle"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operations"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/origin"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
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

	deposits := make([]types.Deposit, 0, len(msg.Indexes))
	content := make([]merkle.Content, 0, len(msg.Indexes))

	for _, index := range msg.Indexes {
		deposit, ok := k.GetDeposit(ctx, index)
		if !ok {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("deposit %s not found", index))
		}

		if deposit.Signed {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("deposit %s is already signed", index))
		}

		deposits = append(deposits, deposit)

		depositContent, err := k.contentFromDeposit(ctx, deposit)
		if err != nil {
			return nil, err
		}

		content = append(content, depositContent)
	}

	if err := crypto.VerifyMerkleRoot(content, msg.Root); err != nil {
		return nil, err
	}

	var confirmation = types.Confirmation{
		Creator:        msg.Creator,
		Root:           msg.Root,
		Indexes:        msg.Indexes,
		SignatureECDSA: msg.SignatureECDSA,
	}

	for _, deposit := range deposits {
		deposit.Signed = true
		k.SetDeposit(ctx, deposit)
	}

	k.SetConfirmation(
		ctx,
		confirmation,
	)

	return &types.MsgCreateConfirmationResponse{}, nil
}

func (k *Keeper) contentFromDeposit(ctx sdk.Context, deposit types.Deposit) (crypto.HashContent, error) {
	info, ok := k.tm.GetInfo(ctx, deposit.TokenIndex)
	if !ok {
		return crypto.HashContent{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("token info %s not found", deposit.Index))
	}

	item, ok := k.tm.GetItem(ctx, info.Chains[deposit.ToChain].TokenAddress, info.Chains[deposit.ToChain].TokenId, deposit.ToChain)
	if !ok {
		return crypto.HashContent{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("token idem %s not found", deposit.Index))
	}

	chainParams, ok := k.tm.GetParams(ctx).Networks[deposit.ToChain]
	if !ok {
		return crypto.HashContent{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("deposit network not found: %s", deposit.ToChain))
	}

	return crypto.HashContent{
		Origin:        origin.NewDefaultOrigin(deposit.Tx, deposit.EventId, deposit.FromChain).GetOrigin(),
		TargetNetwork: deposit.ToChain,
		Receiver:      hexutil.MustDecode(deposit.Receiver),
		Data: operations.NewTransferOperation(
			info.Chains[deposit.ToChain].TokenAddress,
			info.Chains[deposit.ToChain].TokenId,
			deposit.Amount, item.Name, item.Symbol, item.Uri).GetContent(),
		TargetContract: hexutil.MustDecode(chainParams.Contract),
	}, nil
}
