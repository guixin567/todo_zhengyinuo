package router

import (
	"github.com/gin-gonic/gin"
	"todo_zhengyinuo/controller"
)

func Init() *gin.Engine {

	engine := gin.Default()

	userGroup := engine.Group("/user")
	userGroup.POST("/register", controller.Register)
	userGroup.POST("/login", controller.Login)

	taskGroup := engine.Group("/task")
	taskGroup.POST("/add", func(context *gin.Context) {

	})
	return engine
}