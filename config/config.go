package config

import (
	"github.com/spf13/viper"
)

type GlobalConfig struct {
	MysqlConfig *MysqlConfig `mapstructure:"mysql"`
	RedisConfig *RedisConfig `mapstructure:"redis"`
}

type MysqlConfig struct {
	Db         string `mapstructure:"db"`
	DbHost     string `mapstructure:"dbHost"`
	DbPort     string `mapstructure:"dbPort"`
	DbUser     string `mapstructure:"dbUser"`
	DbPassword string `mapstructure:"dbPassword"`
	DbName     string `mapstructure:"dbName"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     int    `mapstructure:"name"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"poolSize"`
}

var Global = new(GlobalConfig)

func Init() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()

	if err != nil {
		println(err)
	}
	err = viper.Unmarshal(Global)
	if err != nil {
		println("zhengxin222", err)
	}
	println("zhengxin", Global.MysqlConfig, viper.GetString("mysql.db"))
}
