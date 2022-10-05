package service

import (
	"todo_zhengyinuo/dao/mysql"
	"todo_zhengyinuo/domain"
	"todo_zhengyinuo/pkg/snow_flake"
)

func Register(param *domain.RegisterParam) {
	//1 判断用户是否存在
	mysql.QueryUserByUserName()
	//2 生成用户ID
	snow_flake.GetSingleID()
	//3 密码加密

	//4 保存数据库
	mysql.InsertUser()
}
