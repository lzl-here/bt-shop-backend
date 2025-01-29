package model

import (
	"encoding/json"

	"github.com/lzl-here/bt-shop-backend/pkg/model"
)

type PayFlow struct {
	ID            uint64 `gorm:"column:id"`
	PayState      int    `gorm:"column:pay_state"`
	TradeNo       string `gorm:"column:trade_no"`
	ThirdTradeNo  string `gorm:"column:third_trade_no"`
	ThirdBuyerID  string `gorm:"column:third_buyer_id"`
	ThirdSellerID string `gorm:"column:third_seller_id"`
	TotalAmount   string `gorm:"column:total_amount"`

	// TODO 店铺信息
	// TODO 卖家信息
	// TODO 买家信息
	// TODO 支付方式：1.支付宝 2.微信 3.银行卡 4.其他
	// TODO 退款信息
	model.BaseModel
}

func (p *PayFlow) TableName() string {
	return "t_pay_flow"
}

func (u *PayFlow) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
