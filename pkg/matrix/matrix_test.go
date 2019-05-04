package matrix

import (
	"fmt"
	"testing"
)

// func TestPerformance(t *testing.T) {
// 	bigZ := Zeros(12345, 12345)
// 	bigI := Identity(12345, 12345)
// 	bigZ.subtract(bigI)

// 	m, _ := bigZ.Dims()
// 	if false {
// 		t.Errorf("Size was incorrect, got: %d, want: %d", m, 12345)
// 	}
// }

func TestZeros(t *testing.T) {
	// make matrices using Zero
	z1x1 := Zeros(1, 1)
	z1x4 := Zeros(1, 4)
	z4x1 := Zeros(4, 1)
	z4x4 := Zeros(4, 4)
	z999x999 := Zeros(999, 999)

	// make el field vectors manually
	el1x1 := []float64{0}
	el4x1 := []float64{0, 0, 0, 0}
	el1x4 := []float64{0, 0, 0, 0}
	el4x4 := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	// check if correct
	if !Equal(z1x1.el, el1x1) {
		panic("Matrix el field doesn't match float64 slice")
	}
	if !Equal(z1x4.el, el1x4) {
		panic("Matrix el field doesn't match float64 slice")
	}
	if !Equal(z4x1.el, el4x1) {
		panic("Matrix el field doesn't match float64 slice")
	}
	if !Equal(z4x4.el, el4x4) {
		panic("Matrix el field doesn't match float64 slice")
	}

	m, n := z999x999.Dims()
	if m != 999 || n != 999 {
		panic("Matrix dimensions are not correct.")
	}
}

func TestOnes(t *testing.T) {
	// make matrices using Zero
	o1x1 := Ones(1, 1)
	o1x4 := Ones(1, 4)
	o4x1 := Ones(4, 1)
	o4x4 := Ones(4, 4)
	o999x999 := Ones(999, 999)

	// make el field vectors manually
	el1x1 := []float64{1}
	el4x1 := []float64{1, 1, 1, 1}
	el1x4 := []float64{1, 1, 1, 1}
	el4x4 := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	// check if correct
	if !Equal(o1x1.el, el1x1) {
		panic("Matrix el field doesn't match float64 slice")
	}
	if !Equal(o1x4.el, el1x4) {
		panic("Matrix el field doesn't match float64 slice")
	}
	if !Equal(o4x1.el, el4x1) {
		panic("Matrix el field doesn't match float64 slice")
	}
	if !Equal(o4x4.el, el4x4) {
		panic("Matrix el field doesn't match float64 slice")
	}

	m, n := o999x999.Dims()
	if m != 999 || n != 999 {
		panic("Matrix dimensions are not correct.")
	}
}

func TestDiagonal(t *testing.T) {
	v := Vec(1, 2, 3)
	d := Diagonal(v)

	el := Vec(1, 0, 0, 0, 2, 0, 0, 0, 3)

	if !Equal(d.el, el) {
		panic("Diagonal matrix not generated correctly.")
	}

}

func TestTranspose(t *testing.T) {
	m := Mat(Vec(1, 2, 3), Vec(4, 5, 6), Vec(7, 8, 9))
	mt := Vec(1, 4, 7, 2, 5, 8, 3, 6, 9)
	m.transpose()

	fmt.Println(m.toStringPlain())
	if !Equal(m.el, mt) {
		panic("Transpose of matrix does not match what it should be.")
	}
}
