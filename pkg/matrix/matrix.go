// Package matrix here
package matrix

import (
	"fmt"
	"strings"

	"github.com/MH15/gomatrix/pkg/vector"
)

// Matrix is the basic matrix interface type.
type Matrix struct {
	el     vector.Vector
	stride int
	class  Class
}

// Class is a future-proof field to use when we add logical functionality
type Class int

const (
	// Logical true (1) or false (0)
	Logical Class = iota
	// FloatingPoint in float64 format
	FloatingPoint
)

// Mat creates a matrix (nxm) of equal sized vector rows
func Mat(v ...vector.Vector) Matrix {
	mat := Matrix{}
	if vector.EqualSize(v...) {
		// row and col sizes
		mat = Zeros(len(v), len(v[0]))
		// flatten v into the matrix
		iterateRows(&mat, func(i int, j int) {
			mat.set(i, j, v[i].At(j))
		})
	}
	return mat
}

// Dims returns the dimensions of the matrix
func (a *Matrix) Dims() (int, int) {
	return len(a.el) / a.stride, a.stride
}

// At returns the value at row i, column j.
// Panics if i or j are out of bounds for the m by n matrix.
func (a *Matrix) at(i, j int) float64 {
	panicMatrixSize(a, i, j)
	return a.el[i*a.stride+j]
}

// Sets the value at i, j
// Panics if i or j are out of bounds for the m by n matrix
func (a *Matrix) set(i, j int, val float64) {
	panicMatrixSize(a, i, j)
	// fmt.Printf("Setting %f at pos(%d,%d)\n", val, i, j)
	a.el[i*a.stride+j] = val
}

// Return the rows of a matrix
func (a *Matrix) rows() []vector.Vector {
	r := make([]vector.Vector, len(a.el)/a.stride)
	i := 0
	for i < len(a.el) {
		r[i/a.stride] = vector.Vec(a.el[i : i+a.stride]...)
		i += a.stride
	}
	return r
}

// Return the columns of a matrix
func (a *Matrix) columns() [][]float64 {
	c := make([][]float64, a.stride)
	m, n := a.Dims()
	for j := 0; j < n; j++ {
		c[j] = make([]float64, m)
		for i := 0; i < m; i++ {
			c[j][i] = a.at(i, j)
		}
	}
	return c
}

// call a function on each position in the matrix
// TODO: could be refactored to include concurrency
func iterateRows(a *Matrix, cb func(i int, j int)) {
	m, n := a.Dims()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cb(i, j)
		}
	}
}

/// PANICS
func panicMatrixSize(a *Matrix, i, j int) {
	m, n := a.Dims()
	if i > m || j > n {
		panic("ErrIndexOutOfRange")
	}
}

func panicMatrixDimMatch(a *Matrix, b Matrix) {
	ma, na := a.Dims()
	mb, nb := b.Dims()

	if ma != mb && na != nb {
		panic("matrix.DimsMustMatch")
	}
}

// Printers
func (a *Matrix) ToStringPlain() string {
	s := ""
	r := a.rows()
	for _, row := range r {
		sRow := arrayToString(row, ", ")
		s += "\n" + sRow
	}
	return s
}

func (a *Matrix) ToString() string {
	s := "a = [ "

	r := a.rows()
	for _, row := range r {
		sRow := arrayToString(row, ", ")
		s += sRow + "\n" + "      "
	}

	return string([]rune(s)[0:len(s)-7]) + " ]"
}

func arrayToString(a []float64, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
