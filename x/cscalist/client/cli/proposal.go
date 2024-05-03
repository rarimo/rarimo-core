package cli

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/iden3/go-iden3-crypto/keccak256"
	"github.com/rarimo/ldif-sdk/ldif"
	"github.com/rarimo/rarimo-core/x/cscalist/types"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func cmdPrepareProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prepare-proposal <file> [output-file]",
		Short: "Prepare content for EditCSCAListProposal from LDIF file",
		Long: `Provide a LDIF file to parse CSCA certificates from it. The list can be manually
retrieved from https://pkddownloadsg.icao.int/. The command builds one tree from
LDIF, fetches another one from chain and selects leaves to add or remove.

If output-file is provided, the output is written to the specified file instead
of stdout. This is useful to copy-paste from a file.

Use cases:
- Copy-paste the output to your EditCSCAListProposal draft`,
		SuggestionsMinimumDistance: 2,
		Args:                       cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			dst := os.Stdout
			if len(args) == 2 {
				dst, err = os.OpenFile(args[1], os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
				if err != nil {
					return fmt.Errorf("open destination file: %w", err)
				}
				defer func() { _ = dst.Close() }()
			}

			data, err := ldif.FromFile(args[0])
			if err != nil {
				return fmt.Errorf("parse LDIF from file: %w", err)
			}

			cliCtx := client.GetClientContextFromCmd(cmd)
			cli := types.NewQueryClient(cliCtx)

			built, stored, err := getBuiltAndStoredRoots(data, cli)
			if err != nil {
				return fmt.Errorf("get built and stored roots: %w", err)
			}

			if built == stored {
				_, err = fmt.Fprintf(dst, "Tree are same: root=%s. Nothing to add or remove.\n", built)
				if err != nil {
					return fmt.Errorf("write to destination: %w", err)
				}
				return nil
			}

			oldLeaves, err := fetchHashLeavesFromCosmos(cmd.Context(), cli)
			if err != nil {
				return fmt.Errorf("fetch hash leaves from cosmos: %w", err)
			}

			pubKeys, _ := data.RawPubKeys() // error was handled in getBuiltAndStoredRoots
			newLeaves := make([]string, len(pubKeys))
			for i, pk := range pubKeys {
				newLeaves[i] = hexutil.Encode(keccak256.Hash(pk))
			}

			toAdd, toRemove := selectProposalHashes(oldLeaves, newLeaves)

			_, err = fmt.Fprintf(dst, "=== To add ===\n%s=== To remove ===\n%s", formatHashes(toAdd), formatHashes(toRemove))
			if err != nil {
				return fmt.Errorf("write to destination: %w", err)
			}
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// optimized copy of github.com/rarimo/ldif-sdk/utils.FetchHashLeavesFromCosmos
func fetchHashLeavesFromCosmos(ctx context.Context, client types.QueryClient) ([]string, error) {
	var (
		limit  = uint64(100)
		offset = uint64(0)
		leaves []string
	)

	for {
		resp, err := client.Tree(ctx, &types.QueryTreeRequest{
			Pagination: &query.PageRequest{
				CountTotal: true,
				Limit:      limit,
				Offset:     offset,
			},
		}, grpc.EmptyCallOption{})
		if err != nil {
			return nil, fmt.Errorf("query paginated tree: %w", err)
		}

		if cap(leaves) == 0 {
			leaves = make([]string, 0, resp.Pagination.Total)
		}

		for _, node := range resp.Tree {
			leaves = append(leaves, node.Key)
		}

		if uint64(len(leaves)) == resp.Pagination.Total {
			break
		}

		offset += limit
	}

	return leaves, nil
}

func selectProposalHashes(oldLeaves, newLeaves []string) (toAdd []string, toRemove []string) {
	toAdd = make([]string, 0, len(newLeaves))
	toRemove = make([]string, 0, len(oldLeaves))

	for _, leaf := range newLeaves {
		if !existsInSlice(oldLeaves, leaf) {
			toAdd = append(toAdd, leaf)
		}
	}

	for _, leaf := range oldLeaves {
		if !existsInSlice(newLeaves, leaf) {
			toRemove = append(toRemove, leaf)
		}
	}

	return
}

func formatHashes(hashes []string) string {
	var sb strings.Builder

	for i, h := range hashes {
		var comma string
		if i < len(hashes)-1 {
			comma = ","
		}
		sb.WriteString(fmt.Sprintf("%q%s\n", h, comma))
	}

	return sb.String()
}

func existsInSlice(slice []string, elem string) bool {
	for _, e := range slice {
		if e == elem {
			return true
		}
	}
	return false
}
