package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gitlab.com/rarimo/rarimo-core/x/rarimocore/types"
)

func (k msgServer) ChangePartyAddress(goCtx context.Context, msg *types.MsgChangePartyAddress) (*types.MsgChangePartyAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)

	// Party can change only self address
	party := getPartyByAccount(msg.Creator, params.Parties)
	if party == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "party not found")
	}

	party.Address = msg.NewAddress
	k.SetParams(ctx, params)
	return &types.MsgChangePartyAddressResponse{}, nil
}
