package pack

import (
	"errors"
	"time"

	"github.com/sjxiang/biz-demo/easy-note/kitex_gen/note"
	"github.com/sjxiang/biz-demo/easy-note/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *note.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *note.BaseResp {
	
	return &note.BaseResp{
		StatusCode:    err.ErrCode, 
		StatusMessage: err.ErrMsg, 
		ServiceTime:   time.Now().Unix(),
	}
}
