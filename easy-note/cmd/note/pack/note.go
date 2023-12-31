package pack

import (
	"github.com/sjxiang/biz-demo/easy-note/cmd/note/dal/db"
	"github.com/sjxiang/biz-demo/easy-note/gen/pb"
)

// 序列化

// po 转 dto
func Note(m *db.Note) *pb.Note {
	
	// 判空处理
	if m == nil {
		return nil
	}

	return &pb.Note{
		NoteId:     int64(m.ID),
		UserId:     m.UserID,
		Title:      m.Title,
		Content:    m.Content,
		CreateTime: m.CreatedAt.Unix(),
	}
}

// Notes pack list of note info
func Notes(ms []*db.Note) []*pb.Note {
	notes := make([]*pb.Note, 0, len(ms))
	for _, m := range ms {
		if n := Note(m); n != nil {
			notes = append(notes, n)
		}
	}
	return notes
}


func UserIds(ms []*db.Note) []int64 {

	// var uIds = make([]int64, 0)

	uIds := make([]int64, 0)
	if len(ms) == 0 {
		return uIds
	}

	// map 去重
	uIdMap := make(map[int64]struct{})
	for _, m := range ms {
		if m != nil {
			uIdMap[m.UserID] = struct{}{}
		}
	}
	for uId := range uIdMap {
		uIds = append(uIds, uId)
	}
	return uIds
}
