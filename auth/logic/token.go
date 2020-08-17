package logic

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type token struct {
}

func Generator(username string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"expire":time.Date()
	})
}
