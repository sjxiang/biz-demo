package service

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/sjxiang/biz-demo/easy-note/cmd/user/dal/db"
	"github.com/sjxiang/biz-demo/easy-note/gen/pb"
	"github.com/sjxiang/biz-demo/easy-note/pkg/errno"
)

type CreateUserService struct {
	ctx context.Context
}

func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

func (s *CreateUserService) CreateUser(req *pb.CreateUserRequest) error {
	users, err := db.QueryUser(s.ctx, req.UserName)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	// 加密
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return db.CreateUser(s.ctx, []*db.User{{
		UserName: req.UserName,
		Password: string(hash),
	}})
}
