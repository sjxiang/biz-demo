package dal

import (
	"github.com/sjxiang/biz-demo/book-shop/app/user/dal/db"
	"github.com/sjxiang/biz-demo/book-shop/app/user/dal/redis"
)

func Init() {
	db.Init()
	redis.Init()
}