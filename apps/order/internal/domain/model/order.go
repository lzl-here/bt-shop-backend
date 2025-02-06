package model

import (
	"github.com/lzl-here/bt-shop-backend/pkg/model"
	"github.com/lzl-here/bt-shop-backend/pkg/utils"
	ogen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/order"
)

type Order struct {
	OrderNo     string `gorm:"column:order_no" json:"order_no"`
	TradeNo     string `gorm:"column:trade_no" json:"trade_no"`
	ShopID      uint64 `gorm:"column:shop_id" json:"shop_id"`
	OrderAmount string `gorm:"column:order_amount" json:"order_amount"`
	SellerID    uint64 `gorm:"column:seller_id" json:"seller_id"`
	OrderStatus string `gorm:"column:order_status" json:"order_status"`
	BuyerID     uint64 `gorm:"column:buyer_id" json:"buyer_id"`
	PayType     string `gorm:"column:pay_type" json:"pay_type"`
	model.BaseModel
}

func (*Order) TableName() string {
	return "order"
}

func (t *Order) ToGen() (*ogen.BaseOrder, error) {
	res := &ogen.BaseOrder{}
	var err error
	if err = utils.Copy(t, res); err != nil {
		return nil, err
	}
	res.CreatedAt = t.CreatedString()
	return res, nil
}
