package main

import (
	"context"

	"github.com/sjxiang/biz-demo/book-shop/app/user/service"
	"github.com/sjxiang/biz-demo/book-shop/grpc_gen/pb"
	"github.com/sjxiang/biz-demo/book-shop/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{
	pb.UnimplementedUserServiceServer
}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (resp *pb.CreateUserResponse, err error) {

	resp = new(pb.CreateUserResponse)

	if len(req.GetUserName()) == 0 || len(req.GetPassword()) == 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewUserService(context.Background()).CreateUser(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return resp, nil
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *pb.MGetUserRequest) (resp *pb.MGetUserResponse, err error) {
	resp = new(pb.MGetUserResponse)

	if len(req.UserIds) == 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	users, err := service.NewUserService(ctx).MGetUser(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}

	resp.Users = users
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *pb.CheckUserRequest) (resp *pb.CheckUserResponse, err error) {
	resp = new(pb.CheckUserResponse)

	if len(req.Password) == 0 || len(req.UserName) == 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	userId, err := service.NewUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}

	resp.UserId = userId
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return resp, nil
}
