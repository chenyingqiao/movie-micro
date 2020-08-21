package logic

import (
	"auth/model"
	"auth/rpc/protos"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

//Auth 鉴权
type Auth interface {
	//GetToken 获取token
	GetToken(request *protos.TokenRequest) (string, error)
	//Validate 校验token是否有效
	Validate(token string) bool
	//GetInfo 获取用户的信息
	GetInfo(token string) (jwt.MapClaims, error)
}

//JwtAuth auth
type JwtAuth struct {
}

//NewJwtAuth NewJwtAuth
func NewJwtAuth() JwtAuth {
	return JwtAuth{}
}

//GetToken 获取token
func (ja *JwtAuth) GetToken(request *protos.TokenRequest) (string, error) {
	tokenEntry := NewToken()
	token, err := tokenEntry.Generator(request.Username)
	if err != nil {
		return "", err
	}
	return token.GetToken(), nil
}

//Validate 校验token是否有效
func (ja *JwtAuth) Validate(token string) bool {
	tokenEntry := NewToken()
	return tokenEntry.Validate(token)
}

//GetInfo 获取用户的信息
func (ja *JwtAuth) GetInfo(token string) (map[string]string, error) {
	tokenEntry := NewToken()
	mapC, err := tokenEntry.Parse(token)
	if err != nil {
		return nil, err
	}

	authInfo := mapC.(jwt.MapClaims)

	user := model.NewEmptyUser()
	userInfo, err := user.Info(authInfo[ParseKeyUsername].(string))

	result := map[string]string{
		"userid":   userInfo.ID.Hex(),
		"username": userInfo.Username,
		"power":    strings.Join(userInfo.Power, ","),
		"expire":   authInfo[ParseKeyExpire].(string),
	}

	return result, nil
}
