package service

import (
	"context"

	"github.com/sjxiang/biz-demo/easy-note/cmd/note/dal/db"
	pb "github.com/sjxiang/biz-demo/easy-note/gen/note"
)

// 业务逻辑
type CreateNoteService struct {
	ctx context.Context
}

func NewCreateNoteService(ctx context.Context) *CreateNoteService {
	return &CreateNoteService{ctx: ctx}
}

func (s *CreateNoteService) CreateNote(req *pb.CreateNoteRequest) error {
	noteModel := &db.Note{
		UserID:  req.UserId,
		Title:   req.Title,
		Content: req.Content,
	}

	return db.CreateNote(s.ctx, []*db.Note{noteModel})
}
