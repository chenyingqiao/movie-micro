package client

import (
	"context"
	"talk/errs"
	"talk/rpc/protos"
	"talk/utils"
	"time"
)

//AuthClient AuthClient
type AuthClient struct {
	protos.UnimplementedAuthServer
}

//NewAuthClient NewAuthClient
func NewAuthClient() AuthClient {
	return AuthClient{}
}

//GetToken 获取token
func (a *AuthClient) GetToken(tokenRequest *protos.TokenRequest) (*protos.TokenResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := utils.OpenGrpcClientConnect(utils.AuthGrpcAddress)
	if err != nil {
		return &protos.TokenResponse{}, errs.NewGrpcError("获取grpc链接失败", err)
	}

	authClient := protos.NewAuthClient(conn)
	return authClient.GetToken(ctx, tokenRequest)
}

//校验token
func (a *AuthClient) Validate(request *protos.ValidateRequest) (*protos.ValidateResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := utils.OpenGrpcClientConnect(utils.AuthGrpcAddress)
	if err != nil {
		return &protos.ValidateResponse{}, errs.NewGrpcError("获取grpc链接失败", err)
	}

	client := protos.NewAuthClient(conn)
	return client.Validate(ctx, request)
}

//GetInfo 获取用户信息
func (a *AuthClient) GetInfo(request *protos.GetInfoRequest) (*protos.GetInfoResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := utils.OpenGrpcClientConnect(utils.AuthGrpcAddress)
	if err != nil {
		return &protos.GetInfoResponse{}, errs.NewGrpcError("获取grpc链接失败", err)
	}

	return protos.NewAuthClient(conn).GetInfo(ctx, request)
}
