package pkg

import (
	"fmt"

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

func GetTransferContent(item *tokentypes.Item, collection *tokentypes.CollectionInfo, transfer *types.Transfer) (*operation.TransferContent, error) {
	builder := data.NewTransferDataBuilder()

	destinationChainParams, ok := collection.ChainParams[transfer.ToChain]
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("chain [%s] not found for collection [%s]", transfer.ToChain, collection.Index))
	}

	destinationItemParams, ok := item.ChainParams[transfer.ToChain]
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("chain [%s] not found for item [%s]", transfer.ToChain, item.Index))
	}

	switch collection.TokenType {
	case tokentypes.Type_NEAR_FT:
		builder.
			SetAddress(destinationChainParams.Address).
			SetAmount(transfer.Amount)
	case tokentypes.Type_NEAR_NFT:
		builder.
			SetAddress(destinationChainParams.Address).
			SetId(destinationItemParams.TokenID).
			SetName(collection.Metadata.Name).
			SetImageURI(item.Metadata.ImageUri).
			SetImageHash(item.Metadata.ImageHash)
	case tokentypes.Type_NATIVE:
		builder.SetAmount(transfer.Amount)
	case tokentypes.Type_ERC20:
		builder.
			SetAddress(destinationChainParams.Address).
			SetAmount(transfer.Amount)
	case tokentypes.Type_ERC721:
		builder.
			SetAddress(destinationChainParams.Address).
			SetId(destinationItemParams.TokenID).
			SetURI(item.Metadata.Uri)
	case tokentypes.Type_ERC1155:
		builder.
			SetAddress(destinationChainParams.Address).
			SetId(destinationItemParams.TokenID).
			SetAmount(transfer.Amount).
			SetURI(item.Metadata.Uri)
	case tokentypes.Type_METAPLEX_FT:
		builder.
			SetAddress(destinationChainParams.Address).
			SetAmount(transfer.Amount).
			SetName(collection.Metadata.Name).
			SetSymbol(collection.Metadata.Symbol).
			SetURI(item.Metadata.Uri).
			SetDecimals(uint8(destinationChainParams.Decimals))
	case tokentypes.Type_METAPLEX_NFT:
		builder.
			SetAddress(destinationChainParams.Address).
			SetId(destinationItemParams.TokenID).
			SetName(collection.Metadata.Name).
			SetSymbol(collection.Metadata.Symbol).
			SetURI(item.Metadata.Uri)
	default:
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "unsupported token type")
	}

	return &operation.TransferContent{
		Origin: origin.NewDefaultOriginBuilder().
			SetTxHash(transfer.Tx).
			SetOpId(transfer.EventId).
			SetCurrentNetwork(transfer.FromChain).
			Build().
			GetOrigin(),
		TargetNetwork:  transfer.ToChain,
		Receiver:       hexutil.MustDecode(transfer.Receiver),
		Data:           builder.Build().GetContent(),
		TargetContract: hexutil.MustDecode(destinationChainParams.Address),
		Bundle: bundle.NewDefaultBundleBuilder().
			SetBundle(transfer.BundleData).
			SetSalt(transfer.BundleSalt).
			Build().
			GetBundle(),
	}, nil
}
