// Common matrix operations
package main

import "fmt"

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

// Diagonal creates a n-by-n matrix with a Vector of n elements spread along the
// diagonal entries of the matrix
func Diagonal(v Vector) Matrix {
	mat := Zeros(len(v), len(v))
	iterateRows(&mat, func(i, j int) {
		if i == j {
			mat.set(i, j, v.At(i))
		}
	})
	return mat
}

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
func (a *Matrix) add(b Matrix) {
	panicMatrixDimMatch(a, b)
	a.el.add(b.el)
}

// Subtract Matrix b from Matrix a
// Dimensions of a and b must match
func (a *Matrix) subtract(b Matrix) {
	panicMatrixDimMatch(a, b)
	a.el.subtract(b.el)
}

// Multiply Matrix a by Matrix b
// Inner dimensions of a and b must match
func (a *Matrix) multiply(b Matrix) {
	m, n := a.Dims()
	new := Zeros(m, n)
	rows, columns := a.rows(), b.columns()
	iterateRows(a, func(i int, j int) {
		d := Dot(rows[i], columns[j])
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
		rows[i].scalarMult(1.0 / rows[i].At(i))

		// clear all other rows that are not of row i
		for j := 0; j < len(columns[i]); j++ {
			if j != i {
				clearRow(rows[j], rows[i], i)
			}

		}
		// break if the last row is zero to avoid inconsistencies
		zvec := ZeroVec(len(columns[i]))
		if Equal(rows[len(rows)-1], zvec) {
			break
		}
	}
	fmt.Printf("rows: %v\n", rows)

	reduced := Mat(rows...)

	return reduced

}

// subtract  row2 from row1 until index column of row1 is 0
func clearRow(row1, row2 Vector, column int) {
	cp := Vec(row2...)
	row1c := row1[column]
	// row2c := row2[column]
	ratio := row1c
	cp.scalarMult(ratio)
	row1.subtract(cp)
}

// swap two rows in the []Vector slice
func swapRows(rows []Vector, i1, i2 int) {
	row1 := rows[i1]
	row2 := rows[i2]

	rows[i1] = row2
	rows[i2] = row1
}

// return the index of the row with maximum value at column position
func maxColumn(rows []Vector, column int) int {
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

/**
 * Method:
 *  - find row with maximum entry in first column
 *  - swap this row with R1
 *  - divide all elements in R1 by the value of col 0 in R1
 *  - subtract R1 until col 0 or R2 is empty
 *
 *
 *
 *
 *
 */
