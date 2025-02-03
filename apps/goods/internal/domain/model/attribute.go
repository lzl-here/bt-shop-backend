package model

import "github.com/lzl-here/bt-shop-backend/pkg/model"

// 商品属性
type Attribute struct {
	ID    uint64 `gorm:"column:id;" json:"id"`         // ID
	SpuID uint64 `gorm:"column:spu_id;" json:"spu_id"` // spuID
	Name  string `gorm:"column:name;" json:"name"`     // 属性名
	Value string `gorm:"column:value;" json:"value"`   // 属性值
	model.BaseModel
}

func (*Attribute) TableName() string {
	return "attribute"
}
