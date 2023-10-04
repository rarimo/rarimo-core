package keeper

import (
	"github.com/rarimo/rarimo-core/x/vestingmint/types"
)

var _ types.QueryServer = Keeper{}
