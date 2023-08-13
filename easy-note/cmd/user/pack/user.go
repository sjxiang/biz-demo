package pack

import (
	

	"github.com/sjxiang/biz-demo/easy-note/cmd/user/dal/db"
	"github.com/sjxiang/biz-demo/easy-note/gen/pb"
)

// User pack user info
func User(u *db.User) *pb.User {
	if u == nil {
		return nil
	}

	return &pb.User{UserId: int64(u.ID), UserName: u.UserName, Avatar: "test"}
}

// Users pack list of user info
func Users(us []*db.User) []*pb.User {
	users := make([]*pb.User, 0)
	for _, u := range us {
		if user2 := User(u); user2 != nil {
			users = append(users, user2)
		}
	}
	return users
}
