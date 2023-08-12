package service

import (
	"context"

	
	"github.com/sjxiang/biz-demo/easy-note/cmd/note/dal/db"
	pb "github.com/sjxiang/biz-demo/easy-note/gen/note"
)

type DelNoteService struct {
	ctx context.Context
}

func NewDelNoteService(ctx context.Context) *DelNoteService {
	return &DelNoteService{
		ctx: ctx,
	}
}

func (s *DelNoteService) DelNote(req *pb.DeleteNoteRequest) error {
	return db.DeleteNote(s.ctx, req.NoteId, req.UserId)
}
