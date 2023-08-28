package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/biz-demo/easy-note/pkg/errno"
)

func AddProduct(ctx *gin.Context) {
	// 参数校验
	// 错误处理
	
	var pid int64 = 1
	SendResponse(ctx, errno.Success, map[string]interface{}{"product_id": strconv.FormatInt(pid, 10)})
}