package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rarimo/rarimo-core/x/rarimocore/types"
)

type RarimocoreKeeper interface {
	CreateRootUpdateOperation(ctx sdk.Context, creator string, update *types.RootUpdate) (string, error)
}
