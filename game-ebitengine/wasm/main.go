package main

import (
	"log"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":8088", http.FileServer(http.Dir("."))); err != nil {
		log.Fatal(err)
	}
}
