package errno

import (
	"errors"
	"time"

	"github.com/sjxiang/biz-demo/book-shop/grpc_gen/pb"
)

func BuildBaseResp(err error) *pb.BaseResp {
	if err == nil {
		return baseResp(Success)
	}

	e := ErrNo{}
	if errors.As(err, &e) {  // 用于检查 error chain 中的每个错误是否可以转换为特定类型
		return baseResp(e)
	}

	// 常量，默认统一处理
	s := ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err ErrNo) *pb.BaseResp {
	return &pb.BaseResp{
		StatusCode:    int64(err.ErrCode),
		StatusMessage: err.ErrMsg,
		ServiceTime:   time.Now().Unix(), 
	}
}
