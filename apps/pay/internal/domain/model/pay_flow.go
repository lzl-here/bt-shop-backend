package model

import (
	"encoding/json"

	"github.com/lzl-here/bt-shop-backend/pkg/model"
)

// 支付流水，一笔支付流水代表一次资金流向，用于商家进行结算
// 一笔订单对应一笔支付流水（注意不是交易，一个交易可能涉及多个商家，资金会流向多个账户，为了方便结算一个订单项对应一笔支付流水）
// 支付时生成正向
// 退款时生成逆向
type PayFlow struct {
	PayNo         string `gorm:"column:pay_no"`          // 支付流水号
	TradeNo       string `gorm:"column:trade_no"`        // 电商交易号
	OrderNo       string `gorm:"column:order_no"`        // 电商订单号
	ThirdTradeNo  string `gorm:"column:third_trade_no"`  // 支付机构交易号
	Status        string `gorm:"column:status"`          // 流转状态 支付成功：success，取消：cancel，退款: refund, 完成结账：completed
	OrderItemID   uint64 `gorm:"column:order_item_id"`   //	关联的订单项号
	InOutType     string `gorm:"column:in_out_type"`     // 出入账类型：入帐: in，出账: out
	ThirdBuyerID  string `gorm:"column:third_buyer_id"`  // 支付机构买家id
	ThirdSellerID string `gorm:"column:third_seller_id"` // 支付机构卖家id
	TotalAmount   string `gorm:"column:total_amount"`    // 流转的资金金额
	ShopID        uint64 `gorm:"column:shop_id"`         // 店铺id
	SellerID      uint64 `gorm:"column:seller_id"`       //	卖家id
	BuyerID       uint64 `gorm:"column:buyer_id"`        //	买家id
	PayType       string `gorm:"column:pay_type"`        // 支付方式：1.支付宝 2.微信 3.银行卡 4.其他

	model.BaseModel
}

func (p *PayFlow) TableName() string {
	return "t_pay_flow"
}

func (u *PayFlow) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
