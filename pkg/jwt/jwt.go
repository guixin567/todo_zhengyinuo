package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

//1,声明一个实体
//2,设置jwt参数

type TodoClaim struct {
	UserId int64 `json:"user_id"`
	*jwt.StandardClaims
}

var secret = []byte("zhengyinuo")

func GetToken(userId int64) (token string, err error) {
	claims := TodoClaim{
		UserId: userId,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
			Issuer:    "todo",
		},
	}
	claimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return claimToken.SignedString(secret)
}

func ParseToken(tokenString string) (claim *TodoClaim, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &TodoClaim{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if todoClaim, ok := token.Claims.(*TodoClaim); ok && token.Valid {
		return todoClaim, nil
	}
	return nil, errors.New("invalid token")

}
