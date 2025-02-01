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
	ItemImgUrl   string `gorm:"column:item_img_url"`       // 图片地址
	/*冗余存储sku信息*/
	SpecValues string `gorm:"column:spec_values"` // json存储，商品规格
	SkuAmount  string `gorm:"column:sku_amount"`  // sku金额

	IsDelivered bool `gorm:"column:is_delivered"` //是否发货
	IsConfirmed bool `gorm:"column:is_confirmed"` //是否签收
}
