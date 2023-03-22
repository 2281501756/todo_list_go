package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JWTSecret = "QWESAQWE123da"

type Claims struct {
	Id       uint   `json:"id"`
	UserName string `json:"userName"`
	jwt.StandardClaims
}

// 生成token

func GeneratorToken(id uint, username string, password string) (string, error) {
	claims := Claims{
		Id:             id,
		UserName:       username,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), Issuer: "todo_list"},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWTSecret))
	return tokenString, err
}

// 解析token

func ParseToken(token string) (*Claims, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, err
}
