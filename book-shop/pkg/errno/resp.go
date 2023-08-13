package errno

import (
	"errors"

	"github.com/sjxiang/biz-demo/book-shop/grpc_gen/pb"
)

func BuildBaseResp(err error) *pb.BaseResp {
	if err == nil {
		return baseResp(Success)
	}

	e := ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err ErrNo) *pb.BaseResp {
	return &pb.BaseResp{
		StatusCode:    int64(err.ErrCode),
		StatusMessage: err.ErrMsg}
}
