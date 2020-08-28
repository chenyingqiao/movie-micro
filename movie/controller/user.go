package controller

import (
	"movie/rpc/client"
	"movie/rpc/protos"
	"movie/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	username := c.PostForm("username")
	password := c.PostForm("password")
	passwordRe := c.PostForm("password_re")

	userClient := client.NewUserClient()
	response, err := userClient.Register(&protos.RegisterRequest{
		Username:       username,
		Password:       password,
		PasswordRepeat: passwordRe,
	})
	if err != nil {
		c.JSON(http.StatusOK, utils.JSONResult("注册失败", gin.H{
			"status": false,
		}, 500))
		return
	}
	if response.Status {
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
	username := c.PostForm("username")
	password := c.PostForm("password")
	userClient := client.NewUserClient()
	response, err := userClient.GetToken(&protos.TokenRequest{
		Username: username,
		Password: password,
	})
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
