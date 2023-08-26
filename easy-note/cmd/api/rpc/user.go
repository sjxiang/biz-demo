package rpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/sjxiang/biz-demo/easy-note/gen/pb"
	"github.com/sjxiang/biz-demo/easy-note/pkg/consts"
	"github.com/sjxiang/biz-demo/easy-note/pkg/errno"
)

var userClient pb.UserServiceClient

func initUserRpc() {

	// etcd 服务发现

	c, err := grpc.Dial(consts.UserServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	userClient = pb.NewUserServiceClient(c)

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

// CreateUser create user info
func CreateUser(ctx context.Context, req *pb.CreateUserRequest) error {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// CheckUser check user info
func CheckUser(ctx context.Context, req *pb.CheckUserRequest) (int64, error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.UserId, nil
}
