package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sjxiang/biz-demo/book-shop/pkg/errno"
)


// serializer，序列化器

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`  
	Data    interface{} `json:"data,omitempty"`
}

// SendResponse pack response
func SendResponse(ctx *gin.Context, err error, data interface{}) {
	Err := errno.ConvertErr(err)

	// 缺陷，HTTP 状态码
	ctx.JSON(http.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}









