package main

import (
	"log"
)

func main() {
	if err := NewGame().Run(); err != nil {
		log.Fatal(err)
	}
}
