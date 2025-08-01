//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/zhaoxlchn/gin-basic-structure/internal/conf"
	"github.com/zhaoxlchn/gin-basic-structure/internal/controller"
	"github.com/zhaoxlchn/gin-basic-structure/internal/dao"
	"github.com/zhaoxlchn/gin-basic-structure/internal/data"
	"github.com/zhaoxlchn/gin-basic-structure/internal/router"
	"github.com/zhaoxlchn/gin-basic-structure/internal/server"
	"github.com/zhaoxlchn/gin-basic-structure/internal/service"
)

// wireApp init application.
func wireApp(
	serverConf *conf.Server,
	dbConf *conf.Database,
	redisConf *conf.Redis) (func() error, func(), error) {
	panic(
		wire.Build(
			data.ProviderSetData,
			router.ProviderSetRouter,
			server.ProviderSetServer,
			controller.ProviderSetController,
			dao.ProviderSetDao,
			service.ProviderSetService,
			newApp))
}
