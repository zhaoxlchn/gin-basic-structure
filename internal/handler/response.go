package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const Success = 200
const SuccessMsg = "success"
const Err = 0
const ErrorMsg = "network error"

// Response body
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    any    `json:"data"`
}

func ResSuccess(ctx *gin.Context, msg string, data ...any) {
	res := &Response{
		Code:    Success,
		Message: msg,
	}
	if 0 < len(data) {
		res.Data = data[0]
	}
	ctx.JSON(http.StatusOK, res)
}

func ResError(ctx *gin.Context, msg ...string) {
	res := &Response{
		Code: Err,
	}
	if 0 < len(msg) {
		res.Message = msg[0]
	}
	ctx.JSON(http.StatusOK, res)
}
