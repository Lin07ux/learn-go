package main

import (
	"bytes"
	"encoding/json"
	"github.com/lin07ux/learn-go/game-ebitengine/resources"
	"image/color"
	"log"
)

type Config struct {
	ScreenWidth  int
	ScreenHeight int
	Title        string
	BgColor      color.RGBA

	ShipSpeedFactor  float64
	ShipFireInterval int64

	BulletWidth       int
	BulletHeight      int
	BulletSpeedFactor float64
	BulletColor       color.RGBA
	MaxBulletNum      int

	AlienSpeedFactor float64

	TitleFontSize  int
	CommonFontSize int
}

func LoadConfig() *Config {
	var cfg Config
	if err := json.NewDecoder(bytes.NewReader(resources.ConfigJson)).Decode(&cfg); err != nil {
		log.Fatalf("json.Decode failed: %v\n", err)
	}

	return &cfg
}
