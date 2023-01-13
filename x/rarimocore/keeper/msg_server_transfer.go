package keeper

import (
	"context"
	"math/big"

	"cosmossdk.io/errors"

	tokentypes "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"

	cosmostypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/origin"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/saver"
	savermsg "gitlab.com/rarimo/savers/saver-grpc-lib/grpc"
)

func (k msgServer) CreateTransferOperation(ctx context.Context, msg *types.MsgCreateTransferOp) (*types.MsgCreateTransferOpResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if k.checkCreatorIsValidator(sdkCtx, msg.Creator) {
		defer k.disableFee(sdkCtx.GasMeter().GasConsumed(), sdkCtx.GasMeter())
	}

	origin := origin.NewDefaultOriginBuilder().
		SetTxHash(msg.Tx).
		SetOpId(msg.EventId).
		SetCurrentNetwork(msg.FromChain).
		Build().
		GetOrigin()

	index := hexutil.Encode(origin[:])

	if _, ok := k.GetOperation(sdkCtx, index); ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	depositInfo, err := k.getDepositInfo(sdkCtx, msg)
	if err != nil {
		return nil, err
	}

	if _, ok := k.tm.GetNetwork(sdkCtx, depositInfo.TargetNetwork); !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "network not found: %s", depositInfo.TargetNetwork)
	}

	collection := k.tm.GetCollectionInfo(sdkCtx, depositInfo.Collection)
	if collection == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrLogic, "collection not found by index [%d]", depositInfo.Collection)
	}

	sourceChainParams, ok := collection.ChainParams[msg.FromChain]
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "chain [%s] not supported by collection", msg.FromChain)
	}

	targetChainParams, ok := collection.ChainParams[depositInfo.TargetNetwork]
	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "chain [%s] not supported by collection", depositInfo.TargetNetwork)
	}

	itemKey := tokentypes.ItemKey(sourceChainParams.Address, depositInfo.TokenId, msg.FromChain)

	_, err = k.ensureItem(ctx, itemKey, *collection, msg.FromChain, depositInfo.TokenId, collection.TokenType)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrLogic, "error creating item: %s", err.Error())
	}

	// TODO ensure that item has network to transfer from and to

	if _, err := hexutil.Decode(depositInfo.Receiver); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid receiver format")
	}

	var transferOp = types.Transfer{
		Origin:          index,
		Tx:              msg.Tx,
		EventId:         msg.EventId,
		FromChain:       msg.FromChain,
		ToChain:         depositInfo.TargetNetwork,
		Receiver:        depositInfo.Receiver,
		Amount:          castAmount(depositInfo.Amount, uint8(sourceChainParams.Decimals), uint8(targetChainParams.Decimals)),
		BundleData:      getBundle(depositInfo),
		BundleSalt:      getSalt(depositInfo),
		CollectionIndex: collection.Index,
	}

	details, err := cosmostypes.NewAnyWithValue(&transferOp)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error parsing details %s", err.Error())
	}

	var operation = types.Operation{
		Index:         index,
		OperationType: types.OpType_TRANSFER,
		Details:       details,
		Signed:        false,
		Creator:       msg.Creator,
		Timestamp:     uint64(sdkCtx.BlockHeight()),
	}

	k.SetOperation(
		sdkCtx,
		operation,
	)

	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(types.EventTypeNewOperation,
		sdk.NewAttribute(types.AttributeKeyOperationId, operation.Index),
		sdk.NewAttribute(types.AttributeKeyOperationType, types.OpType_TRANSFER.String()),
	))
	return &types.MsgCreateTransferOpResponse{}, nil
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

func (k msgServer) ensureItem(ctx context.Context,
	itemKey []byte,
	collection tokentypes.CollectionInfo,
	networkName string,
	tokenID string,
	tokenType tokentypes.Type,
) (*tokentypes.Item, error) {
	itemIdx, ok := collection.Items[string(itemKey)]
	if ok {
		item, ok := k.tm.GetItemByIndex(sdk.UnwrapSDKContext(ctx), itemIdx)
		if ok { // i suppose we should have this item stored but just in case
			return &item, nil
		}
	}

	item := tokentypes.Item{
		Index:      tokentypes.ItemKeyToIndex(itemKey),
		Collection: collection.Index,
		ChainParams: map[string]*tokentypes.ItemChainParams{
			networkName: {
				TokenID:   tokenID,
				TokenType: tokenType,
			},
		},
	}

	saverClient, err := saver.GetClient(networkName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init saver client")
	}

	itemMeta, err := saverClient.GetMetadata(ctx, &savermsg.MsgMetadataRequest{
		TokenAddress: collection.ChainParams[networkName].Address,
		TokenId:      tokenID,
		TokenType:    uint32(tokenType),
	})

	if err != nil {
		return nil, errors.Wrap(err, "failed to get metadata")
	}

	item.Metadata = &tokentypes.ItemMetadata{
		Uri:       itemMeta.Uri,
		ImageUri:  itemMeta.ImageUri,
		ImageHash: itemMeta.ImageHash,
	}

	if itemMeta.SolanaSeed != "" {
		item.ChainParams[networkName].Seed = itemMeta.SolanaSeed
	}

	k.tm.PutItem(sdk.UnwrapSDKContext(ctx), item)

	return &item, nil
}

// TODO put deposit info into MsgCreateTransferOp to make core independent from savers
func (k msgServer) getDepositInfo(ctx sdk.Context, msg *types.MsgCreateTransferOp) (*savermsg.MsgDepositResponse, error) {
	infoRequest := &savermsg.MsgTransactionInfoRequest{
		Hash:    msg.Tx,
		EventId: msg.EventId,
		Type:    k.tm.GetSaverType(ctx, msg.FromChain, msg.TokenType),
	}

	saverClient, err := saver.GetClient(msg.FromChain)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error getting saver connection", err.Error())
	}

	depositInfo, err := saverClient.GetDepositInfo(ctx.Context(), infoRequest)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "error searching deposit %s", err.Error())
	}

	return depositInfo, nil
}

func getSalt(response *savermsg.MsgDepositResponse) string {
	if response.BundleData == "" || response.BundleSalt == "" {
		return ""
	}

	return hexutil.Encode(crypto.Keccak256(hexutil.MustDecode(response.BundleSalt), hexutil.MustDecode(response.Receiver)))
}

func getBundle(response *savermsg.MsgDepositResponse) string {
	if response.BundleSalt == "" {
		return ""
	}

	return response.BundleData
}
