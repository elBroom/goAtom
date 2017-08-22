package model

import (
	"github.com/jinzhu/gorm"
)

type Query struct {
	gorm.Model

	Query            string
	Database 		 Database 	`gorm:"ForeignKey:DatabaseID"`
}

func (Query) TableName() string {
	return "query"
}