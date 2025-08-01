package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaoxlchn/gin-basic-structure/internal/controller"
)

type ApiRouter struct {
	indexController *controller.Index
}

func NewApIRouter(indexController *controller.Index) *ApiRouter {
	return &ApiRouter{indexController: indexController}
}

func (a ApiRouter) RegisterIndexRouter(r *gin.RouterGroup) {
	r.GET("/home", a.indexController.Home)
	r.GET("/info", a.indexController.Info)
}
