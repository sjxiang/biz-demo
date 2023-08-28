package conf


// 常量配置

const (
	UserTableName    = "t_user"
	ProductTableName = "t_product"
	OrderTableName   = "t_order"

	SecretKey   = "secret key"
	IdentityKey = "id"

	ShopLoginName     = "admin"
	ShopLoginPassword = "123"

	MySQLDefaultDSN = "root:123456@tcp(localhost:3306)/book-shop?charset=utf8&parseTime=True&loc=Local"
	ESAddress       = "http://localhost:9200"

	RedisAddress        = "127.0.0.1:6379"
	RedisConnPoolSize   = 20
	RedisKey_User       = "user-"

	ProductESIndex = "product"
)
