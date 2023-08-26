package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/biz-demo/easy-note/pkg/errno"
	"github.com/sjxiang/biz-demo/easy-note/pkg/jwt"
	"github.com/sjxiang/biz-demo/easy-note/cmd/api/rpc"
	"github.com/sjxiang/biz-demo/easy-note/gen/pb"
)


func Login(ctx *gin.Context) {
	var loginVar UserParam
	if err := ctx.ShouldBind(&loginVar); err != nil {
		SendResponse(ctx, errno.ConvertErr(err), nil)
	}
	if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
		SendResponse(ctx, errno.ParamErr, nil)
		return
	}

	uid , err := rpc.CheckUser(context.Background(), &pb.CheckUserRequest{
		UserName: loginVar.UserName, 
		Password: loginVar.PassWord,
	})
	if err != nil {
		SendResponse(ctx, errno.ConvertErr(err), nil)
		return
	}

	// 生成 jwt
	token, err := jwt.GenerateAuth2Token(uid)
	if err != nil {
		SendResponse(ctx, errno.ConvertErr(err), nil)
		return
	}

	SendResponse(ctx, errno.Success, map[string]interface{}{"token":  token})
}
