package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Firework struct {
	positionX, positionY, velocity int32

	radius float32

	color rl.Color

	particalAmount int32

	particalSpeed int32
}

func NewFirework(posx int32) Firework {

	minVel := (HEIGHT * 1 / 22)
	maxVel := (HEIGHT * 1 / 20)
	vel := -(rand.Int31n(int32(maxVel-minVel)) + int32(minVel))

	return Firework{
		positionX:      posx,
		positionY:      HEIGHT + 10,
		velocity:       vel,
		radius:         4.0,
		color:          FireworkColors[rand.Int31n((int32)(len(FireworkColors)))],
		particalAmount: rand.Int31n(50) + 50,
		particalSpeed:  rand.Int31n(5) + 5,
	}

}

func (f Firework) Show() {
	rl.DrawCircle(f.positionX, f.positionY, f.radius, f.color)
}

func (f *Firework) Update() {
	f.velocity += GRAVITY
	f.positionY += f.velocity
}

func (f Firework) ShouldExplode() bool {
	return f.velocity >= 5
}

func (f *Firework) GetParticals() []Partical {
	newParticals := make([]Partical, f.particalAmount)

	if rand.Float32() > 0.5 {
		// circular explosition
		for i := range newParticals {
			newParticals[i] = NewPartical(f.positionX, f.positionY, f.particalSpeed, f.particalSpeed, f.color)
		}
	} else {
		// chaotic explosition
		for i := range newParticals {
			newParticals[i] = NewPartical(f.positionX, f.positionY, f.particalSpeed-5, f.particalSpeed, f.color)
		}
	}

	return newParticals
}
