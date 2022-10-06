package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_zhengyinuo/domain"
	"todo_zhengyinuo/service"
)

func Register(ctx *gin.Context) {
	//1,参数获取和校验
	registerParam := new(domain.RegisterParam)
	//var registerParam domain.RegisterParam
	if err := ctx.ShouldBindJSON(registerParam); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	//2,业务逻辑处理
	err := service.Register(registerParam)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	//3，数据返回

	//对数据封装
	ctx.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "register success",
	})
}

func Login(ctx *gin.Context) {
	//1,参数获取和校验
	param := new(domain.LoginParam)
	if err := ctx.ShouldBindJSON(param); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "username or password is wrong",
		})
		return
	}

	//2,业务逻辑处理
	if err := service.Login(param); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	//3，数据返回
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "login success",
	})
	return
}
