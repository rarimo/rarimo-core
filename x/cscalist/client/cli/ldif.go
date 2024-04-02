package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func cmdParseLDIF() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "parse-ldif <file>",
		Short: "Parse CSCA certificates from LDIF file",
		Long: `Provide a LDIF file to parse CSCA certificates from it. The list can be manually
retrieved from https://pkddownloadsg.icao.int/. By providing --output-format you
can get a list of hashes for a Merkle tree leaves or a full tree, by default the
keys themselves are printed.

Use cases (--output-format=hash):
- On CSCA list update, submit a proposal with the new list
- Compare proposed and actual downloaded list on voting
- Manually retrieve the tree and compare to the current list (can be done
automatically, see ldif-tree-diff)`,
		SuggestionsMinimumDistance: 2,
		Args:                       cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: use the new SDK to implement
			return fmt.Errorf("not implemented")
		},
	}

	cmd.LocalFlags().String("output-format", "raw", "Output format (raw, hash, tree)")
	return cmd
}

func cmdLDIFTreeDiff() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ldif-tree-diff <file>",
		Short: "Parse LDIF file, query the tree and compare leaves",
		Long: `Provide a LDIF file to parse. The list can be manually retrieved from
https://pkddownloadsg.icao.int/. When the file is parsed, the tree is queried
and built. The leaves are compared to the parsed list. The differences are
shown. Comparing to parse-ldif command, this one works only with hashed leaves.

Use cases:
- Automatically validate proposals: ensure that stored and proposed list differ
- Periodically check for list updates and compare to the tree`,
		SuggestionsMinimumDistance: 2,
		Args:                       cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: use the new SDK to implement
			return fmt.Errorf("not implemented")
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
