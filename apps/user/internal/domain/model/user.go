package model

import (
	"github.com/lzl-here/bt-shop-backend/pkg/model"
	"github.com/lzl-here/bt-shop-backend/pkg/utils"
	ugen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/user"
)

type User struct {
	ID           uint64 `gorm:"column:id" json:"id"`
	Username     string `gorm:"column:username" json:"username"`
	AvatarUrl    string `gorm:"column:avatar_url" json:"avatar_url"`
	Password     string `gorm:"column:password" json:"password"`
	Salt         string `gorm:"column:salt" json:"salt"`                   // 密码盐值
	RegisterType string `gorm:"column:register_type" json:"register_type"` // 注册类型
	ShopID       uint64 `gorm:"column:shop_id" json:"shop_id"`
	model.BaseModel
}

func (*User) TableName() string {
	return "user"
}

func (u *User) ToGen() (*ugen.BaseUser, error) {
	user := &ugen.BaseUser{}
	if err := utils.Copy(u, user); err != nil {
		return nil, err
	}
	return user, nil
}
