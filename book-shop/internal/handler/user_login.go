package handler

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/sjxiang/biz-demo/book-shop/internal/service"
	"github.com/sjxiang/biz-demo/book-shop/internal/types"
	"github.com/sjxiang/biz-demo/book-shop/pkg/errno"
)

// 用户登录
func UserLogin(ctx *gin.Context) {
	// 请求校验 
	var loginParam types.UserParam
	if err := ctx.ShouldBind(&loginParam); err != nil {
		SendResponse(ctx, errno.BadRequestErr, nil)
	}
	if (len(loginParam.PassWord) == 0) || len(loginParam.PassWord) == 0 {
		SendResponse(ctx, errno.ParamErr, nil)
	}

	// 调用服务
	userId, err := service.NewUserService(context.Background()).CheckUser(&loginParam)	
	if err != nil {
		SendResponse(ctx, errno.ServiceErr, nil)
	}

	// jwt 生成 token

	
	// 返回响应
	SendResponse(ctx, errno.Success,  map[string]interface{}{"uid": userId, "token": ""})
}

	