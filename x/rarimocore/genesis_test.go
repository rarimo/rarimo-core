package rarimocore_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "gitlab.com/rarify-protocol/rarimo-core/testutil/keeper"
	"gitlab.com/rarify-protocol/rarimo-core/testutil/nullify"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DepositList: []types.Deposit{
			{
				Tx: "0",
			},
			{
				Tx: "1",
			},
		},
		ConfirmationList: []types.Confirmation{
			{
				Height: "0",
			},
			{
				Height: "1",
			},
		},
		ChangeKeyECDSAList: []types.ChangeKeyECDSA{
			{
				NewKey: "0",
			},
			{
				NewKey: "1",
			},
		},
		ChangeKeyEdDSAList: []types.ChangeKeyEdDSA{
			{
				NewKey: "0",
			},
			{
				NewKey: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RarimocoreKeeper(t)
	rarimocore.InitGenesis(ctx, *k, genesisState)
	got := rarimocore.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DepositList, got.DepositList)
	require.ElementsMatch(t, genesisState.ConfirmationList, got.ConfirmationList)
	require.ElementsMatch(t, genesisState.ChangeKeyECDSAList, got.ChangeKeyECDSAList)
	require.ElementsMatch(t, genesisState.ChangeKeyEdDSAList, got.ChangeKeyEdDSAList)
	// this line is used by starport scaffolding # genesis/test/assert
}
