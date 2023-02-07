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
	ebitenutil.DebugPrint(screen, g.input.msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / 2, outsideHeight / 2
}
