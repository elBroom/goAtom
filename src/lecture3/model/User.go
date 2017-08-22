package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Login            string `gorm:"type:varchar(64);unique_index"`
	Password         string
	Name             string
}

func (User) TableName() string {
	return "user"
}