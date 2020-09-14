package controller

import (
	"movie/controller/params"
	"movie/logic"
	"movie/rpc/client"
	"movie/rpc/protos"
	"movie/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	captchaLogic = logic.NewCaptchaLogic()
	validate     = utils.GetDefaultValidate()
)

//UserCotnroller 用户控制器
type UserCotnroller struct{}

//NewUserCotroller NewUserCotroller
func NewUserCotroller() *UserCotnroller {
	return &UserCotnroller{}
}

//Register 注册路由
func (u *UserCotnroller) Register(engin *gin.Engine) {
	engin.POST("/register", u.reg)
	engin.POST("/login", u.login)
}

//Reg 注册新的用户
func (u *UserCotnroller) reg(c *gin.Context) {
	//参数校验
	registerParam := params.Register{}
	if err := c.ShouldBind(&registerParam); err != nil {
		c.JSON(http.StatusOK, utils.JSONResult("参数错误，请确认", gin.H{
			"status": false,
		}, 500))
		return
	}
	if err := validate.Struct(registerParam); err != nil {
		errs := err.(validator.ValidationErrors)
		errStr := utils.ValidateErrorsf(errs)
		c.JSON(http.StatusOK, utils.JSONResult(errStr, gin.H{
			"status": false,
		}, 500))
		return
	}

	//验证码校验
	capVerify := captchaLogic.Verify(registerParam.Capid, registerParam.Answer)
	if !capVerify {
		c.JSON(http.StatusOK, utils.JSONResult("验证码校验错误", gin.H{
			"status": false,
		}, 500))
		return
	}

	//调用第三方服务注册
	userClient := client.NewUserClient()
	response, err := userClient.Register(&protos.RegisterRequest{
		Username:       registerParam.Username,
		Password:       registerParam.Password,
		PasswordRepeat: registerParam.PasswordRe,
	})
	if err != nil {
		c.JSON(http.StatusOK, utils.JSONResult("注册失败", gin.H{
			"status": false,
		}, 500))
		return
	}
	if !response.Status {
		c.JSON(http.StatusOK, utils.JSONResult("注册失败", gin.H{
			"status": false,
		}, 500))
		return
	}
	c.JSON(http.StatusOK, utils.JSONResult("注册成功", gin.H{
		"status": true,
	}, 200))
	return
}

func (u *UserCotnroller) login(c *gin.Context) {
	loginParam := params.Login{}

	if err := c.ShouldBind(&loginParam); err != nil {
		c.JSON(http.StatusOK, utils.JSONResult("参数错误，请确认", gin.H{
			"token": "",
		}, 500))
		return
	}

	if err := validate.Struct(loginParam); err != nil {
		errs := err.(validator.ValidationErrors)
		errStr := utils.ValidateErrorsf(errs)
		c.JSON(http.StatusOK, utils.JSONResult(errStr, gin.H{
			"token": "",
		}, 500))
		return
	}

	userClient := client.NewUserClient()
	response, err := userClient.GetToken(&protos.TokenRequest{
		Username: loginParam.Username,
		Password: loginParam.Password,
	})
	capVerify := captchaLogic.Verify(loginParam.Capid, loginParam.Answer)
	if !capVerify {
		c.JSON(http.StatusOK, utils.JSONResult("验证码校验错误", gin.H{
			"token": "",
		}, 500))
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, utils.JSONResult("登录失败，请重试", gin.H{
			"token": "",
		}, 500))
		return
	}
	c.JSON(http.StatusOK, utils.JSONResult("注册成功", gin.H{
		"token": response.Token,
	}, 200))
	return
}
