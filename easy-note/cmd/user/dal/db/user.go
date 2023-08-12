package db


import (
	"context"

	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

func (u *User) TableName() string {
	return "user"
}


// MGetUsers multiple get list of user info
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

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, username string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("user_name = ?", username).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func FindUserByNameOrEmail(username, email string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.Where("user_name = ?", username).Or("email = ?", email).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func CheckUser(username, password string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.Where("user_name = ? AND password = ?", username, password).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

