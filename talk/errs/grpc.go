package errs

import "github.com/pkg/errors"

//GrpcError jwt过期
type GrpcError struct {
	error
}

//NewGrpcError 生成错误
func NewGrpcError(message string, err error) error {
	if err == nil {
		err = errors.New("GrpcError")
	}
	return errors.Wrap(GrpcError{
		error: err,
	}, "wrap")
}

//ValidataGrpcErrorError 校验错误
func ValidataGrpcErrorError(err error) bool {
	if err == nil {
		return false
	}
	switch errors.Cause(err).(type) {
	case GrpcError:
		return true
	default:
		return false
	}
}
