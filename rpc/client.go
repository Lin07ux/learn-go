package main

import (
	"context"
	"encoding/gob"
	"log"

	"github.com/learn-go/rpc/consumer"
)

var GetUserById func(id int) (User, error)
var Hello func() string

func main() {
	gob.Register(User{})

	client := consumer.NewClientProxy(consumer.DefaultOption)
	ctx := context.Background()

	u1, err := client.Call(ctx, "UserService.User.GetUserById", &GetUserById, 1)
	log.Println("user 1 result:", u1, err)

	u3, err := GetUserById(3)
	log.Println("user 3 result:", u3, err)

	r1, err := client.Call(ctx, "UserService.Test.Hello", &Hello)
	log.Println("hello result:", r1, err)

	r2 := Hello()
	log.Println("hello 2 result:", r2)
}
