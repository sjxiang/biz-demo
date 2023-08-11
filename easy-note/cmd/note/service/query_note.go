package service

import (
	"context"

	
	"github.com/sjxiang/biz-demo/easy-note/cmd/note/dal/db"
	"github.com/sjxiang/biz-demo/easy-note/cmd/note/pack"
	"github.com/sjxiang/biz-demo/easy-note/cmd/note/rpc"
	"github.com/sjxiang/biz-demo/easy-note/kitex_gen/note"
	"github.com/sjxiang/biz-demo/easy-note/kitex_gen/user"
)

type QueryNoteService struct {
	ctx context.Context
}

// NewQueryNoteService new QueryNoteService
func NewQueryNoteService(ctx context.Context) *QueryNoteService {
	return &QueryNoteService{ctx: ctx}
}

// QueryNoteService query list of note info
func (s *QueryNoteService) QueryNoteService(req *note.QueryNoteRequest) ([]*note.Note, int64, error) {
	noteModels, total, err := db.QueryNote(s.ctx, req.UserId, req.SearchKey, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, 0, err
	}
	userMap, err := rpc.MGetUser(s.ctx, &user.MGetUserRequest{UserIds: []int64{req.UserId}})
	if err != nil {
		return nil, 0, err
	}
	notes := pack.Notes(noteModels)
	for i := 0; i < len(notes); i++ {
		if u, ok := userMap[notes[i].UserId]; ok {
			notes[i].UserName = u.UserName
			notes[i].UserAvatar = u.Avatar
		}
	}
	return notes, total, nil
}
