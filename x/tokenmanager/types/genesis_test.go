package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/types"
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

				ItemList: []types.Item{
					{
						TokenAddress: "0",
						TokenId:      "0",
					},
					{
						TokenAddress: "1",
						TokenId:      "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated item",
			genState: &types.GenesisState{
				ItemList: []types.Item{
					{
						TokenAddress: "0",
						TokenId:      "0",
					},
					{
						TokenAddress: "0",
						TokenId:      "0",
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
