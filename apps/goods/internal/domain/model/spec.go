package model

import (
	"github.com/lzl-here/bt-shop-backend/pkg/model"
)

// 规格
type Spec struct {
	ID    uint64 `gorm:"column:id"`
	Name  string `gorm:"column:name"`   // 规格名称
	Value string 
	SpuID uint64 `gorm:"column:spu_id"` // spu id
	model.BaseModel
}

