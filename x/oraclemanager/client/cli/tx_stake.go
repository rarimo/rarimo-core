package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"gitlab.com/rarimo/rarimo-core/x/oraclemanager/types"
)

func CmdStake() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stake [chain] [amount] [oracle-account]",
		Short: "Stake tokens for oracle",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			chain := args[0]

			// Get value arguments
			amount := args[1]

			oracleAccount := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress().String()
			if oracleAccount == from {
				from = ""
			}

			msg := &types.MsgStake{
				Index: &types.OracleIndex{
					Chain:   chain,
					Account: oracleAccount,
				},
				Amount: amount,
				From:   from,
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
