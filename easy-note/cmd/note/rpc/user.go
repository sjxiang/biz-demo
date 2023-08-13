package rpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/sjxiang/biz-demo/easy-note/pkg/consts"
	"github.com/sjxiang/biz-demo/easy-note/pkg/errno"
	"github.com/sjxiang/biz-demo/easy-note/gen/pb"
)

var userClient pb.UserServiceClient

func initUserRpc() {

	// etcd

	cc, err := grpc.Dial(consts.UserServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	userClient = pb.NewUserServiceClient(cc)

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
