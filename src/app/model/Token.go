package model

import (
	"github.com/jinzhu/gorm"
)

type Token struct {
	gorm.Model

	Token         	 float64	`sql:"not null`
	UserID			 int		`sql:"not null`
	User             User 		`gorm:"ForeignKey:UserID"`
}

func (Token) TableName() string {
	return "token"
}