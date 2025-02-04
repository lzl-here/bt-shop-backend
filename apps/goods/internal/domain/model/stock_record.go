package model

import "github.com/lzl-here/bt-shop-backend/pkg/model"

// 库操作记录
type StockOpRecord struct {
	ID      uint64 `gorm:"column:id"`
	TradeNo string `gorm:"column:trade_no"`
	OpType  string `gorm:"column:op_type"`
	model.BaseModel
}

func (*StockOpRecord) TableName() string {
	return "stock_op_record"
}
