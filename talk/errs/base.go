package errs

//BaseErr 基础错误
type BaseErr struct {
	code    int
	message string
}

//GetMessage 基础错误
func (b *BaseErr) GetMessage() string {
	return b.message
}

//SetMessage 基础错误
func (b *BaseErr) SetMessage(message string) {
	b.message = message
}

//NewBaseErr 创建baseErr
func NewBaseErr(code int, message string) BaseErr {
	return BaseErr{
		code:    code,
		message: message,
	}
}
