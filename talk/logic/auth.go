package logic

import (
	"talk/errs"
	"talk/rpc/client"
	"talk/rpc/protos"
)

type AuthLogic struct{}

//NewAuthLogic NewAuthLogic
func NewAuthLogic() AuthLogic {
	return AuthLogic{}
}

//GetUserInfo 获取用户信息
func (auth *AuthLogic) GetUserInfo(jwt string) (protos.GetInfoResponse, error) {
	client := client.NewAuthClient()
	response, err := client.GetInfo(&protos.GetInfoRequest{
		Token: jwt,
	})
	if err != nil {
		return protos.GetInfoResponse{}, errs.NewGrpcError("获取用户信息失败", err)
	}
	return *response, nil
}
