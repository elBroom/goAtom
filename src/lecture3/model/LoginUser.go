package model

import (
	"github.com/jinzhu/gorm"
)

type LoginUser struct {
	gorm.Model

	User             User `gorm:"ForeignKey:Login"`
	Token         	 float64
}

func (LoginUser) TableName() string {
	return "login_user"
}