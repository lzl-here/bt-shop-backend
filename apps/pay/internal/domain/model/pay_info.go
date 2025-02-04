package model

import "github.com/lzl-here/bt-shop-backend/pkg/model"

// 支付详情表
type PayInfo struct {
	ID         uint64 `gorm:"column:id" json:"id"`
	TradeNo    string `gorm:"column:trade_no" json:"trade_no"`
	PayPageUrl string `gorm:"column:pay_page_url" json:"pay_page_url"`
	model.BaseModel
}

func (*PayInfo) TableName() string {
	return "pay_info"
}
