package main

import (
	"log"
)

func main() {
	if err := NewGame(LoadConfig()).Run(); err != nil {
		log.Fatal(err)
	}
}
