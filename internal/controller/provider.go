package controller

import "github.com/google/wire"

var ProviderSetController = wire.NewSet(NewIndex)
