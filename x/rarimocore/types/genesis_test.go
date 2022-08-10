package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated deposit",
			genState: &types.GenesisState{
				DepositList: []types.Deposit{
					{
						Tx: "0",
					},
					{
						Tx: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated confirmation",
			genState: &types.GenesisState{
				ConfirmationList: []types.Confirmation{
					{
						Height: "0",
					},
					{
						Height: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated changeKeyECDSA",
			genState: &types.GenesisState{
				ChangeKeyECDSAList: []types.ChangeKeyECDSA{
					{
						NewKey: "0",
					},
					{
						NewKey: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
