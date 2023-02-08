package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ship    *Ship
	config  *Config
	bullets map[*Bullet]struct{}
}

func NewGame() *Game {
	config := LoadConfig()

	return &Game{
		ship:    NewShip(config.ShipSpeedFactor, config.ScreenWidth, config.ScreenHeight),
		config:  config,
		bullets: make(map[*Bullet]struct{}),
	}
}

func (g *Game) Run() error {
	ebiten.SetWindowTitle(g.config.Title)
	ebiten.SetWindowSize(g.config.ScreenWidth, g.config.ScreenHeight)

	return ebiten.RunGame(g)
}

func (g *Game) Update() error {
	g.updateShip()
	g.updateBullets()
	return nil
}

func (g *Game) updateShip() {
	var deltas int
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		deltas = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		deltas = 1
	}

	g.ship.Move(deltas, g.config.ScreenWidth)
}

func (g *Game) updateBullets() {
	for bullet := range g.bullets {
		bullet.Move(-1)
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		bullet := g.ship.FireBullet(g.config.BulletWidth, g.config.BulletHeight, g.config.BulletSpeedFactor, g.config.BulletColor)
		g.bullets[bullet] = struct{}{}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.config.BgColor)
	g.ship.Draw(screen)
	for bullet := range g.bullets {
		bullet.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
