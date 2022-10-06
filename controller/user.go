package controller

import (
	"github.com/gin-gonic/gin"
	"todo_zhengyinuo/controller/common"
	"todo_zhengyinuo/domain"
	"todo_zhengyinuo/service"
)

func Register(ctx *gin.Context) {
	//1,参数获取和校验
	registerParam := new(domain.RegisterParam)
	//var registerParam domain.RegisterParam
	if err := ctx.ShouldBindJSON(registerParam); err != nil {
		common.ToErrorJson(ctx, common.CodeInvalidParam, nil)
		return
	}

	//2,业务逻辑处理
	err := service.Register(registerParam)
	if err != nil {
		common.ToErrorJson(ctx, common.CodeInvalidParam, nil)
		return
	}
	//3，数据返回

	//对数据封装
	common.ToErrorJson(ctx, common.CodeSuccess, nil)
}

func Login(ctx *gin.Context) {
	//1,参数获取和校验
	param := new(domain.LoginParam)
	if err := ctx.ShouldBindJSON(param); err != nil {
		common.ToErrorJson(ctx, common.CodeInvalidParam, nil)
		return
	}

	//2,业务逻辑处理
	token, err := service.Login(param)
	if err != nil {
		common.ToErrorJson(ctx, common.CodeUserNameOrPasswordError, nil)
		return
	}

	//3，数据返回
	common.ToSuccessJson(ctx, token)
	return
}
