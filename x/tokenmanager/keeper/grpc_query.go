package keeper

import (
	"github.com/rarimo/rarimo-core/x/tokenmanager/types"
)

var _ types.QueryServer = Keeper{}
