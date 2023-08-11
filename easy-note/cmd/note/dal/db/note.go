package db

import (
	"context"

	"gorm.io/gorm"
)


type Note struct {
	gorm.Model

	UserID  int64  `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (n *Note) TableName() string {
	return "note"
}


// CreateNote create note info
func CreateNote(ctx context.Context, notes []*Note) error {
	return DB.WithContext(ctx).Create(notes).Error
}

// MGetNotes multiple get list of note info
func MGetNotes(ctx context.Context, noteIDs []int64) ([]*Note, error) {
	var res []*Note
	if len(noteIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", noteIDs).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

// UpdateNote update note info
func UpdateNote(ctx context.Context, noteID, userID int64, title, content *string) error {
	params := map[string]interface{}{}
	if title != nil {
		params["title"] = *title
	}
	if content != nil {
		params["content"] = *content
	}

	return DB.WithContext(ctx).Model(&Note{}).Where("id = ? and user_id = ?", noteID, userID).
		Updates(params).Error
}

func UpdateNoteV1(ctx context.Context, note *Note) error {
	return DB.WithContext(ctx).Updates(note).Error
}

// DeleteNote delete note info
func DeleteNote(ctx context.Context, noteID, userID int64) error {
	return DB.WithContext(ctx).Where("id = ? and user_id = ? ", noteID, userID).Delete(&Note{}).Error
}

// QueryNote query list of note info
func QueryNote(ctx context.Context, userID int64, searchKey *string, limit, offset int) ([]*Note, int64, error) {
	var total int64
	var res []*Note
	conn := DB.WithContext(ctx).Model(&Note{}).Where("user_id = ?", userID)

	if searchKey != nil {
		conn = conn.Where("title like ?", "%" + *searchKey + "%")
	}

	if err := conn.Count(&total).Error; err != nil {
		return res, total, err
	}

	if err := conn.Limit(limit).Offset(offset).Find(&res).Error; err != nil {
		return res, total, err
	}

	return res, total, nil
}
