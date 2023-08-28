package cache

import (
	"time"

	redigo "github.com/gomodule/redigo/redis"
	
	"github.com/sjxiang/biz-demo/book-shop/internal/conf"
)

var (
	pool *redigo.Pool
)

func Init(c *conf.Data) {
	pool = &redigo.Pool{
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", c.RedisAddr+""+c.RedisPort,
				redigo.DialConnectTimeout(200 * time.Millisecond),
				redigo.DialReadTimeout(200 * time.Millisecond),
				redigo.DialWriteTimeout(200 * time.Millisecond))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		MaxIdle: c.RedisConnPoolSize,
	}
}
