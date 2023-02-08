package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"
	_ "image/png"
	"log"
)

type Ship struct {
	image       *ebiten.Image
	width       int
	height      int
	speedFactor float64
	x           float64
	y           float64
}

func NewShip(speedFactor float64, screenWidth, screenHeight int) *Ship {
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

	if speedFactor <= 0 {
		speedFactor = 1
	}

	return &Ship{
		image:       img,
		width:       width,
		height:      height,
		speedFactor: speedFactor,
		x:           float64(screenWidth-width) / 2,
		y:           float64(screenHeight - height),
	}
}

func (s *Ship) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(s.x, s.y)
	screen.DrawImage(s.image, op)
}

func (s *Ship) Move(deltas int, screenWidth int) {
	if deltas == 0 {
		return
	}

	minX, maxX := -float64(s.width)/2, float64(screenWidth)-float64(s.width)/2
	s.x += float64(deltas) * s.speedFactor

	if s.x < minX {
		s.x = minX
	} else if s.x > maxX {
		s.x = maxX
	}
}

func (s *Ship) FireBullet(width, height int, speedFactor float64, bgColor color.RGBA) *Bullet {
	rect := image.Rect(0, 0, width, height)
	img := ebiten.NewImageWithOptions(rect, nil)
	img.Fill(bgColor)

	return NewBullet(img, speedFactor, s.x+float64(s.width-width)/2, s.y-float64(height))
}
