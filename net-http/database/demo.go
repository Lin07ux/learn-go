package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func MySqlDemo() {
	db := getDbConnection()
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	if err := db.Ping(); err != nil {
		log.Fatal("ping mysql failed:", err)
	}

	log.Println("open and ping mysql success")
}

func CreateTable() string {
	db := getDbConnection()
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	query := `CREATE TABLE users (
       id INT AUTO_INCREMENT,
       username TEXT NOT NULL,
       password TEXT NOT NULL,
       created_at DATETIME,
       PRIMARY KEY (id)
   );`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("create table failed:", err)
	}

	return query
}
