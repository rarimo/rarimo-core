package pkg

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gogo/protobuf/proto"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/bundle"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/data"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/operation/origin"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
	tokentypes "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func GetTransfer(operation types.Operation) (*types.Transfer, error) {
	if operation.OperationType == types.OpType_TRANSFER {
		transfer := new(types.Transfer)
		return transfer, proto.Unmarshal(operation.Details.Value, transfer)
	}
	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "invalid operation type")
}

func GetTransferContent(collectionData tokentypes.CollectionData, item tokentypes.Item, transfer *types.Transfer) (*operation.TransferContent, error) {
	builder := data.NewTransferDataBuilder()

	switch collectionData.TokenType {
	case tokentypes.Type_NEAR_FT:
		builder.
			SetAddress(collectionData.Index.Address).
			SetAmount(transfer.Amount)
	case tokentypes.Type_NEAR_NFT:
		builder.
			SetAddress(collectionData.Index.Address).
			SetId(transfer.To.TokenID).
			SetName(item.Index.Name).
			SetImageURI(item.Meta.ImageUri).
			SetImageHash(item.Meta.ImageHash)
	case tokentypes.Type_NATIVE:
		builder.SetAmount(transfer.Amount)
	case tokentypes.Type_ERC20:
		builder.
			SetAddress(collectionData.Index.Address).
			SetAmount(transfer.Amount)
	case tokentypes.Type_ERC721:
		builder.
			SetAddress(collectionData.Index.Address).
			SetId(transfer.To.TokenID).
			SetURI(item.Index.Uri)
	case tokentypes.Type_ERC1155:
		builder.
			SetAddress(collectionData.Index.Address).
			SetId(transfer.To.TokenID).
			SetAmount(transfer.Amount).
			SetURI(item.Index.Uri)
	case tokentypes.Type_METAPLEX_FT:
		builder.
			SetAddress(collectionData.Index.Address).
			SetAmount(transfer.Amount).
			SetName(item.Index.Name).
			SetSymbol(item.Index.Symbol).
			SetURI(item.Index.Uri).
			SetDecimals(uint8(collectionData.Decimals))
	case tokentypes.Type_METAPLEX_NFT:
		builder.
			SetAddress(collectionData.Index.Address).
			SetId(transfer.To.TokenID).
			SetName(item.Index.Name).
			SetSymbol(item.Index.Symbol).
			SetURI(item.Index.Uri)
	default:
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "unsupported token type")
	}

	return &operation.TransferContent{
		Origin: origin.NewDefaultOriginBuilder().
			SetTxHash(transfer.Tx).
			SetOpId(transfer.EventId).
			SetCurrentNetwork(transfer.From.Chain).
			Build().
			GetOrigin(),
		TargetNetwork:  transfer.To.Chain,
		Receiver:       hexutil.MustDecode(transfer.Receiver),
		Data:           builder.Build().GetContent(),
		TargetContract: hexutil.MustDecode(collectionData.Index.Chain),
		Bundle: bundle.NewDefaultBundleBuilder().
			SetBundle(transfer.BundleData).
			SetSalt(transfer.BundleSalt).
			Build().
			GetBundle(),
	}, nil
}
