package main

import (
	"todo_zhengyinuo/config"
	"todo_zhengyinuo/dao/mysql"
	"todo_zhengyinuo/router/router"
)

func main() {
	config.Init()
	mysql.Init(config.Global.MysqlConfig)
	app := router.Init()
	app.Run(":8083")
}
