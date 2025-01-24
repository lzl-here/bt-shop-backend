package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	email    string `gorm:"uniqueIndex;type:varchar(255) not null"`
	password string `gorm:"type:varchar(255) not null"`
}

func (User) TableName() string {
	return "user"
}
