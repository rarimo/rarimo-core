package keeper

import (
	"fmt"
	"math/big"

	"cosmossdk.io/errors"
	cosmostypes "github.com/cosmos/cosmos-sdk/codec/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	oracletypes "github.com/rarimo/rarimo-core/x/oraclemanager/types"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/origin"
	"github.com/rarimo/rarimo-core/x/rarimocore/crypto/pkg"
	tokenmanagertypes "github.com/rarimo/rarimo-core/x/tokenmanager/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/rarimocore/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey
		memKey   storetypes.StoreKey
		tm       types.TokenmanagerKeeper
		staking  types.StakingKeeper
		bank     types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	tmkeeper types.TokenmanagerKeeper,
	staking types.StakingKeeper,
	bank types.BankKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		tm:       tmkeeper,
		staking:  staking,
		bank:     bank,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) CreateContractUpgradeOperation(ctx sdk.Context, upgradeDetails *tokenmanagertypes.ContractUpgradeDetails) error {
	network, ok := k.tm.GetNetwork(ctx, upgradeDetails.Chain)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "network not found")
	}

	upgrade := &types.ContractUpgrade{
		TargetContract:            upgradeDetails.TargetContract,
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

	// Index is HASH(block, content [depending on chain])
	index := hexutil.Encode(crypto.Keccak256(big.NewInt(ctx.BlockHeight()).Bytes(), content.CalculateHash()))

	var operation = types.Operation{
		Index:         index,
		OperationType: types.OpType_CONTRACT_UPGRADE,
		Details:       details,
		Status:        types.OpStatus_INITIALIZED,
		Creator:       types.ModuleName,
		Timestamp:     uint64(ctx.BlockTime().Unix()),
	}

	// Mostly impossible case cause operation index depends on block height.
	if _, ok := k.GetOperation(ctx, operation.Index); ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "can not recreate operation with type: %s", operation.OperationType.String())
	}

	k.SetOperation(
		ctx,
		operation,
	)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeNewOperation,
		sdk.NewAttribute(types.AttributeKeyOperationId, operation.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, operation.OperationType.String()),
	))

	// Operation is auto-approved (cause created by proposal)
	if err := k.ApproveOperation(ctx, operation); err != nil {
		return sdkerrors.Wrap(err, "failed to auto-approve operation")
	}

	return nil
}

func (k Keeper) CreateFeeTokenManagementOperation(ctx sdk.Context, op *types.FeeTokenManagement) error {
	// Index is HASH(block height, chain, fee token contract, fee amount)
	index := hexutil.Encode(crypto.Keccak256(big.NewInt(ctx.BlockHeight()).Bytes(), []byte(op.Chain), []byte(op.Token.Contract), []byte(op.Token.Amount)))

	details, err := cosmostypes.NewAnyWithValue(op)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error parsing details %s", err.Error())
	}

	var operation = types.Operation{
		Index:         index,
		OperationType: types.OpType_FEE_TOKEN_MANAGEMENT,
		Details:       details,
		Status:        types.OpStatus_INITIALIZED,
		Creator:       types.ModuleName,
		Timestamp:     uint64(ctx.BlockTime().Unix()),
	}

	// Mostly impossible case cause operation index depends on block height.
	if _, ok := k.GetOperation(ctx, operation.Index); ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "can not recreate operation with type: %s", operation.OperationType.String())
	}

	k.SetOperation(
		ctx,
		operation,
	)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeNewOperation,
		sdk.NewAttribute(types.AttributeKeyOperationId, operation.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, operation.OperationType.String()),
	))

	// Operation is auto-approved (cause created by proposal)
	if err := k.ApproveOperation(ctx, operation); err != nil {
		return sdkerrors.Wrap(err, "failed to auto-approve operation")
	}

	return nil
}

func (k Keeper) CreateChangePartiesOperation(ctx sdk.Context, creator string, change *types.ChangeParties) error {
	details, err := cosmostypes.NewAnyWithValue(change)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error parsing details %s", err.Error())
	}

	content, _ := pkg.GetChangePartiesContent(change)

	var operation = types.Operation{
		Index:         hexutil.Encode(content.CalculateHash()),
		OperationType: types.OpType_CHANGE_PARTIES,
		Details:       details,
		Status:        types.OpStatus_APPROVED, // Auto approve without event emitting
		Creator:       creator,
		Timestamp:     uint64(ctx.BlockTime().Unix()),
	}

	if _, ok := k.GetOperation(ctx, operation.Index); ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "can not recreate operation with type: %s", operation.OperationType.String())
	}

	k.SetOperation(
		ctx,
		operation,
	)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeNewOperation,
		sdk.NewAttribute(types.AttributeKeyOperationId, operation.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, types.OpType_CHANGE_PARTIES.String()),
	))

	return nil
}

func (k Keeper) CreateTransferOperation(ctx sdk.Context, creator string, transfer *types.Transfer, approved bool) error {
	network, ok := k.tm.GetNetwork(ctx, transfer.From.Chain)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "source network not found")
	}

	if network.GetBridgeParams() == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "token transfers is not supported due to lack of parameters")
	}

	// Index is HASH(tx, event, chain)
	index := hexutil.Encode(crypto.Keccak256([]byte(transfer.Tx), []byte(transfer.EventId), []byte(transfer.From.Chain)))

	details, err := cosmostypes.NewAnyWithValue(transfer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error parsing details %s", err.Error())
	}

	var operation = types.Operation{
		Index:         index,
		OperationType: types.OpType_TRANSFER,
		Details:       details,
		Status:        types.OpStatus_INITIALIZED,
		Creator:       creator,
		Timestamp:     uint64(ctx.BlockTime().Unix()),
	}

	// Only not approved operation can be changed
	if op, ok := k.GetOperation(ctx, index); ok {
		if op.Status != types.OpStatus_NOT_APPROVED {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "to change operation it should be unapproved")
		}

		// Otherwise - clear votes
		k.IterateVotes(ctx, op.Index, func(vote types.Vote) (stop bool) {
			k.RemoveVote(ctx, vote.Index)
			return false
		})
	}

	k.SetOperation(
		ctx,
		operation,
	)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeNewOperation,
		sdk.NewAttribute(types.AttributeKeyOperationId, operation.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, types.OpType_TRANSFER.String()),
	))

	if approved {
		// Auto-approve operation if requested
		if err := k.ApproveOperation(ctx, operation); err != nil {
			return sdkerrors.Wrap(err, "failed to auto-approve operation")
		}
	}

	return nil
}

func (k Keeper) CreateIdentityAggregatedTransferOperation(ctx sdk.Context, creator string, transfer *types.IdentityAggregatedTransfer) (string, error) {
	details, err := cosmostypes.NewAnyWithValue(transfer)
	if err != nil {
		return "", sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error parsing details %s", err.Error())
	}

	content, err := pkg.GetIdentityAggregatedTransferContent(transfer)
	if err != nil {
		return "", sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error creating content %s", err.Error())
	}

	var operation = types.Operation{
		Index:         hexutil.Encode(content.CalculateHash()),
		OperationType: types.OpType_IDENTITY_AGGREGATED_TRANSFER,
		Details:       details,
		Status:        types.OpStatus_INITIALIZED,
		Creator:       creator,
		Timestamp:     uint64(ctx.BlockTime().Unix()),
	}

	if _, ok := k.GetOperation(ctx, operation.Index); ok {
		return "", sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "that operation can not be changed")
	}

	k.SetOperation(
		ctx,
		operation,
	)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeNewOperation,
		sdk.NewAttribute(types.AttributeKeyOperationId, operation.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, operation.OperationType.String()),
	))

	// Operation is auto-approved (cause created by EndBlock)
	if err := k.ApproveOperation(ctx, operation); err != nil {
		return "", sdkerrors.Wrap(err, "failed to auto-approve operation")
	}

	return operation.Index, nil
}

func (k Keeper) CreateIdentityDefaultTransferOperation(ctx sdk.Context, creator string, transfer *types.IdentityDefaultTransfer) error {
	network, ok := k.tm.GetNetwork(ctx, transfer.Chain)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "source network not found")
	}

	if network.GetIdentityParams() == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "identity transfers is not supported due to lack of parameters")
	}

	// TODO validate state contract using saver one in identity params

	details, err := cosmostypes.NewAnyWithValue(transfer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error parsing details %s", err.Error())
	}

	content, err := pkg.GetIdentityDefaultTransferContent(transfer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error creating content %s", err.Error())
	}

	var operation = types.Operation{
		Index:         hexutil.Encode(content.CalculateHash()),
		OperationType: types.OpType_IDENTITY_DEFAULT_TRANSFER,
		Details:       details,
		Status:        types.OpStatus_INITIALIZED,
		Creator:       creator,
		Timestamp:     uint64(ctx.BlockTime().Unix()),
	}

	// Only not approved operation can be changed
	if op, ok := k.GetOperation(ctx, operation.Index); ok {
		if op.Status != types.OpStatus_NOT_APPROVED {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "to change operation it should be unapproved")
		}

		// Otherwise - clear votes
		k.IterateVotes(ctx, op.Index, func(vote types.Vote) (stop bool) {
			k.RemoveVote(ctx, vote.Index)
			return false
		})
	}

	k.SetOperation(
		ctx,
		operation,
	)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeNewOperation,
		sdk.NewAttribute(types.AttributeKeyOperationId, operation.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, operation.OperationType.String()),
	))

	return nil
}

func (k Keeper) CreateIdentityGISTTransferOperation(ctx sdk.Context, creator string, transfer *types.IdentityGISTTransfer) error {
	network, ok := k.tm.GetNetwork(ctx, transfer.Chain)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "source network not found")
	}

	if network.GetIdentityParams() == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "identity transfers is not supported due to lack of parameters")
	}

	details, err := cosmostypes.NewAnyWithValue(transfer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error parsing details %s", err.Error())
	}

	content, err := pkg.GetIdentityGISTTransferContent(transfer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error creating content %s", err.Error())
	}

	var operation = types.Operation{
		Index:         hexutil.Encode(content.CalculateHash()),
		OperationType: types.OpType_IDENTITY_GIST_TRANSFER,
		Details:       details,
		Status:        types.OpStatus_INITIALIZED,
		Creator:       creator,
		Timestamp:     uint64(ctx.BlockTime().Unix()),
	}

	// Only not approved operation can be changed
	if op, ok := k.GetOperation(ctx, operation.Index); ok {
		if op.Status != types.OpStatus_NOT_APPROVED {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "to change operation it should be unapproved")
		}

		// Otherwise - clear votes
		k.IterateVotes(ctx, op.Index, func(vote types.Vote) (stop bool) {
			k.RemoveVote(ctx, vote.Index)
			return false
		})
	}

	k.SetOperation(
		ctx,
		operation,
	)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeNewOperation,
		sdk.NewAttribute(types.AttributeKeyOperationId, operation.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, operation.OperationType.String()),
	))

	return nil
}

func (k Keeper) CreateIdentityStateTransferOperation(ctx sdk.Context, creator string, transfer *types.IdentityStateTransfer) error {
	network, ok := k.tm.GetNetwork(ctx, transfer.Chain)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "source network not found")
	}

	if network.GetIdentityParams() == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "identity transfers is not supported due to lack of parameters")
	}

	details, err := cosmostypes.NewAnyWithValue(transfer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error parsing details %s", err.Error())
	}

	content, err := pkg.GetIdentityStateTransferContent(transfer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error creating content %s", err.Error())
	}

	var operation = types.Operation{
		Index:         hexutil.Encode(content.CalculateHash()),
		OperationType: types.OpType_IDENTITY_STATE_TRANSFER,
		Details:       details,
		Status:        types.OpStatus_INITIALIZED,
		Creator:       creator,
		Timestamp:     uint64(ctx.BlockTime().Unix()),
	}

	// Only not approved operation can be changed
	if op, ok := k.GetOperation(ctx, operation.Index); ok {
		if op.Status != types.OpStatus_NOT_APPROVED {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "to change operation it should be unapproved")
		}

		// Otherwise - clear votes
		k.IterateVotes(ctx, op.Index, func(vote types.Vote) (stop bool) {
			k.RemoveVote(ctx, vote.Index)
			return false
		})
	}

	k.SetOperation(
		ctx,
		operation,
	)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeNewOperation,
		sdk.NewAttribute(types.AttributeKeyOperationId, operation.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, operation.OperationType.String()),
	))

	return nil
}

func (k Keeper) CreateWorldCoinIdentityTransferOperation(ctx sdk.Context, creator string, transfer *types.WorldCoinIdentityTransfer) error {
	details, err := cosmostypes.NewAnyWithValue(transfer)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "error parsing details: %s", err.Error())
	}

	content, err := pkg.GetWorldCoinIdentityTransferContent(transfer)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "error creating content: %s", err.Error())
	}

	operation := types.Operation{
		Index:         hexutil.Encode(content.CalculateHash()),
		OperationType: types.OpType_WORLDCOIN_IDENTITY_TRANSFER,
		Details:       details,
		Status:        types.OpStatus_INITIALIZED,
		Creator:       creator,
		Timestamp:     uint64(ctx.BlockTime().Unix()),
	}

	if op, ok := k.GetOperation(ctx, operation.Index); ok {
		if op.Status != types.OpStatus_NOT_APPROVED {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "to change operation it should be unapproved")
		}

		k.IterateVotes(ctx, op.Index, func(vote types.Vote) (stop bool) {
			k.RemoveVote(ctx, vote.Index)
			return false
		})
	}

	k.SetOperation(ctx, operation)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeNewOperation,
		sdk.NewAttribute(types.AttributeKeyOperationId, operation.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, operation.OperationType.String()),
	))

	return nil
}

func (k Keeper) CreateCSCARootUpdateOperation(ctx sdk.Context, creator string, update *types.CSCARootUpdate) (string, error) {
	details, err := cosmostypes.NewAnyWithValue(update)
	if err != nil {
		return "", errors.Wrapf(sdkerrors.ErrInvalidRequest, "error parsing details: %s", err.Error())
	}

	content, err := pkg.GetCSCARootUpdateContent(update)
	if err != nil {
		return "", errors.Wrapf(sdkerrors.ErrInvalidRequest, "error creating content: %s", err.Error())
	}

	operation := types.Operation{
		Index:         hexutil.Encode(content.CalculateHash()),
		OperationType: types.OpType_CSCA_ROOT_UPDATE,
		Details:       details,
		Status:        types.OpStatus_INITIALIZED,
		Creator:       creator,
		Timestamp:     uint64(ctx.BlockTime().Unix()),
	}

	if _, ok := k.GetOperation(ctx, operation.Index); ok {
		return "", errors.Wrapf(sdkerrors.ErrInvalidRequest, "that operation can not be changed")
	}
	k.SetOperation(ctx, operation)

	ctx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeNewOperation,
		sdk.NewAttribute(types.AttributeKeyOperationId, operation.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, operation.OperationType.String()),
	))

	// Operation is auto-approved (cause created by EndBlock)
	if err = k.ApproveOperation(ctx, operation); err != nil {
		return "", errors.Wrap(err, "failed to auto-approve operation")
	}

	return operation.Index, nil
}

func (k Keeper) GetTransfer(ctx sdk.Context, msg *oracletypes.MsgCreateTransferOp) (*types.Transfer, error) {
	hash := origin.NewDefaultOriginBuilder().
		SetTxHash(msg.Tx).
		SetOpId(msg.EventId).
		SetCurrentNetwork(msg.From.Chain).
		Build().
		GetOrigin()

	currentData, ok := k.tm.GetCollectionData(ctx, &tokenmanagertypes.CollectionDataIndex{Chain: msg.From.Chain, Address: msg.From.Address})
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "collection data not found")
	}

	targetData, ok := k.tm.GetCollectionData(ctx, &tokenmanagertypes.CollectionDataIndex{Chain: msg.To.Chain, Address: msg.To.Address})
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "collection data not found")
	}

	if _, ok = k.tm.GetOnChainItem(ctx, &msg.From); !ok && msg.Meta == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "metadata should be provided")
	}

	return &types.Transfer{
		Origin:     hexutil.Encode(hash[:]),
		Tx:         msg.Tx,
		EventId:    msg.EventId,
		Sender:     msg.Sender,
		Receiver:   msg.Receiver,
		Amount:     castAmount(msg.Amount, uint8(currentData.Decimals), uint8(targetData.Decimals)),
		BundleData: msg.BundleData,
		BundleSalt: msg.BundleSalt,
		From:       msg.From,
		To:         msg.To,
		Meta:       msg.Meta,
	}, nil
}

func (k Keeper) GetIdentityDefaultTransfer(_ sdk.Context, msg *oracletypes.MsgCreateIdentityDefaultTransferOp) (*types.IdentityDefaultTransfer, error) {
	return &types.IdentityDefaultTransfer{
		Contract:                msg.Contract,
		Chain:                   msg.Chain,
		GISTHash:                msg.GISTHash,
		Id:                      msg.Id,
		StateHash:               msg.StateHash,
		StateCreatedAtTimestamp: msg.StateCreatedAtTimestamp,
		StateCreatedAtBlock:     msg.StateCreatedAtBlock,
		StateReplacedBy:         msg.StateReplacedBy,
		GISTReplacedBy:          msg.GISTReplacedBy,
		GISTCreatedAtTimestamp:  msg.GISTCreatedAtTimestamp,
		GISTCreatedAtBlock:      msg.GISTCreatedAtBlock,
		ReplacedStateHash:       msg.ReplacedStateHash,
		ReplacedGISTHash:        msg.ReplacedGISTtHash,
	}, nil
}

func (k Keeper) GetIdentityGISTTransfer(_ sdk.Context, msg *oracletypes.MsgCreateIdentityGISTTransferOp) (*types.IdentityGISTTransfer, error) {
	return &types.IdentityGISTTransfer{
		Contract:               msg.Contract,
		Chain:                  msg.Chain,
		GISTHash:               msg.GISTHash,
		GISTCreatedAtTimestamp: msg.GISTCreatedAtTimestamp,
		GISTCreatedAtBlock:     msg.GISTCreatedAtBlock,
		ReplacedGISTHash:       msg.ReplacedGISTtHash,
	}, nil
}

func (k Keeper) GetIdentityStateTransfer(_ sdk.Context, msg *oracletypes.MsgCreateIdentityStateTransferOp) (*types.IdentityStateTransfer, error) {
	return &types.IdentityStateTransfer{
		Contract:                msg.Contract,
		Chain:                   msg.Chain,
		Id:                      msg.Id,
		StateHash:               msg.StateHash,
		StateCreatedAtTimestamp: msg.StateCreatedAtTimestamp,
		StateCreatedAtBlock:     msg.StateCreatedAtBlock,
		ReplacedStateHash:       msg.ReplacedStateHash,
	}, nil
}

func (k Keeper) GetWorldCoinIdentityTransfer(
	_ sdk.Context,
	msg *oracletypes.MsgCreateWorldCoinIdentityTransferOp,
) (*types.WorldCoinIdentityTransfer, error) {

	return &types.WorldCoinIdentityTransfer{
		Contract:    msg.Contract,
		Chain:       msg.Chain,
		PrevState:   msg.PrevState,
		State:       msg.State,
		Timestamp:   msg.Timestamp,
		BlockNumber: msg.BlockNumber,
	}, nil
}

func castAmount(currentAmount string, currentDecimals uint8, targetDecimals uint8) string {
	if currentDecimals == targetDecimals {
		return currentAmount
	}

	value, _ := new(big.Int).SetString(currentAmount, 10)

	if currentDecimals < targetDecimals {
		for i := uint8(0); i < targetDecimals-currentDecimals; i++ {
			value.Mul(value, new(big.Int).SetInt64(10))
		}

		return value.String()
	}

	for i := uint8(0); i < currentDecimals-targetDecimals; i++ {
		value.Div(value, new(big.Int).SetInt64(10))
	}

	return value.String()
}

func getPartyByAccount(account string, parties []*types.Party) *types.Party {
	for _, party := range parties {
		if party.Account == account {
			return party
		}
	}

	return nil
}

func getActivePartiesAmount(parties []*types.Party) int {
	activePartyCount := 0

	for _, party := range parties {
		if party.Status == types.PartyStatus_Active {
			activePartyCount++
		}
	}

	return activePartyCount
}
