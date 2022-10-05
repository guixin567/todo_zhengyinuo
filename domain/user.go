package domain

type User struct {
	ID       int64  `gorm:"column:id" db:"id" json:"id"`
	UserId   int64  `gorm:"column:user_id" db:"user_id" json:"user_id"`
	Username string `gorm:"column:username" db:"username" json:"username"`
	Password string `gorm:"column:password" db:"password" json:"password"`
	Email    string `gorm:"column:email" db:"email" json:"email"`
	Gender   int64  `gorm:"column:gender" db:"gender" json:"gender"`
}

func (User) TableName() string {
	return "user"
}
