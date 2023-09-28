package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k msgServer) SetupInitial(goCtx context.Context, msg *types.MsgSetupInitial) (*types.MsgSetupInitialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check that party is valid (active or inactive).
	// check for already active party gets an opportunity to re-submit initial
	// message in case of wrong global public key data.
	if err := k.checkIsValidParty(ctx, msg.Creator); err != nil {
		return nil, err
	}

	params := k.GetParams(ctx)
	if params.KeyECDSA != "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "can not set up: params are already initialized")
	}

	// Flag indicates that all parties were initialized
	isFullyInitialized := true

	for _, party := range params.Parties {
		if party.Account == msg.Creator {
			party.PubKey = msg.PartyPublicKey
			party.Status = types.PartyStatus_Active
			party.CommittedGlobalPublicKey = msg.NewPublicKey
		}

		isFullyInitialized = isFullyInitialized && (party.Status == types.PartyStatus_Active)
	}

	if !isFullyInitialized {
		k.SetParams(ctx, params)
		return &types.MsgSetupInitialResponse{}, nil
	}

	// If all parties was initialized -> setting up new public key and thresold.
	if params.KeyECDSA = getThresholdPublicKey(params.Parties); params.KeyECDSA != "" {
		params.IsUpdateRequired = false
		params.Threshold = uint64(crypto.GetThreshold(getActivePartiesAmount(params.Parties)))
	}

	k.SetParams(ctx, params)
	return &types.MsgSetupInitialResponse{}, nil
}

// getThresholdPublicKey checks all committed pub keys and if all is equal returns it, otherwise - returns empty string
func getThresholdPublicKey(parties []*types.Party) string {
	key := parties[0].CommittedGlobalPublicKey
	for _, party := range parties {
		if key != party.CommittedGlobalPublicKey {
			return ""
		}
	}

	return key
}
