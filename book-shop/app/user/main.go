package main

import (
	"net"

	"google.golang.org/grpc"

	"github.com/sjxiang/biz-demo/book-shop/app/user/dal"
	"github.com/sjxiang/biz-demo/book-shop/pkg/conf"
	"github.com/sjxiang/biz-demo/book-shop/grpc_gen/pb"
)

func Init() {
	// 数据访问层，初始化
	dal.Init()
}

func main() {


	// etcd 服务注册
	
	
	ln, err := net.Listen("tcp", conf.UserServiceAddress)
	if err != nil {
		panic(err)
	}

	Init()

	svr := grpc.NewServer()
	pb.RegisterUserServiceServer(svr, new(UserServiceImpl))

	if err := svr.Serve(ln); err != nil {
		panic(err)
	}
}