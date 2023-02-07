package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	game := NewGame()

	ebiten.SetWindowTitle(game.Title())
	ebiten.SetWindowSize(game.ScreenSize())
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
