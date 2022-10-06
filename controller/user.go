package controller

import (
	"github.com/gin-gonic/gin"
	"todo_zhengyinuo/controller/response"
	"todo_zhengyinuo/domain"
	"todo_zhengyinuo/service"
)

func Register(ctx *gin.Context) {
	//1,参数获取和校验
	registerParam := new(domain.RegisterParam)
	//var registerParam domain.RegisterParam
	if err := ctx.ShouldBindJSON(registerParam); err != nil {
		response.ToErrorJson(ctx, response.CodeInvalidParam, nil)
		return
	}

	//2,业务逻辑处理
	err := service.Register(registerParam)
	if err != nil {
		response.ToErrorJson(ctx, response.CodeInvalidParam, nil)
		return
	}
	//3，数据返回

	//对数据封装
	response.ToErrorJson(ctx, response.CodeSuccess, nil)
}

func Login(ctx *gin.Context) {
	//1,参数获取和校验
	param := new(domain.LoginParam)
	if err := ctx.ShouldBindJSON(param); err != nil {
		response.ToErrorJson(ctx, response.CodeInvalidParam, nil)
		return
	}

	//2,业务逻辑处理
	token, err := service.Login(param)
	if err != nil {
		response.ToErrorJson(ctx, response.CodeUserNameOrPasswordError, nil)
		return
	}

	//3，数据返回
	response.ToSuccessJson(ctx, token)
	return
}
