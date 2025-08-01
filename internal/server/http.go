package server

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/zhaoxlchn/gin-basic-structure/internal/router"
)

var ProviderSetServer = wire.NewSet(NewHTTPServer)

func NewHTTPServer(apiRouter *router.ApiRouter) *gin.Engine {
	r := gin.Default()
	r.GET("/health", func(ctx *gin.Context) { ctx.String(200, "OK") })

	g := r.Group("api")
	apiRouter.RegisterIndexRouter(g)
	return r
}
