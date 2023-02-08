package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	image       *ebiten.Image
	width       int
	height      int
	x           float64
	y           float64
	speedFactor float64
}

func NewBullet(img *ebiten.Image, speedFactor, x, y float64) *Bullet {
	width, height := img.Size()

	return &Bullet{
		image:       img,
		width:       width,
		height:      height,
		x:           x,
		y:           y,
		speedFactor: speedFactor,
	}
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.x, b.y)
	screen.DrawImage(b.image, op)
}

func (b *Bullet) Move(deltas int) {
	b.y += float64(deltas) * b.speedFactor
}
