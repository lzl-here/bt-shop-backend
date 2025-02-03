package model

import "github.com/lzl-here/bt-shop-backend/pkg/model"

type GoodsSpu struct {
	ID           uint64 `gorm:"column:id" json:"id"`                       // ID
	SpuName      string `gorm:"column:spu_name" json:"spu_name"`           // spu名称
	SpuDesc      string `gorm:"column:spu_desc" json:"spu_desc"`           // 描述
	CategoryID   uint64 `gorm:"column:category_id" json:"category_id"`     // 分类id
	CategoryName string `gorm:"column:category_name" json:"category_name"` // 分类名称
	BrandID      uint64 `gorm:"column:brand_id" json:"brand_id"`           // 品牌ID
	BrandName    string `gorm:"column:brand_name" json:"brand_name"`       // 品牌名称
	Price        string `gorm:"column:price" json:"price"`                 // 价格，展示使用，不用于交易
	Enabled      bool   `gorm:"column:enabled" json:"enabled"`             // 是否上架
	SpuImgUrl    string `gorm:"column:spu_img_url" json:"spu_img_url"`         // 图片地址
	AttributeIDs string `gorm:"column:attribute_ids" json:"attribute_ids"` // 属性id列表 ,隔开 不额外用一张表存关系
	SpecIDs      string `gorm:"column:spec_ids" json:"spec_ids"`           // 规格id列表 ,隔开 不额外用一张表存关系
	model.BaseModel
}

func (*GoodsSpu) TableName() string {
	return "goods_spu"
}
