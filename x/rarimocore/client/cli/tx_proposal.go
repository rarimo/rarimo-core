package cli

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cosmostypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/spf13/cobra"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func CmdAddSignerPartyProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-signer-proposal [title] [description] [account] [address] [trial-prv-key] [deposit]",
		Short: "Create add signer party proposal",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			title := args[0]
			description := args[1]
			account := args[2]
			address := args[3]
			prv := args[4]
			deposit, ok := sdk.NewIntFromString(args[5])
			if !ok {
				return errors.New("invalid deposit amount")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			content := &types.AddSignerPartyProposal{
				Title:          title,
				Description:    description,
				Account:        account,
				Address:        address,
				TrialPublicKey: prv,
			}

			contentDetails, err := cosmostypes.NewAnyWithValue(content)
			if err != nil {
				panic(err)
			}

			msg := &govtypes.MsgSubmitProposal{
				Content: contentDetails,
				InitialDeposit: []sdk.Coin{
					{
						Denom:  "stake",
						Amount: deposit,
					},
				},
				Proposer: clientCtx.GetFromAddress().String(),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
