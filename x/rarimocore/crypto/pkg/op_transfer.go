package pkg

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gogo/protobuf/proto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/bundle"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/content"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/operation"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto/origin"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
	tokentypes "gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/types"
)

func GetTransfer(operation types.Operation) (*types.Transfer, error) {
	if operation.OperationType == types.OpType_TRANSFER {
		transfer := new(types.Transfer)
		return transfer, proto.Unmarshal(operation.Details.Value, transfer)
	}
	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "invalid operation type")
}

func GetTransferContent(info *tokentypes.Info, item *tokentypes.Item, params *tokentypes.ChainParams, transfer *types.Transfer) (*content.TransferContent, error) {
	builder := operation.NewTransferOperationBuilder()

	switch item.TokenType {
	case tokentypes.Type_NEAR_FT:
		builder.
			SetAddress(info.Chains[transfer.ToChain].TokenAddress).
			SetAmount(transfer.Amount)
	case tokentypes.Type_NEAR_NFT:
		builder.
			SetAddress(info.Chains[transfer.ToChain].TokenAddress).
			SetId(info.Chains[transfer.ToChain].TokenId).
			SetAmount(transfer.Amount).
			SetName(item.Name).
			SetImageURI(item.ImageUri).
			SetImageHash(item.ImageHash)
	case tokentypes.Type_NATIVE:
		builder.SetAmount(transfer.Amount)
	case tokentypes.Type_ERC20:
		builder.
			SetAddress(info.Chains[transfer.ToChain].TokenAddress).
			SetAmount(transfer.Amount)
	case tokentypes.Type_ERC721:
		builder.
			SetAddress(info.Chains[transfer.ToChain].TokenAddress).
			SetId(info.Chains[transfer.ToChain].TokenId).
			SetAmount(transfer.Amount).
			SetURI(item.Uri)
	case tokentypes.Type_ERC1155:
		builder.
			SetAddress(info.Chains[transfer.ToChain].TokenAddress).
			SetId(info.Chains[transfer.ToChain].TokenId).
			SetAmount(transfer.Amount).
			SetURI(item.Uri)
	case tokentypes.Type_METAPLEX_FT:
		builder.
			SetAddress(info.Chains[transfer.ToChain].TokenAddress).
			SetId(info.Chains[transfer.ToChain].TokenId).
			SetAmount(transfer.Amount).
			SetName(item.Name).
			SetSymbol(item.Symbol).
			SetURI(item.Uri).
			SetDecimals(uint8(item.Decimals))
	case tokentypes.Type_METAPLEX_NFT:
		builder.
			SetAddress(info.Chains[transfer.ToChain].TokenAddress).
			SetId(info.Chains[transfer.ToChain].TokenId).
			SetName(item.Name).
			SetSymbol(item.Symbol).
			SetURI(item.Uri)
	default:
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, "unsupported token type")
	}

	return &content.TransferContent{
		Origin:         origin.NewDefaultOriginBuilder().SetTxHash(transfer.Tx).SetOpId(transfer.EventId).SetCurrentNetwork(transfer.FromChain).Build().GetOrigin(),
		TargetNetwork:  transfer.ToChain,
		Receiver:       hexutil.MustDecode(transfer.Receiver),
		Data:           builder.Build().GetContent(),
		TargetContract: hexutil.MustDecode(params.Contract),
		Bundle:         bundle.NewDefaultBundleBuilder().SetBundle(transfer.GetBundleData()).SetSalt(transfer.BundleSalt).Build().GetBundle(),
	}, nil
}
