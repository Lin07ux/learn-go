package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ship   *Ship
	input  *Input
	config *Config
}

func NewGame() *Game {
	config := LoadConfig()

	return &Game{
		ship:   NewShip(config.ScreenWidth, config.ScreenHeight),
		input:  &Input{"Hello, World"},
		config: config,
	}
}

func (g *Game) Run() error {
	ebiten.SetWindowTitle(g.config.Title)
	ebiten.SetWindowSize(g.config.ScreenWidth, g.config.ScreenHeight)

	return ebiten.RunGame(g)
}

func (g *Game) Update() error {
	var deltas float64
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		deltas = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		deltas = 1
	}

	g.ship.Move(deltas * g.config.ShipSpeedFactor)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.config.BgColor)
	g.ship.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
