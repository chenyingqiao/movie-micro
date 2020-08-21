package errs

import "github.com/pkg/errors"

//JwtExpiredError jwt过期
type JwtExpiredError struct {
	error
	BaseErr
}

//NewJwtExpiredError 生成错误
func NewJwtExpiredError(code int, message string, err error) error {
	if err == nil {
		err = errors.New("NewJwtExpiredError")
	}
	return errors.Wrap(error(JwtExpiredError{
		err,
		NewBaseErr(code, message),
	}), "wrap")
}

//ValidataJwtExpiredError 校验错误
func ValidataJwtExpiredError(err error) bool {
	switch errors.Cause(err).(type) {
	case JwtExpiredError:
		return true
	default:
		return false
	}
}
