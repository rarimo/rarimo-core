package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/types"
)

func CmdCreateInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-info [index] [name] [symbol] [image] [current-chain] [current-address-hex] [current-id-hex]",
		Short: "Create a new info",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argName := args[1]
			argSymbol := args[2]
			argImage := args[3]
			argCurrentChain := args[4]
			argCurrentAddress := args[5]
			argCurrentId := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateInfo(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argName,
				argSymbol,
				argImage,
				argCurrentId,
				argCurrentAddress,
				argCurrentChain,
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

func CmdUpdateInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-info [index] [name] [symbol] [image]",
		Short: "Update a info",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argName := args[1]
			argSymbol := args[2]
			argImage := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateInfo(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argName,
				argSymbol,
				argImage,
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

func CmdDeleteInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-info [index]",
		Short: "Delete a info",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexIndex := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteInfo(
				clientCtx.GetFromAddress().String(),
				indexIndex,
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

func CmdAddChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-chain [index] [chain] [token-address-hex] [token-id-hex]",
		Short: "Add a chain to info",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argChain := args[1]
			argTokenAddress := args[2]
			argTokenId := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddChain(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argChain,
				argTokenAddress,
				argTokenId,
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
