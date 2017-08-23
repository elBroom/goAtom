package db

import (
	"log"

	"../model"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm"
)

var sqlite_connect  (*gorm.DB)  = nil

func Sqlite_connect() *gorm.DB {
	if sqlite_connect != nil {
		return sqlite_connect
	}

	sqlite_connect, err := gorm.Open("sqlite3", "./radish.db")
	//defer connect.Close() in main  =(

	if err != nil {
		log.Fatal(err)
	}

	sqlite_connect.LogMode(true)

	sqlite_connect.AutoMigrate(&model.User{}, &model.Database{}, &model.Query{}, &model.LoginUser{})

	return sqlite_connect
}