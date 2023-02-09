package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lin07ux/learn-go/game-ebitengine/resources"
	"log"
)

type Alien struct {
	image       *ebiten.Image
	width       int
	height      int
	x           float64
	y           float64
	speedFactor float64
}

func NewAlien(x, y, speedFactor float64) *Alien {
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(resources.AlienPng))
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()

	return &Alien{
		image:       img,
		width:       width,
		height:      height,
		x:           x,
		y:           y,
		speedFactor: speedFactor,
	}
}

func (a *Alien) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(a.x, a.y)
	screen.DrawImage(a.image, op)
}

func (a *Alien) Move(deltas int) {
	a.y += float64(deltas) * a.speedFactor
}

func (a *Alien) OutOfScreen(screenHeight int) bool {
	return a.y >= float64(screenHeight)
}

func (a *Alien) Size() (width, height int) {
	return a.width, a.height
}

func (a *Alien) Coordinate() (x, y float64) {
	return a.x, a.y
}
