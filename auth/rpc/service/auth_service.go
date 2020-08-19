package service

import (
	"auth/rpc/protos"
	"context"
)

//AuthService 校验rpc服务
type AuthService struct {
	protos.UnimplementedAuthServer
}

// 获取token
func (auth *AuthService) GetToken(context.Context, *protos.TokenRequest) (*protos.TokenResponse, error) {
}

// 校验是否是登录状态
func (auth *AuthService) Validate(context.Context, *protos.ValidateRequest) (*protos.ValidateResponse, error) {

}

// 获取用户信息
func (auth *AuthService) GetInfo(context.Context, *protos.GetInfoRequest) (*protos.GetInfoResponse, error) {

}
