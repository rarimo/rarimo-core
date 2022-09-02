package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func CmdCreateDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-deposit [tx] [event-id] [from-chain] [to-chain] [receiver-hex] [token-address-hex] [token-id-hex] [token-type]",
		Short: "Create a new deposit",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexTx := args[0]
			eventId := args[1]

			// Get value arguments
			argFromChain := args[2]
			argToChain := args[3]
			argReceiver := args[4]
			argTokenAddress := args[5]
			argTokenId := args[6]
			argTokenType, err := strconv.Atoi(args[7])
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
