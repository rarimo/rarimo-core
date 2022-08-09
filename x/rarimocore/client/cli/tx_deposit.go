package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func CmdCreateDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-deposit [tx] [from-chain] [to-chain] [receiver] [token-address] [token-id]",
		Short: "Create a new deposit",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexTx := args[0]

			// Get value arguments
			argFromChain := args[1]
			argToChain := args[2]
			argReceiver := args[3]
			argTokenAddress := args[4]
			argTokenId := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateDeposit(
				clientCtx.GetFromAddress().String(),
				indexTx,
				argFromChain,
				argToChain,
				argReceiver,
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

func CmdUpdateDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-deposit [tx] [from-chain] [to-chain] [receiver] [token-address] [token-id]",
		Short: "Update a deposit",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexTx := args[0]

			// Get value arguments
			argFromChain := args[1]
			argToChain := args[2]
			argReceiver := args[3]
			argTokenAddress := args[4]
			argTokenId := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateDeposit(
				clientCtx.GetFromAddress().String(),
				indexTx,
				argFromChain,
				argToChain,
				argReceiver,
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

func CmdDeleteDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-deposit [tx]",
		Short: "Delete a deposit",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexTx := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteDeposit(
				clientCtx.GetFromAddress().String(),
				indexTx,
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
