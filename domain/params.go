package domain

type RegisterParam struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"rePassword" binding:"required,eqfield=Password"`
}

type LoginParam struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
