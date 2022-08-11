package keeper

import (
	"github.com/lubtd/planet/x/planet/types"
)

var _ types.QueryServer = Keeper{}
