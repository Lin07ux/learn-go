package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
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
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "" && password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, `{"code": -1, "message": "username and password can not be empty"}`)
		return
	}
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
		log.Println("invalid user id:", vars["id"])
		w.WriteHeader(http.StatusNotFound)
		return
	}

	user := database.QueryUser(int64(id))
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	_, _ = fmt.Fprintf(w, "%s", res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux2.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("invalid user id:", vars["id"])
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = database.DeleteUser(int64(id))
	if err != nil {
		log.Println("delete user failed:", err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type Teacher struct {
	Name    string
	Subject string
}

type Student struct {
	Id      int
	Name    string
	Country string
}

type Rooster struct {
	Teacher  Teacher
	Students []Student
}

func ShowRooster(w http.ResponseWriter, r *http.Request) {
	teacher := Teacher{
		Name:    "Alex",
		Subject: "Physics",
	}
	students := []Student{
		{Id: 1001, Name: "Peter", Country: "China"},
		{Id: 1002, Name: "Jeniffer", Country: "Sweden"},
	}
	rooster := Rooster{
		Teacher:  teacher,
		Students: students,
	}

	tmpl, err := template.ParseFiles("./templates/layout.gohtml",
		"./templates/nav.gohtml",
		"./templates/content.gohtml")
	if err != nil {
		log.Println("parse rooster template failed:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_ = tmpl.Execute(w, rooster)
}
