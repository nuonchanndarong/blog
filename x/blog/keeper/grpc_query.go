package keeper

import (
	"github.com/nuonchanndarong/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
