package params

//Login 登录参数
type Login struct {
	Username string `form:"username" validate:"required,max=15,min=5"`
	Password string `form:"password" validate:"required,max=15,min=5,alphanumunicode"`
	Answer   string `form:"answer" validate:"required"`
	Capid    string `form:"capid" validate:"required"`
}
