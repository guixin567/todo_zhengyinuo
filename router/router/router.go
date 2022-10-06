package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_zhengyinuo/controller"
	"todo_zhengyinuo/controller/response"
	"todo_zhengyinuo/router/middleware"
)

func Init() *gin.Engine {

	engine := gin.Default()

	userGroup := engine.Group("/user")
	userGroup.POST("/register", controller.Register)
	userGroup.POST("/login", controller.Login)

	taskGroup := engine.Group("/task")
	taskGroup.POST("/add", middleware.JwtAuth, func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"userId": context.GetInt64(response.KeyUserId)})
	})

	taskGroup.GET("/category", controller.Category)
	return engine
}
