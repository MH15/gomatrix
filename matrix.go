// Matrix type
package main

// Matrix is the basic matrix interface type.
type Matrix struct {
	el     Vector
	stride int
	class  MatrixClass
}

// MatrixClass is a future-proof field to use when we add logical functionality
type MatrixClass int

const (
	// Logical true (1) or false (0)
	Logical MatrixClass = iota
	// FloatingPoint in float64 format
	FloatingPoint
)

// Mat creates a matrix (nxm) of equal sized vector rows
func Mat(v ...Vector) Matrix {
	mat := Matrix{}
	if EqualSize(v...) {
		// row and col sizes
		mat = Zeros(len(v), len(v[0]))
		// flatten v into the matrix
		iterateRows(&mat, func(i int, j int) {
			mat.set(i, j, v[i].At(j))
		})
	}
	return mat
}

// Zeros creates a zero-filled matrix
func Zeros(m, n int) Matrix {
	if notPositive(m, n) {
		panic("matrix.dimensionsMustbePositive")
	}
	return Matrix{el: make(Vector, m*n), stride: n, class: FloatingPoint}
}

// Ones creates a one-filled matrix
func Ones(m, n int) Matrix {
	mat := Zeros(m, n)
	iterateRows(&mat, func(i, j int) {
		mat.set(i, j, 1)
	})
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
func (a *Matrix) rows() []Vector {
	r := make([]Vector, len(a.el)/a.stride)
	i := 0
	for i < len(a.el) {
		r[i/a.stride] = Vec(a.el[i : i+a.stride]...)
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

func (a *Matrix) toStringPlain() string {
	s := ""
	r := a.rows()
	for _, row := range r {
		sRow := arrayToString(row, ", ")
		s += "\n" + sRow
	}
	return s
}

func (a *Matrix) toString() string {
	s := "a = [ "

	r := a.rows()
	for _, row := range r {
		sRow := arrayToString(row, ", ")
		s += sRow + "\n" + "      "
	}

	return string([]rune(s)[0:len(s)-7]) + " ]"
}

/// PANICS
func panicMatrixSize(a *Matrix, i, j int) {
	m, n := a.Dims()
	if i > m || j > n {
		panic(ErrIndexOutOfRange)
	}
}

func panicMatrixDimMatch(a *Matrix, b Matrix) {
	ma, na := a.Dims()
	mb, nb := b.Dims()

	if ma != mb && na != nb {
		panic("matrix.DimsMustMatch")
	}
}
