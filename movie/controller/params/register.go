package params

//Register 注册参数
type Register struct {
	Username   string `form:"username" validate:"required,max=15,min=5"`
	Password   string `form:"password" validate:"required,max=15,min=5,alphanumunicode" `
	PasswordRe string `form:"password_re" validate:"required,max=15,min=5,alphanumunicode"`
	Answer     string `form:"answer" validate:"required"`
	Capid      string `form:"capid" validate:"required"`
}
