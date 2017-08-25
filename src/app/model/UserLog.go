package model

import (
	"github.com/jinzhu/gorm"
)

type UserLog struct {
	gorm.Model

	UserID			 uint		`sql:"not null`
	User 			 User 		`gorm:"ForeignKey:UserID"`
}

func (UserLog) TableName() string {
	return "user_log"
}