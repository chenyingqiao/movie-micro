package errs

import "github.com/pkg/errors"

//DbError jwt过期
type DbError struct {
	error
}

//NewDbError 生成错误
func NewDbError(message string, err error) error {
	if err == nil {
		err = errors.New("DbError")
	}
	return errors.Wrap(DbError{
		error: err,
	}, "wrap")
}

//ValidataErrorError 校验错误
func ValidataErrorError(err error) bool {
	if err == nil {
		return false
	}
	switch errors.Cause(err).(type) {
	case DbError:
		return true
	default:
		return false
	}
}
