package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func StructExample() {
	fmt.Println("------------------------")
	fmt.Println("Struct Example:")

	fmt.Println(Vertex{1, 2})

	v := Vertex{3, 4}
	v.X = 5
	fmt.Println(v)

	p := &v
	p.X = 1e9

	fmt.Println(v)

	v2 := Vertex{X: 1} // Y:0 is implicit
	v3 := Vertex{}     // X:0 Y:0
	q := &Vertex{1, 2} // has type *Vertex

	fmt.Println(v2, v3, q, *q, q.X, (*q).X)
}
