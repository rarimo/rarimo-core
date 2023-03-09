package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "gitlab.com/rarimo/rarimo-core/testutil/keeper"
	"gitlab.com/rarimo/rarimo-core/x/oraclemanager/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OraclemanagerKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
