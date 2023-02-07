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
	return &Game{
		ship:   NewShip(),
		input:  &Input{"Hello, World"},
		config: LoadConfig(),
	}
}

func (g *Game) Run() error {
	ebiten.SetWindowTitle(g.config.Title)
	ebiten.SetWindowSize(g.config.ScreenWidth, g.config.ScreenHeight)

	return ebiten.RunGame(g)
}

func (g *Game) Update() error {
	g.input.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.config.BgColor)
	g.ship.Draw(screen, g.config.ScreenWidth, g.config.ScreenHeight)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
