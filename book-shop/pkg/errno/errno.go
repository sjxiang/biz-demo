
package errno

import (
	"errors"
	"fmt"
)

const (
	// System Code
	SuccessCode       = 0
	ServiceErrCode    = 10001
	ParamErrCode      = 10002
	BadRequestErrCode = 10003

	// User ErrCode
	LoginErrCode            = 11001
	UserNotExistErrCode     = 11002
	UserAlreadyExistErrCode = 11003
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success             = NewErrNo(SuccessCode, "Success")
	ServiceErr          = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr            = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	BadRequestErr       = NewErrNo(BadRequestErrCode, "请求解析错误")
	LoginErr            = NewErrNo(LoginErrCode, "Wrong username or password")
	UserNotExistErr     = NewErrNo(UserNotExistErrCode, "User does not exists")
	UserAlreadyExistErr = NewErrNo(UserAlreadyExistErrCode, "User already exists")
)

// 错误统一处理
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {  // 用于检查 error chain 中的每个错误是否可以转换为特定类型
		return Err
	}

	// 默认，系统内部错误，常量
	s := ServiceErr  
	s.ErrMsg = err.Error()
	return s
}
