package handlers

import (	
	"context"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/sjxiang/biz-demo/easy-note/cmd/api/rpc"
	"github.com/sjxiang/biz-demo/easy-note/pkg/errno"
	pb "github.com/sjxiang/biz-demo/easy-note/gen/note"

)

// DeleteNote delete note info
func DeleteNote(ctx *gin.Context) {
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

	err = rpc.DeleteNote(context.Background(), &pb.DeleteNoteRequest{
		NoteId: noteID, UserId: uid,
	})
	if err != nil {
		SendResponse(ctx, errno.ConvertErr(err), nil)
		return
	}

	SendResponse(ctx, errno.Success, nil)
}
