package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Recovery 使用 zap.Error() 来记录 Panic 和 call stack
func Recovery(logger *zap.Logger) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		
		defer func ()  {
			if err := recover(); err != nil {

				// 获取用户的请求信息
				httpRequest, _ := httputil.DumpRequest(ctx.Request, true)

				// 连接中断，如果客户端中断连接（tcp）为正常行为，不需要记录堆栈信息
				var brokenPipe bool
				if ne, ok := err.(net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						errStr := strings.ToLower(se.Error())
						if strings.Contains(errStr, "broken pipe") || strings.Contains(errStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				// 连接中断的情况
				if brokenPipe {
					logger.Error(ctx.Request.URL.Path, 
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)

					ctx.Error(err.(error))
					ctx.Abort()

					// 连接已断开，无法写入状态
					return
				}

				// 如果不是连接中断，就开始记录堆栈信息
				logger.Error("recovery from panic",
					zap.Time("time", time.Now()),                 // 时间
					zap.Any("error", err),                        // 错误信息 
					zap.String("request", string(httpRequest)),   // 请求信息
					zap.Stack("stacktrace"),                      // 调用堆栈信息
				)

				// 返回 500 状态码
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "服务器内部错误，请稍后再试。",
				})
			}	
		}()
		
		ctx.Next()

	}
}