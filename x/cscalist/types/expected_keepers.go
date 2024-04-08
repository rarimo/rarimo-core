package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	rarimo "github.com/rarimo/rarimo-core/x/rarimocore/types"
)

type RarimocoreKeeper interface {
	CreateCSCARootUpdateOperation(ctx sdk.Context, creator string, update *rarimo.CSCARootUpdate) error
}
