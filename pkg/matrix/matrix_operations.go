// Package matrix operations
package matrix

import "github.com/MH15/gomatrix/pkg/vector"

// Transpose Matrix a
func (a *Matrix) transpose() {
	m, n := a.Dims()
	new := Zeros(n, m)
	iterateRows(a, func(i int, j int) {
		new.set(j, i, a.at(i, j))
	})
	a.el = new.el
	a.stride = new.stride
}

// Add Matrix b to Matrix a
// Dimensions of a and b must match
func (a *Matrix) Add(b Matrix) {
	panicMatrixDimMatch(a, b)
	a.el.Add(b.el)
}

// Subtract Matrix b from Matrix a
// Dimensions of a and b must match
func (a *Matrix) Subtract(b Matrix) {
	panicMatrixDimMatch(a, b)
	a.el.Subtract(b.el)
}

// Multiply Matrix a by Matrix b
// Inner dimensions of a and b must match
func (a *Matrix) Multiply(b Matrix) {
	m, n := a.Dims()
	new := Zeros(m, n)
	rows, columns := a.rows(), b.columns()
	iterateRows(a, func(i int, j int) {
		d := vector.Dot(rows[i], columns[j])
		new.set(i, j, d)
	})
	a.el = new.el
	a.stride = new.stride
}

// RREF performs Gauss-Jordan Elimination to return the row-reduced echelon form
// while maintining the values in the original matrix.
func RREF(a Matrix) Matrix {
	rows := a.rows()
	columns := a.rows()
	// find the largest row, based on the column
	largest := maxColumn(rows, 0)
	// swap the largest row with the top (or top + i) row
	swapRows(rows, 0, largest)
	for i := 0; i < len(columns); i++ {
		// scale the row to lead with 1
		rows[i].ScalarMult(1.0 / rows[i].At(i))
		// clear all other rows that are not of row i
		for j := 0; j < len(columns[i]); j++ {
			if j != i {
				clearRow(rows[j], rows[i], i)
			}
		}
		// break if the last row is zero to avoid inconsistencies
		if vector.Equal(rows[len(rows)-1], vector.ZeroVec(len(columns[i]))) {
			break
		}
	}
	return Mat(rows...)
}

// subtract  row2 from row1 until index column of row1 is 0
func clearRow(row1, row2 vector.Vector, column int) {
	cp := vector.Vec(row2...)
	row1c := row1[column]
	// row2c := row2[column]
	ratio := row1c
	cp.ScalarMult(ratio)
	row1.Subtract(cp)
}

// swap two rows in the []Vector slice
func swapRows(rows []vector.Vector, i1, i2 int) {
	row1 := rows[i1]
	row2 := rows[i2]

	rows[i1] = row2
	rows[i2] = row1
}

// return the index of the row with maximum value at column position
func maxColumn(rows []vector.Vector, column int) int {
	index := 0
	max := rows[0][column]
	for i, row := range rows {
		if row[column] > max {
			max = row[column]
			index = i
		}
	}
	return index
}
