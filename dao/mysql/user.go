package mysql

import "todo_zhengyinuo/domain"

func Register() {

}

// QueryUserByUserName 判断用户是否存在
func QueryUserByUserName() {

}

func CheckUserExit(userName string) bool {
	user := domain.User{Username: userName}
	db.First(&user)
	println("userid:", user.UserId)
	return user.UserId != 0
}

func InsertUser(user *domain.User) {
	db.Create(user)
}
