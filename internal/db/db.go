package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func InitDB(path string) {
	var err error
	DB, err = sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}
	CreateTables()
}

func CreateTables() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS links (
			id 			INTEGER 	PRIMARY KEY,
			user_id 	INTEGER 		NOT NULL,
			url 		VARCHAR(255) 	NOT NULL,
			added_at	DATETIME	DEFAULT	CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
}