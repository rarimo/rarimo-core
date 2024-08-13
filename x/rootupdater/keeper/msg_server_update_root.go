package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/rootupdater/types"
)

func (m msgServer) UpdateRoot(goctx context.Context, req *types.MsgUpdateRoot) (*types.MsgUpdateRootResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	params := m.GetParams(ctx)

	params.Root = req.Root

	m.SetParams(ctx, params)

	return &types.MsgUpdateRootResponse{}, nil
}
