package main

import (
	"time"
	"todo_zhengyinuo/config"
	"todo_zhengyinuo/dao/mysql"
	"todo_zhengyinuo/dao/redis"
	"todo_zhengyinuo/router/router"
)

func main() {
	config.Init()
	mysql.Init(config.Global.MysqlConfig)
	redis.Init(config.Global.RedisConfig)
	cmd := redis.Db.Set("zhengxin", "jiayou", 1*time.Minute)
	result, err := cmd.Result()
	if err != nil {
		println("zhengyibo", err.Error(), result)
	}

	app := router.Init()
	app.Run(":8083")
}
