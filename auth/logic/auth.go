package logic

//auth 鉴权
type auth interface {
	//GetToken 获取token
	GetToken() (string, error)
	//Validate 校验token是否有效
	Validate(token string) (bool, error)
	GetInfo(token string) (map[string]string, error)
}
