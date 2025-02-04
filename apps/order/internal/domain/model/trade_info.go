package model

import "github.com/lzl-here/bt-shop-backend/pkg/model"

// 支付详情表
type TradeInfo struct {
	ID      uint64 `json:"id"`
	TradeNo string `json:"trade_no"`
	PageUrl string `json:"page_url"`
	model.BaseModel
}

func (*TradeInfo) TableName() string {
	return "trade_info"
}
