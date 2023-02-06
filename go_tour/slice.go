package main

import "fmt"

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
}
