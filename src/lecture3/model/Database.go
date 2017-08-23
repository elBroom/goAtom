package model

import (
	"github.com/jinzhu/gorm"
)

type Database struct {
	gorm.Model

	Name             string
	User 			 User 	`gorm:"ForeignKey:Login"`
}

func (Database) TableName() string {
	return "database"
}