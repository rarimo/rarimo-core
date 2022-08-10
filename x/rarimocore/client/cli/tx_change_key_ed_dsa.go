package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func CmdCreateChangeKeyEdDSA() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-change-key-ed-dsa [new-key] [signature]",
		Short: "Create a new changeKeyEdDSA",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexNewKey := args[0]

			// Get value arguments
			argSignature := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateChangeKeyEdDSA(
				clientCtx.GetFromAddress().String(),
				indexNewKey,
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

func CmdUpdateChangeKeyEdDSA() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-change-key-ed-dsa [new-key] [signature]",
		Short: "Update a changeKeyEdDSA",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexNewKey := args[0]

			// Get value arguments
			argSignature := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateChangeKeyEdDSA(
				clientCtx.GetFromAddress().String(),
				indexNewKey,
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

func CmdDeleteChangeKeyEdDSA() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-change-key-ed-dsa [new-key]",
		Short: "Delete a changeKeyEdDSA",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexNewKey := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteChangeKeyEdDSA(
				clientCtx.GetFromAddress().String(),
				indexNewKey,
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
