package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
	tokentypes "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k Keeper) CreateAddFeeTokenOperation(ctx sdk.Context, token tokentypes.FeeToken, chain string, nonce string) error {
	if _, ok := k.tm.GetNetwork(ctx, chain); !ok {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "network not found")
	}

	op := &types.FeeTokenManagement{
		OpType:   types.FeeTokenManagementType_ADD_FEE_TOKEN,
		Token:    token,
		Chain:    chain,
		Receiver: "",
		Nonce:    nonce,
	}

	return k.CreateFeeTokenManagementOperation(ctx, op)
}

func (k Keeper) CreateRemoveFeeTokenOperation(ctx sdk.Context, token tokentypes.FeeToken, chain string, nonce string) error {
	if _, ok := k.tm.GetNetwork(ctx, chain); !ok {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "network not found")
	}

	op := &types.FeeTokenManagement{
		OpType:   types.FeeTokenManagementType_REMOVE_FEE_TOKEN,
		Token:    token,
		Chain:    chain,
		Receiver: "",
		Nonce:    nonce,
	}

	return k.CreateFeeTokenManagementOperation(ctx, op)
}

func (k Keeper) CreateUpdateFeeTokenOperation(ctx sdk.Context, token tokentypes.FeeToken, chain string, nonce string) error {
	if _, ok := k.tm.GetNetwork(ctx, chain); !ok {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "network not found")
	}

	op := &types.FeeTokenManagement{
		OpType:   types.FeeTokenManagementType_UPDATE_FEE_TOKEN,
		Token:    token,
		Chain:    chain,
		Receiver: "",
		Nonce:    nonce,
	}

	return k.CreateFeeTokenManagementOperation(ctx, op)
}

func (k Keeper) CreateWithdrawFeeOperation(ctx sdk.Context, token tokentypes.FeeToken, chain string, receiver string, nonce string) error {
	op := &types.FeeTokenManagement{
		OpType:   types.FeeTokenManagementType_WITHDRAW_FEE_TOKEN,
		Token:    token,
		Chain:    chain,
		Receiver: receiver,
		Nonce:    nonce,
	}

	return k.CreateFeeTokenManagementOperation(ctx, op)
}
