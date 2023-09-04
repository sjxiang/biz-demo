package dal

import (
	"github.com/sjxiang/biz-demo/tiktok_demo/biz/dal/db"
	"github.com/sjxiang/biz-demo/tiktok_demo/biz/mw/redis"
)

// Init init dal
func Init() {
	db.Init() // mysql init
	redis.InitRedis()
}
