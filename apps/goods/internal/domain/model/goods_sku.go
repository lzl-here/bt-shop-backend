package model

import "github.com/lzl-here/bt-shop-backend/pkg/model"

type GoodsSku struct {
	ID         uint64 `gorm:"column:id"`          //ID
	SpuID      uint64 `gorm:"column:spu_id"`      // spu id
	StockNum   uint64 `gorm:"column:stock_num"`   // 库存
	SpecValues string `gorm:"column:spec_values"` // json存储，商品规格
	Enabled    bool   `gorm:"column:enabled"`     // 是否上架
	model.BaseModel
}
