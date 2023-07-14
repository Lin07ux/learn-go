package main

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/", &HelloWorld{content: "Hello World Struct"})

	visitorRouter := r.PathPrefix("/visitors").Subrouter()
	visitorRouter.HandleFunc("/{name}/countries/{country}", ShowVisitor)

	return r
}
