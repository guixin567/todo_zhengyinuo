package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"todo_zhengyinuo/config"
)

var Db *redis.Client

func Init(config *config.RedisConfig) (err error) {
	Db = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		DB:       config.Name,
		Password: config.Password,
		PoolSize: config.PoolSize,
	})
	_, err = Db.Ping().Result()

	return

}
