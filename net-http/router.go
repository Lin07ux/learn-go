package main

import (
	"github.com/gorilla/mux"

	"github.com/learn-go/net-http/middleware"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.Logging())
	r.Handle("/", &HelloWorld{content: "Hello World Struct"})

	mysqlRouter := r.PathPrefix("/mysql").Subrouter()
	mysqlRouter.HandleFunc("/ping", MySqlDemo)
	mysqlRouter.HandleFunc("/table", CreateMySqlTable)

	userRouter := r.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/", InsertUser).Methods("POST")
	userRouter.HandleFunc("/", UserList).Methods("GET")
	userRouter.HandleFunc("/{id}", UserDetail).Methods("GET")

	visitorRouter := r.PathPrefix("/visitors").Subrouter()
	visitorRouter.Use(middleware.Method("GET"))
	visitorRouter.HandleFunc("/{name}/countries/{country}", ShowVisitor)

	return r
}
