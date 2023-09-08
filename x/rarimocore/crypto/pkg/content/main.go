package content

import (
	"context"

	"github.com/pkg/errors"
	merkle "gitlab.com/rarimo/go-merkle"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto/pkg"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
	token "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
	"google.golang.org/grpc"
)

var (
	// ErrOpShouldBeApproved appears when someone tries to add operation that has been already signed
	ErrOpShouldBeApproved = errors.New("operation should be approved")
	ErrUnsupportedContent = errors.New("unsupported content")
)

func GetContents(client *grpc.ClientConn, operations ...*types.Operation) ([]merkle.Content, error) {
	contents := make([]merkle.Content, 0, len(operations))

	for _, op := range operations {
		if op.Status != types.OpStatus_APPROVED {
			return nil, ErrOpShouldBeApproved
		}

		switch op.OperationType {
		case types.OpType_TRANSFER:
			content, err := GetTransferContent(client, op)
			if err != nil {
				return nil, err
			}

			if content != nil {
				contents = append(contents, content)
			}
		case types.OpType_CHANGE_PARTIES:
			return nil, ErrUnsupportedContent
		case types.OpType_FEE_TOKEN_MANAGEMENT:
			content, err := GetFeeManagementContent(client, op)
			if err != nil {
				return nil, err
			}

			if content != nil {
				contents = append(contents, content)
			}
		case types.OpType_CONTRACT_UPGRADE:
			content, err := GetContractUpgradeContent(client, op)
			if err != nil {
				return nil, err
			}

			if content != nil {
				contents = append(contents, content)
			}
		case types.OpType_IDENTITY_DEFAULT_TRANSFER:
			content, err := GetIdentityDefaultTransferContent(op)
			if err != nil {
				return nil, err
			}

			if content != nil {
				contents = append(contents, content)
			}
		case types.OpType_IDENTITY_AGGREGATED_TRANSFER:
			content, err := GetIdentityAggregatedTransferContent(op)
			if err != nil {
				return nil, err
			}

			if content != nil {
				contents = append(contents, content)
			}
		default:
			return nil, ErrUnsupportedContent
		}
	}

	return contents, nil
}

func GetTransferContent(client *grpc.ClientConn, op *types.Operation) (merkle.Content, error) {
	transfer, err := pkg.GetTransfer(*op)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing operation details")
	}

	collectionDataResp, err := token.NewQueryClient(client).CollectionData(context.TODO(), &token.QueryGetCollectionDataRequest{Chain: transfer.To.Chain, Address: transfer.To.Address})
	if err != nil {
		return nil, errors.Wrap(err, "error getting collection data entry")
	}

	collectionResp, err := token.NewQueryClient(client).Collection(context.TODO(), &token.QueryGetCollectionRequest{Index: collectionDataResp.Data.Collection})
	if err != nil {
		return nil, errors.Wrap(err, "error getting collection data entry")
	}

	onChainItemResp, err := token.NewQueryClient(client).OnChainItem(context.TODO(), &token.QueryGetOnChainItemRequest{Chain: transfer.To.Chain, Address: transfer.To.Address, TokenID: transfer.To.TokenID})
	if err != nil {
		return nil, errors.Wrap(err, "error getting on chain item entry")
	}

	itemResp, err := token.NewQueryClient(client).Item(context.TODO(), &token.QueryGetItemRequest{Index: onChainItemResp.Item.Item})
	if err != nil {
		return nil, errors.Wrap(err, "error getting item entry")
	}

	networkResp, err := token.NewQueryClient(client).NetworkParams(context.TODO(), &token.QueryNetworkParamsRequest{Name: transfer.To.Chain})
	if err != nil {
		return nil, errors.Wrap(err, "error getting network param entry")
	}

	bridgeparams := networkResp.Params.GetBridgeParams()
	if err != nil {
		return nil, errors.New("bridge params not found")
	}

	content, err := pkg.GetTransferContent(collectionResp.Collection, collectionDataResp.Data, itemResp.Item, bridgeparams, transfer)
	return content, errors.Wrap(err, "error creating content")
}

func GetFeeManagementContent(client *grpc.ClientConn, op *types.Operation) (merkle.Content, error) {
	manage, err := pkg.GetFeeTokenManagement(*op)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing operation details")
	}

	networkResp, err := token.NewQueryClient(client).NetworkParams(context.TODO(), &token.QueryNetworkParamsRequest{Name: manage.Chain})
	if err != nil {
		return nil, errors.Wrap(err, "error getting network param entry")
	}

	feeparams := networkResp.Params.GetFeeParams()
	if err != nil {
		return nil, errors.New("bridge params not found")
	}

	content, err := pkg.GetFeeTokenManagementContent(feeparams, manage)
	return content, errors.Wrap(err, "error creating content")
}

func GetContractUpgradeContent(client *grpc.ClientConn, op *types.Operation) (merkle.Content, error) {
	upgrade, err := pkg.GetContractUpgrade(*op)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing operation details")
	}

	networkResp, err := token.NewQueryClient(client).NetworkParams(context.TODO(), &token.QueryNetworkParamsRequest{Name: upgrade.Chain})
	if err != nil {
		return nil, errors.Wrap(err, "error getting network param entry")
	}

	content, err := pkg.GetContractUpgradeContent(networkResp.Params, upgrade)
	return content, errors.Wrap(err, "error creating content")
}

func GetIdentityDefaultTransferContent(op *types.Operation) (merkle.Content, error) {
	transfer, err := pkg.GetIdentityDefaultTransfer(*op)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing operation details")
	}

	content, err := pkg.GetIdentityDefaultTransferContent(transfer)
	return content, errors.Wrap(err, "error creating content")
}

func GetIdentityAggregatedTransferContent(op *types.Operation) (merkle.Content, error) {
	transfer, err := pkg.GetIdentityAggregatedTransfer(*op)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing operation details")
	}

	content, err := pkg.GetIdentityAggregatedTransferContent(transfer)
	return content, errors.Wrap(err, "error creating content")
}
