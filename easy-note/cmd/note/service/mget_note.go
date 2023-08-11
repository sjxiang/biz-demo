package service

import (
	"context"

	"github.com/sjxiang/biz-demo/easy-note/cmd/note/dal/db"
	"github.com/sjxiang/biz-demo/easy-note/cmd/note/pack"
	"github.com/sjxiang/biz-demo/easy-note/cmd/note/rpc"
	"github.com/sjxiang/biz-demo/easy-note/kitex_gen/note"
	"github.com/sjxiang/biz-demo/easy-note/kitex_gen/user"
)

type MGetNoteService struct {
	ctx context.Context
}

// NewMGetNoteService new MGetNoteService
func NewMGetNoteService(ctx context.Context) *MGetNoteService {
	return &MGetNoteService{ctx: ctx}
}

// MGetNote multiple get list of note info
func (s *MGetNoteService) MGetNote(req *note.MGetNoteRequest) ([]*note.Note, error) {
	// 获取笔记信息
	noteModels, err := db.MGetNotes(s.ctx, req.NoteIds)
	if err != nil {
		return nil, err
	}

	// 获取用户信息
	uIds := pack.UserIds(noteModels)
	userMap, err := rpc.MGetUser(s.ctx, &user.MGetUserRequest{UserIds: uIds})
	if err != nil {
		return nil, err
	}
	
	// 搞个大表不就行了，非要拼接 PO 组装 DTO
	notes := pack.Notes(noteModels)
	for i := 0; i < len(notes); i++ {
		if u, ok := userMap[notes[i].UserId]; ok {
			notes[i].UserName = u.UserName
			notes[i].UserAvatar = u.Avatar
		}
	}
	return notes, nil
}
