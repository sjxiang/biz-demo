package service

import (
	"context"

	
	"github.com/sjxiang/biz-demo/easy-note/cmd/note/dal/db"
	"github.com/sjxiang/biz-demo/easy-note/kitex_gen/note"
)

type DelNoteService struct {
	ctx context.Context
}

// NewDelNoteService new DelNoteService
func NewDelNoteService(ctx context.Context) *DelNoteService {
	return &DelNoteService{
		ctx: ctx,
	}
}

// DelNote delete note info
func (s *DelNoteService) DelNote(req *note.DeleteNoteRequest) error {
	return db.DeleteNote(s.ctx, req.NoteId, req.UserId)
}
