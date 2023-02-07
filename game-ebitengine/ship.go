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
	x      float64
	y      float64
}

func NewShip(screenWidth, screenHeight int) *Ship {
	img, _, err := ebitenutil.NewImageFromFile("./assets/ship.png")
	if err != nil {
		log.Fatalf("ship: load failed: %v\n", err)
	}

	width, height := img.Size()
	if screenWidth < width*2 {
		log.Fatalf("ship: the width is too wide: %d/%d\n", width, screenWidth)
	}

	if screenHeight < height*2 {
		log.Fatalf("ship: the height is too high: %d/%d\n", height, screenHeight)
	}

	return &Ship{
		image:  img,
		width:  width,
		height: height,
		x:      float64(screenWidth-width) / 2,
		y:      float64(screenHeight - height),
	}
}

func (s *Ship) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.x, s.y)
	screen.DrawImage(s.image, op)
}
