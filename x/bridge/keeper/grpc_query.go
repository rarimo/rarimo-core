package keeper

import (
	"github.com/rarimo/rarimo-core/x/bridge/types"
)

var _ types.QueryServer = Keeper{}
