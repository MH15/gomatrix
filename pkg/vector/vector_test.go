package vector

import (
	"fmt"
	"testing"
)

func TestLen(t *testing.T) {
	vec1 := Vec(1, 2, 3, 4, 5)

	length := vec1.Len()

	if length != 5 {
		t.Errorf("Length was incorrect, got: %d, want: %d", length, 5)
	}
}

func TestAt(t *testing.T) {
	vec1 := Vec(1, 2, 3, 4, 5)

	pos3 := vec1.At(3)

	if pos3 != 4 {
		t.Errorf("Length was incorrect, got: %f, want: %d", pos3, 4)
	}
}

func TestDotStatic(t *testing.T) {
	vec1 := Vec(1, 2, 3, 4, 5)
	vec2 := Vec(5, 4, 3, 2, 1)

	dot := Dot(vec1, vec2)
	if dot != 35 {
		t.Errorf("Dot product was incorrect, got: %f, want: %d", dot, 35)
	}
}

func TestEqual(t *testing.T) {
	vec1 := Vec(1, 2, 3, 4, 5)
	vec2 := Vec(1, 2, 3, 4, 5)

	equals := Equal(vec1, vec2)
	fmt.Println(equals)
	if equals != true {
		t.Errorf("Vectors should be equal. They were not.")
	}
}

func TestEqualSize(t *testing.T) {
	vec1 := Vec(1, 2, 3, 4, 5)
	vec2 := Vec(1, 2, 3, 4, 5)

	equals := EqualSize(vec1, vec2)
	if equals != true {
		t.Errorf("Vectors should be equal size. They were not.")
	}
}

/// OPERATIONS
func TestAdd(t *testing.T) {
	vec1 := Vec(1, 2, 3, 4, 5)
	vec2 := Vec(1, 2, 3, 4, 5)
	result := Vec(2, 4, 6, 8, 10)
	vec1.add(vec2)

	if !Equal(vec1, result) {
		t.Errorf("Vector after add was %s, should be %s.", arrayToString(vec1, ","), arrayToString(result, ","))
	}

}
func TestSubtract(t *testing.T) {
	vec1 := Vec(1, 2, 3, 4, 5)
	vec2 := Vec(1, 2, 3, 4, 5)
	result := Vec(0, 0, 0, 0, 0)
	vec1.subtract(vec2)

	if !Equal(vec1, result) {
		t.Errorf("Vector after subtract was %s, should be %s.", arrayToString(vec1, ","), arrayToString(result, ","))
	}

}
func TestScalarMult(t *testing.T) {
	vec1 := Vec(1, 2, 3, 4, 5)
	scalar := 16.0
	result := Vec(16, 32, 48, 64, 80)
	vec1.scalarMult(scalar)

	if !Equal(vec1, result) {
		t.Errorf("Vector after scalar mult was %s, should be %s.", arrayToString(vec1, ","), arrayToString(result, ","))
	}
}
