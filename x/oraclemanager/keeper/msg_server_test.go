package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "gitlab.com/rarimo/rarimo-core/testutil/keeper"
	"gitlab.com/rarimo/rarimo-core/x/oraclemanager/keeper"
	"gitlab.com/rarimo/rarimo-core/x/oraclemanager/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.OraclemanagerKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
