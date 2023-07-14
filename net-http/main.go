package main

import (
	"fmt"
	"log"
	"net/http"
)

type HelloWorld struct {
	content string
}

func (h *HelloWorld) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, h.content)
}

func main() {
	http.Handle("/", &HelloWorld{content: "Hello World Struct"})

	log.Println("Http server listening on: http://localhost:8800")
	err := http.ListenAndServe(":8800", nil)
	if err != nil {
		log.Fatal(err)
	}
}
