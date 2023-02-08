package main

import (
	"encoding/json"
	"image/color"
	"log"
	"os"
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
}

func LoadConfig() *Config {
	f, err := os.Open("./config.json")
	if err != nil {
		log.Fatalf("os.Open failed: %v\n", err)
	}

	var cfg Config
	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil {
		log.Fatalf("json.Decode failed: %v\n", err)
	}

	return &cfg
}
