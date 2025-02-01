package model

type OrderItem struct {
	ID      string `gorm:"column:id"`
	TradeNo string `gorm:"column:trade_no"`
	OrderNo string `gorm:"column:order_no"`
	SpuID   uint64 `gorm:"column:spu_id"`
	SkuID   uint64 `gorm:"column:sku_id"`
	/*冗余存储spu信息*/
	SpuName      string `gorm:"column:name"`          // spu名称
	CategoryID   uint64 `gorm:"column:category_id"`   // 分类id
	CategoryName string `gorm:"column:category_name"` // 分类名称
	BrandID      uint64 `gorm:"column:brand_id"`      // 品牌ID
	BrandName    string `gorm:"column:brand_name"`    // 品牌名称
	MinAmount    string `gorm:"column:max_amount"`    // 最低价
	MaxAmount    string `gorm:"column:max_amount"`    // 最高价
	ImgUrl       string `gorm:"column:img_url"`       // 图片地址
	/*冗余存储sku信息*/
	SpecValues string `gorm:"column:spec_values"` // json存储，商品规格

	IsDelivered int `gorm:"column:is_delivered"` //是否发货
	IsConfirmed int `gorm:"column:is_confirmed"` //是否签收
}
