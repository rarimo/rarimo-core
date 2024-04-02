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
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func cmdQueryTree() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tree <limit> <offset>",
		Short: "Fetch the CSCA-based Merkle tree with pagination",
		Long: `Get all the nodes of CSCA Merkle tree. You might need to query params beforehand
to acquire the Merkle root. Having a root, you can build the tree from the list,
where each node has links to children.

Paginate over a list with <offset> parameter, adding <limit> to it on each new
request. Example: tree 30 0; tree 30 30; tree 30 60; ...`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			// argument count is ensured by cobra.ExactArgs(2)
			limit, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("parse limit: %w", err)
			}
			offset, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("parse offset: %w", err)
			}

			cliCtx := client.GetClientContextFromCmd(cmd)
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

	res, err := cli.Tree(context.Background(), &types.QueryTreeRequest{
		Pagination: &query.PageRequest{
			Limit:      limit,
			Offset:     offset,
			CountTotal: true,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("query tree: %w", err)
	}

	return res, nil
}
