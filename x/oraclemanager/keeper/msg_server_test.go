package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "gitlab.com/rarimo/rarimo-core/x/oraclemanager/types"
    "gitlab.com/rarimo/rarimo-core/x/oraclemanager/keeper"
    keepertest "gitlab.com/rarimo/rarimo-core/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.OraclemanagerKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
