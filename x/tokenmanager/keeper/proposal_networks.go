package keeper

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

var upgradeVerifiers = map[types.NetworkType]func(details *types.ContractUpgradeDetails) error{
	types.NetworkType_EVM: func(details *types.ContractUpgradeDetails) error {
		if _, err := hexutil.Decode(details.TargetContract); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid target contract address %s", err.Error())
		}

		if _, err := hexutil.Decode(details.NewImplementationContract); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid new contract address %s", err.Error())
		}

		if _, ok := new(big.Int).SetString(details.Nonce, 10); !ok {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid nonce")
		}

		return nil
	},
	types.NetworkType_Solana: func(details *types.ContractUpgradeDetails) error {
		if _, err := hexutil.Decode(details.TargetContract); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid target contract address %s", err.Error())
		}

		if _, err := hexutil.Decode(details.BufferAccount); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid buffer address %s", err.Error())
		}

		if _, ok := new(big.Int).SetString(details.Nonce, 10); !ok {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid nonce")
		}

		return nil
	},
	types.NetworkType_Near: func(details *types.ContractUpgradeDetails) error {
		if _, err := hexutil.Decode(details.TargetContract); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid target contract address %s", err.Error())
		}

		if _, err := hexutil.Decode(details.Hash); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid hash %s", err.Error())
		}

		if _, ok := new(big.Int).SetString(details.Nonce, 10); !ok {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid nonce")
		}

		return nil
	},
}

func (k Keeper) HandleUpgradeContractProposal(ctx sdk.Context, proposal *types.UpgradeContractProposal) error {
	network, ok := k.GetNetwork(ctx, proposal.Details.Chain)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "network not found")
	}

	if err := upgradeVerifiers[network.Type](&proposal.Details); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "proposal validation failed %s", err.Error())
	}

	return k.rarimo.CreateContractUpgradeOperation(ctx, &proposal.Details)
}

func (k Keeper) HandleAddNetworkProposal(ctx sdk.Context, proposal *types.AddNetworkProposal) error {
	params := k.GetParams(ctx)
	params.Networks = append(params.Networks, &proposal.Network)
	k.SetParams(ctx, params)
	return nil
}

func (k Keeper) HandleRemoveNetworkProposal(ctx sdk.Context, proposal *types.RemoveNetworkProposal) error {
	if _, ok := k.GetNetwork(ctx, proposal.Chain); !ok {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "network not found")
	}

	params := k.GetParams(ctx)

	networks := make([]*types.Network, 0, len(params.Networks)-1)
	for _, network := range params.Networks {
		if network.Name != proposal.Chain {
			networks = append(networks, network)
		}
	}

	params.Networks = networks
	k.SetParams(ctx, params)
	return nil
}

func (k Keeper) HandleAddFeeTokenProposal(ctx sdk.Context, proposal *types.AddFeeTokenProposal) error {
	if _, ok := k.GetNetwork(ctx, proposal.Chain); !ok {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "network not found")
	}

	params := k.GetParams(ctx)
	for _, network := range params.Networks {
		if network.Name == proposal.Chain {
			feeparams := network.GetFeeParams()
			if feeparams == nil {
				return sdkerrors.Wrap(sdkerrors.ErrNotFound, "fee params not found")
			}

			feeparams.SetFeeToken(&proposal.Token)
			break
		}
	}

	if err := k.rarimo.CreateAddFeeTokenOperation(ctx, proposal.Token, proposal.Chain, proposal.Nonce); err != nil {
		return sdkerrors.Wrap(err, "failed to create operation")
	}

	k.SetParams(ctx, params)
	return nil
}

func (k Keeper) HandleUpdateFeeTokenProposal(ctx sdk.Context, proposal *types.UpdateFeeTokenProposal) error {
	if _, ok := k.GetNetwork(ctx, proposal.Chain); !ok {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "network not found")
	}

	params := k.GetParams(ctx)

	for _, network := range params.Networks {
		if network.Name == proposal.Chain {
			feeparams := network.GetFeeParams()
			if feeparams == nil {
				return sdkerrors.Wrap(sdkerrors.ErrNotFound, "fee params not found")
			}

			feeparams.SetFeeToken(&proposal.Token)
			break
		}
	}

	if err := k.rarimo.CreateUpdateFeeTokenOperation(ctx, proposal.Token, proposal.Chain, proposal.Nonce); err != nil {
		return sdkerrors.Wrap(err, "failed to create operation")
	}

	k.SetParams(ctx, params)
	return nil
}

func (k Keeper) HandleRemoveFeeTokenProposal(ctx sdk.Context, proposal *types.RemoveFeeTokenProposal) error {
	if _, ok := k.GetNetwork(ctx, proposal.Chain); !ok {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "network not found")
	}

	params := k.GetParams(ctx)

	var feeTokenToRemove *types.FeeToken

	for _, network := range params.Networks {
		if network.Name == proposal.Chain {
			feeparams := network.GetFeeParams()
			if feeparams == nil {
				return sdkerrors.Wrap(sdkerrors.ErrNotFound, "fee params not found")
			}

			feeTokenToRemove = feeparams.GetFeeToken(proposal.Contract)
			feeparams.RemoveFeeToken(feeTokenToRemove)
			break
		}
	}

	if feeTokenToRemove == nil {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "fee token not found")
	}

	if err := k.rarimo.CreateRemoveFeeTokenOperation(ctx, *feeTokenToRemove, proposal.Chain, proposal.Nonce); err != nil {
		return sdkerrors.Wrap(err, "failed to create operation")
	}

	k.SetParams(ctx, params)
	return nil
}

func (k Keeper) HandleWithdrawFeeProposal(ctx sdk.Context, proposal *types.WithdrawFeeProposal) error {
	network, ok := k.GetNetwork(ctx, proposal.Chain)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "network not found")
	}

	feeparams := network.GetFeeParams()
	if feeparams == nil {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "fee params not found")
	}

	feeTokenToWithdraw := feeparams.GetFeeToken(proposal.Token.Contract)
	if feeTokenToWithdraw == nil {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "fee token not found")
	}

	if err := k.rarimo.CreateWithdrawFeeOperation(ctx, proposal.Token, proposal.Chain, proposal.Receiver, proposal.Nonce); err != nil {
		return sdkerrors.Wrap(err, "failed to create operation")
	}

	return nil
}
