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

func (b *Bullet) OutOfScreen() bool {
	return b.y < -float64(b.height)
}

func (b *Bullet) Hit(target Rectangle) bool {
	width, height := target.Size()
	left, top := target.Coordinate()
	right, bottom := left+float64(width), top+float64(height)

	// 左上角
	x, y := b.x, b.y
	if top < y && y < bottom && left < x && x < right {
		return true
	}

	// 右上角
	x, y = b.x+float64(b.width), b.y
	if top < y && y < bottom && left < x && x < right {
		return true
	}

	// 左下角
	x, y = b.x, b.y+float64(b.height)
	if top < y && y < bottom && left < x && x < right {
		return true
	}

	// 右下角
	x, y = b.x+float64(b.width), b.y+float64(b.height)
	if top < y && y < bottom && left < x && x < right {
		return true
	}

	return false
}
