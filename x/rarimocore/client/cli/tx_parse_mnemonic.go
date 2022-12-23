package cli

import (
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"
)

func CmdParseMnemonic() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "parse-mnemonic [mnemonic]",
		Short: "Get hex private key from mnemonic",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			mnemonic := args[0]

			path := hd.CreateHDPath(118, 0, 0)
			algo := hd.Secp256k1

			derivedPrv, err := algo.Derive()(mnemonic, "", path.String())
			if err != nil {
				return err
			}

			privateKey := algo.Generate()(derivedPrv)
			coreAddress, _ := bech32.ConvertAndEncode("rarimo", privateKey.PubKey().Address().Bytes())

			cmd.Println("Private Key: ", hexutil.Encode(privateKey.Bytes()))
			cmd.Println("Rarimo address: ", coreAddress)
			cmd.Println("Public key:", hexutil.Encode(privateKey.PubKey().Bytes()))
			return nil
		},
	}

	return cmd
}
