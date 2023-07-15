package main

import (
	"fmt"
	"net/http"

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
	_, _ = fmt.Fprintf(w, "success")
}
