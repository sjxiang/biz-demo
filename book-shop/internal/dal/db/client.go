package db

import (
	"github.com/sjxiang/biz-demo/book-shop/pkg/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 连接
var (
	DB *gorm.DB
)

func Init() {
	var err error

	DB, err = gorm.Open(mysql.Open(conf.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
}

