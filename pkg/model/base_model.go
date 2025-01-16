package model

import "time"

type BaseModel struct {
	ID        uint64      `json:"id,omitempty" gorm:"column:id;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"column:deleted_at;"`
}

func (t *BaseModel) CreatedString() string {
	if t.CreatedAt == nil || t.CreatedAt.IsZero() {
		return ""
	}
	return t.CreatedAt.Format("2006-01-02 15:04:05")
}

func (t *BaseModel) UpdatedString() string {
	if t.UpdatedAt == nil || t.UpdatedAt.IsZero() {
		return ""
	}
	return t.UpdatedAt.Format("2006-01-02 15:04:05")
}

func (t *BaseModel) DeletedString() string {
	if t.DeletedAt == nil || t.DeletedAt.IsZero() {
		return ""
	}
	return t.DeletedAt.Format("2006-01-02 15:04:05")
}
