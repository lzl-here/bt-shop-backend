package model

import "github.com/lzl-here/bt-shop-backend/pkg/model"

// 规格值, 和规格1:N关系
type SpecValue struct {
	ID        uint64 `gorm:"column:id" json:"id"`            //ID
	SpecID    uint64 `gorm:"column:spec_id" json:"spec_id"`  // 关联的规格
	SpuID     uint64 `gorm:"column:spu_id" json:"spu_id"`    // spu_id
	SkuID     uint64 `gorm:"column:sku_id" json:"sku_id"`    // sku_id
	SpecValue string `gorm:"column:value" json:"spec_value"` // 规格值
	SpecName  string `gorm:"column:name" json:"spec_name"`   // 规格项名称(spec的name，冗余存储)
	model.BaseModel
}

func (*SpecValue) TableName() string {
	return "spec_value"
}
