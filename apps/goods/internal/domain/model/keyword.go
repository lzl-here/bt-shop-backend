package model

import "github.com/lzl-here/bt-shop-backend/pkg/model"

type Keyword struct {
	ID          uint64 `gorm:"column:id"`           //ID
	Value       string `gorm:"column:value"`        // 搜索关键字
	SearchTimes uint64 `gorm:"column:search_times"` // 搜索次数
	model.BaseModel
}
