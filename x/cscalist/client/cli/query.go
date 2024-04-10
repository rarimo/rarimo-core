package cli

import (
	"context"
	"fmt"
	"strconv"

	// "strings"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/rarimo/rarimo-core/x/cscalist/types"
)

const maxPageLimit = 100000

// GetQueryCmd returns the CLI query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group cscalist queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		cmdQueryParams(),
		cmdQueryTree(),
		cmdParseLDIF(),
		cmdLDIFTreeDiff(),
	)
	// this line is used by starport scaffolding # 1

	return cmd
}

func cmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "shows the parameters of the module",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return fmt.Errorf("query params: %w", err)
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func cmdQueryTree() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tree [limit] [offset]",
		Short: "Fetch the CSCA-based Merkle",
		Long: `Get nodes of CSCA Merkle tree. You might need to fetch params beforehand to
acquire the Merkle root. Having a root, you can build the tree from the list,
where each node has links to children.

If neither limit nor offset is provided, the entire tree is fetched at once.

You can paginate over a list with [offset] parameter, adding [limit] to it on
each new request. Example: tree 30 0; tree 30 30; tree 30 60; ...`,
		Args: cobra.MaximumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := client.GetClientContextFromCmd(cmd)

			if len(args) == 0 {
				res, err := queryTree(cliCtx, maxPageLimit, 0)
				if err != nil {
					return fmt.Errorf("query entire tree: %w", err)
				}

				res.Pagination = nil
				return cliCtx.PrintProto(res)
			}

			limit, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("parse limit: %w", err)
			}

			var offset uint64
			if len(args) == 2 {
				offset, err = strconv.ParseUint(args[1], 10, 64)
				if err != nil {
					return fmt.Errorf("parse offset: %w", err)
				}
			}

			res, err := queryTree(cliCtx, limit, offset)
			if err != nil {
				return fmt.Errorf("query tree: %w", err)
			}

			return cliCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func queryTree(ctx client.Context, limit, offset uint64) (*types.QueryTreeResponse, error) {
	cli := types.NewQueryClient(ctx)

	return cli.Tree(context.Background(), &types.QueryTreeRequest{
		Pagination: &query.PageRequest{
			Limit:      limit,
			Offset:     offset,
			CountTotal: true,
		},
	})
}
