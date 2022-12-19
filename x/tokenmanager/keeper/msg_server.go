package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) checkCreatorIsValidator(ctx sdk.Context, creator string) bool {
	addr, _ := sdk.ValAddressFromBech32(creator)
	return k.staking.Validator(ctx, addr) != nil
}
