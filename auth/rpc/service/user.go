package service

import (
	"auth/logic"
	"auth/rpc/protos"
	"context"

	"github.com/sirupsen/logrus"
)

//UserService 用户注册服务
type UserService struct {
	protos.UnimplementedUserServer
}

//Register 注册
func (u *UserService) Register(ctx context.Context, registerRequest *protos.RegisterRequest) (*protos.RegisterResponse, error) {
	userLogic := logic.NewUser()
	err := userLogic.Register(registerRequest)
	if err != nil {
		return &protos.RegisterResponse{
			Status: false,
		}, err
	}
	logrus.WithField("info", registerRequest).Info("注册用户")

	return &protos.RegisterResponse{
		Status: true,
	}, err
}
