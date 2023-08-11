package main

import (
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"

	"github.com/sjxiang/biz-demo/easy-note/cmd/note/dal"
	"github.com/sjxiang/biz-demo/easy-note/cmd/note/rpc"
	"github.com/sjxiang/biz-demo/easy-note/pkg/consts"
	"github.com/sjxiang/biz-demo/easy-note/pkg/middleware"
	"github.com/sjxiang/biz-demo/easy-note/kitex_gen/note/noteservice"
)

func Init() {
	rpc.InitRPC()
	dal.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{consts.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", consts.NoteServiceAddr)
	if err != nil {
		panic(err)
	}
	Init()
	
	svr := noteservice.NewServer(new(NoteServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.NoteServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithRegistry(r),                                             // registry
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
