package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	input  *Input
	config *Config
}

func NewGame() *Game {
	return &Game{
		input:  &Input{"Hello, World"},
		config: LoadConfig(),
	}
}

func (g *Game) Update() error {
	g.input.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.config.BgColor)
	ebitenutil.DebugPrint(screen, g.input.msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / 2, outsideHeight / 2
}

func (g *Game) Title() string {
	return g.config.Title
}

func (g *Game) ScreenSize() (width, height int) {
	return g.config.ScreenWidth, g.config.ScreenHeight
}
