package keeper

import (
	"gitlab.com/rarimo/rarimo-core/x/oraclemanager/types"
)

var _ types.QueryServer = Keeper{}
