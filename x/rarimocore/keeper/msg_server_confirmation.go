package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	merkle "gitlab.com/rarify-protocol/go-merkle"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
	tokentypes "gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/types"
)

func (k msgServer) CreateConfirmation(goCtx context.Context, msg *types.MsgCreateConfirmation) (*types.MsgCreateConfirmationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := crypto.VerifyECDSA(msg.SigECDSA, msg.Root, k.GetKeyECDSA(ctx)); err != nil {
		return nil, err
	}

	// Check if the value already exists
	_, isFound := k.GetConfirmation(
		ctx,
		msg.Root,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	deposits := make([]types.Deposit, 0, len(msg.Hashes))
	infos := make([]tokentypes.Info, 0, len(msg.Hashes))
	content := make([]merkle.Content, 0, len(msg.Hashes))

	for _, hash := range msg.Hashes {
		deposit, ok := k.GetDeposit(ctx, hash)
		if !ok {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("deposit %s not found", hash))
		}

		if deposit.Signed {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("deposit %s is already signed", hash))
		}

		deposits = append(deposits, deposit)

		item, ok := k.tm.GetItem(ctx, deposit.TokenAddress, deposit.TokenId)
		if !ok {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("token item %s %s not dound", deposit.TokenAddress, deposit.TokenId))
		}

		info, ok := k.tm.GetInfo(ctx, item.Index)
		if !ok {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("token info %s not dound", item.Index))
		}

		infos = append(infos, info)

		content = append(content, crypto.HashContent{
			TxHash:         hash,
			EventId:        deposit.EventId,
			TargetNetwork:  deposit.ToChain,
			CurrentNetwork: deposit.FromChain,

			Receiver:      hexutil.MustDecode(deposit.Receiver),
			TargetAddress: hexutil.MustDecode(info.Chains[deposit.ToChain].TokenAddress),
			TargetId:      hexutil.MustDecode(info.Chains[deposit.ToChain].TokenId),
		})

	}

	if err := crypto.VerifyMerkleRoot(content, msg.Root); err != nil {
		return nil, err
	}

	var confirmation = types.Confirmation{
		Creator:  msg.Creator,
		Root:     msg.Root,
		Hashes:   msg.Hashes,
		SigECDSA: msg.SigECDSA,
	}

	for _, deposit := range deposits {
		deposit.Signed = true
		k.SetDeposit(ctx, deposit)
	}

	for i, info := range infos {
		info.CurrentChain = deposits[i].ToChain
		k.tm.SetInfo(ctx, info)
	}

	k.SetConfirmation(
		ctx,
		confirmation,
	)

	return &types.MsgCreateConfirmationResponse{}, nil
}
