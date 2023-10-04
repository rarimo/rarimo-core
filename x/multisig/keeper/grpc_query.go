package keeper

import (
	"github.com/rarimo/rarimo-core/x/multisig/types"
)

var _ types.QueryServer = Keeper{}
