package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	WIDTH  = 1000
	HEIGHT = 800
	TITLE  = "FireworksDisplay"

	FPS = 60

	GRAVITY = 1

	FIREWORK_DISPLAY = 0.1
)

var BgColor rl.Color = rl.Black

var FireworkColors []rl.Color = []rl.Color{
	rl.White,
	rl.Red,
	rl.Green,
	rl.Blue,
	rl.Lime,
	rl.Yellow,
	rl.Pink,
	rl.Purple,
	rl.Gold,
	rl.Maroon,
	rl.Magenta,
	rl.Violet,
	rl.SkyBlue,
}
