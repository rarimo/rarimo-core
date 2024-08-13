package keeper

import (
	"github.com/rarimo/rarimo-core/x/rootupdater/types"
)

var _ types.QueryServer = Keeper{}
