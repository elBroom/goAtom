package db

import (
	"database/sql"
	"log"
)

func Init() *sql.DB {
	db, err := sql.Open("sqlite3", "./radish.db")
	//defer db.Close() in main  =(

	if err != nil {
		log.Fatal(err)
	}

	db.Exec(`
		CREATE TABLE IF NOT EXISTS User (
			id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			login VARCHAR (32) NOT NULL,
			password VARCHAR (32) NOT NULL
		);
	`)

	db.Exec(`
		CREATE TABLE IF NOT EXISTS Database (
			id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			name VARCHAR (32) NOT NULL,
			user_id INTEGER NOT NULL,
			CONSTRAINT FK_User_Database FOREIGN KEY (user_id) REFERENCES User(id)
		);
	`)

	db.Exec(`
		CREATE TABLE IF NOT EXISTS Query (
			id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			time DATATIME (32) NOT NULL,
			query VARCHAR (32) NOT NULL,
			database_id INTEGER NOT NULL,
			CONSTRAINT FK_Database_Query FOREIGN KEY (database_id) REFERENCES Database(id)
		);
	`)

	return db
}