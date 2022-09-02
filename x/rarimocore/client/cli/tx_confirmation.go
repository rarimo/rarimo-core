package cli

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func CmdCreateConfirmation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-confirmation [height] [root] [hashes] [sig-ecdsa]",
		Short: "Create a new confirmation",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexHeight := args[0]

			// Get value arguments
			argRoot := args[1]
			argHashes := strings.Split(args[2], listSeparator)
			argSigECDSA := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateConfirmation(
				clientCtx.GetFromAddress().String(),
				indexHeight,
				argRoot,
				argHashes,
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
