// the gomatrix library
package main

import (
	"fmt"

	"github.com/MH15/gomatrix/pkg/matrix"
	"github.com/MH15/gomatrix/pkg/vector"
)

func main() {
	fmt.Println("starting...")

	vec1 := vector.Vec(1, 2, 3)
	vec2 := vector.Vec(4, 5, 6)
	vec3 := vector.Vec(7, 8, 9)

	a := matrix.Mat(vec1, vec2, vec3)

	vecA := vector.Vec(3, 1, 3, 6)
	vecB := vector.Vec(1, 2, 7, 6)
	vecC := vector.Vec(1, 4, 5, 7)
	vecD := vector.Vec(0, 0, 1, 2)
	b := matrix.Mat(vecA, vecB, vecC, vecD)

	ar := matrix.RREF(a)
	br := matrix.RREF(b)

	fmt.Println(a.ToString())
	fmt.Println(ar.ToString())

	fmt.Println(b.ToString())
	fmt.Println(br.ToString())

}
