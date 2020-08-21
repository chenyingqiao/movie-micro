package errs

import "github.com/pkg/errors"

//UserOperatorError jwt过期
type UserOperatorError struct {
	error
	BaseErr
}

//NewUserOperatorError 生成错误
func NewUserOperatorError(code int, message string, err error) error {
	if err == nil {
		err = errors.New("NewUserOperatorError")
	}
	return errors.Wrap(error(UserOperatorError{
		err,
		NewBaseErr(code, message),
	}), "wrap")
}

//ValidataUserOperatorErrorError 校验错误
func ValidataUserOperatorErrorError(err error) bool {
	switch errors.Cause(err).(type) {
	case UserOperatorError:
		return true
	default:
		return false
	}
}
