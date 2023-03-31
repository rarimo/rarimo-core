package keeper

import (
	"gitlab.com/rarimo/rarimo-core/x/multisig/types"
)

var _ types.QueryServer = Keeper{}
