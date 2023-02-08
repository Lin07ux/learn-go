package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Game struct {
	ship    *Ship
	config  *Config
	aliens  map[*Alien]struct{}
	bullets map[*Bullet]struct{}
}

func NewGame() *Game {
	config := LoadConfig()

	game := &Game{
		ship:    NewShip(config.ShipSpeedFactor, config.ScreenWidth, config.ScreenHeight),
		config:  config,
		aliens:  make(map[*Alien]struct{}),
		bullets: make(map[*Bullet]struct{}),
	}

	game.createAliens(2)

	return game
}

func (g *Game) Run() error {
	ebiten.SetWindowTitle(g.config.Title)
	ebiten.SetWindowSize(g.config.ScreenWidth, g.config.ScreenHeight)

	return ebiten.RunGame(g)
}

func (g *Game) Update() error {
	g.updateShip()
	g.updateAliens()
	g.updateBullets()
	g.checkCollision()
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
		if bullet.OutOfScreen() {
			delete(g.bullets, bullet)
		}
	}

	cfg := g.config
	if ebiten.IsKeyPressed(ebiten.KeySpace) && len(g.bullets) < cfg.MaxBulletNum && g.ship.LastFireAfter(cfg.ShipFireInterval) {
		bullet := g.ship.FireBullet(cfg.BulletWidth, cfg.BulletHeight, cfg.BulletSpeedFactor, cfg.BulletColor)
		g.bullets[bullet] = struct{}{}
	}
}

func (g *Game) updateAliens() {
	for alien := range g.aliens {
		alien.Move(1)
	}
}

func (g *Game) checkCollision() {
	for alien := range g.aliens {
		for bullet := range g.bullets {
			if bullet.Hit(alien) {
				delete(g.aliens, alien)
				delete(g.bullets, bullet)
			}
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.config.BgColor)
	g.ship.Draw(screen)

	for bullet := range g.bullets {
		bullet.Draw(screen)
	}

	for alien := range g.aliens {
		alien.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) createAliens(rows int) {
	alien := NewAlien(0, 0, g.config.AlienSpeedFactor)

	// 屏幕左右两侧、每个外星人飞碟两侧，各留半个外星人飞碟宽度的空白
	alienWidth, alienHeight := alien.Size()
	aliensNum := (g.config.ScreenWidth - alienWidth) / (2 * alienWidth)
	if aliensNum <= 0 {
		log.Fatalf("game: screen width is too small: %d, minimum need: %d\n", g.config.ScreenWidth, 3*alienWidth)
	}

	top := 5.0
	for row := 0; row < rows; row++ {
		top += float64(row*alienHeight) * 1.5
		for i := 0; i < aliensNum; i++ {
			// alienWidth/2 + i * alienWidth + alienWidth/2
			alien = NewAlien(float64(2*i*alienWidth+alienWidth), top, g.config.AlienSpeedFactor)
			g.aliens[alien] = struct{}{}
		}
	}
}
