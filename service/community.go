package service

import (
	"todo_zhengyinuo/dao/mysql"
	"todo_zhengyinuo/domain"
)

func Categories() (communities []domain.Community, err error) {
	return mysql.CommunityCategory()
}
