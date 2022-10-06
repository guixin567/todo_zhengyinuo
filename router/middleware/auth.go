package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"todo_zhengyinuo/controller/response"
	"todo_zhengyinuo/pkg/jwt"
)

// JwtAuth
// 1,从请求头获取Authorization
// 2,判断是否有token
// 3,解析token
// 4，获取userId，并添加在context里面，其他的请求里面有需要用到userId的时候直接获取就可以了
func JwtAuth(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		response.ToErrorJson(ctx, response.CodeAuthInvalid, nil)
		ctx.Abort()
		return
	}

	authHeader = authHeader[7:]
	println("zhuwei ", authHeader)
	token, err := jwt.ParseToken(authHeader)
	if err != nil || token == nil {
		response.ToErrorJson(ctx, response.CodeNeedLogin, nil)
		ctx.Abort()
		return
	}
	userId := token.UserId
	ctx.Set(response.KeyUserId, userId)
	ctx.Next()
}
