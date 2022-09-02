package cli

import (
	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func CmdCreateDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-deposit [tx] [event-id] [from-chain] [to-chain] [receiver-hex] [token-address-hex] [token-id-hex]",
		Short: "Create a new deposit",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexTx := args[0]
			eventId := args[1]

			// Get value arguments
			argFromChain := args[2]
			argToChain := args[3]
			argReceiver, err := hex.DecodeString(args[4])
			if err != nil {
				return err
			}
			argTokenAddress, err := hex.DecodeString(args[5])
			if err != nil {
				return err
			}
			argTokenId, err := hex.DecodeString(args[6])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateDeposit(
				clientCtx.GetFromAddress().String(),
				indexTx,
				eventId,
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
