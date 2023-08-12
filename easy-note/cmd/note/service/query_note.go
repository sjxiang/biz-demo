package service

import (
	"context"

	
	"github.com/sjxiang/biz-demo/easy-note/cmd/note/dal/db"
	"github.com/sjxiang/biz-demo/easy-note/cmd/note/pack"
	"github.com/sjxiang/biz-demo/easy-note/cmd/note/rpc"
	pb "github.com/sjxiang/biz-demo/easy-note/gen/note"
	proto "github.com/sjxiang/biz-demo/easy-note/gen/user"
)

type QueryNoteService struct {
	ctx context.Context
}

func NewQueryNoteService(ctx context.Context) *QueryNoteService {
	return &QueryNoteService{ctx: ctx}
}

func (s *QueryNoteService) QueryNoteService(req *pb.QueryNoteRequest) ([]*pb.Note, int64, error) {
	// 查询笔记信息
	noteModels, total, err := db.QueryNote(s.ctx, req.UserId, req.SearchKey, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, 0, err
	}

	// 查询用户信息
	userMap, err := rpc.MGetUser(s.ctx, &proto.MGetUserRequest{UserIds: []int64{req.UserId}})
	if err != nil {
		return nil, 0, err
	}

	// po 拼装 dto
	notes := pack.Notes(noteModels)
	for i := 0; i < len(notes); i++ {
		if u, ok := userMap[notes[i].UserId]; ok {
			notes[i].UserName = u.UserName
			notes[i].UserAvatar = u.Avatar
		}
	}
	return notes, total, nil
}
