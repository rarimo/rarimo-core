package keeper

import (
	"gitlab.com/rarimo/rarimo-core/x/bridge/types"
)

var _ types.QueryServer = Keeper{}
