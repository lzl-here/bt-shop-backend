package model

import "github.com/lzl-here/bt-shop-backend/pkg/model"

type Trade struct {
	TradeNo     string `gorm:"column:trade_no"`
	TradeAmount string `gorm:"column:trade_amount"`
	TradeStatus string `gorm:"column:trade_status"`
	BuyerID     uint64 `gorm:"column:buyer_id"`
	PayType     string `gorm:"column:pay_type"`
	model.BaseModel
}

func (t *Trade) TableName() string {
	return "trade"
}
