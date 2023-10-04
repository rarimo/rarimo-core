package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/rarimo/rarimo-core/x/oraclemanager/types"
	"github.com/spf13/cobra"
)

func CmdUnjail() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unjail [chain]",
		Short: "Unjail oracle",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			chain := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgUnjail{
				Index: &types.OracleIndex{
					Chain:   chain,
					Account: clientCtx.GetFromAddress().String(),
				},
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
