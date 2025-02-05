package model

import (
	"github.com/lzl-here/bt-shop-backend/pkg/model"
	"github.com/lzl-here/bt-shop-backend/pkg/utils"
	ogen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/order"
)

type Trade struct {
	TradeNo     string `gorm:"column:trade_no" json:"trade_no"`
	TradeAmount string `gorm:"column:trade_amount" json:"trade_amount"`
	TradeStatus string `gorm:"column:trade_status" json:"trade_status"`
	BuyerID     uint64 `gorm:"column:buyer_id" json:"buyer_id"`
	PayType     string `gorm:"column:pay_type" json:"pay_type"`
	model.BaseModel
}

func (t *Trade) TableName() string {
	return "trade"
}

func (t *Trade) ToGen() (*ogen.BaseTrade, error) {
	res := &ogen.BaseTrade{}
	var err error
	if err = utils.Copy(t, res); err != nil {
		return nil, err
	}
	return res, nil
}
