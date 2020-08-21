package service

import (
	"auth/logic"
	"auth/rpc/protos"
	"context"
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

	return &protos.RegisterResponse{
		Status: true,
	}, err
}
