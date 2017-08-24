package db

import (
	"log"

	"../config"
	"../model"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm"
)

var sqlite_connect  (*gorm.DB)  = nil

func Sqlite_connect() *gorm.DB {
	if sqlite_connect != nil {
		return sqlite_connect
	}

	cfg := config.GetSqlite()
	sqlite_connect, err := gorm.Open("sqlite3", cfg.DatabasePath)
	//defer connect.Close() in main  =(

	if err != nil {
		log.Fatal(err)
	}

	sqlite_connect.LogMode(true)

	sqlite_connect.AutoMigrate(&model.User{}, &model.Token{}, &model.UserLog{}, &model.QueryLog{})

	return sqlite_connect
}