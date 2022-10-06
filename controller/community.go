package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
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

func CategoryDetail(ctx *gin.Context) {
	//1,参数绑定和校验
	id, err := strconv.Atoi(ctx.Param(response.KeyCommunityId))
	communityId := int64(id)
	if err != nil {
		println("zhuwei ", err.Error())
		response.ToErrorJson(ctx, response.CodeInvalidParam, nil)
		return
	}
	//2,业务逻辑处理
	category, err := service.CategoryDetail(communityId)
	if err != nil {
		response.ToErrorJson(ctx, response.CodeServerError, nil)
		return
	}
	//3，返回数据
	response.ToSuccessJson(ctx, category)
}
