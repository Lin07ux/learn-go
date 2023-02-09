package main

import (
	"log"
)

//go:generate go install github.com/hajimehoshi/file2byteslice/cmd/file2byteslice@latest
//go:generate mkdir resources
//go:generate file2byteslice -input ./assets/ship.png -output resources/ship.go -package resources -var ShipPng
//go:generate file2byteslice -input ./assets/alien.png -output resources/alien.go -package resources -var AlienPng
//go:generate file2byteslice -input config.json -output resources/config.go -package resources -var ConfigJson
func main() {
	if err := NewGame(LoadConfig()).Run(); err != nil {
		log.Fatal(err)
	}
}
