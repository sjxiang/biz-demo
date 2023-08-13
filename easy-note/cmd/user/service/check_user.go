package service


import (
	"context"
	
	"golang.org/x/crypto/bcrypt"
	
	"github.com/sjxiang/biz-demo/easy-note/cmd/user/dal/db"
	"github.com/sjxiang/biz-demo/easy-note/gen/pb"
	"github.com/sjxiang/biz-demo/easy-note/pkg/errno"
)

type CheckUserService struct {
	ctx context.Context
}

func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{
		ctx: ctx,
	}
}

func (s *CheckUserService) CheckUser(req *pb.CheckUserRequest) (int64, error) {

	var (
		userName = req.UserName
		password = req.Password
	)

	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.AuthorizationFailedErr
	}
	
	u := users[0]
	// 校验密码
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return 0, errno.AuthorizationFailedErr
	}

	return int64(u.ID), nil
}
