package model

import "github.com/lzl-here/bt-shop-backend/pkg/model"

type Order struct {
	TradeNo     string `gorm:"column:trade_no"`
	OrderNo     string `gorm:"column:order_no"`
	OrderAmount string `gorm:"column:order_amount"`
	OrderStatus string `gorm:"column:order_status"`
	model.BaseModel
}
