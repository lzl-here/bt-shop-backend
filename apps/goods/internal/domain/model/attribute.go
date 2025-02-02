package model

import "github.com/lzl-here/bt-shop-backend/pkg/model"

// 商品属性
type Attribute struct {
	Id          int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name        string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Value       string `gorm:"column:value;type:varchar(255);not null" json:"value"`
	Description string `gorm:"column:description;type:varchar(255);not null" json:"description"`
	model.BaseModel
}
