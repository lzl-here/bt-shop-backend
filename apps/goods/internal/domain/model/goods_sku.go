package model

import (
	"github.com/lzl-here/bt-shop-backend/pkg/model"
)

type GoodsSku struct {
	ID         uint64 `gorm:"column:id" json:"id"`                   // ID
	SpuID      uint64 `gorm:"column:spu_id" json:"spu_id"`           // spu id
	StockNum   uint64 `gorm:"column:stock_num" json:"stock_num"`     // 库存
	SkuPrice   string `gorm:"column:sku_price" json:"sku_price"`     //sku价格
	SpecValues string `gorm:"column:spec_values" json:"spec_values"` // json存储，商品规格
	Enabled    bool   `gorm:"column:enabled" json:"enabled"`         // 是否上架
	model.BaseModel
}

func (*GoodsSku) TableName() string {
	return "goods_sku"
}
