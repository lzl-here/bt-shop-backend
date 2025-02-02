package model

import "github.com/lzl-here/bt-shop-backend/pkg/model"

type GoodsSpu struct {
	ID           uint64 `gorm:"column:id"`
	SpuName      string `gorm:"column:name"`          // spu名称
	SpuDesc      string `gorm:"column:spu_desc"`      //描述
	CategoryID   uint64 `gorm:"column:category_id"`   // 分类id
	CategoryName string `gorm:"column:category_name"` // 分类名称
	BrandID      uint64 `gorm:"column:brand_id"`      // 品牌ID
	BrandName    string `gorm:"column:brand_name"`    // 品牌名称
	MinAmount    string `gorm:"column:max_amount"`    // 最低价
	MaxAmount    string `gorm:"column:max_amount"`    // 最高价
	Enabled      uint64 `gorm:"column:enabled"`       // 是否上架
	SpuImgUrl    string `gorm:"column:img_url"`       // 图片地址
	Attributes   string `gorm:"column:attributes"`    // 属性id列表 ,隔开 不额外用一张表存关系
	Specs        string `gorm:"column:specs"`         // 规格id列表 ,隔开 不额外用一张表存关系
	model.BaseModel
}
