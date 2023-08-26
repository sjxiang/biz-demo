package service

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"

	"github.com/sjxiang/biz-demo/book-shop/app/user/dal/db"
	"github.com/sjxiang/biz-demo/book-shop/app/user/dal/redis"
	"github.com/sjxiang/biz-demo/book-shop/grpc_gen/pb"
	"github.com/sjxiang/biz-demo/book-shop/pkg/errno"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{
		ctx: ctx,
	}
}

func (s *UserService) CreateUser(req *pb.CreateUserRequest) error {
	users, err := db.QueryUser(s.ctx, req.UserName)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	
	return db.CreateUser(s.ctx, []*db.User{{
		UserName: req.UserName,
		Password: passWord,
	}})
}

// MGetUser using cache mode: Cache Aside
func (s *UserService) MGetUser(req *pb.MGetUserRequest) ([]*pb.User, error) {
	ret := make([]*pb.User, 0)
	idNotCached := make([]int64, 0)

	userInfoStr, err := redis.MGet(req.UserIds)
	// 降级
	if err != nil || userInfoStr == nil {
		idNotCached = req.UserIds
	} else {
		for index, item := range userInfoStr {
			if item == "" {
				idNotCached = append(idNotCached, req.UserIds[index])
			} else {
				ret = append(ret, s.getDtoFromString(item))
			}
		}
	}

	// 缺失的，从 db 补上
	users, err := db.MGetUsers(s.ctx, idNotCached)
	if err != nil {
		return nil, err
	}

	for _, userModel := range users {
		// 适配 pb
		userCur := &pb.User{
			UserId:   int64(userModel.ID),
			UserName: userModel.UserName,
		}
		ret = append(ret, userCur)

		// redis 补上缺失的
		str, _ := json.Marshal(userCur)
		_ = redis.Upsert(int64(userModel.ID), string(str))
	}

	return ret, nil
}

func (s *UserService) CheckUser(req *pb.CheckUserRequest) (int64, error) {
	// md5 加密
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.UserName
	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.UserNotExistErr
	}

	// 校验密码
	u := users[0]
	if u.Password != passWord {
		return 0, errno.LoginErr
	}
	return int64(u.ID), nil
}

func (s *UserService) getDtoFromString(userInfo string) *pb.User {
	ret := &pb.User{}
	_ = json.Unmarshal([]byte(userInfo), ret)
	return ret
}

