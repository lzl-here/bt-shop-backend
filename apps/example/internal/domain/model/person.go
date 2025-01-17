package model

import (
	"encoding/json"

	"github.com/lzl-here/bt-shop-backend/pkg/model"
)

type Person struct {
	Name            string `gorm:"column:name"`
	Age             int32  `gorm:"column:age"`
	Email           string `gorm:"column:email"`
	Password        string `gorm:"column:password"`
	model.BaseModel        
}

func (p *Person) TableName() string {
	return "t_person"
}

func (u *Person) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
