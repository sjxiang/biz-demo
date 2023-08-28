package data

import (
	"time"
	"github.com/google/wire"
	"github.com/sjxiang/biz-demo/book-shop/internal/conf"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	redigo "github.com/gomodule/redigo/redis"
)

var ProviderSet = wire.NewSet(NewData, NewDB)

// Data .
type Data struct {
	db   *gorm.DB
	pool *redigo.Pool
}

// NewData 构造方法，初始化了数据库 client
func NewData(c *conf.Data, logger *zap.SugaredLogger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		logger.Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

func NewDB(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.GetDatabaseDsn()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	InitDB(db)
	return db
}

func NewES(c *conf.Data) {

}

func NewCache(c *conf.Data) *redigo.Pool {
	pool :=  &redigo.Pool{
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", c.RedisAddr+ ":" + c.RedisPort,
				redigo.DialConnectTimeout(200*time.Millisecond),
				redigo.DialReadTimeout(200*time.Millisecond),
				redigo.DialWriteTimeout(200*time.Millisecond))
			if err != nil {
				panic(err)
			}

			return c, nil
		},
		MaxIdle: c.RedisConnPoolSize,
	}

	return pool
}

func InitDB(db *gorm.DB) {
	if err := db.AutoMigrate(
		// &User{},
		// &Article{},
		// &Comment{},
		// &ArticleFavorite{},
		// &Following{},
		); err != nil {
		panic(err)
	}
}

