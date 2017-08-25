package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Login            string `gorm:"type:varchar(64);unique_index" sql:"not null`
	Password         string `sql:"not null`
	Name             string
}

func (User) TableName() string {
	return "user"
}