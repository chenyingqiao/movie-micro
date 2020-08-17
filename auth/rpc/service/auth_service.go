package service

import (
	"auth/rpc/protos"
)

//AuthService 校验rpc服务
type AuthService struct {
	protos.UnimplementedAuthServer
}

// //GetToken 获取token
// func (auth *AuthService) GetToken(context.Context, *protos.TokenRequest) (*protos.TokenRequest, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
// }

// //Validate 校验
// func (auth *AuthService) Validate(context.Context, *protos.ValidateRequest) (*protos.ValidateResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
// }

// //GetInfo 获取用户信息
// func (auth *AuthService) GetInfo(context.Context, *protos.GetInfoRequest) (*protos.GetInfoResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
// }
