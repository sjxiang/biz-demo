package main

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/biz-demo/tiktok_demo/biz/handler"
	"github.com/sjxiang/biz-demo/tiktok_demo/biz/handler/comment"
	"github.com/sjxiang/biz-demo/tiktok_demo/biz/handler/favorite"
)

func Init() {
	// dal.Init()
	// jwt.Init()
	// minio.Init()
}

func main() {
	Init()

	r := gin.New()

	register(r)

	r.Run(":18005")
}


// 设置 /src/*name 路由转发，以便从外部网络访问 minio
func minioReverseProxy(c *gin.Context) {
	
	// 创建 minio 的代理地址
	url, _ := url.Parse("http://localhost:18001")
	proxy := httputil.NewSingleHostReverseProxy(url)

	// 重写请求路径
	c.Request.URL.Path = c.Param("name")

	// 转发请求到代理
	proxy.ServeHTTP(c.Writer, c.Request)
}


// register 注册业务路由
func register(r *gin.Engine) {

	r.GET("/ping", handler.Ping)
	r.GET("/src/*name", minioReverseProxy)

	
	douyin := r.Group("/douyin")
	comment := douyin.Group("/comment")
	// jwt
	{
		action := comment.Group("/action")
		action.POST("/", comment.CommentAction)
	}
	// jwt
	{
		list := comment.Group("/list")
		list.GET("/", comment.CommentList)
	}
	
	favorite := douyin.Group("/favorite")
	// jwt
	{
		action := favorite.Group("/action")
		action.POST("/", favorite.CommentAction)
	}
	// jwt
	{
		list := favorite.Group("/list")
		list.GET("/", favorite.CommentList)
	}
	
	
	
}
