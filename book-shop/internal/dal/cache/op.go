package cache

import (
	"strconv"
	
	redigo "github.com/gomodule/redigo/redis"
	
	"github.com/sjxiang/biz-demo/book-shop/pkg/conf"
)

// 操作
func GetClient() redigo.Conn {
	return pool.Get()
}

func Upsert(userId int64, userInfo string) error {
	c := GetClient()
	defer c.Close()

	_, err := c.Do("SET", conf.RedisKey_User + strconv.FormatInt(userId, 10), userInfo) // SET k, v
	return err
}

func Del(userId int64) error {
	c := GetClient()
	defer c.Close()

	_, err := c.Do("DEL", conf.RedisKey_User + strconv.FormatInt(userId, 10))  // DEL k
	return err
}

func IsExist(userId int64) (bool, error) {
	c := GetClient()
	defer c.Close()

	isExist, err := redigo.Bool(c.Do("EXISTS", conf.RedisKey_User+strconv.FormatInt(userId, 10)))  // EXISTS k
	return isExist, err
}

func MGet(userIds []int64) ([]string, error) {
	c := GetClient()
	defer c.Close()

	keys := make([]interface{}, 0)
	for _, id := range userIds {
		keys = append(keys, conf.RedisKey_User+strconv.FormatInt(id, 10))
	}

	ret, err := redigo.Strings(c.Do("MGET", keys...))  // MGET k1 k2 k3  
	return ret, err
}
