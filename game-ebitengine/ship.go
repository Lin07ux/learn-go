package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

type Ship struct {
	image  *ebiten.Image
	width  int
	height int
}

func NewShip() *Ship {
	img, _, err := ebitenutil.NewImageFromFile("./assets/ship.png")
	if err != nil {
		log.Fatalf("ship: load failed: %v\n", err)
	}

	width, height := img.Size()

	return &Ship{
		image:  img,
		width:  width,
		height: height,
	}
}

func (s *Ship) Draw(screen *ebiten.Image, screenWidth, screenHeight int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(screenWidth-s.width)/2, float64(screenHeight-s.height))
	screen.DrawImage(s.image, op)
}
