package database

import (
	"database/sql"
	"log"
)

func getDbConnection() *sql.DB {
	db, err := sql.Open("mysql", "go_web:go_web@tcp(database:3306)/go_web?parseTime=true")
	if err != nil {
		log.Fatal("open mysql failed:", err)
	}

	return db
}
