package handlers

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/sjxiang/biz-demo/easy-note/cmd/api/rpc"
	"github.com/sjxiang/biz-demo/easy-note/pkg/errno"
	proto "github.com/sjxiang/biz-demo/easy-note/gen/user"
)

// Register register user info
func Register(ctx *gin.Context) {
	var registerVar UserParam
	if err := ctx.ShouldBind(&registerVar); err != nil {
		SendResponse(ctx, errno.ConvertErr(err), nil)
		return
	}

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendResponse(ctx, errno.ParamErr, nil)
		return
	}

	err := rpc.CreateUser(context.Background(), &proto.CreateUserRequest{
		UserName: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		SendResponse(ctx, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(ctx, errno.Success, nil)
}
