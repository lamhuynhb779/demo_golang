package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	running bool

	displayOn bool

	fireworks []Firework

	particals []Partical

	liftSounds    []rl.Sound
	explodeSounds []rl.Sound
}

func (g *Game) Init() {
	rl.InitWindow(WIDTH, HEIGHT, TITLE)
	rl.SetTargetFPS(FPS) // optional
	rl.InitAudioDevice()

	g.liftSounds = []rl.Sound{
		rl.LoadSound("sfx/lift1.mp3"),
		rl.LoadSound("sfx/lift2.mp3"),
		rl.LoadSound("sfx/lift3.mp3"),
		rl.LoadSound("sfx/lift4.mp3"),
	}

	g.explodeSounds = []rl.Sound{
		rl.LoadSound("sfx/explode1.mp3"),
		rl.LoadSound("sfx/explode2.mp3"),
		rl.LoadSound("sfx/explode3.mp3"),
		rl.LoadSound("sfx/explode4.mp3"),
	}

	g.running = true

	g.fireworks = make([]Firework, 0)

	g.particals = make([]Partical, 0)
}

func (g *Game) Close() {
	rl.CloseWindow()
	rl.CloseAudioDevice()
}

func (g *Game) Loop() {
	for g.running {
		if rl.WindowShouldClose() {
			g.running = false
			return
		}
		if rl.IsKeyPressed(rl.KeySpace) {
			g.displayOn = !g.displayOn
		}

		g.update()

		g.draw()
	}
}

func (g *Game) update() {

	if g.displayOn && rand.Float32() < FIREWORK_DISPLAY {
		g.CreateFirework(rand.Int31n(WIDTH))
	}

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		g.CreateFirework(rl.GetMouseX())
	}

	for i := len(g.fireworks) - 1; i >= 0; i-- {
		g.fireworks[i].Update()

		if g.fireworks[i].ShouldExplode() {

			rl.PlaySound(g.explodeSounds[rand.Int31n(int32(len(g.explodeSounds)))])

			newParticals := g.fireworks[i].GetParticals()
			g.particals = append(g.particals, newParticals...)

			g.fireworks[i] = g.fireworks[len(g.fireworks)-1]
			/**
			* 4,3,2,1
			* 4,1,2,1
			* 4,1,2
			**/
			g.fireworks = g.fireworks[:len(g.fireworks)-1]

		}
	}

	for i := len(g.particals) - 1; i >= 0; i-- {
		g.particals[i].Update()

		if g.particals[i].HasFizzled() {
			g.particals[i] = g.particals[len(g.particals)-1]
			g.particals = g.particals[:len(g.particals)-1]
		}

	}

}

func (g *Game) draw() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(BgColor)

	for i := range g.fireworks {
		g.fireworks[i].Show()
	}

	for i := range g.particals {
		g.particals[i].Show()
	}
}

func (g *Game) CreateFirework(x int32) {
	rl.PlaySound(g.liftSounds[rand.Int31n(int32(len(g.liftSounds)))])
	g.fireworks = append(g.fireworks, NewFirework(x))
}
