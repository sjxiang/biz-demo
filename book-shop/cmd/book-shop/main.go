package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sjxiang/biz-demo/book-shop/internal/conf"
)

func main() {
	conf.Load()
}


func RegisterRouters(engine *gin.Engine) {
	// // user service
	// userGroup := h.Group("/user")
	// userGroup.POST("/register", handler_user.UserRegister)
	// userGroup.POST("/login", handler_user.UserLogin)

	// // shop service
	// shopGroup := h.Group("/shop")
	// shopGroup.POST("/login", handler_user.ShopLogin)

	// // item-2b service
	// item2BGroup := h.Group("/item2b")
	// item2BGroup.Use(model.ShopAuthMiddleware.MiddlewareFunc())
	// item2BGroup.POST("/add", handler_item.AddProduct)
	// item2BGroup.POST("/edit", handler_item.EditProduct)
	// item2BGroup.POST("/del", handler_item.DelProduct)
	// item2BGroup.POST("/offline", handler_item.OfflineProduct)
	// item2BGroup.POST("/online", handler_item.OnlineProduct)
	// item2BGroup.GET("/get", handler_item.GetProduct)
	// item2BGroup.POST("/list", handler_item.ListProduct)

	// // item-2c service
	// item2CGroup := h.Group("/item2c")
	// item2CGroup.Use(model.UserAuthMiddleware.MiddlewareFunc())
	// item2CGroup.GET("/mget", handler_item.MGetProduct2C)
	// item2CGroup.POST("/search", handler_item.SearchProduct)

	// // order service
	// orderGroup := h.Group("/order")
	// orderGroup.Use(model.UserAuthMiddleware.MiddlewareFunc())
	// orderGroup.POST("/create", handler_order.CreateOrder)
	// orderGroup.POST("/cancel", handler_order.CancelOrder)
	// orderGroup.POST("/list", handler_order.ListOrder)
	// orderGroup.GET("/get", handler_order.GetOrder)
}

