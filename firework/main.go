package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello")

	rand.Seed(time.Now().Unix())

	var game Game
	game.Init()
	defer game.Close()

	game.Loop()
}
