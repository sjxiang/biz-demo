package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	
	"github.com/sjxiang/biz-demo/easy-note/pkg/errno"
)

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse pack response
func SendResponse(ctx *gin.Context, err error, data interface{}) {
	Err := errno.ConvertErr(err)

	ctx.JSON(http.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}

type NoteParam struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}
