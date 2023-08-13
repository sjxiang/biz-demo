package handlers

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/sjxiang/biz-demo/easy-note/cmd/api/rpc"
	"github.com/sjxiang/biz-demo/easy-note/pkg/errno"
	"github.com/sjxiang/biz-demo/easy-note/gen/pb"
)

// QueryNote query list of note info
func QueryNote(ctx *gin.Context) {
	uid := ctx.MustGet("userId").(int64)


	var queryVar struct {
		Limit         int64  `json:"limit" form:"limit" query:"limit"`
		Offset        int64  `json:"offset" form:"offset" query:"offset"`
		SearchKeyword string `json:"search_keyword" form:"search_keyword" query:"search_keyword"`
	}
	if err := ctx.ShouldBind(&queryVar); err != nil {
		SendResponse(ctx, errno.ConvertErr(err), nil)
	}

	if queryVar.Limit < 0 || queryVar.Offset < 0 {
		SendResponse(ctx, errno.ParamErr, nil)
		return
	}

	req := &pb.QueryNoteRequest{UserId: uid, Offset: queryVar.Offset, Limit: queryVar.Limit}
	if len(queryVar.SearchKeyword) != 0 {
		req.SearchKey = &queryVar.SearchKeyword
	}
	notes, total, err := rpc.QueryNotes(context.Background(), req)
	if err != nil {
		SendResponse(ctx, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(ctx, errno.Success, map[string]interface{}{"total": total, "notes": notes})
}
