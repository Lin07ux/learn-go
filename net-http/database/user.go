package database

import (
	"database/sql"
	"fmt"
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

func QueryUserList() []User {
	db := getDbConnection()
	defer func() {
		_ = db.Close()
	}()

	var users []User
	rows, err := db.Query("SELECT * FROM users ORDER BY id")
	if err != nil {
		log.Println("query users failed:", err)
		return users
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	for rows.Next() {
		var u User
		err = rows.Scan(&u.Id, &u.Username, &u.Password, &u.CreatedAt)
		if err != nil {
			log.Println("scan users rows failed:", err)
		} else {
			users = append(users, u)
		}
	}
	return users
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

func DeleteUser(id int64) error {
	db := getDbConnection()
	defer func() {
		_ = db.Close()
	}()

	result, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("delete user[%d] failed: %w", id, err)
	}

	var affected int64
	if affected, err = result.RowsAffected(); err != nil {
		return fmt.Errorf("get deleted users count failed: %w", err)
	}
	if affected == 0 {
		return fmt.Errorf("undeleted user")
	}
	return nil
}

func AuthenticateUser(username, password string) (user User, err error) {
	db := getDbConnection()
	defer func() {
		_ = db.Close()
	}()

	row := db.QueryRow("SELECT id, username FROM users WHERE username = ? AND password = ? limit 1", username, password)
	err = row.Scan(&user.Id, &user.Username)
	return
}
