package handlers

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/sjxiang/biz-demo/easy-note/cmd/api/rpc"
	"github.com/sjxiang/biz-demo/easy-note/pkg/errno"
	pb "github.com/sjxiang/biz-demo/easy-note/gen/note"
)

// CreateNote create note info
func CreateNote(ctx *gin.Context) {
	var noteVar NoteParam
	if err := ctx.ShouldBind(&noteVar); err != nil {
		SendResponse(ctx, errno.ConvertErr(err), nil)
		return
	}

	if len(noteVar.Title) == 0 || len(noteVar.Content) == 0 {
		SendResponse(ctx, errno.ParamErr, nil)
		return
	}
	
	uid := ctx.MustGet("userId").(int64)
	
	err := rpc.CreateNote(context.Background(), &pb.CreateNoteRequest{
		UserId:  uid,
		Content: noteVar.Content, 
		Title:   noteVar.Title,
	})
	if err != nil {
		SendResponse(ctx, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(ctx, errno.Success, nil)
}
