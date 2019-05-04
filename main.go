package main

import (
	"fmt"
)

func main() {
	fmt.Println("starting...")

	vec1 := Vec(1, 2, 3)
	vec2 := Vec(4, 5, 6)
	vec3 := Vec(7, 8, 9)

	a := Mat(vec1, vec2, vec3)

	vecA := Vec(3, 1, 3, 6)
	vecB := Vec(1, 2, 7, 6)
	vecC := Vec(1, 4, 5, 7)
	vecD := Vec(0, 0, 1, 2)
	b := Mat(vecA, vecB, vecC, vecD)

	ar := RREF(a)
	br := RREF(b)

	fmt.Println(a.toString())
	fmt.Println(ar.toString())

	fmt.Println(b.toString())
	fmt.Println(br.toString())

}
