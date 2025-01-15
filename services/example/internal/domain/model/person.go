package model

import (
	"encoding/json"

	"github.com/lzl-here/bt-shop-backend/pkg/model"
)

type Person struct {
	ID              uint64 `gorm:"column:id"`
	Name            string `gorm:"column:name"`
	Age             int32  `gorm:"column:age"`
	Email           string `gorm:"column:email"`
	Password        string `gorm:"column:password"`
	model.TimeModel        // 时间类型
}

func (p *Person) TableName() string {
	return "t_person"
}

func (u *Person) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
