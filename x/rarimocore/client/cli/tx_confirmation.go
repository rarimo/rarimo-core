package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
	"strings"
)

func CmdCreateConfirmation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-confirmation [height] [root] [hashes] [signature]",
		Short: "Create a new confirmation",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexHeight := args[0]

			// Get value arguments
			argRoot := args[1]
			argHashes := strings.Split(args[2], listSeparator)
			argSignature := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateConfirmation(
				clientCtx.GetFromAddress().String(),
				indexHeight,
				argRoot,
				argHashes,
				argSignature,
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

func CmdUpdateConfirmation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-confirmation [height] [root] [hashes] [signature]",
		Short: "Update a confirmation",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexHeight := args[0]

			// Get value arguments
			argRoot := args[1]
			argHashes := strings.Split(args[2], listSeparator)
			argSignature := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateConfirmation(
				clientCtx.GetFromAddress().String(),
				indexHeight,
				argRoot,
				argHashes,
				argSignature,
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

func CmdDeleteConfirmation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-confirmation [height]",
		Short: "Delete a confirmation",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexHeight := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteConfirmation(
				clientCtx.GetFromAddress().String(),
				indexHeight,
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
