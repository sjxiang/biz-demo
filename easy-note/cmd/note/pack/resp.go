package pack

import (
	"errors"
	"time"

	"github.com/sjxiang/biz-demo/easy-note/pkg/errno"
	"github.com/sjxiang/biz-demo/easy-note/gen/pb"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *pb.BaseResp {
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

func baseResp(err errno.ErrNo) *pb.BaseResp {
	
	return &pb.BaseResp{
		StatusCode:    err.ErrCode, 
		StatusMessage: err.ErrMsg, 
		ServiceTime:   time.Now().Unix(),
	}
}
