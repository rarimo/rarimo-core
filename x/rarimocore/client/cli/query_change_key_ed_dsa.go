package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func CmdListChangeKeyEdDSA() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-change-key-ed-dsa",
		Short: "list all changeKeyEdDSA",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllChangeKeyEdDSARequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ChangeKeyEdDSAAll(context.Background(), params)
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

func CmdShowChangeKeyEdDSA() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-change-key-ed-dsa [new-key]",
		Short: "shows a changeKeyEdDSA",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argNewKey := args[0]

			params := &types.QueryGetChangeKeyEdDSARequest{
				NewKey: argNewKey,
			}

			res, err := queryClient.ChangeKeyEdDSA(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
