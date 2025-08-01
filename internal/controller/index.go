package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaoxlchn/gin-basic-structure/internal/handler"
	"github.com/zhaoxlchn/gin-basic-structure/internal/service"
)

type Index struct {
	userService *service.UserService
}

func NewIndex(userService *service.UserService) *Index {
	return &Index{userService: userService}
}

func (i Index) Home(ctx *gin.Context) {
	handler.ResSuccess(ctx, "hello")
}

func (i Index) Info(ctx *gin.Context) {
	id := ctx.GetInt("id")
	user, err := i.userService.GetUser(ctx, int64(id))
	if err != nil {
		handler.ResError(ctx, handler.ErrorMsg+err.Error())
		return
	}
	handler.ResSuccess(ctx, handler.SuccessMsg, user)
}
