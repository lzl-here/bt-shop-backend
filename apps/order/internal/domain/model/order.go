package model

import "github.com/lzl-here/bt-shop-backend/pkg/model"

type Order struct {
	OrderNo     string `gorm:"column:order_no"`
	TradeNo     string `gorm:"column:trade_no"`
	OrderAmount string `gorm:"column:order_amount"`
	OrderStatus string `gorm:"column:order_status"`
	SellerID    uint64 `gorm:"column:seller_id"`
	BuyerID     uint64 `gorm:"column:buyer_id"`
	PayType     string `gorm:"column:pay_type"`
	model.BaseModel
}

func (*Order) TableName() string {
	return "order"
}
