package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k msgServer) SetupInitial(goCtx context.Context, msg *types.MsgSetupInitial) (*types.MsgSetupInitialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.checkSenderIsAParty(ctx, msg.Creator); err != nil {
		return nil, err
	}

	params := k.GetParams(ctx)
	if params.KeyECDSA == "" {
		params.KeyECDSA = msg.NewPublicKey
	}

	verifications := 0
	for i := range params.Parties {
		if params.Parties[i].Account == msg.Creator {
			params.Parties[i].PubKey = msg.PartyPublicKey
		}

		params.Parties[i].Verified = (params.KeyECDSA == msg.NewPublicKey && params.Parties[i].Verified) || (params.Parties[i].Account == msg.Creator)
		if params.Parties[i].Verified {
			verifications++
		}
	}

	params.KeyECDSA = msg.NewPublicKey
	if verifications == len(params.Parties) {
		params.Threshold = uint64(crypto.GetThreshold(len(params.Parties)))
		params.IsUpdateRequired = false
	}

	k.SetParams(ctx, params)
	return &types.MsgSetupInitialResponse{}, nil
}
