package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
)

type mode int

const (
	modePending mode = iota
	modePlaying
	modeFinished
)

const (
	resultWin  = "You Win!"
	resultFail = "Game Over!"
)

var (
	titleFont  font.Face
	commonFont font.Face
)

type Game struct {
	mode
	ship       *Ship
	config     *Config
	aliens     map[*Alien]struct{}
	bullets    map[*Bullet]struct{}
	collisions int // 碰撞次数（子弹、飞船、屏幕下边界与外星人飞碟的碰撞）
	result     string
}

func NewGame(config *Config) *Game {
	return &Game{
		mode:   modePending,
		ship:   NewShip(config.ShipSpeedFactor, config.ScreenWidth, config.ScreenHeight),
		config: config,
	}
}

func (g *Game) Run() error {
	ebiten.SetWindowTitle(g.config.Title)
	ebiten.SetWindowSize(g.config.ScreenWidth, g.config.ScreenHeight)

	if err := g.createFonts(); err != nil {
		return err
	}

	return ebiten.RunGame(g)
}

func (g *Game) createFonts() error {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		return err
	}

	const dpi = 72
	titleFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		DPI:     dpi,
		Size:    float64(g.config.TitleFontSize),
		Hinting: font.HintingFull,
	})
	if err != nil {
		return err
	}

	commonFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		DPI:     dpi,
		Size:    float64(g.config.CommonFontSize),
		Hinting: font.HintingFull,
	})
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) Update() error {
	switch g.mode {
	case modePending:
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.result = ""
			g.collisions = 0
			g.bullets = make(map[*Bullet]struct{})
			if err := g.createAliens(3); err != nil {
				return err
			}
			g.mode = modePlaying
		}
	case modeFinished:
		if ebiten.IsKeyPressed(ebiten.KeyR) {
			g.mode = modePending
		}
	case modePlaying:
		g.updateShip()
		g.updateAliens()
		g.updateBullets()
		g.checkCollision()
		g.checkResult()
	}

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
		// 跑到屏幕外
		if alien.OutOfScreen(g.config.ScreenHeight) {
			g.collisions++
			delete(g.aliens, alien)
			continue
		}

		// 与飞船碰撞
		if checkElementCollision(alien, g.ship) {
			g.collisions++
			delete(g.aliens, alien)
			continue
		}

		// 被子弹击中
		for bullet := range g.bullets {
			if checkElementCollision(alien, bullet) {
				delete(g.aliens, alien)
				delete(g.bullets, bullet)
			}
		}
	}
}

func (g *Game) checkResult() {
	if g.collisions >= 3 {
		g.result = resultFail
	} else if len(g.aliens) == 0 {
		g.result = resultWin
	}

	if len(g.result) > 0 {
		g.mode = modeFinished
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.config.BgColor)

	switch g.mode {
	case modePending:
		g.drawText(screen, []string{"ALIEN INVASION"}, []string{"PRESS SPACE KEY TO START"})
	case modeFinished:
		g.drawText(screen, []string{g.result}, []string{"PRESS R KEY TO RESTART"})
	case modePlaying:
		g.ship.Draw(screen)
		for bullet := range g.bullets {
			bullet.Draw(screen)
		}
		for alien := range g.aliens {
			alien.Draw(screen)
		}
	}
}

func (g *Game) drawText(screen *ebiten.Image, titles, content []string) {
	y := 3 * g.config.TitleFontSize

	for _, t := range titles {
		x := (g.config.ScreenWidth - len(t)*g.config.TitleFontSize) / 2
		y += g.config.TitleFontSize * 18 / 10
		text.Draw(screen, t, titleFont, x, y, color.White)
	}

	y += g.config.TitleFontSize + 3*g.config.CommonFontSize
	for _, c := range content {
		x := (g.config.ScreenWidth - len(c)*g.config.CommonFontSize) / 2
		y += g.config.CommonFontSize * 15 / 10
		text.Draw(screen, c, commonFont, x, y, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) createAliens(rows int) error {
	alien := NewAlien(0, 0, g.config.AlienSpeedFactor)

	// 屏幕左右两侧、每个外星人飞碟两侧，各留半个外星人飞碟宽度的空白
	alienWidth, alienHeight := alien.Size()
	aliensNum := (g.config.ScreenWidth - alienWidth) / (2 * alienWidth)
	if aliensNum <= 0 {
		return fmt.Errorf("game: screen width is too small: %d, minimum need: %d\n", g.config.ScreenWidth, 3*alienWidth)
	}

	top := 5.0
	g.aliens = make(map[*Alien]struct{})
	for row := 0; row < rows; row++ {
		top -= float64(alienHeight) * 1.5
		for i := 0; i < aliensNum; i++ {
			// alienWidth/2 + i * alienWidth + alienWidth/2
			alien = NewAlien(float64(2*i*alienWidth+alienWidth), top, g.config.AlienSpeedFactor)
			g.aliens[alien] = struct{}{}
		}
	}

	return nil
}
