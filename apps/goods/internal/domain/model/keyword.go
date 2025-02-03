package model

import "github.com/lzl-here/bt-shop-backend/pkg/model"

type Keyword struct {
	ID          uint64 `gorm:"column:id" json:"id"`                     //ID
	Value       string `gorm:"column:value" json:"value"`               // 搜索关键字
	SearchTimes uint64 `gorm:"column:search_times" json:"search_times"` // 搜索次数
	model.BaseModel
}

func (*Keyword) TableName() string {
	return "keyword"
}
