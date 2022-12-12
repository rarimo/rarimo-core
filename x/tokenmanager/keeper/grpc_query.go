package keeper

import (
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

var _ types.QueryServer = Keeper{}
