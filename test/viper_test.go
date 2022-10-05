package test

import (
	"github.com/spf13/viper"
	"testing"
	"todo_zhengyinuo/service"
	"todo_zhengyinuo/test/config"
)

func TestViper(test *testing.T) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	config1 := &config.TestConfig{}
	err = viper.Unmarshal(config1)
	if err != nil {
		println("zhuwei", err)
	}
	redisName := viper.GetString("service.app_mode")

	println("zhengxin", redisName)
	println("zhengxin222", config1.AppMode)
}

func TestMd5(test *testing.T) {
	password := service.EncryptPassword("1234")
	println(password)
}
