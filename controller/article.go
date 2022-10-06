package controller

import (
	"github.com/gin-gonic/gin"
	"todo_zhengyinuo/controller/response"
	"todo_zhengyinuo/domain"
	"todo_zhengyinuo/service"
)

func AddArticle(ctx *gin.Context) {
	//1,参数绑定和校验
	param := new(domain.AddArticleParam)
	if err := ctx.ShouldBindJSON(param); err != nil {
		response.ToErrorJson(ctx, response.CodeInvalidParam, nil)
		return
	}

	userId := ctx.GetInt64(response.KeyUserId)
	param.AuthorId = userId
	//2，添加数据到数据库
	article, err := service.AddArticle(param)
	if err != nil {
		response.ToErrorJson(ctx, response.CodeServerError, nil)
		return
	}
	//3，返回数据
	response.ToSuccessJson(ctx, article)

}
