package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const salt = "&(go_spider)&"

type JwtUtil struct {
}

type Claim struct {
	Id   int
	Name string
	jwt.StandardClaims
}

func (j JwtUtil) GetToken(id int, name string, expire time.Duration) string {
	c := Claim{Id: id, Name: name, StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(expire).Unix(), Issuer: "learn-GSP"}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	res, _ := token.SignedString([]byte(salt))
	return res
}

func (j JwtUtil) ValidateToken(token string) bool {
	t, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(salt), nil
	})
	if err != nil {
		return false
	}
	if t != nil {
		if _, ok := t.Claims.(*Claim); ok && t.Valid {
			return true
		}
	}
	return false
}
