package cli

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/rarimo/ldif-sdk/ldif"
	"github.com/rarimo/ldif-sdk/mt"
	"github.com/rarimo/ldif-sdk/utils"
	"github.com/rarimo/rarimo-core/x/cscalist/types"
	"github.com/spf13/cobra"
)

const outputFormatFlag = "output-format"

func cmdParseLDIF() *cobra.Command {
	const (
		rawFormat     = "raw"
		hashFormat    = "hash"
		rootFormat    = "root"
		defaultFormat = rawFormat
	)
	outputFormatValues := fmt.Sprintf("%s, %s, %s", rawFormat, hashFormat, rootFormat)

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
					if _, err = fmt.Fprintln(dst, hex.EncodeToString(pk)); err != nil {
						return fmt.Errorf("write to destination: %w", err)
					}
				}

			case hashFormat:
				for _, pk := range pubKeys {
					hash, err := utils.PoseidonHashBig(pk)
					if err != nil {
						return fmt.Errorf("hash public key: %w", err)
					}
					if _, err = fmt.Fprintln(dst, hash.Text(16)); err != nil {
						return fmt.Errorf("write to destination: %w", err)
					}
				}

			case rootFormat:
				sPubKeys := make([]string, len(pubKeys))
				for i, pk := range pubKeys {
					sPubKeys[i] = string(pk)
				}
				tree, err := mt.BuildFromRaw(sPubKeys)
				if err != nil {
					return fmt.Errorf("build tree: %w", err)
				}
				if _, err = fmt.Fprintln(dst, tree.Root()); err != nil {
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
		Short: "Build one tree from LDIF file, query another tree from chain and compare roots",
		Long: `Provide a LDIF file to parse. The list can be manually retrieved from
https://pkddownloadsg.icao.int/. When the file is parsed, one tree is built,
another one is queried. Roots of 2 trees are compared. Use parse-ldif and tree
subcommands to find out the exact differences of data.

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

			raw, err := json.Marshal(data.ToPem())
			if err != nil {
				return fmt.Errorf("marshal pem string array: %w", err)
			}

			tree, err := mt.BuildTree(raw)
			if err != nil {
				return fmt.Errorf("build Merkle tree: %w", err)
			}

			cliCtx := client.GetClientContextFromCmd(cmd)
			cli := types.NewQueryClient(cliCtx)

			params, err := cli.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return fmt.Errorf("query params: %w", err)
			}

			// TODO @violog use a different tree implementation, or adjust this tree to a treap
			if tree.Root() != params.Params.RootKey {
				fmt.Printf("Trees differ: built_root=%s stored_root=%s\n", tree.Root(), params.Params.RootKey)
				return nil
			}

			fmt.Printf("Tree are same: root=%s", tree.Root())
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
