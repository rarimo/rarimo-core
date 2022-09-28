package keeper

import (
	"context"
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	merkle "gitlab.com/rarify-protocol/go-merkle"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
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

		info, ok := k.tm.GetInfo(ctx, deposit.TokenIndex)
		if !ok {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("token info %s not found", deposit.Index))
		}

		chainParams, ok := k.tm.GetParams(ctx).Networks[deposit.ToChain]
		if !ok {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("deposit network not found: %s", deposit.ToChain))
		}

		content = append(content, crypto.HashContent{
			TxHash:         deposit.Tx,
			EventId:        deposit.EventId,
			TargetNetwork:  deposit.ToChain,
			CurrentNetwork: deposit.FromChain,
			Receiver:       hexutil.MustDecode(deposit.Receiver),
			TargetAddress:  tryHexDecode(info.Chains[deposit.ToChain].TokenAddress, []byte{}),
			TargetId:       tryHexDecode(info.Chains[deposit.ToChain].TokenId, []byte{}),
			Amount:         amountBytes(deposit.Amount),
			ProgramId:      hexutil.MustDecode(chainParams.Contract),
			Data:           append([]byte(info.Name), append([]byte(info.Symbol), []byte(info.Uri)...)...),
		})
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

func amountBytes(amount string) []byte {
	am, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		// it is NFT
		am = new(big.Int).SetInt64(1)
	}

	amBytes := am.Bytes()
	result := make([]byte, 32)

	for i := range amBytes {
		result[31-i] = amBytes[len(amBytes)-1-i]
	}
	return result
}

func tryHexDecode(hexStr string, defResp []byte) []byte {
	resp, err := hexutil.Decode(hexStr)
	if err != nil {
		return defResp
	}

	return resp
}
