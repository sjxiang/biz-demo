package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/sjxiang/biz-demo/book-shop/pkg/conf"
)

var DB *gorm.DB

// Init init DB
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
