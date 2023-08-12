package main

import (
	"net"

	"google.golang.org/grpc"

	"github.com/sjxiang/biz-demo/easy-note/cmd/note/dal"
	"github.com/sjxiang/biz-demo/easy-note/cmd/note/rpc"
	pb "github.com/sjxiang/biz-demo/easy-note/gen/note"
	"github.com/sjxiang/biz-demo/easy-note/pkg/consts"
)

func Init() {
	rpc.InitRPC()
	dal.Init()
}

func main() {
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
