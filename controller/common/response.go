package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int64  `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

func ToSuccessJson(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Msg:  getMsg(CodeSuccess),
		Data: data,
	})
}

func ToErrorJson(ctx *gin.Context, code int64, data any) {
	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  getMsg(code),
		Data: data,
	})
}

//func (ctx * gin.Context) ToSuccessJson(data any) {
//	ctx.JSON(http.StatusOK, Response{
//		Code: CodeSuccess,
//		Msg:  getMsg(CodeSuccess),
//		Data: data,
//	})
//}
