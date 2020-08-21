package logic

import (
	"auth/errs"
	"auth/model"
	"auth/rpc/protos"

	"golang.org/x/crypto/bcrypt"
)

//User 用户逻辑
type User struct{}

//NewUser newUser
func NewUser() User {
	return User{}
}

//Register 注册用户
func (u *User) Register(user *protos.RegisterRequest) error {
	if user.Password != user.PasswordRepeat {
		return errs.NewUserOperatorError(errs.RegisterPasswordNotMatchCode, "两次输入的密码不匹配！", nil)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	userModel := model.NewUser(user.Username, string(password))
	err = userModel.Add()
	if err != nil {
		return errs.NewUserOperatorError(errs.DbErrorCode, "用户信息注册失败", err)
	}
	return nil
}
