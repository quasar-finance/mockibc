package keeper

import (
	"mockibc/x/mockibc/types"
)

var _ types.QueryServer = Keeper{}
