package model

import (
	"github.com/lzl-here/bt-shop-backend/pkg/model"
)

// 规格
type Spec struct {
	ID    uint64 `gorm:"column:id"`     // ID
	Name  string `gorm:"column:name"`   // 规格名称
	SpuID uint64 `gorm:"column:spu_id"` // spu id
	model.BaseModel
}

func (*Spec) TableName() string {
	return "spec"
}
