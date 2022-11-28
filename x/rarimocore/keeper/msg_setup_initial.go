package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/crypto"
	"gitlab.com/rarify-protocol/rarimo-core/x/rarimocore/types"
)

func (k msgServer) SetupInitial(goCtx context.Context, msg *types.MsgSetupInitial) (*types.MsgSetupInitialResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.checkSenderIsAParty(ctx, msg.Creator); err != nil {
		return nil, err
	}

	params := k.GetParams(ctx)

	if len(msg.Set) != len(params.Parties) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid set size")
	}

	if params.KeyECDSA == "" {
		for i := range msg.Set {
			if msg.Set[i].Account != params.Parties[i].Account {
				return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid set accounts")
			}
		}
	}

	key, err := getECDSAPubKey(msg.Set)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid parties pub keys")
	}

	if err := crypto.VerifyECDSA(msg.Signature, hexutil.Encode(crypto.GetPartiesHash(msg.Set)), key); err != nil {
		return nil, err
	}

	params.KeyECDSA = key
	params.Threshold = (uint64(len(msg.Set)+2) / 3) * 2
	params.IsUpdateRequired = false
	params.Parties = msg.Set
	k.SetParams(ctx, params)
	return &types.MsgSetupInitialResponse{}, nil
}
