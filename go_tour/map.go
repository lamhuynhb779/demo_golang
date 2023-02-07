package main

import (
	"fmt"

	"golang.org/x/tour/wc"
)

type Location struct {
	Lat, Long float64
}

var m map[string]Location

var m1 = map[string]Location{
	"HCM": Location{
		40.68433, -74.39967,
	},
	"HN": Location{
		37.42202, -122.08408,
	},
}

func MapExample() {
	fmt.Println("------------------------")
	fmt.Println("Map Example:")

	m = make(map[string]Location)

	m["Bell Labs"] = Location{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])

	// Map literals
	fmt.Println(m1)

	// Mutating Maps
	m2 := make(map[string]int)

	m2["Answer"] = 42
	fmt.Println("The value:", m2["Answer"])

	m2["Answer"] = 48
	fmt.Println("The value:", m2["Answer"])

	delete(m2, "Answer")
	fmt.Println("The value:", m2["Answer"])

	v, ok := m2["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	// Exercise: Maps
	wc.Test(WordCount)
	fmt.Println()
}

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}
