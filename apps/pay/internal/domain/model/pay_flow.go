package model

import (
	"encoding/json"

	"github.com/lzl-here/bt-shop-backend/pkg/model"
)

// 支付流水，一笔支付流水代表一次资金流向，用于商家进行结算
// 一笔订单对应一笔支付流水（注意不是交易，一个交易可能涉及多个商家，资金会流向多个账户，为了方便结算一笔订单对应一笔支付流水）
// 支付时生成正向
// 退款时生成逆向
type PayFlow struct {
	ID            uint64 `gorm:"column:id"`              // 支付流水号
	PayState      int    `gorm:"column:pay_state"`       // 支付状态
	TradeNo       string `gorm:"column:trade_no"`        // 电商交易号
	OrderNo       string `gorm:"column:order_no"`        // 电商订单号
	ThirdTradeNo  string `gorm:"column:third_trade_no"`  // 支付机构交易号
	ThirdBuyerID  string `gorm:"column:third_buyer_id"`  // 支付机构买家id
	ThirdSellerID string `gorm:"column:third_seller_id"` // 支付机构卖家id
	TotalAmount   string `gorm:"column:total_amount"`    // 流转的资金金额

	// TODO 正向还是逆向
	// TODO 店铺信息
	// TODO 卖家信息
	// TODO 买家信息
	// TODO 支付方式：1.支付宝 2.微信 3.银行卡 4.其他
	// TODO 订单项
	model.BaseModel
}

func (p *PayFlow) TableName() string {
	return "t_pay_flow"
}

func (u *PayFlow) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
