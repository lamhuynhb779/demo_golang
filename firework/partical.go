package main

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Partical struct {
	positionX, positionY, velocityX, velocityY float64

	radius float32

	color rl.Color

	alpha int32
}

func NewPartical(x, y, minSpeed, maxSpeed int32, color rl.Color) Partical {
	angle := rand.Float64() * (2 * math.Pi)

	var speed float64

	if minSpeed == maxSpeed {
		speed = float64(minSpeed)
	} else {
		speed = float64(rand.Int31n(maxSpeed-minSpeed) + minSpeed)
	}

	velx := math.Cos(angle) * speed
	vely := math.Sin(angle) * speed

	return Partical{
		positionX: float64(x),
		positionY: float64(y),
		color:     color,
		radius:    1.0,
		velocityX: velx,
		velocityY: vely,
		alpha:     255,
	}

}

func (p *Partical) Update() {
	p.positionX += p.velocityX
	p.positionY += p.velocityY

	p.alpha -= 10

	if p.alpha >= 0 {
		p.color.A = uint8(p.alpha)
	}
}

func (p Partical) Show() {
	rl.DrawCircle(int32(p.positionX), int32(p.positionY), p.radius, p.color)
}

func (p Partical) HasFizzled() bool {
	return p.alpha <= 0
}
