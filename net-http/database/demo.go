package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func MySqlDemo() {
	db, err := sql.Open("mysql", "go_web:go_web@tcp(database:3306)/go_web")
	if err != nil {
		log.Fatal("open mysql failed:", err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	if err = db.Ping(); err != nil {
		log.Fatal("ping mysql failed:", err)
	}

	log.Println("open and ping mysql success")
}
