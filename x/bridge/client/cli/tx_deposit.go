package cli

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"gitlab.com/rarimo/rarimo-core/x/bridge/types"
	tokenmanagertypes "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func CmdDepositNative() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit-native [seed] [receiver] [amount] [chain] [address] [bundleData] [bundleSalt]",
		Short: "Deposit native token",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			seed := args[0]
			receiver := args[1]

			amount, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}

			chain := args[3]
			address := args[4]
			bundleData := args[5]
			bundleSalt := args[6]

			queryClient := tokenmanagertypes.NewQueryClient(clientCtx)

			to, err := queryClient.OnChainItem(context.Background(), &tokenmanagertypes.QueryGetOnChainItemRequest{
				Chain:   chain,
				Address: address,
			})
			if err != nil {
				return err
			}

			msg := types.NewMsgDepositNative(
				clientCtx.GetFromAddress().String(),
				seed,
				receiver,
				&amount,
				bundleData,
				bundleSalt,
				to.Item.Index,
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
