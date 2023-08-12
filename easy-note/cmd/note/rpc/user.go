package rpc

import (
	"context"

	"google.golang.org/grpc"

	pb "github.com/sjxiang/biz-demo/easy-note/gen/user"
	"github.com/sjxiang/biz-demo/easy-note/pkg/consts"
	"github.com/sjxiang/biz-demo/easy-note/pkg/errno"
)

var userClient pb.UserServiceClient

func initUserRpc() {

	cc, err := grpc.Dial(consts.UserServiceAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	userClient = pb.NewUserServiceClient(cc)

	// r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
	// if err != nil {
	// 	panic(err)
	// }
// 	// cc, err := userservice.NewClient(
	// 	consts.UserServiceName,
	// 	client.WithMiddleware(middleware.CommonMiddleware),
	// 	client.WithMuxConnection(1),                       // mux
	// 	client.WithRPCTimeout(3*time.Second),              // rpc timeout
	// 	client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
	// 	client.WithFailureRetry(retry.NewFailurePolicy()), // retry
	// 	client.WithResolver(r),                            // resolver
	// )
}

// MGetUser multiple get list of user info
func MGetUser(ctx context.Context, req *pb.MGetUserRequest) (map[int64]*pb.User, error) {
	resp, err := userClient.MGetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	res := make(map[int64]*pb.User)
	for _, u := range resp.Users {
		res[u.UserId] = u
	}
	return res, nil
}
