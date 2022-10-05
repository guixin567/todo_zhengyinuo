package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"todo_zhengyinuo/dao/mysql"
	"todo_zhengyinuo/domain"
	"todo_zhengyinuo/pkg/snow_flake"
)

func Register(param *domain.RegisterParam) (err error) {
	//1 判断用户是否存在
	if mysql.CheckUserExit(param.Username) {
		err = errors.New("用户已经存在")
		return
	}
	//2 生成用户ID
	uid := snow_flake.GetSingleID()
	//3 密码加密
	password := EncryptPassword(param.Password)
	user := &domain.User{
		UserId:   uid,
		Username: param.Username,
		Password: password,
	}

	fmt.Println(user)
	//4 保存数据库
	mysql.InsertUser(user)
	return
}

var secret = "zhuwei"

func EncryptPassword(password string) string {
	hash := md5.New()
	hash.Write([]byte(secret))
	return hex.EncodeToString(hash.Sum([]byte(password)))
}
