package cli

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/rarimo/rarimo-core/x/rarimocore/types"
	"github.com/spf13/cobra"
)

func CmdCreateConfirmation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-confirmation [root] [indexes] [sig-ecdsa]",
		Short: "Create a new confirmation",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexRoot := args[0]

			// Get value arguments
			argIndexes := strings.Split(args[1], listSeparator)
			argSigECDSA := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateConfirmation(
				clientCtx.GetFromAddress().String(),
				indexRoot,
				argIndexes,
				argSigECDSA,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
