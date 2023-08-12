package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	eth "github.com/ethereum/go-ethereum/crypto"
	merkle "gitlab.com/rarimo/go-merkle"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/pkg"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
	tokentypes "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k msgServer) CreateConfirmation(goCtx context.Context, msg *types.MsgCreateConfirmation) (*types.MsgCreateConfirmationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := crypto.VerifyECDSA(msg.SignatureECDSA, msg.Root, k.GetKeyECDSA(ctx)); err != nil {
		return nil, sdkerrors.Wrap(err, "invalid signature")
	}

	if err := k.checkIsAnActiveParty(ctx, msg.Creator); err != nil {
		return nil, sdkerrors.Wrap(err, "creator is not an active party")
	}

	if _, isFound := k.GetConfirmation(ctx, msg.Root); isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	operations := make([]types.Operation, 0, len(msg.Indexes))
	contents := make([]merkle.Content, 0, len(msg.Indexes))

	for _, index := range msg.Indexes {
		op, ok := k.GetOperation(ctx, index)
		if !ok {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("operation %s not found", index))
		}

		if op.Status != types.OpStatus_APPROVED {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("operation %s is not approved", index))
		}

		operations = append(operations, op)

		content, err := k.getContent(ctx, op)
		if err != nil {
			return nil, sdkerrors.Wrap(err, "failed to get content")
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
		err := k.ApplyOperation(ctx, op, msg.Root)
		if err != nil {
			return nil, err
		}

		k.SetOperationConfirmationId(ctx, op.Index, confirmation.Root)
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

func (k Keeper) ApplyOperation(ctx sdk.Context, op types.Operation, confirmationId string) error {
	canBeAppliedByInactivePartiesSet := map[types.OpType]struct{}{
		types.OpType_CHANGE_PARTIES: {},
	}

	if _, canBeApplied := canBeAppliedByInactivePartiesSet[op.OperationType]; k.GetParams(ctx).IsUpdateRequired && !canBeApplied {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "can not apply operation: parties update needed")
	}

	switch op.OperationType {
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

	op.Status = types.OpStatus_SIGNED
	k.SetOperation(ctx, op)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeOperationSigned,
		sdk.NewAttribute(types.AttributeKeyOperationId, op.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, op.OperationType.String()),
		sdk.NewAttribute(types.AttributeKeyConfirmationId, confirmationId),
	))

	return nil
}

func (k Keeper) applyChangeParties(ctx sdk.Context, op *types.ChangeParties) error {
	params := k.GetParams(ctx)

	hash := hexutil.Encode(eth.Keccak256(hexutil.MustDecode(op.NewPublicKey)))
	if err := crypto.VerifyECDSA(op.Signature, hash, params.KeyECDSA); err != nil {
		return err
	}

	existingParties := make(map[string]*types.Party)
	for _, party := range params.Parties {
		existingParties[party.Account] = party
	}

	for _, party := range op.Parties {
		existingParty, ok := existingParties[party.Account]
		if !ok {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid parties: party does not exist")
		}

		if existingParty.Status == types.PartyStatus_Frozen || existingParty.Status == types.PartyStatus_Slashed {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid parties: can not change party status")
		}

		existingParty.PubKey = party.PubKey
		existingParty.Status = party.Status
	}

	params.KeyECDSA = op.NewPublicKey
	params.IsUpdateRequired = false
	params.Threshold = uint64(crypto.GetThreshold(getActivePartiesAmount(params.Parties)))
	k.SetParams(ctx, params)
	return nil
}

func (k Keeper) getContent(ctx sdk.Context, op types.Operation) (merkle.Content, error) {
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
		return k.getChangePartiesContent(ctx, change)
	case types.OpType_FEE_TOKEN_MANAGEMENT:
		manage, err := pkg.GetFeeTokenManagement(op)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "failed to unmarshal details")
		}
		return k.getFeeTokenManagementContent(ctx, manage)
	case types.OpType_CONTRACT_UPGRADE:
		upgrade, err := pkg.GetContractUpgrade(op)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "failed to unmarshal details")
		}
		return k.getContractUpgradeContent(ctx, upgrade)
	case types.OpType_IDENTITY_DEFAULT_TRANSFER:
		transfer, err := pkg.GetIdentityDefaultTransfer(op)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "failed to unmarshal details")
		}
		return k.getIdentityDefaultTransferContent(ctx, transfer)
	case types.OpType_IDENTITY_AGGREGATED_TRANSFER:
		transfer, err := pkg.GetIdentityAggregatedTransfer(op)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "failed to unmarshal details")
		}
		return k.getIdentityAggregatedTransferContent(ctx, transfer)
	default:
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid operation")
	}
}

func (k Keeper) getTransferOperationContent(ctx sdk.Context, transfer *types.Transfer) (*operation.TransferContent, error) {
	data, ok := k.tm.GetCollectionData(ctx, &tokentypes.CollectionDataIndex{Chain: transfer.To.Chain, Address: transfer.To.Address})
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "collection data not found")
	}

	collection, ok := k.tm.GetCollection(ctx, data.Collection)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "collection not found")
	}

	onChainItem, ok := k.tm.GetOnChainItem(ctx, &transfer.To)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "on chain item not found")
	}

	item, ok := k.tm.GetItem(ctx, onChainItem.Item)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "item not found")
	}

	network, ok := k.tm.GetNetwork(ctx, transfer.To.Chain)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "network not found")
	}

	bridgeparams := network.GetBridgeParams()
	if bridgeparams == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "network bridge params not found")
	}

	return pkg.GetTransferContent(collection, data, item, bridgeparams, transfer)
}

func (k Keeper) getFeeTokenManagementContent(ctx sdk.Context, manage *types.FeeTokenManagement) (*operation.FeeTokenManagementContent, error) {
	network, ok := k.tm.GetNetwork(ctx, manage.Chain)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "network not found")
	}

	feeparams := network.GetFeeParams()
	if feeparams == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "network fee params not found")
	}

	return pkg.GetFeeTokenManagementContent(feeparams, manage)
}

func (k Keeper) getContractUpgradeContent(ctx sdk.Context, upgrade *types.ContractUpgrade) (*operation.ContractUpgradeContent, error) {
	networkParams, ok := k.tm.GetNetwork(ctx, upgrade.Chain)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "network params not found")
	}

	return pkg.GetContractUpgradeContent(networkParams, upgrade)
}

func (k Keeper) getIdentityDefaultTransferContent(_ sdk.Context, transfer *types.IdentityDefaultTransfer) (*operation.IdentityDefaultTransferContent, error) {
	return pkg.GetIdentityDefaultTransferContent(transfer)
}

func (k Keeper) getChangePartiesContent(_ sdk.Context, change *types.ChangeParties) (*operation.ChangePartiesContent, error) {
	return pkg.GetChangePartiesContent(change)
}

func (k Keeper) getIdentityAggregatedTransferContent(_ sdk.Context, transfer *types.IdentityAggregatedTransfer) (*operation.IdentityAggregatedTransferContent, error) {
	return pkg.GetIdentityAggregatedTransferContent(transfer)
}
