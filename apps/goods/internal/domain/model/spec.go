package model

import (
	"github.com/lzl-here/bt-shop-backend/pkg/model"
)

// 规格
type Spec struct {
	ID       uint64 `gorm:"column:id" json:"id"`         // ID
	SpecName string `gorm:"column:name" json:"spec_name"`     // 规格名称
	SpuID    uint64 `gorm:"column:spu_id" json:"spu_id"` // spu id
	model.BaseModel
}

func (*Spec) TableName() string {
	return "spec"
}
