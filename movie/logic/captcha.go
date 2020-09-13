package logic

import "github.com/mojocn/base64Captcha"

var (
	store  = base64Captcha.DefaultMemStore
	driver = base64Captcha.NewDriverAudio(6, "zh")
)

//CaptchaLogic 验证码校验逻辑
type CaptchaLogic struct {
}

//NewCaptchaLogic 获取实例
func NewCaptchaLogic() *CaptchaLogic {
	return &CaptchaLogic{}
}

//Generator 生成
func (capt *CaptchaLogic) Generator() (string, string, error) {
	captcha := base64Captcha.NewCaptcha(driver, store)
	return captcha.Generate()
}

//Verify 校验
func (capt *CaptchaLogic) Verify(id string, input string) bool {
	return store.Verify(id, input, true)
}
