package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/rarimo/rarimo-core/x/rarimocore/types"
	"github.com/spf13/cobra"
)

func CmdListConfirmation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-confirmation",
		Short: "list all confirmation",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllConfirmationRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ConfirmationAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowConfirmation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-confirmation [root]",
		Short: "shows a confirmation",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argRoot := args[0]

			params := &types.QueryGetConfirmationRequest{
				Root: argRoot,
			}

			res, err := queryClient.Confirmation(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
