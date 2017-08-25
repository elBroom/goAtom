package model

import (
	"github.com/jinzhu/gorm"
)

type QueryLog struct {
	gorm.Model

	Query            string		`sql:"not null`
	Params		     string
	UserID			 uint		`sql:"not null`
	User 			 User 		`gorm:"ForeignKey:UserID"`
}

func (QueryLog) TableName() string {
	return "query_log"
}