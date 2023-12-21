package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/learn-go/rpc/provider"
)

var ip = flag.String("ip", "127.0.0.1", "rcp service ip")
var port = flag.Int("port", 8899, "rpc service port")

func main() {
	flag.Parse()
	if *ip == "" || *port == 0 {
		panic("init ip and port error")
	}

	srv := provider.NewRPCServer(*ip, *port)
	srv.RegisterName("Test", &TestHandler{})
	srv.RegisterName("User", &UserHandler{})
	srv.Register(User{})
	go srv.Run()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	<-quit
	srv.Close()
}

type TestHandler struct{}

func (t *TestHandler) Hello() string {
	return "hello world"
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = map[int]User{
	1: {1, "Hero", 11},
	2: {2, "Lin07ux", 18},
}

type UserHandler struct{}

func (u *UserHandler) GetUserById(id int) (User, error) {
	if u, ok := users[id]; ok {
		return u, nil
	}
	return User{}, fmt.Errorf("user %d not found", id)
}
