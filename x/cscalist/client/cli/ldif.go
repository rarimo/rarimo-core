package cli

import (
	"context"
	"encoding/hex"
	"fmt"

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
		rawFormat  = "raw"
		hashFormat = "hash"
		rootFormat = "root"
	)
	outputFormatValues := fmt.Sprintf("%s, %s, %s", rawFormat, hashFormat, rootFormat)

	cmd := &cobra.Command{
		Use:   "parse-ldif <file>",
		Short: "Parse CSCA certificates from LDIF file",
		Long: `Provide a LDIF file to parse CSCA certificates from it. The list can be manually
retrieved from https://pkddownloadsg.icao.int/. By providing --output-format you
can get a list of hashes for a Merkle tree leaves or a full tree, by default the
keys themselves are printed.

Use cases:
- On CSCA list update, submit a proposal with the new list (--output-format=hash)
- Compare proposed and actual downloaded list on voting (--output-format=root|raw)
- Manually retrieve the tree and compare to the current list (can be done
automatically, see ldif-tree-diff)`,
		SuggestionsMinimumDistance: 2,
		Args:                       cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			certs, err := ldif.LDIFToX509(args[0])
			if err != nil {
				return fmt.Errorf("parse certificates from LDIF: %w", err)
			}

			pubKeys := make([]string, 0, len(certs))
			for _, cert := range certs {
				pubKey, err := utils.ParseCertPublicKey(cert.RawSubjectPublicKeyInfo)
				if err != nil {
					return fmt.Errorf("parse public key: %w", err)
				}
				pubKeys = append(pubKeys, hex.EncodeToString(pubKey.N.Bytes()))
			}

			outputFormat := cmd.Flag(outputFormatFlag).Value.String()
			switch outputFormat {
			case rawFormat:
				fmt.Print("Raw public keys: ")
				for _, pubKey := range pubKeys {
					fmt.Printf("%s ", pubKey)
				}
			case hashFormat:
				fmt.Print("Hashed leaves of the tree: ")
				for _, cert := range certs {
					hash, err := utils.HashCertificate(cert)
					if err != nil {
						return fmt.Errorf("hash certificate: %w", err)
					}
					fmt.Printf("%s ", hash.Text(16))
				}
			case rootFormat:
				tree, err := mt.BuildTree(pubKeys)
				if err != nil {
					return fmt.Errorf("build tree: %w", err)
				}
				fmt.Printf("Tree root: %s", tree.Root())
			default:
				return fmt.Errorf("invalid output format: %s. Valid values are: %s", outputFormat, outputFormatValues)
			}

			fmt.Println()
			return nil
		},
	}

	cmd.LocalFlags().String(outputFormatFlag, rawFormat, fmt.Sprintf("Output format (%s)", outputFormatValues))
	return cmd
}

func cmdLDIFTreeDiff() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ldif-tree-diff <file>",
		Short: "Build one tree from LDIF file, query another tree from chain and compare roots",
		Long: `Provide a LDIF file to parse. The list can be manually retrieved from
https://pkddownloadsg.icao.int/. When the file is parsed, the tree is queried
and built. Roots of 2 trees are compared. Use parse-ldif and tree subcommands to
find out the exact differences of data.

Use cases:
- Automatically validate proposals: ensure that stored and proposed list differ
- Periodically check for list updates and compare to the tree`,
		SuggestionsMinimumDistance: 2,
		Args:                       cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			certs, err := ldif.LDIFToPEM(args[0])
			if err != nil {
				return fmt.Errorf("parse certificates as PEM from LDIF: %w", err)
			}

			tree, err := mt.BuildTree(certs)
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
