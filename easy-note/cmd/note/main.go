package main

import (
	"net"

	"google.golang.org/grpc"

	"github.com/sjxiang/biz-demo/easy-note/cmd/note/dal"
	"github.com/sjxiang/biz-demo/easy-note/cmd/note/rpc"
	"github.com/sjxiang/biz-demo/easy-note/pkg/consts"
	"github.com/sjxiang/biz-demo/easy-note/gen/pb"
)

func Init() {
	rpc.InitRPC()
	dal.Init()
}

func main() {
	
	// etcd 服务注册
	
	ln, err := net.Listen("tcp", consts.NoteServiceAddr)
	if err != nil {
		panic(err)
	}
	
	Init()

	svr := grpc.NewServer()
	pb.RegisterNoteServiceServer(svr, new(NoteServiceImpl))
	
	if err := svr.Serve(ln); err != nil {
		panic(err)
	}
}
