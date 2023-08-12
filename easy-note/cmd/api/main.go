package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/sjxiang/biz-demo/easy-note/cmd/api/handlers"
	"github.com/sjxiang/biz-demo/easy-note/cmd/api/middleware"
	"github.com/sjxiang/biz-demo/easy-note/cmd/api/rpc"
	"github.com/sjxiang/biz-demo/easy-note/pkg/consts"
)

func Init() {
	rpc.InitRPC()
}

func main() {
	Init()

	r := gin.New()

	// 绑定中间件 recovery
	r.Use(middleware.Recovery(zap.L()))

	// 路由分组
	v1 := r.Group("/v1")
	user1 := v1.Group("/user")
	user1.POST("/login", handlers.Login)
	user1.POST("/register", handlers.Register)

	note1 := v1.Group("/note")
	note1.Use(middleware.Auth())  // jwt
	note1.GET("/query", handlers.QueryNote)
	note1.POST("", handlers.CreateNote)
	note1.PUT("/:note_id", handlers.UpdateNote)
	note1.DELETE("/:note_id", handlers.DeleteNote)

	r.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "no route")
	})
	r.NoMethod(func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "no method")
	})

	r.Run(consts.ApiServiceAddr)
}