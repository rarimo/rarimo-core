package keeper

import (
	"context"
	"encoding/binary"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"gitlab.com/rarimo/rarimo-core/x/multisig/types"
)

func (k msgServer) CreateGroup(goCtx context.Context, msg *types.MsgCreateGroup) (*types.MsgCreateGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	for _, member := range msg.Members {
		addr, _ := sdk.AccAddressFromBech32(member)
		acc := k.accountKeeper.GetAccount(ctx, addr)
		if acc == nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "member's account %s does not exist", addr)
		}
	}

	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	params := k.GetParams(ctx)

	var groupAccAddr sdk.AccAddress
	// loop here in the rare case where an ADR-028-derived address creates a
	// collision with an existing address.
	for {
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, params.GroupSequence+1)

		parentAcc := address.Module(types.ModuleName, []byte{types.GroupAccountKey})
		groupAccAddr = address.Derive(parentAcc, buf)

		if k.accountKeeper.GetAccount(ctx, groupAccAddr) != nil {
			// handle a rare collision, in which case we just go on to the
			// next sequence value and derive a new address.
			continue
		}
		acc := k.accountKeeper.NewAccount(ctx, &authtypes.ModuleAccount{
			BaseAccount: &authtypes.BaseAccount{
				Address: groupAccAddr.String(),
			},
			Name: groupAccAddr.String(),
		})
		k.accountKeeper.SetAccount(ctx, acc)

		break
	}

	k.SetGroup(ctx, types.Group{
		Account:   moduleAddr.String(),
		Members:   msg.Members,
		Threshold: msg.Threshold,
	})

	params.GroupSequence++
	k.SetParams(ctx, params)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreateGroup,
			sdk.NewAttribute(types.AttributeKeyGroup, groupAccAddr.String()),
		),
	)

	return &types.MsgCreateGroupResponse{
		Group: groupAccAddr.String(),
	}, nil
}
