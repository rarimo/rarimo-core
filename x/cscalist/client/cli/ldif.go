package cli

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/iden3/go-iden3-crypto/keccak256"
	"github.com/rarimo/ldif-sdk/ldif"
	"github.com/rarimo/ldif-sdk/mt"
	"github.com/rarimo/rarimo-core/x/cscalist/types"
	"github.com/spf13/cobra"
)

const outputFormatFlag = "output-format"

func cmdParseLDIF() *cobra.Command {
	const (
		rawFormat      = "raw"
		hashFormat     = "hash"
		proposalFormat = "proposal"
		rootFormat     = "root"
		defaultFormat  = proposalFormat
	)
	outputFormatValues := strings.Join([]string{rawFormat, hashFormat, proposalFormat, rootFormat}, ", ")

	cmd := &cobra.Command{
		Use:   "parse-ldif <file> [output-file]",
		Short: "Parse CSCA certificates from LDIF file",
		Long: `Provide a LDIF file to parse CSCA certificates from it. The list can be manually
retrieved from https://pkddownloadsg.icao.int/. By providing --output-format you
can get a list of hashes for a Merkle tree leaves or a full tree, by default the
keys themselves are printed.

If output-file is provided, the output is written to the specified file instead
of stdout. This is useful for raw and hash formats to perform manual comparison
on need.

Use cases:
- On CSCA list update, submit a proposal with the new list (--output-format=hash)
- Compare proposed and actual downloaded list on voting (--output-format=root|raw)
- Manually retrieve the tree and compare to the current list (can be done
automatically, see ldif-tree-diff)`,
		SuggestionsMinimumDistance: 2,
		Args:                       cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			data, err := ldif.FromFile(args[0])
			if err != nil {
				return fmt.Errorf("parse LDIF from file: %w", err)
			}

			pubKeys, err := data.RawPubKeys()
			if err != nil {
				return fmt.Errorf("extract raw public keys: %w", err)
			}
			if len(pubKeys) == 0 {
				return errors.New("no public keys found in file")
			}

			dst := os.Stdout
			if len(args) == 2 {
				dst, err = os.OpenFile(args[1], os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
				if err != nil {
					return fmt.Errorf("open destination file: %w", err)
				}
				defer func() { _ = dst.Close() }()
			}

			outputFormat, _ := cmd.Flags().GetString(outputFormatFlag)
			switch outputFormat {
			case rawFormat:
				for _, pk := range pubKeys {
					if _, err = fmt.Fprintln(dst, hexutil.Encode(pk)); err != nil {
						return fmt.Errorf("write to destination: %w", err)
					}
				}

			case hashFormat:
				return printHashes(pubKeys, dst, false)

			case proposalFormat:
				return printHashes(pubKeys, dst, true)

			case rootFormat:
				sPubKeys := make([]string, len(pubKeys))
				for i, pk := range pubKeys {
					sPubKeys[i] = string(pk)
				}
				tree, err := mt.BuildFromRaw(sPubKeys)
				if err != nil {
					return fmt.Errorf("build tree: %w", err)
				}
				if _, err = fmt.Fprintln(dst, hexutil.Encode(tree.Root())); err != nil {
					return fmt.Errorf("write to destination: %w", err)
				}

			default:
				return fmt.Errorf("invalid output format: %s. Valid values are: %s", outputFormat, outputFormatValues)
			}

			return nil
		},
	}

	cmd.Flags().String(outputFormatFlag, defaultFormat, fmt.Sprintf("Output format (%s)", outputFormatValues))
	return cmd
}

func cmdLDIFTreeDiff() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ldif-tree-diff <file>",
		Short: "Build one tree from LDIF file, query tree root from chain and compare roots",
		Long: `Provide a LDIF file to parse. The list can be manually retrieved from
https://pkddownloadsg.icao.int/. Use parse-ldif and tree subcommands to find out
the exact differences of data.

Use cases:
- Automatically validate proposals: ensure that stored and proposed list differ
- Periodically check for list updates and compare to the tree`,
		SuggestionsMinimumDistance: 2,
		Args:                       cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
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

			if built != stored {
				fmt.Printf("Trees differ: built_root=%s stored_root=%s\n", built, stored)
				return nil
			}

			fmt.Printf("Tree are same: root=%s\n", built)
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func printHashes(pubKeys [][]byte, dst *os.File, forProposal bool) error {
	for i, pk := range pubKeys {
		hash := keccak256.Hash(pk)

		out := hexutil.Encode(hash)
		if forProposal {
			var comma string
			if i < len(pubKeys)-1 {
				comma = ","
			}
			out = fmt.Sprintf("%q%s", out, comma)
		}

		if _, err := fmt.Fprintln(dst, out); err != nil {
			return fmt.Errorf("write to destination: %w", err)
		}
	}

	return nil
}

func getBuiltAndStoredRoots(data ldif.LDIF, cli types.QueryClient) (string, string, error) {
	pubKeys, err := data.RawPubKeys()
	if err != nil {
		return "", "", fmt.Errorf("extract raw public keys: %w", err)
	}

	sPubKeys := make([]string, len(pubKeys))
	for i, pk := range pubKeys {
		sPubKeys[i] = string(pk)
	}

	tree, err := mt.BuildFromRaw(sPubKeys)
	if err != nil {
		return "", "", fmt.Errorf("build Merkle tree: %w", err)
	}

	root, err := getRootNode(cli)
	if err != nil {
		return "", "", fmt.Errorf("get root node: %w", err)
	}

	return hexutil.Encode(tree.Root()), root.Hash, nil
}

func getRootNode(cli types.QueryClient) (*types.Node, error) {
	params, err := cli.Params(context.Background(), &types.QueryParamsRequest{})
	if err != nil {
		return nil, fmt.Errorf("query params: %w", err)
	}

	resp, err := cli.Tree(context.Background(), &types.QueryTreeRequest{
		Pagination: &query.PageRequest{
			Key:   []byte(params.Params.RootKey),
			Limit: 1,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("query tree: %w", err)
	}

	if len(resp.Tree) == 0 {
		return nil, errors.New("root node not found")
	}

	return &resp.Tree[0], nil
}
