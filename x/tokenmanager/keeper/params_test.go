package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "gitlab.com/rarify-protocol/rarimo-core/testutil/keeper"
	"gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TokenmanagerKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
