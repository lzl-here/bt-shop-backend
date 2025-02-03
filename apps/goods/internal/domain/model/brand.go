package model

import "github.com/lzl-here/bt-shop-backend/pkg/model"

type Brand struct {
	ID      uint64 `gorm:"column:id" json:"id"`             // ID
	Name    string `gorm:"column:name" json:"name"`         // 品牌名称
	IconUrl string `gorm:"column:icon_url" json:"icon_url"` // 品牌图标
	model.BaseModel
}

func (*Brand) TableName() string {
	return "brand"
}
