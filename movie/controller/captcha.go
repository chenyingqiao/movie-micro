package controller

import (
	"movie/logic"
	"movie/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//CaptchaController 验证码校验控制器
type CaptchaController struct {
	captchatLogic *logic.CaptchaLogic
}

//NewCaptchaController 创建验证码校验
func NewCaptchaController() *CaptchaController {
	return &CaptchaController{
		captchatLogic: logic.NewCaptchaLogic(),
	}
}

//Register 被动注册
func (cap *CaptchaController) Register(engin *gin.Engine) {
	engin.GET("/captcha", cap.get) //获取验证码
}

func (cap *CaptchaController) get(c *gin.Context) {
	id, base64Str, err := cap.captchatLogic.Generator()
	if err != nil {
		c.JSON(http.StatusOK, utils.JSONResult("验证码生成失败，请重试", gin.H{
			"status": false,
		}, 500))
		return
	}
	c.JSON(http.StatusOK, utils.JSONResult("success", gin.H{
		"id":     id,
		"base64": base64Str,
	}, 200))
	return
}
