package model

import (
	"encoding/json"

	"github.com/lzl-here/bt-shop-backend/pkg/model"
	ogen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/order"

	"github.com/lzl-here/bt-shop-backend/pkg/utils"
)

type OrderItem struct {
	ID      uint64 `gorm:"column:id" json:"id"`
	TradeNo string `gorm:"column:trade_no" json:"trade_no"`
	OrderNo string `gorm:"column:order_no" json:"order_no"`
	SpuID   uint64 `gorm:"column:spu_id" json:"spu_id"`
	SkuID   uint64 `gorm:"column:sku_id" json:"sku_id"`

	ShopID   uint64 `gorm:"column:shop_id" json:"shop_id"`
	SellerID uint64 `gorm:"column:seller_id" json:"seller_id"`
	BuyerID  uint64 `gorm:"column:buyer_id" json:"buyer_id"`

	/*冗余存储spu信息*/
	SpuName      string `gorm:"column:spu_name" json:"spu_name"`           // spu名称
	CategoryID   uint64 `gorm:"column:category_id" json:"category_id"`     // 分类id
	CategoryName string `gorm:"column:category_name" json:"category_name"` // 分类名称
	BrandID      uint64 `gorm:"column:brand_id" json:"brand_id"`           // 品牌ID
	BrandName    string `gorm:"column:brand_name" json:"brand_name"`       // 品牌名称
	ItemImgUrl   string `gorm:"column:item_img_url" json:"item_img_url"`   // 图片地址
	/*冗余存储sku信息*/
	SpecValues string `gorm:"column:spec_values" json:"spec_values"` // json存储，商品规格
	SkuAmount  string `gorm:"column:sku_amount" json:"sku_amount"`   // sku金额

	IsDelivered bool `gorm:"column:is_delivered" json:"is_delivered"` //是否发货
	IsConfirmed bool `gorm:"column:is_confirmed" json:"is_confirmed"` //是否签收

	model.BaseModel
}

func (*OrderItem) TableName() string {
	return "order_item"
}

func (t *OrderItem) ToGen() (*ogen.BaseOrderItem, error) {
	res := &ogen.BaseOrderItem{}
	var err error
	if err = utils.Copy(t, res); err != nil {
		return nil, err
	}
	specs := []*ogen.BaseSpecValue{}
	if err = json.Unmarshal([]byte(t.SpecValues), &specs); err != nil {
		return nil, err
	}
	res.SpecValueList = specs
	return res, err
}
