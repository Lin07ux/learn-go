package handlers

import (
	"net/http"

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

	usersRouter := r.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("/", InsertUser).Methods("POST")
	usersRouter.HandleFunc("/", UserList).Methods("GET")
	usersRouter.HandleFunc("/{id}", UserDetail).Methods("GET")
	usersRouter.HandleFunc("/{id}", DeleteUser).Methods("DELETE")

	visitorRouter := r.PathPrefix("/visitors").Subrouter()
	visitorRouter.Use(middleware.Method("GET"))
	visitorRouter.HandleFunc("/{name}/countries/{country}", ShowVisitor)

	viewRouter := r.PathPrefix("/view").Subrouter()
	viewRouter.HandleFunc("/rooster", ShowRooster)

	fs := http.FileServer(http.Dir("assets/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/login", UserLogin).Methods("POST")
	userRouter.HandleFunc("/logout", UserLogout)
	userRouter.HandleFunc("/profile", UserProfile)

	return r
}
