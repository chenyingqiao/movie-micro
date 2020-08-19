package logic

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtExpireTimeFormat = "2006-01-02 15:04:05"
)

type Token struct {
	token string
	time  string
}

//NewToken 获取token
func NewToken() Token {
	return Token{}
}

//GetSignString 获取jwt加密字符串
func (t *Token) GetSignString() string {
	return os.Getenv("JWT_SING")
}

//Generator 生成Token
func (t *Token) Generator(username string) (Token, error) {
	tokenEntry := Token{}

	expire := time.Now().Add(60 * 24 * 365 * time.Minute)
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"expire":   expire,
	}).SignedString([]byte(t.GetSignString()))
	if err != nil {
		return tokenEntry, err
	}

	//返回token数据
	tokenEntry.token = token
	tokenEntry.time = expire.Format(jwtExpireTimeFormat)
	return tokenEntry, nil
}

//Parse 解析 token
func (t *Token) Parse(token string) (jwt.MapClaims, error) {
	tokenEntry, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(t.GetSignString()), nil
	})

	if claims, ok := tokenEntry.Claims.(jwt.MapClaims); ok && tokenEntry.Valid {
		if t.isExpired(claims) {
			return nil, errors.New("token已经过期")
		}
		return claims, nil
	}
	return nil, err
}

//isExpired 判断是否过期
func (t *Token) isExpired(mapClaims jwt.MapClaims) bool {
	expireTime, err := time.ParseInLocation(jwtExpireTimeFormat, mapClaims["expire"].(string), time.Local)
	if err != nil {
		return false
	}
	if expireTime.After(time.Now()) {
		return true
	}
	return false
}
