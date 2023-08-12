package main

import (
	"context"

	"github.com/sjxiang/biz-demo/easy-note/cmd/user/service"
	"github.com/sjxiang/biz-demo/easy-note/cmd/user/pack"
	pb "github.com/sjxiang/biz-demo/easy-note/gen/user"
	"github.com/sjxiang/biz-demo/easy-note/pkg/errno"
)

type UserServiceImpl struct{
	pb.UnimplementedUserServiceServer
}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (resp *pb.CreateUserResponse, err error) {
	resp = new(pb.CreateUserResponse)

	if len(req.UserName) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *pb.MGetUserRequest) (resp *pb.MGetUserResponse, err error) {
	resp = new(pb.MGetUserResponse)

	if len(req.UserIds) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	users, err := service.NewMGetUserService(ctx).MGetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Users = users
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *pb.CheckUserRequest) (resp *pb.CheckUserResponse, err error) {
	resp = new(pb.CheckUserResponse)

	if len(req.UserName) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserId = uid
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
