package model

import (
	"time"

	"github.com/lzl-here/bt-shop-backend/pkg/model"
)

// 结算单
type Bill struct {
	ID          uint64     `gorm:"column:id"`
	SellerID    uint64     `gorm:"column:seller_id"`    // 商家ID
	ShopID      uint64     `gorm:"column:shop_id"`      // 店铺ID
	TotalAmount string     `gorm:"column:total_amount"` // 金额
	Type        int        `gorm:"column:type"`         // 类型，1-入帐，2-出账
	BeginTime   *time.Time `gorm:"column:begin_time"`   // 结算的开始时间
	EndTime     *time.Time `gorm:"column:end_time"`     // 结算的结束时间
	Status      int        `gorm:"column:status"`       // 状态，1-生成，2-已完成(商家确认)
	model.BaseModel
}
