package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func CmdCreateInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-info [index] [current-address-hex] [current-id-hex] [current-chain] [token-type]",
		Short: "Create a new info",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argCurrentAddress := args[1]
			argCurrentId := args[2]
			argChain := args[3]
			argTokenType, err := strconv.Atoi(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateInfo(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argCurrentId,
				argCurrentAddress,
				argChain,
				types.Type(argTokenType),
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
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argChain := args[1]
			argTokenAddress := args[2]
			argTokenId := args[3]
			argTokenType, err := strconv.Atoi(args[4])
			if err != nil {
				return err
			}

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
				types.Type(argTokenType),
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
