package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gitlab.com/rarimo/rarimo-core/x/identity/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) SetIdentityContractAddress(goCtx context.Context, msg *types.MsgSetIdentityContractAddress) (*types.MsgSetIdentityContractAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)

	if decoded, err := hexutil.Decode(params.IdentityContractAddress); err != nil || len(decoded) == 0 {
		params.IdentityContractAddress = msg.Address
		k.SetParams(ctx, params)
	}

	return &types.MsgSetIdentityContractAddressResponse{}, nil
}
