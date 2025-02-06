package model

import (
	"github.com/lzl-here/bt-shop-backend/pkg/model"
	"github.com/lzl-here/bt-shop-backend/pkg/utils"
	ugen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/user"
)

// 店铺
type Shop struct {
	ID         uint64 `gorm:"column:id" json:"id"`
	ShopName   string `gorm:"column:shop_name" json:"shop_name"`
	ShopDesc   string `gorm:"column:shop_desc" json:"shop_desc"`
	ShopImgUrl string `gorm:"column:shop_img_url" json:"shop_img_url"`
	UserID     uint64 `gorm:"column:user_id" json:"user_id"`
	model.BaseModel
}

func (*Shop) TableName() string {
	return "shop"
}

func (s *Shop) ToGen() (*ugen.BaseShop, error) {
	shop := &ugen.BaseShop{}
	utils.Copy(s, shop)
	return shop, nil
}
