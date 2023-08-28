package db

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return "t_user"
}


// MGetUsers - multiple get list of user info
func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// InsertUser - insert user info
func InsertUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// QueryUser -  query list of user info
func QueryUserByName(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return res, nil
}

