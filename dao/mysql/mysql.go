package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"todo_zhengyinuo/config"
)

var db *gorm.DB

func Init(config *config.MysqlConfig) (err error) {
	//dsn := "root:root@tcp(127.0.0.1:3306)/todo_list1?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn1 := "root:root@tcp(127.0.0.1:3306)/todo_list1?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := strings.Join([]string{config.DbUser, ":", config.DbPassword, "@tcp(", config.DbHost, ":", config.DbPort, ")/", config.DbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	println(dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}
