package controller

import (
	"github.com/gin-gonic/gin"
	"todo_zhengyinuo/controller/response"
	"todo_zhengyinuo/service"
)

func Category(ctx *gin.Context) {
	//1,参数绑定和校验
	//2,业务逻辑处理
	categories, err := service.Categories()
	if err != nil {
		response.ToErrorJson(ctx, response.CodeServerError, nil)
		return
	}
	//3，返回数据
	response.ToSuccessJson(ctx, categories)
}
