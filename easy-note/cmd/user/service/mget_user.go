package service

import (
	"context"

	"github.com/sjxiang/biz-demo/easy-note/cmd/user/dal/db"
	"github.com/sjxiang/biz-demo/easy-note/cmd/user/pack"
	"github.com/sjxiang/biz-demo/easy-note/gen/pb"
)

type MGetUserService struct {
	ctx context.Context
}

func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

func (s *MGetUserService) MGetUser(req *pb.MGetUserRequest) ([]*pb.User, error) {
	modelUsers, err := db.MGetUsers(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	
	return pack.Users(modelUsers), nil
}
