package keeper

import (
	"github.com/rarimo/rarimo-core/x/oraclemanager/types"
)

var _ types.QueryServer = Keeper{}
