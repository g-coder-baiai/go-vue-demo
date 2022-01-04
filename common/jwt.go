package common

import (
	"gin_vue/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//密钥
var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	//设置token的过期时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			//token发放的时间
			IssuedAt: time.Now().Unix(),
			//token的发放者
			Issuer: "oceanlearn.tech",
			//主题
			Subject: "user token",
		},
	}
	//使用jwt密钥生成我们的token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil

}

// ParseToken 解析tokenString中的claims，并返回
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err

}
