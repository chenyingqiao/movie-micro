package utils

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	validate = validator.New()
	uni      = ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
)

//GetDefaultValidate 获取默认的中文校验器
func GetDefaultValidate() *validator.Validate {
	zh_translations.RegisterDefaultTranslations(validate, trans)
	return validate
}

//ValidateErrorsf 格式化错误信息
func ValidateErrorsf(errs validator.ValidationErrors) string {
	errStr := ""
	for _, v := range errs {
		errStr += "<br>\n" + v.Translate(trans)
	}
	return errStr
}
