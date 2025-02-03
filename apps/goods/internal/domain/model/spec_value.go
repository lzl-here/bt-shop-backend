package model

import "github.com/lzl-here/bt-shop-backend/pkg/model"

// 规格值, 和规格1:N关系
type SpecValue struct {
	ID     uint64 `gorm:"column:id"`      //ID
	SpecID uint64 `gorm:"column:spec_id"` // 关联的规格
	SpuID  uint64 `gorm:"column:spu_id"`  // spu_id
	SkuID  uint64 `gorm:"column:sku_id"`  // sku_id
	Value  string `gorm:"column:value"`   // 规格值
	Name   string `gorm:"column:name"`    // 规格项名称(spec的name，冗余存储)
	model.BaseModel
}

func (*SpecValue) TableName() string {
	return "spec_value"
}
