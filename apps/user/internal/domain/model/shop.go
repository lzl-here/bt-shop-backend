package model

import "github.com/lzl-here/bt-shop-backend/pkg/model"

// 店铺
type Shop struct {
	ID         uint64 `gorm:"column:id" json:"id"`
	ShopName   string `gorm:"column:shop_name" json:"shop_name"`
	ShopDesc   string `gorm:"column:shop_desc" json:"shop_desc"`
	ShopImgUrl string `gorm:"column:shop_img_url" json:"shop_img_url"`
	model.BaseModel
}

func (*Shop) TableName() string {
	return "shop"
}
