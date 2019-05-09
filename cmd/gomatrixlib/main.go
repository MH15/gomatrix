// the gomatrix library
package main

import (
	"fmt"

	mat "github.com/MH15/gomatrix/pkg/matrix"
	vec "github.com/MH15/gomatrix/pkg/vector"
)

func main() {
	fmt.Println("starting...")

	vec1 := vec.Vec(1, 2, 3)
	vec2 := vec.Vec(4, 5, 6)
	vec3 := vec.Vec(7, 8, 9)

	a := mat.Mat(vec1, vec2, vec3)

	vecA := vec.Vec(3, 1, 3, 6)
	vecB := vec.Vec(1, 2, 7, 6)
	vecC := vec.Vec(1, 4, 5, 7)
	vecD := vec.Vec(0, 0, 1, 2)
	b := mat.Mat(vecA, vecB, vecC, vecD)

	ar := mat.RREF(a)
	br := mat.RREF(b)

	fmt.Println(a.ToString())
	fmt.Println(ar.ToString())

	fmt.Println(b.ToString())
	fmt.Println(br.ToString())

}
