package model

import "github.com/lzl-here/bt-shop-backend/pkg/model"

type Trade struct {
	TradeNo     string `gorm:"column:trade_no"`
	TradeAmount string `gorm:"column:trade_amount"`
	TradeStatus string `gorm:"column:trade_status"`
	model.BaseModel
}


func (t *Trade) TableName() string {
	return "t_trade"
}