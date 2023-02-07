package main

import (
	"fmt"
	"math/cmplx"
)

// Constants cannot be declared using the := syntax.
const Pi = 3.14

var c, java, python bool

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// x int, y int => x, y int
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// Constant
	const Truth = true

	fmt.Println(add(5, 6))
	fmt.Println(swap("5", "6"))
	fmt.Println(split(9))

	var i int
	fmt.Println(i, c, java, python)

	var a, b int = 1, 2
	// Short variable declarations
	// can be used in place of a var declaration with implicit type.
	// := only can use inside of function
	k := 3
	fmt.Println(a, b, k)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	PointerExample()
	StructExample()
	ArrayExample()
	SliceExample()
	MapExample()
	FunctionValuesExample()
}
