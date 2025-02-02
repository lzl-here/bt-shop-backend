package model

import "github.com/lzl-here/bt-shop-backend/pkg/model"

type Category struct {
	ID       uint64 `gorm:"column:id"`        // ID
	Name     string `gorm:"column:name"`      // 名称
	ParentID uint64 `gorm:"column:parent_id"` // 父分类
	Level    int32  `gorm:"column:level"`     // 第几层
	model.BaseModel
}
