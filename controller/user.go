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
	service.Register(registerParam)
	//3，数据返回

	//对数据封装
	ctx.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "",
		"data": "ddd",
	})
}

func Login(ctx *gin.Context) {
	//1,参数获取和校验
	//2,业务逻辑处理
	//3，数据返回
}
