package router

import "github.com/google/wire"

var ProviderSetRouter = wire.NewSet(NewApIRouter)
