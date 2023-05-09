package keeper

import (
	cosmostypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/pkg"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
	tokenmanagertypes "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k Keeper) CreateContractUpgradeOperation(ctx sdk.Context, upgradeDetails *tokenmanagertypes.ContractUpgradeDetails) error {
	network, ok := k.tm.GetNetwork(ctx, upgradeDetails.Chain)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "network not found")
	}

	upgrade := &types.ContractUpgrade{
		Chain:                     upgradeDetails.Chain,
		NewImplementationContract: upgradeDetails.NewImplementationContract,
		Hash:                      upgradeDetails.Hash,
		BufferAccount:             upgradeDetails.BufferAccount,
		Nonce:                     upgradeDetails.Nonce,
		Type:                      upgradeDetails.Type,
	}

	content, err := pkg.GetContractUpgradeContent(network, upgrade)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error creating content %s", err.Error())
	}

	details, err := cosmostypes.NewAnyWithValue(upgrade)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error parsing details %s", err.Error())
	}

	var operation = types.Operation{
		Index:         hexutil.Encode(content.CalculateHash()),
		OperationType: types.OpType_CONTRACT_UPGRADE,
		Details:       details,
		Status:        types.OpStatus_INITIALIZED,
		Creator:       SystemCreator,
		Timestamp:     uint64(ctx.BlockTime().Unix()),
	}

	if op, ok := k.GetOperation(ctx, operation.Index); ok {
		if op.Status != types.OpStatus_SIGNED {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "to change operation it should be signed")
		}
	}

	k.SetOperation(
		ctx,
		operation,
	)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeNewOperation,
		sdk.NewAttribute(types.AttributeKeyOperationId, operation.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, types.OpType_FEE_TOKEN_MANAGEMENT.String()),
	))

	if err := k.ApproveOperation(ctx, operation); err != nil {
		return sdkerrors.Wrap(err, "failed to auto-approve operation")
	}

	return nil
}
