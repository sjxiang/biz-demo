package handlers

import (
	
	"context"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/sjxiang/biz-demo/easy-note/cmd/api/rpc"
	"github.com/sjxiang/biz-demo/easy-note/pkg/errno"
	"github.com/sjxiang/biz-demo/easy-note/gen/pb"
)

// UpdateNote update user info
func UpdateNote(ctx *gin.Context) {
	var noteVar NoteParam
	if err := ctx.ShouldBind(&noteVar); err != nil {
		SendResponse(ctx, errno.ConvertErr(err), nil)
		return
	}

	uid := ctx.MustGet("userId").(int64)

	noteIDStr := ctx.Param("note_id")
	noteID, err := strconv.ParseInt(noteIDStr, 10, 64)
	if err != nil {
		SendResponse(ctx, errno.ConvertErr(err), nil)
		return
	}

	if noteID <= 0 {
		SendResponse(ctx, errno.ParamErr, nil)
		return
	}

	req := &pb.UpdateNoteRequest{NoteId: noteID, UserId: uid}
	
	if len(noteVar.Title) != 0 {
		req.Title = &noteVar.Title
	}
	if len(noteVar.Content) != 0 {
		req.Content = &noteVar.Content
	}
	if err = rpc.UpdateNote(context.Background(), req); err != nil {
		SendResponse(ctx, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(ctx, errno.Success, nil)
}
