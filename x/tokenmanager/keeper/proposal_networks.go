package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

func (k Keeper) HandleAddNetworkProposal(ctx sdk.Context, proposal *types.AddNetworkProposal) error {
	params := k.GetParams(ctx)
	params.Networks = append(params.Networks, &proposal.NetworkParams)
	k.SetParams(ctx, params)
	return nil
}

func (k Keeper) HandleRemoveNetworkProposal(ctx sdk.Context, proposal *types.RemoveNetworkProposal) error {
	params := k.GetParams(ctx)
	networks := make([]*types.NetworkParams, 0, len(params.Networks)-1)
	for _, network := range params.Networks {
		if network.Name != proposal.Chain {
			networks = append(networks, network)
		}
	}

	params.Networks = networks
	k.SetParams(ctx, params)
	return nil
}

func (k Keeper) HandleUpdateContractAddressProposal(ctx sdk.Context, proposal *types.UpdateContractAddressProposal) error {
	params := k.GetParams(ctx)
	for _, network := range params.Networks {
		if network.Name == proposal.Chain {
			network.Contract = proposal.Contract
			break
		}
	}

	k.SetParams(ctx, params)
	return nil
}

func (k Keeper) HandleAddFeeTokenProposal(ctx sdk.Context, proposal *types.AddFeeTokenProposal) error {
	params := k.GetParams(ctx)
	for _, network := range params.Networks {
		if network.Name == proposal.Chain {
			network.Fee.FeeTokens = append(network.Fee.FeeTokens, &proposal.Token)
			break
		}
	}

	if err := k.rarimo.CreateAddFeeTokenOperation(ctx, proposal.Token, proposal.Chain); err != nil {
		return sdkerrors.Wrap(err, "failed to create operation")
	}

	k.SetParams(ctx, params)
	return nil
}

func (k Keeper) HandleUpdateFeeTokenProposal(ctx sdk.Context, proposal *types.UpdateFeeTokenProposal) error {
	params := k.GetParams(ctx)

	for _, network := range params.Networks {
		if network.Name == proposal.Chain {
			network.Fee.FeeTokens = append(network.Fee.FeeTokens, &proposal.Token)
			break
		}
	}

	if err := k.rarimo.CreateUpdateFeeTokenOperation(ctx, proposal.Token, proposal.Chain); err != nil {
		return sdkerrors.Wrap(err, "failed to create operation")
	}

	k.SetParams(ctx, params)
	return nil
}

func (k Keeper) HandleRemoveFeeTokenProposal(ctx sdk.Context, proposal *types.RemoveFeeTokenProposal) error {
	params := k.GetParams(ctx)

	var feeTokenToRemove *types.FeeToken

	for _, network := range params.Networks {
		if network.Name == proposal.Chain {
			tokens := make([]*types.FeeToken, 0, len(network.Fee.FeeTokens)-1)

			for _, token := range network.Fee.FeeTokens {
				if token.Contract != proposal.Contract {
					tokens = append(tokens, token)
					continue
				}
				feeTokenToRemove = token
			}

			network.Fee.FeeTokens = tokens
			break
		}
	}

	if feeTokenToRemove == nil {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "fee token not found")
	}

	if err := k.rarimo.CreateRemoveFeeTokenOperation(ctx, *feeTokenToRemove, proposal.Chain); err != nil {
		return sdkerrors.Wrap(err, "failed to create operation")
	}

	k.SetParams(ctx, params)
	return nil
}
