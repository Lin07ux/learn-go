package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	mux2 "github.com/gorilla/mux"

	"github.com/learn-go/net-http/database"
)

type HelloWorld struct {
	content string
}

func (h *HelloWorld) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, h.content)
}

func ShowVisitor(w http.ResponseWriter, r *http.Request) {
	vars := mux2.Vars(r)
	_, _ = fmt.Fprintf(w, "Hello, %s of %s", vars["name"], vars["country"])
}

func MySqlDemo(w http.ResponseWriter, r *http.Request) {
	database.MySqlDemo()
	_, _ = fmt.Fprintf(w, "ping success")
}

func CreateMySqlTable(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, database.CreateTable())
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	username := "Lin07ux"
	password := "secret"
	id := database.InsertData(username, password)
	_, _ = fmt.Fprintf(w, "%d", id)
}

func UserList(w http.ResponseWriter, r *http.Request) {
	users := database.QueryUserList()
	res, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	_, _ = fmt.Fprintf(w, "%s", res)
}

func UserDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux2.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("invalid user id:", vars["id"])
		_, _ = fmt.Fprintf(w, "Not Found")
		return
	}

	user := database.QueryUser(int64(id))
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	_, _ = fmt.Fprintf(w, "%s", res)
}
