package service

import (
	"auth/logic"
	"auth/model"
	"auth/rpc/protos"
	"context"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

//AuthService 校验rpc服务
type AuthService struct {
	protos.UnimplementedAuthServer
}

//GetToken 获取token
func (auth *AuthService) GetToken(ctx context.Context, tokenRequest *protos.TokenRequest) (*protos.TokenResponse, error) {

	userInfo := model.NewEmptyUser()
	user, err := userInfo.Info(tokenRequest.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(tokenRequest.Password))
	if err != nil {
		return nil, err
	}

	//生成token
	tokenLogic := logic.NewJwtAuth()
	token, err := tokenLogic.GetToken(tokenRequest)
	if err != nil {
		return &protos.TokenResponse{
			Token: token,
		}, err
	}

	response := &protos.TokenResponse{
		Token: token,
	}
	logrus.WithField("info", tokenRequest).Info("获取token")

	return response, nil
}

//Validate 校验是否是登录状态
func (auth *AuthService) Validate(ctx context.Context, validateRequest *protos.ValidateRequest) (*protos.ValidateResponse, error) {
	tokenLogic := logic.NewJwtAuth()
	isSuccess := tokenLogic.Validate(validateRequest.Token)
	response := &protos.ValidateResponse{
		Status: isSuccess,
	}

	return response, nil
}

//GetInfo 获取用户信息
func (auth *AuthService) GetInfo(ctx context.Context, getInfoRequest *protos.GetInfoRequest) (*protos.GetInfoResponse, error) {
	tokenLogic := logic.NewJwtAuth()
	authInfo, err := tokenLogic.GetInfo(getInfoRequest.Token)
	if err != nil {
		return &protos.GetInfoResponse{
			Status: "success",
			Info:   &protos.Info{},
			Code:   200,
		}, err
	}

	response := &protos.GetInfoResponse{
		Status: "success",
		Info: &protos.Info{
			Userid:   authInfo["userid"],
			Username: authInfo["username"],
			Power:    strings.Split(authInfo["power"], ","),
			Expire:   authInfo["expire"],
		},
		Code: 200,
	}

	logrus.WithField("info", getInfoRequest).Info("获取用户信息")

	return response, nil
}
