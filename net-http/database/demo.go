package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id        int64
	Username  string
	Password  string
	CreatedAt time.Time
}

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

func InsertData(username, password string) int64 {
	db := getDbConnection()
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	result, err := db.Exec("INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)", username, password, time.Now())
	if err != nil {
		log.Println("insert user failed: ", err)
		return 0
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("get insert user id failed", err)
		return 0
	}
	return id
}

func QueryUser(id int64) User {
	db := getDbConnection()
	defer func() {
		_ = db.Close()
	}()

	user := User{}
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		log.Println("query user failed:", err)
	}
	return user
}
