package client

import (
	"context"
	"movie/errs"
	"movie/rpc/protos"
	"movie/utils"
	"time"
)

//UserClient 用户注册rpc链接
type UserClient struct{}

//NewUserClient NewUserClient
func NewUserClient() UserClient {
	return UserClient{}
}

//Register 用户注册
func (u *UserClient) Register(request *protos.RegisterRequest) (protos.RegisterResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := utils.OpenGrpcClientConnect()
	if err != nil {
		return protos.RegisterResponse{}, errs.NewGrpcError("获取grpc链接失败", err)
	}

	client := protos.NewUserClient(conn)
	response, err := client.Register(ctx, request)
	if err != nil {
		return protos.RegisterResponse{}, errs.NewGrpcError("用户注册rpc请求错误", err)
	}
	return *response, nil
}

//GetToken getToken
func (u *UserClient) GetToken(request *protos.TokenRequest) (protos.TokenResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := utils.OpenGrpcClientConnect()
	if err != nil {
		return protos.TokenResponse{}, errs.NewGrpcError("获取grpc链接失败", err)
	}

	client := protos.NewAuthClient(conn)
	response, err := client.GetToken(ctx, request)
	if err != nil {
		return protos.TokenResponse{}, errs.NewGrpcError("获取token失败", err)
	}

	return *response, nil
}
