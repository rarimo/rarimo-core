package types

import (
	"cosmossdk.io/errors"
)

var ErrUnknownRequest = errors.Register("cscalist", 6, "unknown request")
