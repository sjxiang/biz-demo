package main

import (

	"net"

	"google.golang.org/grpc"

	"github.com/sjxiang/biz-demo/easy-note/cmd/user/dal"
	pb "github.com/sjxiang/biz-demo/easy-note/gen/user"
	"github.com/sjxiang/biz-demo/easy-note/pkg/consts"
)


func Init() {
	dal.Init()
}

func main() {
	ln, err := net.Listen("tcp", consts.UserServiceAddr)
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