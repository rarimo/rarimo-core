package keeper

import (
	"gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/types"
)

var _ types.QueryServer = Keeper{}
