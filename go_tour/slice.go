package main

import (
	"fmt"
	"strings"
)

func SliceExample() {
	fmt.Println("------------------------")
	fmt.Println("Slice Example:")

	arr := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = arr[1:4] // [3 5 7]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names) // [John Paul George Ringo]

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"

	// A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	fmt.Println(a, b)  // [John XXX] [XXX George]
	fmt.Println(names) // [John XXX George Ringo]

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	// r[0] = false
	fmt.Println(r)

	x := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(x)

	// The default is zero for the low bound and the length of the slice for the high bound.
	// For the array
	// var a [10]int
	// these slice expressions are equivalent:
	/**
		a[0:10]
		a[:10]
		a[0:]
		a[:]
	**/

	// Slice length and capacity
	z := []int{2, 3, 5, 7, 11, 13}
	printSlice(z)

	// Slice the slice to give it zero length.
	z = z[:0]
	printSlice(z)

	// Extend its length.
	z = z[:4]
	printSlice(z)

	// Drop its first two values.
	z = z[2:]
	printSlice(z)

	// Nil slices
	var y []int
	fmt.Println(y, len(y), cap(y))
	if y == nil {
		fmt.Println("nil!")
	}

	// Creating a slice with make
	n := make([]int, 5)
	printSlice1("n", n)

	m := make([]int, 0, 5)
	printSlice1("m", m)

	v := m[:2]
	printSlice1("v", v)

	t := v[2:5]
	printSlice1("t", t)

	// Slices of slices
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// Appending to a slice
	var slice1 []int
	printSlice2(slice1)

	// append works on nil slices.
	slice1 = append(slice1, 0)
	printSlice2(slice1)

	// The slice grows as needed.
	slice1 = append(slice1, 1)
	printSlice2(slice1)

	// We can add more than one element at a time.
	slice1 = append(slice1, 2, 3, 4)
	printSlice2(slice1)

	// Range
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// Range continued
	/**
		for i, _ := range pow
		for _, value := range pow
		for i := range pow
	**/
	pow1 := make([]int, 10)
	for i := range pow1 {
		pow1[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow1 {
		fmt.Printf("%d\n", value)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func printSlice2(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
