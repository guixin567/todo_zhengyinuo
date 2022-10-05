package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"todo_zhengyinuo/config"
)

var db *gorm.DB

func Init(config *config.MysqlConfig) (err error) {
	dsn := strings.Join([]string{config.DbUser, ":", config.DbPassword, "@tcp(", config.DbHost, ":", config.DbPort, ")/", config.DbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	println(dsn)
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,
		DisableDatetimePrecision: true,
	}))
	return
}
